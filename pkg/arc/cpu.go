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

func (cpu *CPU) Reset(resetVector uint16){

    // Reset procedure does not follow accurate Commodor 64, it acts like a computer that's like a 
    // Commodor 64.

    // Reset vector address
    cpu.PC = resetVector

    // Clear all flags
    cpu.PS.C = 0
    cpu.PS.Z = 0
    cpu.PS.I = 0
    cpu.PS.D = 0
    cpu.PS.B = 0
    cpu.PS.U = 0
    cpu.PS.V = 0
    cpu.PS.N = 0

    // After the Reset, there's 9 post-reset cycles, which execute three fake push into the stack.
    // The final SP is therefore 00 - 1 = FF, FF - 1 = FE, FF - 1 = FD
    // https://www.c64-wiki.com/wiki/Reset_(Process)
    cpu.SP = 0xFD

    // Not sure if we want this to happen for now.
    cpu.A = 0
    cpu.X = 0
    cpu.Y = 0

    cpu.Memory.Initialise()
}


func (cpu *CPU) PrintStatus(){
    fmt.Println("PC:", cpu.PC)
    fmt.Println("SP:", cpu.SP)
    fmt.Println("PS:", cpu.PS)
}



// Execute runs the fetch-decode loop.
// It fetches the instruction byte and then, based on the opcode fetched, 
// executes the corresponding instruction.
// It returns the number of cycles used, for Testing purposes.
func (cpu *CPU) Execute( cycles int ) ( cyclesUsed int) {

    // At the beginning, initialise cyclesUsed as the number of cycles passed when calling the
    // method Execute().
    cyclesUsed = cycles

    // Can we get stuck in infinite loop if we pass more cycles than expected?
    // Not for now because since memory is initialised to 0, if we try to fetch a 
    // byte from one more cell memory where we are not supposed to be, it fetches 0 and
    // exits the switch loop with the default case
    for cycles > 0 {


        // Fetch instruction, takes up one clock cycle
        // PC++
        ins := cpu.FetchByte(&cycles)
        
        // Decode instruction
        switch (ins) {


            // Execute operations based on the instruction opcode
        case instructions.INS_LDA_IM:

            cpu.A = cpu.FetchByte(&cycles)

            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)

            // Total cycles: 2
            // Total bytes: 2
            break;

        case instructions.INS_LDA_ZP:

            // First byte is the ZeroPage address
            // Second byte is the value to load

            // First cycle to fetch the instruction
            // Second cycle to fetch the address
            // The third cycle to read the data from the address
            zeroPageAddress := cpu.AddressZeroPage(&cycles)
            
            LoadRegisterAndSetStatusFlags(cpu, &cycles, zeroPageAddress, &cpu.A)

            // Total cycles: 3
            // Total bytes: 2
            break;

        case instructions.INS_LDA_ZPX:

            zeroPageAddress := cpu.AddressZeroPageX(&cycles)

            LoadRegisterAndSetStatusFlags(cpu, &cycles, zeroPageAddress, &cpu.A)

            // Total cycles: 4
            // Total bytes: 2
            break;
        case instructions.INS_LDA_ABS:

            // Fetch the target location using a full 16 bit address
            targetAddress := cpu.AddressAbsolute(&cycles)

            LoadRegisterAndSetStatusFlags(cpu, &cycles, targetAddress, &cpu.A)

            // Total cycles: 4
            // Total bytes: 3
            break;
        case instructions.INS_LDA_ABSX:

            // TODO: cycles count it's not right
            // Fetch 16-bit address
            targetAddress := cpu.AddressAbsoluteX(&cycles)

            LoadRegisterAndSetStatusFlags(cpu, &cycles, targetAddress, &cpu.A)

            // Total cycles: 4(+1 if page crossed)
            // Total bytes: 3
            break;
        case instructions.INS_LDA_ABSY:

            targetAddress := cpu.AddressAbsoluteY(&cycles)

            LoadRegisterAndSetStatusFlags(cpu, &cycles, targetAddress, &cpu.A)

            // Total cycles: 4(+1 if page crossed)
            // Total bytes: 3
            break;
        case instructions.INS_LDA_INDX:

            // In this mode the X register is used to offset the zero page vector,
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

            // This takes 4 cycles
            effectiveAddress := cpu.AddressIndirectX(&cycles)

            // This takes 1 cycle
            LoadRegisterAndSetStatusFlags(cpu, &cycles,effectiveAddress, &cpu.A)

            // Total cycles: 6
            // Total bytes: 2
            break;
        case instructions.INS_LDA_INDY:

            // This takes 3+1 cycles
            effectiveAddress := cpu.AddressIndirectY(&cycles)

            // This takes 1 cycle
            LoadRegisterAndSetStatusFlags(cpu, &cycles,effectiveAddress, &cpu.A)

            // Total cycles: 5(+1 if page crossed)
            // Total bytes: 2
            break;

        case instructions.INS_LDX_IM:

            // Load value into X
            cpu.X = cpu.FetchByte(&cycles)

            // Set LDX status flags
            SetZeroAndNegativeFlags(cpu, cpu.X)

            // Total cycles: 2
            // Total bytes: 2
            break;

        case instructions.INS_LDX_ZP:

            // First byte is the ZeroPage address
            // Second byte is the value to load

            // First cycle to fetch the instruction
            // Second cycle to fetch the address
            // The third cycle to read the data from the address
            zeroPageAddress := cpu.AddressZeroPage(&cycles)
            
            LoadRegisterAndSetStatusFlags(cpu, &cycles, zeroPageAddress, &cpu.X)

            // Total cycles: 3
            // Total bytes: 2
            break;

        case instructions.INS_LDX_ZPY:

            zeroPageAddress := cpu.AddressZeroPageY(&cycles)

            LoadRegisterAndSetStatusFlags(cpu, &cycles, zeroPageAddress, &cpu.X)
            // Total cycles: 4
            // Total bytes: 2
            break;

        case instructions.INS_LDX_ABS:

            // Fetch the target location using a full 16 bit address
            targetAddress := cpu.AddressAbsolute(&cycles)

            LoadRegisterAndSetStatusFlags(cpu, &cycles, targetAddress, &cpu.X)

            // Total cycles: 4
            // Total bytes: 3
            break;

        case instructions.INS_LDX_ABSY:

            // Fetch 16-bit address
            targetAddress := cpu.AddressAbsoluteY(&cycles)

            LoadRegisterAndSetStatusFlags(cpu, &cycles, targetAddress, &cpu.X)

            // Total cycles: 4(+1 if page crossed)
            // Total bytes: 3
            break;

        case instructions.INS_LDY_IM:

            // Load value into Y
            cpu.Y = cpu.FetchByte(&cycles)

            // Set LDX status flags
            SetZeroAndNegativeFlags(cpu, cpu.Y)

            // Total cycles: 2
            // Total bytes: 2
            break;

        case instructions.INS_LDY_ZP:

            // First byte is the ZeroPage address
            // Second byte is the value to load

            // First cycle to fetch the instruction
            // Second cycle to fetch the address
            // The third cycle to read the data from the address
            zeroPageAddress := cpu.AddressZeroPage(&cycles)
            
            LoadRegisterAndSetStatusFlags(cpu, &cycles, zeroPageAddress, &cpu.Y)

            // Total cycles: 3
            // Total bytes: 2
            break;

        case instructions.INS_LDY_ZPX:

            zeroPageAddress := cpu.AddressZeroPageX(&cycles)
            
            LoadRegisterAndSetStatusFlags(cpu, &cycles, zeroPageAddress, &cpu.Y)

            // Total cycles: 4
            // Total bytes: 2
            break;
        case instructions.INS_LDY_ABS:

            // Fetch the target location using a full 16 bit address
            targetAddress := cpu.AddressAbsolute(&cycles)

            LoadRegisterAndSetStatusFlags(cpu, &cycles, targetAddress, &cpu.Y)
            // Total cycles: 4
            // Total bytes: 3
            break;

        case instructions.INS_LDY_ABSX:

            // Fetch the target location using a full 16 bit address
            targetAddress := cpu.AddressAbsoluteX(&cycles)

            LoadRegisterAndSetStatusFlags(cpu, &cycles, targetAddress, &cpu.Y)

            // Total cycles: 4(+1 if page crossed)
            // Total bytes: 3
            break;

        case instructions.INS_STA_ZP:

            zeroPageAddress := cpu.AddressZeroPage(&cycles)

            cpu.WriteByte(&cycles, cpu.A, zeroPageAddress)
            
            // Total cycles: 3
            // Total bytes: 2
            break;

        case instructions.INS_STA_ZPX:

            zeroPageAddress := cpu.AddressZeroPageX(&cycles)

            cpu.WriteByte(&cycles, cpu.A, zeroPageAddress)

            // Total cycles: 4
            // Total bytes: 2

            break;

        case instructions.INS_STA_ABS:

            // Fetch the target location using a full 16 bit address
            targetAddress := cpu.AddressAbsolute(&cycles)

            cpu.WriteByte(&cycles, cpu.A, targetAddress)

            // Total cycles: 4
            // Total bytes: 3
            break;

        case instructions.INS_STA_ABSX:

            // Why does this take 5 cycles flat? Weird

            // Fetch the target location using a full 16 bit address
            targetAddress := cpu.FetchWord(&cycles)

            // Add X to the target address
            // AddRegValueToTarget16Address() not used because totale cycles amount to 5, 
            // independently from page crossing.
            targetAddress += uint16(cpu.X)
            cycles--

            cpu.WriteByte(&cycles, cpu.A, targetAddress)

            // Total cycles: 5
            // Total bytes: 3
            break;

        case instructions.INS_STA_ABSY:

            // Fetch the target location using a full 16 bit address
            targetAddress := cpu.FetchWord(&cycles)

            // Add Y to the target address
            // AddRegValueToTarget16Address() not used because totale cycles amount to 5, 
            // independently from page crossing.
            targetAddress += uint16(cpu.Y)
            cycles--

            cpu.WriteByte(&cycles, cpu.A, targetAddress)

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

            effectiveAddress := cpu.AddressIndirectX(&cycles)

            cpu.WriteByte(&cycles, cpu.A, effectiveAddress)

            // Total cycles: 6
            // Total bytes: 2
            break;
        case instructions.INS_STA_INDY:

            // Fetch the Zero Page Address
            zeroPageAddress := cpu.AddressZeroPage(&cycles)

            effectiveAddress := cpu.ReadWord(&cycles, zeroPageAddress)

            // Add Y to the Zero Page Address
            // AddRegValueToTarget16Address() not used because totale cycles amount to 6, 
            // independently from page crossing.
            effectiveAddress += uint16(cpu.Y)
            cycles--

            cpu.WriteByte(&cycles, cpu.A, effectiveAddress)

            // Total cycles: 6
            // Total bytes: 2
            break;

        case instructions.INS_STX_ZP:

            zeroPageAddress := cpu.FetchByte(&cycles)

            cpu.WriteByte(&cycles, cpu.X, uint16(zeroPageAddress))
            
            // Total cycles: 3
            // Total bytes: 2
            break;

        case instructions.INS_STX_ZPY:

            zeroPageAddress := cpu.FetchByte(&cycles)

            // TODO: handle address overflow
            // Wrap Around
            zeroPageAddress = byte(uint16(cpu.Y + zeroPageAddress) % 0x100)
            cycles--

            cpu.WriteByte(&cycles, cpu.X, uint16(zeroPageAddress))

            // Total cycles: 4
            // Total bytes: 2

            break;

        case instructions.INS_STX_ABS:

            // Fetch the target location using a full 16 bit address
            targetAddress := cpu.FetchWord(&cycles)

            cpu.WriteByte(&cycles, cpu.X, targetAddress)

            // Total cycles: 4
            // Total bytes: 3
            break;

        case instructions.INS_STY_ZP:

            zeroPageAddress := cpu.FetchByte(&cycles)

            cpu.WriteByte(&cycles, cpu.Y, uint16(zeroPageAddress))
            
            // Total cycles: 3
            // Total bytes: 2
            break;

        case instructions.INS_STY_ZPX:

            zeroPageAddress := cpu.FetchByte(&cycles)

            // TODO: handle address overflow
            // Wrap Around
            zeroPageAddress = byte(uint16(cpu.X + zeroPageAddress) % 0x100)
            cycles--

            cpu.WriteByte(&cycles, cpu.Y, uint16(zeroPageAddress))

            // Total cycles: 4
            // Total bytes: 2

            break;

        case instructions.INS_STY_ABS:

            // Fetch the target location using a full 16 bit address
            targetAddress := cpu.FetchWord(&cycles)

            cpu.WriteByte(&cycles, cpu.Y, targetAddress)

            // Total cycles: 4
            // Total bytes: 3
            break;

        case instructions.INS_TAX_IMP:

            // Copy the current contents of the accumulator into the X register and sets the zero and negative flags as appropriate.
            // Implicit:
            // For many 6502 instructions the source and destination of the information to be manipulated
            // is implied directly by the function of the instruction itself and no further operand needs to be specified.
            // Operations like 'Clear Carry Flag' (CLC) and 'Return from Subroutine' (RTS) are implicit.

            cpu.X = cpu.A
            cycles--

            SetZeroAndNegativeFlags(cpu, cpu.X)




            // Total cycles: 2
            // Total bytes: 1
            break;

        case instructions.INS_TAY_IMP:

            // Copy the current contents of the accumulator into the X register and sets the zero and negative flags as appropriate.
            // Implicit:
            // For many 6502 instructions the source and destination of the information to be manipulated
            // is implied directly by the function of the instruction itself and no further operand needs to be specified.
            // Operations like 'Clear Carry Flag' (CLC) and 'Return from Subroutine' (RTS) are implicit.

            cpu.Y = cpu.A
            cycles--

            SetZeroAndNegativeFlags(cpu, cpu.Y)




            // Total cycles: 2
            // Total bytes: 1
            break;

        case instructions.INS_TXA_IMP:

            // Copy the current contents of the accumulator into the X register and sets the zero and negative flags as appropriate.
            // Implicit:
            // For many 6502 instructions the source and destination of the information to be manipulated
            // is implied directly by the function of the instruction itself and no further operand needs to be specified.
            // Operations like 'Clear Carry Flag' (CLC) and 'Return from Subroutine' (RTS) are implicit.

            cpu.A = cpu.X
            cycles--

            SetZeroAndNegativeFlags(cpu, cpu.A)




            // Total cycles: 2
            // Total bytes: 1
            break;

        case instructions.INS_TYA_IMP:

            // Copy the current contents of the accumulator into the X register and sets the zero and negative flags as appropriate.
            // Implicit:
            // For many 6502 instructions the source and destination of the information to be manipulated
            // is implied directly by the function of the instruction itself and no further operand needs to be specified.
            // Operations like 'Clear Carry Flag' (CLC) and 'Return from Subroutine' (RTS) are implicit.

            cpu.A = cpu.Y
            cycles--

            SetZeroAndNegativeFlags(cpu, cpu.A)

            // Total cycles: 2
            // Total bytes: 1
            break;

        case instructions.INS_TSX_IMP:

            // Copy the current contents of the accumulator into the X register and sets the zero and negative flags as appropriate.
            // Implicit:
            // For many 6502 instructions the source and destination of the information to be manipulated
            // is implied directly by the function of the instruction itself and no further operand needs to be specified.
            // Operations like 'Clear Carry Flag' (CLC) and 'Return from Subroutine' (RTS) are implicit.

            cpu.X = cpu.SP
            cycles--

            SetZeroAndNegativeFlags(cpu, cpu.X)

            // Total cycles: 2
            // Total bytes: 1
            break;

        case instructions.INS_TXS_IMP:

            // Copy the current contents of the accumulator into the X register and sets the zero and negative flags as appropriate.
            // Implicit:
            // For many 6502 instructions the source and destination of the information to be manipulated
            // is implied directly by the function of the instruction itself and no further operand needs to be specified.
            // Operations like 'Clear Carry Flag' (CLC) and 'Return from Subroutine' (RTS) are implicit.

            cpu.SP = cpu.X
            cycles--

            // Total cycles: 2
            // Total bytes: 1
            break;

        case instructions.INS_PHA_IMP:

            cpu.PushByteToStack(&cycles, cpu.A)

            cycles--

            // Total cycles: 3
            // Total bytes: 1
            break;

        case instructions.INS_PHP_IMP:

            cpu.PushByteToStack(&cycles, cpu.PSToByte())
            cycles--

            // Total cycles: 3
            // Total bytes: 1
            break;

        case instructions.INS_PLA_IMP:

            cpu.A = cpu.PopByteFromStack(&cycles)

            SetZeroAndNegativeFlags(cpu, cpu.A)

            cycles-=2

            // Total cycles: 4
            // Total bytes: 1
            break;

        case instructions.INS_PLP_IMP:

            PSByte := cpu.PopByteFromStack(&cycles)
            cpu.PS = cpu.ByteToPS(PSByte)

            cycles-=2

            // Total cycles: 4
            // Total bytes: 1
            break;

        case instructions.INS_JSR_ABS:

            // Example:
            // I read opcode at FF00. PC is now FF01
            // targetAddress := cpu.FetchWord(&cycles)
            // I read 00 at FF01 and 80 at FF02. PC is now FF03
            // I store PC - 1 = FF02 in the SP which is FD
            // Which means 02 at 01FD and FF at 01FC
            // PC is 8000 

            // Fetch the targetMemoryAddress, which is where we have to jump to
            targetAddress := cpu.FetchWord(&cycles)

            // This takes 2 cycles
            cpu.PushWordToStack(&cycles, cpu.PC-1)

            cpu.PC = targetAddress
            cycles--

            // Total cycles: 6
            // Total bytes: 3

            break;

        case instructions.INS_RTS_IMP:
            cpu.PC = cpu.PopWordFromStack(&cycles)

            // This is necessary since we want to Execute next instruction in the next loop iteration
            // If I don't increase the PC, it will run the same execution stored in the SP
            cpu.PC++

            // TODO: why is this necessary to reach a total of 6 cycles? What happens?

            // Total cycles: 6
            // Total bytes: 1
            cycles -=3
            break;
        case instructions.INS_JMP_ABS:

            cpu.PC = cpu.AddressAbsolute(&cycles)

            // Total cycles: 3
            // Total bytes: 3
            break;

        case instructions.INS_JMP_IND:

            targetAddress := cpu.AddressAbsolute(&cycles)

            cpu.PC = cpu.ReadWord(&cycles, targetAddress)

            // Total cycles: 5
            // Total bytes: 3
            break;

        case instructions.INS_AND_IM:

            cpu.A = cpu.A & cpu.FetchByte(&cycles)

            // Total cycles: 2
            // Total bytes: 2
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_AND_ZP:

            zeroPageAddress := cpu.AddressZeroPage(&cycles)

            cpu.A = cpu.A & cpu.ReadByte(&cycles, zeroPageAddress)

            // Total cycles: 3
            // Total bytes: 2
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_AND_ZPX:

            zeroPageAddress := cpu.AddressZeroPageX(&cycles)

            cpu.A = cpu.A & cpu.ReadByte(&cycles, zeroPageAddress)

            // Total cycles: 4
            // Total bytes: 2
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;
        
        case instructions.INS_AND_ABS:

            absoluteAddress := cpu.AddressAbsolute(&cycles)

            cpu.A = cpu.A & cpu.ReadByte(&cycles, absoluteAddress)

            // Total cycles: 4
            // Total bytes: 3
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_AND_ABSX:

            absoluteAddress := cpu.AddressAbsoluteX(&cycles)

            cpu.A = cpu.A & cpu.ReadByte(&cycles, absoluteAddress)

            // Total cycles: 4+1
            // Total bytes: 3
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_AND_ABSY:

            absoluteAddress := cpu.AddressAbsoluteY(&cycles)

            cpu.A = cpu.A & cpu.ReadByte(&cycles, absoluteAddress)

            // Total cycles: 4+1
            // Total bytes: 3
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_AND_INDX:

            effectiveAddress := cpu.AddressIndirectX(&cycles)

            cpu.A = cpu.A & cpu.ReadByte(&cycles, effectiveAddress)

            // Total cycles: 6
            // Total bytes: 2
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_AND_INDY:

            effectiveAddress := cpu.AddressIndirectY(&cycles)

            cpu.A = cpu.A & cpu.ReadByte(&cycles, effectiveAddress)

            // Total cycles: 5+1
            // Total bytes: 2
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_EOR_IM:

            cpu.A = cpu.A ^ cpu.FetchByte(&cycles)

            // Total cycles: 2
            // Total bytes: 2
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_EOR_ZP:

            zeroPageAddress := cpu.AddressZeroPage(&cycles)

            cpu.A = cpu.A ^ cpu.ReadByte(&cycles, zeroPageAddress)

            // Total cycles: 3
            // Total bytes: 2
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_EOR_ZPX:

            zeroPageAddress := cpu.AddressZeroPageX(&cycles)

            cpu.A = cpu.A ^ cpu.ReadByte(&cycles, zeroPageAddress)

            // Total cycles: 4
            // Total bytes: 2
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;
        
        case instructions.INS_EOR_ABS:

            absoluteAddress := cpu.AddressAbsolute(&cycles)

            cpu.A = cpu.A ^ cpu.ReadByte(&cycles, absoluteAddress)

            // Total cycles: 4
            // Total bytes: 3
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_EOR_ABSX:

            absoluteAddress := cpu.AddressAbsoluteX(&cycles)

            cpu.A = cpu.A ^ cpu.ReadByte(&cycles, absoluteAddress)

            // Total cycles: 4+1
            // Total bytes: 3
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_EOR_ABSY:

            absoluteAddress := cpu.AddressAbsoluteY(&cycles)

            cpu.A = cpu.A ^ cpu.ReadByte(&cycles, absoluteAddress)

            // Total cycles: 4+1
            // Total bytes: 3
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_EOR_INDX:

            effectiveAddress := cpu.AddressIndirectX(&cycles)

            cpu.A = cpu.A ^ cpu.ReadByte(&cycles, effectiveAddress)

            // Total cycles: 6
            // Total bytes: 2
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_EOR_INDY:

            effectiveAddress := cpu.AddressIndirectY(&cycles)

            cpu.A = cpu.A ^ cpu.ReadByte(&cycles, effectiveAddress)

            // Total cycles: 5+1
            // Total bytes: 2
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_ORA_IM:

            cpu.A = cpu.A | cpu.FetchByte(&cycles)

            // Total cycles: 2
            // Total bytes: 2
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_ORA_ZP:

            zeroPageAddress := cpu.AddressZeroPage(&cycles)

            cpu.A = cpu.A | cpu.ReadByte(&cycles, zeroPageAddress)

            // Total cycles: 3
            // Total bytes: 2
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_ORA_ZPX:

            zeroPageAddress := cpu.AddressZeroPageX(&cycles)

            cpu.A = cpu.A | cpu.ReadByte(&cycles, zeroPageAddress)

            // Total cycles: 4
            // Total bytes: 2
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;
        
        case instructions.INS_ORA_ABS:

            absoluteAddress := cpu.AddressAbsolute(&cycles)

            cpu.A = cpu.A | cpu.ReadByte(&cycles, absoluteAddress)

            // Total cycles: 4
            // Total bytes: 3
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_ORA_ABSX:

            absoluteAddress := cpu.AddressAbsoluteX(&cycles)

            cpu.A = cpu.A | cpu.ReadByte(&cycles, absoluteAddress)

            // Total cycles: 4+1
            // Total bytes: 3
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_ORA_ABSY:

            absoluteAddress := cpu.AddressAbsoluteY(&cycles)

            cpu.A = cpu.A | cpu.ReadByte(&cycles, absoluteAddress)

            // Total cycles: 4+1
            // Total bytes: 3
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_ORA_INDX:

            effectiveAddress := cpu.AddressIndirectX(&cycles)

            cpu.A = cpu.A | cpu.ReadByte(&cycles, effectiveAddress)

            // Total cycles: 6
            // Total bytes: 2
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;

        case instructions.INS_ORA_INDY:

            effectiveAddress := cpu.AddressIndirectY(&cycles)

            cpu.A = cpu.A | cpu.ReadByte(&cycles, effectiveAddress)

            // Total cycles: 5+1
            // Total bytes: 2
            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, cpu.A)
            break;
        
        case instructions.INS_BIT_ZP:

            zeropageAddress := cpu.AddressZeroPage(&cycles)

            memValue := cpu.ReadByte(&cycles, zeropageAddress)

            if (cpu.A & memValue) == 0 {
                cpu.PS.Z = 1
            }else {
                cpu.PS.Z = 0
            }

            // Flag V is set to bit 6 of the memory value
            cpu.PS.V = uint((memValue << 1) >> 7)


            // Flag N is set to bit 7 of the memory value
            cpu.PS.N = uint(memValue >> 7)

            // Total cycles: 3
            // Total bytes: 2
            break;
        case instructions.INS_BIT_ABS:

            targetAddress := cpu.AddressAbsolute(&cycles)

            memValue := cpu.ReadByte(&cycles, targetAddress)

            if (cpu.A & memValue) == 0 {
                cpu.PS.Z = 1
            }else {
                cpu.PS.Z = 0
            }

            // Flag V is set to bit 6 of the memory value
            cpu.PS.V = uint((memValue << 1) >> 7)


            // Flag N is set to bit 7 of the memory value
            cpu.PS.N = uint(memValue >> 7)

            // Total cycles: 4
            // Total bytes: 3
            break;
        case instructions.INS_INC_ZP:

            zeroPageAddress := cpu.AddressZeroPage(&cycles)

            memValue := cpu.ReadByte(&cycles, zeroPageAddress)

            memValue++

            cycles--

            cpu.WriteByte(&cycles, memValue, zeroPageAddress)

            SetZeroAndNegativeFlags(cpu, memValue)

            // Total cycles: 5
            // Total bytes: 2
            break;

        case instructions.INS_INC_ZPX:

            zeroPageAddress := cpu.AddressZeroPageX(&cycles)

            memValue := cpu.ReadByte(&cycles, zeroPageAddress)

            memValue++

            cycles--

            cpu.WriteByte(&cycles, memValue, zeroPageAddress)

            SetZeroAndNegativeFlags(cpu, memValue)

            // Total cycles: 6
            // Total bytes: 2
            break;

        case instructions.INS_INC_ABS:

            targetAddress := cpu.AddressAbsolute(&cycles)

            memValue := cpu.ReadByte(&cycles, targetAddress)

            memValue++

            cycles--

            cpu.WriteByte(&cycles, memValue, targetAddress)

            SetZeroAndNegativeFlags(cpu, memValue)

            // Total cycles: 6
            // Total bytes: 3
            break;

        case instructions.INS_INC_ABSX:

            targetAddress := cpu.AddressAbsolute(&cycles)

            targetAddress += uint16(cpu.X)
            cycles--

            memValue := cpu.ReadByte(&cycles, targetAddress)
            memValue++

            cycles--

            cpu.WriteByte(&cycles, memValue, targetAddress)

            SetZeroAndNegativeFlags(cpu, memValue)

            // Total cycles: 7
            // Total bytes: 3
            break;

        case instructions.INS_INX_IMP:

            cpu.X++
            cycles--

            SetZeroAndNegativeFlags(cpu, cpu.X)

            // Total cycles: 2
            // Total bytes: 1
            break;
        case instructions.INS_INY_IMP:

            cpu.Y += 1
            cycles--
            SetZeroAndNegativeFlags(cpu, cpu.Y)
            // Total cycles: 2
            // Total bytes: 1
            break;

        case instructions.INS_DEC_ZP:

            zeroPageAddress := cpu.AddressZeroPage(&cycles)

            memValue := cpu.ReadByte(&cycles, zeroPageAddress)

            memValue--

            cycles--

            cpu.WriteByte(&cycles, memValue, zeroPageAddress)

            SetZeroAndNegativeFlags(cpu, memValue)

            // Total cycles: 5
            // Total bytes: 2
            break;

        case instructions.INS_DEC_ZPX:

            zeroPageAddress := cpu.AddressZeroPageX(&cycles)

            memValue := cpu.ReadByte(&cycles, zeroPageAddress)

            memValue--

            cycles--

            cpu.WriteByte(&cycles, memValue, zeroPageAddress)

            SetZeroAndNegativeFlags(cpu, memValue)

            // Total cycles: 6
            // Total bytes: 2
            break;

        case instructions.INS_DEC_ABS:

            targetAddress := cpu.AddressAbsolute(&cycles)

            memValue := cpu.ReadByte(&cycles, targetAddress)

            memValue--

            cycles--

            cpu.WriteByte(&cycles, memValue, targetAddress)

            SetZeroAndNegativeFlags(cpu, memValue)

            // Total cycles: 6
            // Total bytes: 3
            break;

        case instructions.INS_DEC_ABSX:

            targetAddress := cpu.AddressAbsolute(&cycles)

            targetAddress += uint16(cpu.X)
            cycles--

            memValue := cpu.ReadByte(&cycles, targetAddress)
            memValue--

            cycles--

            cpu.WriteByte(&cycles, memValue, targetAddress)

            SetZeroAndNegativeFlags(cpu, memValue)

            // Total cycles: 7
            // Total bytes: 3
            break;

        case instructions.INS_DEX_IMP:

            cpu.X -= 1
            cycles--
            SetZeroAndNegativeFlags(cpu, cpu.X)
            // Total cycles: 2
            // Total bytes: 1
            break;
        case instructions.INS_DEY_IMP:

            cpu.Y -= 1
            cycles--
            SetZeroAndNegativeFlags(cpu, cpu.Y)
            // Total cycles: 2
            // Total bytes: 1
            break;

        case instructions.INS_BEQ_REL:

            signedOffset := cpu.FetchSignedByte(&cycles)

            // If zero flag is set, add signed int to program counter
            if cpu.PS.Z == 1{

                cycles--

                // Add 8 signed int to uint 16. How?
                // Cast int8 and uint16 to int16
                // Sum 
                // Cast back result to uint16

                
                // If original high byte is different from new high byte, there's been 
                // a page crossing. +1 cycles
                originalHi := cpu.PC >> 8

                cpu.PC = uint16(int16(signedOffset) + int16(cpu.PC))

                if originalHi != cpu.PC >> 8 {
                    cycles--    
                }

                // If to a new page, takes one cycle

            }

            // Total cycles: 2(+1 if branch succeeds, +2 if to a new page)
            // Total bytes: 2
            break;
        default:
            log.Println("At memory address: ", cpu.PC)

            // TODO: Should it stop and Fatal or just keep going till next valid instruction?
            log.Fatalln("Unknown opcode: ", ins)
        }
    }

    // If the number of cycles used is correct, respectively to the instruction used, 
    // the return should be the original value, passed when calling Execute().
    // When testing the instruction, we make sure that the expected value returned by Execute()
    // matches the cycles needed for the instructions, based on official documentation.
    cyclesUsed -= cycles

    return
}

