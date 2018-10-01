// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"
import "unsafe"

 // The SDL haptic subsystem allows you to control haptic (force feedback)
 // devices.
 // 
 // The basic usage is as follows:
 //   
 //   - Initialize the subsystem (SDL_INIT_HAPTIC).
 //   - Open a haptic device. SDL_HapticOpen() to open from
 //     index.SDL_HapticOpenFromJoystick() to open from an existing joystick.
 //   - Create an effect (SDL_HapticEffect).
 //   - Upload the effect with SDL_HapticNewEffect().
 //   - Run the effect with SDL_HapticRunEffect().
 //   - (optional) Free the effect with SDL_HapticDestroyEffect().
 //   - Close the haptic device with SDL_HapticClose().
 // 
 // Simple rumble example:
 //   SDL_Haptic *haptic;
 //   
 //   // Open the device
 //   haptic = SDL_HapticOpen( 0 );
 //   if (haptic == NULL)
 //      return -1;
 //   
 //   // Initialize simple rumble
 //   if (SDL_HapticRumbleInit( haptic ) != 0)
 //      return -1;
 //   
 //   // Play effect at 50% strength for 2 seconds
 //   if (SDL_HapticRumblePlay( haptic, 0.5, 2000 ) != 0)
 //      return -1;
 //   SDL_Delay( 2000 );
 //   
 //   // Clean up
 //   SDL_HapticClose( haptic );
 // 
 // Complete example:
 //   int test_haptic( SDL_Joystick * joystick ) {
 //      SDL_Haptic *haptic;
 //      SDL_HapticEffect effect;
 //      int effect_id;
 //   
 //      // Open the device
 //      haptic = SDL_HapticOpenFromJoystick( joystick );
 //      if (haptic == NULL) return -1; // Most likely joystick isn't haptic
 //   
 //      // See if it can do sine waves
 //      if ((SDL_HapticQuery(haptic) & SDL_HAPTIC_SINE)==0) {
 //         SDL_HapticClose(haptic); // No sine effect
 //         return -1;
 //      }
 //   
 //      // Create the effect
 //      memset( &effect, 0, sizeof(SDL_HapticEffect) ); // 0 is safe default
 //      effect.type = SDL_HAPTIC_SINE;
 //      effect.periodic.direction.type = SDL_HAPTIC_POLAR; // Polar coordinates
 //      effect.periodic.direction.dir[0] = 18000; // Force comes from south
 //      effect.periodic.period = 1000; // 1000 ms
 //      effect.periodic.magnitude = 20000; // 20000/32767 strength
 //      effect.periodic.length = 5000; // 5 seconds long
 //      effect.periodic.attack_length = 1000; // Takes 1 second to get max strength
 //      effect.periodic.fade_length = 1000; // Takes 1 second to fade away
 //   
 //      // Upload the effect
 //      effect_id = SDL_HapticNewEffect( haptic, &effect );
 //   
 //      // Test the effect
 //      SDL_HapticRunEffect( haptic, effect_id, 1 );
 //      SDL_Delay( 5000); // Wait for the effect to finish
 //   
 //      // We destroy the effect, although closing the device also does this
 //      SDL_HapticDestroyEffect( haptic, effect_id );
 //   
 //      // Close the device
 //      SDL_HapticClose(haptic);
 //   
 //      return 0; // Success
 //   }
 // 

 // Structure that represents a haptic direction.
 // 
 // This is the direction where the force comes from, instead of the
 // direction in which the force is exerted.
 // 
 // Directions can be specified by:
 //   
 //   - SDL_HAPTIC_POLAR : Specified by polar coordinates.
 //   - SDL_HAPTIC_CARTESIAN : Specified by cartesian coordinates.
 //   - SDL_HAPTIC_SPHERICAL : Specified by spherical coordinates.
 // 
 // Cardinal directions of the haptic device are relative to the
 // positioning of the device. North is considered to be away from the
 // user.
 // 
 // The following diagram represents the cardinal directions:
 //                 .--.
 //                 |__| .-------.
 //                 |=.| |.-----.|
 //                 |--| ||     ||
 //                 |  | |'-----'|
 //                 |__|~')_____('
 //                   [ COMPUTER ]
 //   
 //   
 //                     North (0,-1)
 //                         ^
 //                         |
 //                         |
 //   (-1,0)  West <----[ HAPTIC ]----> East (1,0)
 //                         |
 //                         |
 //                         v
 //                      South (0,1)
 //   
 //   
 //                      [ USER ]
 //                        \|||/
 //                        (o o)
 //                  ---ooO-(_)-Ooo---
 // 
 // If type is SDL_HAPTIC_POLAR, direction is encoded by hundredths of a
 // degree starting north and turning clockwise. SDL_HAPTIC_POLAR only
 // uses the first dir parameter. The cardinal directions would be:
 //   
 //   - North: 0 (0 degrees)
 //   - East: 9000 (90 degrees)
 //   - South: 18000 (180 degrees)
 //   - West: 27000 (270 degrees)
 // 
 // If type is SDL_HAPTIC_CARTESIAN, direction is encoded by three
 // positions (X axis, Y axis and Z axis (with 3 axes)).
 // SDL_HAPTIC_CARTESIAN uses the first three dir parameters. The cardinal
 // directions would be:
 //   
 //   - North: 0,-1, 0
 //   - East: 1, 0, 0
 //   - South: 0, 1, 0
 //   - West: -1, 0, 0
 // 
 // The Z axis represents the height of the effect if supported, otherwise
 // it's unused. In cartesian encoding (1, 2) would be the same as (2, 4),
 // you can use any multiple you want, only the direction matters.
 // 
 // If type is SDL_HAPTIC_SPHERICAL, direction is encoded by two
 // rotations. The first two dir parameters are used. The dir parameters
 // are as follows (all values are in hundredths of degrees):
 //   
 //   - Degrees from (1, 0) rotated towards (0, 1).
 //   - Degrees towards (0, 0, 1) (device needs at least 3 axes).
 // 
 // Example of force coming from the south with all encodings (force
 // coming from the south means the user will have to pull the stick to
 // counteract):
 //   SDL_HapticDirection direction;
 //   
 //   // Cartesian directions
 //   direction.type = SDL_HAPTIC_CARTESIAN; // Using cartesian direction encoding.
 //   direction.dir[0] = 0; // X position
 //   direction.dir[1] = 1; // Y position
 //   // Assuming the device has 2 axes, we don't need to specify third parameter.
 //   
 //   // Polar directions
 //   direction.type = SDL_HAPTIC_POLAR; // We'll be using polar direction encoding.
 //   direction.dir[0] = 18000; // Polar only uses first parameter
 //   
 //   // Spherical coordinates
 //   direction.type = SDL_HAPTIC_SPHERICAL; // Spherical encoding
 //   direction.dir[0] = 9000; // Since we only have two axes we don't need more parameters.
 // 
 // See also: SDL_HAPTIC_POLAR
 // 
 // See also: SDL_HAPTIC_CARTESIAN
 // 
 // See also: SDL_HAPTIC_SPHERICAL
 // 
 // See also: SDL_HapticEffect
 // 
 // See also: SDL_HapticNumAxes
 // 
 // ↪ https://wiki.libsdl.org/SDL_HapticDirection
