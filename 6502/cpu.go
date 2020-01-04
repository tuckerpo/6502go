package nes

// 6502 CPU

import (
    "fmt"
)

// 16 bit registers
type CPU struct {
    Cycles  uint64
    PC      uint16
    SP      byte
    A       byte
    X       byte
    Y       byte
    C       byte
    Z       byte
    I       byte
    D       byte
    B       byte
    V       byte
    N       byte
    inter   byte
    stall   byte

}

func NewCPU () *CPU {
    cpu := CPU
    cpu.reset()
    return &cpu
}
