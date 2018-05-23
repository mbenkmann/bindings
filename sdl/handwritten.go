// Copyright (c) 2018 Matthias S. Benkmann
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// This file contains handwritten code for the parts of the SDL bindings that
// cannot reasonably be generated.

// Bindings for Simple DirectMedia Layer (www.libsdl.org)
package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
// #include <SDL_shape.h>
//
// void GoSetError(const char *fmt) {
//     SDL_SetError("%s", fmt);
// }
//
import "C"
import "unsafe"

func deref_int_ptr(i *C.int) int            { return int(*i) }
func deref_float32_ptr(i *C.float) float32  { return float32(*i) }
func deref_uint16_ptr(i *C.uint16_t) uint16 { return uint16(*i) }
func deref_uint8_ptr(i *C.uint8_t) uint8    { return uint8(*i) }
func deref_uint32_ptr(i *C.uint32_t) uint32 { return uint32(*i) }
func bool2bool(b bool) C.SDL_bool {
    if b { return C.SDL_TRUE } else { return C.SDL_FALSE }
}
func freeGoString(s *C.char) string {
    defer C.free(unsafe.Pointer(s))
    return C.GoString(s)
}

func fromC2WindowShapeMode(s C.SDL_WindowShapeMode) WindowShapeMode {
    return WindowShapeMode{ShapeMode(s.mode), WindowShapeParams(s.parameters)}
}

func toCFromWindowShapeMode(s WindowShapeMode) (d C.SDL_WindowShapeMode) {
    d.mode = C.WindowShapeMode(s.Mode)
    d.parameters = C.SDL_WindowShapeParams(s.Parameters)
    return
}

type RWops C.SDL_RWops
type Surface C.SDL_Surface
type PixelFormat C.SDL_PixelFormat
type Palette C.SDL_Palette
type GameControllerButtonBind C.SDL_GameControllerButtonBind
type HapticCustom C.SDL_HapticCustom

// Performs a fast blit from the source surface to the destination surface.
//
// This assumes that the source and destination rectangles are
// the same size.  If either srcrect or dstrect are NULL, the entire
// surface (src or  dst) is copied.  The final blit rectangles are saved
// in srcrect and dstrect after all clipping is performed.
//
// Returns: If the blit is successful, it returns 0, otherwise it returns -1.
//
// The blit function should not be called on a locked surface.
//
// The blit semantics for surfaces with and without blending and colorkey
// are defined as follows:
//
//     RGBA->RGB:
//     Source surface blend mode set to SDL_BLENDMODE_BLEND:
//         alpha-blend (using the source alpha-channel and per-surface alpha)
//         SDL_SRCCOLORKEY ignored.
//     Source surface blend mode set to SDL_BLENDMODE_NONE:
//         copy RGB.
//         if SDL_SRCCOLORKEY set, only copy the pixels matching the
//         RGB values of the source color key, ignoring alpha in the
//         comparison.
//
//     RGB->RGBA:
//     Source surface blend mode set to SDL_BLENDMODE_BLEND:
//         alpha-blend (using the source per-surface alpha)
//     Source surface blend mode set to SDL_BLENDMODE_NONE:
//         copy RGB, set destination alpha to source per-surface alpha value.
//     both:
//         if SDL_SRCCOLORKEY set, only copy the pixels matching the
//         source color key.
//
//     RGBA->RGBA:
//     Source surface blend mode set to SDL_BLENDMODE_BLEND:
//         alpha-blend (using the source alpha-channel and per-surface alpha)
//         SDL_SRCCOLORKEY ignored.
//     Source surface blend mode set to SDL_BLENDMODE_NONE:
//         copy all of RGBA to the destination.
//         if SDL_SRCCOLORKEY set, only copy the pixels matching the
//         RGB values of the source color key, ignoring alpha in the
//         comparison.
//
//     RGB->RGB:
//     Source surface blend mode set to SDL_BLENDMODE_BLEND:
//         alpha-blend (using the source per-surface alpha)
//     Source surface blend mode set to SDL_BLENDMODE_NONE:
//         copy RGB.
//     both:
//         if SDL_SRCCOLORKEY set, only copy the pixels matching the
//         source color key.
//
// You should call SDL_BlitSurface() unless you know exactly how SDL
// blitting works internally and how to use the other blit functions.
func BlitSurface(src *Surface, srcrect Rect, dst *Surface, dstrect Rect) (retval int) {
    return UpperBlit(src, srcrect, dst, dstrect)
}

