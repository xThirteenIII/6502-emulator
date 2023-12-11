package arc

import (
	"emulator/pkg/instructions"
	"fmt"
	"log"
)

// Typical of the PS:
// +---+---+---+---+---+---+---+---+
// | N | V |   | B | D | I | Z | C |
// +---+---+---+---+---+---+---+---+
// The missing bit is the Unused (U) or Expansion (E) bit.
// This bit is reserved for future use and should always be set to 1 when writing to the PSR.
// When reading the PSR, this bit should be ignored.

// The 6502 is an 8-bit processor internally, but it has a 16-bit address bus for addressing memory locations.
// This allows it to work with a larger address space even though its data bus is 8 bits wide.

const MaxMem = 1024 * 64

type Memory struct {

   Data [MaxMem]byte
}



func (mem *Memory) Initialise(){
    
    for i := 0; i < MaxMem; i++{
        mem.Data[i] = 0
    }
}
type ProcessorStatus struct {
    C uint
    Z uint
    I uint
    D uint
    B uint
    U uint  // This bit is reserved for future use and should always be set to 1 when writing to the PSR.
    V uint
    N uint
}

type CPU struct {

    // Program Counter points to the next instruction.
    PC uint16 

    // Stack Pointer is a 8-bit register that holds the low 8bits of the next free
    // location on the stack. Pushing bytes to the stack causes the Stack Pointer to be 
    // decremented. The CPU does not detect if the stack is overflowed by excessive pushing or
    // pulling operations and will most likely result in the program crashing.
    SP uint16 

    // Accumulator register is used for all arithmetical and logical operations.
    // Exception made for intcrements and decrements. The contents of the accumulator
    // can be stored and retrieved either from memory or the stack.
    A byte

    // The X register is most commonly used to hold counters and offsets for accessing.
    // It has one special purpose, which is to get a copy of the SP or change its value.
    X byte

    // The Y register is similar to the X register but has no special functions.
    Y byte

    // The Processor Status is a 8-bit status which holds a bunch of bits, that get set in it after operations.
    PS ProcessorStatus

    Memory Memory
}

func (cpu *CPU) Reset(){

    // Reset procedure does not follow accurate Commodor 64, it acts like a computer that's like a 
    // Commodor 64.

    // Reset vector address
    cpu.PC = 0xFFFC

    // Clear all flags
    cpu.PS.C = 0
    cpu.PS.Z = 0
    cpu.PS.I = 0
    cpu.PS.D = 0
    cpu.PS.B = 0
    cpu.PS.U = 0
    cpu.PS.V = 0
    cpu.PS.N = 0

    // The first stack access happens at address $0100 – a push first stores the value at $0100 + SP, then decrements SP.
    cpu.SP = 0x0100

    // Not sure if we want this to happen for now.
    cpu.A = 0
    cpu.X = 0
    cpu.Y = 0

    cpu.Memory.Initialise()
}


func (cpu *CPU) PrintValues(){
    fmt.Println("PC:", cpu.PC)
    fmt.Println("SP:", cpu.SP)
    fmt.Println("PS:", cpu.PS)
}

// FetchByte reads the byte located at the PC address
// It increases the program counter and takes a clock cycle
func (cpu *CPU) FetchByte( cycles *int) (byte, error){

    // TODO:Check if PC exceeds MAX_MEM
    data := cpu.Memory.Data[cpu.PC] 

    fmt.Println("fetched byte: ", data)

    cpu.PC++
    *cycles--

    return data, nil
}

func (cpu *CPU) FetchWord( cycles *int) (uint16, error){

    // TODO:Check if PC exceeds MAX_MEM

    // 6502 is little endian so first byte is the least significant byte of the data
    data := uint16(cpu.Memory.Data[cpu.PC])
    cpu.PC++
    *cycles--

    // second byte is the msb
    // e.g. data = 00000000 10011010 << 8 = 10011010 00000000
    data = data | (uint16(cpu.Memory.Data[cpu.PC]) << 8 )
    cpu.PC++
    *cycles--

    return data, nil
}

// ReadByte reads a piece of memory, without increasing the PC.
// It takes a clock cycle
func (cpu *CPU) ReadByte( cycles *int, address uint16) byte{

    // TODO:Check if PC exceeds MAX_MEM
    data := cpu.Memory.Data[address] 

    *cycles--

    return data
}

func (cpu *CPU) Execute( cycles *int ) error {

    for *cycles > 0 {

        fmt.Println(*cycles)

        // Fetch instruction, takes up one clock cycle
        ins, err := cpu.FetchByte(cycles)
        if err != nil {
            return err
        }
        
        // Execute code based on the opcode
        switch (ins) {

        case instructions.INS_LDA_IM:

            // Load value into A
            // Is it ok to modify directly the cycles in the ins variable? What if i need it after?
            var err error
            cpu.A , err = cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ", err.Error())
            }

            LDASetStatus(cpu)
            break;

        case instructions.INS_LDA_ZP:

            // First byte is the ZeroPage address
            // Second byte is the value to load

            // First cycle to fetch the instruction
            // Second cycle to fetch the address
            // The third cycle to read the data from the address
            zeroPageAddress, err := cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }
            
            // Load the data at the zeroPageAddress in the A register
            cpu.A = cpu.ReadByte(cycles, uint16(zeroPageAddress))
            fmt.Println(zeroPageAddress,cpu.Memory.Data[zeroPageAddress])

            LDASetStatus(cpu)
            break;

        case instructions.INS_LDA_ZPX:
            zeroPageAddress, err := cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // TODO: handle address overflow
            // The address calculation wraps around if the sum of the base address and the register exceed $FF.
            // ZeroPage address $80 and $FF in the X register then the accumulator will be loaded from $007F (e.g. $80 + $FF => $7F) and not $017F.
            *cycles--
            cpu.A = cpu.ReadByte(cycles, uint16(zeroPageAddress))

            LDASetStatus(cpu)
            break;

        case instructions.INS_JSR:

            // Fetch the targetMemoryAddress, which is where we have to jump to
            targetMemoryAddress, err := cpu.FetchWord(cycles)
            if err != nil {
                return err
            }

            // PC - 1 is the return address, where we return after the subRoutine exec
            cpu.Memory.WriteWord(cycles, cpu.SP, cpu.PC-1)
            cpu.SP++

            cpu.PC = targetMemoryAddress
            *cycles--

            break;

        default:
            log.Fatalln("Unknown opcode: ", ins)
        }
    }

    return nil
}

func (memory *Memory) WriteWord(cycles *int, word ,address uint16){

    // Little endian: we store LSB first
    memory.Data[address] = byte(word & 0xFF)
    *cycles--


    // Store MSB
    memory.Data[address+1] = byte(word >> 8)
    *cycles--

}

func LDASetStatus(cpu *CPU) {

            // Set Z flag if A is 0
            if cpu.A == 0 {
                cpu.PS.Z = 1
            }

            // Set N flag if the bit 7 of A is set
            // byte(1 << 7) is a bitmask that has the 7 bit set to 1
            // it left-shifts the 00000001 seven positions left
            if (cpu.A & byte(1 << 7) != 0) {
                cpu.PS.N = 1
            }
}
