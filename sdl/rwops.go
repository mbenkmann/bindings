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

package sdl

// #include "includes.h"
import "C"

import (
    "io"
    "sync"
    "unsafe"
)

// Sizer is an interface that complements io.Reader with the ability to
// determine the total size of the data. This must not be confused with
// a function like bytes.Buffer.Len() that returns the remaining number
// of bytes. The Size() method always returns the same number, even after
// reading parts of the data. If an object implements io.Seeker, it does
// not need to provide Size() because it can be emulated.
type Sizer interface {
    Size() int64
}

// Obtain this via sdl.AllocRW() and fill in at least Read or Write,
// then you can use CustomRWops.RWops() to get an *sdl.RWops which
// you can use with any of the SDL functions that use it.
//
// NOTE: Unless you have special needs, you should use
// sdlutil.RWFromReader() or sdlutil.RWFromWriter() instead of rolling
// your own CustomRWops.
type CustomRWops struct {
    Size  Sizer
    Seek  io.Seeker
    Read  io.Reader
    Write io.Writer
    Close io.Closer
    rwops *RWops
    id    int
}

// Make sure you call Free() on the CustomRWops object or Close() on
// its RWops() when you're done to avoid leaking memory.
// Alternatively many functions that accept *sdl.RWops arguments support a
// parameter that tells them to close the RWops when they're done. If you use
// this parameter that is sufficient. Unlike RWops.Close() CustomRWops.Free()
// may be called multiple times, so "defer customRW.Free()" is often a good idea.
func (rw *CustomRWops) Free() {
    if rw == nil || rw.rwops == nil { return }
    rw.rwops.Close()
}

// Returns a *RWops whose read/write/seek/close/size functions use rw's
// Read/Write/Seek/Close/Size.
func (rw *CustomRWops) RWops() *RWops {
    return rw.rwops
}

// NOTE: Unless you have special needs, use sdlutil.RWFromReader() or
// sdlutil.RWFromWriter() instead of sdl.AllocRW().
//
// Many functions from SDL and related libraries accept *sdl.RWops as IO
// abstractions. AllocRW() builds a bridge between sdl.RWops and Go's io
// interfaces. Fill in the appropriate fields of the returned structure
// with your object(s) and leave the others nil.
//
// ATTENTION: If you fill in Seek and leave Size empty, calls to Seek will be
// used to implement Size().
//
// Make sure you call Free() on the returned object or Close() on
// its RWops() when you're done to avoid leaking memory.
// Alternatively many functions that accept *sdl.RWops arguments support a
// parameter that tells them to close the RWops when they're done. If you use
// this parameter that is sufficient. Unlike RWops.Close() CustomRWops.Free()
// may be called multiple times, so "defer customRW.Free()" is often a good idea.
//
// If something goes wrong in the SDL library (such as an out of memory on
// the C heap), this function calls panic(sdl.GetError()).
func AllocRW() *CustomRWops {
    rw := &CustomRWops{}
    rwops_mutex.Lock()
    rw.id = rwops_next
    rwops_next++
    rwops[rw.id] = rw
    rwops_mutex.Unlock()

    sdl_rwops := C.SDL_AllocRW()
    if sdl_rwops == nil {
        panic(GetError())
    }

    sdl_rwops._type = 0
    *((*unsafe.Pointer)(unsafe.Pointer(&(sdl_rwops.size)))) = go_rwops_size
    *((*unsafe.Pointer)(unsafe.Pointer(&(sdl_rwops.seek)))) = go_rwops_seek
    *((*unsafe.Pointer)(unsafe.Pointer(&(sdl_rwops.read)))) = go_rwops_read
    *((*unsafe.Pointer)(unsafe.Pointer(&(sdl_rwops.write)))) = go_rwops_write
    *((*unsafe.Pointer)(unsafe.Pointer(&(sdl_rwops.close)))) = go_rwops_close
    // Note: It might seem intuitive to just store a pointer to CustomRWops
    // directly instead of using id and a map. However cgo rules state that
    // pointers to Go objects must not be stored in space allocated by C.
    // The reason is that the garbage collector is permitted to move Go
    // objects around and fix up pointers to them, but the GC does not look
    // at C allocated memory.
    *((*int)(unsafe.Pointer(&(sdl_rwops.hidden)))) = rw.id

    rw.rwops = (*RWops)(sdl_rwops)

    return rw
}

var rwops = map[int]*CustomRWops{}
var rwops_mutex sync.Mutex
var rwops_next int = 1

func (rw *CustomRWops) internalFree() {
    rw.Close = nil
    rw.Write = nil
    rw.Read = nil
    rw.Seek = nil
    rw.Size = nil
    rw.rwops.size = nil
    rw.rwops.seek = nil
    rw.rwops.read = nil
    rw.rwops.write = nil
    rw.rwops.close = nil
    rw.rwops = nil
    rwops_mutex.Lock()
    defer rwops_mutex.Unlock()
    delete(rwops, rw.id)
}