// Like BlitSurface() but srcrect and dstrect need not be the same size.
func BlitScaled(src *Surface, srcrect Rect, dst *Surface, dstrect Rect) (retval int) {
    return UpperBlitScaled(src, srcrect, dst, dstrect)
}

// Convenience shortcut for LoadBMP_RW(RWFromFile(fpath, "rb"), 1)
func LoadBMP(fpath string) (retval *Surface) {
    return LoadBMP_RW(RWFromFile(fpath, "rb"), 1)
}

// Convenience shortcut for LoadWAV_RW(RWFromFile(fpath, "rb"), 1)
func LoadWAV(fpath string) (retval *AudioSpec, audio_buf []byte) {
    return LoadWAV_RW(RWFromFile(fpath, "rb"), 1)
}

// SDL_SetError supports extra parameters. At present these are not supported.
// If you have a use case, file an issue.
func SetError(fo string) int {
    st := C.CString(fo)
    defer C.free(unsafe.Pointer(st))
    C.GoSetError(st)
    return -1
}

// Checks the event queue for messages and optionally returns them.
//
// If action is SDL_ADDEVENT, up to numevents events will be added to the
// back of the event queue. numevents > len(events) is not an error, but only
// up to len(events) will be added, of course.
//
// If action is SDL_PEEKEVENT, up to numevents events at the front of the
// event queue, with minType <= type <= maxType, will be
// copied to events and will not be removed from the queue.
// numevents > len(events) is not an error, but only
// up to len(events) will be copied.
//
// If action is SDL_GETEVENT, up to numevents events at the front of the
// event queue, with minType <= type <= maxType, will be
// copied to events and will be removed from the queue.
// numevents > len(events) is not an error, but only
// up to len(events) will be copied.
//
// Returns: The number of events actually stored, or -1 if there was an
// error.
//
// This function is thread-safe.
func PeepEvents(events []Event, numevents int, action Eventaction, minType uint32, maxType uint32) int {
    if numevents > len(events) {
        numevents = len(events)
    }
    if numevents <= 0 {
        return 0
    }
    return int(C.SDL_PeepEvents((*C.SDL_Event)(&(events[0])), C.int(numevents), C.SDL_eventaction(action), C.Uint32(minType), C.Uint32(maxType)))
}

// Calculate a minimal rectangle enclosing a set of points.
//
// Returns: SDL_TRUE if any points were within the clipping rect
//
func EnclosePoints(points []Point, clip Rect) (retval bool, result Rect) {
    pts := make([]C.SDL_Point, len(points)+1) // +1 to make sure pts[0] does not trip bounds checking
    for i := range points {
        pts[i] = toCFromPoint(points[i])
    }
    tmp_clip := toCFromRect(clip)
    tmp_result := toCFromRect(result)
    retval = C.SDL_TRUE == C.SDL_EnclosePoints(&(pts[0]), C.int(len(points)), &tmp_clip, &tmp_result)
    return
}

// Draw multiple points on the current rendering target.
//
// Returns: 0 on success, or -1 on error
//
//   renderer
//     The renderer which should draw multiple points.
//
//   points
//     The points to draw
//
func (renderer *Renderer) DrawPoints(points []Point) (retval int) {
    pts := make([]C.SDL_Point, len(points)+1) // +1 to make sure pts[0] does not trip bounds checking
    for i := range points {
        pts[i] = toCFromPoint(points[i])
    }
    retval = int(C.SDL_RenderDrawPoints((*C.SDL_Renderer)(renderer), &(pts[0]), C.int(len(points))))
    return
}

