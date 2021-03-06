// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"
import "unsafe"

 // Header file for SDL 2D rendering functions.
 // 
 // This API supports the following features:
 //   
 //   - single pixel points
 //   - single pixel lines
 //   - filled rectangles
 //   - texture images
 // 
 // The primitives may be drawn in opaque, blended, or additive modes.
 // 
 // The texture images may be drawn in opaque, blended, or additive modes.
 // They can have an additional color tint or alpha modulation applied to
 // them, and may also be stretched with linear interpolation.
 // 
 // This API is designed to accelerate simple 2D operations. You may want
 // more functionality such as polygons and particle effects and in that
 // case you should use SDL's OpenGL/Direct3D support or one of the many
 // good 3D engines.
 // 
 // These functions must be called from the main thread. See this bug for
 // details: http://bugzilla.libsdl.org/show_bug.cgi?id=1995

 // Information on the capabilities of a render driver or context.
 // ↪ https://wiki.libsdl.org/SDL_RendererInfo
type RendererInfo struct {
     // The name of the renderer
    Name string

     // Supported SDL_RendererFlags
    Flags uint32

     // The number of available texture formats
    Num_texture_formats uint32

     // The available texture formats
    Texture_formats [16]uint32

     // The maximum texture width
    Max_texture_width int

     // The maximum texture height
    Max_texture_height int
}

func fromC2RendererInfo(s C.SDL_RendererInfo) RendererInfo {
    return RendererInfo{C.GoString(s.name), uint32(s.flags), uint32(s.num_texture_formats), *(*[16]uint32)(unsafe.Pointer(&(s.texture_formats))), int(s.max_texture_width), int(s.max_texture_height)}
}

 // Flags used when creating a rendering context.
 // ↪ https://wiki.libsdl.org/SDL_RendererFlags
type RendererFlags int
const (
     // The renderer is a software fallback
    RENDERER_SOFTWARE RendererFlags = C.SDL_RENDERER_SOFTWARE

     // The renderer uses hardware acceleration
    RENDERER_ACCELERATED RendererFlags = C.SDL_RENDERER_ACCELERATED

     // Present is synchronized with the refresh rate
    RENDERER_PRESENTVSYNC RendererFlags = C.SDL_RENDERER_PRESENTVSYNC

     // The renderer supports rendering to texture
    RENDERER_TARGETTEXTURE RendererFlags = C.SDL_RENDERER_TARGETTEXTURE
)

 // The access pattern allowed for a texture.
 // ↪ https://wiki.libsdl.org/SDL_TextureAccess
type TextureAccess int
const (
     // Changes rarely, not lockable
    TEXTUREACCESS_STATIC TextureAccess = C.SDL_TEXTUREACCESS_STATIC

     // Changes frequently, lockable
    TEXTUREACCESS_STREAMING TextureAccess = C.SDL_TEXTUREACCESS_STREAMING

     // Texture can be used as a render target
    TEXTUREACCESS_TARGET TextureAccess = C.SDL_TEXTUREACCESS_TARGET
)

 // The texture channel modulation used in SDL_RenderCopy().
 // ↪ https://wiki.libsdl.org/SDL_TextureModulate
type TextureModulate int
const (
     // No modulation
    TEXTUREMODULATE_NONE TextureModulate = C.SDL_TEXTUREMODULATE_NONE

     // srcC = srcC * color
    TEXTUREMODULATE_COLOR TextureModulate = C.SDL_TEXTUREMODULATE_COLOR

     // srcA = srcA * alpha
    TEXTUREMODULATE_ALPHA TextureModulate = C.SDL_TEXTUREMODULATE_ALPHA
)

 // Flip constants for SDL_RenderCopyEx.
 // ↪ https://wiki.libsdl.org/SDL_RendererFlip
type RendererFlip int
const (
     // Do not flip
    FLIP_NONE RendererFlip = C.SDL_FLIP_NONE

     // flip horizontally
    FLIP_HORIZONTAL RendererFlip = C.SDL_FLIP_HORIZONTAL

     // flip vertically
    FLIP_VERTICAL RendererFlip = C.SDL_FLIP_VERTICAL
)

type Renderer C.SDL_Renderer

type Texture C.SDL_Texture


 // Get the number of 2D rendering drivers available for the current
 // display.
 // 
 // A render driver is a set of code that handles rendering and texture
 // management on a particular display. Normally there is only one, but
 // some drivers may have several available with different capabilities.
 // 
 // See also: SDL_GetRenderDriverInfo()
 // 
 // See also: SDL_CreateRenderer()
 // 
 // ↪ https://wiki.libsdl.org/SDL_GetNumRenderDrivers
func GetNumRenderDrivers() (retval int) {
    retval = int(C.SDL_GetNumRenderDrivers())
    return
}

 // Get information about a specific 2D rendering driver for the current
 // display.
 // 
 // Returns: 0 on success, -1 if the index was out of range.
 // 
 // See also: SDL_CreateRenderer()
 // 
 //   index
 //     The index of the driver to query information about.
 //   
 //   info
 //     A pointer to an SDL_RendererInfo struct to be filled with information
 //     on the rendering driver.
 //   
 // ↪ https://wiki.libsdl.org/SDL_GetRenderDriverInfo
func GetRenderDriverInfo(index int) (retval int, info *RendererInfo) {
    tmp_info := new(C.SDL_RendererInfo)
    retval = int(C.SDL_GetRenderDriverInfo(C.int(index), (*C.SDL_RendererInfo)(tmp_info)))
    tmp2_info := fromC2RendererInfo(*(tmp_info)); info = &tmp2_info
    return
}


 // Create a 2D rendering context for a window.
 // 
 // Returns: A valid rendering context or NULL if there was an error.
 // 
 // See also: SDL_CreateSoftwareRenderer()
 // 
 // See also: SDL_GetRendererInfo()
 // 
 // See also: SDL_DestroyRenderer()
 // 
 //   window
 //     The window where rendering is displayed.
 //   
 //   index
 //     The index of the rendering driver to initialize, or -1 to initialize
 //     the first one supporting the requested flags.
 //   
 //   flags
 //     SDL_RendererFlags.
 //   
 // ↪ https://wiki.libsdl.org/SDL_CreateRenderer
