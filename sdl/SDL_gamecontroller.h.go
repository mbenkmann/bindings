// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"
import "unsafe"

 // Include file for SDL game controller event handling
 // 
 // In order to use these functions, SDL_Init() must have been called with
 // the SDL_INIT_GAMECONTROLLER flag. This causes SDL to scan the system
 // for game controllers, and load appropriate drivers.
 // 
 // If you would like to receive controller updates while the application
 // is in the background, you should set the following hint before calling
 // SDL_Init(): SDL_HINT_JOYSTICK_ALLOW_BACKGROUND_EVENTS


type GameControllerBindType int
const (
    CONTROLLER_BINDTYPE_NONE GameControllerBindType = C.SDL_CONTROLLER_BINDTYPE_NONE

    CONTROLLER_BINDTYPE_BUTTON GameControllerBindType = C.SDL_CONTROLLER_BINDTYPE_BUTTON

    CONTROLLER_BINDTYPE_AXIS GameControllerBindType = C.SDL_CONTROLLER_BINDTYPE_AXIS

    CONTROLLER_BINDTYPE_HAT GameControllerBindType = C.SDL_CONTROLLER_BINDTYPE_HAT
)

 // The list of axes available from a controller
 // 
 // Thumbstick axis values range from SDL_JOYSTICK_AXIS_MIN to
 // SDL_JOYSTICK_AXIS_MAX, and are centered within ~8000 of zero, though
 // advanced UI will allow users to set or autodetect the dead zone, which
 // varies between controllers.
 // 
 // Trigger axis values range from 0 to SDL_JOYSTICK_AXIS_MAX.
type GameControllerAxis int
const (
    CONTROLLER_AXIS_INVALID GameControllerAxis = C.SDL_CONTROLLER_AXIS_INVALID

    CONTROLLER_AXIS_LEFTX GameControllerAxis = C.SDL_CONTROLLER_AXIS_LEFTX

    CONTROLLER_AXIS_LEFTY GameControllerAxis = C.SDL_CONTROLLER_AXIS_LEFTY

    CONTROLLER_AXIS_RIGHTX GameControllerAxis = C.SDL_CONTROLLER_AXIS_RIGHTX

    CONTROLLER_AXIS_RIGHTY GameControllerAxis = C.SDL_CONTROLLER_AXIS_RIGHTY

    CONTROLLER_AXIS_TRIGGERLEFT GameControllerAxis = C.SDL_CONTROLLER_AXIS_TRIGGERLEFT

    CONTROLLER_AXIS_TRIGGERRIGHT GameControllerAxis = C.SDL_CONTROLLER_AXIS_TRIGGERRIGHT

    CONTROLLER_AXIS_MAX GameControllerAxis = C.SDL_CONTROLLER_AXIS_MAX
)

 // The list of buttons available from a controller
 // ↪ https://wiki.libsdl.org/SDL_GameControllerButton
type GameControllerButton int
const (
    CONTROLLER_BUTTON_INVALID GameControllerButton = C.SDL_CONTROLLER_BUTTON_INVALID

    CONTROLLER_BUTTON_A GameControllerButton = C.SDL_CONTROLLER_BUTTON_A

    CONTROLLER_BUTTON_B GameControllerButton = C.SDL_CONTROLLER_BUTTON_B

    CONTROLLER_BUTTON_X GameControllerButton = C.SDL_CONTROLLER_BUTTON_X

    CONTROLLER_BUTTON_Y GameControllerButton = C.SDL_CONTROLLER_BUTTON_Y

    CONTROLLER_BUTTON_BACK GameControllerButton = C.SDL_CONTROLLER_BUTTON_BACK

    CONTROLLER_BUTTON_GUIDE GameControllerButton = C.SDL_CONTROLLER_BUTTON_GUIDE

    CONTROLLER_BUTTON_START GameControllerButton = C.SDL_CONTROLLER_BUTTON_START

    CONTROLLER_BUTTON_LEFTSTICK GameControllerButton = C.SDL_CONTROLLER_BUTTON_LEFTSTICK

    CONTROLLER_BUTTON_RIGHTSTICK GameControllerButton = C.SDL_CONTROLLER_BUTTON_RIGHTSTICK

    CONTROLLER_BUTTON_LEFTSHOULDER GameControllerButton = C.SDL_CONTROLLER_BUTTON_LEFTSHOULDER

    CONTROLLER_BUTTON_RIGHTSHOULDER GameControllerButton = C.SDL_CONTROLLER_BUTTON_RIGHTSHOULDER

    CONTROLLER_BUTTON_DPAD_UP GameControllerButton = C.SDL_CONTROLLER_BUTTON_DPAD_UP

    CONTROLLER_BUTTON_DPAD_DOWN GameControllerButton = C.SDL_CONTROLLER_BUTTON_DPAD_DOWN

    CONTROLLER_BUTTON_DPAD_LEFT GameControllerButton = C.SDL_CONTROLLER_BUTTON_DPAD_LEFT

    CONTROLLER_BUTTON_DPAD_RIGHT GameControllerButton = C.SDL_CONTROLLER_BUTTON_DPAD_RIGHT

    CONTROLLER_BUTTON_MAX GameControllerButton = C.SDL_CONTROLLER_BUTTON_MAX
)

