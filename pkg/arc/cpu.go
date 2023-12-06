package arc

import "fmt"

// Typical of the PS:
// +---+---+---+---+---+---+---+---+
// | N | V |   | B | D | I | Z | C |
// +---+---+---+---+---+---+---+---+
// The missing bit is the Unused (U) or Expansion (E) bit.
// This bit is reserved for future use and should always be set to 1 when writing to the PSR.
// When reading the PSR, this bit should be ignored.

// The 6502 is an 8-bit processor internally, but it has a 16-bit address bus for addressing memory locations.
// This allows it to work with a larger address space even though its data bus is 8 bits wide.


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
}

func (cpu *CPU) Reset( memory *Memory){

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
}

func (cpu *CPU) FetchByte( cycles *int,  memory *Memory) byte{

    Data := memory.Data[cpu.PC] 

    *cycles--

    return Data
}

func (cpu *CPU) PrintValues(){
    fmt.Println("PC:", cpu.PC)
    fmt.Println("SP:", cpu.SP)
    fmt.Println("PS:", cpu.PS)
}