// Draw a series of connected lines on the current rendering target.
//
// Returns: 0 on success, or -1 on error
//
//   renderer
//     The renderer which should draw multiple lines.
//
//   points
//     The points along the lines
//
func (renderer *Renderer) DrawLines(points []Point) (retval int) {
    pts := make([]C.SDL_Point, len(points)+1) // +1 to make sure pts[0] does not trip bounds checking
    for i := range points {
        pts[i] = toCFromPoint(points[i])
    }
    retval = int(C.SDL_RenderDrawLines((*C.SDL_Renderer)(renderer), &(pts[0]), C.int(len(points))))
    return
}

// Draw some number of rectangles on the current rendering target.
//
// Returns: 0 on success, or -1 on error
//
//   renderer
//     The renderer which should draw multiple rectangles.
//
//   rects
//     A pointer to an array of destination rectangles.
//
func (renderer *Renderer) DrawRects(rects []Rect) (retval int) {
    rcts := make([]C.SDL_Rect, len(rects)+1) // +1 to make sure rcts[0] does not trip bounds checking
    for i := range rects {
        rcts[i] = toCFromRect(rects[i])
    }
    retval = int(C.SDL_RenderDrawRects((*C.SDL_Renderer)(renderer), &(rcts[0]), C.int(len(rects))))
    return
}

// Fill some number of rectangles on the current rendering target with
// the drawing color.
//
// Returns: 0 on success, or -1 on error
//
//   renderer
//     The renderer which should fill multiple rectangles.
//
//   rects
//     A pointer to an array of destination rectangles.
//
func (renderer *Renderer) FillRects(rects []Rect) (retval int) {
    rcts := make([]C.SDL_Rect, len(rects)+1) // +1 to make sure rcts[0] does not trip bounds checking
    for i := range rects {
        rcts[i] = toCFromRect(rects[i])
    }
    retval = int(C.SDL_RenderFillRects((*C.SDL_Renderer)(renderer), &(rcts[0]), C.int(len(rects))))
    return
}

// Set a range of colors in a palette.
//
// Returns: 0 on success, or -1 if not all of the colors could be set.
//
//   palette
//     The palette to modify.
//
//   colors
//     An array of colors to copy into the palette.
//
//   firstcolor
//     The index of the first palette entry to modify.
//
func (palette *Palette) SetColors(colors []Color, firstcolor int) (retval int) {
    cols := make([]C.SDL_Color, len(colors)+1) // +1 to make sure cols[0] does not trip bounds checking
    for i := range colors {
        cols[i] = toCFromColor(colors[i])
    }
    retval = int(C.SDL_SetPaletteColors((*C.SDL_Palette)(palette), (*C.SDL_Color)(&(cols[0])), C.int(firstcolor), C.int(len(colors))))
    return
}

// Copy a number of rectangles on the window surface to the screen.
//
// Returns: 0 on success, or -1 on error.
//
// See also: SDL_GetWindowSurface()
//
// See also: SDL_UpdateWindowSurfaceRect()
//
func (window *Window) UpdateSurfaceRects(rects []Rect) (retval int) {
    rcts := make([]C.SDL_Rect, len(rects)+1) // +1 to make sure rcts[0] does not trip bounds checking
    for i := range rects {
        rcts[i] = toCFromRect(rects[i])
    }
    retval = int(C.SDL_UpdateWindowSurfaceRects((*C.SDL_Window)(window), (*C.SDL_Rect)(&(rcts[0])), C.int(len(rects))))
    return
}

// Performs a fast fill of the given rectangles with color.
//
// The color should be a pixel of the format used by the surface, and can
// be generated by the SDL_MapRGB() function.
//
// Returns: 0 on success, or -1 on error.
//
func (dst *Surface) FillRects(rects []Rect, color uint32) (retval int) {
    rcts := make([]C.SDL_Rect, len(rects)+1) // +1 to make sure rcts[0] does not trip bounds checking
    for i := range rects {
        rcts[i] = toCFromRect(rects[i])
    }
    retval = int(C.SDL_FillRects((*C.SDL_Surface)(dst), (*C.SDL_Rect)(&(rcts[0])), C.int(len(rects)), C.Uint32(color)))
    return
}