func SetZeroAndNegativeFlags(cpu *CPU, register byte) {

            // Set Z flag if A is 0
            if register == 0 {
                cpu.PS.Z = 1
            }else{
                cpu.PS.Z = 0
            }

            // Set N flag if the bit 7 of A is set
            // byte(1 << 7) is a bitmask that has the 7 bit set to 1
            // it left-shifts the 00000001 seven positions left
            if (register & byte(1 << 7) != 0) {
                cpu.PS.N = 1
            }else {
                cpu.PS.N = 0
            }
}

func AddRegValueToTarget16Address(value byte, address *uint16, cycles *int){

    // "+1 if page crossed": a page boundary is crossed if the high byte of the original absolute address is
    // different from the high byte of the calculated address after adding the X register.
    // If a page boundary is crossed, an additional cycle is required.

    originalHighByte := (*address >> 8) // Add value to address *address += uint16(value)

    *address += uint16(value)

    newHighByte := (*address >> 8)


    if originalHighByte != newHighByte {
        *cycles--
    }
}

func LoadRegisterAndSetStatusFlags(cpu *CPU, cycles *int, address uint16, register *byte){

            // Load the data at the zeroPageAddress in the A register
            *register = cpu.ReadByte(cycles, address)

            // Set LDA status flags
            SetZeroAndNegativeFlags(cpu, *register)
}