func (window *Window) CreateRenderer(index int, flags RendererFlags) (retval *Renderer) {
    retval = (*Renderer)(unsafe.Pointer(C.SDL_CreateRenderer((*C.SDL_Window)(window), C.int(index), C.Uint32(flags))))
    return
}

 // Create a 2D software rendering context for a surface.
 // 
 // Returns: A valid rendering context or NULL if there was an error.
 // 
 // See also: SDL_CreateRenderer()
 // 
 // See also: SDL_DestroyRenderer()
 // 
 //   surface
 //     The surface where rendering is done.
 //   
 // ↪ https://wiki.libsdl.org/SDL_CreateSoftwareRenderer
func (surface *Surface) CreateSoftwareRenderer() (retval *Renderer) {
    retval = (*Renderer)(unsafe.Pointer(C.SDL_CreateSoftwareRenderer((*C.SDL_Surface)(surface))))
    return
}

 // Get the renderer associated with a window.
 // ↪ https://wiki.libsdl.org/SDL_GetRenderer
func (window *Window) GetRenderer() (retval *Renderer) {
    retval = (*Renderer)(unsafe.Pointer(C.SDL_GetRenderer((*C.SDL_Window)(window))))
    return
}

 // Get information about a rendering context.
 // ↪ https://wiki.libsdl.org/SDL_GetRendererInfo
func (renderer *Renderer) GetInfo() (retval int, info *RendererInfo) {
    tmp_info := new(C.SDL_RendererInfo)
    retval = int(C.SDL_GetRendererInfo((*C.SDL_Renderer)(renderer), (*C.SDL_RendererInfo)(tmp_info)))
    tmp2_info := fromC2RendererInfo(*(tmp_info)); info = &tmp2_info
    return
}

 // Get the output size in pixels of a rendering context.
 // ↪ https://wiki.libsdl.org/SDL_GetRendererOutputSize
func (renderer *Renderer) GetOutputSize() (retval int, w int, h int) {
    tmp_w := new(C.int)
    tmp_h := new(C.int)
    retval = int(C.SDL_GetRendererOutputSize((*C.SDL_Renderer)(renderer), (*C.int)(tmp_w), (*C.int)(tmp_h)))
    w = deref_int_ptr(tmp_w)
    h = deref_int_ptr(tmp_h)
    return
}

 // Create a texture for a rendering context.
 // 
 // Returns: The created texture is returned, or NULL if no rendering
 // context was active, the format was unsupported, or the width or height
 // were out of range.
 // 
 // See also: SDL_QueryTexture()
 // 
 // See also: SDL_UpdateTexture()
 // 
 // See also: SDL_DestroyTexture()
 // 
 //   renderer
 //     The renderer.
 //   
 //   format
 //     The format of the texture.
 //   
 //   access
 //     One of the enumerated values in SDL_TextureAccess.
 //   
 //   w
 //     The width of the texture in pixels.
 //   
 //   h
 //     The height of the texture in pixels.
 //   
 // ↪ https://wiki.libsdl.org/SDL_CreateTexture
func (renderer *Renderer) CreateTexture(format uint32, access int, w int, h int) (retval *Texture) {
    retval = (*Texture)(unsafe.Pointer(C.SDL_CreateTexture((*C.SDL_Renderer)(renderer), C.Uint32(format), C.int(access), C.int(w), C.int(h))))
    return
}

 // Create a texture from an existing surface.
 // 
 // Returns: The created texture is returned, or NULL on error.
 // 
 // Note: The surface is not modified or freed by this function.
 // 
 // See also: SDL_QueryTexture()
 // 
 // See also: SDL_DestroyTexture()
 // 
 //   renderer
 //     The renderer.
 //   
 //   surface
 //     The surface containing pixel data used to fill the texture.
 //   
 // ↪ https://wiki.libsdl.org/SDL_CreateTextureFromSurface
func (renderer *Renderer) CreateTextureFromSurface(surface *Surface) (retval *Texture) {
    retval = (*Texture)(unsafe.Pointer(C.SDL_CreateTextureFromSurface((*C.SDL_Renderer)(renderer), (*C.SDL_Surface)(surface))))
    return
}

 // Query the attributes of a texture.
 // 
 // Returns: 0 on success, or -1 if the texture is not valid.
 // 
 //   texture
 //     A texture to be queried.
 //   
 //   format
 //     A pointer filled in with the raw format of the texture. The actual
 //     format may differ, but pixel transfers will use this format.
 //   
 //   access
 //     A pointer filled in with the actual access to the texture.
 //   
 //   w
 //     A pointer filled in with the width of the texture in pixels.
 //   
 //   h
 //     A pointer filled in with the height of the texture in pixels.
 //   
 // ↪ https://wiki.libsdl.org/SDL_QueryTexture
