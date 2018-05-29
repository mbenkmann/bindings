#!/usr/bin/python3

import lib
from bs4 import BeautifulSoup
import sys
import os.path

SDL_COMMON_FILE_HEADER = '''// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"'''

SDL_TYPE_MAPPING = {
    "SDL_bool": "bool",
    "uint": "uint",
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

    # We use uintptr as type for void* because
    # a) we don't want to force users to import "unsafe"
    # b) we want to make users think twice when using such a pointer to C space
    # c) a typical place where this occurs is userdata for callbacks and we
    #    want to encourage people to use a uintptr as key into a map instead of
    #    as an actual pointer (which would in fact be illegal if it's a ptr to
    #    a Go object stored in C space)
    "void*": "uintptr",
    "float": "float32",
    "double": "float64",
    "const Uint8*": "*byte",
    "Uint8*": "*byte",
    "WindowShapeMode": "ShapeMode"
}

SDL_GOCAST = {"SDL_bool": "C.SDL_TRUE=="}
SDL_CCAST = {"SDL_bool": "bool2bool"}

SDL_POINTER_ARG = {
    "Uint8": {
        "default":
        "in",
        "out": [
            "SDL_GetTextureColorMod", "SDL_GetSurfaceColorMod", "SDL_GetTextureAlphaMod",
            "SDL_GetSurfaceAlphaMod", "SDL_GetRenderDrawColor", "SDL_GetRGB", "SDL_GetRGBA"
        ]
    },
    "SDL_Event": {
        "default": "out",
        "receiver": {"SDL_PushEvent"}
    },
    "SDL_Joystick": {
        "default": "receiver"
    },
    "SDL_GameController": {
        "default": "receiver"
    },
    "SDL_Window": {
        "default": "receiver",
        "in": {
            "SDL_WarpMouseInWindow", "SDL_ShowSimpleMessageBox",
            "SDL_SetWindowModalFor.parent_window"
        }
    },
    "SDL_RWops": {
        "default": "in",
        "receiver": {
            "SDL_ReadU8", "SDL_ReadLE16", "SDL_ReadBE16", "SDL_ReadLE32", "SDL_ReadBE32",
            "SDL_ReadLE64", "SDL_ReadBE64", "SDL_WriteU8", "SDL_WriteLE16", "SDL_WriteBE16",
            "SDL_WriteLE32", "SDL_WriteBE32", "SDL_WriteLE64", "SDL_WriteBE64"
        }
    },
    "SDL_Surface": {
        "default": "receiver",
        "in": {
            "SDL_SetWindowIcon", "SDL_SaveBMP_RW", "SDL_UpperBlit", "SDL_UpperBlitScaled",
            "SDL_LowerBlit", "SDL_LowerBlitScaled", "SDL_SoftStretch",
            "SDL_CreateTextureFromSurface", "SDL_SetWindowShape", "SDL_CreateColorCursor"
        }
    },
    "SDL_Renderer": {
        "default": "receiver",
    },
    "SDL_Texture": {
        "default": "receiver",
        "in": {
            "SDL_SetRenderTarget",
            "SDL_RenderCopy",
            "SDL_RenderCopyEx",
        }
    },
    "SDL_Rect": {
        "by-value": True,
        "by-ptr":
        {"SDL_RenderCopy", "SDL_UpdateTexture", "SDL_UpdateYUVTexture", "SDL_RenderReadPixels"},
        "default": "in",  # not as receiver to keep option open to write native Go methods
        "out": {
            "SDL_IntersectRect.result", "SDL_UnionRect.result", "SDL_GetDisplayBounds",
            "SDL_GetDisplayUsableBounds", "SDL_GetClipRect", "SDL_RenderGetViewport",
            "SDL_RenderGetClipRect", "SDL_EnclosePoints.result"
        }
    },
    "SDL_Point": {
        "by-value": True,
        "default": "in",  # not as receiver to keep option open to write native Go methods
    },
    "SDL_Color": {
        "by-value": True,
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
    "SDL_WindowShapeMode": {
        "default": "in",
        "out": {"SDL_GetShapedWindowMode"}
    },
    "SDL_Cursor": {
        "default": "receiver",
        "in": {"SDL_SetCursor"},
    },
    "SDL_Haptic": {
        "default": "receiver",
    },
    "SDL_HapticEffect": {
        "default": "in",
    },
    "SDL_AudioCVT": {
        "default": "in"
    },
    "SDL_AudioStream": {
        "default": "receiver"
    },
    "SDL_AudioSpec": {
        "default": "in",
        "inout": {"SDL_OpenAudio.obtained", "SDL_OpenAudioDevice.obtained"}
    }
}

SDL_BLACKLIST = frozenset(
    ("SDL_SetError", "SDL_OutOfMemory", "SDL_Unsupported", "SDL_InvalidParamError",
     "SDL_GetEventState", "SDL_SysWMEvent", "SDL_Event.syswm", "Event.SetDrop", "toCFromDropEvent",
     "SDL_PeepEvents", "SDL_bool", "SDL_SetEventFilter", "SDL_GetEventFilter", "SDL_AddEventWatch",
     "SDL_DelEventWatch", "SDL_FilterEvents", "SDL_JoystickGetGUIDString",
     "SDL_IntersectRectAndLine", "SDL_RWops", "SDL_RWFromFP", "SDL_RWFromMem", "SDL_RWFromConstMem",
     "SDL_AllocRW", "SDL_FreeRW", "SDL_SetWindowGammaRamp", "SDL_GetWindowGammaRamp",
     "SDL_PixelFormat", "SDL_Surface", "SDL_BlitSurface", "SDL_BlitScaled", "SDL_Palette",
     "SDL_Colour", "SDL_CalculateGammaRamp", "toCFromRendererInfo", "SDL_CreateWindowAndRenderer",
     "SDL_LockTexture", "fromC2WindowShapeMode", "toCFromWindowShapeMode",
     "SDL_GameControllerButtonBind", "fromC2MessageBoxButtonData", "fromC2MessageBoxColorScheme",
     "toCFromMessageBoxColorScheme", "SDL_MessageBoxData", "SDL_ShowMessageBox", "SDL_HapticCustom",
     "SDL_AudioCVT", "SDL_BuildAudioCVT", "SDL_ConvertAudio", "SDL_LoadWAV_RW", "SDL_FreeWAV",
     "SDL_LoadFile_RW", "SDL_GetKeyboardState", "SDL_CreateCursor"))

SDL_IGNORED_TYPE_ELEMENTS = frozenset(("SDL_FORCE_INLINE", ))

SDL_RECEIVER_ALIASES = {"Renderer": ["Renderer", "Render"]}

SDL_FREE_STRINGS = frozenset(("SDL_GetBasePath", "SDL_GetPrefPath", "SDL_GetClipboardText",
                              "SDL_GameControllerMapping"))

SDL_GOTYPE_OVERRIDE = {
    "SDL_CreateWindow.flags": "WindowFlags",
    "SDL_CreateShapedWindow.flags": "WindowFlags",
    "SDL_CreateRenderer.flags": "RendererFlags",
    "SDL_*Event.type": "EventType",
    "SDL_FlushEvents.*Type": "EventType",
}

SDL_ENUM_TYPES = {"SDL_EventType": "uint32"}


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
    lib.receiver_aliases = SDL_RECEIVER_ALIASES
    lib.free_strings = SDL_FREE_STRINGS
    lib.gotype_override = SDL_GOTYPE_OVERRIDE
    lib.enum_types = SDL_ENUM_TYPES


if __name__ == "__main__":
    sdl()
    lib.out.append(SDL_COMMON_FILE_HEADER)
    import_idx = len(lib.out)
    lib.out.append("")

    lib.generate_bindings()

    lib.additional_boilerplate(import_idx)

    print("\n".join(lib.out))
