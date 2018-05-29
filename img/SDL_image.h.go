// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package img

// #include "includes.h"
import "C"

import "winterdrache.de/bindings/sdl"
import "unsafe"


const (
     // ↪ https://wiki.libsdl.org/SDL_IMAGE_MAJOR_VERSION
    IMAGE_MAJOR_VERSION = C.SDL_IMAGE_MAJOR_VERSION

     // ↪ https://wiki.libsdl.org/SDL_IMAGE_MINOR_VERSION
    IMAGE_MINOR_VERSION = C.SDL_IMAGE_MINOR_VERSION

     // ↪ https://wiki.libsdl.org/SDL_IMAGE_PATCHLEVEL
    IMAGE_PATCHLEVEL = C.SDL_IMAGE_PATCHLEVEL

     // This is the version number macro for the current SDL_image version.
    IMAGE_COMPILEDVERSION = C.SDL_IMAGE_COMPILEDVERSION
)

type InitFlags int
const (
    INIT_JPG InitFlags = C.IMG_INIT_JPG

    INIT_PNG InitFlags = C.IMG_INIT_PNG

    INIT_TIF InitFlags = C.IMG_INIT_TIF

    INIT_WEBP InitFlags = C.IMG_INIT_WEBP
)


func Linked_Version() (retval *sdl.Version) {
    tmp_retval  := fromC2Version(*(C.IMG_Linked_Version()))
    retval  = &tmp_retval 
    return
}

func Init(flags int) (retval int) {
    retval = int(C.IMG_Init(C.int(flags)))
    return
}

func Quit() {
    C.IMG_Quit()
}

func LoadTyped_RW(src *sdl.RWops, freesrc int, _type string) (retval *sdl.Surface) {
    tmp__type := C.CString(_type); defer C.free(unsafe.Pointer(tmp__type))
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadTyped_RW((*C.SDL_RWops)(unsafe.Pointer(src)), C.int(freesrc), (*C.char)(tmp__type))))
    return
}

func Load(file string) (retval *sdl.Surface) {
    tmp_file := C.CString(file); defer C.free(unsafe.Pointer(tmp_file))
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_Load((*C.char)(tmp_file))))
    return
}

func Load_RW(src *sdl.RWops, freesrc int) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_Load_RW((*C.SDL_RWops)(unsafe.Pointer(src)), C.int(freesrc))))
    return
}

func LoadTexture(renderer *sdl.Renderer, file string) (retval *sdl.Texture) {
    tmp_file := C.CString(file); defer C.free(unsafe.Pointer(tmp_file))
    retval = (*sdl.Texture)(unsafe.Pointer(C.IMG_LoadTexture((*C.SDL_Renderer)(unsafe.Pointer(renderer)), (*C.char)(tmp_file))))
    return
}

func LoadTexture_RW(renderer *sdl.Renderer, src *sdl.RWops, freesrc int) (retval *sdl.Texture) {
    retval = (*sdl.Texture)(unsafe.Pointer(C.IMG_LoadTexture_RW((*C.SDL_Renderer)(unsafe.Pointer(renderer)), (*C.SDL_RWops)(unsafe.Pointer(src)), C.int(freesrc))))
    return
}

func LoadTextureTyped_RW(renderer *sdl.Renderer, src *sdl.RWops, freesrc int, _type string) (retval *sdl.Texture) {
    tmp__type := C.CString(_type); defer C.free(unsafe.Pointer(tmp__type))
    retval = (*sdl.Texture)(unsafe.Pointer(C.IMG_LoadTextureTyped_RW((*C.SDL_Renderer)(unsafe.Pointer(renderer)), (*C.SDL_RWops)(unsafe.Pointer(src)), C.int(freesrc), (*C.char)(tmp__type))))
    return
}

func IsICO(src *sdl.RWops) (retval int) {
    retval = int(C.IMG_isICO((*C.SDL_RWops)(unsafe.Pointer(src))))
    return
}

func IsCUR(src *sdl.RWops) (retval int) {
    retval = int(C.IMG_isCUR((*C.SDL_RWops)(unsafe.Pointer(src))))
    return
}

func IsBMP(src *sdl.RWops) (retval int) {
    retval = int(C.IMG_isBMP((*C.SDL_RWops)(unsafe.Pointer(src))))
    return
}

