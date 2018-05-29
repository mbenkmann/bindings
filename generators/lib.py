#!/usr/bin/python3

# Copyright (c) 2018 Matthias S. Benkmann
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

from bs4 import BeautifulSoup
import textwrap
import fnmatch
import urllib.request
import time
import os

# Set of prefixes that are removed from C names to produce Go names.
prefixes = []

# The parsed input file data everything operates on.
soup = BeautifulSoup("", "xml")

# List of strings that will be "\n".join()ed to form the output Go file.
out = []

# A stack of indentation added to lines output. Each line is indented by "".join(indent)
indent = []

# List of objects with a typeinfo() method. The global lib.typeinfo() function
# goes through this list and calls each heuristic to gather typeinfo.
typeinfo_heuristics = []

# Maps a C type name to a gocast value as for typeinfo().
custom_gocast = {}

# Maps a C type name to a ccast value as for typeinfo().
custom_ccast = {}

# Maps a prefix in a C type name to a Go package name, for cross-package references.
cprefix_to_go_package = {}

# List of names for which not to generate wrappers.
blacklist = frozenset()

# These words are ignored when they occur in a type
ignored_type_elements = frozenset()

# Maps a Go type name to a list of name components that will be removed from
# generated function names that use this type as a receiver.
# All lists must be sorted from longest to shortest name and must include
# the Go type name (if it is to be considered).
# Types whose names are not in this map default to a list consisting of just
# their name.
receiver_aliases = {}

# Directory of doxygen xml output files.
doxyxml = "./xml"

# additional boilerplate lines
boilerplate = set()

# Pointer arguments to C functions can be used for different things:
#  "out" Parameters: The storage pointed to is not read by the function and
#                    only used to transfer results to the caller.
#  "receiver" Parameters: The storage pointed to is operated on by the function
#                         in some way so that it makes sense to treat the object
#                         as receiver in Go.
# This map maps a C type name to another map that describes how pointers to this
# type that are used as function arguments are to be handled.
#  "default": May be "in", "out","inout" or "receiver" and specifies the default treatment,
#  "receiver": An iterable of function names of functions for which the
#              pointers should be treated as "receiver" arguments instead of
#              the default.
#              It's even possible to specify an individual argument of a function,
#              e.g. "SDL_OpenAudio.obtained"
#  "out": An iterable of function names of functions for which the
#         pointers should be treated as "out" arguments instead of
#         the default.
#  "in": Like "out", but for "in"
#  "inout": Like "out", but for "inout"
#  "by-value": True => Pass arg by value in Go func even if C signature has it by pointer.
#  "by-ptr": iterable of C function names for which by-value:True is ignored.
pointer_arg_treatment = {}

# Names of functions whose returned C strings need to be freed by the caller.
free_strings = frozenset()

# Mapping for individual function parameter types and structure member types, e.g.
# { "SDL_CreateWindow.flags": "WindowFlags" }
# The argument name is a glob pattern (see fnmatch module).
gotype_override = {}

# Maps a C enum name to the integer type to be used on the Go side (if not "int").
enum_types = {}

# A mapping from a glob pattern to a URL which may contain the string "$0" that
# will be replaced with the complete name that matches the pattern.
# This map is used to include links to external URLs in documentation comments.
# ATTENTION: Links are simply copied into the documentation unless the
# environment variable CHECK_EXTERNAL_LINKS is set.
external_link_urls = {}


def generate_bindings():
    '''
    Performs standard binding generation, filling lib.out with the result.
    '''
    describe(soup.doxygen.compounddef)

    structs()
    unions()

    for section in soup.doxygen.compounddef.find_all("sectiondef"):
        out.append("")
        describe(section)
        define2const(section)
        enum2const(section)
        aliastypedefs(section)
        simpletypedefs(section)
        wrapfunctions(section)


