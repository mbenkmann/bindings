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

// #include "includes.h"
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
func deref_int16_ptr(i *C.int16_t) int16    { return int16(*i) }
func deref_uint8_ptr(i *C.uint8_t) uint8    { return uint8(*i) }
func deref_byte_ptr(i *C.uint8_t) byte      { return byte(*i) }
func deref_uint32_ptr(i *C.uint32_t) uint32 { return uint32(*i) }
func deref_uint64_ptr(i *C.uint64_t) uint64 { return uint64(*i) }
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

// Get a snapshot of the current state of the keyboard.
//
// Returns: An array of key states. Indexes into this array are obtained
// by using SDL_Scancode values. The pointer returned is a pointer to an
// internal SDL array. It will be valid for the whole lifetime of the
// application.
//
// Example:
//   state := SDL_GetKeyboardState();
//   if state[sdl.SCANCODE_RETURN] != 0   {
//       fmt.Printf("<RETURN> is pressed.\n");
//   }
func GetKeyboardState() []byte {
    tmp_numkeys := new(C.int)
    states := C.SDL_GetKeyboardState((*C.int)(tmp_numkeys))
    numkeys := deref_int_ptr(tmp_numkeys)
    return (*((*[999999999]byte)(unsafe.Pointer(states))))[0:numkeys]
}

const check_failed = "Function called with parameters that may call overwriting of random memory"

func checkParametersForSDL_UpdateTexture(texture *Texture, rect *Rect, pixels []byte, pitch int) {
    valid, format, _, w, h := texture.Query()
    if valid < 0 { return } // texture invalid
    if rect != nil {
        w, h = rect.W, rect.H
    }
    line := int(BYTESPERPIXEL(format)) * w

    if line >= 0 && w >= 0 && h >= 0 && pitch >= line && len(pixels) >= pitch*h {
        return
    }

    panic(check_failed)
}

func checkParametersForSDL_UpdateYUVTexture(texture *Texture, rect *Rect, Yplane []byte, Ypitch int, Uplane []byte, Upitch int, Vplane []byte, Vpitch int) {
    valid, _, _, w, h := texture.Query()
    if valid < 0 { return } // texture invalid
    if rect != nil {
        w, h = rect.W, rect.H
    }

    if w >= 0 && h >= 0 && (Ypitch >= 0 && Upitch >= 0 && Vpitch >= 0) &&
        (Ypitch >= w && len(Yplane) >= Ypitch*h) &&
        (Upitch >= w && len(Uplane) >= Upitch*h) &&
        (Vpitch >= w && len(Vplane) >= Vpitch*h) {
        return
    }

    panic(check_failed)
}

func checkParametersForSDL_CreateRGBSurfaceFrom(pixels []byte, width int, height int, depth int, pitch int, Rmask uint32, Gmask uint32, Bmask uint32, Amask uint32) {
    bpp := (depth + 7) >> 3
    line := width * bpp
    if pitch >= 0 && pitch >= line && width >= 0 && height >= 0 && len(pixels) >= pitch*height {
        return
    }
    panic(check_failed)
}

func checkParametersForSDL_CreateRGBSurfaceWithFormatFrom(pixels []byte, width int, height int, depth int, pitch int, format uint32) {
    bpp := int(BYTESPERPIXEL(format))
    line := width * bpp
    if pitch >= 0 && pitch >= line && width >= 0 && height >= 0 && len(pixels) >= pitch*height {
        return
    }
    panic(check_failed)
}

func checkParametersForSDL_RenderReadPixels(renderer *Renderer, rect *Rect, format uint32, pixels []byte, pitch int) {
    r := renderer.GetViewport()
    width := r.W
    height := r.H
    if rect != nil {
        width = rect.W
        height = rect.H
    }

    bpp := int(BYTESPERPIXEL(format))
    line := width * bpp
    if pitch >= 0 && pitch >= line && width >= 0 && height >= 0 && len(pixels) >= pitch*height {
        return
    }
    panic(check_failed)
}

func checkParametersForSDL_ConvertPixels(width int, height int, src_format uint32, src []byte, src_pitch int, dst_format uint32, dst []byte, dst_pitch int) {
    src_bpp := int(BYTESPERPIXEL(src_format))
    dst_bpp := int(BYTESPERPIXEL(dst_format))
    src_line := width * src_bpp
    dst_line := width * dst_bpp
    if width >= 0 && height >= 0 && src_pitch >= src_line && dst_pitch >= dst_line &&
        len(src) >= src_pitch*height && len(dst) >= dst_pitch*height {
        return
    }
    panic(check_failed)
}

