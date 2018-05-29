// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #include "includes.h"
import "C"
import "unsafe"

 // Header for the enumerated pixel format definitions.

 // ↪ https://wiki.libsdl.org/SDL_Color
type Color struct {
    R uint8

    G uint8

    B uint8

    A uint8
}

func fromC2Color(s C.SDL_Color) Color {
    return Color{uint8(s.r), uint8(s.g), uint8(s.b), uint8(s.a)}
}

func toCFromColor(s Color) (d C.SDL_Color) {
    d.r = C.Uint8(s.R)
    d.g = C.Uint8(s.G)
    d.b = C.Uint8(s.B)
    d.a = C.Uint8(s.A)
    return
}

 // Transparency definitions
 // 
 // These define alpha as the opacity of a surface.
const (
     // ↪ https://wiki.libsdl.org/SDL_ALPHA_OPAQUE
    ALPHA_OPAQUE = C.SDL_ALPHA_OPAQUE

     // ↪ https://wiki.libsdl.org/SDL_ALPHA_TRANSPARENT
    ALPHA_TRANSPARENT = C.SDL_ALPHA_TRANSPARENT
)


 // Pixel type.
const (
    PIXELTYPE_UNKNOWN  = C.SDL_PIXELTYPE_UNKNOWN

    PIXELTYPE_INDEX1  = C.SDL_PIXELTYPE_INDEX1

    PIXELTYPE_INDEX4  = C.SDL_PIXELTYPE_INDEX4

    PIXELTYPE_INDEX8  = C.SDL_PIXELTYPE_INDEX8

    PIXELTYPE_PACKED8  = C.SDL_PIXELTYPE_PACKED8

    PIXELTYPE_PACKED16  = C.SDL_PIXELTYPE_PACKED16

    PIXELTYPE_PACKED32  = C.SDL_PIXELTYPE_PACKED32

    PIXELTYPE_ARRAYU8  = C.SDL_PIXELTYPE_ARRAYU8

    PIXELTYPE_ARRAYU16  = C.SDL_PIXELTYPE_ARRAYU16

    PIXELTYPE_ARRAYU32  = C.SDL_PIXELTYPE_ARRAYU32

    PIXELTYPE_ARRAYF16  = C.SDL_PIXELTYPE_ARRAYF16

    PIXELTYPE_ARRAYF32  = C.SDL_PIXELTYPE_ARRAYF32
)

 // Bitmap pixel order, high bit -> low bit.
const (
    BITMAPORDER_NONE  = C.SDL_BITMAPORDER_NONE

    BITMAPORDER_4321  = C.SDL_BITMAPORDER_4321

    BITMAPORDER_1234  = C.SDL_BITMAPORDER_1234
)

 // Packed component order, high bit -> low bit.
const (
    PACKEDORDER_NONE  = C.SDL_PACKEDORDER_NONE

    PACKEDORDER_XRGB  = C.SDL_PACKEDORDER_XRGB

    PACKEDORDER_RGBX  = C.SDL_PACKEDORDER_RGBX

    PACKEDORDER_ARGB  = C.SDL_PACKEDORDER_ARGB

    PACKEDORDER_RGBA  = C.SDL_PACKEDORDER_RGBA

    PACKEDORDER_XBGR  = C.SDL_PACKEDORDER_XBGR

    PACKEDORDER_BGRX  = C.SDL_PACKEDORDER_BGRX

    PACKEDORDER_ABGR  = C.SDL_PACKEDORDER_ABGR

    PACKEDORDER_BGRA  = C.SDL_PACKEDORDER_BGRA
)

 // Array component order, low byte -> high byte.
const (
    ARRAYORDER_NONE  = C.SDL_ARRAYORDER_NONE

    ARRAYORDER_RGB  = C.SDL_ARRAYORDER_RGB

    ARRAYORDER_RGBA  = C.SDL_ARRAYORDER_RGBA

    ARRAYORDER_ARGB  = C.SDL_ARRAYORDER_ARGB

    ARRAYORDER_BGR  = C.SDL_ARRAYORDER_BGR

    ARRAYORDER_BGRA  = C.SDL_ARRAYORDER_BGRA

    ARRAYORDER_ABGR  = C.SDL_ARRAYORDER_ABGR
)

 // Packed component layout.
