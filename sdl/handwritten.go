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

func BlitScaled(src *Surface, srcrect Rect, dst *Surface, dstrect Rect) (retval int) {
    return UpperBlitScaled(src, srcrect, dst, dstrect)
}

func BlitSurface(src *Surface, srcrect Rect, dst *Surface, dstrect Rect) (retval int) {
    return UpperBlit(src, srcrect, dst, dstrect)
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
