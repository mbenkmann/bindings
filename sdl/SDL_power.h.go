// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"

 // Header for the SDL power management routines.

 // The basic state for the system's power supply.
 // ↪ https://wiki.libsdl.org/SDL_PowerState
type PowerState int
const (
     // cannot determine power status
    POWERSTATE_UNKNOWN PowerState = C.SDL_POWERSTATE_UNKNOWN

     // Not plugged in, running on the battery
    POWERSTATE_ON_BATTERY PowerState = C.SDL_POWERSTATE_ON_BATTERY

     // Plugged in, no battery available
    POWERSTATE_NO_BATTERY PowerState = C.SDL_POWERSTATE_NO_BATTERY

     // Plugged in, charging battery
    POWERSTATE_CHARGING PowerState = C.SDL_POWERSTATE_CHARGING

     // Plugged in, battery charged
    POWERSTATE_CHARGED PowerState = C.SDL_POWERSTATE_CHARGED
)


 // Get the current power supply details.
 // 
 // Returns: The state of the battery (if any).
 // 
 //   secs
 //     Seconds of battery life left. You can pass a NULL here if you don't
 //     care. Will return -1 if we can't determine a value, or we're not
 //     running on a battery.
 //   
 //   pct
 //     Percentage of battery life left, between 0 and 100. You can pass a
 //     NULL here if you don't care. Will return -1 if we can't determine a
 //     value, or we're not running on a battery.
 //   
 // ↪ https://wiki.libsdl.org/SDL_GetPowerInfo
func GetPowerInfo() (retval PowerState, secs int, pct int) {
    tmp_secs := new(C.int)
    tmp_pct := new(C.int)
    retval = PowerState(C.SDL_GetPowerInfo((*C.int)(tmp_secs), (*C.int)(tmp_pct)))
    secs = deref_int_ptr(tmp_secs)
    pct = deref_int_ptr(tmp_pct)
    return
}
