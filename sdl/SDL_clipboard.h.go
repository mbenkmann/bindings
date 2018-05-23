// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"
import "unsafe"

 // Include file for SDL clipboard handling


 // Put UTF-8 text into the clipboard.
 // 
 // See also: SDL_GetClipboardText()
 // 
func SetClipboardText(text string) (retval int) {
    tmp_text := C.CString(text); defer C.free(unsafe.Pointer(tmp_text))
    retval = int(C.SDL_SetClipboardText((*C.char)(tmp_text)))
    return
}

 // Get UTF-8 text from the clipboard, which must be freed with SDL_free()
 // 
 // See also: SDL_SetClipboardText()
 // 
func GetClipboardText() (retval string) {
    retval = freeGoString(C.SDL_GetClipboardText())
    return
}

 // Returns a flag indicating whether the clipboard exists and contains a
 // text string that is non-empty.
 // 
 // See also: SDL_GetClipboardText()
 // 
func HasClipboardText() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasClipboardText())
    return
}
