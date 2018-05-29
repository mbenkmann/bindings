// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"

 // Include file for SDL touch event handling.

 // ↪ https://wiki.libsdl.org/SDL_Finger
type Finger struct {
    Id FingerID

    X float32

    Y float32

    Pressure float32
}

func fromC2Finger(s C.SDL_Finger) Finger {
    return Finger{FingerID(s.id), float32(s.x), float32(s.y), float32(s.pressure)}
}

func toCFromFinger(s Finger) (d C.SDL_Finger) {
    d.id = C.SDL_FingerID(s.Id)
    d.x = C.float(s.X)
    d.y = C.float(s.Y)
    d.pressure = C.float(s.Pressure)
    return
}

const (
    TOUCH_MOUSEID = C.SDL_TOUCH_MOUSEID
)

type TouchID int64

type FingerID int64


 // Get the number of registered touch devices.
 // ↪ https://wiki.libsdl.org/SDL_GetNumTouchDevices
func GetNumTouchDevices() (retval int) {
    retval = int(C.SDL_GetNumTouchDevices())
    return
}

 // Get the touch ID with the given index, or 0 if the index is invalid.
 // ↪ https://wiki.libsdl.org/SDL_GetTouchDevice
func GetTouchDevice(index int) (retval TouchID) {
    retval = TouchID(C.SDL_GetTouchDevice(C.int(index)))
    return
}

 // Get the number of active fingers for a given touch device.
 // ↪ https://wiki.libsdl.org/SDL_GetNumTouchFingers
func GetNumTouchFingers(touchID TouchID) (retval int) {
    retval = int(C.SDL_GetNumTouchFingers(C.SDL_TouchID(touchID)))
    return
}

 // Get the finger object of the given touch, with the given index.
 // ↪ https://wiki.libsdl.org/SDL_GetTouchFinger
func GetTouchFinger(touchID TouchID, index int) (retval Finger) {
    retval = fromC2Finger(*(C.SDL_GetTouchFinger(C.SDL_TouchID(touchID), C.int(index))))
    return
}
