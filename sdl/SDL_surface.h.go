// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
import "C"
import "unsafe"

 // Header file for SDL_Surface definition and management functions.

 // Surface flags
 // 
 // These are the currently supported flags for the SDL_Surface.
const (
     // Just here for compatibility
    SWSURFACE = C.SDL_SWSURFACE

     // Surface uses preallocated memory
    PREALLOC = C.SDL_PREALLOC

     // Surface is RLE encoded
    RLEACCEL = C.SDL_RLEACCEL

     // Surface is referenced internally
    DONTFREE = C.SDL_DONTFREE
)


 // The type of function used for surface blitting functions.
type Blit C.SDL_blit


 // Allocate and free an RGB surface.
 // 
 // If the depth is 4 or 8 bits, an empty palette is allocated for the
 // surface. If the depth is greater than 8 bits, the pixel format is set
 // using the flags '[RGB]mask'.
 // 
 // If the function runs out of memory, it will return NULL.
 // 
 //   flags
 //     The flags are obsolete and should be set to 0.
 //   
 //   width
 //     The width in pixels of the surface to create.
 //   
 //   height
 //     The height in pixels of the surface to create.
 //   
 //   depth
 //     The depth in bits of the surface to create.
 //   
 //   Rmask
 //     The red mask of the surface to create.
 //   
 //   Gmask
 //     The green mask of the surface to create.
 //   
 //   Bmask
 //     The blue mask of the surface to create.
 //   
 //   Amask
 //     The alpha mask of the surface to create.
 //   
func CreateRGBSurface(flags uint32, width int, height int, depth int, Rmask uint32, Gmask uint32, Bmask uint32, Amask uint32) (retval *Surface) {
    retval = (*Surface)(unsafe.Pointer(C.SDL_CreateRGBSurface(C.Uint32(flags), C.int(width), C.int(height), C.int(depth), C.Uint32(Rmask), C.Uint32(Gmask), C.Uint32(Bmask), C.Uint32(Amask))))
    return
}


func (surface *Surface) Free() {
    C.SDL_FreeSurface((*C.SDL_Surface)(surface))
}

 // Set the palette used by a surface.
 // 
 // Returns: 0, or -1 if the surface format doesn't use a palette.
 // 
 // Note: A single palette can be shared with many surfaces.
 // 
func (surface *Surface) SetPalette(palette *Palette) (retval int) {
    retval = int(C.SDL_SetSurfacePalette((*C.SDL_Surface)(surface), (*C.SDL_Palette)(palette)))
    return
}

 // Sets up a surface for directly accessing the pixels.
 // 
 // Between calls to SDL_LockSurface() / SDL_UnlockSurface(), you can
 // write to and read from surface->pixels, using the pixel format stored
 // in surface->format. Once you are done accessing the surface, you
 // should use SDL_UnlockSurface() to release it.
 // 
 // Not all surfaces require locking. If SDL_MUSTLOCK(surface) evaluates
 // to 0, then you can read and write to the surface at any time, and the
 // pixel format of the surface will not change.
 // 
 // No operating system or library calls should be made between
 // lock/unlock pairs, as critical system locks may be held during this
 // time.
 // 
 // SDL_LockSurface() returns 0, or -1 if the surface couldn't be locked.
 // 
 // See also: SDL_UnlockSurface()
 // 
func (surface *Surface) Lock() (retval int) {
    retval = int(C.SDL_LockSurface((*C.SDL_Surface)(surface)))
    return
}

 // See also: SDL_LockSurface()
 // 
func (surface *Surface) Unlock() {
    C.SDL_UnlockSurface((*C.SDL_Surface)(surface))
}

 // Load a surface from a seekable SDL data stream (memory or file).
 // 
 // If freesrc is non-zero, the stream will be closed after being read.
 // 
 // The new surface should be freed with SDL_FreeSurface().
 // 
 // Returns: the new surface, or NULL if there was an error.
 // 
func (src *RWops) LoadBMP_RW(freesrc int) (retval *Surface) {
    retval = (*Surface)(unsafe.Pointer(C.SDL_LoadBMP_RW((*C.SDL_RWops)(src), C.int(freesrc))))
    return
}

 // Save a surface to a seekable SDL data stream (memory or file).
 // 
 // If freedst is non-zero, the stream will be closed after being written.
 // 
 // Returns: 0 if successful or -1 if there was an error.
 // 
