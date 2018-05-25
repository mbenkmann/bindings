// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"
import "unsafe"

 // Include file for SDL mouse event handling.

const (
    BUTTON_LEFT = C.SDL_BUTTON_LEFT

    BUTTON_MIDDLE = C.SDL_BUTTON_MIDDLE

    BUTTON_RIGHT = C.SDL_BUTTON_RIGHT

    BUTTON_X1 = C.SDL_BUTTON_X1

    BUTTON_X2 = C.SDL_BUTTON_X2

    BUTTON_LMASK = C.SDL_BUTTON_LMASK

    BUTTON_MMASK = C.SDL_BUTTON_MMASK

    BUTTON_RMASK = C.SDL_BUTTON_RMASK

    BUTTON_X1MASK = C.SDL_BUTTON_X1MASK

    BUTTON_X2MASK = C.SDL_BUTTON_X2MASK
)

 // Cursor types for SDL_CreateSystemCursor.
type SystemCursor int
const (
     // Arrow
    SYSTEM_CURSOR_ARROW SystemCursor = C.SDL_SYSTEM_CURSOR_ARROW

     // I-beam
    SYSTEM_CURSOR_IBEAM SystemCursor = C.SDL_SYSTEM_CURSOR_IBEAM

     // Wait
    SYSTEM_CURSOR_WAIT SystemCursor = C.SDL_SYSTEM_CURSOR_WAIT

     // Crosshair
    SYSTEM_CURSOR_CROSSHAIR SystemCursor = C.SDL_SYSTEM_CURSOR_CROSSHAIR

     // Small wait cursor (or Wait if not available)
    SYSTEM_CURSOR_WAITARROW SystemCursor = C.SDL_SYSTEM_CURSOR_WAITARROW

     // Double arrow pointing northwest and southeast
    SYSTEM_CURSOR_SIZENWSE SystemCursor = C.SDL_SYSTEM_CURSOR_SIZENWSE

     // Double arrow pointing northeast and southwest
    SYSTEM_CURSOR_SIZENESW SystemCursor = C.SDL_SYSTEM_CURSOR_SIZENESW

     // Double arrow pointing west and east
    SYSTEM_CURSOR_SIZEWE SystemCursor = C.SDL_SYSTEM_CURSOR_SIZEWE

     // Double arrow pointing north and south
    SYSTEM_CURSOR_SIZENS SystemCursor = C.SDL_SYSTEM_CURSOR_SIZENS

     // Four pointed arrow pointing north, south, east, and west
    SYSTEM_CURSOR_SIZEALL SystemCursor = C.SDL_SYSTEM_CURSOR_SIZEALL

     // Slashed circle or crossbones
    SYSTEM_CURSOR_NO SystemCursor = C.SDL_SYSTEM_CURSOR_NO

     // Hand
    SYSTEM_CURSOR_HAND SystemCursor = C.SDL_SYSTEM_CURSOR_HAND

    NUM_SYSTEM_CURSORS SystemCursor = C.SDL_NUM_SYSTEM_CURSORS
)

 // Scroll direction types for the Scroll event.
type MouseWheelDirection int
const (
     // The scroll direction is normal
    MOUSEWHEEL_NORMAL MouseWheelDirection = C.SDL_MOUSEWHEEL_NORMAL

     // The scroll direction is flipped / natural
    MOUSEWHEEL_FLIPPED MouseWheelDirection = C.SDL_MOUSEWHEEL_FLIPPED
)

type Cursor C.SDL_Cursor


 // Get the window which currently has mouse focus.
func GetMouseFocus() (retval *Window) {
    retval = (*Window)(unsafe.Pointer(C.SDL_GetMouseFocus()))
    return
}

 // Retrieve the current state of the mouse.
 // 
 // The current button state is returned as a button bitmask, which can be
 // tested using the SDL_BUTTON(X) macros, and x and y are set to the
 // mouse cursor position relative to the focus window for the currently
 // selected mouse. You can pass NULL for either x or y.
func GetMouseState() (retval uint32, x int, y int) {
    tmp_x := new(C.int)
    tmp_y := new(C.int)
    retval = uint32(C.SDL_GetMouseState((*C.int)(tmp_x), (*C.int)(tmp_y)))
    x = deref_int_ptr(tmp_x)
    y = deref_int_ptr(tmp_y)
    return
}

 // Get the current state of the mouse, in relation to the desktop.
 // 
 // This works just like SDL_GetMouseState(), but the coordinates will be
 // reported relative to the top-left of the desktop. This can be useful
 // if you need to track the mouse outside of a specific window and
 // SDL_CaptureMouse() doesn't fit your needs. For example, it could be
 // useful if you need to track the mouse while dragging a window, where
 // coordinates relative to a window might not be in sync at all times.
 // 
 // Note: SDL_GetMouseState() returns the mouse position as SDL
 // understands it from the last pump of the event queue. This function,
 // however, queries the OS for the current mouse position, and as such,
 // might be a slightly less efficient function. Unless you know what
 // you're doing and have a good reason to use this function, you probably
 // want SDL_GetMouseState() instead.
 // 
 // Returns: The current button state as a bitmask, which can be tested
 // using the SDL_BUTTON(X) macros.
 // 
 // See also: SDL_GetMouseState
 // 
 //   x
 //     Returns the current X coord, relative to the desktop. Can be NULL.
 //   
 //   y
 //     Returns the current Y coord, relative to the desktop. Can be NULL.
 //   