class BaseTypeinfo(object):
    '''
    A simple translation from source to destination types based on a dict.
    Also fills in basic info to make sure the result dict is complete.
    This heuristic is usually the first in typeinfo_heuristics.
    '''

    def __init__(self, mapping):
        '''
        Takes a dict from str->str that translates a source type into a destination type.
        '''
        self.mapping = mapping

    def typeinfo(self, funcname, argidx, argname, argtype, argtypeargs, result):
        '''
        Fills in the result (dict). See global typeinfo() function
        for details on the other arguments.
        '''
        if argname == "":
            if argidx == 0:
                result["name"] = "retval"
            else:
                result["name"] = "arg%s" % argidx
        else:
            result["name"] = fix(argname, True)

        result["allocarg"] = result["name"]
        result["alloc"] = ""
        result["dealloc"] = ""

        argtype_stars = ""
        while argtype[-1] == "*":
            argtype_stars += "*"
            argtype = argtype[:-1]

        result["struct"] = False
        const = ""
        at = argtype.split()
        if "const" in at:
            at.remove("const")
            const = "const "
        if "struct" in at:
            at.remove("struct")
            result["struct"] = True
        argtype = "".join(at)

        if argtype == "unsignedint":
            argtype = "uint"

        result["array"] = []
        if argtypeargs != "":
            argtypeargs = argtypeargs.replace("[", "").replace("]", " ")
            for a in argtypeargs.split():
                if a[0].isalpha():
                    result["array"].append(fix(a))
                else:
                    result["array"].append(a)

        arr = ""
        if len(result["array"]) > 0:
            arr = "[" + "][".join(result["array"]) + "]"

        try:
            gotype = self.mapping[const + argtype + argtype_stars + arr]
        except:
            try:
                gotype = self.mapping[argtype + argtype_stars + arr]
            except:
                try:
                    gotype = argtype_stars + arr + self.mapping[argtype]
                except:
                    gotype = argtype_stars + arr + fix(argtype)

        gotype = apply_gotype_override(funcname, argname, gotype)

        package = ""
        for prefix, pkg in cprefix_to_go_package.items():
            if argtype.startswith(prefix):
                package = pkg + "."
                break

        i = min((x for x in range(len(gotype)) if gotype[x].isalpha()), default=0)
        gotype = gotype[0:i] + package + gotype[i:]
        result["gotype"] = gotype
        result["gocastend"] = ""

        if argtype in custom_gocast:
            result["gocast"] = custom_gocast[argtype]
        elif gotype == "string":
            if funcname in free_strings:
                result["gocast"] = "freeGoString"
            else:
                result["gocast"] = "C.GoString"
            if argidx != 0:
                result["allocarg"] = "tmp_" + result["allocarg"]
                result["alloc"] = "%s := C.CString(%s); defer C.free(unsafe.Pointer(%s))" % (
                    result["allocarg"], result["name"], result["allocarg"])
        else:
            result["gocast"] = gotype
            if result["struct"]:
                if argtype in pointer_arg_treatment and "by-value" in pointer_arg_treatment[argtype] and pointer_arg_treatment[argtype]["by-value"]:
                    if "by-ptr" in pointer_arg_treatment[argtype] and funcname in pointer_arg_treatment[argtype]["by-ptr"]:
                        pass  # by-ptr is an exception to by-value:True
                    else:
                        gotype = gotype.lstrip("*")
                        result["gotype"] = gotype

                result["gocast"] = "fromC2" + result["gocast"].lstrip("*")
                if package != "":
                    result["gocast"] = result["gocast"].replace(package, "", 1)
                if argtype_stars != "":
                    result["gocast"] = result["gocast"] + "(*"
                    result["gocastend"] = ")"
            elif "*" in result["gocast"]:
                result["gocast"] = "(%s)(unsafe.Pointer" % result["gocast"]
                result["gocastend"] = ")"
            elif "[" in result["gocast"]:
                result["gocast"] = "*(*%s)(unsafe.Pointer(&" % result["gocast"]
                result["gocastend"] = "))"

        result["ctype"] = "%s%sC.%s" % (argtype_stars, arr, fix(argtype, True))
        if result["ctype"] == "*C.void":
            result["ctype"] = "unsafe.Pointer"

        result["ccastend"] = ""
        if argtype in custom_ccast:
            result["ccast"] = custom_ccast[argtype]
        else:
            if gotype == "string" and result["alloc"] == "":
                result["ccast"] = "C.CString"
            else:
                result["ccast"] = result["ctype"]

            if "[" in result["ccast"]:
                result["ccast"] = "(" + result["ccast"] + ")"
            elif "*" in result["ccast"]:
                if "[" in result["gotype"] or package != "":
                    result["ccast"] = "(%s)(unsafe.Pointer" % result["ctype"]
                    result["ccastend"] = ")"
                else:
                    result["ccast"] = "(" + result["ccast"] + ")"

        if len(result["array"]) > 0:
            result["ccast"] = "*(*%s)(unsafe.Pointer(&" % (result["ctype"], )
            result["ccastend"] = "))"

        result["cidx"] = argidx
        if argidx != 0 and (argtype_stars == "" or argtype == "void" or const != ""):
            treat = "in"
        else:
            treat = "out"

        if argidx != 0 and argtype in pointer_arg_treatment:
            if "default" in pointer_arg_treatment[argtype]:
                treat = pointer_arg_treatment[argtype]["default"]
            if "out" in pointer_arg_treatment[argtype] and funcname in pointer_arg_treatment[argtype]["out"]:
                treat = "out"
            if "in" in pointer_arg_treatment[argtype] and funcname in pointer_arg_treatment[argtype]["in"]:
                treat = "in"
            if "inout" in pointer_arg_treatment[argtype] and funcname in pointer_arg_treatment[argtype]["inout"]:
                treat = "inout"
            if "receiver" in pointer_arg_treatment[argtype] and funcname in pointer_arg_treatment[argtype]["receiver"]:
                treat = "receiver"
            if "out" in pointer_arg_treatment[argtype] and (
                    funcname + "." + argname) in pointer_arg_treatment[argtype]["out"]:
                treat = "out"
            if "in" in pointer_arg_treatment[argtype] and (
                    funcname + "." + argname) in pointer_arg_treatment[argtype]["in"]:
                treat = "in"
            if "inout" in pointer_arg_treatment[argtype] and (
                    funcname + "." + argname) in pointer_arg_treatment[argtype]["inout"]:
                treat = "inout"
            if "receiver" in pointer_arg_treatment[argtype] and (
                    funcname + "." + argname) in pointer_arg_treatment[argtype]["receiver"]:
                treat = "receiver"

        result["retval"] = treat == "out"

        if treat == "receiver":
            result["goidx"] = 0

        elif treat == "out":
            # +100 to make sure that out args have higher index than in args
            result["goidx"] = argidx + 100
            if argtype_stars != "" and argidx != 0:
                result["allocarg"] = "tmp_" + result["allocarg"]
                result[
                    "alloc"] = result["allocarg"] + " := new(" + result["ctype"].lstrip("*") + ")"
                result["dealloc"] = "%s = %s(%s)%s" % (result["name"], result["gocast"],
                                                       result["allocarg"], result["gocastend"])
                if need_temp_to_get_pointer(result):
                    result["dealloc"] = "tmp2_%s; %s = &tmp2_%s" % (result["dealloc"],
                                                                    result["name"], result["name"])
                    result["dealloc"] = result["dealloc"].replace("=", ":=", 1)

        elif treat.startswith("in"):
            result["goidx"] = argidx
            if result["struct"]:
                if result["gotype"].startswith("*"):
                    result["allocarg"] = "tmp_" + result["allocarg"]
                    result["alloc"] = "var %s %s; if %s != nil { x := toCFrom%s(*%s); %s = &x }" % (
                        result["allocarg"], result["ctype"], result["name"], gotype.lstrip("*"),
                        result["name"], result["allocarg"])
                    if treat.endswith("out"):
                        result["dealloc"] = "if %s != nil { *%s = %s(%s)%s }" % (
                            result["name"], result["name"], result["gocast"], result["allocarg"],
                            result["gocastend"])
                else:
                    result["allocarg"] = "tmp_" + result["allocarg"]
                    result["alloc"] = result["allocarg"] + " := toCFrom%s(%s)" % (
                        gotype.lstrip("*"), result["name"])
                    if argtype_stars != "":
                        result["allocarg"] = "&" + result["allocarg"]

        # Special case for pointer to primitive data type as out parameter
        if treat == "out" and result["gotype"][0] == "*" and result["gotype"].islower():
            result["gotype"] = result["gotype"].lstrip("*")
            new_go_cast = "deref_" + result["gotype"] + "_ptr"
            result["dealloc"] = result["dealloc"].replace(result["gocast"], new_go_cast)
            if result["gocastend"] != "":
                result["dealloc"] = result["dealloc"][0:-len(result["gocastend"])]
            result["gocast"] = new_go_cast
            result["gocastend"] = ""

        if "unsafe." in result["gotype"] or "unsafe" in result["alloc"] or "unsafe" in result["ccast"] or (
                treat == "out" and "unsafe" in result["gocast"]):
            boilerplate.add('import "unsafe"')