func StoreRegistersIntoMemoryAddress(cpu *CPU, cycles *int, address uint16, register *byte){

            cpu.WriteByte(cycles, *register, address)
}

// Fetches ZeroPage Address when in Addressing Mode - Zero Page 
// Consumes 1 clock cycle
func (cpu *CPU) AddressZeroPage(cycles *int) uint16{

    return uint16(cpu.FetchByte(cycles))
}

// Fetches ZeroPage Address when in Addressing Mode - Zero Page with X offset
// Consumes 2 clock cycles
func (cpu *CPU) AddressZeroPageX(cycles *int) uint16{

            zeroPageAddress := cpu.FetchByte(cycles)

            // Wrap Around
            zeroPageAddress = byte(uint16(cpu.X + zeroPageAddress) % 0x100)
            *cycles--
            return uint16(zeroPageAddress)
}

// Fetches ZeroPage Address when in Addressing Mode - Zero Page with Y offset
// Consumes 2 clock cycles
func (cpu *CPU) AddressZeroPageY(cycles *int) uint16{

            zeroPageAddress := cpu.FetchByte(cycles)

            // Wrap Around
            zeroPageAddress = byte(uint16(cpu.Y + zeroPageAddress) % 0x100)
            *cycles--

            return uint16(zeroPageAddress)
}

