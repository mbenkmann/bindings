// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
// #include <SDL_shape.h>
import "C"
import "unsafe"

 // This file provides a general interface for SDL to read and write data
 // streams. It can easily be extended to files, memory, etc.

 // RWFrom functions
 // 
 // Functions to create SDL_RWops structures from various data streams.

func RWFromFile(file string, mode string) (retval *RWops) {
    tmp_file := C.CString(file); defer C.free(unsafe.Pointer(tmp_file))
    tmp_mode := C.CString(mode); defer C.free(unsafe.Pointer(tmp_mode))
    retval = (*RWops)(unsafe.Pointer(C.SDL_RWFromFile((*C.char)(tmp_file), (*C.char)(tmp_mode))))
    return
}




 // Read/write macros
 // 
 // Macros to easily read and write from an SDL_RWops structure.

 // Read endian functions
 // 
 // Read an item of the specified endianness and return in native format.

func (src *RWops) ReadU8() (retval uint8) {
    retval = uint8(C.SDL_ReadU8((*C.SDL_RWops)(src)))
    return
}

func (src *RWops) ReadLE16() (retval uint16) {
    retval = uint16(C.SDL_ReadLE16((*C.SDL_RWops)(src)))
    return
}

func (src *RWops) ReadBE16() (retval uint16) {
    retval = uint16(C.SDL_ReadBE16((*C.SDL_RWops)(src)))
    return
}

func (src *RWops) ReadLE32() (retval uint32) {
    retval = uint32(C.SDL_ReadLE32((*C.SDL_RWops)(src)))
    return
}

func (src *RWops) ReadBE32() (retval uint32) {
    retval = uint32(C.SDL_ReadBE32((*C.SDL_RWops)(src)))
    return
}

func (src *RWops) ReadLE64() (retval uint64) {
    retval = uint64(C.SDL_ReadLE64((*C.SDL_RWops)(src)))
    return
}

func (src *RWops) ReadBE64() (retval uint64) {
    retval = uint64(C.SDL_ReadBE64((*C.SDL_RWops)(src)))
    return
}

 // Write endian functions
 // 
 // Write an item of native format to the specified endianness.

func (dst *RWops) WriteU8(value uint8) (retval uint64) {
    retval = uint64(C.SDL_WriteU8((*C.SDL_RWops)(dst), C.Uint8(value)))
    return
}

func (dst *RWops) WriteLE16(value uint16) (retval uint64) {
    retval = uint64(C.SDL_WriteLE16((*C.SDL_RWops)(dst), C.Uint16(value)))
    return
}

func (dst *RWops) WriteBE16(value uint16) (retval uint64) {
    retval = uint64(C.SDL_WriteBE16((*C.SDL_RWops)(dst), C.Uint16(value)))
    return
}

func (dst *RWops) WriteLE32(value uint32) (retval uint64) {
    retval = uint64(C.SDL_WriteLE32((*C.SDL_RWops)(dst), C.Uint32(value)))
    return
}

func (dst *RWops) WriteBE32(value uint32) (retval uint64) {
    retval = uint64(C.SDL_WriteBE32((*C.SDL_RWops)(dst), C.Uint32(value)))
    return
}

func (dst *RWops) WriteLE64(value uint64) (retval uint64) {
    retval = uint64(C.SDL_WriteLE64((*C.SDL_RWops)(dst), C.Uint64(value)))
    return
}

func (dst *RWops) WriteBE64(value uint64) (retval uint64) {
    retval = uint64(C.SDL_WriteBE64((*C.SDL_RWops)(dst), C.Uint64(value)))
    return
}

const (
    RWOPS_UNKNOWN = C.SDL_RWOPS_UNKNOWN

    RWOPS_WINFILE = C.SDL_RWOPS_WINFILE

    RWOPS_STDFILE = C.SDL_RWOPS_STDFILE

    RWOPS_JNIFILE = C.SDL_RWOPS_JNIFILE

    RWOPS_MEMORY = C.SDL_RWOPS_MEMORY

    RWOPS_MEMORY_RO = C.SDL_RWOPS_MEMORY_RO

     // Seek from the beginning of data
    RW_SEEK_SET = C.RW_SEEK_SET

     // Seek relative to current read point
    RW_SEEK_CUR = C.RW_SEEK_CUR

     // Seek relative to the end of data
    RW_SEEK_END = C.RW_SEEK_END
)