// Calculate the intersection of a rectangle and line segment.
//
// Returns: SDL_TRUE if there is an intersection, SDL_FALSE otherwise.
//
func IntersectRectAndLine(rect Rect, lineX1 int, lineY1 int, lineX2 int, lineY2 int) (retval bool, X1 int, Y1 int, X2 int, Y2 int) {
    tmp_rect := toCFromRect(rect)
    tmp_X1 := C.int(lineX1)
    tmp_Y1 := C.int(lineY1)
    tmp_X2 := C.int(lineX2)
    tmp_Y2 := C.int(lineY2)
    retval = C.SDL_TRUE == C.SDL_IntersectRectAndLine(&tmp_rect, &tmp_X1, &tmp_Y1, &tmp_X2, &tmp_Y2)
    X1 = int(tmp_X1)
    Y1 = int(tmp_Y1)
    X2 = int(tmp_X2)
    Y2 = int(tmp_Y2)
    return
}

// Set the gamma ramp for a window.
//
// Returns: 0 on success, or -1 if gamma ramps are unsupported.
//
//   window
//     The window for which the gamma ramp should be set.
//
//   red
//     The translation table for the red channel, or NULL.
//
//   green
//     The translation table for the green channel, or NULL.
//
//   blue
//     The translation table for the blue channel, or NULL.
//
// Set the gamma translation table for the red, green, and blue channels
// of the video hardware. Each table is an array of 256 16-bit
// quantities, representing a mapping between the input and output for
// that channel. The input is the index into the array, and the output is
// the 16-bit gamma value at that index, scaled to the output color
// precision.
//
// See also: SDL_GetWindowGammaRamp()
//
func (window *Window) SetGammaRamp(red *[256]uint16, green *[256]uint16, blue *[256]uint16) (retval int) {
    retval = int(C.SDL_SetWindowGammaRamp((*C.SDL_Window)(window), (*C.Uint16)(unsafe.Pointer(red)), (*C.Uint16)(unsafe.Pointer(green)), (*C.Uint16)(unsafe.Pointer(blue))))
    return
}

// Get the gamma ramp for a window.
//
// Returns: 0 on success, or -1 if gamma ramps are unsupported.
//
// See also: SDL_SetWindowGammaRamp()
//
//   window
//     The window from which the gamma ramp should be queried.
//
//   red
//     A pointer to a 256 element array of 16-bit quantities to hold the
//     translation table for the red channel, or NULL.
//
//   green
//     A pointer to a 256 element array of 16-bit quantities to hold the
//     translation table for the green channel, or NULL.
//
//   blue
//     A pointer to a 256 element array of 16-bit quantities to hold the
//     translation table for the blue channel, or NULL.
//
func (window *Window) GetGammaRamp(red *[256]uint16, green *[256]uint16, blue *[256]uint16) (retval int) {
    retval = int(C.SDL_SetWindowGammaRamp((*C.SDL_Window)(window), (*C.Uint16)(unsafe.Pointer(red)), (*C.Uint16)(unsafe.Pointer(green)), (*C.Uint16)(unsafe.Pointer(blue))))
    return
}

// Allocate and free an RGB surface.
//
// If the depth is 4 or 8 bits, an empty palette is allocated for the
// surface. If the depth is greater than 8 bits, the pixel format is set
// using the '[RGBA]mask'.
//
// If the function runs out of memory, it will return NULL.
//
//   width
//     The width in pixels of the surface to create.
//
//   height
//     The height in pixels of the surface to create.
//
//   depth
//     The depth in bits of the surface to create.
//
//   pitch
//     Number of bytes per row of pixels
//
//   Rmask
//     The red mask of the surface to create.
//
//   Gmask
//     The green mask of the surface to create.
//
//   Bmask
//     The blue mask of the surface to create.
//
//   Amask
//     The alpha mask of the surface to create.
//
func CreateRGBSurfaceFrom(pixels []byte, width int, height int, depth int, pitch int, Rmask uint32, Gmask uint32, Bmask uint32, Amask uint32) (retval *Surface) {
    retval = (*Surface)(unsafe.Pointer(C.SDL_CreateRGBSurfaceFrom(unsafe.Pointer(&(pixels[0])), C.int(width), C.int(height), C.int(depth), C.int(pitch), C.Uint32(Rmask), C.Uint32(Gmask), C.Uint32(Bmask), C.Uint32(Amask))))
    return
}

