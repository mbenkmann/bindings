// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
// #include <SDL_shape.h>
import "C"
import "unsafe"

 // Include file for SDL keyboard event handling

 // The SDL keysym structure, used in key events.
 // 
 // Note: If you are looking for translated character input, see the
 // SDL_TEXTINPUT event.
 // 
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
func GetKeyboardFocus() (retval *Window) {
    retval = (*Window)(unsafe.Pointer(C.SDL_GetKeyboardFocus()))
    return
}

 // Get a snapshot of the current state of the keyboard.
 // 
 // Returns: An array of key states. Indexes into this array are obtained
 // by using SDL_Scancode values.
 // 
 //   numkeys
 //     if non-NULL, receives the length of the returned array.
 //   
 // Example:
 //   const Uint8 *state = SDL_GetKeyboardState(NULL);
 //   
 //   if ( state[SDL_SCANCODE_RETURN] )   {
 //   
 //       printf("<RETURN> is pressed.\n");
 //   
 //   }
 //   
func GetKeyboardState() (retval *[999999999]byte, numkeys int) {
    tmp_numkeys := new(C.int)
    retval = (*[999999999]byte)(unsafe.Pointer(C.SDL_GetKeyboardState((*C.int)(tmp_numkeys))))
    numkeys = deref_int_ptr(tmp_numkeys)
    return
}

 // Get the current key modifier state for the keyboard.
func GetModState() (retval Keymod) {
    retval = Keymod(C.SDL_GetModState())
    return
}

 // Set the current key modifier state for the keyboard.
 // 
 // Note: This does not change the keyboard state, only the key modifier
 // flags.
 // 
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
func StartTextInput() {
    C.SDL_StartTextInput()
}

 // Return whether or not Unicode text input events are enabled.
 // 
 // See also: SDL_StartTextInput()
 // 
 // See also: SDL_StopTextInput()
 // 
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
func StopTextInput() {
    C.SDL_StopTextInput()
}

 // Set the rectangle used to type Unicode text inputs. This is used as a
 // hint for IME and on-screen keyboard placement.
 // 
 // See also: SDL_StartTextInput()
 // 
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
func (window *Window) IsScreenKeyboardShown() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_IsScreenKeyboardShown((*C.SDL_Window)(window)))
    return
}
