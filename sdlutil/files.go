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

// Utility functions to complement Simple DirectMedia Layer (www.libsdl.org)
package sdlutil

import (
    "io"
    "path"
    "winterdrache.de/bindings/sdl"
)

var base_path string

// Returns the concatenation $BASE/path/compo/nents where
// $BASE is sdl.GetBasePath() (which is usually the location of the binary).
func GetResourcePath(pathcomponents ...string) string {
    if base_path == "" { // cache GetBasePath() because it may be slow
        base_path = sdl.GetBasePath()
    }
    pth := []string{base_path}
    pth = append(pth, pathcomponents...)
    return path.Join(pth...)
}

// Even though the argument is only required to implement io.Reader, this
// function will return an sdl.RWops that offers the functions corresponding
// to ALL of the following interfaces if supported by r:
//   io.Reader
//   io.Writer
//   io.Seeker
//   io.Closer
//   sdl.Sizer
// If something goes wrong in the SDL library (such as an out of memory on
// the C heap), this function calls panic(sdl.GetError()).
func RWFromReader(r io.Reader) *sdl.RWops {
    return rwFromIO(r)
}

// Even though the argument is only required to implement io.Writer, this
// function will return an sdl.RWops that offers the functions corresponding
// to ALL of the following interfaces if supported by w:
//   io.Reader
//   io.Writer
//   io.Seeker
//   io.Closer
//   sdl.Sizer
// If something goes wrong in the SDL library (such as an out of memory on
// the C heap), this function calls panic(sdl.GetError()).
func RWFromWriter(w io.Writer) *sdl.RWops {
    return rwFromIO(w)
}

// Not exported because it's too easy to mess up foo and *foo. The
// other RWFrom*() functions prevent that with compile time type checking
// and I see no application for creating an RWops that does not support
// either io.Reader or io.Writer.
func rwFromIO(iosomething interface{}) *sdl.RWops {
    rw := sdl.AllocRW()

    if x, ok := iosomething.(io.Reader); ok {
        rw.Read = x
    }
    if x, ok := iosomething.(io.Writer); ok {
        rw.Write = x
    }
    if x, ok := iosomething.(io.Closer); ok {
        rw.Close = x
    }
    if x, ok := iosomething.(io.Seeker); ok {
        rw.Seek = x
    }
    if x, ok := iosomething.(sdl.Sizer); ok {
        rw.Size = x
    }

    return rw.RWops()
}