const (
    PACKEDLAYOUT_NONE  = C.SDL_PACKEDLAYOUT_NONE

    PACKEDLAYOUT_332  = C.SDL_PACKEDLAYOUT_332

    PACKEDLAYOUT_4444  = C.SDL_PACKEDLAYOUT_4444

    PACKEDLAYOUT_1555  = C.SDL_PACKEDLAYOUT_1555

    PACKEDLAYOUT_5551  = C.SDL_PACKEDLAYOUT_5551

    PACKEDLAYOUT_565  = C.SDL_PACKEDLAYOUT_565

    PACKEDLAYOUT_8888  = C.SDL_PACKEDLAYOUT_8888

    PACKEDLAYOUT_2101010  = C.SDL_PACKEDLAYOUT_2101010

    PACKEDLAYOUT_1010102  = C.SDL_PACKEDLAYOUT_1010102
)

const (
    PIXELFORMAT_UNKNOWN  = C.SDL_PIXELFORMAT_UNKNOWN

    PIXELFORMAT_INDEX1LSB  = C.SDL_PIXELFORMAT_INDEX1LSB

    PIXELFORMAT_INDEX1MSB  = C.SDL_PIXELFORMAT_INDEX1MSB

    PIXELFORMAT_INDEX4LSB  = C.SDL_PIXELFORMAT_INDEX4LSB

    PIXELFORMAT_INDEX4MSB  = C.SDL_PIXELFORMAT_INDEX4MSB

    PIXELFORMAT_INDEX8  = C.SDL_PIXELFORMAT_INDEX8

    PIXELFORMAT_RGB332  = C.SDL_PIXELFORMAT_RGB332

    PIXELFORMAT_RGB444  = C.SDL_PIXELFORMAT_RGB444

    PIXELFORMAT_RGB555  = C.SDL_PIXELFORMAT_RGB555

    PIXELFORMAT_BGR555  = C.SDL_PIXELFORMAT_BGR555

    PIXELFORMAT_ARGB4444  = C.SDL_PIXELFORMAT_ARGB4444

    PIXELFORMAT_RGBA4444  = C.SDL_PIXELFORMAT_RGBA4444

    PIXELFORMAT_ABGR4444  = C.SDL_PIXELFORMAT_ABGR4444

    PIXELFORMAT_BGRA4444  = C.SDL_PIXELFORMAT_BGRA4444

    PIXELFORMAT_ARGB1555  = C.SDL_PIXELFORMAT_ARGB1555

    PIXELFORMAT_RGBA5551  = C.SDL_PIXELFORMAT_RGBA5551

    PIXELFORMAT_ABGR1555  = C.SDL_PIXELFORMAT_ABGR1555

    PIXELFORMAT_BGRA5551  = C.SDL_PIXELFORMAT_BGRA5551

    PIXELFORMAT_RGB565  = C.SDL_PIXELFORMAT_RGB565

    PIXELFORMAT_BGR565  = C.SDL_PIXELFORMAT_BGR565

    PIXELFORMAT_RGB24  = C.SDL_PIXELFORMAT_RGB24

    PIXELFORMAT_BGR24  = C.SDL_PIXELFORMAT_BGR24

    PIXELFORMAT_RGB888  = C.SDL_PIXELFORMAT_RGB888

    PIXELFORMAT_RGBX8888  = C.SDL_PIXELFORMAT_RGBX8888

    PIXELFORMAT_BGR888  = C.SDL_PIXELFORMAT_BGR888

    PIXELFORMAT_BGRX8888  = C.SDL_PIXELFORMAT_BGRX8888

    PIXELFORMAT_ARGB8888  = C.SDL_PIXELFORMAT_ARGB8888

    PIXELFORMAT_RGBA8888  = C.SDL_PIXELFORMAT_RGBA8888

    PIXELFORMAT_ABGR8888  = C.SDL_PIXELFORMAT_ABGR8888

    PIXELFORMAT_BGRA8888  = C.SDL_PIXELFORMAT_BGRA8888

    PIXELFORMAT_ARGB2101010  = C.SDL_PIXELFORMAT_ARGB2101010

    PIXELFORMAT_RGBA32  = C.SDL_PIXELFORMAT_RGBA32

    PIXELFORMAT_ARGB32  = C.SDL_PIXELFORMAT_ARGB32

    PIXELFORMAT_BGRA32  = C.SDL_PIXELFORMAT_BGRA32

    PIXELFORMAT_ABGR32  = C.SDL_PIXELFORMAT_ABGR32

     // Planar mode: Y + V + U (3 planes)
    PIXELFORMAT_YV12  = C.SDL_PIXELFORMAT_YV12

     // Planar mode: Y + U + V (3 planes)
    PIXELFORMAT_IYUV  = C.SDL_PIXELFORMAT_IYUV

     // Packed mode: Y0+U0+Y1+V0 (1 plane)
    PIXELFORMAT_YUY2  = C.SDL_PIXELFORMAT_YUY2

     // Packed mode: U0+Y0+V0+Y1 (1 plane)
    PIXELFORMAT_UYVY  = C.SDL_PIXELFORMAT_UYVY

     // Packed mode: Y0+V0+Y1+U0 (1 plane)
    PIXELFORMAT_YVYU  = C.SDL_PIXELFORMAT_YVYU

     // Planar mode: Y + U/V interleaved (2 planes)
    PIXELFORMAT_NV12  = C.SDL_PIXELFORMAT_NV12

     // Planar mode: Y + V/U interleaved (2 planes)
    PIXELFORMAT_NV21  = C.SDL_PIXELFORMAT_NV21

     // Android video texture format
    PIXELFORMAT_EXTERNAL_OES  = C.SDL_PIXELFORMAT_EXTERNAL_OES
)


 // Get the human readable name of a pixel format.
 // ↪ https://wiki.libsdl.org/SDL_GetPixelFormatName
