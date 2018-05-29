// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"

 // Simple error message routines for SDL.

 // Internal error functions
type Errorcode int
const (
    ENOMEM Errorcode = C.SDL_ENOMEM

    EFREAD Errorcode = C.SDL_EFREAD

    EFWRITE Errorcode = C.SDL_EFWRITE

    EFSEEK Errorcode = C.SDL_EFSEEK

    UNSUPPORTED Errorcode = C.SDL_UNSUPPORTED

    LASTERROR Errorcode = C.SDL_LASTERROR
)

func Error(code Errorcode) (retval int) {
    retval = int(C.SDL_Error(C.SDL_errorcode(code)))
    return
}



 // ↪ https://wiki.libsdl.org/SDL_GetError
func GetError() (retval string) {
    retval = C.GoString(C.SDL_GetError())
    return
}

 // ↪ https://wiki.libsdl.org/SDL_ClearError
func ClearError() {
    C.SDL_ClearError()
}