def apply_gotype_override(name, subname, gotype):
    '''
    If gotype_override contains an override for name.subname (or just name if
    subname==""), returns it; otherwise  returns gotype.
    '''
    if subname != "":
        name += "." + subname
    for pattern in gotype_override:
        if fnmatch.fnmatchcase(name, pattern):
            return gotype_override[pattern]

    return gotype


def typeinfo(funcname, argidx, argname, argtype, argtypeargs):
    '''
    In:
      funcname: Name of the function in whose signature the type to get info on occurs.
      argidx: Index in the function's argument list where the type occurs.
             argidx == 0 refers to the return type.
             argidx == 1 is the first argument inside the parentheses.
      argname: Name of the argument in the signature. Empty string if no name.
      argtype: Type name in the function signature. All whitespace must be removed,
               i.e. "char *" => "char*".
      argtypeargs: If the type is an array of length X this is "[X]"
    
    Out: A dictionary containing the following data
      name: (str) Name to use for the argument in the generated Go code.
            Usually the same as argname, unless argname == "".
      struct: (bool) True if the type is a struct with known fields (i.e. not
                     a library-internal struct).
      alloc: (str) If not "", this is a statement that needs to be output
                   before the function call to allocate a variable.
      allocarg: (str) If alloc=="" this is the same as name, otherwise it is
                the name of the allocated variable, potentially prefixed with "&".
                Should be used as the function call argument in the C function call.
      dealloc: (str) If not "", this is a statement to execute after the function
               call. It will usually correspond to alloc in some way.
      gotype: (str) Go type to use for the argument.
      gocast: (str) Cast to use for converting C to Go. Often the same as gotype,
                    but might be something like "C.GoString"
      gocastend: (str) A string with additional closing ")" required after
                 the closing ")" that every cast has.
      ctype: (str) C type of the argument within Go, e.g. "C.char".
      ccast: (str) Cast to use for converting Go to C. Often the same as ctype,
                   but could include parentheses, e.g. "(*C.char)". Code using
                   ccast never has to add parentheses around ccast.
      ccastend: (str) A string with additional closing ")" required after
                   the closing ")" that every cast has.
      array: (list of str) If the argument type is an array, this
                   is a list containing the array dimensions
                   (e.g. "char[2][3]" => ["2","3"])
      retval: (bool) True => this argument is a return value of the Go function
      goidx: (int) 0 => on the Go side the argument is a method receiver.
                   1 => (if not retval) 1st argument to the Go function.
                   ...
                   n => (if retval) 1st return value
                   ...
                   NOTE: typeinfo() cannot know about actual argument indexes,
                   so this is just filled in from argidx and has to be fixed up
                   later once info on all arguments has been collected.
      cidx: (int) 0 => on the C side the argument is the return value
                  1 => on the C side the argument is the 1st function parameter.
                  ...
                  NOTE: typeinfo() cannot know about actual argument indexes,
                  so this is just filled in from argidx and has to be fixed up
                  later once info on all arguments has been collected.

    '''
    result = {}
    for heuristic in typeinfo_heuristics:
        heuristic.typeinfo(funcname, argidx, argname, argtype, argtypeargs, result)

    return result


def fixup_params(params):
    '''
    Takes a list of dicts as returned from typeinfo() and
    * fixes the included goidx numbers.
    * sorts params according to goidx
    * if there is no receiver, inserts None as params[0]
    Returns 
      num_recv  number of receivers (Could be >1 if params is incorrect)
      num_args  number of pure input arguments
      num_ret   number of return values (C retval and ptr params)
    '''
    if len(params) == 0:
        return 0, 0, 0

    params.sort(key=lambda p: p["goidx"])

    if params[0]["goidx"] > 0:
        params.insert(0, None)

    num_recv = 0
    num_args = 0
    num_ret = 0

    for p in params:
        if p is not None:
            if p["goidx"] == 0:
                num_recv += 1
            else:
                p["goidx"] = num_args + num_ret + 1

                if p["retval"]:
                    num_ret += 1
                else:
                    num_args += 1

    return num_recv, num_args, num_ret


# The function slices_instead_of_pointers() deals with the following cases:
#
# *Struct+count => []Struct
# Patterns to recognize: 1) "const Struct*" followed by "int count"
#                           (also recognize "len" and "length")
#                        2) "const Struct*" named "foo" followed by "int nfoo"
#                           (also recognize "numfoo" and "num_foo")
#                  In both cases, allow 1 optional int between the 2 arguments
#SDL_RenderDrawPoints(SDL_Renderer* renderer, const SDL_Point * points,                 int count);
#SDL_RenderDrawLines(SDL_Renderer * renderer, const SDL_Point * points,                 int count);
#SDL_RenderDrawRects(SDL_Renderer * renderer, const SDL_Rect  * rects,                  int count);
#SDL_RenderFillRects(SDL_Renderer * renderer, const SDL_Rect  * rects,                  int count);
#SDL_FillRects(SDL_Surface * dst,             const SDL_Rect  * rects,                  int count, Uint32 color);
#SDL_SetPaletteColors(SDL_Palette * palette,  const SDL_Color * colors, int firstcolor, int ncolors);

# *byte+count => []byte
# Patterns to recognize: [const] void* followed by Uint32/int named "len" (also recognize "length")
#SDL_QueueAudio(SDL_AudioDeviceID    dev,    const void  *data,                       Uint32 len);
#SDL_AudioStreamPut(SDL_AudioStream *stream, const void  *buf,                           int len);
#SDL_DequeueAudio(SDL_AudioDeviceID  dev,          void  *data,                       Uint32 len);
#SDL_AudioStreamGet(SDL_AudioStream *stream,       void  *buf,                           int len);

# Variant of the previous case with an additional *byte and count that applies to both
# Patterns to recognize: [[const] Uint8*] [const] Uint8* followed by Uint32/int named "len"
#                        (also recognize "length")
#                        allow 1 optional non-pointer before "len"
#SDL_MixAudio(                Uint8 * dst,   const Uint8 *src,                        Uint32 len, int volume);
#SDL_MixAudioFormat(          Uint8 * dst,   const Uint8 *src, SDL_AudioFormat format,Uint32 len, int volume);

