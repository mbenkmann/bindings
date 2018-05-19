// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
import "C"
import "unsafe"

 // Include file for SDL event handling.

 // Fields shared by every event.
type CommonEvent struct {
    Type uint32

    Timestamp uint32
}

func fromC2CommonEvent(s C.SDL_CommonEvent) CommonEvent {
    return CommonEvent{uint32(s._type), uint32(s.timestamp)}
}

func toCFromCommonEvent(s CommonEvent) (d C.SDL_CommonEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    return
}

 // Window state change event data (event.window.*)
type WindowEvent struct {
     // SDL_WINDOWEVENT
    Type uint32

    Timestamp uint32

     // The associated window
    WindowID uint32

     // SDL_WindowEventID
    Event uint8

    Padding1 uint8

    Padding2 uint8

    Padding3 uint8

     // event dependent data
    Data1 int32

     // event dependent data
    Data2 int32
}

func fromC2WindowEvent(s C.SDL_WindowEvent) WindowEvent {
    return WindowEvent{uint32(s._type), uint32(s.timestamp), uint32(s.windowID), uint8(s.event), uint8(s.padding1), uint8(s.padding2), uint8(s.padding3), int32(s.data1), int32(s.data2)}
}

func toCFromWindowEvent(s WindowEvent) (d C.SDL_WindowEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.windowID = C.Uint32(s.WindowID)
    d.event = C.Uint8(s.Event)
    d.padding1 = C.Uint8(s.Padding1)
    d.padding2 = C.Uint8(s.Padding2)
    d.padding3 = C.Uint8(s.Padding3)
    d.data1 = C.Sint32(s.Data1)
    d.data2 = C.Sint32(s.Data2)
    return
}

 // Keyboard button event structure (event.key.*)
type KeyboardEvent struct {
     // SDL_KEYDOWN or SDL_KEYUP
    Type uint32

    Timestamp uint32

     // The window with keyboard focus, if any
    WindowID uint32

     // SDL_PRESSED or SDL_RELEASED
    State uint8

     // Non-zero if this is a key repeat
    Repeat uint8

    Padding2 uint8

    Padding3 uint8

     // The key that was pressed or released
    Keysym Keysym
}

func fromC2KeyboardEvent(s C.SDL_KeyboardEvent) KeyboardEvent {
    return KeyboardEvent{uint32(s._type), uint32(s.timestamp), uint32(s.windowID), uint8(s.state), uint8(s.repeat), uint8(s.padding2), uint8(s.padding3), Keysym{Scancode(s.keysym.scancode), Keycode(s.keysym.sym), uint16(s.keysym.mod), uint32(s.keysym.unused)}}
}

func toCFromKeyboardEvent(s KeyboardEvent) (d C.SDL_KeyboardEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.windowID = C.Uint32(s.WindowID)
    d.state = C.Uint8(s.State)
    d.repeat = C.Uint8(s.Repeat)
    d.padding2 = C.Uint8(s.Padding2)
    d.padding3 = C.Uint8(s.Padding3)
    d.keysym.scancode = C.SDL_Scancode(s.Keysym.Scancode)
    d.keysym.sym = C.SDL_Keycode(s.Keysym.Sym)
    d.keysym.mod = C.Uint16(s.Keysym.Mod)
    d.keysym.unused = C.Uint32(s.Keysym.Unused)
    return
}

 // Keyboard text editing event structure (event.edit.*)
type TextEditingEvent struct {
     // SDL_TEXTEDITING
    Type uint32

    Timestamp uint32

     // The window with keyboard focus, if any
    WindowID uint32

     // The editing text
    Text [TEXTEDITINGEVENT_TEXT_SIZE]int8

     // The start cursor of selected editing text
    Start int32

     // The length of selected editing text
    Length int32
}

func fromC2TextEditingEvent(s C.SDL_TextEditingEvent) TextEditingEvent {
    return TextEditingEvent{uint32(s._type), uint32(s.timestamp), uint32(s.windowID), *(*[TEXTEDITINGEVENT_TEXT_SIZE]int8)(unsafe.Pointer(&(s.text))), int32(s.start), int32(s.length)}
}

func toCFromTextEditingEvent(s TextEditingEvent) (d C.SDL_TextEditingEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.windowID = C.Uint32(s.WindowID)
    d.text = *(*[TEXTEDITINGEVENT_TEXT_SIZE]C.char)(unsafe.Pointer(&(s.Text)))
    d.start = C.Sint32(s.Start)
    d.length = C.Sint32(s.Length)
    return
}

 // Keyboard text input event structure (event.text.*)
type TextInputEvent struct {
     // SDL_TEXTINPUT
    Type uint32

    Timestamp uint32

     // The window with keyboard focus, if any
    WindowID uint32

     // The input text
    Text [TEXTINPUTEVENT_TEXT_SIZE]int8
}

func fromC2TextInputEvent(s C.SDL_TextInputEvent) TextInputEvent {
    return TextInputEvent{uint32(s._type), uint32(s.timestamp), uint32(s.windowID), *(*[TEXTINPUTEVENT_TEXT_SIZE]int8)(unsafe.Pointer(&(s.text)))}
}

func toCFromTextInputEvent(s TextInputEvent) (d C.SDL_TextInputEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.windowID = C.Uint32(s.WindowID)
    d.text = *(*[TEXTINPUTEVENT_TEXT_SIZE]C.char)(unsafe.Pointer(&(s.Text)))
    return
}

 // Mouse motion event structure (event.motion.*)
type MouseMotionEvent struct {
     // SDL_MOUSEMOTION
    Type uint32

    Timestamp uint32

     // The window with mouse focus, if any
    WindowID uint32

     // The mouse instance id, or SDL_TOUCH_MOUSEID
    Which uint32

     // The current button state
    State uint32

     // X coordinate, relative to window
    X int32

     // Y coordinate, relative to window
    Y int32

     // The relative motion in the X direction
    Xrel int32

     // The relative motion in the Y direction
    Yrel int32
}

func fromC2MouseMotionEvent(s C.SDL_MouseMotionEvent) MouseMotionEvent {
    return MouseMotionEvent{uint32(s._type), uint32(s.timestamp), uint32(s.windowID), uint32(s.which), uint32(s.state), int32(s.x), int32(s.y), int32(s.xrel), int32(s.yrel)}
}

func toCFromMouseMotionEvent(s MouseMotionEvent) (d C.SDL_MouseMotionEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.windowID = C.Uint32(s.WindowID)
    d.which = C.Uint32(s.Which)
    d.state = C.Uint32(s.State)
    d.x = C.Sint32(s.X)
    d.y = C.Sint32(s.Y)
    d.xrel = C.Sint32(s.Xrel)
    d.yrel = C.Sint32(s.Yrel)
    return
}

 // Mouse button event structure (event.button.*)
