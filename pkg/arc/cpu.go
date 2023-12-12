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



func (cpu *CPU) Execute( cycles *int ) error {

    for *cycles > 0 {


        // Fetch instruction, takes up one clock cycle
        /*LOOP:*/ ins, err := cpu.FetchByte(cycles)

        if err != nil {
            return err
        }
        
        // Decode instruction
        switch (ins) {


            // Execute operations based on the instruction opcode
        case instructions.INS_LDA_IM:

            // Load value into A
            // Is it ok to modify directly the cycles in the ins variable? What if i need it after?
            var err error
            cpu.A , err = cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ", err.Error())
            }

            // Set LDA status flags
            LDASetStatus(cpu)

            // Total cycles: 2
            // Total bytes: 2
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

            // Set LDA status flags
            LDASetStatus(cpu)

            // Total cycles: 3
            // Total bytes: 2
            break;

        case instructions.INS_LDA_ZPX:
            zeroPageAddress, err := cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // TODO: handle address overflow
            // The address calculation wraps around if the sum of the base address and the register exceed $FF.
            // ZeroPage address $80 and $FF in the X register then the accumulator will be loaded from $007F (e.g. $80 + $FF => $7F) and not $017F.

            // Does adding X to ZP take one cycle?
            AddRegValueToTarget8Address(cpu.X, &zeroPageAddress, cycles)
            cpu.A = cpu.ReadByte(cycles, uint16(zeroPageAddress))

            // Set LDA status flags
            LDASetStatus(cpu)

            // Total cycles: 4
            // Total bytes: 2
            break;

        case instructions.INS_JSR:

            // Fetch the targetMemoryAddress, which is where we have to jump to
            targetAddress, err := cpu.FetchWord(cycles)
            if err != nil {
                return err
            }

            // PC - 1 is the return address, where we return after the subRoutine exec
            cpu.Memory.WriteWord(cycles, cpu.SP, cpu.PC-1)
            cpu.SP++

            cpu.PC = targetAddress
            *cycles--

            // Total cycles: 6
            // Total bytes: 3

            break;
        case instructions.INS_LDA_ABS:

            // Fetch the target location using a full 16 bit address
            targetAddress, err := cpu.FetchWord(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // Load value at the target location into the A register
            cpu.A = cpu.ReadByte(cycles, targetAddress)

            // Set LDA status flags
            LDASetStatus(cpu)

            // Total cycles: 4
            // Total bytes: 3
            break;
        case instructions.INS_LDA_ABSX:

            // Fetch 16-bit address
            targetAddress, err := cpu.FetchWord(cycles) 
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // Add X value to the fetched address
            AddRegValueToTarget16Address(cpu.X, &targetAddress, cycles)


            // Load value stored at the address+X into the A register
            cpu.A = cpu.ReadByte(cycles, targetAddress)

            // Set LDA status flags
            LDASetStatus(cpu)

            // Total cycles: 4(+1 if page crossed)
            // Total bytes: 3
            break;
        case instructions.INS_LDA_ABSY:

            // Fetch 16-bit address
            targetAddress, err := cpu.FetchWord(cycles) 
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // Add X value to the fetched address
            AddRegValueToTarget16Address(cpu.Y, &targetAddress, cycles)


            // Load value stored at the address+X into the A register
            cpu.A = cpu.ReadByte(cycles, targetAddress)

            // Set LDA status flags
            LDASetStatus(cpu)

            // Total cycles: 4(+1 if page crossed)
            // Total bytes: 3
            break;
        case instructions.INS_LDA_INDX:
            // Set LDA status flags
            LDASetStatus(cpu)

            // Total cycles: 6
            // Total bytes: 2
            break;
        case instructions.INS_LDA_INDY:
            // Set LDA status flags
            LDASetStatus(cpu)

            // Total cycles: 5(+1 if page crossed)
            // Total bytes: 2
            break;

        default:
            log.Fatalln("Unknown opcode: ", ins)
        }
    }

    return nil
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

func ZeroPageWrapAround(value byte, address *byte) {

    // If the sum of the address and the value exceeds 0xFF, wrap around.
    // E.g. if I have to add 0x80 to 0xFF it results in 0x7F and not 0x017F

    // Use modulo operator to get the wrapped around value.
    *address = byte(uint16(value + *address) % 0x100)
}

func AddRegValueToTarget16Address(value byte, address *uint16, cycles *int){

    // "+1 if page crossed": a page boundary is crossed if the high byte of the original absolute address is
    // different from the high byte of the calculated address after adding the X register.
    // If a page boundary is crossed, an additional cycle is required.

    originalHighByte := (*address >> 8)
    
    // Add value to address
    *address += uint16(value)

    newHighByte := (*address >> 8)


    if originalHighByte != newHighByte {
        *cycles--
    }

    *cycles--
}

func AddRegValueToTarget8Address(value byte, address *byte, cycles *int){

    // "+1 if page crossed" can't happen

    *address += value
    *cycles--
}
