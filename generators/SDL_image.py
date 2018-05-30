#!/usr/bin/python3

import lib
from bs4 import BeautifulSoup
import sys
import os.path
from SDL import sdl

IMG_COMMON_FILE_HEADER = '''// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package img

// #include "includes.h"
import "C"

import "winterdrache.de/bindings/sdl"'''

IMG_BLACKLIST = frozenset(("IMG_GetError", "IMG_SetError", "IMG_ReadXPMFromArray"))
IMG_POINTER_ARG = {
    "SDL_RWops": {
        "default": "in"
    },
    "SDL_Renderer": {
        "default": "in"
    },
    "SDL_Surface": {
        "default": "in"
    },
}


def sdl_image():
    '''
    Called at the very beginning of processing an SDL* header.
    Checks sys.argv and exits the program on error.
    Loads the data for the header into the variable soup.
    '''

    sdl()
    lib.prefixes.append("IMG_")
    lib.blacklist = lib.blacklist.union(IMG_BLACKLIST)
    lib.cprefix_to_go_package = {"SDL_": "sdl"}
    lib.pointer_arg_treatment = IMG_POINTER_ARG


if __name__ == "__main__":
    sdl_image()
    lib.out.append(IMG_COMMON_FILE_HEADER)
    import_idx = len(lib.out)
    lib.out.append("")

    lib.generate_bindings()

    lib.additional_boilerplate(import_idx)

    print("\n".join(lib.out))
