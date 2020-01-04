package nes

import (
    "os"
    "io"
)

/* 16 bit addresses byte map */

const (
    MEMSIZE uint16 = (1 << 16) - 1 // 6502 has 16 bit addr bus
)

type Memory interface {
    Read(address uint16) (value uint8) // Read byte from address
    Write(address uint16, value uint8) uint8 // Write byte to address
}

type RAM struct {
    m      []uint8 // memory map
    reads  bool // reads enabled
    writes bool // writes enabled
}

func NewRAM(size uint32) *RAM {
    return &RAM{
        m: make([]uint8, size),
    }
}

func (mem* RAM) SetMemoryReadable(enable bool) {
    mem.reads = enable
}

func (mem* RAM) SetMemoryWritable(enable bool) {
    mem.writes = enable
}

func (mem* RAM) Reset() {
    for i := range mem.m {
        mem.m[i] = 0
    }
}

func (mem* RAM) Read(address uint16) (value uint8) {
    if !mem.reads {
        value = mem.m[address]
    }
    return
}

func (mem* RAM) Write(address uint16, value uint8) {
    if !mem.writes {
        mem.m[address] = value
    }
}

// Open rom at path into high memory
func (mem *RAM) loadRom(path string) {
    file, err := os.Open(path)
    if err != nil {
        panic (err)
    }

    defer func() {
        if  err := file.Close(); err != nil {
            panic (err)
            }
        }()

        total := 0
        buf := make([]byte, (1 << 16) - 1)

        for {
            n, err := file.Read(buf)
            if err != nil  && err != io.EOF {
                panic (err)
            }

            if n == 0 {
                break
            }
            total ++
        }

        j := 0xc000

        for i, b := range buf {
            if i <= 15 {
                continue
            }
            mem.m[j] = b
            j++
            // filled upper memory with rom
            if j == 0xffff {
                break
            }
        }
    return
}

func SamePage(addr1 uint16, addr2 uint16) bool {
    mask1 := addr1 & 0xFF
    mask2 := addr2 & 0xFF
    return mask1 == mask2
}