// Copy a block of pixels of one format to another format.
//
// Returns: 0 on success, or -1 if there was an error
//
func ConvertPixels(width int, height int, src_format uint32, src []byte, src_pitch int, dst_format uint32, dst []byte, dst_pitch int) (retval int) {
    retval = int(C.SDL_ConvertPixels(C.int(width), C.int(height), C.Uint32(src_format), unsafe.Pointer(&(src[0])), C.int(src_pitch), C.Uint32(dst_format), unsafe.Pointer(&(dst[0])), C.int(dst_pitch)))
    return
}

// Calculate a 256 entry gamma ramp for a gamma value.
func CalculateGammaRamp(gamma float32, ramp *[256]uint16) {
    C.SDL_CalculateGammaRamp(C.float(gamma), (*C.Uint16)(unsafe.Pointer(ramp)))
    return
}

// Create a window and default renderer.
//
// Returns: 0 on success, or -1 on error
//
//   width
//     The width of the window
//
//   height
//     The height of the window
//
//   window_flags
//     The flags used to create the window
//
//   window
//     A pointer filled with the window, or NULL on error
//
//   renderer
//     A pointer filled with the renderer, or NULL on error
//
func CreateWindowAndRenderer(width int, height int, window_flags uint32) (retval int, window *Window, renderer *Renderer) {
    var tmp_window *C.SDL_Window
    var tmp_renderer *C.SDL_Renderer
    retval = int(C.SDL_CreateWindowAndRenderer(C.int(width), C.int(height), C.Uint32(window_flags), &tmp_window, &tmp_renderer))
    window = (*Window)(unsafe.Pointer(tmp_window))
    renderer = (*Renderer)(unsafe.Pointer(tmp_renderer))
    return
}

// Update the given texture rectangle with new pixel data.
//
// Returns: 0 on success, or -1 if the texture is not valid.
//
// Note: This is a fairly slow function.
//
//   texture
//     The texture to update
//
//   rect
//     A pointer to the rectangle of pixels to update, or NULL to update the
//     entire texture.
//
//   pixels
//     The raw pixel data.
//
//   pitch
//     The number of bytes in a row of pixel data, including padding between
//     lines.
//
func (texture *Texture) Update(rect *Rect, pixels []byte, pitch int) (retval int) {
    var tmp_rect *C.SDL_Rect
    if rect != nil {
        tmp_rect2 := toCFromRect(*rect)
        tmp_rect = &tmp_rect2
    }
    retval = int(C.SDL_UpdateTexture((*C.SDL_Texture)(texture), tmp_rect, unsafe.Pointer(&(pixels[0])), C.int(pitch)))
    return
}