// Create a cursor, using the specified bitmap data and mask (in MSB
// format).
//
// The cursor width must be a multiple of 8.
//
// The cursor is created in black and white according to the following:
//   data | mask | resulting pixel on screen
//   0    | 1    | White
//   1    | 1    | Black
//   0    | 0    | Transparent
//   1    | 0    | Inverted color if possible, black if not.
//
// See also: SDL_FreeCursor()
//
func CreateCursor(data []byte, mask []byte, width int, height int, hot_x int, hot_y int) (retval *Cursor) {
    width = (width + 7) &^ 7
    wb := width >> 3
    if len(data) < wb || len(mask) < wb || width <= 0 || height <= 0 {
        return (*Cursor)(C.SDL_CreateCursor(nil, nil, 0, 0, 0, 0))
    }
    if len(data) < len(mask) {
        mask = mask[:len(data)]
    } else {
        data = data[:len(mask)]
    }
    for wb*height > len(data) {
        height--
    }
    retval = (*Cursor)(C.SDL_CreateCursor((*C.Uint8)(&(data[0])), (*C.Uint8)(&(mask[0])), C.int(width), C.int(height), C.int(hot_x), C.int(hot_y)))
    return
}

func DEFINE_PIXELFORMAT(typ, order, layout, bits, bytes uint32) uint32 {
    return ((1 << 28) | ((typ) << 24) | ((order) << 20) | ((layout) << 16) |
        ((bits) << 8) | ((bytes) << 0))
}

func PIXELFLAG(X uint32) uint32 {
    return (((X) >> 28) & 0x0F)
}

func PIXELTYPE(X uint32) uint32    { return (((X) >> 24) & 0x0F) }
func PIXELORDER(X uint32) uint32   { return (((X) >> 20) & 0x0F) }
func PIXELLAYOUT(X uint32) uint32  { return (((X) >> 16) & 0x0F) }
func BITSPERPIXEL(X uint32) uint32 { return (((X) >> 8) & 0xFF) }
func BYTESPERPIXEL(X uint32) uint32 {
    if ISPIXELFORMAT_FOURCC(X) {
        if ((X) == PIXELFORMAT_YUY2) ||
            ((X) == PIXELFORMAT_UYVY) ||
            ((X) == PIXELFORMAT_YVYU) {
            return 2
        } else {
            return 1
        }
    } else {
        return (((X) >> 0) & 0xFF)
    }
}

func ISPIXELFORMAT_INDEXED(format uint32) bool {
    return (!ISPIXELFORMAT_FOURCC(format) &&
        ((PIXELTYPE(format) == PIXELTYPE_INDEX1) ||
            (PIXELTYPE(format) == PIXELTYPE_INDEX4) ||
            (PIXELTYPE(format) == PIXELTYPE_INDEX8)))
}

func ISPIXELFORMAT_PACKED(format uint32) bool {
    return (!ISPIXELFORMAT_FOURCC(format) &&
        ((PIXELTYPE(format) == PIXELTYPE_PACKED8) ||
            (PIXELTYPE(format) == PIXELTYPE_PACKED16) ||
            (PIXELTYPE(format) == PIXELTYPE_PACKED32)))
}
func ISPIXELFORMAT_ARRAY(format uint32) bool {
    return (!ISPIXELFORMAT_FOURCC(format) &&
        ((PIXELTYPE(format) == PIXELTYPE_ARRAYU8) ||
            (PIXELTYPE(format) == PIXELTYPE_ARRAYU16) ||
            (PIXELTYPE(format) == PIXELTYPE_ARRAYU32) ||
            (PIXELTYPE(format) == PIXELTYPE_ARRAYF16) ||
            (PIXELTYPE(format) == PIXELTYPE_ARRAYF32)))
}
func ISPIXELFORMAT_ALPHA(format uint32) bool {
    return ((ISPIXELFORMAT_PACKED(format) &&
        ((PIXELORDER(format) == PACKEDORDER_ARGB) ||
            (PIXELORDER(format) == PACKEDORDER_RGBA) ||
            (PIXELORDER(format) == PACKEDORDER_ABGR) ||
            (PIXELORDER(format) == PACKEDORDER_BGRA))) ||
        (ISPIXELFORMAT_ARRAY(format) &&
            ((PIXELORDER(format) == ARRAYORDER_ARGB) ||
                (PIXELORDER(format) == ARRAYORDER_RGBA) ||
                (PIXELORDER(format) == ARRAYORDER_ABGR) ||
                (PIXELORDER(format) == ARRAYORDER_BGRA))))
}

/* The flag is set to 1 because 0x1? is not in the printable ASCII range */
func ISPIXELFORMAT_FOURCC(format uint32) bool {
    return ((format != 0) && (PIXELFLAG(format) != 1))
}
