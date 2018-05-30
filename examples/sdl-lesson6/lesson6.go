// https://www.willusher.io/pages/sdl2/
// https://github.com/Twinklebear/TwinklebearDev-Lessons/tree/master/Lesson6

package main

import (
    "fmt"
    "io"
    "os"
    "runtime"

    "winterdrache.de/bindings/sdl"
    "winterdrache.de/bindings/sdlutil"
    "winterdrache.de/bindings/ttf"
)

/*
 * Lesson 6: True Type Fonts with SDL_ttf
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

/*
 * Render the message we want to display to a texture for drawing
 * @param message The message we want to display
 * @param fontFile The font we want to use to render the text
 * @param color The color we want the text to be
 * @param fontSize The size we want the font to be
 * @param renderer The renderer to load the texture in
 * @return An sdl.Texture containing the rendered message, or nil if something went wrong
 */
func renderText(message string, fontFile string, color sdl.Color,
    fontSize int, renderer *sdl.Renderer) *sdl.Texture {
    //Open the font
    font := ttf.OpenFont(fontFile, fontSize)
    if font == nil {
        logSDLError(os.Stderr, "ttf.OpenFont")
        return nil
    }
    defer font.Close()

    //We need to first render to a surface as that's what font.RenderText returns, then
    //load that surface into a texture
    surf := font.RenderUTF8_Blended(message, color)
    if surf == nil {
        logSDLError(os.Stderr, "ttf.Font.RenderText_Blended")
        return nil
    }
    defer surf.Free()

    texture := renderer.CreateTextureFromSurface(surf)
    if texture == nil {
        logSDLError(os.Stderr, "CreateTextureFromSurface")
        return nil
    }
    // don't defer texture.Destory()! We're returning it!

    return texture
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

    //Also need to init SDL_ttf
    if ttf.Init() != 0 {
        logSDLError(os.Stderr, "ttf.Init")
        return 1
    }
    defer ttf.Quit()

    //Setup our window and renderer, this time let's put our window in the center
    //of the screen
    window := sdl.CreateWindow("Lesson 6", sdl.WINDOWPOS_CENTERED,
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

    fontPath := sdlutil.GetResourcePath("../examples/resources", "Lesson6", "sample.ttf")
    //We'll render the string "TTF fonts are cool!" in white
    //Color is in RGB format
    color := sdl.Color{255, 255, 255, 255}
    image := renderText("TTF fonts are cool!", fontPath, color, 64, renderer)
    if image == nil {
        // error message has been output by renderText()
        return 1
    }
    defer image.Destroy()

    //Get the texture w/h so we can center it in the screen
    _, _, _, iW, iH := image.Query()
    x := SCREEN_WIDTH/2 - iW/2
    y := SCREEN_HEIGHT/2 - iH/2

    var e sdl.Event
    quit := false
    for !quit {

        //Event Polling
        for sdl.PollEvent(&e) {
            if e.Type() == sdl.QUIT {
                quit = true
            }

            if e.Type() == sdl.KEYDOWN && e.Key().Keysym.Sym == sdl.K_ESCAPE {
                quit = true
            }
        }

        //Rendering
        renderer.Clear()
        //We can draw our message as we do any other texture, since it's been
        //rendered to a texture
        renderTexture(image, renderer, x, y, nil)
        //Update the screen
        renderer.Present()
    }

    return 0
}