// Update a rectangle within a planar YV12 or IYUV texture with new pixel
// data.
//
// Returns: 0 on success, or -1 if the texture is not valid.
//
// Note: You can use SDL_UpdateTexture() as long as your pixel data is a
// contiguous block of Y and U/V planes in the proper order, but this
// function is available if your pixel data is not contiguous.
//
//   texture
//     The texture to update
//
//   rect
//     A pointer to the rectangle of pixels to update, or NULL to update the
//     entire texture.
//
//   Yplane
//     The raw pixel data for the Y plane.
//
//   Ypitch
//     The number of bytes between rows of pixel data for the Y plane.
//
//   Uplane
//     The raw pixel data for the U plane.
//
//   Upitch
//     The number of bytes between rows of pixel data for the U plane.
//
//   Vplane
//     The raw pixel data for the V plane.
//
//   Vpitch
//     The number of bytes between rows of pixel data for the V plane.
//
func (texture *Texture) UpdateYUV(rect *Rect, Yplane []byte, Ypitch int, Uplane []byte, Upitch int, Vplane []byte, Vpitch int) (retval int) {
    var tmp_rect *C.SDL_Rect
    if rect != nil {
        tmp_rect2 := toCFromRect(*rect)
        tmp_rect = &tmp_rect2
    }
    retval = int(C.SDL_UpdateYUVTexture((*C.SDL_Texture)(texture), tmp_rect, (*C.Uint8)(unsafe.Pointer(&(Yplane[0]))), C.int(Ypitch), (*C.Uint8)(unsafe.Pointer(&(Uplane[0]))), C.int(Upitch), (*C.Uint8)(unsafe.Pointer(&(Vplane[0]))), C.int(Vpitch)))
    return
}

// Lock a portion of the texture for write-only pixel access.
//
// Returns: 0 on success, or -1 if the texture is not valid or was not
// created with SDL_TEXTUREACCESS_STREAMING.
//
// See also: SDL_UnlockTexture()
//
//   texture
//     The texture to lock for access, which was created with
//     SDL_TEXTUREACCESS_STREAMING.
//
//   rect
//     A pointer to the rectangle to lock for access. If the rect is NULL,
//     the entire texture will be locked.
//
//   pixels
//     This is filled in with a pointer to the locked pixels, appropriately
//     offset by the locked area.
//
//   pitch
//     This is filled in with the pitch of the locked pixels.
//
func (texture *Texture) Lock(rect *Rect) (retval int, pixels *[999999999]byte, pitch int) {
    var tmp_rect *C.SDL_Rect
    if rect != nil {
        tmp_rect2 := toCFromRect(*rect)
        tmp_rect = &tmp_rect2
    }
    tmp_pitch := new(C.int)
    var pix unsafe.Pointer
    retval = int(C.SDL_LockTexture((*C.SDL_Texture)(texture), tmp_rect, &pix, (*C.int)(tmp_pitch)))
    pitch = deref_int_ptr(tmp_pitch)
    pixels = (*[999999999]byte)(pix)
    return
}

// MessageBox structure containing title, text, window, etc.
type MessageBoxData struct {
    // SDL_MessageBoxFlags
    Flags uint32

    // Parent window, can be NULL
    Window *Window

    // UTF-8 title
    Title string

    // UTF-8 message text
    Message string

    Buttons []MessageBoxButtonData

    // SDL_MessageBoxColorScheme, can be NULL to use system settings
    ColorScheme *MessageBoxColorScheme
}

// Create a modal message box.
//
// Returns: -1 on error, otherwise 0 and buttonid contains user id of
// button hit or -1 if dialog was closed.
//
// Note: This function should be called on the thread that created the
// parent window, or on the main thread if the messagebox has no parent.
// It will block execution of that thread until the user clicks a button
// or closes the messagebox.
//
//   messageboxdata
//     The SDL_MessageBoxData structure with title, text, etc.
//
//   buttonid
//     The pointer to which user id of hit button should be copied.
//
func ShowMessageBox(messageboxdata *MessageBoxData) (retval int, buttonid int) {
    if messageboxdata == nil { return -1, -1 }
    mbox := new(C.SDL_MessageBoxData)
    mbox.flags = C.Uint32(messageboxdata.Flags)
    mbox.window = (*C.SDL_Window)(messageboxdata.Window)
    tmp_title := C.CString(messageboxdata.Title)
    defer C.free(unsafe.Pointer(tmp_title))
    tmp_message := C.CString(messageboxdata.Message)
    defer C.free(unsafe.Pointer(tmp_message))
    mbox.title = tmp_title
    mbox.message = tmp_message
    tmp_buttonid := new(C.int)
    if messageboxdata.ColorScheme != nil {
        mbox.colorScheme = new(C.SDL_MessageBoxColorScheme)
        for i, col := range messageboxdata.ColorScheme.Colors {
            mbox.colorScheme.colors[i] = toCFromMessageBoxColor(col)
        }
    }
    mbox.numbuttons = C.int(len(messageboxdata.Buttons))
    buttons := make([]C.SDL_MessageBoxButtonData, len(messageboxdata.Buttons))
    for i, butt := range messageboxdata.Buttons {
        buttons[i] = toCFromMessageBoxButtonData(butt)
    }
    defer func() {
        for i := range buttons {
            C.free(unsafe.Pointer(buttons[i].text))
        }
    }()
    mbox.buttons = &(buttons[0])
    retval = int(C.SDL_ShowMessageBox(mbox, (*C.int)(tmp_buttonid)))
    buttonid = deref_int_ptr(tmp_buttonid)
    return
}

