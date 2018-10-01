// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"
import "unsafe"

 // Access to the raw audio mixing buffer for the SDL library.

 // The calculated values in this structure are calculated by
 // SDL_OpenAudio().
 // 
 // For multi-channel audio, the default SDL channel mapping is: 2: FL FR
 // (stereo) 3: FL FR LFE (2.1 surround) 4: FL FR BL BR (quad) 5: FL FR FC
 // BL BR (quad + center) 6: FL FR FC LFE SL SR (5.1 surround - last two
 // can also be BL BR) 7: FL FR FC LFE BC SL SR (6.1 surround) 8: FL FR FC
 // LFE BL BR SL SR (7.1 surround)
 // ↪ https://wiki.libsdl.org/SDL_AudioSpec
type AudioSpec struct {
     // DSP frequency -- samples per second
    Freq int

     // Audio data format
    Format AudioFormat

     // Number of channels: 1 mono, 2 stereo
    Channels uint8

     // Audio buffer silence value (calculated)
    Silence uint8

     // Audio buffer size in sample FRAMES (total samples divided by channel
     // count)
    Samples uint16

     // Necessary for some compile environments
    Padding uint16

     // Audio buffer size in bytes (calculated)
    Size uint32

     // Callback that feeds the audio device (NULL to use SDL_QueueAudio()).
    Callback AudioCallback

     // Userdata passed to callback (ignored for NULL callbacks).
    Userdata uintptr
}

func fromC2AudioSpec(s C.SDL_AudioSpec) AudioSpec {
    return AudioSpec{int(s.freq), AudioFormat(s.format), uint8(s.channels), uint8(s.silence), uint16(s.samples), uint16(s.padding), uint32(s.size), AudioCallback(s.callback), uintptr(s.userdata)}
}

func toCFromAudioSpec(s AudioSpec) (d C.SDL_AudioSpec) {
    d.freq = C.int(s.Freq)
    d.format = C.SDL_AudioFormat(s.Format)
    d.channels = C.Uint8(s.Channels)
    d.silence = C.Uint8(s.Silence)
    d.samples = C.Uint16(s.Samples)
    d.padding = C.Uint16(s.Padding)
    d.size = C.Uint32(s.Size)
    d.callback = C.SDL_AudioCallback(s.Callback)
    d.userdata = unsafe.Pointer(s.Userdata)
    return
}

 // Audio flags
const (
    AUDIO_MASK_BITSIZE = C.SDL_AUDIO_MASK_BITSIZE

    AUDIO_MASK_DATATYPE = C.SDL_AUDIO_MASK_DATATYPE

    AUDIO_MASK_ENDIAN = C.SDL_AUDIO_MASK_ENDIAN

    AUDIO_MASK_SIGNED = C.SDL_AUDIO_MASK_SIGNED
)

 // Audio format flags
 // 
 // Defaults to LSB byte order.
const (
     // Unsigned 8-bit samples
    AUDIO_U8 = C.AUDIO_U8

     // Signed 8-bit samples
    AUDIO_S8 = C.AUDIO_S8

     // Unsigned 16-bit samples
    AUDIO_U16LSB = C.AUDIO_U16LSB

     // Signed 16-bit samples
    AUDIO_S16LSB = C.AUDIO_S16LSB

     // As above, but big-endian byte order
    AUDIO_U16MSB = C.AUDIO_U16MSB

     // As above, but big-endian byte order
    AUDIO_S16MSB = C.AUDIO_S16MSB

    AUDIO_U16 = C.AUDIO_U16

    AUDIO_S16 = C.AUDIO_S16
)

 // int32 support
const (
     // 32-bit integer samples
    AUDIO_S32LSB = C.AUDIO_S32LSB

     // As above, but big-endian byte order
    AUDIO_S32MSB = C.AUDIO_S32MSB

    AUDIO_S32 = C.AUDIO_S32
)

 // float32 support
const (
     // 32-bit floating point samples
    AUDIO_F32LSB = C.AUDIO_F32LSB

     // As above, but big-endian byte order
    AUDIO_F32MSB = C.AUDIO_F32MSB

    AUDIO_F32 = C.AUDIO_F32
)

 // Native audio byte ordering