func (texture *Texture) Query() (retval int, format uint32, access int, w int, h int) {
    tmp_format := new(C.Uint32)
    tmp_access := new(C.int)
    tmp_w := new(C.int)
    tmp_h := new(C.int)
    retval = int(C.SDL_QueryTexture((*C.SDL_Texture)(texture), (*C.Uint32)(tmp_format), (*C.int)(tmp_access), (*C.int)(tmp_w), (*C.int)(tmp_h)))
    format = deref_uint32_ptr(tmp_format)
    access = deref_int_ptr(tmp_access)
    w = deref_int_ptr(tmp_w)
    h = deref_int_ptr(tmp_h)
    return
}

 // Set an additional color value used in render copy operations.
 // 
 // Returns: 0 on success, or -1 if the texture is not valid or color
 // modulation is not supported.
 // 
 // See also: SDL_GetTextureColorMod()
 // 
 //   texture
 //     The texture to update.
 //   
 //   r
 //     The red color value multiplied into copy operations.
 //   
 //   g
 //     The green color value multiplied into copy operations.
 //   
 //   b
 //     The blue color value multiplied into copy operations.
 //   
 // ↪ https://wiki.libsdl.org/SDL_SetTextureColorMod
func (texture *Texture) SetColorMod(r uint8, g uint8, b uint8) (retval int) {
    retval = int(C.SDL_SetTextureColorMod((*C.SDL_Texture)(texture), C.Uint8(r), C.Uint8(g), C.Uint8(b)))
    return
}

 // Get the additional color value used in render copy operations.
 // 
 // Returns: 0 on success, or -1 if the texture is not valid.
 // 
 // See also: SDL_SetTextureColorMod()
 // 
 //   texture
 //     The texture to query.
 //   
 //   r
 //     A pointer filled in with the current red color value.
 //   
 //   g
 //     A pointer filled in with the current green color value.
 //   
 //   b
 //     A pointer filled in with the current blue color value.
 //   
 // ↪ https://wiki.libsdl.org/SDL_GetTextureColorMod
func (texture *Texture) GetColorMod() (retval int, r byte, g byte, b byte) {
    tmp_r := new(C.Uint8)
    tmp_g := new(C.Uint8)
    tmp_b := new(C.Uint8)
    retval = int(C.SDL_GetTextureColorMod((*C.SDL_Texture)(texture), (*C.Uint8)(tmp_r), (*C.Uint8)(tmp_g), (*C.Uint8)(tmp_b)))
    r = deref_byte_ptr(tmp_r)
    g = deref_byte_ptr(tmp_g)
    b = deref_byte_ptr(tmp_b)
    return
}

 // Set an additional alpha value used in render copy operations.
 // 
 // Returns: 0 on success, or -1 if the texture is not valid or alpha
 // modulation is not supported.
 // 
 // See also: SDL_GetTextureAlphaMod()
 // 
 //   texture
 //     The texture to update.
 //   
 //   alpha
 //     The alpha value multiplied into copy operations.
 //   
 // ↪ https://wiki.libsdl.org/SDL_SetTextureAlphaMod
func (texture *Texture) SetAlphaMod(alpha uint8) (retval int) {
    retval = int(C.SDL_SetTextureAlphaMod((*C.SDL_Texture)(texture), C.Uint8(alpha)))
    return
}

 // Get the additional alpha value used in render copy operations.
 // 
 // Returns: 0 on success, or -1 if the texture is not valid.
 // 
 // See also: SDL_SetTextureAlphaMod()
 // 
 //   texture
 //     The texture to query.
 //   
 //   alpha
 //     A pointer filled in with the current alpha value.
 //   
 // ↪ https://wiki.libsdl.org/SDL_GetTextureAlphaMod
func (texture *Texture) GetAlphaMod() (retval int, alpha byte) {
    tmp_alpha := new(C.Uint8)
    retval = int(C.SDL_GetTextureAlphaMod((*C.SDL_Texture)(texture), (*C.Uint8)(tmp_alpha)))
    alpha = deref_byte_ptr(tmp_alpha)
    return
}

 // Set the blend mode used for texture copy operations.
 // 
 // Returns: 0 on success, or -1 if the texture is not valid or the blend
 // mode is not supported.
 // 
 // Note: If the blend mode is not supported, the closest supported mode
 // is chosen.
 // 
 // See also: SDL_GetTextureBlendMode()
 // 
 //   texture
 //     The texture to update.
 //   
 //   blendMode
 //     SDL_BlendMode to use for texture blending.
 //   
 // ↪ https://wiki.libsdl.org/SDL_SetTextureBlendMode
func (texture *Texture) SetBlendMode(blendMode BlendMode) (retval int) {
    retval = int(C.SDL_SetTextureBlendMode((*C.SDL_Texture)(texture), C.SDL_BlendMode(blendMode)))
    return
}

 // Get the blend mode used for texture copy operations.
 // 
 // Returns: 0 on success, or -1 if the texture is not valid.
 // 
 // See also: SDL_SetTextureBlendMode()
 // 
 //   texture
 //     The texture to query.
 //   
 //   blendMode
 //     A pointer filled in with the current blend mode.
 //   
 // ↪ https://wiki.libsdl.org/SDL_GetTextureBlendMode
func (texture *Texture) GetBlendMode() (retval int, blendMode *BlendMode) {
    tmp_blendMode := new(C.SDL_BlendMode)
    retval = int(C.SDL_GetTextureBlendMode((*C.SDL_Texture)(texture), (*C.SDL_BlendMode)(tmp_blendMode)))
    blendMode = (*BlendMode)(unsafe.Pointer(tmp_blendMode))
    return
}

 // Update the given texture rectangle with new pixel data.
 // 
 // Returns: 0 on success, or -1 if the texture is not valid.
 // 
 // Note: This is a fairly slow function.
 // 
 //   texture
 //     The texture to update
 //   
 //   rect
 //     A pointer to the rectangle of pixels to update, or NULL to update the
 //     entire texture.
 //   
 //   pixels
 //     The raw pixel data.
 //   
 //   pitch
 //     The number of bytes in a row of pixel data, including padding between
 //     lines.
 //   
 // ↪ https://wiki.libsdl.org/SDL_UpdateTexture