# *byte WITHOUT count => []byte + handwritten checkParametersFor...() function
# Patterns to recognize: [const] void* followed by an "int pitch" (with potentially other parameters in between)
#SDL_CreateRGBSurfaceWithFormatFrom(                                void *pixels, int width, int height, int depth, int pitch, Uint32 format);
#SDL_CreateRGBSurfaceFrom(                                          void *pixels, int width, int height, int depth, int pitch, Uint32 Rmask,Uint32 Gmask, Uint32 Bmask, Uint32 Amask);
#SDL_UpdateTexture(SDL_Texture* texture,const SDL_Rect* rect, const void *pixels, int pitch);

# Variant of the previous case with multiple *byte
# Patterns to recognize: [const] Uint8* named "foo" followed by an int whose name
#                         ends with "pitch" and starts with a prefix of "foo"
#SDL_UpdateYUVTexture(SDL_Texture* texture,const SDL_Rect* rect, const Uint8 *Yplane, int Ypitch, const Uint8 *Uplane, int Upitch, const Uint8 *Vplane, int Vpitch);
#
#


def slices_instead_of_pointers(funcname, params):
    '''
    Takes a function name and its parameters as a list of dicts as returned from
    typeinfo() that has already been through fixup_params(); and checks if there
    are any pointer arguments that should be converted to slices instead.
    These arguments are converted to slices and if necessary, accompanying
    length/count arguments are marked by adding a mapping "nogo":True, so that
    they can be omitted from the Go wrapper's signature.
    '''
    fix = []
    for i in range(len(params)):
        p = params[i]
        if p is None or p["retval"]:
            continue
        if p["ctype"].startswith("*") or p["ctype"] == "unsafe.Pointer":
            if p["struct"]:
                # if there either is no next argument or it is a pointer, then
                # it doesn't match the pattern
                if i + 1 >= len(params) or (params[i + 1]["ctype"].startswith("*")
                                            or params[i + 1]["ctype"] == "unsafe.Pointer"):
                    continue
                for k in (1, 2):
                    if i + k < len(params) and params[i + k]["name"] in ("len", "length", "count",
                                                                         "n" + p["name"],
                                                                         "num" + p["name"],
                                                                         "num_" + p["name"]):
                        fix.append((True, i, i + k))
                        break
            elif p["ctype"] == "unsafe.Pointer" or isbyteptr(p["gotype"]):
                one_non_pointer_allowed = True
                for k in range(1, len(params) - i):
                    if params[i + k]["name"] in ("len", "length",
                                                 "count") and params[i + k]["gotype"] in ("uint32",
                                                                                          "int"):
                        fix.append((False, i, i + k))
                        break
                    elif params[i + k]["ctype"] == "unsafe.Pointer" or isbyteptr(
                            params[i + k]["gotype"]):
                        continue  # multiple subsequent pointers are allowed in this pattern
                    else:
                        if one_non_pointer_allowed and not (
                                params[i + k]["ctype"].startswith("*")
                                or params[i + k]["ctype"] == "unsafe.Pointer"):
                            one_non_pointer_allowed = False
                            continue
                        else:
                            break  # other stuff breaks the pattern

                for k in range(1, len(params) - i):
                    if params[i + k]["gotype"] in ("uint32", "int"):
                        n = params[i + k]["name"]
                        if n.endswith("pitch") and p["name"].startswith(n[0:len(n) - 5].strip("_")):
                            fix.append((False, i, -1))
                            break

    have_output_check = False

    for struct, idx, lenidx in fix:
        raw_gotype = params[idx]["gotype"].lstrip("*")

        p = params[idx]
        p["allocarg"] = "tmp_" + p["name"]
        alloc = []
        if lenidx < 0 and not have_output_check:
            have_output_check = True
            alloc.append("checkParametersFor" + funcname + "(")
            first = True
            for parm in params:
                if parm is None or parm["retval"]:
                    continue
                if not first:
                    alloc[-1] += ", "
                first = False
                alloc[-1] += parm["name"]
            alloc[-1] += ")"
        alloc.append("var %s %s" % (p["allocarg"], p["ctype"]))
        alloc.append("if len(%s) > 0 {" % p["name"])
        if struct:
            p["gotype"] = "[]" + raw_gotype
            alloc.append("    sl_%s := make([]%s, len(%s))" % (p["allocarg"],
                                                               p["ctype"].lstrip("*"), p["name"]))
            alloc.append("    for i := range %s {" % p["name"])
            alloc.append("        sl_%s[i] = toCFrom%s(%s[i])" % (p["allocarg"], raw_gotype,
                                                                  p["name"]))
            alloc.append("    }")
            alloc.append("    %s = &(sl_%s[0])" % (p["allocarg"], p["allocarg"]))
        else:
            p["gotype"] = "[]byte"
            alloc.append("    %s = (%s)(unsafe.Pointer(&(%s[0])))" % (p["allocarg"], p["ctype"],
                                                                      p["name"]))
        alloc.append("}")

        p["alloc"] = "\n    ".join(alloc)
        p["dealloc"] = ""
        p["ccast"] = ""
        p["ccastend"] = ""

        if lenidx >= 0:
            p = params[lenidx]
            if "nogo" not in p:
                p["nogo"] = True
                p["dealloc"] = ""
                p["allocarg"] = "tmp_" + p["name"]
                p["alloc"] = "%s := len(%s)" % (p["allocarg"], params[idx]["name"])
            else:
                p["alloc"] += "; if len(%s) < %s { %s = len(%s) }" % (params[idx]["name"],
                                                                      p["allocarg"], p["allocarg"],
                                                                      params[idx]["name"])


def isbyteptr(gotype):
    return gotype.startswith("*") and (gotype.endswith("byte") or gotype.endswith("uint8"))


def additional_boilerplate(idx):
    '''
    When called after the whole file is processed, this function will insert
    import statements and other boilerplate at lib.out[idx], whose necessity
    is only known after processing the complete file.
    '''
    for add in boilerplate:
        out.insert(idx, add)


def external_link(name):
    '''
    If external_link_urls contains a pattern that matches name, this function
    tests the resulting URL for a 404 and if it doesn't give a 404, a comment
    with the URL is appended to out.
    '''
    append_link = False
    for pattern in external_link_urls:
        if fnmatch.fnmatchcase(name, pattern):
            url = external_link_urls[pattern].replace("$0", name)
            if "CHECK_EXTERNAL_LINKS" in os.environ:
                time.sleep(1000)  # throttling to avoid triggering DDOS protection
                request = urllib.request.Request(url, method='HEAD')
                try:
                    with urllib.request.urlopen(request) as r:
                        if r.getcode() != 404:
                            append_link = True
                except:
                    pass
            else:
                append_link = True

    if append_link:
        pre = indentation() + " // ↪ "
        out.append(pre + url)


