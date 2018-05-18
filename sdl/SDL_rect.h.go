// Machine-generated file.
// See http://winterdrache.de/bindings for details.

// Bindings for Simple DirectMedia Layer (www.libsdl.org)
package sdl

// #cgo linux freebsd darwin pkg-config: sdl2
// #include <SDL.h>
import "C"

 // Header file for SDL_rect definition and management functions.

 // The structure that defines a point.
 // 
 // See also: SDL_EnclosePoints
 // 
 // See also: SDL_PointInRect
 // 
type Point struct {
    X int

    Y int
}

func fromC2Point(s C.SDL_Point) Point {
    return Point{int(s.x), int(s.y)}
}

func toCFromPoint(s Point) (d C.SDL_Point) {
    d.x = C.int(s.X)
    d.y = C.int(s.Y)
    return
}

 // A rectangle, with the origin at the upper left.
 // 
 // See also: SDL_RectEmpty
 // 
 // See also: SDL_RectEquals
 // 
 // See also: SDL_HasIntersection
 // 
 // See also: SDL_IntersectRect
 // 
 // See also: SDL_UnionRect
 // 
 // See also: SDL_EnclosePoints
 // 
type Rect struct {
    X int

    Y int

    W int

    H int
}

func fromC2Rect(s C.SDL_Rect) Rect {
    return Rect{int(s.x), int(s.y), int(s.w), int(s.h)}
}

func toCFromRect(s Rect) (d C.SDL_Rect) {
    d.x = C.int(s.X)
    d.y = C.int(s.Y)
    d.w = C.int(s.W)
    d.h = C.int(s.H)
    return
}


 // Returns true if point resides inside a rectangle.
func PointInRect(p Point, r Rect) (retval bool) {
    tmp_p := toCFromPoint(p)
    tmp_r := toCFromRect(r)
    retval = C.SDL_TRUE==(C.SDL_PointInRect((*C.SDL_Point)(&tmp_p), (*C.SDL_Rect)(&tmp_r)))
    return
}

 // Returns true if the rectangle has no area.
func RectEmpty(r Rect) (retval bool) {
    tmp_r := toCFromRect(r)
    retval = C.SDL_TRUE==(C.SDL_RectEmpty((*C.SDL_Rect)(&tmp_r)))
    return
}

 // Returns true if the two rectangles are equal.
func RectEquals(a Rect, b Rect) (retval bool) {
    tmp_a := toCFromRect(a)
    tmp_b := toCFromRect(b)
    retval = C.SDL_TRUE==(C.SDL_RectEquals((*C.SDL_Rect)(&tmp_a), (*C.SDL_Rect)(&tmp_b)))
    return
}

 // Determine whether two rectangles intersect.
 // 
 // Returns: SDL_TRUE if there is an intersection, SDL_FALSE otherwise.
 // 
func HasIntersection(A Rect, B Rect) (retval bool) {
    tmp_A := toCFromRect(A)
    tmp_B := toCFromRect(B)
    retval = C.SDL_TRUE==(C.SDL_HasIntersection((*C.SDL_Rect)(&tmp_A), (*C.SDL_Rect)(&tmp_B)))
    return
}

 // Calculate the intersection of two rectangles.
 // 
 // Returns: SDL_TRUE if there is an intersection, SDL_FALSE otherwise.
 // 
func IntersectRect(A Rect, B Rect) (retval bool, result Rect) {
    tmp_A := toCFromRect(A)
    tmp_B := toCFromRect(B)
    tmp_result := new(C.SDL_Rect)
    retval = C.SDL_TRUE==(C.SDL_IntersectRect((*C.SDL_Rect)(&tmp_A), (*C.SDL_Rect)(&tmp_B), (*C.SDL_Rect)(tmp_result)))
    result = fromC2Rect(*(tmp_result))
    return
}

 // Calculate the union of two rectangles.
func UnionRect(A Rect, B Rect) (result Rect) {
    tmp_A := toCFromRect(A)
    tmp_B := toCFromRect(B)
    tmp_result := new(C.SDL_Rect)
    C.SDL_UnionRect((*C.SDL_Rect)(&tmp_A), (*C.SDL_Rect)(&tmp_B), (*C.SDL_Rect)(tmp_result))
    result = fromC2Rect(*(tmp_result))
    return
}