func (texture *Texture) Update(rect *Rect, pixels []byte, pitch int) (retval int) {
    var tmp_rect *C.SDL_Rect; if rect != nil { x := toCFromRect(*rect); tmp_rect = &x }
    checkParametersForSDL_UpdateTexture(texture, rect, pixels, pitch)
    var tmp_pixels unsafe.Pointer
    if len(pixels) > 0 {
        tmp_pixels = (unsafe.Pointer)(unsafe.Pointer(&(pixels[0])))
    }
    retval = int(C.SDL_UpdateTexture((*C.SDL_Texture)(texture), (*C.SDL_Rect)(tmp_rect), (tmp_pixels), C.int(pitch)))
    return
}

 // Update a rectangle within a planar YV12 or IYUV texture with new pixel
 // data.
 // 
 // Returns: 0 on success, or -1 if the texture is not valid.
 // 
 // Note: You can use SDL_UpdateTexture() as long as your pixel data is a
 // contiguous block of Y and U/V planes in the proper order, but this
 // function is available if your pixel data is not contiguous.
 // 
 //   texture
 //     The texture to update
 //   
 //   rect
 //     A pointer to the rectangle of pixels to update, or NULL to update the
 //     entire texture.
 //   
 //   Yplane
 //     The raw pixel data for the Y plane.
 //   
 //   Ypitch
 //     The number of bytes between rows of pixel data for the Y plane.
 //   
 //   Uplane
 //     The raw pixel data for the U plane.
 //   
 //   Upitch
 //     The number of bytes between rows of pixel data for the U plane.
 //   
 //   Vplane
 //     The raw pixel data for the V plane.
 //   
 //   Vpitch
 //     The number of bytes between rows of pixel data for the V plane.
 //   
 // ↪ https://wiki.libsdl.org/SDL_UpdateYUVTexture
func (texture *Texture) UpdateYUV(rect *Rect, Yplane []byte, Ypitch int, Uplane []byte, Upitch int, Vplane []byte, Vpitch int) (retval int) {
    var tmp_rect *C.SDL_Rect; if rect != nil { x := toCFromRect(*rect); tmp_rect = &x }
    checkParametersForSDL_UpdateYUVTexture(texture, rect, Yplane, Ypitch, Uplane, Upitch, Vplane, Vpitch)
    var tmp_Yplane *C.Uint8
    if len(Yplane) > 0 {
        tmp_Yplane = (*C.Uint8)(unsafe.Pointer(&(Yplane[0])))
    }
    var tmp_Uplane *C.Uint8
    if len(Uplane) > 0 {
        tmp_Uplane = (*C.Uint8)(unsafe.Pointer(&(Uplane[0])))
    }
    var tmp_Vplane *C.Uint8
    if len(Vplane) > 0 {
        tmp_Vplane = (*C.Uint8)(unsafe.Pointer(&(Vplane[0])))
    }
    retval = int(C.SDL_UpdateYUVTexture((*C.SDL_Texture)(texture), (*C.SDL_Rect)(tmp_rect), (tmp_Yplane), C.int(Ypitch), (tmp_Uplane), C.int(Upitch), (tmp_Vplane), C.int(Vpitch)))
    return
}


 // Unlock a texture, uploading the changes to video memory, if needed.
 // 
 // See also: SDL_LockTexture()
 // 
 // ↪ https://wiki.libsdl.org/SDL_UnlockTexture
func (texture *Texture) Unlock() {
    C.SDL_UnlockTexture((*C.SDL_Texture)(texture))
}

 // Determines whether a window supports the use of render targets.
 // 
 // Returns: SDL_TRUE if supported, SDL_FALSE if not.
 // 
 //   renderer
 //     The renderer that will be checked
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderTargetSupported
func (renderer *Renderer) TargetSupported() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_RenderTargetSupported((*C.SDL_Renderer)(renderer)))
    return
}

 // Set a texture as the current rendering target.
 // 
 // Returns: 0 on success, or -1 on error
 // 
 // See also: SDL_GetRenderTarget()
 // 
 //   renderer
 //     The renderer.
 //   
 //   texture
 //     The targeted texture, which must be created with the
 //     SDL_TEXTUREACCESS_TARGET flag, or NULL for the default render target
 //   
 // ↪ https://wiki.libsdl.org/SDL_SetRenderTarget
func (renderer *Renderer) SetTarget(texture *Texture) (retval int) {
    retval = int(C.SDL_SetRenderTarget((*C.SDL_Renderer)(renderer), (*C.SDL_Texture)(texture)))
    return
}

 // Get the current render target or NULL for the default render target.
 // 
 // Returns: The current render target
 // 
 // See also: SDL_SetRenderTarget()
 // 
 // ↪ https://wiki.libsdl.org/SDL_GetRenderTarget
func (renderer *Renderer) GetTarget() (retval *Texture) {
    retval = (*Texture)(unsafe.Pointer(C.SDL_GetRenderTarget((*C.SDL_Renderer)(renderer))))
    return
}

 // Set device independent resolution for rendering.
 // 
 //   renderer
 //     The renderer for which resolution should be set.
 //   
 //   w
 //     The width of the logical resolution
 //   
 //   h
 //     The height of the logical resolution
 //   
 // This function uses the viewport and scaling functionality to allow a
 // fixed logical resolution for rendering, regardless of the actual
 // output resolution. If the actual output resolution doesn't have the
 // same aspect ratio the output rendering will be centered within the
 // output display.
 // 
 // If the output display is a window, mouse events in the window will be
 // filtered and scaled so they seem to arrive within the logical
 // resolution.
 // 
 // Note: If this function results in scaling or subpixel drawing by the
 // rendering backend, it will be handled using the appropriate quality
 // hints.
 // 
 // See also: SDL_RenderGetLogicalSize()
 // 
 // See also: SDL_RenderSetScale()
 // 
 // See also: SDL_RenderSetViewport()
 // 
 // ↪ https://wiki.libsdl.org/SDL_RenderSetLogicalSize