// A structure to hold a set of audio conversion filters and buffers.
type AudioCVT struct {
    // Set to 1 if conversion possible
    Needed int

    // Source audio format
    Src_format AudioFormat

    // Target audio format
    Dst_format AudioFormat

    // Rate conversion increment (dst_rate/src_rate)
    Rate_incr float64

    // Buffer to hold entire audio data
    Buf []byte

    cvt C.SDL_AudioCVT
}

// This function takes a source format and rate and a destination format
// and rate, and initializes the cvt structure with information needed by
// SDL_ConvertAudio() to convert a buffer of audio data from one format
// to the other.
//
// Returns: -1 if the format conversion is not supported, 0 if there's no
// conversion needed, or 1 if the audio filter is set up.
//
func BuildAudioCVT(src_format AudioFormat, src_channels uint8, src_rate int, dst_format AudioFormat, dst_channels uint8, dst_rate int) (retval int, cvt *AudioCVT) {
    cvt = new(AudioCVT)
    retval = int(C.SDL_BuildAudioCVT((*C.SDL_AudioCVT)(&cvt.cvt), C.SDL_AudioFormat(src_format), C.Uint8(src_channels), C.int(src_rate), C.SDL_AudioFormat(dst_format), C.Uint8(dst_channels), C.int(dst_rate)))
    cvt.Needed = int(cvt.cvt.needed)
    cvt.Src_format = AudioFormat(cvt.cvt.src_format)
    cvt.Dst_format = AudioFormat(cvt.cvt.dst_format)
    cvt.Rate_incr = float64(cvt.cvt.rate_incr)
    return
}

// Once you have initialized the cvt structure using sdl.BuildAudioCVT(),
// and filled in cvt.Buf of audio data in the source format, this function
// will convert it to the desired format.
//
// The data conversion may expand or shrink the size of the audio data in
// cvt.Buf.
func ConvertAudio(cvt *AudioCVT) (retval int) {
    cvt.cvt.len = C.int(len(cvt.Buf))
    if cvt.cvt.len_mult > 1 {
        new_len := len(cvt.Buf) * int(cvt.cvt.len_mult)
        new_buf := make([]byte, new_len)
        copy(new_buf, cvt.Buf)
        cvt.Buf = new_buf
    }
    cvt.cvt.buf = (*C.Uint8)(&(cvt.Buf[0]))
    retval = int(C.SDL_ConvertAudio((*C.SDL_AudioCVT)(&cvt.cvt)))
    cvt.Buf = cvt.Buf[0:int(cvt.cvt.len_cvt)]
    return
}

