package arc

import "log"

// Write one byte to memory
func (cpu *CPU) WriteByte(cycles *int, b byte ,address uint16){
    
    if address > MaxMem-1 {
        log.Fatalf("Program Counter exceeded max memory")
    }
    cpu.Memory.Data[address] = b
    *cycles--
}

// Write two bytes to memory
func (cpu *CPU) WriteWord(cycles *int, word ,address uint16){

    if address > MaxMem-1 {
        log.Fatalf("Program Counter exceeded max memory")
    }

    // Little endian: we store LSB first
    cpu.Memory.Data[address] = byte(word & 0xFF)
    *cycles--


    // Store MSB
    cpu.Memory.Data[address+1] = byte(word >> 8)
    *cycles--

}

// Write one byte to memory
func (cpu *CPU) WriteByteToStack(cycles *int, b byte){
    
    cpu.Memory.Data[cpu.SPTo16Address(cpu.SP)] = b
    cpu.SP--
    *cycles--
}

// TODO: for now we write MSB first and then LSB.
// That follows 6502 little endian architecture. Don't know if it's correct.
func (cpu *CPU) WriteWordToStack(cycles *int, word uint16){

    // Store MSB
    cpu.Memory.Data[cpu.SPTo16Address(cpu.SP)] = byte(word >> 8)
    cpu.SP--
    *cycles--

    cpu.Memory.Data[cpu.SPTo16Address(cpu.SP)] = byte(word & 0xFF)
    *cycles--
    cpu.SP--
}
