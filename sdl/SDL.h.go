// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"

 // Main include header for the SDL library

 // SDL_INIT_*
 // 
 // These are the flags which may be passed to SDL_Init(). You should
 // specify the subsystems which you will be using in your application.
const (
     // ↪ https://wiki.libsdl.org/SDL_INIT_TIMER
    INIT_TIMER = C.SDL_INIT_TIMER

     // ↪ https://wiki.libsdl.org/SDL_INIT_AUDIO
    INIT_AUDIO = C.SDL_INIT_AUDIO

     // SDL_INIT_VIDEO implies SDL_INIT_EVENTS
     // ↪ https://wiki.libsdl.org/SDL_INIT_VIDEO
    INIT_VIDEO = C.SDL_INIT_VIDEO

     // SDL_INIT_JOYSTICK implies SDL_INIT_EVENTS
     // ↪ https://wiki.libsdl.org/SDL_INIT_JOYSTICK
    INIT_JOYSTICK = C.SDL_INIT_JOYSTICK

     // ↪ https://wiki.libsdl.org/SDL_INIT_HAPTIC
    INIT_HAPTIC = C.SDL_INIT_HAPTIC

     // SDL_INIT_GAMECONTROLLER implies SDL_INIT_JOYSTICK
     // ↪ https://wiki.libsdl.org/SDL_INIT_GAMECONTROLLER
    INIT_GAMECONTROLLER = C.SDL_INIT_GAMECONTROLLER

     // ↪ https://wiki.libsdl.org/SDL_INIT_EVENTS
    INIT_EVENTS = C.SDL_INIT_EVENTS

     // compatibility; this flag is ignored.
     // ↪ https://wiki.libsdl.org/SDL_INIT_NOPARACHUTE
    INIT_NOPARACHUTE = C.SDL_INIT_NOPARACHUTE

     // ↪ https://wiki.libsdl.org/SDL_INIT_EVERYTHING
    INIT_EVERYTHING = C.SDL_INIT_EVERYTHING
)


 // This function initializes the subsystems specified by flags
 // ↪ https://wiki.libsdl.org/SDL_Init
func Init(flags uint32) (retval int) {
    retval = int(C.SDL_Init(C.Uint32(flags)))
    return
}

 // This function initializes specific SDL subsystems
 // 
 // Subsystem initialization is ref-counted, you must call
 // SDL_QuitSubSystem() for each SDL_InitSubSystem() to correctly shutdown
 // a subsystem manually (or call SDL_Quit() to force shutdown). If a
 // subsystem is already loaded then this call will increase the ref-count
 // and return.
 // ↪ https://wiki.libsdl.org/SDL_InitSubSystem
func InitSubSystem(flags uint32) (retval int) {
    retval = int(C.SDL_InitSubSystem(C.Uint32(flags)))
    return
}

 // This function cleans up specific SDL subsystems
 // ↪ https://wiki.libsdl.org/SDL_QuitSubSystem
func QuitSubSystem(flags uint32) {
    C.SDL_QuitSubSystem(C.Uint32(flags))
}

 // This function returns a mask of the specified subsystems which have
 // previously been initialized.
 // 
 // If flags is 0, it returns a mask of all initialized subsystems.
 // ↪ https://wiki.libsdl.org/SDL_WasInit
func WasInit(flags uint32) (retval uint32) {
    retval = uint32(C.SDL_WasInit(C.Uint32(flags)))
    return
}

 // This function cleans up all initialized subsystems. You should call it
 // upon all exit conditions.
 // ↪ https://wiki.libsdl.org/SDL_Quit
func Quit() {
    C.SDL_Quit()
}