// This function loads a WAVE from the data source, automatically freeing
// that source if freesrc is non-zero. For example, to load a WAVE file,
// you could do:
//   SDL_LoadWAV_RW(SDL_RWFromFile("sample.wav", "rb"), 1, ...);
//
// If this function succeeds, it returns an SDL_AudioSpec, filled
// with the audio data format of the wave data, and sets audio_buf to a
// buffer containing the audio data
//
// This function returns NULL and sets the SDL error message if the wave
// file cannot be opened, uses an unknown data format, or is corrupt.
// Currently raw and MS-ADPCM WAVE files are supported.
func LoadWAV_RW(src *RWops, freesrc int) (retval *AudioSpec, audio_buf []byte) {
    tmp_spec := new(C.SDL_AudioSpec)
    tmp_audio_buf := new(C.Uint8)
    audio_len := new(C.Uint32)
    tmp_retval := C.SDL_LoadWAV_RW((*C.SDL_RWops)(src), C.int(freesrc), tmp_spec, &tmp_audio_buf, audio_len)
    if tmp_retval != nil {
        defer C.SDL_FreeWAV(tmp_audio_buf)
        tr := fromC2AudioSpec(*tmp_retval)
        retval = &tr
        audio_buf = make([]byte, *audio_len)
        copy(audio_buf, ((*[999999999]byte)(unsafe.Pointer(tmp_audio_buf)))[0:999999999])
    }
    return
}

// This takes two audio buffers of the playing audio format and mixes
// them, performing addition, volume adjustment, and overflow clipping.
// The volume ranges from 0 - 128, and should be set to SDL_MIX_MAXVOLUME
// for full audio volume. Note this does not change hardware volume. This
// is provided for convenience -- you can mix your own audio data.
//
// Note: If src and dst have different lengths, the shorter length determines
// what will be mixed.
func MixAudio(dst, src []byte, volume int) {
    l := len(dst)
    if len(src) < l { l = len(src) }
    if l == 0       { return } // make sure dst[0] and src[0] don't trip range check
    C.SDL_MixAudio((*C.Uint8)(unsafe.Pointer(&(dst[0]))), (*C.Uint8)(unsafe.Pointer(&(src[0]))), C.Uint32(l), C.int(volume))
}

// This works like SDL_MixAudio(), but you specify the audio format
// instead of using the format of audio device 1. Thus it can be used
// when no audio device is open at all.
func MixAudioFormat(dst, src []byte, format AudioFormat, volume int) {
    l := len(dst)
    if len(src) < l { l = len(src) }
    if l == 0       { return } // make sure dst[0] and src[0] don't trip range check
    C.SDL_MixAudioFormat((*C.Uint8)(unsafe.Pointer(&(dst[0]))), (*C.Uint8)(unsafe.Pointer(&(src[0]))), C.SDL_AudioFormat(format), C.Uint32(l), C.int(volume))
}

// Queue more audio on non-callback devices.
//
// SDL offers two ways to feed audio to the device: you can either supply
// a callback that SDL triggers with some frequency to obtain more audio
// (pull method), or you can supply no callback, and then SDL will expect
// you to supply data at regular intervals (push method) with this
// function.
//
// There are no limits on the amount of data you can queue, short of
// exhaustion of address space. Queued data will drain to the device as
// necessary without further intervention from you. If the device needs
// audio but there is not enough queued, it will play silence to make up
// the difference. This means you will have skips in your audio playback
// if you aren't routinely queueing sufficient data.
//
// This function copies the supplied data, so you are safe to free it
// when the function returns. This function is thread-safe, but queueing
// to the same device from two threads at once does not promise which
// buffer will be queued first.
//
// You may not queue audio on a device that is using an application-
// supplied callback; doing so returns an error. You have to use the
// audio callback or queue audio with this function, but not both.
//
// You should not call SDL_LockAudio() on the device before queueing; SDL
// handles locking internally for this function.
//
// Returns: zero on success, -1 on error.
//
// See also: SDL_GetQueuedAudioSize
//
// See also: SDL_ClearQueuedAudio
//
//   dev
//     The device ID to which we will queue audio.
//
//   data
//     The data to queue to the device for later playback.
//
//   len
//     The number of bytes (not samples!) to which (data) points.
//
func QueueAudio(dev AudioDeviceID, data []byte) (retval int) {
    if len(data) == 0 { return 0 } // make sure data[0] does not trip range check
    retval = int(C.SDL_QueueAudio(C.SDL_AudioDeviceID(dev), unsafe.Pointer(&(data[0])), C.Uint32(len(data))))
    return
}