def describe(tag):
    ''' Converts tag's description into a Go comment. '''
    global out
    paras = []
    for descname in ("header", "briefdescription", "detaileddescription", "inbodydescription",
                     "description"):
        for desc in tag.find_all(descname, recursive=False):
            if desc.find("para") is None:
                s = str(desc.string)
                if s != "" and not s.isspace():
                    paras.append(s)

            # Special handling for <simplesect kind="foo"> because
            # it has a <para> inside, so it's like this:
            # <para><simplesect><para>Text</para></simplesect>More Text</para>
            # We transform this to:
            # <para>Foo: Text</para><para>More Text</para>
            for retsec in desc.find_all("simplesect"):
                kind = retsec["kind"]
                if kind not in ("return", "warning", "note", "see", "par"):
                    continue
                retsec.para.unwrap()
                retsec.name = "para"
                del retsec["kind"]
                p = retsec.parent
                retsec = retsec.extract()
                if kind == "return":
                    retsec.insert(0, "Returns: ")
                elif kind == "warning":
                    retsec.insert(0, "Warning: ")
                elif kind == "note":
                    retsec.insert(0, "Note: ")
                elif kind == "see":
                    retsec.insert(0, "See also: ")
                elif kind == "par":
                    retsec.title.unwrap()
                p.insert_before(retsec)

            for para in desc.find_all("para", recursive=False):
                p = []

                # First we remove all tags where we just use the string contents.
                # We do this before the loop below so that the line wrapping in the
                # NavigableString case works properly.
                for x in para.children:
                    if x.name in ("ref", "computeroutput", "bold", "ulink", "emphasis"):
                        x.unwrap()
                    elif x.name == "ndash":
                        x.replace_with("--")

                # Now merge adjacent NavigableStrings
                kids = list(para.children)  # need to instantiate list, because we use extract()
                for x in kids:
                    if x.name is None:
                        sib = x.previous_sibling
                        if sib is not None and sib.name is None:
                            x.replace_with(str(sib) + str(x.string))
                            sib.extract()

                for x in para.children:
                    if x.name is None:  # NavigableString
                        s = str(x).strip()
                        wrapped = textwrap.wrap(s, 70)
                        if len(s) > 0 and len(wrapped) > 0:
                            if s[-1].isspace():
                                wrapped[-1] += " "
                            if s[0].isspace():
                                wrapped[0] = " " + wrapped[0]

                        p.extend(wrapped)
                    else:  # Tag
                        # translate itemizedlist into verbatim
                        # <para>List:<itemizedlist>
                        # <listitem><para>Item 1</para></listitem>
                        # <listitem><para>Item 2</para></listitem>
                        # </itemizedlist></para>
                        # NOTE: Each <listitem> may contain <ref>, <computeroutput> and stuff like that.
                        if x.name == "itemizedlist":
                            x.name = "verbatim"

                            for item in x("listitem"):
                                txt = "".join(item.strings)
                                txt = " ".join(txt.split())  # reduce ws sequences to 1 space
                                wrapped = textwrap.wrap(txt, 70)
                                if len(wrapped) > 0:
                                    wrapped[0] = "- " + wrapped[0]
                                for i in range(1, len(wrapped)):
                                    wrapped[i] = "  " + wrapped[i]
                                txt = "\n".join(wrapped)
                                item.replace_with(txt + "\n")

                        # translate orderedlist into verbatim
                        # <para>List:<orderedlist>
                        # <listitem><para>Item 1</para></listitem>
                        # <listitem><para>Item 2</para></listitem>
                        # </orderedlist></para>
                        # NOTE: Each <listitem> may contain <ref>, <computeroutput> and stuff like that.
                        if x.name == "orderedlist":
                            x.name = "verbatim"

                            idx = 0
                            for item in x("listitem"):
                                idx += 1
                                txt = "".join(item.strings)
                                txt = " ".join(txt.split())  # reduce ws sequences to 1 space
                                wrapped = textwrap.wrap(txt, 65)
                                if len(wrapped) > 0:
                                    wrapped[0] = "%s. %s" % (str(idx).rjust(2), wrapped[0])
                                for i in range(1, len(wrapped)):
                                    wrapped[i] = "    " + wrapped[i]
                                txt = "\n".join(wrapped)
                                item.replace_with(txt + "\n\n")

                        # translate parameterlist into verbatim
                        # <para>Params:<parameterlist>
                        # <parameteritem>
                        # <parameternamelist><parametername>param1</parametername>
                        # <parametername>param2</parametername>
                        # </parameternamelist>
                        # <parameterdescription><para>Text</para></parameterdescription>
                        # </parameteritem>
                        # <parameteritem>...</parameteritem>
                        # </parameterlist>
                        if x.name == "parameterlist" and x["kind"] == "param":
                            x.name = "verbatim"
                            del x["kind"]

                            for item in x("parameteritem"):
                                pnamelist = ", ".join(item.parameternamelist.stripped_strings)
                                txt = "".join(item.parameterdescription.strings)
                                txt = " ".join(txt.split())  # reduce ws sequences to 1 space
                                wrapped = textwrap.wrap(txt, 70)
                                for i in range(0, len(wrapped)):
                                    wrapped[i] = "  " + wrapped[i]
                                txt = "\n".join(wrapped)
                                item.replace_with(pnamelist + "\n" + txt + "\n")

                        # translate programlisting into verbatim
                        # <para>Code:<programlisting>
                        # <codeline><highlight>word<sp/>word<sp/></highlight></codeline>
                        # <codeline><highlight>word<sp/>word<sp/></highlight></codeline>
                        # </programlisting>
                        if x.name == "programlisting":
                            x.name = "verbatim"

                            for item in x("codeline"):
                                for sp in item("sp"):  # replace non-breaking space tag with space
                                    sp.replace_with(" ")
                                txt = "".join(item.strings)
                                item.replace_with(txt)

                        # translate table into verbatim
                        # <table rows="5" cols="3">
                        # <row><entry thead="no"><para>...</para></entry><entry thead="no"><para>...</para></entry><entry thead="no"><para>...</para></entry></row>
                        # <row><entry thead="no"><para>...</para></entry><entry thead="no"><para>...</para></entry><entry thead="no"><para>...</para></entry></row>
                        # <row><entry thead="no"><para>...</para></entry><entry thead="no"><para>...</para></entry><entry thead="no"><para>...</para></entry></row>
                        # <row><entry thead="no"><para>...</para></entry><entry thead="no"><para>...</para></entry><entry thead="no"><para>...</para></entry></row>
                        # <row><entry thead="no"><para>...</para></entry><entry thead="no"><para>...</para></entry><entry thead="no"><para>...</para></entry></row>
                        # </table>
                        if x.name == "table":
                            x.name = "verbatim"

                            tab = []
                            width = []
                            for row in x("row"):
                                tab.append([])
                                i = 0
                                for entry in row("entry"):
                                    txt = "".join(entry.strings)
                                    txt = " ".join(txt.split())  # reduce ws sequences to 1 space
                                    tab[-1].append(txt)
                                    if i >= len(width):
                                        width.append(0)
                                    width[i] = max(width[i], len(txt))
                                    i += 1

                            k = 0
                            for row in x("row"):
                                line = tab[k]
                                k += 1
                                s = []
                                for i in range(len(line)):
                                    s.append(line[i].ljust(width[i]))
                                    i += 1
                                row.replace_with(" | ".join(s))

                        if x.name == "verbatim":
                            prevline = ""
                            if len(p) > 0:
                                prevline = p[-1]
                            else:
                                if len(paras) > 0:
                                    prevline = paras[-1]

                            prevline = prevline.rstrip("\n")
                            i = prevline.rfind("\n")
                            if i >= 0:
                                prevline = prevline[i + 1:]

                            if prevline == "":
                                prevline = "Notes:"
                                p.append(prevline)

                            i = 0
                            while i < len(prevline) and prevline[i].isspace():
                                i += 1

                            # for godoc to create a <pre> section we need to add more indent than
                            # the previous line
                            additional_indent = prevline[0:i] + "  "
                            for line in "".join(x.strings).splitlines():
                                p.append(additional_indent + line)

                        else:
                            raise NotImplementedError("Unknown tag " + str(x.name))

                paras.append("\n".join(p))

    description = "\n\n".join(paras)
    if description != "" and not description.isspace():
        pre = indentation() + " // "
        out.append(pre + (("\n" + pre).join(description.splitlines())))