type MouseButtonEvent struct {
     // SDL_MOUSEBUTTONDOWN or SDL_MOUSEBUTTONUP
    Type uint32

    Timestamp uint32

     // The window with mouse focus, if any
    WindowID uint32

     // The mouse instance id, or SDL_TOUCH_MOUSEID
    Which uint32

     // The mouse button index
    Button uint8

     // SDL_PRESSED or SDL_RELEASED
    State uint8

     // 1 for single-click, 2 for double-click, etc.
    Clicks uint8

    Padding1 uint8

     // X coordinate, relative to window
    X int32

     // Y coordinate, relative to window
    Y int32
}

func fromC2MouseButtonEvent(s C.SDL_MouseButtonEvent) MouseButtonEvent {
    return MouseButtonEvent{uint32(s._type), uint32(s.timestamp), uint32(s.windowID), uint32(s.which), uint8(s.button), uint8(s.state), uint8(s.clicks), uint8(s.padding1), int32(s.x), int32(s.y)}
}

func toCFromMouseButtonEvent(s MouseButtonEvent) (d C.SDL_MouseButtonEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.windowID = C.Uint32(s.WindowID)
    d.which = C.Uint32(s.Which)
    d.button = C.Uint8(s.Button)
    d.state = C.Uint8(s.State)
    d.clicks = C.Uint8(s.Clicks)
    d.padding1 = C.Uint8(s.Padding1)
    d.x = C.Sint32(s.X)
    d.y = C.Sint32(s.Y)
    return
}

 // Mouse wheel event structure (event.wheel.*)
type MouseWheelEvent struct {
     // SDL_MOUSEWHEEL
    Type uint32

    Timestamp uint32

     // The window with mouse focus, if any
    WindowID uint32

     // The mouse instance id, or SDL_TOUCH_MOUSEID
    Which uint32

     // The amount scrolled horizontally, positive to the right and negative
     // to the left
    X int32

     // The amount scrolled vertically, positive away from the user and
     // negative toward the user
    Y int32

     // Set to one of the SDL_MOUSEWHEEL_* defines. When FLIPPED the values in
     // X and Y will be opposite. Multiply by -1 to change them back
    Direction uint32
}

func fromC2MouseWheelEvent(s C.SDL_MouseWheelEvent) MouseWheelEvent {
    return MouseWheelEvent{uint32(s._type), uint32(s.timestamp), uint32(s.windowID), uint32(s.which), int32(s.x), int32(s.y), uint32(s.direction)}
}

func toCFromMouseWheelEvent(s MouseWheelEvent) (d C.SDL_MouseWheelEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.windowID = C.Uint32(s.WindowID)
    d.which = C.Uint32(s.Which)
    d.x = C.Sint32(s.X)
    d.y = C.Sint32(s.Y)
    d.direction = C.Uint32(s.Direction)
    return
}

 // Joystick axis motion event structure (event.jaxis.*)
type JoyAxisEvent struct {
     // SDL_JOYAXISMOTION
    Type uint32

    Timestamp uint32

     // The joystick instance id
    Which JoystickID

     // The joystick axis index
    Axis uint8

    Padding1 uint8

    Padding2 uint8

    Padding3 uint8

     // The axis value (range: -32768 to 32767)
    Value int16

    Padding4 uint16
}

func fromC2JoyAxisEvent(s C.SDL_JoyAxisEvent) JoyAxisEvent {
    return JoyAxisEvent{uint32(s._type), uint32(s.timestamp), JoystickID(s.which), uint8(s.axis), uint8(s.padding1), uint8(s.padding2), uint8(s.padding3), int16(s.value), uint16(s.padding4)}
}

func toCFromJoyAxisEvent(s JoyAxisEvent) (d C.SDL_JoyAxisEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.which = C.SDL_JoystickID(s.Which)
    d.axis = C.Uint8(s.Axis)
    d.padding1 = C.Uint8(s.Padding1)
    d.padding2 = C.Uint8(s.Padding2)
    d.padding3 = C.Uint8(s.Padding3)
    d.value = C.Sint16(s.Value)
    d.padding4 = C.Uint16(s.Padding4)
    return
}

 // Joystick trackball motion event structure (event.jball.*)
type JoyBallEvent struct {
     // SDL_JOYBALLMOTION
    Type uint32

    Timestamp uint32

     // The joystick instance id
    Which JoystickID

     // The joystick trackball index
    Ball uint8

    Padding1 uint8

    Padding2 uint8

    Padding3 uint8

     // The relative motion in the X direction
    Xrel int16

     // The relative motion in the Y direction
    Yrel int16
}

func fromC2JoyBallEvent(s C.SDL_JoyBallEvent) JoyBallEvent {
    return JoyBallEvent{uint32(s._type), uint32(s.timestamp), JoystickID(s.which), uint8(s.ball), uint8(s.padding1), uint8(s.padding2), uint8(s.padding3), int16(s.xrel), int16(s.yrel)}
}

func toCFromJoyBallEvent(s JoyBallEvent) (d C.SDL_JoyBallEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.which = C.SDL_JoystickID(s.Which)
    d.ball = C.Uint8(s.Ball)
    d.padding1 = C.Uint8(s.Padding1)
    d.padding2 = C.Uint8(s.Padding2)
    d.padding3 = C.Uint8(s.Padding3)
    d.xrel = C.Sint16(s.Xrel)
    d.yrel = C.Sint16(s.Yrel)
    return
}

 // Joystick hat position change event structure (event.jhat.*)
type JoyHatEvent struct {
     // SDL_JOYHATMOTION
    Type uint32

    Timestamp uint32

     // The joystick instance id
    Which JoystickID

     // The joystick hat index
    Hat uint8

     // See also: SDL_HAT_LEFTUP SDL_HAT_UP SDL_HAT_RIGHTUP
     // 
     // See also: SDL_HAT_LEFT SDL_HAT_CENTERED SDL_HAT_RIGHT
     // 
     // See also: SDL_HAT_LEFTDOWN SDL_HAT_DOWN SDL_HAT_RIGHTDOWN
     // 
     // The hat position value.    Note that zero means the POV is centered.
    Value uint8

    Padding1 uint8

    Padding2 uint8
}

func fromC2JoyHatEvent(s C.SDL_JoyHatEvent) JoyHatEvent {
    return JoyHatEvent{uint32(s._type), uint32(s.timestamp), JoystickID(s.which), uint8(s.hat), uint8(s.value), uint8(s.padding1), uint8(s.padding2)}
}

func toCFromJoyHatEvent(s JoyHatEvent) (d C.SDL_JoyHatEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.which = C.SDL_JoystickID(s.Which)
    d.hat = C.Uint8(s.Hat)
    d.value = C.Uint8(s.Value)
    d.padding1 = C.Uint8(s.Padding1)
    d.padding2 = C.Uint8(s.Padding2)
    return
}

 // Joystick button event structure (event.jbutton.*)