func (renderer *Renderer) SetLogicalSize(w int, h int) (retval int) {
    retval = int(C.SDL_RenderSetLogicalSize((*C.SDL_Renderer)(renderer), C.int(w), C.int(h)))
    return
}

 // Get device independent resolution for rendering.
 // 
 // See also: SDL_RenderSetLogicalSize()
 // 
 //   renderer
 //     The renderer from which resolution should be queried.
 //   
 //   w
 //     A pointer filled with the width of the logical resolution
 //   
 //   h
 //     A pointer filled with the height of the logical resolution
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderGetLogicalSize
func (renderer *Renderer) GetLogicalSize() (w int, h int) {
    tmp_w := new(C.int)
    tmp_h := new(C.int)
    C.SDL_RenderGetLogicalSize((*C.SDL_Renderer)(renderer), (*C.int)(tmp_w), (*C.int)(tmp_h))
    w = deref_int_ptr(tmp_w)
    h = deref_int_ptr(tmp_h)
    return
}

 // Set the drawing area for rendering on the current target.
 // 
 //   renderer
 //     The renderer for which the drawing area should be set.
 //   
 //   rect
 //     The rectangle representing the drawing area, or NULL to set the
 //     viewport to the entire target.
 //   
 // The x,y of the viewport rect represents the origin for rendering.
 // 
 // Returns: 0 on success, or -1 on error
 // 
 // Note: If the window associated with the renderer is resized, the
 // viewport is automatically reset.
 // 
 // See also: SDL_RenderGetViewport()
 // 
 // See also: SDL_RenderSetLogicalSize()
 // 
 // ↪ https://wiki.libsdl.org/SDL_RenderSetViewport
func (renderer *Renderer) SetViewport(rect Rect) (retval int) {
    tmp_rect := toCFromRect(rect)
    retval = int(C.SDL_RenderSetViewport((*C.SDL_Renderer)(renderer), (*C.SDL_Rect)(&tmp_rect)))
    return
}

 // Get the drawing area for the current target.
 // 
 // See also: SDL_RenderSetViewport()
 // 
 // ↪ https://wiki.libsdl.org/SDL_RenderGetViewport
func (renderer *Renderer) GetViewport() (rect Rect) {
    tmp_rect := new(C.SDL_Rect)
    C.SDL_RenderGetViewport((*C.SDL_Renderer)(renderer), (*C.SDL_Rect)(tmp_rect))
    rect = fromC2Rect(*(tmp_rect))
    return
}

 // Set the clip rectangle for the current target.
 // 
 // Returns: 0 on success, or -1 on error
 // 
 // See also: SDL_RenderGetClipRect()
 // 
 //   renderer
 //     The renderer for which clip rectangle should be set.
 //   
 //   rect
 //     A pointer to the rectangle to set as the clip rectangle, or NULL to
 //     disable clipping.
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderSetClipRect
func (renderer *Renderer) SetClipRect(rect Rect) (retval int) {
    tmp_rect := toCFromRect(rect)
    retval = int(C.SDL_RenderSetClipRect((*C.SDL_Renderer)(renderer), (*C.SDL_Rect)(&tmp_rect)))
    return
}

 // Get the clip rectangle for the current target.
 // 
 // See also: SDL_RenderSetClipRect()
 // 
 //   renderer
 //     The renderer from which clip rectangle should be queried.
 //   
 //   rect
 //     A pointer filled in with the current clip rectangle, or an empty
 //     rectangle if clipping is disabled.
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderGetClipRect
func (renderer *Renderer) GetClipRect() (rect Rect) {
    tmp_rect := new(C.SDL_Rect)
    C.SDL_RenderGetClipRect((*C.SDL_Renderer)(renderer), (*C.SDL_Rect)(tmp_rect))
    rect = fromC2Rect(*(tmp_rect))
    return
}

 // Get whether clipping is enabled on the given renderer.
 // 
 // See also: SDL_RenderGetClipRect()
 // 
 //   renderer
 //     The renderer from which clip state should be queried.
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderIsClipEnabled
func (renderer *Renderer) IsClipEnabled() (retval bool) {
    retval = C.SDL_TRUE==(C.SDL_RenderIsClipEnabled((*C.SDL_Renderer)(renderer)))
    return
}

 // Set the drawing scale for rendering on the current target.
 // 
 //   renderer
 //     The renderer for which the drawing scale should be set.
 //   
 //   scaleX
 //     The horizontal scaling factor
 //   
 //   scaleY
 //     The vertical scaling factor
 //   
 // The drawing coordinates are scaled by the x/y scaling factors before
 // they are used by the renderer. This allows resolution independent
 // drawing with a single coordinate system.
 // 
 // Note: If this results in scaling or subpixel drawing by the rendering
 // backend, it will be handled using the appropriate quality hints. For
 // best results use integer scaling factors.
 // 
 // See also: SDL_RenderGetScale()
 // 
 // See also: SDL_RenderSetLogicalSize()
 // 
 // ↪ https://wiki.libsdl.org/SDL_RenderSetScale
