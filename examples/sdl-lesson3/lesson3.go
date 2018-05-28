// https://www.willusher.io/pages/sdl2/
// https://github.com/Twinklebear/TwinklebearDev-Lessons/tree/master/Lesson3

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
 * Lesson 3: SDL Extension Libraries
 */

//Screen attributes
const SCREEN_WIDTH = 640
const SCREEN_HEIGHT = 480

//We'll be scaling our tiles to be 40x40
const TILE_SIZE = 40

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

    //Setup our window and renderer
    window := sdl.CreateWindow("Lesson 3", 100, 100, SCREEN_WIDTH, SCREEN_HEIGHT, sdl.WINDOW_SHOWN)
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
    bgPath := sdlutil.GetResourcePath("../examples/resources", "Lesson3", "background.png")
    background := loadTexture(bgPath, renderer)
    if background == nil { return 1 }
    defer background.Destroy()
    imgPath := sdlutil.GetResourcePath("../examples/resources", "Lesson3", "image.png")
    image := loadTexture(imgPath, renderer)
    if image == nil { return 1 }
    defer image.Destroy()

    //A sleepy rendering loop, wait for 10 seconds and render and present the screen each time
    for i := 0; i < 10; i++ {
        //Clear the window
        renderer.Clear()

        //Determine how many tiles we'll need to fill the screen
        xTiles := SCREEN_WIDTH / TILE_SIZE
        yTiles := SCREEN_HEIGHT / TILE_SIZE

        //Draw the tiles by calculating their positions
        for i := 0; i < xTiles*yTiles; i++ {
            x := i % xTiles
            y := i / xTiles
            renderTextureScaled(background, renderer, x*TILE_SIZE, y*TILE_SIZE, TILE_SIZE, TILE_SIZE)
        }

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