type JoyButtonEvent struct {
     // SDL_JOYBUTTONDOWN or SDL_JOYBUTTONUP
    Type uint32

    Timestamp uint32

     // The joystick instance id
    Which JoystickID

     // The joystick button index
    Button uint8

     // SDL_PRESSED or SDL_RELEASED
    State uint8

    Padding1 uint8

    Padding2 uint8
}

func fromC2JoyButtonEvent(s C.SDL_JoyButtonEvent) JoyButtonEvent {
    return JoyButtonEvent{uint32(s._type), uint32(s.timestamp), JoystickID(s.which), uint8(s.button), uint8(s.state), uint8(s.padding1), uint8(s.padding2)}
}

func toCFromJoyButtonEvent(s JoyButtonEvent) (d C.SDL_JoyButtonEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.which = C.SDL_JoystickID(s.Which)
    d.button = C.Uint8(s.Button)
    d.state = C.Uint8(s.State)
    d.padding1 = C.Uint8(s.Padding1)
    d.padding2 = C.Uint8(s.Padding2)
    return
}

 // Joystick device event structure (event.jdevice.*)
type JoyDeviceEvent struct {
     // SDL_JOYDEVICEADDED or SDL_JOYDEVICEREMOVED
    Type uint32

    Timestamp uint32

     // The joystick device index for the ADDED event, instance id for the
     // REMOVED event
    Which int32
}

func fromC2JoyDeviceEvent(s C.SDL_JoyDeviceEvent) JoyDeviceEvent {
    return JoyDeviceEvent{uint32(s._type), uint32(s.timestamp), int32(s.which)}
}

func toCFromJoyDeviceEvent(s JoyDeviceEvent) (d C.SDL_JoyDeviceEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.which = C.Sint32(s.Which)
    return
}

 // Game controller axis motion event structure (event.caxis.*)
type ControllerAxisEvent struct {
     // SDL_CONTROLLERAXISMOTION
    Type uint32

    Timestamp uint32

     // The joystick instance id
    Which JoystickID

     // The controller axis (SDL_GameControllerAxis)
    Axis uint8

    Padding1 uint8

    Padding2 uint8

    Padding3 uint8

     // The axis value (range: -32768 to 32767)
    Value int16

    Padding4 uint16
}

func fromC2ControllerAxisEvent(s C.SDL_ControllerAxisEvent) ControllerAxisEvent {
    return ControllerAxisEvent{uint32(s._type), uint32(s.timestamp), JoystickID(s.which), uint8(s.axis), uint8(s.padding1), uint8(s.padding2), uint8(s.padding3), int16(s.value), uint16(s.padding4)}
}

func toCFromControllerAxisEvent(s ControllerAxisEvent) (d C.SDL_ControllerAxisEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.which = C.SDL_JoystickID(s.Which)
    d.axis = C.Uint8(s.Axis)
    d.padding1 = C.Uint8(s.Padding1)
    d.padding2 = C.Uint8(s.Padding2)
    d.padding3 = C.Uint8(s.Padding3)
    d.value = C.Sint16(s.Value)
    d.padding4 = C.Uint16(s.Padding4)
    return
}

 // Game controller button event structure (event.cbutton.*)
type ControllerButtonEvent struct {
     // SDL_CONTROLLERBUTTONDOWN or SDL_CONTROLLERBUTTONUP
    Type uint32

    Timestamp uint32

     // The joystick instance id
    Which JoystickID

     // The controller button (SDL_GameControllerButton)
    Button uint8

     // SDL_PRESSED or SDL_RELEASED
    State uint8

    Padding1 uint8

    Padding2 uint8
}

func fromC2ControllerButtonEvent(s C.SDL_ControllerButtonEvent) ControllerButtonEvent {
    return ControllerButtonEvent{uint32(s._type), uint32(s.timestamp), JoystickID(s.which), uint8(s.button), uint8(s.state), uint8(s.padding1), uint8(s.padding2)}
}

func toCFromControllerButtonEvent(s ControllerButtonEvent) (d C.SDL_ControllerButtonEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.which = C.SDL_JoystickID(s.Which)
    d.button = C.Uint8(s.Button)
    d.state = C.Uint8(s.State)
    d.padding1 = C.Uint8(s.Padding1)
    d.padding2 = C.Uint8(s.Padding2)
    return
}

 // Controller device event structure (event.cdevice.*)
type ControllerDeviceEvent struct {
     // SDL_CONTROLLERDEVICEADDED, SDL_CONTROLLERDEVICEREMOVED, or
     // SDL_CONTROLLERDEVICEREMAPPED
    Type uint32

    Timestamp uint32

     // The joystick device index for the ADDED event, instance id for the
     // REMOVED or REMAPPED event
    Which int32
}

func fromC2ControllerDeviceEvent(s C.SDL_ControllerDeviceEvent) ControllerDeviceEvent {
    return ControllerDeviceEvent{uint32(s._type), uint32(s.timestamp), int32(s.which)}
}

func toCFromControllerDeviceEvent(s ControllerDeviceEvent) (d C.SDL_ControllerDeviceEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.which = C.Sint32(s.Which)
    return
}

 // Audio device event structure (event.adevice.*)
type AudioDeviceEvent struct {
     // SDL_AUDIODEVICEADDED, or SDL_AUDIODEVICEREMOVED
    Type uint32

    Timestamp uint32

     // The audio device index for the ADDED event (valid until next
     // SDL_GetNumAudioDevices() call), SDL_AudioDeviceID for the REMOVED
     // event
    Which uint32

     // zero if an output device, non-zero if a capture device.
    Iscapture uint8

    Padding1 uint8

    Padding2 uint8

    Padding3 uint8
}

func fromC2AudioDeviceEvent(s C.SDL_AudioDeviceEvent) AudioDeviceEvent {
    return AudioDeviceEvent{uint32(s._type), uint32(s.timestamp), uint32(s.which), uint8(s.iscapture), uint8(s.padding1), uint8(s.padding2), uint8(s.padding3)}
}

func toCFromAudioDeviceEvent(s AudioDeviceEvent) (d C.SDL_AudioDeviceEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.which = C.Uint32(s.Which)
    d.iscapture = C.Uint8(s.Iscapture)
    d.padding1 = C.Uint8(s.Padding1)
    d.padding2 = C.Uint8(s.Padding2)
    d.padding3 = C.Uint8(s.Padding3)
    return
}

 // Touch finger event structure (event.tfinger.*)
type TouchFingerEvent struct {
     // SDL_FINGERMOTION or SDL_FINGERDOWN or SDL_FINGERUP
    Type uint32

    Timestamp uint32

     // The touch device id
    TouchId TouchID

    FingerId FingerID

     // Normalized in the range 0...1
    X float32

     // Normalized in the range 0...1
    Y float32

     // Normalized in the range -1...1
    Dx float32

     // Normalized in the range -1...1
    Dy float32

     // Normalized in the range 0...1
    Pressure float32
}