func (dst *RWops) SaveBMP_RW(surface *Surface, freedst int) (retval int) {
    retval = int(C.SDL_SaveBMP_RW((*C.SDL_Surface)(surface), (*C.SDL_RWops)(dst), C.int(freedst)))
    return
}

 // Sets the RLE acceleration hint for a surface.
 // 
 // Returns: 0 on success, or -1 if the surface is not valid
 // 
 // Note: If RLE is enabled, colorkey and alpha blending blits are much
 // faster, but the surface must be locked before directly accessing the
 // pixels.
 // 
func (surface *Surface) SetRLE(flag int) (retval int) {
    retval = int(C.SDL_SetSurfaceRLE((*C.SDL_Surface)(surface), C.int(flag)))
    return
}

 // Sets the color key (transparent pixel) in a blittable surface.
 // 
 // Returns: 0 on success, or -1 if the surface is not valid
 // 
 //   surface
 //     The surface to update
 //   
 //   flag
 //     Non-zero to enable colorkey and 0 to disable colorkey
 //   
 //   key
 //     The transparent pixel in the native surface format
 //   
 // You can pass SDL_RLEACCEL to enable RLE accelerated blits.
func (surface *Surface) SetColorKey(flag int, key uint32) (retval int) {
    retval = int(C.SDL_SetColorKey((*C.SDL_Surface)(surface), C.int(flag), C.Uint32(key)))
    return
}

 // Gets the color key (transparent pixel) in a blittable surface.
 // 
 // Returns: 0 on success, or -1 if the surface is not valid or colorkey
 // is not enabled.
 // 
 //   surface
 //     The surface to update
 //   
 //   key
 //     A pointer filled in with the transparent pixel in the native surface
 //     format
 //   
func (surface *Surface) GetColorKey() (retval int, key uint32) {
    tmp_key := new(C.Uint32)
    retval = int(C.SDL_GetColorKey((*C.SDL_Surface)(surface), (*C.Uint32)(tmp_key)))
    key = deref_uint32_ptr(tmp_key)
    return
}

 // Set an additional color value used in blit operations.
 // 
 // Returns: 0 on success, or -1 if the surface is not valid.
 // 
 // See also: SDL_GetSurfaceColorMod()
 // 
 //   surface
 //     The surface to update.
 //   
 //   r
 //     The red color value multiplied into blit operations.
 //   
 //   g
 //     The green color value multiplied into blit operations.
 //   
 //   b
 //     The blue color value multiplied into blit operations.
 //   
func (surface *Surface) SetColorMod(r uint8, g uint8, b uint8) (retval int) {
    retval = int(C.SDL_SetSurfaceColorMod((*C.SDL_Surface)(surface), C.Uint8(r), C.Uint8(g), C.Uint8(b)))
    return
}

 // Get the additional color value used in blit operations.
 // 
 // Returns: 0 on success, or -1 if the surface is not valid.
 // 
 // See also: SDL_SetSurfaceColorMod()
 // 
 //   surface
 //     The surface to query.
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
func (surface *Surface) GetColorMod() (retval int, r uint8, g uint8, b uint8) {
    tmp_r := new(C.Uint8)
    tmp_g := new(C.Uint8)
    tmp_b := new(C.Uint8)
    retval = int(C.SDL_GetSurfaceColorMod((*C.SDL_Surface)(surface), (*C.Uint8)(tmp_r), (*C.Uint8)(tmp_g), (*C.Uint8)(tmp_b)))
    r = deref_uint8_ptr(tmp_r)
    g = deref_uint8_ptr(tmp_g)
    b = deref_uint8_ptr(tmp_b)
    return
}

 // Set an additional alpha value used in blit operations.
 // 
 // Returns: 0 on success, or -1 if the surface is not valid.
 // 
 // See also: SDL_GetSurfaceAlphaMod()
 // 
 //   surface
 //     The surface to update.
 //   
 //   alpha
 //     The alpha value multiplied into blit operations.
 //   
