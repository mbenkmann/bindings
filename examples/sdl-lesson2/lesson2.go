// https://www.willusher.io/pages/sdl2/
// https://github.com/Twinklebear/TwinklebearDev-Lessons/tree/master/Lesson2

package main

import (
    "fmt"
    "io"
    "os"
    "runtime"

    "winterdrache.de/bindings/sdl"
    "winterdrache.de/bindings/sdlutil"
)

/*
 * Lesson 2: Don't Put Everything in Main
 */

//Screen attributes
const SCREEN_WIDTH = 640
const SCREEN_HEIGHT = 480

/*
 * Log an SDL error with some error message to the output stream of our choice
 * @param w The output stream to write the message too
 * @param msg The error message to write, format will be msg error: SDL_GetError()
 */
func logSDLError(w io.Writer, msg string) {
    fmt.Fprintf(w, "%v error: %v\n", msg, sdl.GetError())
}

/*
 * Loads a BMP image into a texture on the rendering device
 * @param file The BMP image file to load
 * @param ren The renderer to load the texture onto
 * @return the loaded texture, or nullptr if something went wrong.
 */
func loadTexture(file string, ren *sdl.Renderer) *sdl.Texture {
    var texture *sdl.Texture
    //Load the image
    loadedImage := sdl.LoadBMP(file)
    //If the loading went ok, convert to texture and return the texture
    if loadedImage != nil {
        texture = ren.CreateTextureFromSurface(loadedImage)
        loadedImage.Free()
        //Make sure converting went ok too
        if texture == nil {
            logSDLError(os.Stderr, "CreateTextureFromSurface")
        }
    } else {
        logSDLError(os.Stderr, "LoadBMP")
    }
    return texture
}

/*
 * Draw an SDL_Texture to an SDL_Renderer at position x, y, preserving
 * the texture's width and height
 * @param tex The source texture we want to draw
 * @param ren The renderer we want to draw too
 * @param x The x coordinate to draw too
 * @param y The y coordinate to draw too
 */
func renderTexture(tex *sdl.Texture, ren *sdl.Renderer, x int, y int) {
    //Setup the destination rectangle to be at the position we want
    var dst sdl.Rect
    dst.X = x
    dst.Y = y
    //Query the texture to get its width and height to use
    _, _, _, dst.W, dst.H = tex.Query()
    ren.Copy(tex, nil, &dst)
}

func main() {
    os.Exit(sdlmain())
}

func sdlmain() int {
    runtime.LockOSThread() // SDL functions must all be called from the main thread

    //Start up SDL and make sure it went ok
    if sdl.Init(sdl.INIT_VIDEO) != 0 {
        logSDLError(os.Stderr, "SDL_Init")
        return 1
    }
    defer sdl.Quit()

    //Setup our window and renderer
    window := sdl.CreateWindow("Lesson 2", 100, 100, SCREEN_WIDTH, SCREEN_HEIGHT, sdl.WINDOW_SHOWN)
    if window == nil {
        logSDLError(os.Stderr, "CreateWindow")
        return 1
    }
    defer window.Destroy()

    renderer := window.CreateRenderer(-1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
    if renderer == nil {
        logSDLError(os.Stderr, "CreateRenderer")
        return 1
    }
    defer renderer.Destroy()

    //The textures we'll be using
    bgPath := sdlutil.GetResourcePath("../examples/resources", "Lesson2", "background.bmp")
    background := loadTexture(bgPath, renderer)
    if background == nil { return 1 }
    defer background.Destroy()
    imgPath := sdlutil.GetResourcePath("../examples/resources", "Lesson2", "image.bmp")
    image := loadTexture(imgPath, renderer)
    if image == nil { return 1 }
    defer image.Destroy()

    //A sleepy rendering loop, wait for 3 seconds and render and present the screen each time
    for i := 0; i < 3; i++ {
        //Clear the window
        renderer.Clear()

        //Get the width and height from the texture so we know how much to move x,y by
        //to tile it correctly
        _, _, _, bW, bH := background.Query()
        //We want to tile our background so draw it 4 times
        renderTexture(background, renderer, 0, 0)
        renderTexture(background, renderer, bW, 0)
        renderTexture(background, renderer, 0, bH)
        renderTexture(background, renderer, bW, bH)

        //Draw our image in the center of the window
        //We need the foreground image's width to properly compute the position
        //of it's top left corner so that the image will be centered
        _, _, _, iW, iH := image.Query()
        x := SCREEN_WIDTH/2 - iW/2
        y := SCREEN_HEIGHT/2 - iH/2
        renderTexture(image, renderer, x, y)

        //Update the screen
        renderer.Present()
        //Take a quick break after all that hard work
        sdl.Delay(1000)
    }

    return 0
}
