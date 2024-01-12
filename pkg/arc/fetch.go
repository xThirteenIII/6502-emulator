package arc

import (
	"log"
)

// FetchByte reads the byte located at the PC address
// It increases the program counter and decrements clock cycles by 1
// It return an error if PC exceeds max memory (65535 B)
func (cpu *CPU) FetchByte( cycles *int) byte{

    // Exceeding max memory halts the cpu (Fatal log)
    if cpu.PC > MaxMem-1 {
        log.Fatalf("Program Counter exceeded max memory")
    }
    data := cpu.Memory.Data[cpu.PC] 

    cpu.PC++
    *cycles--

    return data
}

// FetchWord consumes 2 clock cycles
func (cpu *CPU) FetchWord( cycles *int) uint16{

    if cpu.PC > MaxMem-1 {
        log.Fatalf("Program Counter exceeded max memory")
    }

    // 6502 is little endian so first byte is the least significant byte of the data
    // Fetch low byte of address
    data := uint16(cpu.Memory.Data[cpu.PC])
    cpu.PC++
    *cycles--

    // second byte is the msb
    // e.g. data = 00000000 10011010 << 8 = 10011010 00000000
    // Fetch high byte of address
    data = data | (uint16(cpu.Memory.Data[cpu.PC]) << 8 )
    cpu.PC++
    *cycles--

    return data
}