// Fetches Absolute Address when in Addressing Mode - Absolute
// Consumes 2 clock cycles
func (cpu *CPU) AddressAbsolute(cycles *int) uint16{

            return uint16(cpu.FetchWord(cycles))
}

// Fetches Absolute Address when in Addressing Mode - Absolute with X offset
// Consumes 2(+1) clock cycles
func (cpu *CPU) AddressAbsoluteX(cycles *int) uint16{

            targetAddress := uint16(cpu.FetchWord(cycles))

            // Add X value to the fetched address
            AddRegValueToTarget16Address(cpu.X, &targetAddress, cycles)

            return targetAddress
}

// Fetches Absolute Address when in Addressing Mode - Absolute with Y offset
func (cpu *CPU) AddressAbsoluteY(cycles *int) uint16{

            targetAddress := uint16(cpu.FetchWord(cycles))

            // Add Y value to the fetched address
            AddRegValueToTarget16Address(cpu.Y, &targetAddress, cycles)

            return targetAddress
}

func (cpu *CPU) AddressIndirectX(cycles *int) uint16{

        // Fetch the Zero Page Address and add X
        zeroPageAddress := cpu.AddressZeroPageX(cycles)

        // The effective address is the 
        effectiveAddress := cpu.ReadWord(cycles, zeroPageAddress)

        return effectiveAddress
}