func (renderer *Renderer) SetScale(scaleX float32, scaleY float32) (retval int) {
    retval = int(C.SDL_RenderSetScale((*C.SDL_Renderer)(renderer), C.float(scaleX), C.float(scaleY)))
    return
}

 // Get the drawing scale for the current target.
 // 
 // See also: SDL_RenderSetScale()
 // 
 //   renderer
 //     The renderer from which drawing scale should be queried.
 //   
 //   scaleX
 //     A pointer filled in with the horizontal scaling factor
 //   
 //   scaleY
 //     A pointer filled in with the vertical scaling factor
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderGetScale
func (renderer *Renderer) GetScale() (scaleX float32, scaleY float32) {
    tmp_scaleX := new(C.float)
    tmp_scaleY := new(C.float)
    C.SDL_RenderGetScale((*C.SDL_Renderer)(renderer), (*C.float)(tmp_scaleX), (*C.float)(tmp_scaleY))
    scaleX = deref_float32_ptr(tmp_scaleX)
    scaleY = deref_float32_ptr(tmp_scaleY)
    return
}

 // Set the color used for drawing operations (Rect, Line and Clear).
 // 
 // Returns: 0 on success, or -1 on error
 // 
 //   renderer
 //     The renderer for which drawing color should be set.
 //   
 //   r
 //     The red value used to draw on the rendering target.
 //   
 //   g
 //     The green value used to draw on the rendering target.
 //   
 //   b
 //     The blue value used to draw on the rendering target.
 //   
 //   a
 //     The alpha value used to draw on the rendering target, usually
 //     SDL_ALPHA_OPAQUE (255).
 //   
 // ↪ https://wiki.libsdl.org/SDL_SetRenderDrawColor
func (renderer *Renderer) SetDrawColor(r uint8, g uint8, b uint8, a uint8) (retval int) {
    retval = int(C.SDL_SetRenderDrawColor((*C.SDL_Renderer)(renderer), C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)))
    return
}

 // Get the color used for drawing operations (Rect, Line and Clear).
 // 
 // Returns: 0 on success, or -1 on error
 // 
 //   renderer
 //     The renderer from which drawing color should be queried.
 //   
 //   r
 //     A pointer to the red value used to draw on the rendering target.
 //   
 //   g
 //     A pointer to the green value used to draw on the rendering target.
 //   
 //   b
 //     A pointer to the blue value used to draw on the rendering target.
 //   
 //   a
 //     A pointer to the alpha value used to draw on the rendering target,
 //     usually SDL_ALPHA_OPAQUE (255).
 //   
 // ↪ https://wiki.libsdl.org/SDL_GetRenderDrawColor
func (renderer *Renderer) GetDrawColor() (retval int, r byte, g byte, b byte, a byte) {
    tmp_r := new(C.Uint8)
    tmp_g := new(C.Uint8)
    tmp_b := new(C.Uint8)
    tmp_a := new(C.Uint8)
    retval = int(C.SDL_GetRenderDrawColor((*C.SDL_Renderer)(renderer), (*C.Uint8)(tmp_r), (*C.Uint8)(tmp_g), (*C.Uint8)(tmp_b), (*C.Uint8)(tmp_a)))
    r = deref_byte_ptr(tmp_r)
    g = deref_byte_ptr(tmp_g)
    b = deref_byte_ptr(tmp_b)
    a = deref_byte_ptr(tmp_a)
    return
}

 // Set the blend mode used for drawing operations (Fill and Line).
 // 
 // Returns: 0 on success, or -1 on error
 // 
 // Note: If the blend mode is not supported, the closest supported mode
 // is chosen.
 // 
 // See also: SDL_GetRenderDrawBlendMode()
 // 
 //   renderer
 //     The renderer for which blend mode should be set.
 //   
 //   blendMode
 //     SDL_BlendMode to use for blending.
 //   
 // ↪ https://wiki.libsdl.org/SDL_SetRenderDrawBlendMode
func (renderer *Renderer) SetDrawBlendMode(blendMode BlendMode) (retval int) {
    retval = int(C.SDL_SetRenderDrawBlendMode((*C.SDL_Renderer)(renderer), C.SDL_BlendMode(blendMode)))
    return
}

 // Get the blend mode used for drawing operations.
 // 
 // Returns: 0 on success, or -1 on error
 // 
 // See also: SDL_SetRenderDrawBlendMode()
 // 
 //   renderer
 //     The renderer from which blend mode should be queried.
 //   
 //   blendMode
 //     A pointer filled in with the current blend mode.
 //   
 // ↪ https://wiki.libsdl.org/SDL_GetRenderDrawBlendMode
func (renderer *Renderer) GetDrawBlendMode() (retval int, blendMode *BlendMode) {
    tmp_blendMode := new(C.SDL_BlendMode)
    retval = int(C.SDL_GetRenderDrawBlendMode((*C.SDL_Renderer)(renderer), (*C.SDL_BlendMode)(tmp_blendMode)))
    blendMode = (*BlendMode)(unsafe.Pointer(tmp_blendMode))
    return
}

 // Clear the current rendering target with the drawing color.
 // 
 // This function clears the entire rendering target, ignoring the
 // viewport.
 // 
 // Returns: 0 on success, or -1 on error
 // 
 // ↪ https://wiki.libsdl.org/SDL_RenderClear
