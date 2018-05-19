// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
import "C"
import "unsafe"

 // Header for the SDL time management routines.


 // Function prototype for the timer callback function.
 // 
 // The callback function is passed the current timer interval and returns
 // the next timer interval. If the returned value is the same as the one
 // passed in, the periodic alarm continues, otherwise a new alarm is
 // scheduled. If the callback returns 0, the periodic alarm is cancelled.
type TimerCallback C.SDL_TimerCallback
 // Definition of the timer ID type.
type TimerID int


 // Get the number of milliseconds since the SDL library initialization.
 // 
 // Note: This value wraps if the program runs for more than ~49 days.
 // 
func GetTicks() (retval uint32) {
    retval = uint32(C.SDL_GetTicks())
    return
}

 // Get the current value of the high resolution counter.
func GetPerformanceCounter() (retval uint64) {
    retval = uint64(C.SDL_GetPerformanceCounter())
    return
}

 // Get the count per second of the high resolution counter.
func GetPerformanceFrequency() (retval uint64) {
    retval = uint64(C.SDL_GetPerformanceFrequency())
    return
}

 // Wait a specified number of milliseconds before returning.
func Delay(ms uint32) {
    C.SDL_Delay(C.Uint32(ms))
}

 // Add a new timer to the pool of timers already running.
 // 
 // Returns: A timer ID, or 0 when an error occurs.
 // 
func AddTimer(interval uint32, callback TimerCallback, param uintptr) (retval TimerID) {
    retval = TimerID(C.SDL_AddTimer(C.Uint32(interval), C.SDL_TimerCallback(callback), unsafe.Pointer(param)))
    return
}

 // Remove a timer knowing its ID.
 // 
 // Returns: A boolean value indicating success or failure.
 // 
 // Warning: It is not safe to remove a timer multiple times.
 // 
func RemoveTimer(id TimerID) (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_RemoveTimer(C.SDL_TimerID(id)))
    return
}
