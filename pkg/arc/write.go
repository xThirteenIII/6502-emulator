package arc

import "log"

func (memory *Memory) WriteByte(cycles *int, b byte ,address uint16){
    
    if address > MaxMem-1 {
        log.Fatalf("Program Counter exceeded max memory")
    }
    memory.Data[address] = b
    *cycles--
}

func (memory *Memory) WriteWord(cycles *int, word ,address uint16){

    if address > MaxMem-1 {
        log.Fatalf("Program Counter exceeded max memory")
    }

    // Little endian: we store LSB first
    memory.Data[address] = byte(word & 0xFF)
    *cycles--


    // Store MSB
    memory.Data[address+1] = byte(word >> 8)
    *cycles--

}