func fromC2TouchFingerEvent(s C.SDL_TouchFingerEvent) TouchFingerEvent {
    return TouchFingerEvent{uint32(s._type), uint32(s.timestamp), TouchID(s.touchId), FingerID(s.fingerId), float32(s.x), float32(s.y), float32(s.dx), float32(s.dy), float32(s.pressure)}
}

func toCFromTouchFingerEvent(s TouchFingerEvent) (d C.SDL_TouchFingerEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.touchId = C.SDL_TouchID(s.TouchId)
    d.fingerId = C.SDL_FingerID(s.FingerId)
    d.x = C.float(s.X)
    d.y = C.float(s.Y)
    d.dx = C.float(s.Dx)
    d.dy = C.float(s.Dy)
    d.pressure = C.float(s.Pressure)
    return
}

 // Multiple Finger Gesture Event (event.mgesture.*)
type MultiGestureEvent struct {
     // SDL_MULTIGESTURE
    Type uint32

    Timestamp uint32

     // The touch device index
    TouchId TouchID

    DTheta float32

    DDist float32

    X float32

    Y float32

    NumFingers uint16

    Padding uint16
}

func fromC2MultiGestureEvent(s C.SDL_MultiGestureEvent) MultiGestureEvent {
    return MultiGestureEvent{uint32(s._type), uint32(s.timestamp), TouchID(s.touchId), float32(s.dTheta), float32(s.dDist), float32(s.x), float32(s.y), uint16(s.numFingers), uint16(s.padding)}
}

func toCFromMultiGestureEvent(s MultiGestureEvent) (d C.SDL_MultiGestureEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.touchId = C.SDL_TouchID(s.TouchId)
    d.dTheta = C.float(s.DTheta)
    d.dDist = C.float(s.DDist)
    d.x = C.float(s.X)
    d.y = C.float(s.Y)
    d.numFingers = C.Uint16(s.NumFingers)
    d.padding = C.Uint16(s.Padding)
    return
}

 // Dollar Gesture Event (event.dgesture.*)
type DollarGestureEvent struct {
     // SDL_DOLLARGESTURE or SDL_DOLLARRECORD
    Type uint32

    Timestamp uint32

     // The touch device id
    TouchId TouchID

    GestureId GestureID

    NumFingers uint32

    Error float32

     // Normalized center of gesture
    X float32

     // Normalized center of gesture
    Y float32
}

func fromC2DollarGestureEvent(s C.SDL_DollarGestureEvent) DollarGestureEvent {
    return DollarGestureEvent{uint32(s._type), uint32(s.timestamp), TouchID(s.touchId), GestureID(s.gestureId), uint32(s.numFingers), float32(s.error), float32(s.x), float32(s.y)}
}

func toCFromDollarGestureEvent(s DollarGestureEvent) (d C.SDL_DollarGestureEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.touchId = C.SDL_TouchID(s.TouchId)
    d.gestureId = C.SDL_GestureID(s.GestureId)
    d.numFingers = C.Uint32(s.NumFingers)
    d.error = C.float(s.Error)
    d.x = C.float(s.X)
    d.y = C.float(s.Y)
    return
}

 // An event used to request a file open by the system (event.drop.*) This
 // event is enabled by default, you can disable it with SDL_EventState().
 // 
 // Note: If this event is enabled, you must free the filename in the
 // event.
 // 
type DropEvent struct {
     // SDL_DROPFILE
    Type uint32

    Timestamp uint32

     // The file name, which should be freed with SDL_free()
    File string
}

func fromC2DropEvent(s C.SDL_DropEvent) DropEvent {
    return DropEvent{uint32(s._type), uint32(s.timestamp), C.GoString(s.file)}
}

 // The "quit requested" event.
type QuitEvent struct {
     // SDL_QUIT
    Type uint32

    Timestamp uint32
}

func fromC2QuitEvent(s C.SDL_QuitEvent) QuitEvent {
    return QuitEvent{uint32(s._type), uint32(s.timestamp)}
}

func toCFromQuitEvent(s QuitEvent) (d C.SDL_QuitEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    return
}

 // OS Specific event.
type OSEvent struct {
     // SDL_QUIT
    Type uint32

    Timestamp uint32
}

func fromC2OSEvent(s C.SDL_OSEvent) OSEvent {
    return OSEvent{uint32(s._type), uint32(s.timestamp)}
}

func toCFromOSEvent(s OSEvent) (d C.SDL_OSEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    return
}

 // A user-defined event type (event.user.*)
type UserEvent struct {
     // SDL_USEREVENT through SDL_LASTEVENT-1
    Type uint32

    Timestamp uint32

     // The associated window if any
    WindowID uint32

     // User defined event code
    Code int32

     // User defined data pointer
    Data1 uintptr

     // User defined data pointer
    Data2 uintptr
}

func fromC2UserEvent(s C.SDL_UserEvent) UserEvent {
    return UserEvent{uint32(s._type), uint32(s.timestamp), uint32(s.windowID), int32(s.code), uintptr(s.data1), uintptr(s.data2)}
}

func toCFromUserEvent(s UserEvent) (d C.SDL_UserEvent) {
    d._type = C.Uint32(s.Type)
    d.timestamp = C.Uint32(s.Timestamp)
    d.windowID = C.Uint32(s.WindowID)
    d.code = C.Sint32(s.Code)
    d.data1 = unsafe.Pointer(s.Data1)
    d.data2 = unsafe.Pointer(s.Data2)
    return
}

 // General event structure.
type Event C.SDL_Event

 // Event type, shared with all events
func (u *Event) Type() uint32 {
    p := (*C.Uint32)(unsafe.Pointer(u))
    return uint32(*p)
}
 // Event type, shared with all events
func (u *Event) SetType(x uint32) {
    p := (*C.Uint32)(unsafe.Pointer(u))
    *p = C.Uint32(x)
}

 // Common event data
func (u *Event) Common() CommonEvent {
    p := (*C.SDL_CommonEvent)(unsafe.Pointer(u))
    return CommonEvent{uint32(p._type), uint32(p.timestamp)}
}
 // Common event data