type HapticDirection struct {
     // The type of encoding.
    Type uint8

     // The encoded direction.
    Dir [3]int32
}

func fromC2HapticDirection(s C.SDL_HapticDirection) HapticDirection {
    return HapticDirection{uint8(s._type), *(*[3]int32)(unsafe.Pointer(&(s.dir)))}
}

func toCFromHapticDirection(s HapticDirection) (d C.SDL_HapticDirection) {
    d._type = C.Uint8(s.Type)
    d.dir = *(*[3]C.Sint32)(unsafe.Pointer(&(s.Dir)))
    return
}

 // A structure containing a template for a Constant effect.
 // 
 // This struct is exclusively for the SDL_HAPTIC_CONSTANT effect.
 // 
 // A constant effect applies a constant force in the specified direction
 // to the joystick.
 // 
 // See also: SDL_HAPTIC_CONSTANT
 // 
 // See also: SDL_HapticEffect
 // 
 // ↪ https://wiki.libsdl.org/SDL_HapticConstant
type HapticConstant struct {
     // SDL_HAPTIC_CONSTANT
    Type uint16

     // Direction of the effect.
    Direction HapticDirection

     // Duration of the effect.
    Length uint32

     // Delay before starting the effect.
    Delay uint16

     // Button that triggers the effect.
    Button uint16

     // How soon it can be triggered again after button.
    Interval uint16

     // Strength of the constant effect.
    Level int16

     // Duration of the attack.
    Attack_length uint16

     // Level at the start of the attack.
    Attack_level uint16

     // Duration of the fade.
    Fade_length uint16

     // Level at the end of the fade.
    Fade_level uint16
}

func fromC2HapticConstant(s C.SDL_HapticConstant) HapticConstant {
    return HapticConstant{uint16(s._type), HapticDirection{uint8(s.direction._type), *(*[3]int32)(unsafe.Pointer(&(s.direction.dir)))}, uint32(s.length), uint16(s.delay), uint16(s.button), uint16(s.interval), int16(s.level), uint16(s.attack_length), uint16(s.attack_level), uint16(s.fade_length), uint16(s.fade_level)}
}

func toCFromHapticConstant(s HapticConstant) (d C.SDL_HapticConstant) {
    d._type = C.Uint16(s.Type)
    d.direction._type = C.Uint8(s.Direction.Type)
    d.direction.dir = *(*[3]C.Sint32)(unsafe.Pointer(&(s.Direction.Dir)))
    d.length = C.Uint32(s.Length)
    d.delay = C.Uint16(s.Delay)
    d.button = C.Uint16(s.Button)
    d.interval = C.Uint16(s.Interval)
    d.level = C.Sint16(s.Level)
    d.attack_length = C.Uint16(s.Attack_length)
    d.attack_level = C.Uint16(s.Attack_level)
    d.fade_length = C.Uint16(s.Fade_length)
    d.fade_level = C.Uint16(s.Fade_level)
    return
}

 // A structure containing a template for a Periodic effect.
 // 
 // The struct handles the following effects:
 //   
 //   - SDL_HAPTIC_SINE
 //   - SDL_HAPTIC_LEFTRIGHT
 //   - SDL_HAPTIC_TRIANGLE
 //   - SDL_HAPTIC_SAWTOOTHUP
 //   - SDL_HAPTIC_SAWTOOTHDOWN
 // 
 // A periodic effect consists in a wave-shaped effect that repeats itself
 // over time. The type determines the shape of the wave and the
 // parameters determine the dimensions of the wave.
 // 
 // Phase is given by hundredth of a degree meaning that giving the phase
 // a value of 9000 will displace it 25% of its period. Here are sample
 // values:
 //   
 //   - 0: No phase displacement.
 //   - 9000: Displaced 25% of its period.
 //   - 18000: Displaced 50% of its period.
 //   - 27000: Displaced 75% of its period.
 //   - 36000: Displaced 100% of its period, same as 0, but 0 is preferred.
 // 
 // Examples:
 //   SDL_HAPTIC_SINE
 //     __      __      __      __
 //    /  \    /  \    /  \    /
 //   /    \__/    \__/    \__/
 //   
 //   SDL_HAPTIC_SQUARE
 //    __    __    __    __    __
 //   |  |  |  |  |  |  |  |  |  |
 //   |  |__|  |__|  |__|  |__|  |
 //   
 //   SDL_HAPTIC_TRIANGLE
 //     /\    /\    /\    /\    /\
 //    /  \  /  \  /  \  /  \  /
 //   /    \/    \/    \/    \/
 //   
 //   SDL_HAPTIC_SAWTOOTHUP
 //     /|  /|  /|  /|  /|  /|  /|
 //    / | / | / | / | / | / | / |
 //   /  |/  |/  |/  |/  |/  |/  |
 //   
 //   SDL_HAPTIC_SAWTOOTHDOWN
 //   \  |\  |\  |\  |\  |\  |\  |
 //    \ | \ | \ | \ | \ | \ | \ |
 //     \|  \|  \|  \|  \|  \|  \|
 // 
 // See also: SDL_HAPTIC_SINE
 // 
 // See also: SDL_HAPTIC_LEFTRIGHT
 // 
 // See also: SDL_HAPTIC_TRIANGLE
 // 
 // See also: SDL_HAPTIC_SAWTOOTHUP
 // 
 // See also: SDL_HAPTIC_SAWTOOTHDOWN
 // 
 // See also: SDL_HapticEffect
 // 
 // ↪ https://wiki.libsdl.org/SDL_HapticPeriodic
type HapticPeriodic struct {
     // SDL_HAPTIC_SINE, SDL_HAPTIC_LEFTRIGHT, SDL_HAPTIC_TRIANGLE,
     // SDL_HAPTIC_SAWTOOTHUP or SDL_HAPTIC_SAWTOOTHDOWN
    Type uint16

     // Direction of the effect.
    Direction HapticDirection

     // Duration of the effect.
    Length uint32

     // Delay before starting the effect.
    Delay uint16

     // Button that triggers the effect.
    Button uint16

     // How soon it can be triggered again after button.
    Interval uint16

     // Period of the wave.
    Period uint16

     // Peak value; if negative, equivalent to 180 degrees extra phase shift.
    Magnitude int16

     // Mean value of the wave.
    Offset int16

     // Positive phase shift given by hundredth of a degree.
    Phase uint16

     // Duration of the attack.
    Attack_length uint16

     // Level at the start of the attack.
    Attack_level uint16

     // Duration of the fade.
    Fade_length uint16

     // Level at the end of the fade.
    Fade_level uint16
}

func fromC2HapticPeriodic(s C.SDL_HapticPeriodic) HapticPeriodic {
    return HapticPeriodic{uint16(s._type), HapticDirection{uint8(s.direction._type), *(*[3]int32)(unsafe.Pointer(&(s.direction.dir)))}, uint32(s.length), uint16(s.delay), uint16(s.button), uint16(s.interval), uint16(s.period), int16(s.magnitude), int16(s.offset), uint16(s.phase), uint16(s.attack_length), uint16(s.attack_level), uint16(s.fade_length), uint16(s.fade_level)}
}

