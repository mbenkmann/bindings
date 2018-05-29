// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"

 // CPU feature detection for SDL.

const (
     // ↪ https://wiki.libsdl.org/SDL_CACHELINE_SIZE
    CACHELINE_SIZE = C.SDL_CACHELINE_SIZE
)


 // This function returns the number of CPU cores available.
 // ↪ https://wiki.libsdl.org/SDL_GetCPUCount
func GetCPUCount() (retval int) {
    retval = int(C.SDL_GetCPUCount())
    return
}

 // This function returns the L1 cache line size of the CPU
 // 
 // This is useful for determining multi-threaded structure padding or
 // SIMD prefetch sizes.
 // ↪ https://wiki.libsdl.org/SDL_GetCPUCacheLineSize
func GetCPUCacheLineSize() (retval int) {
    retval = int(C.SDL_GetCPUCacheLineSize())
    return
}

 // This function returns true if the CPU has the RDTSC instruction.
 // ↪ https://wiki.libsdl.org/SDL_HasRDTSC
func HasRDTSC() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasRDTSC())
    return
}

 // This function returns true if the CPU has AltiVec features.
 // ↪ https://wiki.libsdl.org/SDL_HasAltiVec
func HasAltiVec() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasAltiVec())
    return
}

 // This function returns true if the CPU has MMX features.
 // ↪ https://wiki.libsdl.org/SDL_HasMMX
func HasMMX() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasMMX())
    return
}

 // This function returns true if the CPU has 3DNow! features.
 // ↪ https://wiki.libsdl.org/SDL_Has3DNow
func Has3DNow() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_Has3DNow())
    return
}

 // This function returns true if the CPU has SSE features.
 // ↪ https://wiki.libsdl.org/SDL_HasSSE
func HasSSE() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasSSE())
    return
}

 // This function returns true if the CPU has SSE2 features.
 // ↪ https://wiki.libsdl.org/SDL_HasSSE2
func HasSSE2() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasSSE2())
    return
}

 // This function returns true if the CPU has SSE3 features.
 // ↪ https://wiki.libsdl.org/SDL_HasSSE3
func HasSSE3() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasSSE3())
    return
}

 // This function returns true if the CPU has SSE4.1 features.
 // ↪ https://wiki.libsdl.org/SDL_HasSSE41
func HasSSE41() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasSSE41())
    return
}

 // This function returns true if the CPU has SSE4.2 features.
 // ↪ https://wiki.libsdl.org/SDL_HasSSE42
func HasSSE42() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasSSE42())
    return
}

 // This function returns true if the CPU has AVX features.
 // ↪ https://wiki.libsdl.org/SDL_HasAVX
func HasAVX() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasAVX())
    return
}

 // This function returns true if the CPU has AVX2 features.
 // ↪ https://wiki.libsdl.org/SDL_HasAVX2
func HasAVX2() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasAVX2())
    return
}

 // This function returns true if the CPU has NEON (ARM SIMD) features.
func HasNEON() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasNEON())
    return
}

 // This function returns the amount of RAM configured in the system, in
 // MB.
 // ↪ https://wiki.libsdl.org/SDL_GetSystemRAM
func GetSystemRAM() (retval int) {
    retval = int(C.SDL_GetSystemRAM())
    return
}