const (
    AUDIO_U16SYS = C.AUDIO_U16SYS

    AUDIO_S16SYS = C.AUDIO_S16SYS

    AUDIO_S32SYS = C.AUDIO_S32SYS

    AUDIO_F32SYS = C.AUDIO_F32SYS
)

 // Allow change flags
 // 
 // Which audio format changes are allowed when opening a device.
const (
    AUDIO_ALLOW_FREQUENCY_CHANGE = C.SDL_AUDIO_ALLOW_FREQUENCY_CHANGE

    AUDIO_ALLOW_FORMAT_CHANGE = C.SDL_AUDIO_ALLOW_FORMAT_CHANGE

    AUDIO_ALLOW_CHANNELS_CHANGE = C.SDL_AUDIO_ALLOW_CHANNELS_CHANGE

    AUDIO_ALLOW_SAMPLES_CHANGE = C.SDL_AUDIO_ALLOW_SAMPLES_CHANGE

    AUDIO_ALLOW_ANY_CHANGE = C.SDL_AUDIO_ALLOW_ANY_CHANGE
)

 // Driver discovery functions
 // 
 // These functions return the list of built in audio drivers, in the
 // order that they are normally initialized by default.

 // ↪ https://wiki.libsdl.org/SDL_GetNumAudioDrivers
func GetNumAudioDrivers() (retval int) {
    retval = int(C.SDL_GetNumAudioDrivers())
    return
}

 // ↪ https://wiki.libsdl.org/SDL_GetAudioDriver
func GetAudioDriver(index int) (retval string) {
    retval = C.GoString(C.SDL_GetAudioDriver(C.int(index)))
    return
}

 // Initialization and cleanup

 // ↪ https://wiki.libsdl.org/SDL_AudioInit
func AudioInit(driver_name string) (retval int) {
    tmp_driver_name := C.CString(driver_name); defer C.free(unsafe.Pointer(tmp_driver_name))
    retval = int(C.SDL_AudioInit((*C.char)(tmp_driver_name)))
    return
}

 // ↪ https://wiki.libsdl.org/SDL_AudioQuit
func AudioQuit() {
    C.SDL_AudioQuit()
}

 // Audio state
 // 
 // Get the current audio state.
 // ↪ https://wiki.libsdl.org/SDL_AudioStatus
type AudioStatus int
const (
    AUDIO_STOPPED AudioStatus = C.SDL_AUDIO_STOPPED

    AUDIO_PLAYING AudioStatus = C.SDL_AUDIO_PLAYING

    AUDIO_PAUSED AudioStatus = C.SDL_AUDIO_PAUSED
)

 // ↪ https://wiki.libsdl.org/SDL_GetAudioStatus
func GetAudioStatus() (retval AudioStatus) {
    retval = AudioStatus(C.SDL_GetAudioStatus())
    return
}

 // ↪ https://wiki.libsdl.org/SDL_GetAudioDeviceStatus
func GetAudioDeviceStatus(dev AudioDeviceID) (retval AudioStatus) {
    retval = AudioStatus(C.SDL_GetAudioDeviceStatus(C.SDL_AudioDeviceID(dev)))
    return
}

 // Pause audio functions
 // 
 // These functions pause and unpause the audio callback processing. They
 // should be called with a parameter of 0 after opening the audio device
 // to start playing sound. This is so you can safely initialize data for
 // your callback function after opening the audio device. Silence will be
 // written to the audio device during the pause.

 // ↪ https://wiki.libsdl.org/SDL_PauseAudio
func PauseAudio(pause_on int) {
    C.SDL_PauseAudio(C.int(pause_on))
}

 // ↪ https://wiki.libsdl.org/SDL_PauseAudioDevice
func PauseAudioDevice(dev AudioDeviceID, pause_on int) {
    C.SDL_PauseAudioDevice(C.SDL_AudioDeviceID(dev), C.int(pause_on))
}

 // Audio lock functions
 // 
 // The lock manipulated by these functions protects the callback
 // function. During a SDL_LockAudio()/SDL_UnlockAudio() pair, you can be
 // guaranteed that the callback function is not running. Do not call
 // these from the callback function or you will cause deadlock.

 // ↪ https://wiki.libsdl.org/SDL_LockAudio
func LockAudio() {
    C.SDL_LockAudio()
}

 // ↪ https://wiki.libsdl.org/SDL_LockAudioDevice
