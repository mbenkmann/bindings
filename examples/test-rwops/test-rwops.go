package main

import (
    "fmt"
    "io"
    "winterdrache.de/bindings/sdl"
    "winterdrache.de/bindings/sdlutil"
)

var my_error = fmt.Errorf("error")

type Buf []byte

func (b *Buf) Seek(offset int64, whence int) (int64, error) {
    fmt.Printf("Seek(%v,%v)\n", offset, whence)
    if len(*b) == 0 {
        *b = []byte{0}
    }
    switch whence {
        case 0:
            if offset >= 0 && offset < int64(len(*b)) {
                (*b)[0] = byte(offset)
                return offset, nil
            }
        case 1:
            pos := int64((*b)[0]) + offset
            if pos >= 0 && pos < 256 {
                (*b)[0] = byte(pos)
                return pos, nil
            }
        case 2:
            pos := int64(len(*b)-1) + offset
            if pos >= 0 && pos < 256 {
                (*b)[0] = byte(pos)
                return pos, nil
            }
    }

    return -1, my_error
}

func (b *Buf) Write(p []byte) (n int, err error) {
    fmt.Printf("Write(%v)\n", p)
    if len(*b) == 0 {
        *b = []byte{0}
    }

    pos := int((*b)[0])
    maxn := 255 - pos
    n = len(p)
    if n > maxn {
        n = maxn
        err = my_error // disk full
    }

    for pos >= len(*b) {
        *b = append(*b, 0)
    }

    for i := 0; i < n; i++ {
        pos++
        if pos < len(*b) {
            (*b)[pos] = p[i]
        } else {
            *b = append(*b, p[i])
        }
    }

    (*b)[0] = byte(pos)

    return n, err
}

func (b *Buf) Read(p []byte) (n int, err error) {
    fmt.Printf("Read(%v)\n", len(p))
    if len(*b) == 0 {
        *b = []byte{0}
    }

    if len(p) == 0 {
        return 0, nil
    }

    pos := int((*b)[0])
    if pos+1 < len(*b) {
        for n = 0; n < len(p) && pos+1 < len(*b); n++ {
            pos++
            p[n] = (*b)[pos]
        }
        (*b)[0] = byte(pos)
        return n, nil
    }

    return 0, io.EOF
}

func (b *Buf) Close() error {
    fmt.Printf("Close()\n")
    *b = nil
    return nil
}

func (b *Buf) SizeX() int64 {
    fmt.Printf("Size()\n")
    if len(*b) == 0 {
        *b = []byte{0}
    }
    return int64(len(*b) - 1)
}

func main() {
    var buf Buf
    var rwops *sdl.RWops
    rwops = sdlutil.RWFromReader(&buf)
    rwops.WriteU8(42)
    rwops.WriteBE32(0x12345678)
    fmt.Println(buf.Seek(-4, 2))
    fmt.Printf("%x\n", rwops.ReadBE32())
    fmt.Println(rwops.Size())
    fmt.Println(buf.Seek(0, 0))
    fmt.Println(rwops.ReadU8())
    fmt.Println(buf.Seek(-1, 1))
    fmt.Println(rwops.ReadU8())
    fmt.Println(buf.Seek(3, 0))
    fmt.Println(rwops.Size())
    fmt.Printf("%x\n", rwops.ReadBE16())
    fmt.Println(buf.Seek(1, 0))
    fmt.Printf("%x\n", rwops.ReadLE32())
    rwops.Close()
}