func toCFromHapticPeriodic(s HapticPeriodic) (d C.SDL_HapticPeriodic) {
    d._type = C.Uint16(s.Type)
    d.direction._type = C.Uint8(s.Direction.Type)
    d.direction.dir = *(*[3]C.Sint32)(unsafe.Pointer(&(s.Direction.Dir)))
    d.length = C.Uint32(s.Length)
    d.delay = C.Uint16(s.Delay)
    d.button = C.Uint16(s.Button)
    d.interval = C.Uint16(s.Interval)
    d.period = C.Uint16(s.Period)
    d.magnitude = C.Sint16(s.Magnitude)
    d.offset = C.Sint16(s.Offset)
    d.phase = C.Uint16(s.Phase)
    d.attack_length = C.Uint16(s.Attack_length)
    d.attack_level = C.Uint16(s.Attack_level)
    d.fade_length = C.Uint16(s.Fade_length)
    d.fade_level = C.Uint16(s.Fade_level)
    return
}

 // A structure containing a template for a Condition effect.
 // 
 // The struct handles the following effects:
 //   
 //   - SDL_HAPTIC_SPRING: Effect based on axes position.
 //   - SDL_HAPTIC_DAMPER: Effect based on axes velocity.
 //   - SDL_HAPTIC_INERTIA: Effect based on axes acceleration.
 //   - SDL_HAPTIC_FRICTION: Effect based on axes movement.
 // 
 // Direction is handled by condition internals instead of a direction
 // member. The condition effect specific members have three parameters.
 // The first refers to the X axis, the second refers to the Y axis and
 // the third refers to the Z axis. The right terms refer to the positive
 // side of the axis and the left terms refer to the negative side of the
 // axis. Please refer to the SDL_HapticDirection diagram for which side
 // is positive and which is negative.
 // 
 // See also: SDL_HapticDirection
 // 
 // See also: SDL_HAPTIC_SPRING
 // 
 // See also: SDL_HAPTIC_DAMPER
 // 
 // See also: SDL_HAPTIC_INERTIA
 // 
 // See also: SDL_HAPTIC_FRICTION
 // 
 // See also: SDL_HapticEffect
 // 
 // ↪ https://wiki.libsdl.org/SDL_HapticCondition
type HapticCondition struct {
     // SDL_HAPTIC_SPRING, SDL_HAPTIC_DAMPER, SDL_HAPTIC_INERTIA or
     // SDL_HAPTIC_FRICTION
    Type uint16

     // Direction of the effect - Not used ATM.
    Direction HapticDirection

     // Duration of the effect.
    Length uint32

     // Delay before starting the effect.
    Delay uint16

     // Button that triggers the effect.
    Button uint16

     // How soon it can be triggered again after button.
    Interval uint16

     // Level when joystick is to the positive side; max 0xFFFF.
    Right_sat [3]uint16

     // Level when joystick is to the negative side; max 0xFFFF.
    Left_sat [3]uint16

     // How fast to increase the force towards the positive side.
    Right_coeff [3]int16

     // How fast to increase the force towards the negative side.
    Left_coeff [3]int16

     // Size of the dead zone; max 0xFFFF: whole axis-range when 0-centered.
    Deadband [3]uint16

     // Position of the dead zone.
    Center [3]int16
}

func fromC2HapticCondition(s C.SDL_HapticCondition) HapticCondition {
    return HapticCondition{uint16(s._type), HapticDirection{uint8(s.direction._type), *(*[3]int32)(unsafe.Pointer(&(s.direction.dir)))}, uint32(s.length), uint16(s.delay), uint16(s.button), uint16(s.interval), *(*[3]uint16)(unsafe.Pointer(&(s.right_sat))), *(*[3]uint16)(unsafe.Pointer(&(s.left_sat))), *(*[3]int16)(unsafe.Pointer(&(s.right_coeff))), *(*[3]int16)(unsafe.Pointer(&(s.left_coeff))), *(*[3]uint16)(unsafe.Pointer(&(s.deadband))), *(*[3]int16)(unsafe.Pointer(&(s.center)))}
}

func toCFromHapticCondition(s HapticCondition) (d C.SDL_HapticCondition) {
    d._type = C.Uint16(s.Type)
    d.direction._type = C.Uint8(s.Direction.Type)
    d.direction.dir = *(*[3]C.Sint32)(unsafe.Pointer(&(s.Direction.Dir)))
    d.length = C.Uint32(s.Length)
    d.delay = C.Uint16(s.Delay)
    d.button = C.Uint16(s.Button)
    d.interval = C.Uint16(s.Interval)
    d.right_sat = *(*[3]C.Uint16)(unsafe.Pointer(&(s.Right_sat)))
    d.left_sat = *(*[3]C.Uint16)(unsafe.Pointer(&(s.Left_sat)))
    d.right_coeff = *(*[3]C.Sint16)(unsafe.Pointer(&(s.Right_coeff)))
    d.left_coeff = *(*[3]C.Sint16)(unsafe.Pointer(&(s.Left_coeff)))
    d.deadband = *(*[3]C.Uint16)(unsafe.Pointer(&(s.Deadband)))
    d.center = *(*[3]C.Sint16)(unsafe.Pointer(&(s.Center)))
    return
}

 // A structure containing a template for a Ramp effect.
 // 
 // This struct is exclusively for the SDL_HAPTIC_RAMP effect.
 // 
 // The ramp effect starts at start strength and ends at end strength. It
 // augments in linear fashion. If you use attack and fade with a ramp the
 // effects get added to the ramp effect making the effect become
 // quadratic instead of linear.
 // 
 // See also: SDL_HAPTIC_RAMP
 // 
 // See also: SDL_HapticEffect
 // 
 // ↪ https://wiki.libsdl.org/SDL_HapticRamp
type HapticRamp struct {
     // SDL_HAPTIC_RAMP
    Type uint16

     // Direction of the effect.
    Direction HapticDirection

     // Duration of the effect.
    Length uint32

     // Delay before starting the effect.
    Delay uint16

     // Button that triggers the effect.
    Button uint16

     // How soon it can be triggered again after button.
    Interval uint16

     // Beginning strength level.
    Start int16

     // Ending strength level.
    End int16

     // Duration of the attack.
    Attack_length uint16

     // Level at the start of the attack.
    Attack_level uint16

     // Duration of the fade.
    Fade_length uint16

     // Level at the end of the fade.
    Fade_level uint16
}

func fromC2HapticRamp(s C.SDL_HapticRamp) HapticRamp {
    return HapticRamp{uint16(s._type), HapticDirection{uint8(s.direction._type), *(*[3]int32)(unsafe.Pointer(&(s.direction.dir)))}, uint32(s.length), uint16(s.delay), uint16(s.button), uint16(s.interval), int16(s.start), int16(s.end), uint16(s.attack_length), uint16(s.attack_level), uint16(s.fade_length), uint16(s.fade_level)}
}

