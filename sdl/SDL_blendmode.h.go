// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"

 // Header file declaring the SDL_BlendMode enumeration

 // The blend mode used in SDL_RenderCopy() and drawing operations.
 // ↪ https://wiki.libsdl.org/SDL_BlendMode
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

    BLENDMODE_INVALID BlendMode = C.SDL_BLENDMODE_INVALID
)

 // The blend operation used when combining source and destination pixel
 // components.
type BlendOperation int
const (
     // dst + src: supported by all renderers
    BLENDOPERATION_ADD BlendOperation = C.SDL_BLENDOPERATION_ADD

     // dst - src : supported by D3D9, D3D11, OpenGL, OpenGLES
    BLENDOPERATION_SUBTRACT BlendOperation = C.SDL_BLENDOPERATION_SUBTRACT

     // src - dst : supported by D3D9, D3D11, OpenGL, OpenGLES
    BLENDOPERATION_REV_SUBTRACT BlendOperation = C.SDL_BLENDOPERATION_REV_SUBTRACT

     // min(dst, src) : supported by D3D11
    BLENDOPERATION_MINIMUM BlendOperation = C.SDL_BLENDOPERATION_MINIMUM

     // max(dst, src) : supported by D3D11
    BLENDOPERATION_MAXIMUM BlendOperation = C.SDL_BLENDOPERATION_MAXIMUM
)

 // The normalized factor used to multiply pixel components.
type BlendFactor int
const (
     // 0, 0, 0, 0
    BLENDFACTOR_ZERO BlendFactor = C.SDL_BLENDFACTOR_ZERO

     // 1, 1, 1, 1
    BLENDFACTOR_ONE BlendFactor = C.SDL_BLENDFACTOR_ONE

     // srcR, srcG, srcB, srcA
    BLENDFACTOR_SRC_COLOR BlendFactor = C.SDL_BLENDFACTOR_SRC_COLOR

     // 1-srcR, 1-srcG, 1-srcB, 1-srcA
    BLENDFACTOR_ONE_MINUS_SRC_COLOR BlendFactor = C.SDL_BLENDFACTOR_ONE_MINUS_SRC_COLOR

     // srcA, srcA, srcA, srcA
    BLENDFACTOR_SRC_ALPHA BlendFactor = C.SDL_BLENDFACTOR_SRC_ALPHA

     // 1-srcA, 1-srcA, 1-srcA, 1-srcA
    BLENDFACTOR_ONE_MINUS_SRC_ALPHA BlendFactor = C.SDL_BLENDFACTOR_ONE_MINUS_SRC_ALPHA

     // dstR, dstG, dstB, dstA
    BLENDFACTOR_DST_COLOR BlendFactor = C.SDL_BLENDFACTOR_DST_COLOR

     // 1-dstR, 1-dstG, 1-dstB, 1-dstA
    BLENDFACTOR_ONE_MINUS_DST_COLOR BlendFactor = C.SDL_BLENDFACTOR_ONE_MINUS_DST_COLOR

     // dstA, dstA, dstA, dstA
    BLENDFACTOR_DST_ALPHA BlendFactor = C.SDL_BLENDFACTOR_DST_ALPHA

     // 1-dstA, 1-dstA, 1-dstA, 1-dstA
    BLENDFACTOR_ONE_MINUS_DST_ALPHA BlendFactor = C.SDL_BLENDFACTOR_ONE_MINUS_DST_ALPHA
)


 // Create a custom blend mode, which may or may not be supported by a
 // given renderer.
 // 
 //   srcColorFactor
 //   
 //   
 //   dstColorFactor
 //   
 //   
 //   colorOperation
 //   
 //   
 //   srcAlphaFactor
 //   
 //   
 //   dstAlphaFactor
 //   
 //   
 //   alphaOperation
 //     The result of the blend mode operation will be: dstRGB = dstRGB *
 //     dstColorFactor colorOperation srcRGB * srcColorFactor and dstA = dstA
 //     * dstAlphaFactor alphaOperation srcA * srcAlphaFactor
 //   
func ComposeCustomBlendMode(srcColorFactor BlendFactor, dstColorFactor BlendFactor, colorOperation BlendOperation, srcAlphaFactor BlendFactor, dstAlphaFactor BlendFactor, alphaOperation BlendOperation) (retval BlendMode) {
    retval = BlendMode(C.SDL_ComposeCustomBlendMode(C.SDL_BlendFactor(srcColorFactor), C.SDL_BlendFactor(dstColorFactor), C.SDL_BlendOperation(colorOperation), C.SDL_BlendFactor(srcAlphaFactor), C.SDL_BlendFactor(dstAlphaFactor), C.SDL_BlendOperation(alphaOperation)))
    return
}
