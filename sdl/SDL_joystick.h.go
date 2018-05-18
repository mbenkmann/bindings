// Machine-generated file.
// See http://winterdrache.de/bindings for details.

// Bindings for Simple DirectMedia Layer (www.libsdl.org)
package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
import "C"
import "unsafe"

 // Include file for SDL joystick event handling
 // 
 // The term "device_index" identifies currently plugged in joystick
 // devices between 0 and SDL_NumJoysticks, with the exact joystick behind
 // a device_index changing as joysticks are plugged and unplugged.
 // 
 // The term "instance_id" is the current instantiation of a joystick
 // device in the system, if the joystick is removed and then re-inserted
 // then it will get a new instance_id, instance_id's are monotonically
 // increasing identifiers of a joystick plugged in.
 // 
 // The term JoystickGUID is a stable 128-bit identifier for a joystick
 // device that does not change over time, it identifies class of the
 // device (a X360 wired controller for example). This identifier is
 // platform dependent.
 // 
 // In order to use these functions, SDL_Init() must have been called with
 // the SDL_INIT_JOYSTICK flag. This causes SDL to scan the system for
 // joysticks, and load appropriate drivers.
 // 
 // If you would like to receive joystick updates while the application is
 // in the background, you should set the following hint before calling
 // SDL_Init(): SDL_HINT_JOYSTICK_ALLOW_BACKGROUND_EVENTS

type JoystickGUID struct {
    Data [16]uint8
}

func fromC2JoystickGUID(s C.SDL_JoystickGUID) JoystickGUID {
    return JoystickGUID{*(*[16]uint8)(unsafe.Pointer(&(s.data)))}
}

func toCFromJoystickGUID(s JoystickGUID) (d C.SDL_JoystickGUID) {
    d.data = *(*[16]C.Uint8)(unsafe.Pointer(&(s.Data)))
    return
}

 // Hat positions
const (
    HAT_CENTERED = C.SDL_HAT_CENTERED

    HAT_UP = C.SDL_HAT_UP

    HAT_RIGHT = C.SDL_HAT_RIGHT

    HAT_DOWN = C.SDL_HAT_DOWN

    HAT_LEFT = C.SDL_HAT_LEFT

    HAT_RIGHTUP = C.SDL_HAT_RIGHTUP

    HAT_RIGHTDOWN = C.SDL_HAT_RIGHTDOWN

    HAT_LEFTUP = C.SDL_HAT_LEFTUP

    HAT_LEFTDOWN = C.SDL_HAT_LEFTDOWN
)

type JoystickPowerLevel int
const (
    JOYSTICK_POWER_UNKNOWN JoystickPowerLevel = C.SDL_JOYSTICK_POWER_UNKNOWN

    JOYSTICK_POWER_EMPTY JoystickPowerLevel = C.SDL_JOYSTICK_POWER_EMPTY

    JOYSTICK_POWER_LOW JoystickPowerLevel = C.SDL_JOYSTICK_POWER_LOW

    JOYSTICK_POWER_MEDIUM JoystickPowerLevel = C.SDL_JOYSTICK_POWER_MEDIUM

    JOYSTICK_POWER_FULL JoystickPowerLevel = C.SDL_JOYSTICK_POWER_FULL

    JOYSTICK_POWER_WIRED JoystickPowerLevel = C.SDL_JOYSTICK_POWER_WIRED

    JOYSTICK_POWER_MAX JoystickPowerLevel = C.SDL_JOYSTICK_POWER_MAX
)

type Joystick C.SDL_Joystick
type JoystickID int32


 // Count the number of joysticks attached to the system right now
func NumJoysticks() (retval int) {
    retval = int(C.SDL_NumJoysticks())
    return
}

 // Get the implementation dependent name of a joystick. This can be
 // called before any joysticks are opened. If no name can be found, this
 // function returns NULL.
func JoystickNameForIndex(device_index int) (retval string) {
    retval = C.GoString(C.SDL_JoystickNameForIndex(C.int(device_index)))
    return
}

 // Open a joystick for use. The index passed as an argument refers to the
 // N'th joystick on the system. This index is not the value which will
 // identify this joystick in future joystick events. The joystick's
 // instance id (SDL_JoystickID) will be used there instead.
 // 
 // Returns: A joystick identifier, or NULL if an error occurred.
 // 
func JoystickOpen(device_index int) (retval *Joystick) {
    retval = (*Joystick)(unsafe.Pointer(C.SDL_JoystickOpen(C.int(device_index))))
    return
}

 // Return the SDL_Joystick associated with an instance id.
func JoystickFromInstanceID(joyid JoystickID) (retval *Joystick) {
    retval = (*Joystick)(unsafe.Pointer(C.SDL_JoystickFromInstanceID(C.SDL_JoystickID(joyid))))
    return
}

 // Return the name for this currently opened joystick. If no name can be
 // found, this function returns NULL.
func (joystick *Joystick) Name() (retval string) {
    retval = C.GoString(C.SDL_JoystickName((*C.SDL_Joystick)(joystick)))
    return
}

 // Return the GUID for the joystick at this index
func JoystickGetDeviceGUID(device_index int) (retval JoystickGUID) {
    retval = fromC2JoystickGUID(C.SDL_JoystickGetDeviceGUID(C.int(device_index)))
    return
}

 // Return the GUID for this opened joystick
