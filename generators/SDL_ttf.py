#!/usr/bin/python3

import lib
from SDL import sdl

TTF_COMMON_FILE_HEADER = '''// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package ttf

// #include "includes.h"
import "C"

import "winterdrache.de/bindings/sdl"'''

TTF_BLACKLIST = frozenset(
    ("TTF_GetError", "TTF_SetError", "TTF_ByteSwappedUNICODE", "TTF_SizeUNICODE",
     "TTF_RenderUNICODE_Solid", "TTF_RenderUNICODE_Shaded", "TTF_RenderUNICODE_Blended",
     "TTF_RenderUNICODE_Blended_Wrapped", "TTF_GetFontKerningSize", "TTF_MAJOR_VERSION",
     "TTF_MINOR_VERSION", "TTF_PATCHLEVEL"))

TTF_POINTER_ARG = {
    "SDL_RWops": {
        "default": "in"
    },
    "TTF_Font": {
        "default": "receiver"
    },
}

TTF_GOTYPE_OVERRIDE = {
    "TTF_WasInit": "bool",
}

TTF_CCAST = {"SDL_Color": "toCFromColor"}


def sdl_ttf():
    '''
    Called at the very beginning of processing an SDL* header.
    Checks sys.argv and exits the program on error.
    Loads the data for the header into the variable soup.
    '''

    sdl()
    lib.prefixes.append("TTF_")
    lib.blacklist = lib.blacklist.union(TTF_BLACKLIST)
    lib.cprefix_to_go_package = {"SDL_": "sdl"}
    lib.pointer_arg_treatment = TTF_POINTER_ARG
    lib.custom_ccast = TTF_CCAST


if __name__ == "__main__":
    sdl_ttf()
    lib.out.append(TTF_COMMON_FILE_HEADER)
    import_idx = len(lib.out)
    lib.out.append("")

    lib.generate_bindings()

    lib.additional_boilerplate(import_idx)

    print("\n".join(lib.out))