func LockAudioDevice(dev AudioDeviceID) {
    C.SDL_LockAudioDevice(C.SDL_AudioDeviceID(dev))
}

 // ↪ https://wiki.libsdl.org/SDL_UnlockAudio
func UnlockAudio() {
    C.SDL_UnlockAudio()
}

 // ↪ https://wiki.libsdl.org/SDL_UnlockAudioDevice
func UnlockAudioDevice(dev AudioDeviceID) {
    C.SDL_UnlockAudioDevice(C.SDL_AudioDeviceID(dev))
}

const (
     // Upper limit of filters in SDL_AudioCVT.
     // 
     // The maximum number of SDL_AudioFilter functions in SDL_AudioCVT is
     // currently limited to 9. The SDL_AudioCVT.filters array has 10
     // pointers, one of which is the terminating NULL pointer.
    AUDIOCVT_MAX_FILTERS = C.SDL_AUDIOCVT_MAX_FILTERS

    MIX_MAXVOLUME = C.SDL_MIX_MAXVOLUME
)

 // This function is called when the audio device needs more data.
 // 
 //   userdata
 //     An application-specific parameter saved in the SDL_AudioSpec structure
 //   
 //   stream
 //     A pointer to the audio data buffer.
 //   
 //   len
 //     The length of that buffer in bytes.
 //   
 // Once the callback returns, the buffer will no longer be valid. Stereo
 // samples are stored in a LRLRLR ordering.
 // 
 // You can choose to avoid callbacks and use SDL_QueueAudio() instead, if
 // you like. Just open your audio device with a NULL callback.
type AudioCallback C.SDL_AudioCallback

type AudioFilter C.SDL_AudioFilter

type AudioStream C.SDL_AudioStream
 // Audio format flags.
 // 
 // These are what the 16 bits in SDL_AudioFormat currently mean...
 // (Unspecified bits are always zero).
 // 
 //   ++-----------------------sample is signed if set
 //   ||
 //   ||       ++-----------sample is bigendian if set
 //   ||       ||
 //   ||       ||          ++---sample is float if set
 //   ||       ||          ||
 //   ||       ||          || +---sample bit size---+
 //   ||       ||          || |                     |
 //   15 14 13 12 11 10 09 08 07 06 05 04 03 02 01 00
 // 
 // There are macros in SDL 2.0 and later to query these bits.
type AudioFormat uint16

 // SDL Audio Device IDs.
 // 
 // A successful call to SDL_OpenAudio() is always device id 1, and legacy
 // SDL audio APIs assume you want this device ID. SDL_OpenAudioDevice()
 // calls always returns devices >= 2 on success. The legacy calls are
 // good both for backwards compatibility and when you don't care about
 // multiple, specific, or capture devices.
type AudioDeviceID uint32


 // This function returns the name of the current audio driver, or NULL if
 // no driver has been initialized.
 // ↪ https://wiki.libsdl.org/SDL_GetCurrentAudioDriver
func GetCurrentAudioDriver() (retval string) {
    retval = C.GoString(C.SDL_GetCurrentAudioDriver())
    return
}

 // This function opens the audio device with the desired parameters, and
 // returns 0 if successful, placing the actual hardware parameters in the
 // structure pointed to by obtained. If obtained is NULL, the audio data
 // passed to the callback function will be guaranteed to be in the
 // requested format, and will be automatically converted to the hardware
 // audio format if necessary. This function returns -1 if it failed to
 // open the audio device, or couldn't set up the audio thread.
 // 
 // When filling in the desired audio spec structure,
 //   
 //   - desired->freq should be the desired audio frequency in samples-per-
 //     second.
 //   - desired->format should be the desired audio format.
 //   - desired->samples is the desired size of the audio buffer, in samples.
 //     This number should be a power of two, and may be adjusted by the audio
 //     driver to a value more suitable for the hardware. Good values seem to
 //     range between 512 and 8096 inclusive, depending on the application and
 //     CPU speed. Smaller values yield faster response time, but can lead to
 //     underflow if the application is doing heavy processing and cannot fill
 //     the audio buffer in time. A stereo sample consists of both right and
 //     left channels in LR ordering. Note that the number of samples is
 //     directly related to time by the following
 //     formula:ms=(samples*1000)/freq
 //   - desired->size is the size in bytes of the audio buffer, and is
 //     calculated by SDL_OpenAudio().
 //   - desired->silence is the value used to set the buffer to silence, and
 //     is calculated by SDL_OpenAudio().
 //   - desired->callback should be set to a function that will be called when
 //     the audio device is ready for more data. It is passed a pointer to the
 //     audio buffer, and the length in bytes of the audio buffer. This
 //     function usually runs in a separate thread, and so you should protect
 //     data structures that it accesses by calling SDL_LockAudio() and
 //     SDL_UnlockAudio() in your code. Alternately, you may pass a NULL
 //     pointer here, and call SDL_QueueAudio() with some frequency, to queue
 //     more audio samples to be played (or for capture devices, call
 //     SDL_DequeueAudio() with some frequency, to obtain audio samples).
 //   - desired->userdata is passed as the first parameter to your callback
 //     function. If you passed a NULL callback, this value is ignored.
 // 
 // The audio device starts out playing silence when it's opened, and
 // should be enabled for playing by calling SDL_PauseAudio(0) when you
 // are ready for your audio callback function to be called. Since the
 // audio driver may modify the requested size of the audio buffer, you
 // should allocate any local mixing buffers after you open the audio
 // device.
 // ↪ https://wiki.libsdl.org/SDL_OpenAudio