func (joystick *Joystick) GetGUID() (retval JoystickGUID) {
    retval = fromC2JoystickGUID(C.SDL_JoystickGetGUID((*C.SDL_Joystick)(joystick)))
    return
}


 // convert a string into a joystick formatted guid
func JoystickGetGUIDFromString(pchGUID string) (retval JoystickGUID) {
    tmp_pchGUID := C.CString(pchGUID); defer C.free(unsafe.Pointer(tmp_pchGUID))
    retval = fromC2JoystickGUID(C.SDL_JoystickGetGUIDFromString((*C.char)(tmp_pchGUID)))
    return
}

 // Returns SDL_TRUE if the joystick has been opened and currently
 // connected, or SDL_FALSE if it has not.
func (joystick *Joystick) GetAttached() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_JoystickGetAttached((*C.SDL_Joystick)(joystick)))
    return
}

 // Get the instance ID of an opened joystick or -1 if the joystick is
 // invalid.
func (joystick *Joystick) InstanceID() (retval JoystickID) {
    retval = JoystickID(C.SDL_JoystickInstanceID((*C.SDL_Joystick)(joystick)))
    return
}

 // Get the number of general axis controls on a joystick.
func (joystick *Joystick) NumAxes() (retval int) {
    retval = int(C.SDL_JoystickNumAxes((*C.SDL_Joystick)(joystick)))
    return
}

 // Get the number of trackballs on a joystick.
 // 
 // Joystick trackballs have only relative motion events associated with
 // them and their state cannot be polled.
func (joystick *Joystick) NumBalls() (retval int) {
    retval = int(C.SDL_JoystickNumBalls((*C.SDL_Joystick)(joystick)))
    return
}

 // Get the number of POV hats on a joystick.
func (joystick *Joystick) NumHats() (retval int) {
    retval = int(C.SDL_JoystickNumHats((*C.SDL_Joystick)(joystick)))
    return
}

 // Get the number of buttons on a joystick.
func (joystick *Joystick) NumButtons() (retval int) {
    retval = int(C.SDL_JoystickNumButtons((*C.SDL_Joystick)(joystick)))
    return
}

 // Update the current state of the open joysticks.
 // 
 // This is called automatically by the event loop if any joystick events
 // are enabled.
func JoystickUpdate() {
    C.SDL_JoystickUpdate()
}

 // Enable/disable joystick event polling.
 // 
 // If joystick events are disabled, you must call SDL_JoystickUpdate()
 // yourself and check the state of the joystick when you want joystick
 // information.
 // 
 // The state can be one of SDL_QUERY, SDL_ENABLE or SDL_IGNORE.
func JoystickEventState(state int) (retval int) {
    retval = int(C.SDL_JoystickEventState(C.int(state)))
    return
}

 // Get the current state of an axis control on a joystick.
 // 
 // The state is a value ranging from -32768 to 32767.
 // 
 // The axis indices start at index 0.
func (joystick *Joystick) GetAxis(axis int) (retval int16) {
    retval = int16(C.SDL_JoystickGetAxis((*C.SDL_Joystick)(joystick), C.int(axis)))
    return
}

 // Get the current state of a POV hat on a joystick.
 // 
 // The hat indices start at index 0.
 // 
 // Returns: The return value is one of the following positions:
 //   
 //   - SDL_HAT_CENTERED
 //   - SDL_HAT_UP
 //   - SDL_HAT_RIGHT
 //   - SDL_HAT_DOWN
 //   - SDL_HAT_LEFT
 //   - SDL_HAT_RIGHTUP
 //   - SDL_HAT_RIGHTDOWN
 //   - SDL_HAT_LEFTUP
 //   - SDL_HAT_LEFTDOWN
 // 
func (joystick *Joystick) GetHat(hat int) (retval uint8) {
    retval = uint8(C.SDL_JoystickGetHat((*C.SDL_Joystick)(joystick), C.int(hat)))
    return
}

 // Get the ball axis change since the last poll.
 // 
 // Returns: 0, or -1 if you passed it invalid parameters.
 // 
 // The ball indices start at index 0.
func (joystick *Joystick) GetBall(ball int) (retval int, dx int, dy int) {
    tmp_dx := new(C.int)
    tmp_dy := new(C.int)
    retval = int(C.SDL_JoystickGetBall((*C.SDL_Joystick)(joystick), C.int(ball), (*C.int)(tmp_dx), (*C.int)(tmp_dy)))
    dx = deref_int_ptr(tmp_dx)
    dy = deref_int_ptr(tmp_dy)
    return
}

 // Get the current state of a button on a joystick.
 // 
 // The button indices start at index 0.
func (joystick *Joystick) GetButton(button int) (retval uint8) {
    retval = uint8(C.SDL_JoystickGetButton((*C.SDL_Joystick)(joystick), C.int(button)))
    return
}

 // Close a joystick previously opened with SDL_JoystickOpen().
func (joystick *Joystick) Close() {
    C.SDL_JoystickClose((*C.SDL_Joystick)(joystick))
}

 // Return the battery level of this joystick
func (joystick *Joystick) CurrentPowerLevel() (retval JoystickPowerLevel) {
    retval = JoystickPowerLevel(C.SDL_JoystickCurrentPowerLevel((*C.SDL_Joystick)(joystick)))
    return
}
