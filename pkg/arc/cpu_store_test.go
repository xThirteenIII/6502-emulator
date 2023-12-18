package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

func TestSTAStoresValueIntoTargetAddress(t *testing.T){
    
    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    cpu.A = 0x09

    CheckStoreRegisterZeroPage(cpu, instructions.INS_STA_ZP, &cpu.A, t)

}

func TestSTAZeroPageXStoresValueIntoTargetAddress(t *testing.T){
    
    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    cpu.A = 0x09

    CheckStoreRegisterZeroPageX(cpu, instructions.INS_STA_ZPX, &cpu.A, t)

}

func TestSTAAbsoluteStoresValueIntoTargetAddress(t *testing.T){
    
    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    cpu.A = 0x09

    CheckStoreRegisterAbsolute(cpu, instructions.INS_STA_ABS, &cpu.A, t)

}

func CheckStoreRegisterZeroPage(cpu *CPU, opcode int, register *byte, t *testing.T){

    cpu.Memory.Data[0xFFFC] = byte(opcode)
    cpu.Memory.Data[0xFFFD] = 0x44

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Memory.Data[0x0044] != *register {
        t.Error("Expected ", *register, "but got: ", cpu.Memory.Data[0x0044])
    }
}

func CheckStoreRegisterZeroPageX(cpu *CPU, opcode int, register *byte, t *testing.T){

    cpu.X = 0x02
    cpu.Memory.Data[0xFFFC] = byte(opcode)
    cpu.Memory.Data[0xFFFD] = 0x44

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Memory.Data[0x0046] != *register {
        t.Error("Expected ", *register, "but got: ", cpu.Memory.Data[0x0044])
    }
}

func CheckStoreRegisterAbsolute(cpu *CPU, opcode int, register *byte, t *testing.T){

    cpu.Memory.Data[0xFFFC] = byte(opcode)
    cpu.Memory.Data[0xFFFD] = 0x44
    cpu.Memory.Data[0xFFFE] = 0x80

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Memory.Data[0x8044] != *register {
        t.Error("Expected ", *register, "but got: ", cpu.Memory.Data[0x8044])
    }
    
}
