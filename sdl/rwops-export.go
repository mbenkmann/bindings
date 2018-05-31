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

//#include "includes.h"
//extern Sint64 go_impl_rwops_size(SDL_RWops* rw);
//extern Sint64 go_impl_rwops_seek(SDL_RWops* rw, Sint64 offset, int whence);
//extern size_t go_impl_rwops_read(SDL_RWops* rw, void* ptr, size_t size, size_t maxnum);
//extern size_t go_impl_rwops_write(SDL_RWops* rw, void* ptr, size_t size, size_t num);
//extern int go_impl_rwops_close(SDL_RWops* rw);
import "C"

import (
    "io"
    "unsafe"
)

//export go_impl_rwops_size
func go_impl_rwops_size(rw *C.SDL_RWops) C.Sint64 {
    id := *((*int)(unsafe.Pointer(&(rw.hidden))))
    rwops_mutex.Lock()
    custom := rwops[id]
    rwops_mutex.Unlock()
    if custom == nil {
        panic("size() called on already free'd or corrupted SDL_RWops")
    }

    // use Size if defined
    if custom.Size != nil {
        return C.Sint64(custom.Size.Size())
    }

    // fall back to emulating Size via Seek
    if custom.Seek != nil {
        pos, err := custom.Seek.Seek(0, io.SeekCurrent)
        if err != nil {
            return C.Sint64(-1)
        }

        size, err := custom.Seek.Seek(0, io.SeekEnd)
        if err != nil {
            size = -1
        }

        _, err = custom.Seek.Seek(pos, io.SeekStart)
        if err != nil {
            // If we can't restore the original position after seeking to
            // the end it's better to abort the program than produce
            // undefined results. This should never happen unless the
            // user has provided an incomplete implementation of io.Seeker.
            panic(err)
        }

        return C.Sint64(size)
    }

    return C.Sint64(-1)
}

//export go_impl_rwops_seek
func go_impl_rwops_seek(rw *C.SDL_RWops, offset C.Sint64, whence C.int) C.Sint64 {
    id := *((*int)(unsafe.Pointer(&(rw.hidden))))
    rwops_mutex.Lock()
    custom := rwops[id]
    rwops_mutex.Unlock()
    if custom == nil {
        panic("seek() called on already free'd or corrupted SDL_RWops")
    }

    if custom.Seek != nil {
        pos, err := custom.Seek.Seek(int64(offset), int(whence))
        if err == nil {
            return C.Sint64(pos)
        }
    }

    return C.Sint64(-1)
}

const bignumber = 999999999999

//export go_impl_rwops_read
func go_impl_rwops_read(rw *C.SDL_RWops, ptr unsafe.Pointer, size C.size_t, maxnum C.size_t) C.size_t {
    id := *((*int)(unsafe.Pointer(&(rw.hidden))))
    rwops_mutex.Lock()
    custom := rwops[id]
    rwops_mutex.Unlock()
    if custom == nil {
        panic("read() called on already free'd or corrupted SDL_RWops")
    }

    maxbytes := int64(size) * int64(maxnum)
    if size > bignumber || maxnum > bignumber || maxbytes > bignumber {
        maxbytes = bignumber
    }

    slice := (*(*[bignumber]byte)(ptr))[0:maxbytes]

    if custom.Read != nil {
        n, err := custom.Read.Read(slice)
        if err == nil && size > 0 {
            return C.size_t(int64(n) / int64(size))
        }
    }

    return C.size_t(0)
}

//export go_impl_rwops_write
func go_impl_rwops_write(rw *C.SDL_RWops, ptr unsafe.Pointer, size C.size_t, num C.size_t) C.size_t {
    id := *((*int)(unsafe.Pointer(&(rw.hidden))))
    rwops_mutex.Lock()
    custom := rwops[id]
    rwops_mutex.Unlock()
    if custom == nil {
        panic("write() called on already free'd or corrupted SDL_RWops")
    }

    maxbytes := int64(size) * int64(num)
    if size > bignumber || num > bignumber || maxbytes > bignumber {
        maxbytes = bignumber
    }

    slice := (*(*[bignumber]byte)(ptr))[0:maxbytes]

    if custom.Write != nil {
        n, err := custom.Write.Write(slice)
        if err == nil && size > 0 {
            return C.size_t(int64(n) / int64(size))
        }
    }

    return C.size_t(0)
}

//export go_impl_rwops_close
func go_impl_rwops_close(rw *C.SDL_RWops) C.int {
    id := *((*int)(unsafe.Pointer(&(rw.hidden))))
    rwops_mutex.Lock()
    custom := rwops[id]
    rwops_mutex.Unlock()
    if custom == nil {
        panic("close() called on already free'd or corrupted SDL_RWops")
    }

    retval := 0

    if custom.Close != nil {
        err := custom.Close.Close()
        if err != nil {
            retval = -1
        }
    }

    custom.internalFree()
    C.SDL_FreeRW(rw)

    return C.int(retval)
}

var go_rwops_size = C.go_impl_rwops_size
var go_rwops_seek = C.go_impl_rwops_seek
var go_rwops_read = C.go_impl_rwops_read
var go_rwops_write = C.go_impl_rwops_write
var go_rwops_close = C.go_impl_rwops_close