def fix(name, keep_name=False):
    '''
    If keep_name == False, removes prefixes from name if present and makes the
    first character uppercase if it's a letter.
    If keep_name == True, prefixes the name with an underscore if it conflicts
    with a Go reserved word but otherwise keeps name untouched.
    '''
    if keep_name:
        if name in ("break", "default", "func", "interface", "select", "case", "defer", "go", "map",
                    "struct", "chan", "else", "goto", "package", "switch", "const", "fallthrough",
                    "if", "range", "type", "continue", "for", "import", "return", "var"):
            name = "_" + name
        return name

    for p in prefixes:
        if name.startswith(p):
            name = name[len(p):]

    return name[0].upper() + name[1:]


def indentation():
    return "".join(indent)


def push_indent():
    global indent
    indent.append("    ")


def pop_indent():
    global indent
    indent.pop()


def structs():
    ''' Converts C structs to Go structs. '''
    global out
    global indent

    for struct in soup.find_all("compounddef", kind="struct"):
        name = str(struct.compoundname.string)
        if name in blacklist:
            continue

        refid = str(struct["id"])

        out.append("")  # empty line to separate entries
        describe(struct)
        external_link(name)
        goname = fix(name)
        out.append("%stype %s struct {" % (indentation(), goname))
        push_indent()

        fields = list(struct.find_all("memberdef", kind="variable", prot="public"))
        fields.sort(key=lambda x: int(x.location["line"]))
        first = True
        for member in fields:
            if not first:
                out.append("")  # empty line to separate entries
            first = False
            typ = get_ctype(member.type)
            typargs = ""
            membname = str(member.find("name").string)
            if member.argsstring is not None:
                typargs = str("".join(member.argsstring.stripped_strings))
            ti = typeinfo(name, 0, membname, typ, typargs)
            describe(member)
            out.append("%s%s %s" % (indentation(), fix(membname), ti["gotype"]))

        pop_indent()
        out.append(indentation() + "}")

        # Now output a function that converts from the C struct to the Go struct
        fromCName = "fromC2" + goname
        if not fromCName in blacklist:
            out.append("")
            out.append("%sfunc %s(s C.%s) %s {" % (indentation(), fromCName, name, goname))
            push_indent()
            ti = typeinfo(name, 0, "", name, "")
            out.append("%sreturn %s" % (indentation(), go_copy_of(refid, "compound", ti, "s")))
            pop_indent()
            out.append(indentation() + "}")

        # Now output a function that converts from the Go struct to the C struct
        toCFromName = "toCFrom" + goname
        if not toCFromName in blacklist:
            out.append("")
            out.append("%sfunc %s(s %s) (d C.%s) {" % (indentation(), toCFromName, goname, name))
            push_indent()
            ti = typeinfo(name, 0, "", name, "")
            recursive_copy(refid, "compound", ti, "d", "s")
            out.append(indentation() + "return")
            pop_indent()
            out.append(indentation() + "}")


def unions():
    ''' Converts C unions to Go types with access methods. '''
    global out
    global indent

    for union in soup.find_all("compounddef", kind="union"):
        name = str(union.compoundname.string)
        if name in blacklist:
            continue

        out.append("")  # empty line to separate entries
        describe(union)
        external_link(name)
        goname = fix(name)
        out.append("%stype %s C.%s" % (indentation(), goname, name))

        fields = list(union.find_all("memberdef", kind="variable", prot="public"))
        fields.sort(key=lambda x: int(x.location["line"]))
        for member in fields:
            typ = get_ctype(member.type)
            typargs = ""
            if member.argsstring is not None:
                typargs = str("".join(member.argsstring.stripped_strings))
            membname = str(member.find("name").string)
            if name + "." + membname in blacklist:
                continue
            ti = typeinfo(name, 0, membname, typ, typargs)
            try:
                refid = str(member.type.ref["refid"])
                kindref = str(member.type.ref["kindref"])
            except:
                refid = ""
                kindref = "primitive"

            boilerplate.add('import "unsafe"')

            out.append("")  # empty line to separate entries

            # Getter
            describe(member)
            out.append("%sfunc (u *%s) %s() %s {" % (indentation(), goname, fix(membname),
                                                     ti["gotype"]))
            push_indent()
            out.append("%sp := (*%s)(unsafe.Pointer(u))" % (indentation(), ti["ctype"]))
            out.append("%sreturn %s" % (indentation(), go_copy_of(refid, kindref, ti, "*p")))
            pop_indent()
            out.append(indentation() + "}")

            # Setter
            settername = "Set" + fix(membname)
            if goname + "." + settername not in blacklist:
                describe(member)
                out.append("%sfunc (u *%s) %s(x %s) {" % (indentation(), goname, settername,
                                                          ti["gotype"]))
                push_indent()
                out.append("%sp := (*%s)(unsafe.Pointer(u))" % (indentation(), ti["ctype"]))
                recursive_copy(refid, kindref, ti, "*p", "x")
                pop_indent()
                out.append(indentation() + "}")


