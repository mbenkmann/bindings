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

    windowShape := sdl.LoadBMP(sdlutil.GetResourcePath("../examples/resources", "images", "troll.bmp"))
    if windowShape == nil {
        fmt.Fprintf(os.Stderr, "sdl.LoadBMP Error: %v\n", sdl.GetError())
        return 1
    }
    defer windowShape.Free()

    dimensions := windowShape.GetClipRect()
    width, height := uint(dimensions.W), uint(dimensions.H)

    win := sdl.CreateShapedWindow("Shaped Window", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
        width, height, sdl.WINDOW_SHOWN|sdl.WINDOW_BORDERLESS)
    if win == nil {
        fmt.Fprintf(os.Stderr, "sdl.CreateShapedWindow Error: %v\n", sdl.GetError())
        return 1
    }
    defer win.Destroy()

    renderer := win.CreateRenderer(-1, sdl.RENDERER_ACCELERATED)
    if renderer == nil {
        fmt.Fprintf(os.Stderr, "sdl.CreateRenderer Error: %v\n", sdl.GetError())
        return 1
    }
    defer renderer.Destroy()

    /*
       ATTENTION! win.SetShape() will only work
       a) AFTER win.CreateRenderer() has been called!
       b) IF then window size matches the shape's size
    */
    shapeMode := &sdl.WindowShapeMode{Mode: sdl.ShapeModeBinarizeAlpha}
    shapeMode.Parameters.SetBinarizationCutoff(100)
    if win.SetShape(windowShape, shapeMode) != 0 {
        fmt.Fprintf(os.Stderr, "sdl.SetShape Error: %v\n", sdl.GetError())
        return 1
    }

    texture := renderer.CreateTextureFromSurface(windowShape)
    if texture == nil {
        fmt.Fprintf(os.Stderr, "sdl.CreateTextureFromSurface Error: %v\n", sdl.GetError())
        return 1
    }
    defer texture.Destroy()

    renderer.SetDrawColor(255, 100, 100, 100)

    for i := 0; i < 100; i++ {
        renderer.Clear()
        renderer.Copy(texture, nil, nil)
        renderer.Present()
        sdl.Delay(100)
    }

    return 0
}
