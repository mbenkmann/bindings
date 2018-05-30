// Copyright (c) 2018 Matthias S. Benkmann
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// This file contains handwritten code for the parts of the SDL_image bindings that
// cannot reasonably be generated.

// Bindings for SDL_ttf (www.libsdl.org/projects/SDL_ttf)
package ttf

// #include "includes.h"
import "C"
import "unsafe"
import "winterdrache.de/bindings/sdl"

func freeGoString(s *C.char) string {
    defer C.free(unsafe.Pointer(s))
    return C.GoString(s)
}

func fromC2Version(s C.SDL_version) sdl.Version {
    return sdl.Version{uint8(s.major), uint8(s.minor), uint8(s.patch)}
}

func deref_int_ptr(i *C.int) int { return int(*i) }

func toCFromColor(s sdl.Color) (d C.SDL_Color) {
    d.r = C.Uint8(s.R)
    d.g = C.Uint8(s.G)
    d.b = C.Uint8(s.B)
    d.a = C.Uint8(s.A)
    return
}