func (u *Event) SetCommon(x CommonEvent) {
    p := (*C.SDL_CommonEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
}

 // Window event data
func (u *Event) Window() WindowEvent {
    p := (*C.SDL_WindowEvent)(unsafe.Pointer(u))
    return WindowEvent{uint32(p._type), uint32(p.timestamp), uint32(p.windowID), uint8(p.event), uint8(p.padding1), uint8(p.padding2), uint8(p.padding3), int32(p.data1), int32(p.data2)}
}
 // Window event data
func (u *Event) SetWindow(x WindowEvent) {
    p := (*C.SDL_WindowEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.windowID = C.Uint32(x.WindowID)
    p.event = C.Uint8(x.Event)
    p.padding1 = C.Uint8(x.Padding1)
    p.padding2 = C.Uint8(x.Padding2)
    p.padding3 = C.Uint8(x.Padding3)
    p.data1 = C.Sint32(x.Data1)
    p.data2 = C.Sint32(x.Data2)
}

 // Keyboard event data
func (u *Event) Key() KeyboardEvent {
    p := (*C.SDL_KeyboardEvent)(unsafe.Pointer(u))
    return KeyboardEvent{uint32(p._type), uint32(p.timestamp), uint32(p.windowID), uint8(p.state), uint8(p.repeat), uint8(p.padding2), uint8(p.padding3), Keysym{Scancode(p.keysym.scancode), Keycode(p.keysym.sym), uint16(p.keysym.mod), uint32(p.keysym.unused)}}
}
 // Keyboard event data
func (u *Event) SetKey(x KeyboardEvent) {
    p := (*C.SDL_KeyboardEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.windowID = C.Uint32(x.WindowID)
    p.state = C.Uint8(x.State)
    p.repeat = C.Uint8(x.Repeat)
    p.padding2 = C.Uint8(x.Padding2)
    p.padding3 = C.Uint8(x.Padding3)
    p.keysym.scancode = C.SDL_Scancode(x.Keysym.Scancode)
    p.keysym.sym = C.SDL_Keycode(x.Keysym.Sym)
    p.keysym.mod = C.Uint16(x.Keysym.Mod)
    p.keysym.unused = C.Uint32(x.Keysym.Unused)
}

 // Text editing event data
func (u *Event) Edit() TextEditingEvent {
    p := (*C.SDL_TextEditingEvent)(unsafe.Pointer(u))
    return TextEditingEvent{uint32(p._type), uint32(p.timestamp), uint32(p.windowID), *(*[TEXTEDITINGEVENT_TEXT_SIZE]int8)(unsafe.Pointer(&(p.text))), int32(p.start), int32(p.length)}
}
 // Text editing event data
func (u *Event) SetEdit(x TextEditingEvent) {
    p := (*C.SDL_TextEditingEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.windowID = C.Uint32(x.WindowID)
    p.text = *(*[TEXTEDITINGEVENT_TEXT_SIZE]C.char)(unsafe.Pointer(&(x.Text)))
    p.start = C.Sint32(x.Start)
    p.length = C.Sint32(x.Length)
}

 // Text input event data
func (u *Event) Text() TextInputEvent {
    p := (*C.SDL_TextInputEvent)(unsafe.Pointer(u))
    return TextInputEvent{uint32(p._type), uint32(p.timestamp), uint32(p.windowID), *(*[TEXTINPUTEVENT_TEXT_SIZE]int8)(unsafe.Pointer(&(p.text)))}
}
 // Text input event data
func (u *Event) SetText(x TextInputEvent) {
    p := (*C.SDL_TextInputEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.windowID = C.Uint32(x.WindowID)
    p.text = *(*[TEXTINPUTEVENT_TEXT_SIZE]C.char)(unsafe.Pointer(&(x.Text)))
}

 // Mouse motion event data
func (u *Event) Motion() MouseMotionEvent {
    p := (*C.SDL_MouseMotionEvent)(unsafe.Pointer(u))
    return MouseMotionEvent{uint32(p._type), uint32(p.timestamp), uint32(p.windowID), uint32(p.which), uint32(p.state), int32(p.x), int32(p.y), int32(p.xrel), int32(p.yrel)}
}
 // Mouse motion event data
func (u *Event) SetMotion(x MouseMotionEvent) {
    p := (*C.SDL_MouseMotionEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.windowID = C.Uint32(x.WindowID)
    p.which = C.Uint32(x.Which)
    p.state = C.Uint32(x.State)
    p.x = C.Sint32(x.X)
    p.y = C.Sint32(x.Y)
    p.xrel = C.Sint32(x.Xrel)
    p.yrel = C.Sint32(x.Yrel)
}

 // Mouse button event data
func (u *Event) Button() MouseButtonEvent {
    p := (*C.SDL_MouseButtonEvent)(unsafe.Pointer(u))
    return MouseButtonEvent{uint32(p._type), uint32(p.timestamp), uint32(p.windowID), uint32(p.which), uint8(p.button), uint8(p.state), uint8(p.clicks), uint8(p.padding1), int32(p.x), int32(p.y)}
}
 // Mouse button event data
func (u *Event) SetButton(x MouseButtonEvent) {
    p := (*C.SDL_MouseButtonEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.windowID = C.Uint32(x.WindowID)
    p.which = C.Uint32(x.Which)
    p.button = C.Uint8(x.Button)
    p.state = C.Uint8(x.State)
    p.clicks = C.Uint8(x.Clicks)
    p.padding1 = C.Uint8(x.Padding1)
    p.x = C.Sint32(x.X)
    p.y = C.Sint32(x.Y)
}

 // Mouse wheel event data
func (u *Event) Wheel() MouseWheelEvent {
    p := (*C.SDL_MouseWheelEvent)(unsafe.Pointer(u))
    return MouseWheelEvent{uint32(p._type), uint32(p.timestamp), uint32(p.windowID), uint32(p.which), int32(p.x), int32(p.y), uint32(p.direction)}
}
 // Mouse wheel event data
func (u *Event) SetWheel(x MouseWheelEvent) {
    p := (*C.SDL_MouseWheelEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.windowID = C.Uint32(x.WindowID)
    p.which = C.Uint32(x.Which)
    p.x = C.Sint32(x.X)
    p.y = C.Sint32(x.Y)
    p.direction = C.Uint32(x.Direction)
}

 // Joystick axis event data
func (u *Event) Jaxis() JoyAxisEvent {
    p := (*C.SDL_JoyAxisEvent)(unsafe.Pointer(u))
    return JoyAxisEvent{uint32(p._type), uint32(p.timestamp), JoystickID(p.which), uint8(p.axis), uint8(p.padding1), uint8(p.padding2), uint8(p.padding3), int16(p.value), uint16(p.padding4)}
}
 // Joystick axis event data
func (u *Event) SetJaxis(x JoyAxisEvent) {
    p := (*C.SDL_JoyAxisEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.which = C.SDL_JoystickID(x.Which)
    p.axis = C.Uint8(x.Axis)
    p.padding1 = C.Uint8(x.Padding1)
    p.padding2 = C.Uint8(x.Padding2)
    p.padding3 = C.Uint8(x.Padding3)
    p.value = C.Sint16(x.Value)
    p.padding4 = C.Uint16(x.Padding4)
}

 // Joystick ball event data
func (u *Event) Jball() JoyBallEvent {
    p := (*C.SDL_JoyBallEvent)(unsafe.Pointer(u))
    return JoyBallEvent{uint32(p._type), uint32(p.timestamp), JoystickID(p.which), uint8(p.ball), uint8(p.padding1), uint8(p.padding2), uint8(p.padding3), int16(p.xrel), int16(p.yrel)}
}
 // Joystick ball event data
func (u *Event) SetJball(x JoyBallEvent) {
    p := (*C.SDL_JoyBallEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.which = C.SDL_JoystickID(x.Which)
    p.ball = C.Uint8(x.Ball)
    p.padding1 = C.Uint8(x.Padding1)
    p.padding2 = C.Uint8(x.Padding2)
    p.padding3 = C.Uint8(x.Padding3)
    p.xrel = C.Sint16(x.Xrel)
    p.yrel = C.Sint16(x.Yrel)
}

 // Joystick hat event data
func (u *Event) Jhat() JoyHatEvent {
    p := (*C.SDL_JoyHatEvent)(unsafe.Pointer(u))
    return JoyHatEvent{uint32(p._type), uint32(p.timestamp), JoystickID(p.which), uint8(p.hat), uint8(p.value), uint8(p.padding1), uint8(p.padding2)}
}
 // Joystick hat event data
func (u *Event) SetJhat(x JoyHatEvent) {
    p := (*C.SDL_JoyHatEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.which = C.SDL_JoystickID(x.Which)
    p.hat = C.Uint8(x.Hat)
    p.value = C.Uint8(x.Value)
    p.padding1 = C.Uint8(x.Padding1)
    p.padding2 = C.Uint8(x.Padding2)
}

 // Joystick button event data
func (u *Event) Jbutton() JoyButtonEvent {
    p := (*C.SDL_JoyButtonEvent)(unsafe.Pointer(u))
    return JoyButtonEvent{uint32(p._type), uint32(p.timestamp), JoystickID(p.which), uint8(p.button), uint8(p.state), uint8(p.padding1), uint8(p.padding2)}
}
 // Joystick button event data
func (u *Event) SetJbutton(x JoyButtonEvent) {
    p := (*C.SDL_JoyButtonEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.which = C.SDL_JoystickID(x.Which)
    p.button = C.Uint8(x.Button)
    p.state = C.Uint8(x.State)
    p.padding1 = C.Uint8(x.Padding1)
    p.padding2 = C.Uint8(x.Padding2)
}

 // Joystick device change event data
func (u *Event) Jdevice() JoyDeviceEvent {
    p := (*C.SDL_JoyDeviceEvent)(unsafe.Pointer(u))
    return JoyDeviceEvent{uint32(p._type), uint32(p.timestamp), int32(p.which)}
}
 // Joystick device change event data
func (u *Event) SetJdevice(x JoyDeviceEvent) {
    p := (*C.SDL_JoyDeviceEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.which = C.Sint32(x.Which)
}

 // Game Controller axis event data
func (u *Event) Caxis() ControllerAxisEvent {
    p := (*C.SDL_ControllerAxisEvent)(unsafe.Pointer(u))
    return ControllerAxisEvent{uint32(p._type), uint32(p.timestamp), JoystickID(p.which), uint8(p.axis), uint8(p.padding1), uint8(p.padding2), uint8(p.padding3), int16(p.value), uint16(p.padding4)}
}
 // Game Controller axis event data
func (u *Event) SetCaxis(x ControllerAxisEvent) {
    p := (*C.SDL_ControllerAxisEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.which = C.SDL_JoystickID(x.Which)
    p.axis = C.Uint8(x.Axis)
    p.padding1 = C.Uint8(x.Padding1)
    p.padding2 = C.Uint8(x.Padding2)
    p.padding3 = C.Uint8(x.Padding3)
    p.value = C.Sint16(x.Value)
    p.padding4 = C.Uint16(x.Padding4)
}

 // Game Controller button event data
func (u *Event) Cbutton() ControllerButtonEvent {
    p := (*C.SDL_ControllerButtonEvent)(unsafe.Pointer(u))
    return ControllerButtonEvent{uint32(p._type), uint32(p.timestamp), JoystickID(p.which), uint8(p.button), uint8(p.state), uint8(p.padding1), uint8(p.padding2)}
}
 // Game Controller button event data
func (u *Event) SetCbutton(x ControllerButtonEvent) {
    p := (*C.SDL_ControllerButtonEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.which = C.SDL_JoystickID(x.Which)
    p.button = C.Uint8(x.Button)
    p.state = C.Uint8(x.State)
    p.padding1 = C.Uint8(x.Padding1)
    p.padding2 = C.Uint8(x.Padding2)
}

 // Game Controller device event data
func (u *Event) Cdevice() ControllerDeviceEvent {
    p := (*C.SDL_ControllerDeviceEvent)(unsafe.Pointer(u))
    return ControllerDeviceEvent{uint32(p._type), uint32(p.timestamp), int32(p.which)}
}
 // Game Controller device event data
func (u *Event) SetCdevice(x ControllerDeviceEvent) {
    p := (*C.SDL_ControllerDeviceEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.which = C.Sint32(x.Which)
}

 // Audio device event data
func (u *Event) Adevice() AudioDeviceEvent {
    p := (*C.SDL_AudioDeviceEvent)(unsafe.Pointer(u))
    return AudioDeviceEvent{uint32(p._type), uint32(p.timestamp), uint32(p.which), uint8(p.iscapture), uint8(p.padding1), uint8(p.padding2), uint8(p.padding3)}
}
 // Audio device event data
func (u *Event) SetAdevice(x AudioDeviceEvent) {
    p := (*C.SDL_AudioDeviceEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.which = C.Uint32(x.Which)
    p.iscapture = C.Uint8(x.Iscapture)
    p.padding1 = C.Uint8(x.Padding1)
    p.padding2 = C.Uint8(x.Padding2)
    p.padding3 = C.Uint8(x.Padding3)
}

 // Quit request event data
func (u *Event) Quit() QuitEvent {
    p := (*C.SDL_QuitEvent)(unsafe.Pointer(u))
    return QuitEvent{uint32(p._type), uint32(p.timestamp)}
}
 // Quit request event data
func (u *Event) SetQuit(x QuitEvent) {
    p := (*C.SDL_QuitEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
}

 // Custom event data
func (u *Event) User() UserEvent {
    p := (*C.SDL_UserEvent)(unsafe.Pointer(u))
    return UserEvent{uint32(p._type), uint32(p.timestamp), uint32(p.windowID), int32(p.code), uintptr(p.data1), uintptr(p.data2)}
}
 // Custom event data
func (u *Event) SetUser(x UserEvent) {
    p := (*C.SDL_UserEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.windowID = C.Uint32(x.WindowID)
    p.code = C.Sint32(x.Code)
    p.data1 = unsafe.Pointer(x.Data1)
    p.data2 = unsafe.Pointer(x.Data2)
}

 // Touch finger event data
func (u *Event) Tfinger() TouchFingerEvent {
    p := (*C.SDL_TouchFingerEvent)(unsafe.Pointer(u))
    return TouchFingerEvent{uint32(p._type), uint32(p.timestamp), TouchID(p.touchId), FingerID(p.fingerId), float32(p.x), float32(p.y), float32(p.dx), float32(p.dy), float32(p.pressure)}
}
 // Touch finger event data
func (u *Event) SetTfinger(x TouchFingerEvent) {
    p := (*C.SDL_TouchFingerEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.touchId = C.SDL_TouchID(x.TouchId)
    p.fingerId = C.SDL_FingerID(x.FingerId)
    p.x = C.float(x.X)
    p.y = C.float(x.Y)
    p.dx = C.float(x.Dx)
    p.dy = C.float(x.Dy)
    p.pressure = C.float(x.Pressure)
}

 // Gesture event data
func (u *Event) Mgesture() MultiGestureEvent {
    p := (*C.SDL_MultiGestureEvent)(unsafe.Pointer(u))
    return MultiGestureEvent{uint32(p._type), uint32(p.timestamp), TouchID(p.touchId), float32(p.dTheta), float32(p.dDist), float32(p.x), float32(p.y), uint16(p.numFingers), uint16(p.padding)}
}
 // Gesture event data
func (u *Event) SetMgesture(x MultiGestureEvent) {
    p := (*C.SDL_MultiGestureEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.touchId = C.SDL_TouchID(x.TouchId)
    p.dTheta = C.float(x.DTheta)
    p.dDist = C.float(x.DDist)
    p.x = C.float(x.X)
    p.y = C.float(x.Y)
    p.numFingers = C.Uint16(x.NumFingers)
    p.padding = C.Uint16(x.Padding)
}

 // Gesture event data
func (u *Event) Dgesture() DollarGestureEvent {
    p := (*C.SDL_DollarGestureEvent)(unsafe.Pointer(u))
    return DollarGestureEvent{uint32(p._type), uint32(p.timestamp), TouchID(p.touchId), GestureID(p.gestureId), uint32(p.numFingers), float32(p.error), float32(p.x), float32(p.y)}
}
 // Gesture event data
func (u *Event) SetDgesture(x DollarGestureEvent) {
    p := (*C.SDL_DollarGestureEvent)(unsafe.Pointer(u))
    p._type = C.Uint32(x.Type)
    p.timestamp = C.Uint32(x.Timestamp)
    p.touchId = C.SDL_TouchID(x.TouchId)
    p.gestureId = C.SDL_GestureID(x.GestureId)
    p.numFingers = C.Uint32(x.NumFingers)
    p.error = C.float(x.Error)
    p.x = C.float(x.X)
    p.y = C.float(x.Y)
}

 // Drag and drop event data
func (u *Event) Drop() DropEvent {
    p := (*C.SDL_DropEvent)(unsafe.Pointer(u))
    return DropEvent{uint32(p._type), uint32(p.timestamp), C.GoString(p.file)}
}

func (u *Event) Padding() [56]uint8 {
    p := (*[56]C.Uint8)(unsafe.Pointer(u))
    return *(*[56]uint8)(unsafe.Pointer(&(*p)))
}
func (u *Event) SetPadding(x [56]uint8) {
    p := (*[56]C.Uint8)(unsafe.Pointer(u))
    *p = *(*[56]C.Uint8)(unsafe.Pointer(&(x)))
}

type Eventaction int
const (
    ADDEVENT Eventaction = C.SDL_ADDEVENT

    PEEKEVENT Eventaction = C.SDL_PEEKEVENT

    GETEVENT Eventaction = C.SDL_GETEVENT
)


const (
    QUERY = C.SDL_QUERY

    IGNORE = C.SDL_IGNORE

    DISABLE = C.SDL_DISABLE

    ENABLE = C.SDL_ENABLE
)

 // This function allows you to set the state of processing certain
 // events.
 //   
 //   - If state is set to SDL_IGNORE, that event will be automatically
 //     dropped from the event queue and will not event be filtered.
 //   - If state is set to SDL_ENABLE, that event will be processed normally.
 //   - If state is set to SDL_QUERY, SDL_EventState() will return the current
 //     processing state of the specified event.
func EventState(_type uint32, state int) (retval uint8) {
    retval = uint8(C.SDL_EventState(C.Uint32(_type), C.int(state)))
    return
}

const (
    RELEASED = C.SDL_RELEASED

    PRESSED = C.SDL_PRESSED

    TEXTEDITINGEVENT_TEXT_SIZE = C.SDL_TEXTEDITINGEVENT_TEXT_SIZE

    TEXTINPUTEVENT_TEXT_SIZE = C.SDL_TEXTINPUTEVENT_TEXT_SIZE
)

 // The types of events that can be delivered.
type EventType int
const (
     // Unused (do not remove)
    FIRSTEVENT EventType = C.SDL_FIRSTEVENT

     // User-requested quit
    QUIT EventType = C.SDL_QUIT

     // The application is being terminated by the OS Called on iOS in
     // applicationWillTerminate() Called on Android in onDestroy()
    APP_TERMINATING EventType = C.SDL_APP_TERMINATING

     // The application is low on memory, free memory if possible. Called on
     // iOS in applicationDidReceiveMemoryWarning() Called on Android in
     // onLowMemory()
    APP_LOWMEMORY EventType = C.SDL_APP_LOWMEMORY

     // The application is about to enter the background Called on iOS in
     // applicationWillResignActive() Called on Android in onPause()
    APP_WILLENTERBACKGROUND EventType = C.SDL_APP_WILLENTERBACKGROUND

     // The application did enter the background and may not get CPU for some
     // time Called on iOS in applicationDidEnterBackground() Called on
     // Android in onPause()
    APP_DIDENTERBACKGROUND EventType = C.SDL_APP_DIDENTERBACKGROUND

     // The application is about to enter the foreground Called on iOS in
     // applicationWillEnterForeground() Called on Android in onResume()
    APP_WILLENTERFOREGROUND EventType = C.SDL_APP_WILLENTERFOREGROUND

     // The application is now interactive Called on iOS in
     // applicationDidBecomeActive() Called on Android in onResume()
    APP_DIDENTERFOREGROUND EventType = C.SDL_APP_DIDENTERFOREGROUND

     // Window state change
    WINDOWEVENT EventType = C.SDL_WINDOWEVENT

     // System specific event
    SYSWMEVENT EventType = C.SDL_SYSWMEVENT

     // Key pressed
    KEYDOWN EventType = C.SDL_KEYDOWN

     // Key released
    KEYUP EventType = C.SDL_KEYUP

     // Keyboard text editing (composition)
    TEXTEDITING EventType = C.SDL_TEXTEDITING

     // Keyboard text input
    TEXTINPUT EventType = C.SDL_TEXTINPUT

     // Keymap changed due to a system event such as an input language or
     // keyboard layout change.
    KEYMAPCHANGED EventType = C.SDL_KEYMAPCHANGED

     // Mouse moved
    MOUSEMOTION EventType = C.SDL_MOUSEMOTION

     // Mouse button pressed
    MOUSEBUTTONDOWN EventType = C.SDL_MOUSEBUTTONDOWN

     // Mouse button released
    MOUSEBUTTONUP EventType = C.SDL_MOUSEBUTTONUP

     // Mouse wheel motion
    MOUSEWHEEL EventType = C.SDL_MOUSEWHEEL

     // Joystick axis motion
    JOYAXISMOTION EventType = C.SDL_JOYAXISMOTION

     // Joystick trackball motion
    JOYBALLMOTION EventType = C.SDL_JOYBALLMOTION

     // Joystick hat position change
    JOYHATMOTION EventType = C.SDL_JOYHATMOTION

     // Joystick button pressed
    JOYBUTTONDOWN EventType = C.SDL_JOYBUTTONDOWN

     // Joystick button released
    JOYBUTTONUP EventType = C.SDL_JOYBUTTONUP

     // A new joystick has been inserted into the system
    JOYDEVICEADDED EventType = C.SDL_JOYDEVICEADDED

     // An opened joystick has been removed
    JOYDEVICEREMOVED EventType = C.SDL_JOYDEVICEREMOVED

     // Game controller axis motion
    CONTROLLERAXISMOTION EventType = C.SDL_CONTROLLERAXISMOTION

     // Game controller button pressed
    CONTROLLERBUTTONDOWN EventType = C.SDL_CONTROLLERBUTTONDOWN

     // Game controller button released
    CONTROLLERBUTTONUP EventType = C.SDL_CONTROLLERBUTTONUP

     // A new Game controller has been inserted into the system
    CONTROLLERDEVICEADDED EventType = C.SDL_CONTROLLERDEVICEADDED

     // An opened Game controller has been removed
    CONTROLLERDEVICEREMOVED EventType = C.SDL_CONTROLLERDEVICEREMOVED

     // The controller mapping was updated
    CONTROLLERDEVICEREMAPPED EventType = C.SDL_CONTROLLERDEVICEREMAPPED

    FINGERDOWN EventType = C.SDL_FINGERDOWN

    FINGERUP EventType = C.SDL_FINGERUP

    FINGERMOTION EventType = C.SDL_FINGERMOTION

    DOLLARGESTURE EventType = C.SDL_DOLLARGESTURE

    DOLLARRECORD EventType = C.SDL_DOLLARRECORD

    MULTIGESTURE EventType = C.SDL_MULTIGESTURE

     // The clipboard changed
    CLIPBOARDUPDATE EventType = C.SDL_CLIPBOARDUPDATE

     // The system requests a file open
    DROPFILE EventType = C.SDL_DROPFILE

     // A new audio device is available
    AUDIODEVICEADDED EventType = C.SDL_AUDIODEVICEADDED

     // An audio device has been removed.
    AUDIODEVICEREMOVED EventType = C.SDL_AUDIODEVICEREMOVED

     // The render targets have been reset and their contents need to be
     // updated
    RENDER_TARGETS_RESET EventType = C.SDL_RENDER_TARGETS_RESET

     // The device has been reset and all textures need to be recreated
    RENDER_DEVICE_RESET EventType = C.SDL_RENDER_DEVICE_RESET

     // Events SDL_USEREVENT through SDL_LASTEVENT are for your use, and
     // should be allocated with SDL_RegisterEvents()
    USEREVENT EventType = C.SDL_USEREVENT

     // This last event is only for bounding internal arrays
    LASTEVENT EventType = C.SDL_LASTEVENT
)

type EventFilter C.SDL_EventFilter


 // Pumps the event loop, gathering events from the input devices.
 // 
 // This function updates the event queue and internal input device state.
 // 
 // This should only be run in the thread that sets the video mode.
func PumpEvents() {
    C.SDL_PumpEvents()
}

 // Checks to see if certain event types are in the event queue.
func HasEvent(_type uint32) (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasEvent(C.Uint32(_type)))
    return
}

func HasEvents(minType uint32, maxType uint32) (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasEvents(C.Uint32(minType), C.Uint32(maxType)))
    return
}

 // This function clears events from the event queue This function only
 // affects currently queued events. If you want to make sure that all
 // pending OS events are flushed, you can call SDL_PumpEvents() on the
 // main thread immediately before the flush call.
func FlushEvent(_type uint32) {
    C.SDL_FlushEvent(C.Uint32(_type))
}

func FlushEvents(minType uint32, maxType uint32) {
    C.SDL_FlushEvents(C.Uint32(minType), C.Uint32(maxType))
}

 // Polls for currently pending events.
 // 
 // Returns: 1 if there are any pending events, or 0 if there are none
 // available.
 // 
 //   event
 //     If not NULL, the next event is removed from the queue and stored in
 //     that area.
 //   
func PollEvent() (retval int, event *Event) {
    tmp_event := new(C.SDL_Event)
    retval = int(C.SDL_PollEvent((*C.SDL_Event)(tmp_event)))
    event = (*Event)(unsafe.Pointer(tmp_event))
    return
}

 // Waits indefinitely for the next available event.
 // 
 // Returns: 1, or 0 if there was an error while waiting for events.
 // 
 //   event
 //     If not NULL, the next event is removed from the queue and stored in
 //     that area.
 //   
func WaitEvent() (retval int, event *Event) {
    tmp_event := new(C.SDL_Event)
    retval = int(C.SDL_WaitEvent((*C.SDL_Event)(tmp_event)))
    event = (*Event)(unsafe.Pointer(tmp_event))
    return
}

 // Waits until the specified timeout (in milliseconds) for the next
 // available event.
 // 
 // Returns: 1, or 0 if there was an error while waiting for events.
 // 
 //   event
 //     If not NULL, the next event is removed from the queue and stored in
 //     that area.
 //   
 //   timeout
 //     The timeout (in milliseconds) to wait for next event.
 //   
func WaitEventTimeout(timeout int) (retval int, event *Event) {
    tmp_event := new(C.SDL_Event)
    retval = int(C.SDL_WaitEventTimeout((*C.SDL_Event)(tmp_event), C.int(timeout)))
    event = (*Event)(unsafe.Pointer(tmp_event))
    return
}

 // Add an event to the event queue.
 // 
 // Returns: 1 on success, 0 if the event was filtered, or -1 if the event
 // queue was full or there was some other error.
 // 
func (event *Event) Push() (retval int) {
    retval = int(C.SDL_PushEvent((*C.SDL_Event)(event)))
    return
}






 // This function allocates a set of user-defined events, and returns the
 // beginning event number for that set of events.
 // 
 // If there aren't enough user-defined events left, this function returns
 // (Uint32)-1
func RegisterEvents(numevents int) (retval uint32) {
    retval = uint32(C.SDL_RegisterEvents(C.int(numevents)))
    return
}