func OpenAudio(desired *AudioSpec, obtained *AudioSpec) (retval int) {
    var tmp_desired *C.SDL_AudioSpec; if desired != nil { x := toCFromAudioSpec(*desired); tmp_desired = &x }
    var tmp_obtained *C.SDL_AudioSpec; if obtained != nil { x := toCFromAudioSpec(*obtained); tmp_obtained = &x }
    retval = int(C.SDL_OpenAudio((*C.SDL_AudioSpec)(tmp_desired), (*C.SDL_AudioSpec)(tmp_obtained)))
    if obtained != nil { *obtained = fromC2AudioSpec(*(tmp_obtained)) }
    return
}

 // Get the number of available devices exposed by the current driver.
 // Only valid after a successfully initializing the audio subsystem.
 // Returns -1 if an explicit list of devices can't be determined; this is
 // not an error. For example, if SDL is set up to talk to a remote audio
 // server, it can't list every one available on the Internet, but it will
 // still allow a specific host to be specified to SDL_OpenAudioDevice().
 // 
 // In many common cases, when this function returns a value <= 0, it can
 // still successfully open the default device (NULL for first argument of
 // SDL_OpenAudioDevice()).
 // ↪ https://wiki.libsdl.org/SDL_GetNumAudioDevices
func GetNumAudioDevices(iscapture int) (retval int) {
    retval = int(C.SDL_GetNumAudioDevices(C.int(iscapture)))
    return
}

 // Get the human-readable name of a specific audio device. Must be a
 // value between 0 and (number of audio devices-1). Only valid after a
 // successfully initializing the audio subsystem. The values returned by
 // this function reflect the latest call to SDL_GetNumAudioDevices();
 // recall that function to redetect available hardware.
 // 
 // The string returned by this function is UTF-8 encoded, read-only, and
 // managed internally. You are not to free it. If you need to keep the
 // string for any length of time, you should make your own copy of it, as
 // it will be invalid next time any of several other SDL functions is
 // called.
 // ↪ https://wiki.libsdl.org/SDL_GetAudioDeviceName
func GetAudioDeviceName(index int, iscapture int) (retval string) {
    retval = C.GoString(C.SDL_GetAudioDeviceName(C.int(index), C.int(iscapture)))
    return
}

 // Open a specific audio device. Passing in a device name of NULL
 // requests the most reasonable default (and is equivalent to calling
 // SDL_OpenAudio()).
 // 
 // The device name is a UTF-8 string reported by
 // SDL_GetAudioDeviceName(), but some drivers allow arbitrary and driver-
 // specific strings, such as a hostname/IP address for a remote audio
 // server, or a filename in the diskaudio driver.
 // 
 // Returns: 0 on error, a valid device ID that is >= 2 on success.
 // 
 // SDL_OpenAudio(), unlike this function, always acts on device ID 1.
 // ↪ https://wiki.libsdl.org/SDL_OpenAudioDevice