func GetGlobalMouseState() (retval uint32, x int, y int) {
    tmp_x := new(C.int)
    tmp_y := new(C.int)
    retval = uint32(C.SDL_GetGlobalMouseState((*C.int)(tmp_x), (*C.int)(tmp_y)))
    x = deref_int_ptr(tmp_x)
    y = deref_int_ptr(tmp_y)
    return
}

 // Retrieve the relative state of the mouse.
 // 
 // The current button state is returned as a button bitmask, which can be
 // tested using the SDL_BUTTON(X) macros, and x and y are set to the
 // mouse deltas since the last call to SDL_GetRelativeMouseState().
func GetRelativeMouseState() (retval uint32, x int, y int) {
    tmp_x := new(C.int)
    tmp_y := new(C.int)
    retval = uint32(C.SDL_GetRelativeMouseState((*C.int)(tmp_x), (*C.int)(tmp_y)))
    x = deref_int_ptr(tmp_x)
    y = deref_int_ptr(tmp_y)
    return
}

 // Moves the mouse to the given position within the window.
 // 
 // Note: This function generates a mouse motion event
 // 
 //   window
 //     The window to move the mouse into, or NULL for the current mouse focus
 //   
 //   x
 //     The x coordinate within the window
 //   
 //   y
 //     The y coordinate within the window
 //   
func WarpMouseInWindow(window *Window, x int, y int) {
    C.SDL_WarpMouseInWindow((*C.SDL_Window)(window), C.int(x), C.int(y))
}

 // Moves the mouse to the given position in global screen space.
 // 
 // Returns: 0 on success, -1 on error (usually: unsupported by a
 // platform).
 // 
 // Note: This function generates a mouse motion event
 // 
 //   x
 //     The x coordinate
 //   
 //   y
 //     The y coordinate
 //   
func WarpMouseGlobal(x int, y int) (retval int) {
    retval = int(C.SDL_WarpMouseGlobal(C.int(x), C.int(y)))
    return
}

 // Set relative mouse mode.
 // 
 // Returns: 0 on success, or -1 if relative mode is not supported.
 // 
 //   enabled
 //     Whether or not to enable relative mode
 //   
 // While the mouse is in relative mode, the cursor is hidden, and the
 // driver will try to report continuous motion in the current window.
 // Only relative motion events will be delivered, the mouse position will
 // not change.
 // 
 // Note: This function will flush any pending mouse motion.
 // 
 // See also: SDL_GetRelativeMouseMode()
 // 
func SetRelativeMouseMode(enabled bool) (retval int) {
    retval = int(C.SDL_SetRelativeMouseMode(bool2bool(enabled)))
    return
}

 // Capture the mouse, to track input outside an SDL window.
 // 
 //   enabled
 //     Whether or not to enable capturing
 //   
 // Capturing enables your app to obtain mouse events globally, instead of
 // just within your window. Not all video targets support this function.
 // When capturing is enabled, the current window will get all mouse
 // events, but unlike relative mode, no change is made to the cursor and
 // it is not restrained to your window.
 // 
 // This function may also deny mouse input to other windows--both those
 // in your application and others on the system--so you should use this
 // function sparingly, and in small bursts. For example, you might want
 // to track the mouse while the user is dragging something, until the
 // user releases a mouse button. It is not recommended that you capture
 // the mouse for long periods of time, such as the entire time your app
 // is running.
 // 
 // While captured, mouse events still report coordinates relative to the
 // current (foreground) window, but those coordinates may be outside the
 // bounds of the window (including negative values). Capturing is only
 // allowed for the foreground window. If the window loses focus while
 // capturing, the capture will be disabled automatically.
 // 
 // While capturing is enabled, the current window will have the
 // SDL_WINDOW_MOUSE_CAPTURE flag set.
 // 
 // Returns: 0 on success, or -1 if not supported.
 // 
func CaptureMouse(enabled bool) (retval int) {
    retval = int(C.SDL_CaptureMouse(bool2bool(enabled)))
    return
}

 // Query whether relative mouse mode is enabled.
 // 
 // See also: SDL_SetRelativeMouseMode()
 // 
func GetRelativeMouseMode() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_GetRelativeMouseMode())
    return
}


 // Create a color cursor.
 // 
 // See also: SDL_FreeCursor()
 // 
func CreateColorCursor(surface *Surface, hot_x int, hot_y int) (retval *Cursor) {
    retval = (*Cursor)(unsafe.Pointer(C.SDL_CreateColorCursor((*C.SDL_Surface)(surface), C.int(hot_x), C.int(hot_y))))
    return
}

 // Create a system cursor.
 // 
 // See also: SDL_FreeCursor()
 // 
func CreateSystemCursor(id SystemCursor) (retval *Cursor) {
    retval = (*Cursor)(unsafe.Pointer(C.SDL_CreateSystemCursor(C.SDL_SystemCursor(id))))
    return
}

 // Set the active cursor.
func SetCursor(cursor *Cursor) {
    C.SDL_SetCursor((*C.SDL_Cursor)(cursor))
}

 // Return the active cursor.
func GetCursor() (retval *Cursor) {
    retval = (*Cursor)(unsafe.Pointer(C.SDL_GetCursor()))
    return
}

 // Return the default cursor.
func GetDefaultCursor() (retval *Cursor) {
    retval = (*Cursor)(unsafe.Pointer(C.SDL_GetDefaultCursor()))
    return
}

 // Frees a cursor created with SDL_CreateCursor().
 // 
 // See also: SDL_CreateCursor()
 // 
func (cursor *Cursor) Free() {
    C.SDL_FreeCursor((*C.SDL_Cursor)(cursor))
}

 // Toggle whether or not the cursor is shown.
 // 
 // Returns: 1 if the cursor is shown, or 0 if the cursor is hidden.
 // 
 //   toggle
 //     1 to show the cursor, 0 to hide it, -1 to query the current state.
 //   
func ShowCursor(toggle int) (retval int) {
    retval = int(C.SDL_ShowCursor(C.int(toggle)))
    return
}
