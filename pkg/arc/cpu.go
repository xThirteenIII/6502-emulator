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
    // Goes from $0100 to $01FF, it's fixed.
    SP byte  

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
    cpu.SP = 0x00

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
        ins, err := cpu.FetchByte(cycles)

        if err != nil {
            return err
        }
        
        // Decode instruction
        switch (ins) {


            // Execute operations based on the instruction opcode
        case instructions.INS_LDA_IM:

            // Load value into A
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
            // Wrap Around
            zeroPageAddress = byte(uint16(cpu.X + zeroPageAddress) % 0x100)
            *cycles--

            cpu.A = cpu.ReadByte(cycles, uint16(zeroPageAddress))


            // Set LDA status flags
            LDASetStatus(cpu)

            // Total cycles: 4
            // Total bytes: 2
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

            // TODO: cycles count it's not right
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

            // Add Y value to the fetched address
            AddRegValueToTarget16Address(cpu.Y, &targetAddress, cycles)


            // Load value stored at the address+X into the A register
            cpu.A = cpu.ReadByte(cycles, targetAddress)

            // Set LDA status flags
            LDASetStatus(cpu)

            // Total cycles: 4(+1 if page crossed)
            // Total bytes: 3
            break;
        case instructions.INS_LDA_INDX:

            // In this mode the X register is used to offste the zero page vector,
            // used to determine the effective address.
            // Put another way, the vector is chosen by adding the value in the X register,
            // to the given zero page address.
            // The resulting zero page address is the vector from which the effective address is read.
            // Weird stuff.

            // Example:
            // LDX #$04
            // LDA ($02, X)

            // In the above case X is loaded with four, so the vector is calculated with 
            // $02 + $04 = $06 (resulting vector)
            // If the zero page memory $06 contains: 00 80, then the effective address from the vector (06)
            // would be $8000

            // Fetch the Zero Page Address
            zeroPageAddress, err := cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // Wrap Around
            zeroPageAddress = byte(uint16(cpu.X + zeroPageAddress) % 0x100)
            *cycles--

            // The effective address is the 
            effectiveAddress := cpu.ReadWord(cycles, uint16(zeroPageAddress))

            cpu.A = cpu.ReadByte(cycles, effectiveAddress)
            // Set LDA status flags

            LDASetStatus(cpu)

            // Total cycles: 6
            // Total bytes: 2
            break;
        case instructions.INS_LDA_INDY:

            // Fetch the Zero Page Address
            zeroPageAddress, err := cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            effectiveAddress := cpu.ReadWord(cycles, uint16(zeroPageAddress))

            // Add X to the Zero Page Address
            AddRegValueToTarget16Address(cpu.Y, &effectiveAddress, cycles)

            cpu.A = cpu.ReadByte(cycles, effectiveAddress)

            // Set LDA status flags
            LDASetStatus(cpu)

            // Total cycles: 5(+1 if page crossed)
            // Total bytes: 2
            break;

        case instructions.INS_LDX_IM:
            // Load value into X
            var err error
            cpu.X , err = cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ", err.Error())
            }

            // Set LDX status flags
            LDXSetStatus(cpu)

            // Total cycles: 2
            // Total bytes: 2
            break;

        case instructions.INS_LDX_ZP:

            // First byte is the ZeroPage address
            // Second byte is the value to load

            // First cycle to fetch the instruction
            // Second cycle to fetch the address
            // The third cycle to read the data from the address
            zeroPageAddress, err := cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }
            
            // Load the data at the zeroPageAddress in the X register
            cpu.X = cpu.ReadByte(cycles, uint16(zeroPageAddress))

            // Set LDX status flags
            LDXSetStatus(cpu)

            // Total cycles: 3
            // Total bytes: 2
            break;

        case instructions.INS_LDX_ZPY:

            zeroPageAddress, err := cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // TODO: handle address overflow
            // Wrap Around
            zeroPageAddress = byte(uint16(cpu.Y + zeroPageAddress) % 0x100)
            *cycles--

            cpu.X = cpu.ReadByte(cycles, uint16(zeroPageAddress))


            // Set LDX status flags
            LDXSetStatus(cpu)

            // Total cycles: 4
            // Total bytes: 2
            break;
        case instructions.INS_LDX_ABS:

            // Fetch the target location using a full 16 bit address
            targetAddress, err := cpu.FetchWord(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // Load value at the target location into the A register
            cpu.X = cpu.ReadByte(cycles, targetAddress)

            // Set LDA status flags
            LDXSetStatus(cpu)

            // Total cycles: 4
            // Total bytes: 3
            break;

        case instructions.INS_LDX_ABSY:

            // Fetch 16-bit address
            targetAddress, err := cpu.FetchWord(cycles) 
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // Add Y value to the fetched address
            AddRegValueToTarget16Address(cpu.Y, &targetAddress, cycles)


            // Load value stored at the address+X into the A register
            cpu.Y = cpu.ReadByte(cycles, targetAddress)

            // Set LDA status flags
            LDASetStatus(cpu)

            // Total cycles: 4(+1 if page crossed)
            // Total bytes: 3
            break;

        case instructions.INS_LDY_IM:
            // Load value into Y
            var err error
            cpu.Y , err = cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ", err.Error())
            }

            // Set LDX status flags
            LDYSetStatus(cpu)

            // Total cycles: 2
            // Total bytes: 2
            break;

        case instructions.INS_LDY_ZP:

            // First byte is the ZeroPage address
            // Second byte is the value to load

            // First cycle to fetch the instruction
            // Second cycle to fetch the address
            // The third cycle to read the data from the address
            zeroPageAddress, err := cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }
            
            // Load the data at the zeroPageAddress in the Y register
            cpu.Y = cpu.ReadByte(cycles, uint16(zeroPageAddress))

            // Set LDX status flags
            LDYSetStatus(cpu)

            // Total cycles: 3
            // Total bytes: 2
            break;

        case instructions.INS_LDY_ZPX:

            zeroPageAddress, err := cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // TODO: handle address overflow
            // Wrap Around
            zeroPageAddress = byte(uint16(cpu.X + zeroPageAddress) % 0x100)
            *cycles--

            cpu.Y = cpu.ReadByte(cycles, uint16(zeroPageAddress))


            // Set LDX status flags
            LDYSetStatus(cpu)

            // Total cycles: 4
            // Total bytes: 2
            break;
        case instructions.INS_LDY_ABS:

            // Fetch the target location using a full 16 bit address
            targetAddress, err := cpu.FetchWord(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // Load value at the target location into the Y register
            cpu.Y = cpu.ReadByte(cycles, targetAddress)

            // Set LDA status flags
            LDYSetStatus(cpu)

            // Total cycles: 4
            // Total bytes: 3
            break;

        case instructions.INS_LDY_ABSX:

            // Fetch 16-bit address
            targetAddress, err := cpu.FetchWord(cycles) 
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // Add X value to the fetched address
            AddRegValueToTarget16Address(cpu.X, &targetAddress, cycles)


            // Load value stored at the address+X into the Y register
            cpu.Y = cpu.ReadByte(cycles, targetAddress)

            // Set LDA status flags
            LDYSetStatus(cpu)

            // Total cycles: 4(+1 if page crossed)
            // Total bytes: 3
            break;

        case instructions.INS_STA_ZP:

            var err error
            zeroPageAddress, err := cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            cpu.Memory.WriteByte(cycles, cpu.A, uint16(zeroPageAddress))
            
            // Total cycles: 3
            // Total bytes: 2
            break;

        case instructions.INS_STA_ZPX:

            zeroPageAddress, err := cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // TODO: handle address overflow
            // Wrap Around
            zeroPageAddress = byte(uint16(cpu.X + zeroPageAddress) % 0x100)
            *cycles--

            cpu.Memory.WriteByte(cycles, cpu.A, uint16(zeroPageAddress))

            // Total cycles: 4
            // Total bytes: 2

            break;

        case instructions.INS_STA_ABS:

            // Fetch the target location using a full 16 bit address
            targetAddress, err := cpu.FetchWord(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            cpu.Memory.WriteByte(cycles, cpu.A, targetAddress)

            // Total cycles: 4
            // Total bytes: 3
            break;

        case instructions.INS_STA_ABSX:

            // Fetch the target location using a full 16 bit address
            targetAddress, err := cpu.FetchWord(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // Add X to the target address
            // AddRegValueToTarget16Address() not used because totale cycles amount to 5, 
            // independently from page crossing.
            targetAddress += uint16(cpu.X)
            *cycles--

            cpu.Memory.WriteByte(cycles, cpu.A, targetAddress)

            // Total cycles: 5
            // Total bytes: 3
            break;

        case instructions.INS_STA_ABSY:

            // Fetch the target location using a full 16 bit address
            targetAddress, err := cpu.FetchWord(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // Add Y to the target address
            // AddRegValueToTarget16Address() not used because totale cycles amount to 5, 
            // independently from page crossing.
            targetAddress += uint16(cpu.Y)
            *cycles--

            cpu.Memory.WriteByte(cycles, cpu.A, targetAddress)

            // Total cycles: 5
            // Total bytes: 3
            break;


        case instructions.INS_STA_INDX:

            // In this mode the X register is used to offste the zero page vector,
            // used to determine the effective address.
            // Put another way, the vector is chosen by adding the value in the X register,
            // to the given zero page address.
            // The resulting zero page address is the vector from which the effective address is read.
            // Weird stuff.

            // Example:
            // LDX #$04
            // LDA ($02, X)

            // In the above case X is loaded with four, so the vector is calculated with 
            // $02 + $04 = $06 (resulting vector)
            // If the zero page memory $06 contains: 00 80, then the effective address from the vector (06)
            // would be $8000

            // Fetch the Zero Page Address
            zeroPageAddress, err := cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // Wrap Around
            zeroPageAddress = byte(uint16(cpu.X + zeroPageAddress) % 0x100)
            *cycles--

            // The effective address is the 
            effectiveAddress := cpu.ReadWord(cycles, uint16(zeroPageAddress))

            cpu.Memory.WriteByte(cycles, cpu.A, effectiveAddress)

            // Total cycles: 6
            // Total bytes: 2
            break;
        case instructions.INS_STA_INDY:

            // Fetch the Zero Page Address
            zeroPageAddress, err := cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            effectiveAddress := cpu.ReadWord(cycles, uint16(zeroPageAddress))

            // Add Y to the Zero Page Address
            AddRegValueToTarget16Address(cpu.Y, &effectiveAddress, cycles)

            cpu.Memory.WriteByte(cycles, cpu.A, effectiveAddress)

            // Total cycles: 6
            // Total bytes: 2
            break;

        case instructions.INS_JSR:

            // Fetch the targetMemoryAddress, which is where we have to jump to
            targetAddress, err := cpu.FetchWord(cycles)
            if err != nil {
                return err
            }

            // Set cpu.SP 16-bit, which is 0100 + cpu.SP (8 bit)

            SP := uint16(cpu.SP) + 0x0100

            // PC - 1 is the return address, where we return after the subRoutine exec
            cpu.Memory.WriteWord(cycles, SP, cpu.PC-1)
            cpu.SP++

            cpu.PC = targetAddress
            *cycles--

            // Total cycles: 6
            // Total bytes: 3

            break;

        case instructions.INS_STX_ZP:

            var err error
            zeroPageAddress, err := cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            cpu.Memory.WriteByte(cycles, cpu.X, uint16(zeroPageAddress))
            
            // Total cycles: 3
            // Total bytes: 2
            break;

        case instructions.INS_STX_ZPY:

            zeroPageAddress, err := cpu.FetchByte(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            // TODO: handle address overflow
            // Wrap Around
            zeroPageAddress = byte(uint16(cpu.Y + zeroPageAddress) % 0x100)
            *cycles--

            cpu.Memory.WriteByte(cycles, cpu.X, uint16(zeroPageAddress))

            // Total cycles: 4
            // Total bytes: 2

            break;

        case instructions.INS_STX_ABS:

            // Fetch the target location using a full 16 bit address
            targetAddress, err := cpu.FetchWord(cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }

            cpu.Memory.WriteByte(cycles, cpu.X, targetAddress)

            // Total cycles: 4
            // Total bytes: 3
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

func LDXSetStatus(cpu *CPU) {

            // Set Z flag if X is 0
            if cpu.X == 0 {
                cpu.PS.Z = 1
            }

            // Set N flag if the bit 7 of X is set
            // byte(1 << 7) is a bitmask that has the 7 bit set to 1
            // it left-shifts the 00000001 seven positions left
            if (cpu.X & byte(1 << 7) != 0) {
                cpu.PS.N = 1
            }
}

func LDYSetStatus(cpu *CPU) {

            // Set Z flag if Y is 0
            if cpu.Y == 0 {
                cpu.PS.Z = 1
            }

            // Set N flag if the bit 7 of Y is set
            // byte(1 << 7) is a bitmask that has the 7 bit set to 1
            // it left-shifts the 00000001 seven positions left
            if (cpu.Y & byte(1 << 7) != 0) {
                cpu.PS.N = 1
            }
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
}