func (surface *Surface) SetAlphaMod(alpha uint8) (retval int) {
    retval = int(C.SDL_SetSurfaceAlphaMod((*C.SDL_Surface)(surface), C.Uint8(alpha)))
    return
}

 // Get the additional alpha value used in blit operations.
 // 
 // Returns: 0 on success, or -1 if the surface is not valid.
 // 
 // See also: SDL_SetSurfaceAlphaMod()
 // 
 //   surface
 //     The surface to query.
 //   
 //   alpha
 //     A pointer filled in with the current alpha value.
 //   
func (surface *Surface) GetAlphaMod() (retval int, alpha uint8) {
    tmp_alpha := new(C.Uint8)
    retval = int(C.SDL_GetSurfaceAlphaMod((*C.SDL_Surface)(surface), (*C.Uint8)(tmp_alpha)))
    alpha = deref_uint8_ptr(tmp_alpha)
    return
}

 // Set the blend mode used for blit operations.
 // 
 // Returns: 0 on success, or -1 if the parameters are not valid.
 // 
 // See also: SDL_GetSurfaceBlendMode()
 // 
 //   surface
 //     The surface to update.
 //   
 //   blendMode
 //     SDL_BlendMode to use for blit blending.
 //   
func (surface *Surface) SetBlendMode(blendMode BlendMode) (retval int) {
    retval = int(C.SDL_SetSurfaceBlendMode((*C.SDL_Surface)(surface), C.SDL_BlendMode(blendMode)))
    return
}

 // Get the blend mode used for blit operations.
 // 
 // Returns: 0 on success, or -1 if the surface is not valid.
 // 
 // See also: SDL_SetSurfaceBlendMode()
 // 
 //   surface
 //     The surface to query.
 //   
 //   blendMode
 //     A pointer filled in with the current blend mode.
 //   
func (surface *Surface) GetBlendMode() (retval int, blendMode *BlendMode) {
    tmp_blendMode := new(C.SDL_BlendMode)
    retval = int(C.SDL_GetSurfaceBlendMode((*C.SDL_Surface)(surface), (*C.SDL_BlendMode)(tmp_blendMode)))
    blendMode = (*BlendMode)(unsafe.Pointer(tmp_blendMode))
    return
}

 // Sets the clipping rectangle for the destination surface in a blit.
 // 
 // If the clip rectangle is NULL, clipping will be disabled.
 // 
 // If the clip rectangle doesn't intersect the surface, the function will
 // return SDL_FALSE and blits will be completely clipped. Otherwise the
 // function returns SDL_TRUE and blits to the surface will be clipped to
 // the intersection of the surface area and the clipping rectangle.
 // 
 // Note that blits are automatically clipped to the edges of the source
 // and destination surfaces.
func (surface *Surface) SetClipRect(rect Rect) (retval bool) {
    tmp_rect := toCFromRect(rect)
    retval = C.SDL_TRUE==(C.SDL_SetClipRect((*C.SDL_Surface)(surface), (*C.SDL_Rect)(&tmp_rect)))
    return
}

 // Gets the clipping rectangle for the destination surface in a blit.
 // 
 // rect must be a pointer to a valid rectangle which will be filled with
 // the correct values.
func (surface *Surface) GetClipRect() (rect Rect) {
    tmp_rect := new(C.SDL_Rect)
    C.SDL_GetClipRect((*C.SDL_Surface)(surface), (*C.SDL_Rect)(tmp_rect))
    rect = fromC2Rect(*(tmp_rect))
    return
}

 // Creates a new surface of the specified format, and then copies and
 // maps the given surface to it so the blit of the converted surface will
 // be as fast as possible. If this function fails, it returns NULL.
 // 
 // The flags parameter is passed to SDL_CreateRGBSurface() and has those
 // semantics. You can also pass SDL_RLEACCEL in the flags parameter and
 // SDL will try to RLE accelerate colorkey and alpha blits in the
 // resulting surface.
func (src *Surface) Convert(fmt *PixelFormat, flags uint32) (retval *Surface) {
    retval = (*Surface)(unsafe.Pointer(C.SDL_ConvertSurface((*C.SDL_Surface)(src), (*C.SDL_PixelFormat)(fmt), C.Uint32(flags))))
    return
}

