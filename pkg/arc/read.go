package arc

// ReadByte reads a piece of memory, without increasing the PC.
// It takes a clock cycle
func (cpu *CPU) ReadByte( cycles *int, address uint16) byte{

    // TODO:Check if PC exceeds MAX_MEM
    data := cpu.Memory.Data[address] 

    *cycles--

    return data
}

func (cpu *CPU) ReadWord( cycles *int, address uint16) uint16{

    // TODO:Check if PC exceeds MAX_MEM
    // Read low byte of address
    data := uint16(cpu.Memory.Data[address])
    *cycles--

    // Read high byte of address
    // e.g. data = 00000000 10011010 << 8 = 10011010 00000000
    data = data | (uint16(cpu.Memory.Data[address+1]) << 8 )
    *cycles--

    return data

}