func (renderer *Renderer) Clear() (retval int) {
    retval = int(C.SDL_RenderClear((*C.SDL_Renderer)(renderer)))
    return
}

 // Draw a point on the current rendering target.
 // 
 // Returns: 0 on success, or -1 on error
 // 
 //   renderer
 //     The renderer which should draw a point.
 //   
 //   x
 //     The x coordinate of the point.
 //   
 //   y
 //     The y coordinate of the point.
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderDrawPoint
func (renderer *Renderer) DrawPoint(x int, y int) (retval int) {
    retval = int(C.SDL_RenderDrawPoint((*C.SDL_Renderer)(renderer), C.int(x), C.int(y)))
    return
}

 // Draw multiple points on the current rendering target.
 // 
 // Returns: 0 on success, or -1 on error
 // 
 //   renderer
 //     The renderer which should draw multiple points.
 //   
 //   points
 //     The points to draw
 //   
 //   count
 //     The number of points to draw
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderDrawPoints
func (renderer *Renderer) DrawPoints(points []Point) (retval int) {
    var tmp_points *C.SDL_Point
    if len(points) > 0 {
        sl_tmp_points := make([]C.SDL_Point, len(points))
        for i := range points {
            sl_tmp_points[i] = toCFromPoint(points[i])
        }
        tmp_points = &(sl_tmp_points[0])
    }
    tmp_count := len(points)
    retval = int(C.SDL_RenderDrawPoints((*C.SDL_Renderer)(renderer), (tmp_points), C.int(tmp_count)))
    return
}

 // Draw a line on the current rendering target.
 // 
 // Returns: 0 on success, or -1 on error
 // 
 //   renderer
 //     The renderer which should draw a line.
 //   
 //   x1
 //     The x coordinate of the start point.
 //   
 //   y1
 //     The y coordinate of the start point.
 //   
 //   x2
 //     The x coordinate of the end point.
 //   
 //   y2
 //     The y coordinate of the end point.
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderDrawLine
func (renderer *Renderer) DrawLine(x1 int, y1 int, x2 int, y2 int) (retval int) {
    retval = int(C.SDL_RenderDrawLine((*C.SDL_Renderer)(renderer), C.int(x1), C.int(y1), C.int(x2), C.int(y2)))
    return
}

 // Draw a series of connected lines on the current rendering target.
 // 
 // Returns: 0 on success, or -1 on error
 // 
 //   renderer
 //     The renderer which should draw multiple lines.
 //   
 //   points
 //     The points along the lines
 //   
 //   count
 //     The number of points, drawing count-1 lines
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderDrawLines
func (renderer *Renderer) DrawLines(points []Point) (retval int) {
    var tmp_points *C.SDL_Point
    if len(points) > 0 {
        sl_tmp_points := make([]C.SDL_Point, len(points))
        for i := range points {
            sl_tmp_points[i] = toCFromPoint(points[i])
        }
        tmp_points = &(sl_tmp_points[0])
    }
    tmp_count := len(points)
    retval = int(C.SDL_RenderDrawLines((*C.SDL_Renderer)(renderer), (tmp_points), C.int(tmp_count)))
    return
}

 // Draw a rectangle on the current rendering target.
 // 
 // Returns: 0 on success, or -1 on error
 // 
 //   renderer
 //     The renderer which should draw a rectangle.
 //   
 //   rect
 //     A pointer to the destination rectangle, or NULL to outline the entire
 //     rendering target.
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderDrawRect
func (renderer *Renderer) DrawRect(rect Rect) (retval int) {
    tmp_rect := toCFromRect(rect)
    retval = int(C.SDL_RenderDrawRect((*C.SDL_Renderer)(renderer), (*C.SDL_Rect)(&tmp_rect)))
    return
}

 // Draw some number of rectangles on the current rendering target.
 // 
 // Returns: 0 on success, or -1 on error
 // 
 //   renderer
 //     The renderer which should draw multiple rectangles.
 //   
 //   rects
 //     A pointer to an array of destination rectangles.
 //   
 //   count
 //     The number of rectangles.
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderDrawRects
func (renderer *Renderer) DrawRects(rects []Rect) (retval int) {
    var tmp_rects *C.SDL_Rect
    if len(rects) > 0 {
        sl_tmp_rects := make([]C.SDL_Rect, len(rects))
        for i := range rects {
            sl_tmp_rects[i] = toCFromRect(rects[i])
        }
        tmp_rects = &(sl_tmp_rects[0])
    }
    tmp_count := len(rects)
    retval = int(C.SDL_RenderDrawRects((*C.SDL_Renderer)(renderer), (tmp_rects), C.int(tmp_count)))
    return
}

 // Fill a rectangle on the current rendering target with the drawing
 // color.
 // 
 // Returns: 0 on success, or -1 on error
 // 
 //   renderer
 //     The renderer which should fill a rectangle.
 //   
 //   rect
 //     A pointer to the destination rectangle, or NULL for the entire
 //     rendering target.
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderFillRect
func (renderer *Renderer) FillRect(rect Rect) (retval int) {
    tmp_rect := toCFromRect(rect)
    retval = int(C.SDL_RenderFillRect((*C.SDL_Renderer)(renderer), (*C.SDL_Rect)(&tmp_rect)))
    return
}

 // Fill some number of rectangles on the current rendering target with
 // the drawing color.
 // 
 // Returns: 0 on success, or -1 on error
 // 
 //   renderer
 //     The renderer which should fill multiple rectangles.
 //   
 //   rects
 //     A pointer to an array of destination rectangles.
 //   
 //   count
 //     The number of rectangles.
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderFillRects
func (renderer *Renderer) FillRects(rects []Rect) (retval int) {
    var tmp_rects *C.SDL_Rect
    if len(rects) > 0 {
        sl_tmp_rects := make([]C.SDL_Rect, len(rects))
        for i := range rects {
            sl_tmp_rects[i] = toCFromRect(rects[i])
        }
        tmp_rects = &(sl_tmp_rects[0])
    }
    tmp_count := len(rects)
    retval = int(C.SDL_RenderFillRects((*C.SDL_Renderer)(renderer), (tmp_rects), C.int(tmp_count)))
    return
}

 // Copy a portion of the texture to the current rendering target.
 // 
 // Returns: 0 on success, or -1 on error
 // 
 //   renderer
 //     The renderer which should copy parts of a texture.
 //   
 //   texture
 //     The source texture.
 //   
 //   srcrect
 //     A pointer to the source rectangle, or NULL for the entire texture.
 //   
 //   dstrect
 //     A pointer to the destination rectangle, or NULL for the entire
 //     rendering target.
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderCopy
func (renderer *Renderer) Copy(texture *Texture, srcrect *Rect, dstrect *Rect) (retval int) {
    var tmp_srcrect *C.SDL_Rect; if srcrect != nil { x := toCFromRect(*srcrect); tmp_srcrect = &x }
    var tmp_dstrect *C.SDL_Rect; if dstrect != nil { x := toCFromRect(*dstrect); tmp_dstrect = &x }
    retval = int(C.SDL_RenderCopy((*C.SDL_Renderer)(renderer), (*C.SDL_Texture)(texture), (*C.SDL_Rect)(tmp_srcrect), (*C.SDL_Rect)(tmp_dstrect)))
    return
}

 // Copy a portion of the source texture to the current rendering target,
 // rotating it by angle around the given center.
 // 
 // Returns: 0 on success, or -1 on error
 // 
 //   renderer
 //     The renderer which should copy parts of a texture.
 //   
 //   texture
 //     The source texture.
 //   
 //   srcrect
 //     A pointer to the source rectangle, or NULL for the entire texture.
 //   
 //   dstrect
 //     A pointer to the destination rectangle, or NULL for the entire
 //     rendering target.
 //   
 //   angle
 //     An angle in degrees that indicates the rotation that will be applied
 //     to dstrect
 //   
 //   center
 //     A pointer to a point indicating the point around which dstrect will be
 //     rotated (if NULL, rotation will be done around dstrect.w/2,
 //     dstrect.h/2).
 //   
 //   flip
 //     An SDL_RendererFlip value stating which flipping actions should be
 //     performed on the texture
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderCopyEx
func (renderer *Renderer) CopyEx(texture *Texture, srcrect Rect, dstrect Rect, angle float64, center Point, flip RendererFlip) (retval int) {
    tmp_srcrect := toCFromRect(srcrect)
    tmp_dstrect := toCFromRect(dstrect)
    tmp_center := toCFromPoint(center)
    retval = int(C.SDL_RenderCopyEx((*C.SDL_Renderer)(renderer), (*C.SDL_Texture)(texture), (*C.SDL_Rect)(&tmp_srcrect), (*C.SDL_Rect)(&tmp_dstrect), C.double(angle), (*C.SDL_Point)(&tmp_center), C.SDL_RendererFlip(flip)))
    return
}

 // Read pixels from the current rendering target.
 // 
 // Returns: 0 on success, or -1 if pixel reading is not supported.
 // 
 // Warning: This is a very slow operation, and should not be used
 // frequently.
 // 
 //   renderer
 //     The renderer from which pixels should be read.
 //   
 //   rect
 //     A pointer to the rectangle to read, or NULL for the entire render
 //     target.
 //   
 //   format
 //     The desired format of the pixel data, or 0 to use the format of the
 //     rendering target
 //   
 //   pixels
 //     A pointer to be filled in with the pixel data
 //   
 //   pitch
 //     The pitch of the pixels parameter.
 //   
 // ↪ https://wiki.libsdl.org/SDL_RenderReadPixels
