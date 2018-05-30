// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package ttf

// #include "includes.h"
import "C"

import "winterdrache.de/bindings/sdl"
import "unsafe"


const (
    MAJOR_VERSION = C.SDL_TTF_MAJOR_VERSION

    MINOR_VERSION = C.SDL_TTF_MINOR_VERSION

    PATCHLEVEL = C.SDL_TTF_PATCHLEVEL

    UNICODE_BOM_NATIVE = C.UNICODE_BOM_NATIVE

    UNICODE_BOM_SWAPPED = C.UNICODE_BOM_SWAPPED

    STYLE_NORMAL = C.TTF_STYLE_NORMAL

    STYLE_BOLD = C.TTF_STYLE_BOLD

    STYLE_ITALIC = C.TTF_STYLE_ITALIC

    STYLE_UNDERLINE = C.TTF_STYLE_UNDERLINE

    STYLE_STRIKETHROUGH = C.TTF_STYLE_STRIKETHROUGH

    HINTING_NORMAL = C.TTF_HINTING_NORMAL

    HINTING_LIGHT = C.TTF_HINTING_LIGHT

    HINTING_MONO = C.TTF_HINTING_MONO

    HINTING_NONE = C.TTF_HINTING_NONE
)

type Font C.TTF_Font


func Linked_Version() (retval *sdl.Version) {
    tmp_retval  := fromC2Version(*(C.TTF_Linked_Version()))
    retval  = &tmp_retval 
    return
}


func Init() (retval int) {
    retval = int(C.TTF_Init())
    return
}

func OpenFont(file string, ptsize int) (retval *Font) {
    tmp_file := C.CString(file); defer C.free(unsafe.Pointer(tmp_file))
    retval = (*Font)(unsafe.Pointer(C.TTF_OpenFont((*C.char)(tmp_file), C.int(ptsize))))
    return
}

func OpenFontIndex(file string, ptsize int, index int64) (retval *Font) {
    tmp_file := C.CString(file); defer C.free(unsafe.Pointer(tmp_file))
    retval = (*Font)(unsafe.Pointer(C.TTF_OpenFontIndex((*C.char)(tmp_file), C.int(ptsize), C.long(index))))
    return
}

func OpenFontRW(src *sdl.RWops, freesrc int, ptsize int) (retval *Font) {
    retval = (*Font)(unsafe.Pointer(C.TTF_OpenFontRW((*C.SDL_RWops)(unsafe.Pointer(src)), C.int(freesrc), C.int(ptsize))))
    return
}

func OpenFontIndexRW(src *sdl.RWops, freesrc int, ptsize int, index int64) (retval *Font) {
    retval = (*Font)(unsafe.Pointer(C.TTF_OpenFontIndexRW((*C.SDL_RWops)(unsafe.Pointer(src)), C.int(freesrc), C.int(ptsize), C.long(index))))
    return
}

func (font *Font) GetStyle() (retval int) {
    retval = int(C.TTF_GetFontStyle((*C.TTF_Font)(font)))
    return
}

func (font *Font) SetStyle(style int) {
    C.TTF_SetFontStyle((*C.TTF_Font)(font), C.int(style))
}

func (font *Font) GetOutline() (retval int) {
    retval = int(C.TTF_GetFontOutline((*C.TTF_Font)(font)))
    return
}

func (font *Font) SetOutline(outline int) {
    C.TTF_SetFontOutline((*C.TTF_Font)(font), C.int(outline))
}

func (font *Font) GetHinting() (retval int) {
    retval = int(C.TTF_GetFontHinting((*C.TTF_Font)(font)))
    return
}

func (font *Font) SetHinting(hinting int) {
    C.TTF_SetFontHinting((*C.TTF_Font)(font), C.int(hinting))
}

func (font *Font) Height() (retval int) {
    retval = int(C.TTF_FontHeight((*C.TTF_Font)(font)))
    return
}

