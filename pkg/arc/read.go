package arc

import "log"

// ReadByte reads a piece of memory, without increasing the PC.
// It takes a clock cycle
func (cpu *CPU) ReadByte( cycles *int, address uint16) byte{

    if address > MaxMem-1 {
        log.Fatalf("Program Counter exceeded max memory")
    }
    data := cpu.Memory.Data[address] 

    *cycles--

    return data
}

func (cpu *CPU) ReadWord( cycles *int, address uint16) uint16{

    if address > MaxMem-1 {
        log.Fatalf("Program Counter exceeded max memory")
    }

    // Read low byte of address (LSB)
    data := uint16(cpu.Memory.Data[address])
    *cycles--

    // Read high byte of address (MSB)
    // e.g. data = 00000000 10011010 << 8 = 10011010 00000000
    data = data | (uint16(cpu.Memory.Data[address+1]) << 8 )
    *cycles--

    return data

}

// TODO: for now we write MSB first and decrement the SP
// That follows 6502 little endian architecture. Don't know if it's correct.
func (cpu *CPU) ReadByteFromStack(cycles *int) byte{

    cpu.SP++
    data := cpu.Memory.Data[cpu.SPTo16Address(cpu.SP)] 

    *cycles--

    return data
}

func (cpu *CPU) ReadWordFromStack( cycles *int) uint16{

    // Read low byte of address (LSB)
    cpu.SP++
    data := uint16(cpu.Memory.Data[cpu.SPTo16Address(cpu.SP)])
    *cycles--
    cpu.SP++

    // Read high byte of address (MSB)
    // e.g. data = 00000000 10011010 << 8 = 10011010 00000000
    data = data | (uint16(cpu.Memory.Data[cpu.SPTo16Address(cpu.SP)]) << 8 )
    *cycles--

    return data

}
