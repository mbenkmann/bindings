package sdl

// disabled#cgo CFLAGS: -D_REENTRANT -I/tmp/SDL2-2.0.9/include
// disabled#cgo LDFLAGS: -L/tmp/SDL2-2.0.9/build/.libs -lSDL2

// #cgo linux freebsd darwin pkg-config: sdl2
import "C"
