package arc

func (memory *Memory) WriteByte(cycles *int, b byte ,address uint16){
    
    // TODO: check that address doesn't exceed MAX_MEM
    memory.Data[address] = b
    *cycles--
}

func (memory *Memory) WriteWord(cycles *int, word ,address uint16){

    // TODO: check that address doesn't exceed MAX_MEM

    // Little endian: we store LSB first
    memory.Data[address] = byte(word & 0xFF)
    *cycles--


    // Store MSB
    memory.Data[address+1] = byte(word >> 8)
    *cycles--

}