type GameController C.SDL_GameController


 // To count the number of game controllers in the system for the
 // following: int nJoysticks = SDL_NumJoysticks(); int nGameControllers =
 // 0; for (int i = 0; i < nJoysticks; i++) { if (SDL_IsGameController(i))
 // { nGameControllers++; } }
 // 
 // Using the SDL_HINT_GAMECONTROLLERCONFIG hint or the
 // SDL_GameControllerAddMapping() you can add support for controllers SDL
 // is unaware of or cause an existing controller to have a different
 // binding. The format is: guid,name,mappings
 // 
 // Where GUID is the string value from SDL_JoystickGetGUIDString(), name
 // is the human readable string for the device and mappings are
 // controller mappings to joystick ones. Under Windows there is a
 // reserved GUID of "xinput" that covers any XInput devices. The mapping
 // format for joystick is: bX - a joystick button, index X hX.Y - hat X
 // with value Y aX - axis X of the joystick Buttons can be used as a
 // controller axis and vice versa.
 // 
 // This string shows an example of a valid mapping for a controller
 // "03000000341a00003608000000000000,PS3 Controller,a:b1,b:b2,y:b3,x:b0,s
 // tart:b9,guide:b12,back:b8,dpup:h0.1,dpleft:h0.8,dpdown:h0.4,dpright:h0
 // .2,leftshoulder:b4,rightshoulder:b5,leftstick:b10,rightstick:b11,leftx
 // :a0,lefty:a1,rightx:a2,righty:a3,lefttrigger:b6,righttrigger:b7", Load
 // a set of mappings from a seekable SDL data stream (memory or file),
 // filtered by the current SDL_GetPlatform() A community sourced database
 // of controllers is available at https://raw.github.com/gabomdq/SDL_Game
 // ControllerDB/master/gamecontrollerdb.txt
 // 
 // If freerw is non-zero, the stream will be closed after being read.
 // 
 // Returns: number of mappings added, -1 on error
 // 
 // ↪ https://wiki.libsdl.org/SDL_GameControllerAddMappingsFromRW
func GameControllerAddMappingsFromRW(rw *RWops, freerw int) (retval int) {
    retval = int(C.SDL_GameControllerAddMappingsFromRW((*C.SDL_RWops)(rw), C.int(freerw)))
    return
}

 // Add or update an existing mapping configuration
 // 
 // Returns: 1 if mapping is added, 0 if updated, -1 on error
 // 
 // ↪ https://wiki.libsdl.org/SDL_GameControllerAddMapping
func GameControllerAddMapping(mappingString string) (retval int) {
    tmp_mappingString := C.CString(mappingString); defer C.free(unsafe.Pointer(tmp_mappingString))
    retval = int(C.SDL_GameControllerAddMapping((*C.char)(tmp_mappingString)))
    return
}

 // Get the number of mappings installed
 // 
 // Returns: the number of mappings
 // 
func GameControllerNumMappings() (retval int) {
    retval = int(C.SDL_GameControllerNumMappings())
    return
}

 // Get the mapping at a particular index.
 // 
 // Returns: the mapping string. Must be freed with SDL_free(). Returns
 // NULL if the index is out of range.
 // 
func GameControllerMappingForIndex(mapping_index int) (retval string) {
    retval = C.GoString(C.SDL_GameControllerMappingForIndex(C.int(mapping_index)))
    return
}

 // Get a mapping string for a GUID
 // 
 // Returns: the mapping string. Must be freed with SDL_free(). Returns
 // NULL if no mapping is available
 // 
 // ↪ https://wiki.libsdl.org/SDL_GameControllerMappingForGUID