func toCFromHapticRamp(s HapticRamp) (d C.SDL_HapticRamp) {
    d._type = C.Uint16(s.Type)
    d.direction._type = C.Uint8(s.Direction.Type)
    d.direction.dir = *(*[3]C.Sint32)(unsafe.Pointer(&(s.Direction.Dir)))
    d.length = C.Uint32(s.Length)
    d.delay = C.Uint16(s.Delay)
    d.button = C.Uint16(s.Button)
    d.interval = C.Uint16(s.Interval)
    d.start = C.Sint16(s.Start)
    d.end = C.Sint16(s.End)
    d.attack_length = C.Uint16(s.Attack_length)
    d.attack_level = C.Uint16(s.Attack_level)
    d.fade_length = C.Uint16(s.Fade_length)
    d.fade_level = C.Uint16(s.Fade_level)
    return
}

 // A structure containing a template for a Left/Right effect.
 // 
 // This struct is exclusively for the SDL_HAPTIC_LEFTRIGHT effect.
 // 
 // The Left/Right effect is used to explicitly control the large and
 // small motors, commonly found in modern game controllers. The small
 // (right) motor is high frequency, and the large (left) motor is low
 // frequency.
 // 
 // See also: SDL_HAPTIC_LEFTRIGHT
 // 
 // See also: SDL_HapticEffect
 // 
 // ↪ https://wiki.libsdl.org/SDL_HapticLeftRight
type HapticLeftRight struct {
     // SDL_HAPTIC_LEFTRIGHT
    Type uint16

     // Duration of the effect in milliseconds.
    Length uint32

     // Control of the large controller motor.
    Large_magnitude uint16

     // Control of the small controller motor.
    Small_magnitude uint16
}

func fromC2HapticLeftRight(s C.SDL_HapticLeftRight) HapticLeftRight {
    return HapticLeftRight{uint16(s._type), uint32(s.length), uint16(s.large_magnitude), uint16(s.small_magnitude)}
}

func toCFromHapticLeftRight(s HapticLeftRight) (d C.SDL_HapticLeftRight) {
    d._type = C.Uint16(s.Type)
    d.length = C.Uint32(s.Length)
    d.large_magnitude = C.Uint16(s.Large_magnitude)
    d.small_magnitude = C.Uint16(s.Small_magnitude)
    return
}

 // The generic template for any haptic effect.
 // 
 // All values max at 32767 (0x7FFF). Signed values also can be negative.
 // Time values unless specified otherwise are in milliseconds.
 // 
 // You can also pass SDL_HAPTIC_INFINITY to length instead of a 0-32767
 // value. Neither delay, interval, attack_length nor fade_length support
 // SDL_HAPTIC_INFINITY. Fade will also not be used since effect never
 // ends.
 // 
 // Additionally, the SDL_HAPTIC_RAMP effect does not support a duration
 // of SDL_HAPTIC_INFINITY.
 // 
 // Button triggers may not be supported on all devices, it is advised to
 // not use them if possible. Buttons start at index 1 instead of index 0
 // like the joystick.
 // 
 // If both attack_length and fade_level are 0, the envelope is not used,
 // otherwise both values are used.
 // 
 // Common parts:
 //   // Replay - All effects have this
 //   Uint32 length;        // Duration of effect (ms).
 //   Uint16 delay;         // Delay before starting effect.
 //   
 //   // Trigger - All effects have this
 //   Uint16 button;        // Button that triggers effect.
 //   Uint16 interval;      // How soon before effect can be triggered again.
 //   
 //   // Envelope - All effects except condition effects have this
 //   Uint16 attack_length; // Duration of the attack (ms).
 //   Uint16 attack_level;  // Level at the start of the attack.
 //   Uint16 fade_length;   // Duration of the fade out (ms).
 //   Uint16 fade_level;    // Level at the end of the fade.
 // 
 // Here we have an example of a constant effect evolution in time:
 //   Strength
 //   ^
 //   |
 //   |    effect level -->  _________________
 //   |                     /                 \
 //   |                    /                   \
 //   |                   /                     \
 //   |                  /                       \
 //   | attack_level --> |                        \
 //   |                  |                        |  <---  fade_level
 //   |
 //   +--------------------------------------------------> Time
 //                      [--]                 [---]
 //                      attack_length        fade_length
 //   
 //   [------------------][-----------------------]
 //   delay               length
 // 
 // Note either the attack_level or the fade_level may be above the actual
 // effect level.
 // 
 // See also: SDL_HapticConstant
 // 
 // See also: SDL_HapticPeriodic
 // 
 // See also: SDL_HapticCondition
 // 
 // See also: SDL_HapticRamp
 // 
 // See also: SDL_HapticLeftRight
 // 
 // See also: SDL_HapticCustom
 // 
 // ↪ https://wiki.libsdl.org/SDL_HapticEffect
type HapticEffect C.SDL_HapticEffect

 // Effect type.
func (u *HapticEffect) Type() uint16 {
    p := (*C.Uint16)(unsafe.Pointer(u))
    return uint16(*p)
}
 // Effect type.
func (u *HapticEffect) SetType(x uint16) {
    p := (*C.Uint16)(unsafe.Pointer(u))
    *p = C.Uint16(x)
}

 // Constant effect.
func (u *HapticEffect) Constant() HapticConstant {
    p := (*C.SDL_HapticConstant)(unsafe.Pointer(u))
    return HapticConstant{uint16(p._type), HapticDirection{uint8(p.direction._type), *(*[3]int32)(unsafe.Pointer(&(p.direction.dir)))}, uint32(p.length), uint16(p.delay), uint16(p.button), uint16(p.interval), int16(p.level), uint16(p.attack_length), uint16(p.attack_level), uint16(p.fade_length), uint16(p.fade_level)}
}
 // Constant effect.
func (u *HapticEffect) SetConstant(x HapticConstant) {
    p := (*C.SDL_HapticConstant)(unsafe.Pointer(u))
    p._type = C.Uint16(x.Type)
    p.direction._type = C.Uint8(x.Direction.Type)
    p.direction.dir = *(*[3]C.Sint32)(unsafe.Pointer(&(x.Direction.Dir)))
    p.length = C.Uint32(x.Length)
    p.delay = C.Uint16(x.Delay)
    p.button = C.Uint16(x.Button)
    p.interval = C.Uint16(x.Interval)
    p.level = C.Sint16(x.Level)
    p.attack_length = C.Uint16(x.Attack_length)
    p.attack_level = C.Uint16(x.Attack_level)
    p.fade_length = C.Uint16(x.Fade_length)
    p.fade_level = C.Uint16(x.Fade_level)
}

 // Periodic effect.
func (u *HapticEffect) Periodic() HapticPeriodic {
    p := (*C.SDL_HapticPeriodic)(unsafe.Pointer(u))
    return HapticPeriodic{uint16(p._type), HapticDirection{uint8(p.direction._type), *(*[3]int32)(unsafe.Pointer(&(p.direction.dir)))}, uint32(p.length), uint16(p.delay), uint16(p.button), uint16(p.interval), uint16(p.period), int16(p.magnitude), int16(p.offset), uint16(p.phase), uint16(p.attack_length), uint16(p.attack_level), uint16(p.fade_length), uint16(p.fade_level)}
}
 // Periodic effect.
func (u *HapticEffect) SetPeriodic(x HapticPeriodic) {
    p := (*C.SDL_HapticPeriodic)(unsafe.Pointer(u))
    p._type = C.Uint16(x.Type)
    p.direction._type = C.Uint8(x.Direction.Type)
    p.direction.dir = *(*[3]C.Sint32)(unsafe.Pointer(&(x.Direction.Dir)))
    p.length = C.Uint32(x.Length)
    p.delay = C.Uint16(x.Delay)
    p.button = C.Uint16(x.Button)
    p.interval = C.Uint16(x.Interval)
    p.period = C.Uint16(x.Period)
    p.magnitude = C.Sint16(x.Magnitude)
    p.offset = C.Sint16(x.Offset)
    p.phase = C.Uint16(x.Phase)
    p.attack_length = C.Uint16(x.Attack_length)
    p.attack_level = C.Uint16(x.Attack_level)
    p.fade_length = C.Uint16(x.Fade_length)
    p.fade_level = C.Uint16(x.Fade_level)
}

 // Condition effect.
