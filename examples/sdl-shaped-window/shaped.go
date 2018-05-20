package main

import (
    "fmt"
    "os"
    "runtime"

    "winterdrache.de/bindings/sdl"
    "winterdrache.de/bindings/sdlutil"
)

/*
 * Creates a non-rectangular window.
 * NOTE: Does not seem to work. But seems to be unrelated to the Go bindings:
 * https://www.reddit.com/r/sdl/comments/8gmy8w/sdl_createshapedwindow/
 * I've also searched github and only found code from a python-sdl test suite
 * with comments that suggest that the author has not made it work.
 * I'm pretty sure that this is a bug in SDL.
 *
 * DO NOT FILE A BUG REPORT FOR THE GO BINDINGS UNLESS YOU HAVE A WORKING
 * C PROGRAM THAT SHOWS A SHAPED WINDOW.
 */
func main() {
    os.Exit(sdlmain())
}

func sdlmain() int {
    runtime.LockOSThread() // SDL functions must all be called from the main thread

    if sdl.Init(sdl.INIT_VIDEO) != 0 {
        fmt.Fprintf(os.Stderr, "sdl.Init Error: %v\n", sdl.GetError())
        return 1
    }
    defer sdl.Quit()

    win := sdl.CreateShapedWindow("Shaped Window", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, 100, 100, sdl.WINDOW_SHOWN|sdl.WINDOW_BORDERLESS)
    if win == nil {
        fmt.Fprintf(os.Stderr, "sdl.CreateShapedWindow Error: %v\n", sdl.GetError())
        return 1
    }
    defer win.Destroy()

    windowShape := sdl.LoadBMP(sdlutil.GetResourcePath("../examples/resources", "images", "shapemask.bmp"))
    if windowShape == nil {
        fmt.Fprintf(os.Stderr, "sdl.LoadBMP Error: %v\n", sdl.GetError())
        return 1
    }
    defer windowShape.Free()

    shapeMode := &sdl.WindowShapeMode{Mode: sdl.ShapeModeBinarizeAlpha}
    shapeMode.Parameters.SetBinarizationCutoff(100)
    if win.SetShape(windowShape, shapeMode) != 0 {
        fmt.Fprintf(os.Stderr, "sdl.SetShape Error: %v\n", sdl.GetError())
        return 1
    }

    renderer := win.CreateRenderer(-1, sdl.RENDERER_ACCELERATED)
    if renderer == nil {
        fmt.Fprintf(os.Stderr, "sdl.CreateRenderer Error: %v\n", sdl.GetError())
        return 1
    }
    defer renderer.Destroy()

    texture := renderer.CreateTextureFromSurface(windowShape)
    if texture == nil {
        fmt.Fprintf(os.Stderr, "sdl.CreateTextureFromSurface Error: %v\n", sdl.GetError())
        return 1
    }
    defer texture.Destroy()

    renderer.SetDrawColor(255, 100, 100, 100)

    for i := 0; i < 10; i++ {
        renderer.Clear()
        renderer.Copy(texture, nil, nil)
        renderer.Present()
        sdl.Delay(1000)
    }

    return 0
}
