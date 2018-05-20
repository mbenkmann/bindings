// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
// #include <SDL_shape.h>
import "C"

 // CPU feature detection for SDL.

const (
    CACHELINE_SIZE = C.SDL_CACHELINE_SIZE
)


 // This function returns the number of CPU cores available.
func GetCPUCount() (retval int) {
    retval = int(C.SDL_GetCPUCount())
    return
}

 // This function returns the L1 cache line size of the CPU
 // 
 // This is useful for determining multi-threaded structure padding or
 // SIMD prefetch sizes.
func GetCPUCacheLineSize() (retval int) {
    retval = int(C.SDL_GetCPUCacheLineSize())
    return
}

 // This function returns true if the CPU has the RDTSC instruction.
func HasRDTSC() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasRDTSC())
    return
}

 // This function returns true if the CPU has AltiVec features.
func HasAltiVec() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasAltiVec())
    return
}

 // This function returns true if the CPU has MMX features.
func HasMMX() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasMMX())
    return
}

 // This function returns true if the CPU has 3DNow! features.
func Has3DNow() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_Has3DNow())
    return
}

 // This function returns true if the CPU has SSE features.
func HasSSE() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasSSE())
    return
}

 // This function returns true if the CPU has SSE2 features.
func HasSSE2() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasSSE2())
    return
}

 // This function returns true if the CPU has SSE3 features.
func HasSSE3() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasSSE3())
    return
}

 // This function returns true if the CPU has SSE4.1 features.
func HasSSE41() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasSSE41())
    return
}

 // This function returns true if the CPU has SSE4.2 features.
func HasSSE42() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasSSE42())
    return
}

 // This function returns true if the CPU has AVX features.
func HasAVX() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasAVX())
    return
}

 // This function returns true if the CPU has AVX2 features.
func HasAVX2() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_HasAVX2())
    return
}

 // This function returns the amount of RAM configured in the system, in
 // MB.
func GetSystemRAM() (retval int) {
    retval = int(C.SDL_GetSystemRAM())
    return
}