func (font *Font) Ascent() (retval int) {
    retval = int(C.TTF_FontAscent((*C.TTF_Font)(font)))
    return
}

func (font *Font) Descent() (retval int) {
    retval = int(C.TTF_FontDescent((*C.TTF_Font)(font)))
    return
}

func (font *Font) LineSkip() (retval int) {
    retval = int(C.TTF_FontLineSkip((*C.TTF_Font)(font)))
    return
}

func (font *Font) GetKerning() (retval int) {
    retval = int(C.TTF_GetFontKerning((*C.TTF_Font)(font)))
    return
}

func (font *Font) SetKerning(allowed int) {
    C.TTF_SetFontKerning((*C.TTF_Font)(font), C.int(allowed))
}

func (font *Font) Faces() (retval int64) {
    retval = int64(C.TTF_FontFaces((*C.TTF_Font)(font)))
    return
}

func (font *Font) FaceIsFixedWidth() (retval int) {
    retval = int(C.TTF_FontFaceIsFixedWidth((*C.TTF_Font)(font)))
    return
}

func (font *Font) FaceFamilyName() (retval string) {
    retval = C.GoString(C.TTF_FontFaceFamilyName((*C.TTF_Font)(font)))
    return
}

func (font *Font) FaceStyleName() (retval string) {
    retval = C.GoString(C.TTF_FontFaceStyleName((*C.TTF_Font)(font)))
    return
}

func (font *Font) GlyphIsProvided(ch uint16) (retval int) {
    retval = int(C.TTF_GlyphIsProvided((*C.TTF_Font)(font), C.Uint16(ch)))
    return
}

func (font *Font) GlyphMetrics(ch uint16) (retval int, minx int, maxx int, miny int, maxy int, advance int) {
    tmp_minx := new(C.int)
    tmp_maxx := new(C.int)
    tmp_miny := new(C.int)
    tmp_maxy := new(C.int)
    tmp_advance := new(C.int)
    retval = int(C.TTF_GlyphMetrics((*C.TTF_Font)(font), C.Uint16(ch), (*C.int)(tmp_minx), (*C.int)(tmp_maxx), (*C.int)(tmp_miny), (*C.int)(tmp_maxy), (*C.int)(tmp_advance)))
    minx = deref_int_ptr(tmp_minx)
    maxx = deref_int_ptr(tmp_maxx)
    miny = deref_int_ptr(tmp_miny)
    maxy = deref_int_ptr(tmp_maxy)
    advance = deref_int_ptr(tmp_advance)
    return
}

func (font *Font) SizeText(text string) (retval int, w int, h int) {
    tmp_text := C.CString(text); defer C.free(unsafe.Pointer(tmp_text))
    tmp_w := new(C.int)
    tmp_h := new(C.int)
    retval = int(C.TTF_SizeText((*C.TTF_Font)(font), (*C.char)(tmp_text), (*C.int)(tmp_w), (*C.int)(tmp_h)))
    w = deref_int_ptr(tmp_w)
    h = deref_int_ptr(tmp_h)
    return
}

func (font *Font) SizeUTF8(text string) (retval int, w int, h int) {
    tmp_text := C.CString(text); defer C.free(unsafe.Pointer(tmp_text))
    tmp_w := new(C.int)
    tmp_h := new(C.int)
    retval = int(C.TTF_SizeUTF8((*C.TTF_Font)(font), (*C.char)(tmp_text), (*C.int)(tmp_w), (*C.int)(tmp_h)))
    w = deref_int_ptr(tmp_w)
    h = deref_int_ptr(tmp_h)
    return
}


func (font *Font) RenderText_Solid(text string, fg sdl.Color) (retval *sdl.Surface) {
    tmp_text := C.CString(text); defer C.free(unsafe.Pointer(tmp_text))
    retval = (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderText_Solid((*C.TTF_Font)(font), (*C.char)(tmp_text), toCFromColor(fg))))
    return
}