func (renderer *Renderer) ReadPixels(rect *Rect, format uint32, pixels []byte, pitch int) (retval int) {
    var tmp_rect *C.SDL_Rect; if rect != nil { x := toCFromRect(*rect); tmp_rect = &x }
    checkParametersForSDL_RenderReadPixels(renderer, rect, format, pixels, pitch)
    var tmp_pixels unsafe.Pointer
    if len(pixels) > 0 {
        tmp_pixels = (unsafe.Pointer)(unsafe.Pointer(&(pixels[0])))
    }
    retval = int(C.SDL_RenderReadPixels((*C.SDL_Renderer)(renderer), (*C.SDL_Rect)(tmp_rect), C.Uint32(format), (tmp_pixels), C.int(pitch)))
    return
}

 // Update the screen with rendering performed.
 // ↪ https://wiki.libsdl.org/SDL_RenderPresent
func (renderer *Renderer) Present() {
    C.SDL_RenderPresent((*C.SDL_Renderer)(renderer))
}

 // Destroy the specified texture.
 // 
 // See also: SDL_CreateTexture()
 // 
 // See also: SDL_CreateTextureFromSurface()
 // 
 // ↪ https://wiki.libsdl.org/SDL_DestroyTexture
func (texture *Texture) Destroy() {
    C.SDL_DestroyTexture((*C.SDL_Texture)(texture))
}

 // Destroy the rendering context for a window and free associated
 // textures.
 // 
 // See also: SDL_CreateRenderer()
 // 
 // ↪ https://wiki.libsdl.org/SDL_DestroyRenderer
func (renderer *Renderer) Destroy() {
    C.SDL_DestroyRenderer((*C.SDL_Renderer)(renderer))
}

 // Bind the texture to the current OpenGL/ES/ES2 context for use with
 // OpenGL instructions.
 // 
 // Returns: 0 on success, or -1 if the operation is not supported
 // 
 //   texture
 //     The SDL texture to bind
 //   
 //   texw
 //     A pointer to a float that will be filled with the texture width
 //   
 //   texh
 //     A pointer to a float that will be filled with the texture height
 //   
 // ↪ https://wiki.libsdl.org/SDL_GL_BindTexture
func (texture *Texture) GL_Bind() (retval int, texw float32, texh float32) {
    tmp_texw := new(C.float)
    tmp_texh := new(C.float)
    retval = int(C.SDL_GL_BindTexture((*C.SDL_Texture)(texture), (*C.float)(tmp_texw), (*C.float)(tmp_texh)))
    texw = deref_float32_ptr(tmp_texw)
    texh = deref_float32_ptr(tmp_texh)
    return
}

 // Unbind a texture from the current OpenGL/ES/ES2 context.
 // 
 // Returns: 0 on success, or -1 if the operation is not supported
 // 
 //   texture
 //     The SDL texture to unbind
 //   
 // ↪ https://wiki.libsdl.org/SDL_GL_UnbindTexture
func (texture *Texture) GL_Unbind() (retval int) {
    retval = int(C.SDL_GL_UnbindTexture((*C.SDL_Texture)(texture)))
    return
}