func (u *HapticEffect) Condition() HapticCondition {
    p := (*C.SDL_HapticCondition)(unsafe.Pointer(u))
    return HapticCondition{uint16(p._type), HapticDirection{uint8(p.direction._type), *(*[3]int32)(unsafe.Pointer(&(p.direction.dir)))}, uint32(p.length), uint16(p.delay), uint16(p.button), uint16(p.interval), *(*[3]uint16)(unsafe.Pointer(&(p.right_sat))), *(*[3]uint16)(unsafe.Pointer(&(p.left_sat))), *(*[3]int16)(unsafe.Pointer(&(p.right_coeff))), *(*[3]int16)(unsafe.Pointer(&(p.left_coeff))), *(*[3]uint16)(unsafe.Pointer(&(p.deadband))), *(*[3]int16)(unsafe.Pointer(&(p.center)))}
}
 // Condition effect.
func (u *HapticEffect) SetCondition(x HapticCondition) {
    p := (*C.SDL_HapticCondition)(unsafe.Pointer(u))
    p._type = C.Uint16(x.Type)
    p.direction._type = C.Uint8(x.Direction.Type)
    p.direction.dir = *(*[3]C.Sint32)(unsafe.Pointer(&(x.Direction.Dir)))
    p.length = C.Uint32(x.Length)
    p.delay = C.Uint16(x.Delay)
    p.button = C.Uint16(x.Button)
    p.interval = C.Uint16(x.Interval)
    p.right_sat = *(*[3]C.Uint16)(unsafe.Pointer(&(x.Right_sat)))
    p.left_sat = *(*[3]C.Uint16)(unsafe.Pointer(&(x.Left_sat)))
    p.right_coeff = *(*[3]C.Sint16)(unsafe.Pointer(&(x.Right_coeff)))
    p.left_coeff = *(*[3]C.Sint16)(unsafe.Pointer(&(x.Left_coeff)))
    p.deadband = *(*[3]C.Uint16)(unsafe.Pointer(&(x.Deadband)))
    p.center = *(*[3]C.Sint16)(unsafe.Pointer(&(x.Center)))
}

 // Ramp effect.
func (u *HapticEffect) Ramp() HapticRamp {
    p := (*C.SDL_HapticRamp)(unsafe.Pointer(u))
    return HapticRamp{uint16(p._type), HapticDirection{uint8(p.direction._type), *(*[3]int32)(unsafe.Pointer(&(p.direction.dir)))}, uint32(p.length), uint16(p.delay), uint16(p.button), uint16(p.interval), int16(p.start), int16(p.end), uint16(p.attack_length), uint16(p.attack_level), uint16(p.fade_length), uint16(p.fade_level)}
}
 // Ramp effect.
func (u *HapticEffect) SetRamp(x HapticRamp) {
    p := (*C.SDL_HapticRamp)(unsafe.Pointer(u))
    p._type = C.Uint16(x.Type)
    p.direction._type = C.Uint8(x.Direction.Type)
    p.direction.dir = *(*[3]C.Sint32)(unsafe.Pointer(&(x.Direction.Dir)))
    p.length = C.Uint32(x.Length)
    p.delay = C.Uint16(x.Delay)
    p.button = C.Uint16(x.Button)
    p.interval = C.Uint16(x.Interval)
    p.start = C.Sint16(x.Start)
    p.end = C.Sint16(x.End)
    p.attack_length = C.Uint16(x.Attack_length)
    p.attack_level = C.Uint16(x.Attack_level)
    p.fade_length = C.Uint16(x.Fade_length)
    p.fade_level = C.Uint16(x.Fade_level)
}

 // Left/Right effect.
func (u *HapticEffect) Leftright() HapticLeftRight {
    p := (*C.SDL_HapticLeftRight)(unsafe.Pointer(u))
    return HapticLeftRight{uint16(p._type), uint32(p.length), uint16(p.large_magnitude), uint16(p.small_magnitude)}
}
 // Left/Right effect.
func (u *HapticEffect) SetLeftright(x HapticLeftRight) {
    p := (*C.SDL_HapticLeftRight)(unsafe.Pointer(u))
    p._type = C.Uint16(x.Type)
    p.length = C.Uint32(x.Length)
    p.large_magnitude = C.Uint16(x.Large_magnitude)
    p.small_magnitude = C.Uint16(x.Small_magnitude)
}

 // Custom effect.
func (u *HapticEffect) Custom() HapticCustom {
    p := (*C.SDL_HapticCustom)(unsafe.Pointer(u))
    return HapticCustom(*p)
}
 // Custom effect.
func (u *HapticEffect) SetCustom(x HapticCustom) {
    p := (*C.SDL_HapticCustom)(unsafe.Pointer(u))
    *p = C.SDL_HapticCustom(x)
}

 // Haptic effects
const (
     // Constant effect supported.
     // 
     // Constant haptic effect.
     // 
     // See also: SDL_HapticCondition
     // 
    HAPTIC_CONSTANT = C.SDL_HAPTIC_CONSTANT

     // Sine wave effect supported.
     // 
     // Periodic haptic effect that simulates sine waves.
     // 
     // See also: SDL_HapticPeriodic
     // 
    HAPTIC_SINE = C.SDL_HAPTIC_SINE

     // Left/Right effect supported.
     // 
     // Haptic effect for direct control over high/low frequency motors.
     // 
     // See also: SDL_HapticLeftRight
     // 
     // Warning: this value was SDL_HAPTIC_SQUARE right before 2.0.0 shipped.
     // Sorry, we ran out of bits, and this is important for XInput devices.
     // 
    HAPTIC_LEFTRIGHT = C.SDL_HAPTIC_LEFTRIGHT

     // Triangle wave effect supported.
     // 
     // Periodic haptic effect that simulates triangular waves.
     // 
     // See also: SDL_HapticPeriodic
     // 
    HAPTIC_TRIANGLE = C.SDL_HAPTIC_TRIANGLE

     // Sawtoothup wave effect supported.
     // 
     // Periodic haptic effect that simulates saw tooth up waves.
     // 
     // See also: SDL_HapticPeriodic
     // 
    HAPTIC_SAWTOOTHUP = C.SDL_HAPTIC_SAWTOOTHUP

     // Sawtoothdown wave effect supported.
     // 
     // Periodic haptic effect that simulates saw tooth down waves.
     // 
     // See also: SDL_HapticPeriodic
     // 
    HAPTIC_SAWTOOTHDOWN = C.SDL_HAPTIC_SAWTOOTHDOWN

     // Ramp effect supported.
     // 
     // Ramp haptic effect.
     // 
     // See also: SDL_HapticRamp
     // 
    HAPTIC_RAMP = C.SDL_HAPTIC_RAMP

     // Spring effect supported - uses axes position.
     // 
     // Condition haptic effect that simulates a spring. Effect is based on
     // the axes position.
     // 
     // See also: SDL_HapticCondition
     // 
    HAPTIC_SPRING = C.SDL_HAPTIC_SPRING

     // Damper effect supported - uses axes velocity.
     // 
     // Condition haptic effect that simulates dampening. Effect is based on
     // the axes velocity.
     // 
     // See also: SDL_HapticCondition
     // 
    HAPTIC_DAMPER = C.SDL_HAPTIC_DAMPER

     // Inertia effect supported - uses axes acceleration.
     // 
     // Condition haptic effect that simulates inertia. Effect is based on the
     // axes acceleration.
     // 
     // See also: SDL_HapticCondition
     // 
    HAPTIC_INERTIA = C.SDL_HAPTIC_INERTIA

     // Friction effect supported - uses axes movement.
     // 
     // Condition haptic effect that simulates friction. Effect is based on
     // the axes movement.
     // 
     // See also: SDL_HapticCondition
     // 
    HAPTIC_FRICTION = C.SDL_HAPTIC_FRICTION

     // Custom effect is supported.
     // 
     // User defined custom haptic effect.
    HAPTIC_CUSTOM = C.SDL_HAPTIC_CUSTOM
)

 // Direction encodings