func OpenAudioDevice(device string, iscapture int, desired *AudioSpec, obtained *AudioSpec, allowed_changes int) (retval AudioDeviceID) {
    tmp_device := C.CString(device); defer C.free(unsafe.Pointer(tmp_device))
    var tmp_desired *C.SDL_AudioSpec; if desired != nil { x := toCFromAudioSpec(*desired); tmp_desired = &x }
    var tmp_obtained *C.SDL_AudioSpec; if obtained != nil { x := toCFromAudioSpec(*obtained); tmp_obtained = &x }
    retval = AudioDeviceID(C.SDL_OpenAudioDevice((*C.char)(tmp_device), C.int(iscapture), (*C.SDL_AudioSpec)(tmp_desired), (*C.SDL_AudioSpec)(tmp_obtained), C.int(allowed_changes)))
    if obtained != nil { *obtained = fromC2AudioSpec(*(tmp_obtained)) }
    return
}





 // Create a new audio stream
 // 
 // Returns: 0 on success, or -1 on error.
 // 
 // See also: SDL_AudioStreamPut
 // 
 // See also: SDL_AudioStreamGet
 // 
 // See also: SDL_AudioStreamAvailable
 // 
 // See also: SDL_AudioStreamFlush
 // 
 // See also: SDL_AudioStreamClear
 // 
 // See also: SDL_FreeAudioStream
 // 
 //   src_format
 //     The format of the source audio
 //   
 //   src_channels
 //     The number of channels of the source audio
 //   
 //   src_rate
 //     The sampling rate of the source audio
 //   
 //   dst_format
 //     The format of the desired audio output
 //   
 //   dst_channels
 //     The number of channels of the desired audio output
 //   
 //   dst_rate
 //     The sampling rate of the desired audio output
 //   
func NewAudioStream(src_format AudioFormat, src_channels uint8, src_rate int, dst_format AudioFormat, dst_channels uint8, dst_rate int) (retval *AudioStream) {
    retval = (*AudioStream)(unsafe.Pointer(C.SDL_NewAudioStream(C.SDL_AudioFormat(src_format), C.Uint8(src_channels), C.int(src_rate), C.SDL_AudioFormat(dst_format), C.Uint8(dst_channels), C.int(dst_rate))))
    return
}

 // Add data to be converted/resampled to the stream
 // 
 // Returns: 0 on success, or -1 on error.
 // 
 // See also: SDL_NewAudioStream
 // 
 // See also: SDL_AudioStreamGet
 // 
 // See also: SDL_AudioStreamAvailable
 // 
 // See also: SDL_AudioStreamFlush
 // 
 // See also: SDL_AudioStreamClear
 // 
 // See also: SDL_FreeAudioStream
 // 
 //   stream
 //     The stream the audio data is being added to
 //   
 //   buf
 //     A pointer to the audio data to add
 //   
 //   len
 //     The number of bytes to write to the stream
 //   
func (stream *AudioStream) Put(buf []byte) (retval int) {
    var tmp_buf unsafe.Pointer
    if len(buf) > 0 {
        tmp_buf = (unsafe.Pointer)(unsafe.Pointer(&(buf[0])))
    }
    tmp_len := len(buf)
    retval = int(C.SDL_AudioStreamPut((*C.SDL_AudioStream)(stream), (tmp_buf), C.int(tmp_len)))
    return
}

 // Get converted/resampled data from the stream
 // 
 // Returns: The number of bytes read from the stream, or -1 on error
 // 
 // See also: SDL_NewAudioStream
 // 
 // See also: SDL_AudioStreamPut
 // 
 // See also: SDL_AudioStreamAvailable
 // 
 // See also: SDL_AudioStreamFlush
 // 
 // See also: SDL_AudioStreamClear
 // 
 // See also: SDL_FreeAudioStream
 // 
 //   stream
 //     The stream the audio is being requested from
 //   
 //   buf
 //     A buffer to fill with audio data
 //   
 //   len
 //     The maximum number of bytes to fill
 //   
