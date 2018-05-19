// https://www.willusher.io/pages/sdl2/
// https://github.com/Twinklebear/TwinklebearDev-Lessons/tree/master/Lesson1

package main

import (
    "fmt"
    "os"
    "runtime"

    "winterdrache.de/bindings/sdl"
    "winterdrache.de/bindings/sdlutil"
)

/*
 * Lesson 1: Hello World!
 */
func main() {
    runtime.LockOSThread() // SDL functions must all be called from the main thread

    //First we need to start up SDL, and make sure it went ok
    if sdl.Init(sdl.INIT_VIDEO) != 0 {
        fmt.Fprintf(os.Stderr, "sdl.Init Error: %v\n", sdl.GetError())
        os.Exit(1)
    }
    defer sdl.Quit()

    //Now create a window with title "Hello World" at 100, 100 on the screen with w:640 h:480 and show it
    win := sdl.CreateWindow("Hello World!", 100, 100, 640, 480, sdl.WINDOW_SHOWN)
    //Make sure creating our window went ok
    if win == nil {
        fmt.Fprintf(os.Stderr, "sdl.CreateWindow Error: %v\n", sdl.GetError())
        os.Exit(1)
    }
    defer win.Destroy()

    //Create a renderer that will draw to the window, -1 specifies that we want to load whichever
    //video driver supports the flags we're passing
    //Flags: sdl.RENDERER_ACCELERATED: We want to use hardware accelerated rendering
    //sdl.RENDERER_PRESENTVSYNC: We want the renderer's present function (update screen) to be
    //synchronized with the monitor's refresh rate
    renderer := win.CreateRenderer(-1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
    if renderer == nil {
        fmt.Fprintf(os.Stderr, "sdl.CreateRenderer Error: %v\n", sdl.GetError())
        os.Exit(1)
    }
    defer renderer.Destroy()

    //SDL 2.0 now uses textures to draw things but sdl.LoadBMP returns a surface
    //this lets us choose when to upload or remove textures from the GPU
    imagePath := sdlutil.GetResourcePath("../examples/resources", "Lesson1", "hello.bmp")
    bmpSurface := sdl.LoadBMP(imagePath)
    if bmpSurface == nil {
        fmt.Fprintf(os.Stderr, "sdl.LoadBMP Error: %v\n", sdl.GetError())
        os.Exit(1)
    }

    //To use a hardware accelerated texture for rendering we can create one from
    //the surface we loaded
    tex := renderer.CreateTextureFromSurface(bmpSurface)
    //We no longer need the surface
    bmpSurface.Free()
    if tex == nil {
        fmt.Fprintf(os.Stderr, "sdl.CreateTextureFromSurface Error: %v\n", sdl.GetError())
        os.Exit(1)
    }
    defer tex.Destroy()

    //A sleepy rendering loop, wait for 3 seconds and render and present the screen each time
    for i := 0; i < 3; i++ {
        //First clear the renderer
        renderer.Clear()
        //Draw the texture
        renderer.Copy(tex, nil, nil)
        //Update the screen
        renderer.Present()
        //Take a quick break after all that hard work
        sdl.Delay(1000)
    }

    os.Exit(0)
}