const (
     // Uses polar coordinates for the direction.
     // 
     // See also: SDL_HapticDirection
     // 
    HAPTIC_POLAR = C.SDL_HAPTIC_POLAR

     // Uses cartesian coordinates for the direction.
     // 
     // See also: SDL_HapticDirection
     // 
    HAPTIC_CARTESIAN = C.SDL_HAPTIC_CARTESIAN

     // Uses spherical coordinates for the direction.
     // 
     // See also: SDL_HapticDirection
     // 
    HAPTIC_SPHERICAL = C.SDL_HAPTIC_SPHERICAL
)

const (
     // Device can set global gain.
     // 
     // Device supports setting the global gain.
     // 
     // See also: SDL_HapticSetGain
     // 
    HAPTIC_GAIN = C.SDL_HAPTIC_GAIN

     // Device can set autocenter.
     // 
     // Device supports setting autocenter.
     // 
     // See also: SDL_HapticSetAutocenter
     // 
    HAPTIC_AUTOCENTER = C.SDL_HAPTIC_AUTOCENTER

     // Device can be queried for effect status.
     // 
     // Device supports querying effect status.
     // 
     // See also: SDL_HapticGetEffectStatus
     // 
    HAPTIC_STATUS = C.SDL_HAPTIC_STATUS

     // Device can be paused.
     // 
     // Devices supports being paused.
     // 
     // See also: SDL_HapticPause
     // 
     // See also: SDL_HapticUnpause
     // 
    HAPTIC_PAUSE = C.SDL_HAPTIC_PAUSE

     // Used to play a device an infinite number of times.
     // 
     // See also: SDL_HapticRunEffect
     // 
    HAPTIC_INFINITY = C.SDL_HAPTIC_INFINITY
)

 // The haptic structure used to identify an SDL haptic.
 // 
 // See also: SDL_HapticOpen
 // 
 // See also: SDL_HapticOpenFromJoystick
 // 
 // See also: SDL_HapticClose
 // 
type Haptic C.SDL_Haptic


 // Count the number of haptic devices attached to the system.
 // 
 // Returns: Number of haptic devices detected on the system.
 // 
 // ↪ https://wiki.libsdl.org/SDL_NumHaptics
func NumHaptics() (retval int) {
    retval = int(C.SDL_NumHaptics())
    return
}

 // Get the implementation dependent name of a haptic device.
 // 
 // This can be called before any joysticks are opened. If no name can be
 // found, this function returns NULL.
 // 
 // Returns: Name of the device or NULL on error.
 // 
 // See also: SDL_NumHaptics
 // 
 //   device_index
 //     Index of the device to get its name.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticName
func HapticName(device_index int) (retval string) {
    retval = C.GoString(C.SDL_HapticName(C.int(device_index)))
    return
}

 // Opens a haptic device for use.
 // 
 // The index passed as an argument refers to the N'th haptic device on
 // this system.
 // 
 // When opening a haptic device, its gain will be set to maximum and
 // autocenter will be disabled. To modify these values use
 // SDL_HapticSetGain() and SDL_HapticSetAutocenter().
 // 
 // Returns: Device identifier or NULL on error.
 // 
 // See also: SDL_HapticIndex
 // 
 // See also: SDL_HapticOpenFromMouse
 // 
 // See also: SDL_HapticOpenFromJoystick
 // 
 // See also: SDL_HapticClose
 // 
 // See also: SDL_HapticSetGain
 // 
 // See also: SDL_HapticSetAutocenter
 // 
 // See also: SDL_HapticPause
 // 
 // See also: SDL_HapticStopAll
 // 
 //   device_index
 //     Index of the device to open.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticOpen
func HapticOpen(device_index int) (retval *Haptic) {
    retval = (*Haptic)(unsafe.Pointer(C.SDL_HapticOpen(C.int(device_index))))
    return
}

 // Checks if the haptic device at index has been opened.
 // 
 // Returns: 1 if it has been opened or 0 if it hasn't.
 // 
 // See also: SDL_HapticOpen
 // 
 // See also: SDL_HapticIndex
 // 
 //   device_index
 //     Index to check to see if it has been opened.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticOpened
func HapticOpened(device_index int) (retval int) {
    retval = int(C.SDL_HapticOpened(C.int(device_index)))
    return
}

 // Gets the index of a haptic device.
 // 
 // Returns: The index of the haptic device or -1 on error.
 // 
 // See also: SDL_HapticOpen
 // 
 // See also: SDL_HapticOpened
 // 
 //   haptic
 //     Haptic device to get the index of.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticIndex
func (haptic *Haptic) Index() (retval int) {
    retval = int(C.SDL_HapticIndex((*C.SDL_Haptic)(haptic)))
    return
}

 // Gets whether or not the current mouse has haptic capabilities.
 // 
 // Returns: SDL_TRUE if the mouse is haptic, SDL_FALSE if it isn't.
 // 
 // See also: SDL_HapticOpenFromMouse
 // 
 // ↪ https://wiki.libsdl.org/SDL_MouseIsHaptic
func MouseIsHaptic() (retval int) {
    retval = int(C.SDL_MouseIsHaptic())
    return
}

 // Tries to open a haptic device from the current mouse.
 // 
 // Returns: The haptic device identifier or NULL on error.
 // 
 // See also: SDL_MouseIsHaptic
 // 
 // See also: SDL_HapticOpen
 // 
 // ↪ https://wiki.libsdl.org/SDL_HapticOpenFromMouse
func HapticOpenFromMouse() (retval *Haptic) {
    retval = (*Haptic)(unsafe.Pointer(C.SDL_HapticOpenFromMouse()))
    return
}

 // Checks to see if a joystick has haptic features.
 // 
 // Returns: SDL_TRUE if the joystick is haptic, SDL_FALSE if it isn't or
 // -1 if an error occurred.
 // 
 // See also: SDL_HapticOpenFromJoystick
 // 
 //   joystick
 //     Joystick to test for haptic capabilities.
 //   
 // ↪ https://wiki.libsdl.org/SDL_JoystickIsHaptic
