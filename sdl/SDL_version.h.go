// Machine-generated file.
// See http://winterdrache.de/bindings for details.

package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
// #include <SDL_shape.h>
import "C"

 // This header defines the current SDL version.

 // Information the version of SDL in use.
 // 
 // Represents the library's version as three levels: major revision
 // (increments with massive changes, additions, and enhancements), minor
 // revision (increments with backwards-compatible changes to the major
 // revision), and patchlevel (increments with fixes to the minor
 // revision).
 // 
 // See also: SDL_VERSION
 // 
 // See also: SDL_GetVersion
 // 
type Version struct {
     // major version
    Major uint8

     // minor version
    Minor uint8

     // update version
    Patch uint8
}

func fromC2Version(s C.SDL_version) Version {
    return Version{uint8(s.major), uint8(s.minor), uint8(s.patch)}
}

func toCFromVersion(s Version) (d C.SDL_version) {
    d.major = C.Uint8(s.Major)
    d.minor = C.Uint8(s.Minor)
    d.patch = C.Uint8(s.Patch)
    return
}

const (
    MAJOR_VERSION = C.SDL_MAJOR_VERSION

    MINOR_VERSION = C.SDL_MINOR_VERSION

    PATCHLEVEL = C.SDL_PATCHLEVEL

     // This is the version number macro for the current SDL version.
    COMPILEDVERSION = C.SDL_COMPILEDVERSION
)


 // Get the version of SDL that is linked against your program.
 // 
 // If you are linking to SDL dynamically, then it is possible that the
 // current version will be different than the version you compiled
 // against. This function returns the current version, while
 // SDL_VERSION() is a macro that tells you what version you compiled
 // with.
 // 
 //   SDL_version compiled;
 //   SDL_version linked;
 //   
 //   SDL_VERSION(&compiled);
 //   SDL_GetVersion(&linked);
 //   printf("We compiled against SDL version %d.%d.%d ...\n",
 //          compiled.major, compiled.minor, compiled.patch);
 //   printf("But we linked against SDL version %d.%d.%d.\n",
 //          linked.major, linked.minor, linked.patch);
 // 
 // This function may be called safely at any time, even before
 // SDL_Init().
 // 
 // See also: SDL_VERSION
 // 
func GetVersion() (ver *Version) {
    tmp_ver := new(C.SDL_version)
    C.SDL_GetVersion((*C.SDL_version)(tmp_ver))
    tmp2_ver := fromC2Version(*(tmp_ver)); ver = &tmp2_ver
    return
}

 // Get the code revision of SDL that is linked against your program.
 // 
 // Returns an arbitrary string (a hash value) uniquely identifying the
 // exact revision of the SDL library in use, and is only useful in
 // comparing against other revisions. It is NOT an incrementing number.
func GetRevision() (retval string) {
    retval = C.GoString(C.SDL_GetRevision())
    return
}

 // Get the revision number of SDL that is linked against your program.
 // 
 // Returns a number uniquely identifying the exact revision of the SDL
 // library in use. It is an incrementing number based on commits to
 // hg.libsdl.org.
func GetRevisionNumber() (retval int) {
    retval = int(C.SDL_GetRevisionNumber())
    return
}