func GetPixelFormatName(format uint32) (retval string) {
    retval = C.GoString(C.SDL_GetPixelFormatName(C.Uint32(format)))
    return
}

 // Convert one of the enumerated pixel formats to a bpp and RGBA masks.
 // 
 // Returns: SDL_TRUE, or SDL_FALSE if the conversion wasn't possible.
 // 
 // See also: SDL_MasksToPixelFormatEnum()
 // 
 // ↪ https://wiki.libsdl.org/SDL_PixelFormatEnumToMasks
func PixelFormatEnumToMasks(format uint32) (retval bool, bpp int, Rmask uint32, Gmask uint32, Bmask uint32, Amask uint32) {
    tmp_bpp := new(C.int)
    tmp_Rmask := new(C.Uint32)
    tmp_Gmask := new(C.Uint32)
    tmp_Bmask := new(C.Uint32)
    tmp_Amask := new(C.Uint32)
    retval = C.SDL_TRUE==(C.SDL_PixelFormatEnumToMasks(C.Uint32(format), (*C.int)(tmp_bpp), (*C.Uint32)(tmp_Rmask), (*C.Uint32)(tmp_Gmask), (*C.Uint32)(tmp_Bmask), (*C.Uint32)(tmp_Amask)))
    bpp = deref_int_ptr(tmp_bpp)
    Rmask = deref_uint32_ptr(tmp_Rmask)
    Gmask = deref_uint32_ptr(tmp_Gmask)
    Bmask = deref_uint32_ptr(tmp_Bmask)
    Amask = deref_uint32_ptr(tmp_Amask)
    return
}

 // Convert a bpp and RGBA masks to an enumerated pixel format.
 // 
 // Returns: The pixel format, or SDL_PIXELFORMAT_UNKNOWN if the
 // conversion wasn't possible.
 // 
 // See also: SDL_PixelFormatEnumToMasks()
 // 
 // ↪ https://wiki.libsdl.org/SDL_MasksToPixelFormatEnum
func MasksToPixelFormatEnum(bpp int, Rmask uint32, Gmask uint32, Bmask uint32, Amask uint32) (retval uint32) {
    retval = uint32(C.SDL_MasksToPixelFormatEnum(C.int(bpp), C.Uint32(Rmask), C.Uint32(Gmask), C.Uint32(Bmask), C.Uint32(Amask)))
    return
}

 // Create an SDL_PixelFormat structure from a pixel format enum.
 // ↪ https://wiki.libsdl.org/SDL_AllocFormat
func AllocFormat(pixel_format uint32) (retval *PixelFormat) {
    retval = (*PixelFormat)(unsafe.Pointer(C.SDL_AllocFormat(C.Uint32(pixel_format))))
    return
}

 // Free an SDL_PixelFormat structure.
 // ↪ https://wiki.libsdl.org/SDL_FreeFormat
func (format *PixelFormat) FreeFormat() {
    C.SDL_FreeFormat((*C.SDL_PixelFormat)(format))
}

 // Create a palette structure with the specified number of color entries.
 // 
 // Returns: A new palette, or NULL if there wasn't enough memory.
 // 
 // Note: The palette entries are initialized to white.
 // 
 // See also: SDL_FreePalette()
 // 
 // ↪ https://wiki.libsdl.org/SDL_AllocPalette