func (stream *AudioStream) Get(buf []byte) (retval int) {
    var tmp_buf unsafe.Pointer
    if len(buf) > 0 {
        tmp_buf = (unsafe.Pointer)(unsafe.Pointer(&(buf[0])))
    }
    tmp_len := len(buf)
    retval = int(C.SDL_AudioStreamGet((*C.SDL_AudioStream)(stream), (tmp_buf), C.int(tmp_len)))
    return
}

 // Get the number of converted/resampled bytes available. The stream may
 // be buffering data behind the scenes until it has enough to resample
 // correctly, so this number might be lower than what you expect, or even
 // be zero. Add more data or flush the stream if you need the data now.
 // 
 // See also: SDL_NewAudioStream
 // 
 // See also: SDL_AudioStreamPut
 // 
 // See also: SDL_AudioStreamGet
 // 
 // See also: SDL_AudioStreamFlush
 // 
 // See also: SDL_AudioStreamClear
 // 
 // See also: SDL_FreeAudioStream
 // 
func (stream *AudioStream) Available() (retval int) {
    retval = int(C.SDL_AudioStreamAvailable((*C.SDL_AudioStream)(stream)))
    return
}

 // Tell the stream that you're done sending data, and anything being
 // buffered should be converted/resampled and made available immediately.
 // 
 // It is legal to add more data to a stream after flushing, but there
 // will be audio gaps in the output. Generally this is intended to signal
 // the end of input, so the complete output becomes available.
 // 
 // See also: SDL_NewAudioStream
 // 
 // See also: SDL_AudioStreamPut
 // 
 // See also: SDL_AudioStreamGet
 // 
 // See also: SDL_AudioStreamAvailable
 // 
 // See also: SDL_AudioStreamClear
 // 
 // See also: SDL_FreeAudioStream
 // 
func (stream *AudioStream) Flush() (retval int) {
    retval = int(C.SDL_AudioStreamFlush((*C.SDL_AudioStream)(stream)))
    return
}

 // Clear any pending data in the stream without converting it
 // 
 // See also: SDL_NewAudioStream
 // 
 // See also: SDL_AudioStreamPut
 // 
 // See also: SDL_AudioStreamGet
 // 
 // See also: SDL_AudioStreamAvailable
 // 
 // See also: SDL_AudioStreamFlush
 // 
 // See also: SDL_FreeAudioStream
 // 
func (stream *AudioStream) Clear() {
    C.SDL_AudioStreamClear((*C.SDL_AudioStream)(stream))
}

 // Free an audio stream
 // 
 // See also: SDL_NewAudioStream
 // 
 // See also: SDL_AudioStreamPut
 // 
 // See also: SDL_AudioStreamGet
 // 
 // See also: SDL_AudioStreamAvailable
 // 
 // See also: SDL_AudioStreamFlush
 // 
 // See also: SDL_AudioStreamClear
 // 
func (stream *AudioStream) Free() {
    C.SDL_FreeAudioStream((*C.SDL_AudioStream)(stream))
}

 // This takes two audio buffers of the playing audio format and mixes
 // them, performing addition, volume adjustment, and overflow clipping.
 // The volume ranges from 0 - 128, and should be set to SDL_MIX_MAXVOLUME
 // for full audio volume. Note this does not change hardware volume. This
 // is provided for convenience -- you can mix your own audio data.
 // ↪ https://wiki.libsdl.org/SDL_MixAudio
func MixAudio(dst []byte, src []byte, volume int) {
    var tmp_dst *C.Uint8
    if len(dst) > 0 {
        tmp_dst = (*C.Uint8)(unsafe.Pointer(&(dst[0])))
    }
    var tmp_src *C.Uint8
    if len(src) > 0 {
        tmp_src = (*C.Uint8)(unsafe.Pointer(&(src[0])))
    }
    tmp_len := len(dst); if len(src) < tmp_len { tmp_len = len(src) }
    C.SDL_MixAudio((tmp_dst), (tmp_src), C.Uint32(tmp_len), C.int(volume))
}

 // This works like SDL_MixAudio(), but you specify the audio format
 // instead of using the format of audio device 1. Thus it can be used
 // when no audio device is open at all.
 // ↪ https://wiki.libsdl.org/SDL_MixAudioFormat
