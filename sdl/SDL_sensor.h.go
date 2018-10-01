// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"
import "unsafe"

 // Include file for SDL sensor event handling

const (
     // Accelerometer sensor
     // 
     // The accelerometer returns the current acceleration in SI meters per
     // second squared. This includes gravity, so a device at rest will have
     // an acceleration of SDL_STANDARD_GRAVITY straight down.
     // 
     // values[0]: Acceleration on the x axis values[1]: Acceleration on the y
     // axis values[2]: Acceleration on the z axis
     // 
     // For phones held in portrait mode, the axes are defined as follows: -X
     // ... +X : left ... right -Y ... +Y : bottom ... top -Z ... +Z : farther
     // ... closer
     // 
     // The axis data is not changed when the phone is rotated.
     // 
     // See also: SDL_GetDisplayOrientation()
     // 
    STANDARD_GRAVITY = C.SDL_STANDARD_GRAVITY
)

type SensorType int
const (
     // Returned for an invalid sensor
    SENSOR_INVALID SensorType = C.SDL_SENSOR_INVALID

     // Unknown sensor type
    SENSOR_UNKNOWN SensorType = C.SDL_SENSOR_UNKNOWN

     // Accelerometer
    SENSOR_ACCEL SensorType = C.SDL_SENSOR_ACCEL

     // Gyroscope
    SENSOR_GYRO SensorType = C.SDL_SENSOR_GYRO
)

type Sensor C.SDL_Sensor
 // This is a unique ID for a sensor for the time it is connected to the
 // system, and is never reused for the lifetime of the application.
 // 
 // The ID value starts at 0 and increments from there. The value -1 is an
 // invalid ID.
type SensorID int32


 // Count the number of sensors attached to the system right now.
 // 
 // Gyroscope sensor
 // 
 // The gyroscope returns the current rate of rotation in radians per
 // second. The rotation is positive in the counter-clockwise direction.
 // That is, an observer looking from a positive location on one of the
 // axes would see positive rotation on that axis when it appeared to be
 // rotating counter-clockwise.
 // 
 // values[0]: Angular speed around the x axis values[1]: Angular speed
 // around the y axis values[2]: Angular speed around the z axis
 // 
 // For phones held in portrait mode, the axes are defined as follows: -X
 // ... +X : left ... right -Y ... +Y : bottom ... top -Z ... +Z : farther
 // ... closer
 // 
 // The axis data is not changed when the phone is rotated.
 // 
 // See also: SDL_GetDisplayOrientation()
 // 
func NumSensors() (retval int) {
    retval = int(C.SDL_NumSensors())
    return
}

 // Get the implementation dependent name of a sensor.
 // 
 // This can be called before any sensors are opened.
 // 
 // Returns: The sensor name, or NULL if device_index is out of range.
 // 
func SensorGetDeviceName(device_index int) (retval string) {
    retval = C.GoString(C.SDL_SensorGetDeviceName(C.int(device_index)))
    return
}

 // Get the type of a sensor.
 // 
 // This can be called before any sensors are opened.
 // 
 // Returns: The sensor type, or SDL_SENSOR_INVALID if device_index is out
 // of range.
 // 
func SensorGetDeviceType(device_index int) (retval SensorType) {
    retval = SensorType(C.SDL_SensorGetDeviceType(C.int(device_index)))
    return
}

 // Get the platform dependent type of a sensor.
 // 
 // This can be called before any sensors are opened.
 // 
 // Returns: The sensor platform dependent type, or -1 if device_index is
 // out of range.
 // 
func SensorGetDeviceNonPortableType(device_index int) (retval int) {
    retval = int(C.SDL_SensorGetDeviceNonPortableType(C.int(device_index)))
    return
}

 // Get the instance ID of a sensor.
 // 
 // This can be called before any sensors are opened.
 // 
 // Returns: The sensor instance ID, or -1 if device_index is out of
 // range.
 // 
func SensorGetDeviceInstanceID(device_index int) (retval SensorID) {
    retval = SensorID(C.SDL_SensorGetDeviceInstanceID(C.int(device_index)))
    return
}

 // Open a sensor for use.
 // 
 // The index passed as an argument refers to the N'th sensor on the
 // system.
 // 
 // Returns: A sensor identifier, or NULL if an error occurred.
 // 
func SensorOpen(device_index int) (retval *Sensor) {
    retval = (*Sensor)(unsafe.Pointer(C.SDL_SensorOpen(C.int(device_index))))
    return
}

 // Return the SDL_Sensor associated with an instance id.
func SensorFromInstanceID(instance_id SensorID) (retval *Sensor) {
    retval = (*Sensor)(unsafe.Pointer(C.SDL_SensorFromInstanceID(C.SDL_SensorID(instance_id))))
    return
}

 // Get the implementation dependent name of a sensor.
 // 
 // Returns: The sensor name, or NULL if the sensor is NULL.
 // 
func (sensor *Sensor) GetName() (retval string) {
    retval = C.GoString(C.SDL_SensorGetName((*C.SDL_Sensor)(sensor)))
    return
}

 // Get the type of a sensor.
 // 
 // This can be called before any sensors are opened.
 // 
 // Returns: The sensor type, or SDL_SENSOR_INVALID if the sensor is NULL.
 // 
func (sensor *Sensor) GetType() (retval SensorType) {
    retval = SensorType(C.SDL_SensorGetType((*C.SDL_Sensor)(sensor)))
    return
}

 // Get the platform dependent type of a sensor.
 // 
 // This can be called before any sensors are opened.
 // 
 // Returns: The sensor platform dependent type, or -1 if the sensor is
 // NULL.
 // 
func (sensor *Sensor) GetNonPortableType() (retval int) {
    retval = int(C.SDL_SensorGetNonPortableType((*C.SDL_Sensor)(sensor)))
    return
}

 // Get the instance ID of a sensor.
 // 
 // This can be called before any sensors are opened.
 // 
 // Returns: The sensor instance ID, or -1 if the sensor is NULL.
 // 
func (sensor *Sensor) GetInstanceID() (retval SensorID) {
    retval = SensorID(C.SDL_SensorGetInstanceID((*C.SDL_Sensor)(sensor)))
    return
}

 // Get the current state of an opened sensor.
 // 
 // The number of values and interpretation of the data is sensor
 // dependent.
 // 
 // Returns: 0 or -1 if an error occurred.
 // 
 //   sensor
 //     The sensor to query
 //   
 //   data
 //     A pointer filled with the current sensor state
 //   
 //   num_values
 //     The number of values to write to data
 //   
func (sensor *Sensor) GetData(num_values int) (retval int, data float32) {
    tmp_data := new(C.float)
    retval = int(C.SDL_SensorGetData((*C.SDL_Sensor)(sensor), (*C.float)(tmp_data), C.int(num_values)))
    data = deref_float32_ptr(tmp_data)
    return
}

 // Close a sensor previously opened with SDL_SensorOpen()
func (sensor *Sensor) Close() {
    C.SDL_SensorClose((*C.SDL_Sensor)(sensor))
}

 // Update the current state of the open sensors.
 // 
 // This is called automatically by the event loop if sensor events are
 // enabled.
 // 
 // This needs to be called from the thread that initialized the sensor
 // subsystem.
func SensorUpdate() {
    C.SDL_SensorUpdate()
}