func GameControllerMappingForGUID(guid JoystickGUID) (retval string) {
    tmp_guid := toCFromJoystickGUID(guid)
    retval = C.GoString(C.SDL_GameControllerMappingForGUID(C.SDL_JoystickGUID(tmp_guid)))
    return
}

 // Get a mapping string for an open GameController
 // 
 // Returns: the mapping string. Must be freed with SDL_free(). Returns
 // NULL if no mapping is available
 // 
 // ↪ https://wiki.libsdl.org/SDL_GameControllerMapping
func (gamecontroller *GameController) Mapping() (retval string) {
    retval = freeGoString(C.SDL_GameControllerMapping((*C.SDL_GameController)(gamecontroller)))
    return
}

 // Is the joystick on this index supported by the game controller
 // interface?
 // ↪ https://wiki.libsdl.org/SDL_IsGameController
func IsGameController(joystick_index int) (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_IsGameController(C.int(joystick_index)))
    return
}

 // Get the implementation dependent name of a game controller. This can
 // be called before any controllers are opened. If no name can be found,
 // this function returns NULL.
 // ↪ https://wiki.libsdl.org/SDL_GameControllerNameForIndex
func GameControllerNameForIndex(joystick_index int) (retval string) {
    retval = C.GoString(C.SDL_GameControllerNameForIndex(C.int(joystick_index)))
    return
}

 // Open a game controller for use. The index passed as an argument refers
 // to the N'th game controller on the system. This index is not the value
 // which will identify this controller in future controller events. The
 // joystick's instance id (SDL_JoystickID) will be used there instead.
 // 
 // Returns: A controller identifier, or NULL if an error occurred.
 // 
 // ↪ https://wiki.libsdl.org/SDL_GameControllerOpen
func GameControllerOpen(joystick_index int) (retval *GameController) {
    retval = (*GameController)(unsafe.Pointer(C.SDL_GameControllerOpen(C.int(joystick_index))))
    return
}

 // Return the SDL_GameController associated with an instance id.
 // ↪ https://wiki.libsdl.org/SDL_GameControllerFromInstanceID
func GameControllerFromInstanceID(joyid JoystickID) (retval *GameController) {
    retval = (*GameController)(unsafe.Pointer(C.SDL_GameControllerFromInstanceID(C.SDL_JoystickID(joyid))))
    return
}

 // Return the name for this currently opened controller
 // ↪ https://wiki.libsdl.org/SDL_GameControllerName
func (gamecontroller *GameController) Name() (retval string) {
    retval = C.GoString(C.SDL_GameControllerName((*C.SDL_GameController)(gamecontroller)))
    return
}

 // Get the USB vendor ID of an opened controller, if available. If the
 // vendor ID isn't available this function returns 0.
func (gamecontroller *GameController) GetVendor() (retval uint16) {
    retval = uint16(C.SDL_GameControllerGetVendor((*C.SDL_GameController)(gamecontroller)))
    return
}

 // Get the USB product ID of an opened controller, if available. If the
 // product ID isn't available this function returns 0.
func (gamecontroller *GameController) GetProduct() (retval uint16) {
    retval = uint16(C.SDL_GameControllerGetProduct((*C.SDL_GameController)(gamecontroller)))
    return
}

 // Get the product version of an opened controller, if available. If the
 // product version isn't available this function returns 0.
func (gamecontroller *GameController) GetProductVersion() (retval uint16) {
    retval = uint16(C.SDL_GameControllerGetProductVersion((*C.SDL_GameController)(gamecontroller)))
    return
}

 // Returns SDL_TRUE if the controller has been opened and currently
 // connected, or SDL_FALSE if it has not.
 // ↪ https://wiki.libsdl.org/SDL_GameControllerGetAttached
func (gamecontroller *GameController) GetAttached() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_GameControllerGetAttached((*C.SDL_GameController)(gamecontroller)))
    return
}

 // Get the underlying joystick object used by a controller
 // ↪ https://wiki.libsdl.org/SDL_GameControllerGetJoystick
func (gamecontroller *GameController) GetJoystick() (retval *Joystick) {
    retval = (*Joystick)(unsafe.Pointer(C.SDL_GameControllerGetJoystick((*C.SDL_GameController)(gamecontroller))))
    return
}

 // Enable/disable controller event polling.
 // 
 // If controller events are disabled, you must call
 // SDL_GameControllerUpdate() yourself and check the state of the
 // controller when you want controller information.
 // 
 // The state can be one of SDL_QUERY, SDL_ENABLE or SDL_IGNORE.
 // ↪ https://wiki.libsdl.org/SDL_GameControllerEventState