func MixAudioFormat(dst []byte, src []byte, format AudioFormat, volume int) {
    var tmp_dst *C.Uint8
    if len(dst) > 0 {
        tmp_dst = (*C.Uint8)(unsafe.Pointer(&(dst[0])))
    }
    var tmp_src *C.Uint8
    if len(src) > 0 {
        tmp_src = (*C.Uint8)(unsafe.Pointer(&(src[0])))
    }
    tmp_len := len(dst); if len(src) < tmp_len { tmp_len = len(src) }
    C.SDL_MixAudioFormat((tmp_dst), (tmp_src), C.SDL_AudioFormat(format), C.Uint32(tmp_len), C.int(volume))
}

 // Queue more audio on non-callback devices.
 // 
 // (If you are looking to retrieve queued audio from a non-callback
 // capture device, you want SDL_DequeueAudio() instead. This will return
 // -1 to signify an error if you use it with capture devices.)
 // 
 // SDL offers two ways to feed audio to the device: you can either supply
 // a callback that SDL triggers with some frequency to obtain more audio
 // (pull method), or you can supply no callback, and then SDL will expect
 // you to supply data at regular intervals (push method) with this
 // function.
 // 
 // There are no limits on the amount of data you can queue, short of
 // exhaustion of address space. Queued data will drain to the device as
 // necessary without further intervention from you. If the device needs
 // audio but there is not enough queued, it will play silence to make up
 // the difference. This means you will have skips in your audio playback
 // if you aren't routinely queueing sufficient data.
 // 
 // This function copies the supplied data, so you are safe to free it
 // when the function returns. This function is thread-safe, but queueing
 // to the same device from two threads at once does not promise which
 // buffer will be queued first.
 // 
 // You may not queue audio on a device that is using an application-
 // supplied callback; doing so returns an error. You have to use the
 // audio callback or queue audio with this function, but not both.
 // 
 // You should not call SDL_LockAudio() on the device before queueing; SDL
 // handles locking internally for this function.
 // 
 // Returns: 0 on success, or -1 on error.
 // 
 // See also: SDL_GetQueuedAudioSize
 // 
 // See also: SDL_ClearQueuedAudio
 // 
 //   dev
 //     The device ID to which we will queue audio.
 //   
 //   data
 //     The data to queue to the device for later playback.
 //   
 //   len
 //     The number of bytes (not samples!) to which (data) points.
 //   
 // ↪ https://wiki.libsdl.org/SDL_QueueAudio
func QueueAudio(dev AudioDeviceID, data []byte) (retval int) {
    var tmp_data unsafe.Pointer
    if len(data) > 0 {
        tmp_data = (unsafe.Pointer)(unsafe.Pointer(&(data[0])))
    }
    tmp_len := len(data)
    retval = int(C.SDL_QueueAudio(C.SDL_AudioDeviceID(dev), (tmp_data), C.Uint32(tmp_len)))
    return
}

 // Dequeue more audio on non-callback devices.
 // 
 // (If you are looking to queue audio for output on a non-callback
 // playback device, you want SDL_QueueAudio() instead. This will always
 // return 0 if you use it with playback devices.)
 // 
 // SDL offers two ways to retrieve audio from a capture device: you can
 // either supply a callback that SDL triggers with some frequency as the
 // device records more audio data, (push method), or you can supply no
 // callback, and then SDL will expect you to retrieve data at regular
 // intervals (pull method) with this function.
 // 
 // There are no limits on the amount of data you can queue, short of
 // exhaustion of address space. Data from the device will keep queuing as
 // necessary without further intervention from you. This means you will
 // eventually run out of memory if you aren't routinely dequeueing data.
 // 
 // Capture devices will not queue data when paused; if you are expecting
 // to not need captured audio for some length of time, use
 // SDL_PauseAudioDevice() to stop the capture device from queueing more
 // data. This can be useful during, say, level loading times. When
 // unpaused, capture devices will start queueing data from that point,
 // having flushed any capturable data available while paused.
 // 
 // This function is thread-safe, but dequeueing from the same device from
 // two threads at once does not promise which thread will dequeued data
 // first.
 // 
 // You may not dequeue audio from a device that is using an application-
 // supplied callback; doing so returns an error. You have to use the
 // audio callback, or dequeue audio with this function, but not both.
 // 
 // You should not call SDL_LockAudio() on the device before queueing; SDL
 // handles locking internally for this function.
 // 
 // Returns: number of bytes dequeued, which could be less than requested.
 // 
 // See also: SDL_GetQueuedAudioSize
 // 
 // See also: SDL_ClearQueuedAudio
 // 
 //   dev
 //     The device ID from which we will dequeue audio.
 //   
 //   data
 //     A pointer into where audio data should be copied.
 //   
 //   len
 //     The number of bytes (not samples!) to which (data) points.
 //   
 // ↪ https://wiki.libsdl.org/SDL_DequeueAudio
