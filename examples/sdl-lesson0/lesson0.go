// https://www.willusher.io/pages/sdl2/
// https://github.com/Twinklebear/TwinklebearDev-Lessons/tree/master/Lesson0

package main

import (
    "fmt"
    "os"
    "runtime"

    "winterdrache.de/bindings/sdl"
)

func main() {
    runtime.LockOSThread() // SDL functions must all be called from the main thread

    if sdl.Init(sdl.INIT_VIDEO) != 0 {
        fmt.Fprintf(os.Stderr, "SDL_Init Error: %v\n", sdl.GetError())
        os.Exit(1)
    }
    defer sdl.Quit()

    version := sdl.GetVersion()
    fmt.Printf("SDL Version: %v.%v.%v\n", version.Major, version.Minor, version.Patch)

    os.Exit(0)
}