func (font *Font) RenderUTF8_Solid(text string, fg sdl.Color) (retval *sdl.Surface) {
    tmp_text := C.CString(text); defer C.free(unsafe.Pointer(tmp_text))
    retval = (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Solid((*C.TTF_Font)(font), (*C.char)(tmp_text), toCFromColor(fg))))
    return
}


func (font *Font) RenderGlyph_Solid(ch uint16, fg sdl.Color) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderGlyph_Solid((*C.TTF_Font)(font), C.Uint16(ch), toCFromColor(fg))))
    return
}

func (font *Font) RenderText_Shaded(text string, fg sdl.Color, bg sdl.Color) (retval *sdl.Surface) {
    tmp_text := C.CString(text); defer C.free(unsafe.Pointer(tmp_text))
    retval = (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderText_Shaded((*C.TTF_Font)(font), (*C.char)(tmp_text), toCFromColor(fg), toCFromColor(bg))))
    return
}

func (font *Font) RenderUTF8_Shaded(text string, fg sdl.Color, bg sdl.Color) (retval *sdl.Surface) {
    tmp_text := C.CString(text); defer C.free(unsafe.Pointer(tmp_text))
    retval = (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Shaded((*C.TTF_Font)(font), (*C.char)(tmp_text), toCFromColor(fg), toCFromColor(bg))))
    return
}


func (font *Font) RenderGlyph_Shaded(ch uint16, fg sdl.Color, bg sdl.Color) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderGlyph_Shaded((*C.TTF_Font)(font), C.Uint16(ch), toCFromColor(fg), toCFromColor(bg))))
    return
}

func (font *Font) RenderText_Blended(text string, fg sdl.Color) (retval *sdl.Surface) {
    tmp_text := C.CString(text); defer C.free(unsafe.Pointer(tmp_text))
    retval = (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderText_Blended((*C.TTF_Font)(font), (*C.char)(tmp_text), toCFromColor(fg))))
    return
}

func (font *Font) RenderUTF8_Blended(text string, fg sdl.Color) (retval *sdl.Surface) {
    tmp_text := C.CString(text); defer C.free(unsafe.Pointer(tmp_text))
    retval = (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Blended((*C.TTF_Font)(font), (*C.char)(tmp_text), toCFromColor(fg))))
    return
}


func (font *Font) RenderText_Blended_Wrapped(text string, fg sdl.Color, wrapLength uint32) (retval *sdl.Surface) {
    tmp_text := C.CString(text); defer C.free(unsafe.Pointer(tmp_text))
    retval = (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderText_Blended_Wrapped((*C.TTF_Font)(font), (*C.char)(tmp_text), toCFromColor(fg), C.Uint32(wrapLength))))
    return
}

func (font *Font) RenderUTF8_Blended_Wrapped(text string, fg sdl.Color, wrapLength uint32) (retval *sdl.Surface) {
    tmp_text := C.CString(text); defer C.free(unsafe.Pointer(tmp_text))
    retval = (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Blended_Wrapped((*C.TTF_Font)(font), (*C.char)(tmp_text), toCFromColor(fg), C.Uint32(wrapLength))))
    return
}


func (font *Font) RenderGlyph_Blended(ch uint16, fg sdl.Color) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderGlyph_Blended((*C.TTF_Font)(font), C.Uint16(ch), toCFromColor(fg))))
    return
}

func (font *Font) Close() {
    C.TTF_CloseFont((*C.TTF_Font)(font))
}

func Quit() {
    C.TTF_Quit()
}

func WasInit() (retval int) {
    retval = int(C.TTF_WasInit())
    return
}


func (font *Font) GetKerningSizeGlyphs(previous_ch uint16, ch uint16) (retval int) {
    retval = int(C.TTF_GetFontKerningSizeGlyphs((*C.TTF_Font)(font), C.Uint16(previous_ch), C.Uint16(ch)))
    return
}