func DequeueAudio(dev AudioDeviceID, data []byte) (retval uint32) {
    var tmp_data unsafe.Pointer
    if len(data) > 0 {
        tmp_data = (unsafe.Pointer)(unsafe.Pointer(&(data[0])))
    }
    tmp_len := len(data)
    retval = uint32(C.SDL_DequeueAudio(C.SDL_AudioDeviceID(dev), (tmp_data), C.Uint32(tmp_len)))
    return
}

 // Get the number of bytes of still-queued audio.
 // 
 // For playback device:
 // 
 // This is the number of bytes that have been queued for playback with
 // SDL_QueueAudio(), but have not yet been sent to the hardware. This
 // number may shrink at any time, so this only informs of pending data.
 // 
 // Once we've sent it to the hardware, this function can not decide the
 // exact byte boundary of what has been played. It's possible that we
 // just gave the hardware several kilobytes right before you called this
 // function, but it hasn't played any of it yet, or maybe half of it,
 // etc.
 // 
 // For capture devices:
 // 
 // This is the number of bytes that have been captured by the device and
 // are waiting for you to dequeue. This number may grow at any time, so
 // this only informs of the lower-bound of available data.
 // 
 // You may not queue audio on a device that is using an application-
 // supplied callback; calling this function on such a device always
 // returns 0. You have to queue audio with
 // SDL_QueueAudio()/SDL_DequeueAudio(), or use the audio callback, but
 // not both.
 // 
 // You should not call SDL_LockAudio() on the device before querying; SDL
 // handles locking internally for this function.
 // 
 // Returns: Number of bytes (not samples!) of queued audio.
 // 
 // See also: SDL_QueueAudio
 // 
 // See also: SDL_ClearQueuedAudio
 // 
 //   dev
 //     The device ID of which we will query queued audio size.
 //   
 // ↪ https://wiki.libsdl.org/SDL_GetQueuedAudioSize
func GetQueuedAudioSize(dev AudioDeviceID) (retval uint32) {
    retval = uint32(C.SDL_GetQueuedAudioSize(C.SDL_AudioDeviceID(dev)))
    return
}

 // Drop any queued audio data. For playback devices, this is any queued
 // data still waiting to be submitted to the hardware. For capture
 // devices, this is any data that was queued by the device that hasn't
 // yet been dequeued by the application.
 // 
 // Immediately after this call, SDL_GetQueuedAudioSize() will return 0.
 // For playback devices, the hardware will start playing silence if more
 // audio isn't queued. Unpaused capture devices will start filling the
 // queue again as soon as they have more data available (which, depending
 // on the state of the hardware and the thread, could be before this
 // function call returns!).
 // 
 // This will not prevent playback of queued audio that's already been
 // sent to the hardware, as we can not undo that, so expect there to be
 // some fraction of a second of audio that might still be heard. This can
 // be useful if you want to, say, drop any pending music during a level
 // change in your game.
 // 
 // You may not queue audio on a device that is using an application-
 // supplied callback; calling this function on such a device is always a
 // no-op. You have to queue audio with
 // SDL_QueueAudio()/SDL_DequeueAudio(), or use the audio callback, but
 // not both.
 // 
 // You should not call SDL_LockAudio() on the device before clearing the
 // queue; SDL handles locking internally for this function.
 // 
 // This function always succeeds and thus returns void.
 // 
 // See also: SDL_QueueAudio
 // 
 // See also: SDL_GetQueuedAudioSize
 // 
 //   dev
 //     The device ID of which to clear the audio queue.
 //   
 // ↪ https://wiki.libsdl.org/SDL_ClearQueuedAudio
func ClearQueuedAudio(dev AudioDeviceID) {
    C.SDL_ClearQueuedAudio(C.SDL_AudioDeviceID(dev))
}

 // This function shuts down audio processing and closes the audio device.
 // ↪ https://wiki.libsdl.org/SDL_CloseAudio
func CloseAudio() {
    C.SDL_CloseAudio()
}

 // ↪ https://wiki.libsdl.org/SDL_CloseAudioDevice
func CloseAudioDevice(dev AudioDeviceID) {
    C.SDL_CloseAudioDevice(C.SDL_AudioDeviceID(dev))
}
