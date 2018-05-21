// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
// #include <SDL_shape.h>
import "C"
import "unsafe"


 // Individual button data.
type MessageBoxButtonData struct {
     // SDL_MessageBoxButtonFlags
    Flags uint32

     // User defined button id (value returned via SDL_ShowMessageBox)
    Buttonid int

     // The UTF-8 button text
    Text string
}

func toCFromMessageBoxButtonData(s MessageBoxButtonData) (d C.SDL_MessageBoxButtonData) {
    d.flags = C.Uint32(s.Flags)
    d.buttonid = C.int(s.Buttonid)
    d.text = C.CString(s.Text)
    return
}

 // RGB value used in a message box color scheme.
type MessageBoxColor struct {
    R uint8

    G uint8

    B uint8
}

func fromC2MessageBoxColor(s C.SDL_MessageBoxColor) MessageBoxColor {
    return MessageBoxColor{uint8(s.r), uint8(s.g), uint8(s.b)}
}

func toCFromMessageBoxColor(s MessageBoxColor) (d C.SDL_MessageBoxColor) {
    d.r = C.Uint8(s.R)
    d.g = C.Uint8(s.G)
    d.b = C.Uint8(s.B)
    return
}

 // A set of colors to use for message box dialogs.
type MessageBoxColorScheme struct {
    Colors [MESSAGEBOX_COLOR_MAX]MessageBoxColor
}

 // SDL_MessageBox flags. If supported will display warning icon, etc.
type MessageBoxFlags int
const (
     // error dialog
    MESSAGEBOX_ERROR MessageBoxFlags = C.SDL_MESSAGEBOX_ERROR

     // warning dialog
    MESSAGEBOX_WARNING MessageBoxFlags = C.SDL_MESSAGEBOX_WARNING

     // informational dialog
    MESSAGEBOX_INFORMATION MessageBoxFlags = C.SDL_MESSAGEBOX_INFORMATION
)

 // Flags for SDL_MessageBoxButtonData.
type MessageBoxButtonFlags int
const (
     // Marks the default button when return is hit
    MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT MessageBoxButtonFlags = C.SDL_MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT

     // Marks the default button when escape is hit
    MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT MessageBoxButtonFlags = C.SDL_MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT
)

type MessageBoxColorType int
const (
    MESSAGEBOX_COLOR_BACKGROUND MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_BACKGROUND

    MESSAGEBOX_COLOR_TEXT MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_TEXT

    MESSAGEBOX_COLOR_BUTTON_BORDER MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_BUTTON_BORDER

    MESSAGEBOX_COLOR_BUTTON_BACKGROUND MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_BUTTON_BACKGROUND

    MESSAGEBOX_COLOR_BUTTON_SELECTED MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_BUTTON_SELECTED

    MESSAGEBOX_COLOR_MAX MessageBoxColorType = C.SDL_MESSAGEBOX_COLOR_MAX
)



 // Create a simple modal message box.
 // 
 // Returns: 0 on success, -1 on error
 // 
 // See also: SDL_ShowMessageBox
 // 
 //   flags
 //     SDL_MessageBoxFlags
 //   
 //   title
 //     UTF-8 title text
 //   
 //   message
 //     UTF-8 message text
 //   
 //   window
 //     The parent window, or NULL for no parent
 //   
func ShowSimpleMessageBox(flags uint32, title string, message string, window *Window) (retval int) {
    tmp_title := C.CString(title); defer C.free(unsafe.Pointer(tmp_title))
    tmp_message := C.CString(message); defer C.free(unsafe.Pointer(tmp_message))
    retval = int(C.SDL_ShowSimpleMessageBox(C.Uint32(flags), (*C.char)(tmp_title), (*C.char)(tmp_message), (*C.SDL_Window)(window)))
    return
}