func (cpu *CPU) AddressIndirectY(cycles *int) uint16{

        // Fetch the Zero Page Address
        zeroPageAddress := cpu.FetchByte(cycles)

        effectiveAddress := cpu.ReadWord(cycles, uint16(zeroPageAddress))

        // Add Y to the Effective Address
        AddRegValueToTarget16Address(cpu.Y, &effectiveAddress, cycles)

        return effectiveAddress
}

func (cpu *CPU) PushByteToStack(cycles *int, data byte) {
    cpu.WriteByteToStack(cycles, data)
}

func (cpu *CPU) PushWordToStack(cycles *int, data uint16) {
    cpu.WriteWordToStack(cycles, data)
}

func (cpu *CPU) PopByteFromStack(cycles *int) (data byte){

    data = cpu.ReadByteFromStack(cycles)
    return
}

func (cpu *CPU) PopWordFromStack(cycles *int) (data uint16){

    data = cpu.ReadWordFromStack(cycles)
    return
}

func (cpu *CPU) SPTo16Address(sp byte) (SP uint16){
    SP = uint16(sp) + 0x0100
    return
}

func (cpu *CPU) PSToByte() (PS byte){

    // e.g.: 200 = 11001000
    // 00000000 | 00000000 | 00000000 | 00001000 | 00000000 | 00000000 | 01000000 | 10000000 = 11001000
    PS = byte(cpu.PS.C) | byte(cpu.PS.Z << 1) | byte(cpu.PS.I << 2) | byte(cpu.PS.D << 3) | byte(cpu.PS.B << 4) | byte(cpu.PS.U << 5) | byte(cpu.PS.V << 6) | byte(cpu.PS.N << 7)

    return
}

func (cpu *CPU) ByteToPS(bytePS byte) (ps ProcessorStatus){
    
    // This is super ugly but works for now
    ps.C = uint((bytePS << 7) >> 7)
    ps.Z = uint((bytePS << 6) >> 7)
    ps.I = uint((bytePS << 5) >> 7)
    ps.D = uint((bytePS << 4) >> 7)
    ps.B = uint((bytePS << 3) >> 7)
    ps.U = uint((bytePS << 2) >> 7)
    ps.V = uint((bytePS << 1) >> 7)
    ps.N = uint(bytePS >> 7)

    return
}

// LoadProgram loads an array of bytes (program) into memory
func (cpu *CPU) LoadProgram(program []byte){

    // first two bytes address the PC
    lo := uint16(program[0])
    hi := uint16(program[1]) << 8
    loadAddress := lo | hi

    // We set PC here, don't know if it's correct
    // TODO: find a place for setting program counter
    cpu.PC = loadAddress

    // Load following bytes into memory
    for i, byte := range program {

        if i > 1 {

            cpu.Memory.Data[loadAddress] = byte

            loadAddress++
        }
    }
}
