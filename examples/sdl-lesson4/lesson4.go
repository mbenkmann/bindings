// https://www.willusher.io/pages/sdl2/
// https://github.com/Twinklebear/TwinklebearDev-Lessons/tree/master/Lesson4

package main

import (
    "fmt"
    "io"
    "os"
    "runtime"

    "winterdrache.de/bindings/img"
    "winterdrache.de/bindings/sdl"
    "winterdrache.de/bindings/sdlutil"
)

/*
 * Lesson 4: Handling Events
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
 * Loads an image into a texture on the rendering device
 * @param file The image file to load
 * @param ren The renderer to load the texture onto
 * @return the loaded texture, or nullptr if something went wrong.
 */
func loadTexture(file string, ren *sdl.Renderer) *sdl.Texture {
    texture := img.LoadTexture(ren, file)
    if texture == nil {
        logSDLError(os.Stderr, "LoadTexture")
    }
    return texture
}

/*
 * Draw an SDL_Texture to an SDL_Renderer at position x, y, with some desired
 * width and height
 * @param tex The source texture we want to draw
 * @param rend The renderer we want to draw too
 * @param x The x coordinate to draw too
 * @param y The y coordinate to draw too
 * @param w The width of the texture to draw
 * @param h The height of the texture to draw
 */
func renderTextureScaled(tex *sdl.Texture, ren *sdl.Renderer, x int, y int, w int, h int) {
    //Set up the destination rectangle to be at the position we want
    var dst sdl.Rect
    dst.X = x
    dst.Y = y
    dst.W = w
    dst.H = h
    ren.Copy(tex, nil, &dst)
}

/*
 * Draw an SDL_Texture to an SDL_Renderer at position x, y, preserving
 * the texture's width and height
 * @param tex The source texture we want to draw
 * @param rend The renderer we want to draw too
 * @param x The x coordinate to draw too
 * @param y The y coordinate to draw too
 */
func renderTexture(tex *sdl.Texture, ren *sdl.Renderer, x int, y int) {
    _, _, _, w, h := tex.Query()
    renderTextureScaled(tex, ren, x, y, w, h)
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
    // img.Init() is optional. It will happen automatically on first use.
    defer img.Quit()

    //Setup our window and renderer, this time let's put our window in the center
    //of the screen
    window := sdl.CreateWindow("Lesson 4", sdl.WINDOWPOS_CENTERED,
        sdl.WINDOWPOS_CENTERED, SCREEN_WIDTH, SCREEN_HEIGHT, sdl.WINDOW_SHOWN)
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

    //The texture we'll be using
    imgPath := sdlutil.GetResourcePath("../examples/resources", "Lesson4", "image.png")
    image := loadTexture(imgPath, renderer)
    if image == nil { return 1 }
    defer image.Destroy()

    //Our texture size won't change, so we can get it here
    //instead of constantly allocating/deleting ints in the loop
    _, _, _, iW, iH := image.Query()
    x := SCREEN_WIDTH/2 - iW/2
    y := SCREEN_HEIGHT/2 - iH/2

    //Our event structure
    var e sdl.Event

    //For tracking if we want to quit
    quit := false
    for !quit {

        //Read any events that occured, for now we'll just quit if any event occurs
        for sdl.PollEvent(&e) {
            //If user closes the window
            if e.Type() == sdl.QUIT {
                quit = true
            }
            //If user presses any key
            if e.Type() == sdl.KEYDOWN {
                quit = true
            }
            //If user clicks the mouse
            if e.Type() == sdl.MOUSEBUTTONDOWN {
                quit = true
            }
        }

        //Rendering
        renderer.Clear()
        //Draw the image
        renderTexture(image, renderer, x, y)
        //Update the screen
        renderer.Present()
    }

    return 0
}
