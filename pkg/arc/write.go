package arc

import "log"

func (cpu *CPU) WriteByte(cycles *int, b byte ,address uint16){
    
    if address > MaxMem-1 {
        log.Fatalf("Program Counter exceeded max memory")
    }
    cpu.Memory.Data[address] = b
    *cycles--
}

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
