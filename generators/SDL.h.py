#!/usr/bin/python3

import lib
from bs4 import BeautifulSoup
import sys
import os.path

SDL_COMMON_FILE_HEADER = '''// Machine-generated file.
// See http://winterdrache.de/bindings for details.

// Bindings for Simple DirectMedia Layer (www.libsdl.org)
package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
import "C"'''

SDL_TYPE_MAPPING = {
    "SDL_bool": "bool",
    "Uint8": "uint8",
    "Uint16": "uint16",
    "Sint16": "int16",
    "Sint32": "int32",
    "Uint32": "uint32",
    "Sint64": "int64",
    "Uint64": "uint64",
    "size_t": "uint64",
    "int": "int",
    "char": "int8",
    "char*": "string",
    "const char*": "string",
    "void*": "uintptr",
    "float": "float32",
    "const Uint8*": "*[999999]byte"
}

SDL_GOCAST = {"SDL_bool": "C.SDL_TRUE=="}
SDL_CCAST = {"SDL_bool": "bool2bool"}

SDL_POINTER_ARG = {
    "SDL_Event": {
        "default": "out",
        "receiver": ["SDL_PushEvent"]
    },
    "SDL_Joystick": {
        "default": "receiver"
    },
    "SDL_Window": {
        "default": "receiver"
    },
    "SDL_RWops": {
        "default": "receiver"
    },
    "SDL_Surface": {
        "default": "receiver",
        "in": {
            "SDL_SetWindowIcon",
            "SDL_SaveBMP_RW",
            "SDL_UpperBlit",
            "SDL_UpperBlitScaled",
            "SDL_LowerBlit",
            "SDL_LowerBlitScaled",
            "SDL_SoftStretch",
        }
    },
    "SDL_Rect": {
        "by-value":
        True,
        "default":
        "in",  # not as receiver to keep option open to write native Go methods
        "out": [
            "SDL_IntersectRect.result", "SDL_UnionRect.result", "SDL_GetDisplayBounds",
            "SDL_GetClipRect"
        ]
    },
    "SDL_Point": {
        "by-value": True,
        "default": "in",  # not as receiver to keep option open to write native Go methods
    },
    "SDL_Finger": {
        "by-value": True,
    },
    "SDL_Palette": {
        "default": "receiver",
        "in": {"SDL_SetSurfacePalette", "SDL_SetPixelFormatPalette"}
    },
    "SDL_PixelFormat": {
        "default": "receiver",
        "in": {"SDL_ConvertSurface"}
    },
}

SDL_BLACKLIST = frozenset(
    ("SDL_SetError", "SDL_OutOfMemory", "SDL_Unsupported", "SDL_InvalidParamError",
     "SDL_GetEventState", "SDL_SysWMEvent", "SDL_Event.syswm", "Event.SetDrop", "toCFromDropEvent",
     "SDL_PeepEvents", "SDL_bool", "SDL_SetEventFilter", "SDL_GetEventFilter", "SDL_AddEventWatch",
     "SDL_DelEventWatch", "SDL_FilterEvents", "SDL_JoystickGetGUIDString",
     "SDL_IntersectRectAndLine", "SDL_EnclosePoints", "SDL_RWops", "SDL_RWFromFP", "SDL_RWFromMem",
     "SDL_RWFromConstMem", "SDL_AllocRW", "SDL_FreeRW", "SDL_UpdateWindowSurfaceRects",
     "SDL_SetWindowGammaRamp", "SDL_GetWindowGammaRamp", "SDL_PixelFormat", "SDL_Surface",
     "SDL_BlitSurface", "SDL_BlitScaled", "SDL_ConvertPixels", "SDL_CreateRGBSurfaceFrom",
     "SDL_FillRects", "SDL_Palette", "SDL_Colour", "SDL_SetPaletteColors",
     "SDL_CalculateGammaRamp"))

SDL_IGNORED_TYPE_ELEMENTS = frozenset(("SDL_FORCE_INLINE", ))


def sdl():
    '''
    Called at the very beginning of processing an SDL* header.
    Checks sys.argv and exits the program on error.
    Loads the data for the header into the variable soup.
    '''

    if len(sys.argv) != 3:
        sys.stderr.write("USAGE: headername.h doxygen-xml-dir\n")
        sys.exit(1)

    headername = sys.argv[1]
    lib.doxyxml = sys.argv[2]

    if not os.path.isdir(lib.doxyxml):
        sys.stderr.write("Not a directory: %s\n" % (lib.doxyxml, ))
        sys.exit(1)

    with open("%s/%s.xml" % (lib.doxyxml, lib.get_doxyname(headername))) as f:
        lib.soup = BeautifulSoup(f, "xml")

    for cls in lib.soup("innerclass", prot="public"):
        with open("%s/%s.xml" % (lib.doxyxml, cls["refid"])) as f:
            clssoup = BeautifulSoup(f, "xml")
            lib.soup.doxygen.append(clssoup.doxygen.compounddef)

    lib.prefixes = ["SDL_"]
    lib.typeinfo_heuristics.append(lib.BaseTypeinfo(SDL_TYPE_MAPPING))
    lib.blacklist = SDL_BLACKLIST
    lib.custom_gocast = SDL_GOCAST
    lib.custom_ccast = SDL_CCAST
    lib.pointer_arg_treatment = SDL_POINTER_ARG
    lib.ignored_type_elements = SDL_IGNORED_TYPE_ELEMENTS


sdl()
lib.out.append(SDL_COMMON_FILE_HEADER)
import_idx = len(lib.out)
lib.out.append("")

lib.describe(lib.soup.doxygen.compounddef)

lib.structs()
lib.unions()

for section in lib.soup.doxygen.compounddef.find_all("sectiondef"):
    lib.out.append("")
    lib.describe(section)
    lib.define2const(section)
    lib.enum2const(section)
    lib.aliastypedefs(section)
    lib.simpletypedefs(section)
    lib.wrapfunctions(section)

lib.additional_boilerplate(import_idx)

print("\n".join(lib.out))