def recursive_copy(refid, kindref, ti, dest, source):
    '''
    Appends to out a series of assignment statements that copy the Go variable source
    to the C variable dest, whose type is described by ti.
    refid/kindref are the doxygen refid/kindref of dest's type, e.g.
    "struct_s_d_l___touch_finger_event"/"compound" for an object of C type
    struct SDL_TouchFingerEvent.
    dest is a string that accesses the C value (not a pointer to it) in the context
    where the output lines are used. E.g. dest could be "*p" if "p" is
    a pointer to the C object in question that is defined in the respective context.
    '''
    if kindref != "compound":
        out.append("%s%s = %s(%s)%s" % (indentation(), dest, ti["ccast"], source, ti["ccastend"]))
    else:
        with open("%s/%s.xml" % (doxyxml, refid)) as f:
            compound = BeautifulSoup(f, "xml").doxygen.compounddef

            name = str(compound.compoundname.string)
            if name in blacklist:
                return recursive_copy(refid, "blacklisted", ti, dest, source)

            fields = list(compound.find_all("memberdef", kind="variable", prot="public"))
            fields.sort(key=lambda x: int(x.location["line"]))

            for member in fields:
                membname = str(member.find("name").string)
                cmembname = fix(membname, True)
                try:
                    refid2 = str(member.type.ref["refid"])
                    kindref2 = str(member.type.ref["kindref"])
                except:
                    refid2 = ""
                    kindref2 = "primitive"
                typ = get_ctype(member.type)
                typargs = ""
                if member.argsstring is not None:
                    typargs = str("".join(member.argsstring.stripped_strings))
                ti2 = typeinfo(name, 0, membname, typ, typargs)
                recursive_copy(refid2, kindref2, ti2, "%s.%s" % (dest.lstrip("*"), cmembname),
                               "%s.%s" % (source, fix(membname)))


def go_copy_of(refid, kindref, ti, value):
    '''
    Returns a string that represents a Go expression that creates a Go object that
    corresponds to value, which is a C object of the type described by typeinfo() dict ti.
    refid/kindref are the doxygen refid/kindref of value's type, e.g.
    "struct_s_d_l___touch_finger_event"/"compound" for an object of C type
    struct SDL_TouchFingerEvent.
    value is a string that accesses the C value (not a pointer to it) in the context
    where the result of go_copy_of is used. E.g. value could be "*p" if "p" is
    a pointer to the C object in question that is defined in the respective context.
    '''
    if kindref != "compound":
        return "%s(%s)%s" % (ti["gocast"], value, ti["gocastend"])
    else:
        s = []
        with open("%s/%s.xml" % (doxyxml, refid)) as f:
            compound = BeautifulSoup(f, "xml").doxygen.compounddef

            name = str(compound.compoundname.string)
            if name in blacklist:
                return go_copy_of(refid, "blacklisted", ti, value)

            fields = list(compound.find_all("memberdef", kind="variable", prot="public"))
            fields.sort(key=lambda x: int(x.location["line"]))

            for member in fields:
                membname = str(member.find("name").string)
                cmembname = fix(membname, True)
                try:
                    refid2 = str(member.type.ref["refid"])
                    kindref2 = str(member.type.ref["kindref"])
                except:
                    refid2 = ""
                    kindref2 = "primitive"
                typ = get_ctype(member.type)
                typargs = ""
                if member.argsstring is not None:
                    typargs = str("".join(member.argsstring.stripped_strings))
                ti2 = typeinfo(name, 0, membname, typ, typargs)
                s.append(
                    go_copy_of(refid2, kindref2, ti2, "%s.%s" % (value.lstrip("*"), cmembname)))

        return ti["gotype"] + "{" + ", ".join(s) + "}"


def define2const(section):
    '''Converts all defines in section (part of a BeautifulSoup) into a const block.'''
    global out
    global indent
    first = True
    for define in section.find_all("memberdef", kind="define", prot="public"):
        name = str(define.find("name").string)
        if name in blacklist:
            continue

        # Function-like defines can't be handled here
        if define.find("param") is not None:
            continue

        if define.find("initializer") is None:
            continue

        if first:
            out.append("const (")
            push_indent()
            first = False
        else:
            out.append("")  # empty line to separate entries

        describe(define)
        external_link(name)
        out.append("%s%s = C.%s" % (indentation(), fix(name), name))

    if not first:
        pop_indent()
        out.append(")")


def enum2const(section):
    '''Converts all enums in section (part of a BeautifulSoup) into const blocks.'''
    global out
    global indent
    firstE = True
    for enu in section.find_all("memberdef", kind="enum", prot="public"):
        name = str(enu.find("name").string)
        if name in blacklist:
            continue

        if not firstE:
            out.append("")
        firstE = False

        describe(enu)
        external_link(name)
        ti = typeinfo(name, 0, "", name, "")
        goname = ti["gotype"]
        if goname.isidentifier():
            enum_type = enum_types.get(name, "int")
            out.append("%stype %s %s" % (indentation(), goname, enum_type))
        else:
            goname = ""
        out.append(indentation() + "const (")
        push_indent()

        first = True
        for enuval in enu("enumvalue"):
            valname = str(enuval.find("name").string)
            if valname in blacklist:
                continue

            if not first:
                out.append("")  # empty line to separate entries

            first = False

            describe(enuval)
            out.append("%s%s %s = C.%s" % (indentation(), fix(valname), goname, valname))

        pop_indent()
        out.append(")")


def simpletypedefs(section):
    '''Converts all simple typedefs in section (part of a BeautifulSoup) into 
    the Go equivalent. A simple typedef is of the form "typedef Type1 Type2;"
    where Type1 and Type2 and single words with no special characters (not even "*").
    '''
    global out
    global indent
    first = True
    for typedef in section.find_all("memberdef", kind="typedef", prot="public"):
        # Test if typedef is simple
        td = str(typedef.definition.string).split()
        if len(td) != 3 or not td[1].isidentifier() or not td[2].isidentifier():
            continue

        ti1 = typeinfo(td[1], 0, "", td[1], "")
        ti2 = typeinfo(td[1], 0, "", td[2], "")

        if not first:
            out.append("")
        first = False

        describe(typedef)
        out.append("%stype %s %s" % (indentation(), ti2["gotype"], ti1["gotype"]))


