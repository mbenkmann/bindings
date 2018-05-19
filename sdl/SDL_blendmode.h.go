// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
import "C"

 // Header file declaring the SDL_BlendMode enumeration

 // The blend mode used in SDL_RenderCopy() and drawing operations.
type BlendMode int
const (
     // no blending dstRGBA = srcRGBA
    BLENDMODE_NONE BlendMode = C.SDL_BLENDMODE_NONE

     // alpha blending dstRGB = (srcRGB * srcA) + (dstRGB * (1-srcA)) dstA =
     // srcA + (dstA * (1-srcA))
    BLENDMODE_BLEND BlendMode = C.SDL_BLENDMODE_BLEND

     // additive blending dstRGB = (srcRGB * srcA) + dstRGB dstA = dstA
    BLENDMODE_ADD BlendMode = C.SDL_BLENDMODE_ADD

     // color modulate dstRGB = srcRGB * dstRGB dstA = dstA
    BLENDMODE_MOD BlendMode = C.SDL_BLENDMODE_MOD
)