func GameControllerEventState(state int) (retval int) {
    retval = int(C.SDL_GameControllerEventState(C.int(state)))
    return
}

 // Update the current state of the open game controllers.
 // 
 // This is called automatically by the event loop if any game controller
 // events are enabled.
 // ↪ https://wiki.libsdl.org/SDL_GameControllerUpdate
func GameControllerUpdate() {
    C.SDL_GameControllerUpdate()
}

 // turn this string into a axis mapping
 // ↪ https://wiki.libsdl.org/SDL_GameControllerGetAxisFromString
func GameControllerGetAxisFromString(pchString string) (retval GameControllerAxis) {
    tmp_pchString := C.CString(pchString); defer C.free(unsafe.Pointer(tmp_pchString))
    retval = GameControllerAxis(C.SDL_GameControllerGetAxisFromString((*C.char)(tmp_pchString)))
    return
}

 // turn this axis enum into a string mapping
 // ↪ https://wiki.libsdl.org/SDL_GameControllerGetStringForAxis
func GameControllerGetStringForAxis(axis GameControllerAxis) (retval string) {
    retval = C.GoString(C.SDL_GameControllerGetStringForAxis(C.SDL_GameControllerAxis(axis)))
    return
}

 // Get the SDL joystick layer binding for this controller button mapping
 // ↪ https://wiki.libsdl.org/SDL_GameControllerGetBindForAxis
func (gamecontroller *GameController) GetBindForAxis(axis GameControllerAxis) (retval GameControllerButtonBind) {
    retval = GameControllerButtonBind(C.SDL_GameControllerGetBindForAxis((*C.SDL_GameController)(gamecontroller), C.SDL_GameControllerAxis(axis)))
    return
}

 // Get the current state of an axis control on a game controller.
 // 
 // The state is a value ranging from -32768 to 32767 (except for the
 // triggers, which range from 0 to 32767).
 // 
 // The axis indices start at index 0.
 // ↪ https://wiki.libsdl.org/SDL_GameControllerGetAxis
func (gamecontroller *GameController) GetAxis(axis GameControllerAxis) (retval int16) {
    retval = int16(C.SDL_GameControllerGetAxis((*C.SDL_GameController)(gamecontroller), C.SDL_GameControllerAxis(axis)))
    return
}

 // turn this string into a button mapping
 // ↪ https://wiki.libsdl.org/SDL_GameControllerGetButtonFromString
func GameControllerGetButtonFromString(pchString string) (retval GameControllerButton) {
    tmp_pchString := C.CString(pchString); defer C.free(unsafe.Pointer(tmp_pchString))
    retval = GameControllerButton(C.SDL_GameControllerGetButtonFromString((*C.char)(tmp_pchString)))
    return
}

 // turn this button enum into a string mapping
 // ↪ https://wiki.libsdl.org/SDL_GameControllerGetStringForButton
func GameControllerGetStringForButton(button GameControllerButton) (retval string) {
    retval = C.GoString(C.SDL_GameControllerGetStringForButton(C.SDL_GameControllerButton(button)))
    return
}

 // Get the SDL joystick layer binding for this controller button mapping
 // ↪ https://wiki.libsdl.org/SDL_GameControllerGetBindForButton
func (gamecontroller *GameController) GetBindForButton(button GameControllerButton) (retval GameControllerButtonBind) {
    retval = GameControllerButtonBind(C.SDL_GameControllerGetBindForButton((*C.SDL_GameController)(gamecontroller), C.SDL_GameControllerButton(button)))
    return
}

 // Get the current state of a button on a game controller.
 // 
 // The button indices start at index 0.
 // ↪ https://wiki.libsdl.org/SDL_GameControllerGetButton
func (gamecontroller *GameController) GetButton(button GameControllerButton) (retval uint8) {
    retval = uint8(C.SDL_GameControllerGetButton((*C.SDL_GameController)(gamecontroller), C.SDL_GameControllerButton(button)))
    return
}

 // Close a controller previously opened with SDL_GameControllerOpen().
 // ↪ https://wiki.libsdl.org/SDL_GameControllerClose
func (gamecontroller *GameController) Close() {
    C.SDL_GameControllerClose((*C.SDL_GameController)(gamecontroller))
}
