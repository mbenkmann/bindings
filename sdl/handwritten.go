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

type RWops C.SDL_RWops
type Surface C.SDL_Surface
type PixelFormat C.SDL_PixelFormat
type Palette C.SDL_Palette

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
