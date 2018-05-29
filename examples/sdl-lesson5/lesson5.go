// https://www.willusher.io/pages/sdl2/
// https://github.com/Twinklebear/TwinklebearDev-Lessons/tree/master/Lesson5

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
 * Lesson 5: Clipping Sprite Sheets
 */

//Screen attributes
const (
    SCREEN_WIDTH  = 640
    SCREEN_HEIGHT = 480
)

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
 * Draw an SDL_Texture to an SDL_Renderer at some destination rect
 * taking a clip of the texture if desired
 * @param tex The source texture we want to draw
 * @param rend The renderer we want to draw too
 * @param dst The destination rectangle to render the texture too
 * @param clip The sub-section of the texture to draw (clipping rect)
 *		  nil draws the entire texture
 */
func renderTextureClipScaled(tex *sdl.Texture, ren *sdl.Renderer, dst sdl.Rect, clip *sdl.Rect) {
    ren.Copy(tex, clip, &dst)
}

/*
 * Draw an SDL_Texture to an SDL_Renderer at position x, y, preserving
 * the texture's width and height and taking a clip of the texture if desired
 * If a clip is passed, the clip's width and height will be used instead of the texture's
 * @param tex The source texture we want to draw
 * @param rend The renderer we want to draw too
 * @param x The x coordinate to draw too
 * @param y The y coordinate to draw too
 * @param clip The sub-section of the texture to draw (clipping rect)
 *		  nil draws the entire texture
 */
func renderTexture(tex *sdl.Texture, ren *sdl.Renderer, x int, y int, clip *sdl.Rect) {
    var dst sdl.Rect
    dst.X = x
    dst.Y = y
    if clip != nil {
        dst.W = clip.W
        dst.H = clip.H
    } else {
        _, _, _, dst.W, dst.H = tex.Query()
    }
    renderTextureClipScaled(tex, ren, dst, clip)
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
    window := sdl.CreateWindow("Lesson 5", sdl.WINDOWPOS_CENTERED,
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
    imgPath := sdlutil.GetResourcePath("../examples/resources", "Lesson5", "image.png")
    image := loadTexture(imgPath, renderer)
    if image == nil { return 1 }
    defer image.Destroy()

    //iW and iH are the clip width and height
    //We'll be drawing only clips so get a center position for the w/h of a clip
    iW := 100
    iH := 100
    x := SCREEN_WIDTH/2 - iW/2
    y := SCREEN_HEIGHT/2 - iH/2

    //Setup the clips for our image
    var clips [4]sdl.Rect
    //Since our clips our uniform in size we can generate a list of their
    //positions using some math (the specifics of this are covered in the lesson)
    for i := 0; i < 4; i++ {
        clips[i].X = i / 2 * iW
        clips[i].Y = i % 2 * iH
        clips[i].W = iW
        clips[i].H = iH
    }
    //Specify a default clip to start with
    useClip := 0

    var e sdl.Event
    quit := false
    for !quit {

        //Event Polling
        for sdl.PollEvent(&e) {
            if e.Type() == sdl.QUIT {
                quit = true
            }

            //Use number input to select which clip should be drawn
            if e.Type() == sdl.KEYDOWN {
                switch e.Key().Keysym.Sym {
                    case sdl.K_1:
                    case sdl.K_KP_1:
                        useClip = 0
                    case sdl.K_2:
                    case sdl.K_KP_2:
                        useClip = 1
                    case sdl.K_3:
                    case sdl.K_KP_3:
                        useClip = 2
                    case sdl.K_4:
                    case sdl.K_KP_4:
                        useClip = 3
                    case sdl.K_ESCAPE:
                        quit = true
                }
            }
        }

        //Rendering
        renderer.Clear()
        //Draw the image
        renderTexture(image, renderer, x, y, &clips[useClip])
        //Update the screen
        renderer.Present()
    }

    return 0
}