func AllocPalette(ncolors int) (retval *Palette) {
    retval = (*Palette)(unsafe.Pointer(C.SDL_AllocPalette(C.int(ncolors))))
    return
}

 // Set the palette for a pixel format structure.
 // ↪ https://wiki.libsdl.org/SDL_SetPixelFormatPalette
func (format *PixelFormat) SetPalette(palette *Palette) (retval int) {
    retval = int(C.SDL_SetPixelFormatPalette((*C.SDL_PixelFormat)(format), (*C.SDL_Palette)(palette)))
    return
}

 // Set a range of colors in a palette.
 // 
 // Returns: 0 on success, or -1 if not all of the colors could be set.
 // 
 //   palette
 //     The palette to modify.
 //   
 //   colors
 //     An array of colors to copy into the palette.
 //   
 //   firstcolor
 //     The index of the first palette entry to modify.
 //   
 //   ncolors
 //     The number of entries to modify.
 //   
 // ↪ https://wiki.libsdl.org/SDL_SetPaletteColors
func (palette *Palette) SetColors(colors []Color, firstcolor int) (retval int) {
    var tmp_colors *C.SDL_Color
    if len(colors) > 0 {
        sl_tmp_colors := make([]C.SDL_Color, len(colors))
        for i := range colors {
            sl_tmp_colors[i] = toCFromColor(colors[i])
        }
        tmp_colors = &(sl_tmp_colors[0])
    }
    tmp_ncolors := len(colors)
    retval = int(C.SDL_SetPaletteColors((*C.SDL_Palette)(palette), (tmp_colors), C.int(firstcolor), C.int(tmp_ncolors)))
    return
}

 // Free a palette created with SDL_AllocPalette().
 // 
 // See also: SDL_AllocPalette()
 // 
 // ↪ https://wiki.libsdl.org/SDL_FreePalette
func (palette *Palette) Free() {
    C.SDL_FreePalette((*C.SDL_Palette)(palette))
}

 // Maps an RGB triple to an opaque pixel value for a given pixel format.
 // 
 // See also: SDL_MapRGBA
 // 
 // ↪ https://wiki.libsdl.org/SDL_MapRGB
func (format *PixelFormat) MapRGB(r uint8, g uint8, b uint8) (retval uint32) {
    retval = uint32(C.SDL_MapRGB((*C.SDL_PixelFormat)(format), C.Uint8(r), C.Uint8(g), C.Uint8(b)))
    return
}

 // Maps an RGBA quadruple to a pixel value for a given pixel format.
 // 
 // See also: SDL_MapRGB
 // 
 // ↪ https://wiki.libsdl.org/SDL_MapRGBA
func (format *PixelFormat) MapRGBA(r uint8, g uint8, b uint8, a uint8) (retval uint32) {
    retval = uint32(C.SDL_MapRGBA((*C.SDL_PixelFormat)(format), C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)))
    return
}

 // Get the RGB components from a pixel of the specified format.
 // 
 // See also: SDL_GetRGBA
 // 
 // ↪ https://wiki.libsdl.org/SDL_GetRGB
func (format *PixelFormat) GetRGB(pixel uint32) (r byte, g byte, b byte) {
    tmp_r := new(C.Uint8)
    tmp_g := new(C.Uint8)
    tmp_b := new(C.Uint8)
    C.SDL_GetRGB(C.Uint32(pixel), (*C.SDL_PixelFormat)(format), (*C.Uint8)(tmp_r), (*C.Uint8)(tmp_g), (*C.Uint8)(tmp_b))
    r = deref_byte_ptr(tmp_r)
    g = deref_byte_ptr(tmp_g)
    b = deref_byte_ptr(tmp_b)
    return
}

 // Get the RGBA components from a pixel of the specified format.
 // 
 // See also: SDL_GetRGB
 // 
 // ↪ https://wiki.libsdl.org/SDL_GetRGBA
func (format *PixelFormat) GetRGBA(pixel uint32) (r byte, g byte, b byte, a byte) {
    tmp_r := new(C.Uint8)
    tmp_g := new(C.Uint8)
    tmp_b := new(C.Uint8)
    tmp_a := new(C.Uint8)
    C.SDL_GetRGBA(C.Uint32(pixel), (*C.SDL_PixelFormat)(format), (*C.Uint8)(tmp_r), (*C.Uint8)(tmp_g), (*C.Uint8)(tmp_b), (*C.Uint8)(tmp_a))
    r = deref_byte_ptr(tmp_r)
    g = deref_byte_ptr(tmp_g)
    b = deref_byte_ptr(tmp_b)
    a = deref_byte_ptr(tmp_a)
    return
}

