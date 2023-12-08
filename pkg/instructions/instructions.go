package instructions

import (
	"emulator/pkg/arc"
	"fmt"
	"log"
)

// Opcodes
const (
    INS_LDA_IM = 0xA9

    // An instruction using zero page addressing mode has only an 8 bit address operand.
    // This limits it to addressing only the first 256 bytes of memory (e.g. $0000 to $00FF) where the most significant byte of the address is always zero.
    // In zero page mode only the least significant byte of the address is held in the instruction making it shorter by one byte
    // (important for space saving) and one less memory fetch during execution (important for speed).
    INS_LDA_ZP = 0xA5
)

// Instruction is an interface which is implemented by GenericInstruction,
// thus being able to call its methods without worrying about the specific instruction.
type Instruction interface {
    execute(cpu *arc.CPU)
    write(cpu *arc.CPU)
    decode(cpu *arc.CPU)
    // This is useful to manage similar operations in different address modes within the same instruction.
    // E.g: All LDA addressing modes instructions set the N and A flags in the same way.
    setStatus(cpu *arc.CPU)
}

// GenericInstruction holds the constant values for each instruction
type GenericInstruction struct {
    Opcode byte
    Bytes int
    Cycles int
}


// execute is a method that executes instruction-specific operations based on the opcode.
func (ins *GenericInstruction) execute(cpu  *arc.CPU){

    for (ins.Cycles > 0){

        switch (ins.Opcode){

        // LDA loads a byte of memory into the accumulator setting the zero and negative flags as appropriate.
        case INS_LDA_IM:

            // Load value into A
            // Is it ok to modify directly the cycles in the ins variable? What if i need it after?
            var err error
            cpu.A , err = cpu.FetchByte( &ins.Cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ", err.Error())
            }

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
            break;
        case INS_LDA_ZP:

            // First byte is the ZeroPage address
            // Second byte is the value to load

            // First cycle to fetch the instruction
            // Second cycle to fetch the address
            // The third cycle to read the data from the address
            zeroPageAddress, err := cpu.FetchByte(&ins.Cycles)
            if err != nil {
                fmt.Println("Error while fetching byte: ",err.Error())
            }
            
            // Load the data at the zeroPageAddress in the A register
            cpu.A = cpu.ReadByte(&ins.Cycles, uint16(zeroPageAddress))
            fmt.Println(cpu.Memory.Data[zeroPageAddress])

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
            break;
        default:
            log.Fatalln("Unknown opcode: ", ins.Opcode)
        }

    }

}

func (ins *GenericInstruction) decode(cpu  *arc.CPU){
    fmt.Println("decoding...")
}

func (ins *GenericInstruction) write(cpu  *arc.CPU){
    fmt.Println("writing...")
}

func (ins *GenericInstruction) setStatus(cpu  *arc.CPU){

}

/*
// next goes to the next instruction updating the PC
func (ins *GenericInstruction) next(cpu  *arc.CPU){

    // Increment the PC based on the instruction size
    cpu.PC += uint16(ins.Bytes)
}*/

// InstructionFactory is a function type to create instances of instructions, it returns a Instruction type.
// This is useful to abstract the creation of instructions and allows flexibility.
type InstructionFactory func() Instruction

// Instructions is a map that associates opcodes to an Instruction instance. 
// Here are defined Opcode, Bytes and Cycles for each instruction.
var Instructions = map[byte]InstructionFactory{
    INS_LDA_IM: func() Instruction { return &GenericInstruction{Opcode: INS_LDA_IM, Bytes: 2, Cycles: 2}},
    INS_LDA_ZP: func() Instruction { return &GenericInstruction{Opcode: INS_LDA_ZP, Bytes: 2, Cycles: 3}},
}


// ExecuteInstruction is a function to execute an instruction based on its opcode
func ExecuteInstruction(opcode byte, cpu *arc.CPU) {
    if factory, ok := Instructions[opcode]; ok {

        // If we find the opcode in the Instructions map, we create a GenericInstruction instance 
        // which has the Opcode, Bytes and Cycles set in the map.
        instruction := factory()

        instruction.execute(cpu)
        instruction.decode(cpu)
        instruction.write(cpu)
    } else {
        fmt.Printf("Unknown opcode: %X\n", opcode)
    }
}


