// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"
import "unsafe"

 // Include file for SDL keyboard event handling

 // The SDL keysym structure, used in key events.
 // 
 // Note: If you are looking for translated character input, see the
 // SDL_TEXTINPUT event.
 // 
 // ↪ https://wiki.libsdl.org/SDL_Keysym
type Keysym struct {
     // SDL physical key code - see SDL_Scancode for details
    Scancode Scancode

     // SDL virtual key code - see SDL_Keycode for details
    Sym Keycode

     // current key modifiers
    Mod uint16

    Unused uint32
}

func fromC2Keysym(s C.SDL_Keysym) Keysym {
    return Keysym{Scancode(s.scancode), Keycode(s.sym), uint16(s.mod), uint32(s.unused)}
}

func toCFromKeysym(s Keysym) (d C.SDL_Keysym) {
    d.scancode = C.SDL_Scancode(s.Scancode)
    d.sym = C.SDL_Keycode(s.Sym)
    d.mod = C.Uint16(s.Mod)
    d.unused = C.Uint32(s.Unused)
    return
}


 // Get the window which currently has keyboard focus.
 // ↪ https://wiki.libsdl.org/SDL_GetKeyboardFocus
func GetKeyboardFocus() (retval *Window) {
    retval = (*Window)(unsafe.Pointer(C.SDL_GetKeyboardFocus()))
    return
}


 // Get the current key modifier state for the keyboard.
 // ↪ https://wiki.libsdl.org/SDL_GetModState
func GetModState() (retval Keymod) {
    retval = Keymod(C.SDL_GetModState())
    return
}

 // Set the current key modifier state for the keyboard.
 // 
 // Note: This does not change the keyboard state, only the key modifier
 // flags.
 // 
 // ↪ https://wiki.libsdl.org/SDL_SetModState
func SetModState(modstate Keymod) {
    C.SDL_SetModState(C.SDL_Keymod(modstate))
}

 // Get the key code corresponding to the given scancode according to the
 // current keyboard layout.
 // 
 // See SDL_Keycode for details.
 // 
 // See also: SDL_GetKeyName()
 // 
 // ↪ https://wiki.libsdl.org/SDL_GetKeyFromScancode
func GetKeyFromScancode(scancode Scancode) (retval Keycode) {
    retval = Keycode(C.SDL_GetKeyFromScancode(C.SDL_Scancode(scancode)))
    return
}

 // Get the scancode corresponding to the given key code according to the
 // current keyboard layout.
 // 
 // See SDL_Scancode for details.
 // 
 // See also: SDL_GetScancodeName()
 // 
 // ↪ https://wiki.libsdl.org/SDL_GetScancodeFromKey
func GetScancodeFromKey(key Keycode) (retval Scancode) {
    retval = Scancode(C.SDL_GetScancodeFromKey(C.SDL_Keycode(key)))
    return
}

 // Get a human-readable name for a scancode.
 // 
 // Returns: A pointer to the name for the scancode. If the scancode
 // doesn't have a name, this function returns an empty string ("").
 // 
 // See also: SDL_Scancode
 // 
 // ↪ https://wiki.libsdl.org/SDL_GetScancodeName
func GetScancodeName(scancode Scancode) (retval string) {
    retval = C.GoString(C.SDL_GetScancodeName(C.SDL_Scancode(scancode)))
    return
}

 // Get a scancode from a human-readable name.
 // 
 // Returns: scancode, or SDL_SCANCODE_UNKNOWN if the name wasn't
 // recognized
 // 
 // See also: SDL_Scancode
 // 
 // ↪ https://wiki.libsdl.org/SDL_GetScancodeFromName
func GetScancodeFromName(name string) (retval Scancode) {
    tmp_name := C.CString(name); defer C.free(unsafe.Pointer(tmp_name))
    retval = Scancode(C.SDL_GetScancodeFromName((*C.char)(tmp_name)))
    return
}

 // Get a human-readable name for a key.
 // 
 // Returns: A pointer to a UTF-8 string that stays valid at least until
 // the next call to this function. If you need it around any longer, you
 // must copy it. If the key doesn't have a name, this function returns an
 // empty string ("").
 // 
 // See also: SDL_Key
 // 
 // ↪ https://wiki.libsdl.org/SDL_GetKeyName
func GetKeyName(key Keycode) (retval string) {
    retval = C.GoString(C.SDL_GetKeyName(C.SDL_Keycode(key)))
    return
}

 // Get a key code from a human-readable name.
 // 
 // Returns: key code, or SDLK_UNKNOWN if the name wasn't recognized
 // 
 // See also: SDL_Keycode
 // 
 // ↪ https://wiki.libsdl.org/SDL_GetKeyFromName
func GetKeyFromName(name string) (retval Keycode) {
    tmp_name := C.CString(name); defer C.free(unsafe.Pointer(tmp_name))
    retval = Keycode(C.SDL_GetKeyFromName((*C.char)(tmp_name)))
    return
}

 // Start accepting Unicode text input events. This function will show the
 // on-screen keyboard if supported.
 // 
 // See also: SDL_StopTextInput()
 // 
 // See also: SDL_SetTextInputRect()
 // 
 // See also: SDL_HasScreenKeyboardSupport()
 // 
 // ↪ https://wiki.libsdl.org/SDL_StartTextInput
func StartTextInput() {
    C.SDL_StartTextInput()
}

 // Return whether or not Unicode text input events are enabled.
 // 
 // See also: SDL_StartTextInput()
 // 
 // See also: SDL_StopTextInput()
 // 
 // ↪ https://wiki.libsdl.org/SDL_IsTextInputActive
func IsTextInputActive() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_IsTextInputActive())
    return
}

 // Stop receiving any text input events. This function will hide the on-
 // screen keyboard if supported.
 // 
 // See also: SDL_StartTextInput()
 // 
 // See also: SDL_HasScreenKeyboardSupport()
 // 
 // ↪ https://wiki.libsdl.org/SDL_StopTextInput
func StopTextInput() {
    C.SDL_StopTextInput()
}

 // Set the rectangle used to type Unicode text inputs. This is used as a
 // hint for IME and on-screen keyboard placement.
 // 
 // See also: SDL_StartTextInput()
 // 
 // ↪ https://wiki.libsdl.org/SDL_SetTextInputRect
func SetTextInputRect(rect Rect) {
    tmp_rect := toCFromRect(rect)
    C.SDL_SetTextInputRect((*C.SDL_Rect)(&tmp_rect))
}

 // Returns whether the platform has some screen keyboard support.
 // 
 // Returns: SDL_TRUE if some keyboard support is available else
 // SDL_FALSE.
 // 
 // Note: Not all screen keyboard functions are supported on all
 // platforms.
 // 
 // See also: SDL_IsScreenKeyboardShown()
 // 
 // ↪ https://wiki.libsdl.org/SDL_HasScreenKeyboardSupport
func HasScreenKeyboardSupport() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasScreenKeyboardSupport())
    return
}

 // Returns whether the screen keyboard is shown for given window.
 // 
 // Returns: SDL_TRUE if screen keyboard is shown else SDL_FALSE.
 // 
 // See also: SDL_HasScreenKeyboardSupport()
 // 
 //   window
 //     The window for which screen keyboard should be queried.
 //   
 // ↪ https://wiki.libsdl.org/SDL_IsScreenKeyboardShown
func (window *Window) IsScreenKeyboardShown() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_IsScreenKeyboardShown((*C.SDL_Window)(window)))
    return
}
