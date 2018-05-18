// Machine-generated file.
// See http://winterdrache.de/bindings for details.

// Bindings for Simple DirectMedia Layer (www.libsdl.org)
package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
import "C"

 // Main include header for the SDL library

 // SDL_INIT_*
 // 
 // These are the flags which may be passed to SDL_Init(). You should
 // specify the subsystems which you will be using in your application.
const (
    INIT_TIMER = C.SDL_INIT_TIMER

    INIT_AUDIO = C.SDL_INIT_AUDIO

     // SDL_INIT_VIDEO implies SDL_INIT_EVENTS
    INIT_VIDEO = C.SDL_INIT_VIDEO

     // SDL_INIT_JOYSTICK implies SDL_INIT_EVENTS
    INIT_JOYSTICK = C.SDL_INIT_JOYSTICK

    INIT_HAPTIC = C.SDL_INIT_HAPTIC

     // SDL_INIT_GAMECONTROLLER implies SDL_INIT_JOYSTICK
    INIT_GAMECONTROLLER = C.SDL_INIT_GAMECONTROLLER

    INIT_EVENTS = C.SDL_INIT_EVENTS

     // compatibility; this flag is ignored.
    INIT_NOPARACHUTE = C.SDL_INIT_NOPARACHUTE

    INIT_EVERYTHING = C.SDL_INIT_EVERYTHING
)


 // This function initializes the subsystems specified by flags
func Init(flags uint32) (retval int) {
    retval = int(C.SDL_Init(C.Uint32(flags)))
    return
}

 // This function initializes specific SDL subsystems
 // 
 // Subsystem initialization is ref-counted, you must call
 // SDL_QuitSubSystem for each SDL_InitSubSystem to correctly shutdown a
 // subsystem manually (or call SDL_Quit to force shutdown). If a
 // subsystem is already loaded then this call will increase the ref-count
 // and return.
func InitSubSystem(flags uint32) (retval int) {
    retval = int(C.SDL_InitSubSystem(C.Uint32(flags)))
    return
}

 // This function cleans up specific SDL subsystems
func QuitSubSystem(flags uint32) {
    C.SDL_QuitSubSystem(C.Uint32(flags))
}

 // This function returns a mask of the specified subsystems which have
 // previously been initialized.
 // 
 // If flags is 0, it returns a mask of all initialized subsystems.
func WasInit(flags uint32) (retval uint32) {
    retval = uint32(C.SDL_WasInit(C.Uint32(flags)))
    return
}

 // This function cleans up all initialized subsystems. You should call it
 // upon all exit conditions.
func Quit() {
    C.SDL_Quit()
}
