package arc

const MaxMem = 1024 * 64

type Memory struct {

   Data [MaxMem]byte
}



func (mem *Memory) Initialise(){
    
    for i := 0; i < MaxMem; i++{
        mem.Data[i] = 0
    }
}
