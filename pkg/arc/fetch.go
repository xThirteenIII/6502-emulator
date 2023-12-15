package arc

// FetchByte reads the byte located at the PC address
// It increases the program counter and takes a clock cycle
func (cpu *CPU) FetchByte( cycles *int) (byte, error){

    // TODO:Check if PC exceeds MAX_MEM
    data := cpu.Memory.Data[cpu.PC] 


    cpu.PC++
    *cycles--

    return data, nil
}

func (cpu *CPU) FetchWord( cycles *int) (uint16, error){

    // TODO: Check if PC exceeds MAX_MEM
    // TODO: Handle error

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

    return data, nil
}