def aliastypedefs(section):
    '''Outputs Go type definitions for all typedefs in section (part of a BeautifulSoup)
    where the Go type is simply an alias for the C.type.
    '''
    global out
    global indent
    first = True
    for typedef in section.find_all("memberdef", kind="typedef", prot="public"):
        # Test if typedef is not one of those handled by simpletypedefs()
        td = str(typedef.definition.string).split()
        if len(td) != 3 or not td[1].isidentifier() or not td[2].isidentifier():
            name = typedef.find("name").string

            if not first:
                out.append("")
            first = False

            describe(typedef)
            out.append("%stype %s C.%s" % (indentation(), fix(name), name))


def get_ctype(tag):
    '''
    Extracts a C type name from tag's contents.
    '''
    strs = []
    for s in tag.stripped_strings:
        strs.extend(x for x in s.split() if not x in ignored_type_elements)
    typ = " ".join(strs)
    typ = typ.replace(" *", "*").strip()
    if typ.rstrip("*").replace(
            "const ",
            "") not in blacklist and tag.ref is not None and tag.ref["refid"].startswith("struct"):
        typ = "struct " + typ
    return typ


def wrapfunctions(section):
    '''Output Go function wrappers for all functions in section (part of BeautifulSoup)'''
    global out
    global indent
    for fun in section.find_all("memberdef", kind="function", prot="public"):
        out.append("")  # empty line between functions
        name = str(fun.find("name").string)
        if name in blacklist:
            continue
        return_type = get_ctype(fun.type)
        return_type_args = ""
        params = []
        if return_type != "void":
            params.append(typeinfo(name, 0, "", return_type, return_type_args))

        argidx = 1
        for p in fun("param"):
            typ = get_ctype(p.type)
            typargs = ""
            if typ != "void":
                dn = p.declname
                if dn is not None:
                    argname = str(dn.string)
                else:
                    argname = ""
                params.append(typeinfo(name, argidx, argname, typ, typargs))
            argidx += 1

        num_recv, num_args, num_ret = fixup_params(params)
        if num_recv > 1:
            raise ValueError(
                "wrapfunctions(): Function '%s' tries to have more than 1 receiver" % name)

        slices_instead_of_pointers(name, params)

        describe(fun)
        external_link(name)
        goname = fix(name)
        s = "func "
        if num_recv > 0:
            s += "(%s %s) " % (params[0]["name"], params[0]["gotype"])
            # if the function has a receiver whose type name is a part
            # of the function name, remove that part
            recvpart = params[0]["gotype"].lstrip("*")
            if recvpart in receiver_aliases:
                for a in receiver_aliases[recvpart]:
                    goname = goname.replace(a, "", 1)
            else:
                goname = goname.replace(recvpart, "", 1)

        s += goname + "("
        first = True
        for i in range(1, num_args + 1):
            if "nogo" in params[i]:
                continue
            if not first:
                s += ", "
            first = False
            s += "%s %s" % (params[i]["name"], params[i]["gotype"])

        s += ")"

        if num_ret > 0:
            s += " ("
            for i in range(num_args + 1, num_args + 1 + num_ret):
                if i > num_args + 1:
                    s += ", "
                s += "%s %s" % (params[i]["name"], params[i]["gotype"])
            s += ")"

        s += " {"

        out.append(indentation() + s)
        push_indent()

        c_ret_idx = -1
        for i in range(1, len(params)):
            if params[i]["cidx"] == 0:
                c_ret_idx = i
            if params[i]["alloc"] != "":
                out.append(indentation() + params[i]["alloc"])

        s = ""
        copy_retval = ""
        if c_ret_idx >= 0:
            ret_name = params[c_ret_idx]["name"] + " "
            if need_temp_to_get_pointer(params[c_ret_idx]):
                copy_retval = "%s = &tmp_%s" % (ret_name, ret_name)
                ret_name = "tmp_" + ret_name + " :"
            s += "%s= %s(" % (ret_name, params[c_ret_idx]["gocast"])

        s += "C." + name + "("

        c_args = list(
            i for i in range(0, len(params)) if params[i] is not None and params[i]["cidx"] != 0)
        c_args.sort(key=lambda i: params[i]["cidx"])
        comma = ""
        for i in c_args:
            s += "%s%s(%s)%s" % (comma, params[i]["ccast"], params[i]["allocarg"],
                                 params[i]["ccastend"])
            comma = ", "

        s += ")"

        if c_ret_idx >= 0:
            s += ")" + params[c_ret_idx]["gocastend"]

        out.append(indentation() + s)
        if copy_retval != "":
            out.append(indentation() + copy_retval)

        for i in range(1, len(params)):
            if params[i]["dealloc"] != "":
                out.append("%s%s" % (indentation(), params[i]["dealloc"]))

        if num_ret > 0:
            out.append(indentation() + "return")

        pop_indent()
        out.append(indentation() + "}")


DOXYGEN_MAPPING = {
    '_': "__",
    ':': "_1",
    '/': "_2",
    '<': "_3",
    '>': "_4",
    '*': "_5",
    '&': "_6",
    '|': "_7",
    '.': "_8",
    '!': "_9",
    ',': "_00",
    ' ': "_01",
    '{': "_02",
    '}': "_03",
    '?': "_04",
    '^': "_05",
    '%': "_06",
    '(': "_07",
    ')': "_08",
    '+': "_09",
    '=': "_0A",
    '$': "_0B",
    '\\': "_0C",
    '@': "_0D"
}


def need_temp_to_get_pointer(ti):
    '''Returns true if the typeinfo ti describes a type that needs
    a temporary because gocast does not produce a pointer but gotype needs one
    and we can't take the address of gocast.'''
    return ti["gotype"].startswith("*") and ti["gocast"][0].islower()


def get_doxyname(name):
    '''
    Applies Doxygen's filename escaping rules to name and returns the result.
    (This is not a complete implementation. It only handles ASCII characters.)
    (Doxygen source ref: src/util.cpp:escapeCharsInString())
    '''
    out = ""
    for ch in name:
        if ch in DOXYGEN_MAPPING:
            out += DOXYGEN_MAPPING[ch]
        else:
            lch = ch.lower()
            if lch != ch:
                out += "_"
            out += lch

    return out
