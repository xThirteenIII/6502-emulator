package instructions

import (
	"6502_emulator/pkg/arc"
	"log"
)

// Opcodes
const (
    INS_LDA_IM = 0xA9
)

type Instruction interface {
    execute(cpu *arc.CPU)
    write(cpu *arc.CPU)
    decode(cpu *arc.CPU)
    next(cpu *arc.CPU)
}

type GenericInstruction struct {
    Opcode byte
    Bytes int
    Cycles int
}

func (ins *GenericInstruction) execute(memory *arc.Memory, cpu  *arc.CPU){

    for (ins.Cycles > 0){

        switch (ins.Opcode){

        case INS_LDA_IM:

            // Load value into A
            cpu.A = arc.FetchByte( &ins.Cycles, memory)

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

func (ins *GenericInstruction) next(memory *arc.Memory, cpu  *arc.CPU){

    // Increment the PC based on the instruction size
    cpu.PC += uint16(ins.Bytes)
}