func IsGIF(src *sdl.RWops) (retval int) {
    retval = int(C.IMG_isGIF((*C.SDL_RWops)(unsafe.Pointer(src))))
    return
}

func IsJPG(src *sdl.RWops) (retval int) {
    retval = int(C.IMG_isJPG((*C.SDL_RWops)(unsafe.Pointer(src))))
    return
}

func IsLBM(src *sdl.RWops) (retval int) {
    retval = int(C.IMG_isLBM((*C.SDL_RWops)(unsafe.Pointer(src))))
    return
}

func IsPCX(src *sdl.RWops) (retval int) {
    retval = int(C.IMG_isPCX((*C.SDL_RWops)(unsafe.Pointer(src))))
    return
}

func IsPNG(src *sdl.RWops) (retval int) {
    retval = int(C.IMG_isPNG((*C.SDL_RWops)(unsafe.Pointer(src))))
    return
}

func IsPNM(src *sdl.RWops) (retval int) {
    retval = int(C.IMG_isPNM((*C.SDL_RWops)(unsafe.Pointer(src))))
    return
}

func IsSVG(src *sdl.RWops) (retval int) {
    retval = int(C.IMG_isSVG((*C.SDL_RWops)(unsafe.Pointer(src))))
    return
}

func IsTIF(src *sdl.RWops) (retval int) {
    retval = int(C.IMG_isTIF((*C.SDL_RWops)(unsafe.Pointer(src))))
    return
}

func IsXCF(src *sdl.RWops) (retval int) {
    retval = int(C.IMG_isXCF((*C.SDL_RWops)(unsafe.Pointer(src))))
    return
}

func IsXPM(src *sdl.RWops) (retval int) {
    retval = int(C.IMG_isXPM((*C.SDL_RWops)(unsafe.Pointer(src))))
    return
}

func IsXV(src *sdl.RWops) (retval int) {
    retval = int(C.IMG_isXV((*C.SDL_RWops)(unsafe.Pointer(src))))
    return
}

func IsWEBP(src *sdl.RWops) (retval int) {
    retval = int(C.IMG_isWEBP((*C.SDL_RWops)(unsafe.Pointer(src))))
    return
}

func LoadICO_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadICO_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}

func LoadCUR_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadCUR_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}

func LoadBMP_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadBMP_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}

func LoadGIF_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadGIF_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}

func LoadJPG_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadJPG_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}

func LoadLBM_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadLBM_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}

func LoadPCX_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadPCX_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}

func LoadPNG_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadPNG_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}

func LoadPNM_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadPNM_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}

func LoadSVG_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadSVG_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}

func LoadTGA_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadTGA_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}

func LoadTIF_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadTIF_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}

func LoadXCF_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadXCF_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}

func LoadXPM_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadXPM_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}

func LoadXV_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadXV_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}

func LoadWEBP_RW(src *sdl.RWops) (retval *sdl.Surface) {
    retval = (*sdl.Surface)(unsafe.Pointer(C.IMG_LoadWEBP_RW((*C.SDL_RWops)(unsafe.Pointer(src)))))
    return
}


func SavePNG(surface *sdl.Surface, file string) (retval int) {
    tmp_file := C.CString(file); defer C.free(unsafe.Pointer(tmp_file))
    retval = int(C.IMG_SavePNG((*C.SDL_Surface)(unsafe.Pointer(surface)), (*C.char)(tmp_file)))
    return
}

func SavePNG_RW(surface *sdl.Surface, dst *sdl.RWops, freedst int) (retval int) {
    retval = int(C.IMG_SavePNG_RW((*C.SDL_Surface)(unsafe.Pointer(surface)), (*C.SDL_RWops)(unsafe.Pointer(dst)), C.int(freedst)))
    return
}

func SaveJPG(surface *sdl.Surface, file string, quality int) (retval int) {
    tmp_file := C.CString(file); defer C.free(unsafe.Pointer(tmp_file))
    retval = int(C.IMG_SaveJPG((*C.SDL_Surface)(unsafe.Pointer(surface)), (*C.char)(tmp_file), C.int(quality)))
    return
}

func SaveJPG_RW(surface *sdl.Surface, dst *sdl.RWops, freedst int, quality int) (retval int) {
    retval = int(C.IMG_SaveJPG_RW((*C.SDL_Surface)(unsafe.Pointer(surface)), (*C.SDL_RWops)(unsafe.Pointer(dst)), C.int(freedst), C.int(quality)))
    return
}
