// Machine-generated file.
// See http://winterdrache.de/bindings for details.

// Bindings for Simple DirectMedia Layer (www.libsdl.org)
package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
import "C"

 // Include file for SDL gesture event handling.

type GestureID int64


 // Begin Recording a gesture on the specified touch, or all touches (-1)
func RecordGesture(touchId TouchID) (retval int) {
    retval = int(C.SDL_RecordGesture(C.SDL_TouchID(touchId)))
    return
}

 // Save all currently loaded Dollar Gesture templates.
func (dst *RWops) SaveAllDollarTemplates() (retval int) {
    retval = int(C.SDL_SaveAllDollarTemplates((*C.SDL_RWops)(dst)))
    return
}

 // Save a currently loaded Dollar Gesture template.
func (dst *RWops) SaveDollarTemplate(gestureId GestureID) (retval int) {
    retval = int(C.SDL_SaveDollarTemplate(C.SDL_GestureID(gestureId), (*C.SDL_RWops)(dst)))
    return
}

 // Load Dollar Gesture templates from a file.
func (src *RWops) LoadDollarTemplates(touchId TouchID) (retval int) {
    retval = int(C.SDL_LoadDollarTemplates(C.SDL_TouchID(touchId), (*C.SDL_RWops)(src)))
    return
}
