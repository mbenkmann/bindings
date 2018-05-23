// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"

 // Include file for SDL gesture event handling.

type GestureID int64


 // Begin Recording a gesture on the specified touch, or all touches (-1)
func RecordGesture(touchId TouchID) (retval int) {
    retval = int(C.SDL_RecordGesture(C.SDL_TouchID(touchId)))
    return
}

 // Save all currently loaded Dollar Gesture templates.
func SaveAllDollarTemplates(dst *RWops) (retval int) {
    retval = int(C.SDL_SaveAllDollarTemplates((*C.SDL_RWops)(dst)))
    return
}

 // Save a currently loaded Dollar Gesture template.
func SaveDollarTemplate(gestureId GestureID, dst *RWops) (retval int) {
    retval = int(C.SDL_SaveDollarTemplate(C.SDL_GestureID(gestureId), (*C.SDL_RWops)(dst)))
    return
}

 // Load Dollar Gesture templates from a file.
func LoadDollarTemplates(touchId TouchID, src *RWops) (retval int) {
    retval = int(C.SDL_LoadDollarTemplates(C.SDL_TouchID(touchId), (*C.SDL_RWops)(src)))
    return
}