func (joystick *Joystick) IsHaptic() (retval int) {
    retval = int(C.SDL_JoystickIsHaptic((*C.SDL_Joystick)(joystick)))
    return
}

 // Opens a haptic device for use from a joystick device.
 // 
 // You must still close the haptic device separately. It will not be
 // closed with the joystick.
 // 
 // When opening from a joystick you should first close the haptic device
 // before closing the joystick device. If not, on some implementations
 // the haptic device will also get unallocated and you'll be unable to
 // use force feedback on that device.
 // 
 // Returns: A valid haptic device identifier on success or NULL on error.
 // 
 // See also: SDL_HapticOpen
 // 
 // See also: SDL_HapticClose
 // 
 //   joystick
 //     Joystick to create a haptic device from.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticOpenFromJoystick
func (joystick *Joystick) HapticOpenFrom() (retval *Haptic) {
    retval = (*Haptic)(unsafe.Pointer(C.SDL_HapticOpenFromJoystick((*C.SDL_Joystick)(joystick))))
    return
}

 // Closes a haptic device previously opened with SDL_HapticOpen().
 // 
 //   haptic
 //     Haptic device to close.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticClose
func (haptic *Haptic) Close() {
    C.SDL_HapticClose((*C.SDL_Haptic)(haptic))
}

 // Returns the number of effects a haptic device can store.
 // 
 // On some platforms this isn't fully supported, and therefore is an
 // approximation. Always check to see if your created effect was actually
 // created and do not rely solely on SDL_HapticNumEffects().
 // 
 // Returns: The number of effects the haptic device can store or -1 on
 // error.
 // 
 // See also: SDL_HapticNumEffectsPlaying
 // 
 // See also: SDL_HapticQuery
 // 
 //   haptic
 //     The haptic device to query effect max.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticNumEffects
func (haptic *Haptic) NumEffects() (retval int) {
    retval = int(C.SDL_HapticNumEffects((*C.SDL_Haptic)(haptic)))
    return
}

 // Returns the number of effects a haptic device can play at the same
 // time.
 // 
 // This is not supported on all platforms, but will always return a
 // value. Added here for the sake of completeness.
 // 
 // Returns: The number of effects the haptic device can play at the same
 // time or -1 on error.
 // 
 // See also: SDL_HapticNumEffects
 // 
 // See also: SDL_HapticQuery
 // 
 //   haptic
 //     The haptic device to query maximum playing effects.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticNumEffectsPlaying
func (haptic *Haptic) NumEffectsPlaying() (retval int) {
    retval = int(C.SDL_HapticNumEffectsPlaying((*C.SDL_Haptic)(haptic)))
    return
}

 // Gets the haptic device's supported features in bitwise manner.
 // 
 // Example:
 //   if (SDL_HapticQuery(haptic) & SDL_HAPTIC_CONSTANT) {
 //       printf("We have constant haptic effect!\n");
 //   }
 // 
 // Returns: Haptic features in bitwise manner (OR'd).
 // 
 // See also: SDL_HapticNumEffects
 // 
 // See also: SDL_HapticEffectSupported
 // 
 //   haptic
 //     The haptic device to query.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticQuery
func (haptic *Haptic) Query() (retval uint) {
    retval = uint(C.SDL_HapticQuery((*C.SDL_Haptic)(haptic)))
    return
}

 // Gets the number of haptic axes the device has.
 // 
 // See also: SDL_HapticDirection
 // 
 // ↪ https://wiki.libsdl.org/SDL_HapticNumAxes
func (haptic *Haptic) NumAxes() (retval int) {
    retval = int(C.SDL_HapticNumAxes((*C.SDL_Haptic)(haptic)))
    return
}

 // Checks to see if effect is supported by haptic.
 // 
 // Returns: SDL_TRUE if effect is supported, SDL_FALSE if it isn't or -1
 // on error.
 // 
 // See also: SDL_HapticQuery
 // 
 // See also: SDL_HapticNewEffect
 // 
 //   haptic
 //     Haptic device to check on.
 //   
 //   effect
 //     Effect to check to see if it is supported.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticEffectSupported
func (haptic *Haptic) EffectSupported(effect *HapticEffect) (retval int) {
    retval = int(C.SDL_HapticEffectSupported((*C.SDL_Haptic)(haptic), (*C.SDL_HapticEffect)(effect)))
    return
}

 // Creates a new haptic effect on the device.
 // 
 // Returns: The identifier of the effect on success or -1 on error.
 // 
 // See also: SDL_HapticUpdateEffect
 // 
 // See also: SDL_HapticRunEffect
 // 
 // See also: SDL_HapticDestroyEffect
 // 
 //   haptic
 //     Haptic device to create the effect on.
 //   
 //   effect
 //     Properties of the effect to create.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticNewEffect
func (haptic *Haptic) NewEffect(effect *HapticEffect) (retval int) {
    retval = int(C.SDL_HapticNewEffect((*C.SDL_Haptic)(haptic), (*C.SDL_HapticEffect)(effect)))
    return
}

 // Updates the properties of an effect.
 // 
 // Can be used dynamically, although behavior when dynamically changing
 // direction may be strange. Specifically the effect may reupload itself
 // and start playing from the start. You cannot change the type either
 // when running SDL_HapticUpdateEffect().
 // 
 // Returns: 0 on success or -1 on error.
 // 
 // See also: SDL_HapticNewEffect
 // 
 // See also: SDL_HapticRunEffect
 // 
 // See also: SDL_HapticDestroyEffect
 // 
 //   haptic
 //     Haptic device that has the effect.
 //   
 //   effect
 //     Identifier of the effect to update.
 //   
 //   data
 //     New effect properties to use.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticUpdateEffect
func (haptic *Haptic) UpdateEffect(effect int, data *HapticEffect) (retval int) {
    retval = int(C.SDL_HapticUpdateEffect((*C.SDL_Haptic)(haptic), C.int(effect), (*C.SDL_HapticEffect)(data)))
    return
}

 // Runs the haptic effect on its associated haptic device.
 // 
 // If iterations are SDL_HAPTIC_INFINITY, it'll run the effect over and
 // over repeating the envelope (attack and fade) every time. If you only
 // want the effect to last forever, set SDL_HAPTIC_INFINITY in the
 // effect's length parameter.
 // 
 // Returns: 0 on success or -1 on error.
 // 
 // See also: SDL_HapticStopEffect
 // 
 // See also: SDL_HapticDestroyEffect
 // 
 // See also: SDL_HapticGetEffectStatus
 // 
 //   haptic
 //     Haptic device to run the effect on.
 //   
 //   effect
 //     Identifier of the haptic effect to run.
 //   
 //   iterations
 //     Number of iterations to run the effect. Use SDL_HAPTIC_INFINITY for
 //     infinity.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticRunEffect
func (haptic *Haptic) RunEffect(effect int, iterations uint32) (retval int) {
    retval = int(C.SDL_HapticRunEffect((*C.SDL_Haptic)(haptic), C.int(effect), C.Uint32(iterations)))
    return
}

 // Stops the haptic effect on its associated haptic device.
 // 
 // Returns: 0 on success or -1 on error.
 // 
 // See also: SDL_HapticRunEffect
 // 
 // See also: SDL_HapticDestroyEffect
 // 
 //   haptic
 //     Haptic device to stop the effect on.
 //   
 //   effect
 //     Identifier of the effect to stop.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticStopEffect
