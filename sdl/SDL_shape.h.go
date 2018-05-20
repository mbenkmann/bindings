// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
// #include <SDL_shape.h>
import "C"
import "unsafe"

 // Header file for the shaped window API.

 // A struct that tags the SDL_WindowShapeParams union with an enum
 // describing the type of its contents.
type WindowShapeMode struct {
     // The mode of these window-shape parameters.
    Mode ShapeMode

     // Window-shape parameters.
    Parameters WindowShapeParams
}

 // A union containing parameters for shaped windows.
type WindowShapeParams C.SDL_WindowShapeParams

 // a cutoff alpha value for binarization of the window shape's alpha
 // channel.
func (u *WindowShapeParams) BinarizationCutoff() uint8 {
    p := (*C.Uint8)(unsafe.Pointer(u))
    return uint8(*p)
}
 // a cutoff alpha value for binarization of the window shape's alpha
 // channel.
func (u *WindowShapeParams) SetBinarizationCutoff(x uint8) {
    p := (*C.Uint8)(unsafe.Pointer(u))
    *p = C.Uint8(x)
}

func (u *WindowShapeParams) ColorKey() Color {
    p := (*C.SDL_Color)(unsafe.Pointer(u))
    return Color{uint8(p.r), uint8(p.g), uint8(p.b), uint8(p.a)}
}
func (u *WindowShapeParams) SetColorKey(x Color) {
    p := (*C.SDL_Color)(unsafe.Pointer(u))
    p.r = C.Uint8(x.R)
    p.g = C.Uint8(x.G)
    p.b = C.Uint8(x.B)
    p.a = C.Uint8(x.A)
}

const (
    NONSHAPEABLE_WINDOW = C.SDL_NONSHAPEABLE_WINDOW

    INVALID_SHAPE_ARGUMENT = C.SDL_INVALID_SHAPE_ARGUMENT

    WINDOW_LACKS_SHAPE = C.SDL_WINDOW_LACKS_SHAPE
)

 // An enum denoting the specific type of contents present in an
 // SDL_WindowShapeParams union.
type ShapeMode int
const (
     // The default mode, a binarized alpha cutoff of 1.
    ShapeModeDefault ShapeMode = C.ShapeModeDefault

     // A binarized alpha cutoff with a given integer value.
    ShapeModeBinarizeAlpha ShapeMode = C.ShapeModeBinarizeAlpha

     // A binarized alpha cutoff with a given integer value, but with the
     // opposite comparison.
    ShapeModeReverseBinarizeAlpha ShapeMode = C.ShapeModeReverseBinarizeAlpha

     // A color key is applied.
    ShapeModeColorKey ShapeMode = C.ShapeModeColorKey
)


 // Create a window that can be shaped with the specified position,
 // dimensions, and flags.
 // 
 // Returns: The window created, or NULL if window creation failed.
 // 
 // See also: SDL_DestroyWindow()
 // 
 //   title
 //     The title of the window, in UTF-8 encoding.
 //   
 //   x
 //     The x position of the window, SDL_WINDOWPOS_CENTERED, or
 //     SDL_WINDOWPOS_UNDEFINED.
 //   
 //   y
 //     The y position of the window, SDL_WINDOWPOS_CENTERED, or
 //     SDL_WINDOWPOS_UNDEFINED.
 //   
 //   w
 //     The width of the window.
 //   
 //   h
 //     The height of the window.
 //   
 //   flags
 //     The flags for the window, a mask of SDL_WINDOW_BORDERLESS with any of
 //     the following: SDL_WINDOW_OPENGL, SDL_WINDOW_INPUT_GRABBED,
 //     SDL_WINDOW_HIDDEN, SDL_WINDOW_RESIZABLE, SDL_WINDOW_MAXIMIZED,
 //     SDL_WINDOW_MINIMIZED, SDL_WINDOW_BORDERLESS is always set, and
 //     SDL_WINDOW_FULLSCREEN is always unset.
 //   
func CreateShapedWindow(title string, x uint, y uint, w uint, h uint, flags WindowFlags) (retval *Window) {
    tmp_title := C.CString(title); defer C.free(unsafe.Pointer(tmp_title))
    retval = (*Window)(unsafe.Pointer(C.SDL_CreateShapedWindow((*C.char)(tmp_title), C.uint(x), C.uint(y), C.uint(w), C.uint(h), C.Uint32(flags))))
    return
}

 // Return whether the given window is a shaped window.
 // 
 // Returns: SDL_TRUE if the window is a window that can be shaped,
 // SDL_FALSE if the window is unshaped or NULL.
 // 
 // See also: SDL_CreateShapedWindow
 // 
 //   window
 //     The window to query for being shaped.
 //   
func (window *Window) IsShaped() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_IsShapedWindow((*C.SDL_Window)(window)))
    return
}

 // Set the shape and parameters of a shaped window.
 // 
 // Returns: 0 on success, SDL_INVALID_SHAPE_ARGUMENT on invalid an
 // invalid shape argument, or SDL_NONSHAPEABLE_WINDOW if the SDL_Window*
 // given does not reference a valid shaped window.
 // 
 // See also: SDL_WindowShapeMode
 // 
 // See also: SDL_GetShapedWindowMode.
 // 
 //   window
 //     The shaped window whose parameters should be set.
 //   
 //   shape
 //     A surface encoding the desired shape for the window.
 //   
 //   shape_mode
 //     The parameters to set for the shaped window.
 //   
func (window *Window) SetShape(shape *Surface, shape_mode *WindowShapeMode) (retval int) {
    var tmp_shape_mode *C.SDL_WindowShapeMode; if shape_mode != nil { x := toCFromWindowShapeMode(*shape_mode); tmp_shape_mode = &x }
    retval = int(C.SDL_SetWindowShape((*C.SDL_Window)(window), (*C.SDL_Surface)(shape), (*C.SDL_WindowShapeMode)(tmp_shape_mode)))
    return
}

 // Get the shape parameters of a shaped window.
 // 
 // Returns: 0 if the window has a shape and, provided shape_mode was not
 // NULL, shape_mode has been filled with the mode data,
 // SDL_NONSHAPEABLE_WINDOW if the SDL_Window given is not a shaped
 // window, or SDL_WINDOW_LACKS_SHAPE if the SDL_Window* given is a
 // shapeable window currently lacking a shape.
 // 
 // See also: SDL_WindowShapeMode
 // 
 // See also: SDL_SetWindowShape
 // 
 //   window
 //     The shaped window whose parameters should be retrieved.
 //   
 //   shape_mode
 //     An empty shape-mode structure to fill, or NULL to check whether the
 //     window has a shape.
 //   
func (window *Window) GetShapedMode() (retval int, shape_mode *WindowShapeMode) {
    tmp_shape_mode := new(C.SDL_WindowShapeMode)
    retval = int(C.SDL_GetShapedWindowMode((*C.SDL_Window)(window), (*C.SDL_WindowShapeMode)(tmp_shape_mode)))
    tmp2_shape_mode := fromC2WindowShapeMode(*(tmp_shape_mode)); shape_mode = &tmp2_shape_mode
    return
}