func (src *Surface) ConvertFormat(pixel_format uint32, flags uint32) (retval *Surface) {
    retval = (*Surface)(unsafe.Pointer(C.SDL_ConvertSurfaceFormat((*C.SDL_Surface)(src), C.Uint32(pixel_format), C.Uint32(flags))))
    return
}


 // Performs a fast fill of the given rectangle with color.
 // 
 // If rect is NULL, the whole surface will be filled with color.
 // 
 // The color should be a pixel of the format used by the surface, and can
 // be generated by the SDL_MapRGB() function.
 // 
 // Returns: 0 on success, or -1 on error.
 // 
func (dst *Surface) FillRect(rect Rect, color uint32) (retval int) {
    tmp_rect := toCFromRect(rect)
    retval = int(C.SDL_FillRect((*C.SDL_Surface)(dst), (*C.SDL_Rect)(&tmp_rect), C.Uint32(color)))
    return
}


 // This is the public blit function, SDL_BlitSurface(), and it performs
 // rectangle validation and clipping before passing it to SDL_LowerBlit()
func UpperBlit(src *Surface, srcrect Rect, dst *Surface, dstrect Rect) (retval int) {
    tmp_srcrect := toCFromRect(srcrect)
    tmp_dstrect := toCFromRect(dstrect)
    retval = int(C.SDL_UpperBlit((*C.SDL_Surface)(src), (*C.SDL_Rect)(&tmp_srcrect), (*C.SDL_Surface)(dst), (*C.SDL_Rect)(&tmp_dstrect)))
    return
}

 // This is a semi-private blit function and it performs low-level surface
 // blitting only.
func LowerBlit(src *Surface, srcrect Rect, dst *Surface, dstrect Rect) (retval int) {
    tmp_srcrect := toCFromRect(srcrect)
    tmp_dstrect := toCFromRect(dstrect)
    retval = int(C.SDL_LowerBlit((*C.SDL_Surface)(src), (*C.SDL_Rect)(&tmp_srcrect), (*C.SDL_Surface)(dst), (*C.SDL_Rect)(&tmp_dstrect)))
    return
}

 // Perform a fast, low quality, stretch blit between two surfaces of the
 // same pixel format.
 // 
 // Note: This function uses a static buffer, and is not thread-safe.
 // 
func SoftStretch(src *Surface, srcrect Rect, dst *Surface, dstrect Rect) (retval int) {
    tmp_srcrect := toCFromRect(srcrect)
    tmp_dstrect := toCFromRect(dstrect)
    retval = int(C.SDL_SoftStretch((*C.SDL_Surface)(src), (*C.SDL_Rect)(&tmp_srcrect), (*C.SDL_Surface)(dst), (*C.SDL_Rect)(&tmp_dstrect)))
    return
}

 // This is the public scaled blit function, SDL_BlitScaled(), and it
 // performs rectangle validation and clipping before passing it to
 // SDL_LowerBlitScaled()
func UpperBlitScaled(src *Surface, srcrect Rect, dst *Surface, dstrect Rect) (retval int) {
    tmp_srcrect := toCFromRect(srcrect)
    tmp_dstrect := toCFromRect(dstrect)
    retval = int(C.SDL_UpperBlitScaled((*C.SDL_Surface)(src), (*C.SDL_Rect)(&tmp_srcrect), (*C.SDL_Surface)(dst), (*C.SDL_Rect)(&tmp_dstrect)))
    return
}

 // This is a semi-private blit function and it performs low-level surface
 // scaled blitting only.
func LowerBlitScaled(src *Surface, srcrect Rect, dst *Surface, dstrect Rect) (retval int) {
    tmp_srcrect := toCFromRect(srcrect)
    tmp_dstrect := toCFromRect(dstrect)
    retval = int(C.SDL_LowerBlitScaled((*C.SDL_Surface)(src), (*C.SDL_Rect)(&tmp_srcrect), (*C.SDL_Surface)(dst), (*C.SDL_Rect)(&tmp_dstrect)))
    return
}