func (haptic *Haptic) StopEffect(effect int) (retval int) {
    retval = int(C.SDL_HapticStopEffect((*C.SDL_Haptic)(haptic), C.int(effect)))
    return
}

 // Destroys a haptic effect on the device.
 // 
 // This will stop the effect if it's running. Effects are automatically
 // destroyed when the device is closed.
 // 
 // See also: SDL_HapticNewEffect
 // 
 //   haptic
 //     Device to destroy the effect on.
 //   
 //   effect
 //     Identifier of the effect to destroy.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticDestroyEffect
func (haptic *Haptic) DestroyEffect(effect int) {
    C.SDL_HapticDestroyEffect((*C.SDL_Haptic)(haptic), C.int(effect))
}

 // Gets the status of the current effect on the haptic device.
 // 
 // Device must support the SDL_HAPTIC_STATUS feature.
 // 
 // Returns: 0 if it isn't playing, 1 if it is playing or -1 on error.
 // 
 // See also: SDL_HapticRunEffect
 // 
 // See also: SDL_HapticStopEffect
 // 
 //   haptic
 //     Haptic device to query the effect status on.
 //   
 //   effect
 //     Identifier of the effect to query its status.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticGetEffectStatus
func (haptic *Haptic) GetEffectStatus(effect int) (retval int) {
    retval = int(C.SDL_HapticGetEffectStatus((*C.SDL_Haptic)(haptic), C.int(effect)))
    return
}

 // Sets the global gain of the device.
 // 
 // Device must support the SDL_HAPTIC_GAIN feature.
 // 
 // The user may specify the maximum gain by setting the environment
 // variable SDL_HAPTIC_GAIN_MAX which should be between 0 and 100. All
 // calls to SDL_HapticSetGain() will scale linearly using
 // SDL_HAPTIC_GAIN_MAX as the maximum.
 // 
 // Returns: 0 on success or -1 on error.
 // 
 // See also: SDL_HapticQuery
 // 
 //   haptic
 //     Haptic device to set the gain on.
 //   
 //   gain
 //     Value to set the gain to, should be between 0 and 100.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticSetGain
func (haptic *Haptic) SetGain(gain int) (retval int) {
    retval = int(C.SDL_HapticSetGain((*C.SDL_Haptic)(haptic), C.int(gain)))
    return
}

 // Sets the global autocenter of the device.
 // 
 // Autocenter should be between 0 and 100. Setting it to 0 will disable
 // autocentering.
 // 
 // Device must support the SDL_HAPTIC_AUTOCENTER feature.
 // 
 // Returns: 0 on success or -1 on error.
 // 
 // See also: SDL_HapticQuery
 // 
 //   haptic
 //     Haptic device to set autocentering on.
 //   
 //   autocenter
 //     Value to set autocenter to, 0 disables autocentering.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticSetAutocenter
func (haptic *Haptic) SetAutocenter(autocenter int) (retval int) {
    retval = int(C.SDL_HapticSetAutocenter((*C.SDL_Haptic)(haptic), C.int(autocenter)))
    return
}

 // Pauses a haptic device.
 // 
 // Device must support the SDL_HAPTIC_PAUSE feature. Call
 // SDL_HapticUnpause() to resume playback.
 // 
 // Do not modify the effects nor add new ones while the device is paused.
 // That can cause all sorts of weird errors.
 // 
 // Returns: 0 on success or -1 on error.
 // 
 // See also: SDL_HapticUnpause
 // 
 //   haptic
 //     Haptic device to pause.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticPause
func (haptic *Haptic) Pause() (retval int) {
    retval = int(C.SDL_HapticPause((*C.SDL_Haptic)(haptic)))
    return
}

 // Unpauses a haptic device.
 // 
 // Call to unpause after SDL_HapticPause().
 // 
 // Returns: 0 on success or -1 on error.
 // 
 // See also: SDL_HapticPause
 // 
 //   haptic
 //     Haptic device to unpause.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticUnpause
func (haptic *Haptic) Unpause() (retval int) {
    retval = int(C.SDL_HapticUnpause((*C.SDL_Haptic)(haptic)))
    return
}

 // Stops all the currently playing effects on a haptic device.
 // 
 // Returns: 0 on success or -1 on error.
 // 
 //   haptic
 //     Haptic device to stop.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticStopAll
func (haptic *Haptic) StopAll() (retval int) {
    retval = int(C.SDL_HapticStopAll((*C.SDL_Haptic)(haptic)))
    return
}

 // Checks to see if rumble is supported on a haptic device.
 // 
 // Returns: SDL_TRUE if effect is supported, SDL_FALSE if it isn't or -1
 // on error.
 // 
 // See also: SDL_HapticRumbleInit
 // 
 // See also: SDL_HapticRumblePlay
 // 
 // See also: SDL_HapticRumbleStop
 // 
 //   haptic
 //     Haptic device to check to see if it supports rumble.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticRumbleSupported
func (haptic *Haptic) RumbleSupported() (retval int) {
    retval = int(C.SDL_HapticRumbleSupported((*C.SDL_Haptic)(haptic)))
    return
}

 // Initializes the haptic device for simple rumble playback.
 // 
 // Returns: 0 on success or -1 on error.
 // 
 // See also: SDL_HapticOpen
 // 
 // See also: SDL_HapticRumbleSupported
 // 
 // See also: SDL_HapticRumblePlay
 // 
 // See also: SDL_HapticRumbleStop
 // 
 //   haptic
 //     Haptic device to initialize for simple rumble playback.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticRumbleInit
func (haptic *Haptic) RumbleInit() (retval int) {
    retval = int(C.SDL_HapticRumbleInit((*C.SDL_Haptic)(haptic)))
    return
}

 // Runs simple rumble on a haptic device.
 // 
 // Returns: 0 on success or -1 on error.
 // 
 // See also: SDL_HapticRumbleSupported
 // 
 // See also: SDL_HapticRumbleInit
 // 
 // See also: SDL_HapticRumbleStop
 // 
 //   haptic
 //     Haptic device to play rumble effect on.
 //   
 //   strength
 //     Strength of the rumble to play as a 0-1 float value.
 //   
 //   length
 //     Length of the rumble to play in milliseconds.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticRumblePlay
func (haptic *Haptic) RumblePlay(strength float32, length uint32) (retval int) {
    retval = int(C.SDL_HapticRumblePlay((*C.SDL_Haptic)(haptic), C.float(strength), C.Uint32(length)))
    return
}

 // Stops the simple rumble on a haptic device.
 // 
 // Returns: 0 on success or -1 on error.
 // 
 // See also: SDL_HapticRumbleSupported
 // 
 // See also: SDL_HapticRumbleInit
 // 
 // See also: SDL_HapticRumblePlay
 // 
 //   haptic
 //     Haptic to stop the rumble on.
 //   
 // ↪ https://wiki.libsdl.org/SDL_HapticRumbleStop
func (haptic *Haptic) RumbleStop() (retval int) {
    retval = int(C.SDL_HapticRumbleStop((*C.SDL_Haptic)(haptic)))
    return
}
