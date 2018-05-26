package main

import (
    "fmt"
    "unsafe"
    "winterdrache.de/bindings/sdl"
)

//extern int HitTest(void* p0, void* p1, void* p2);
import "C"

//export HitTest
func HitTest(w unsafe.Pointer, point unsafe.Pointer, u unsafe.Pointer) C.int {
    win := (*sdl.Window)(w)
    win_x, win_y := win.GetPosition()
    x, y := int((*[2]C.int)(point)[0]), int((*[2]C.int)(point)[1])
    fmt.Printf("Hit test: x=%v y=%v  Window: x=%v y=%v\n", x, y, win_x, win_y)
    return C.int(sdl.HITTEST_DRAGGABLE)
}

var myHitTest = C.HitTest
