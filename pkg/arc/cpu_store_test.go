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

func TestSTAAbsoluteXStoresValueIntoTargetAddress(t *testing.T){
    
    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    cpu.A = 0x09

    CheckStoreRegisterAbsoluteX(cpu, instructions.INS_STA_ABSX, &cpu.A, t)

}

func TestSTAAbsoluteYStoresValueIntoTargetAddress(t *testing.T){
    
    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    cpu.A = 0x09

    CheckStoreRegisterAbsoluteY(cpu, instructions.INS_STA_ABSY, &cpu.A, t)

}

func TestSTAIndirectXCanLoadIntoARegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    // given
    cpu.X = 0x04

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_STA_INDX
    cpu.Memory.Data[0xFFFD] = 0x02
    cpu.Memory.Data[0x0006] = 0x00
    cpu.Memory.Data[0x0007] = 0x80
    // end - inline program

    expectedCycles := 6
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles )
    }

    if cpu.Memory.Data[0x8000] != cpu.A {
        t.Error("Want ", cpu.A, "instead got: ", cpu.Memory.Data[0x8000])
    }
}

// Test if the STA_INDY instruction loads a value succefully into the A register
func TestSTAIndirectYCanLoadIntoARegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    // given
    cpu.Y = 0x04

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_STA_INDY
    cpu.Memory.Data[0xFFFD] = 0x02
    cpu.Memory.Data[0x0002] = 0x00
    cpu.Memory.Data[0x0003] = 0x80
    // end - inline program

    expectedCycles := 6
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Memory.Data[0x8004] != cpu.A {
        t.Error("Want ", cpu.A, "instead got: ", cpu.Memory.Data[0x8004])
    }
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

func CheckStoreRegisterAbsoluteX(cpu *CPU, opcode int, register *byte, t *testing.T){

    cpu.X = 0x02
    cpu.Memory.Data[0xFFFC] = byte(opcode)
    cpu.Memory.Data[0xFFFD] = 0x44
    cpu.Memory.Data[0xFFFE] = 0x80

    expectedCycles := 5
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Memory.Data[0x8046] != *register {
        t.Error("Expected ", *register, "but got: ", cpu.Memory.Data[0x8046])
    }
    
}

func CheckStoreRegisterAbsoluteY(cpu *CPU, opcode int, register *byte, t *testing.T){

    cpu.Y = 0x02
    cpu.Memory.Data[0xFFFC] = byte(opcode)
    cpu.Memory.Data[0xFFFD] = 0x44
    cpu.Memory.Data[0xFFFE] = 0x80

    expectedCycles := 5
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Memory.Data[0x8046] != *register {
        t.Error("Expected ", *register, "but got: ", cpu.Memory.Data[0x8046])
    }
    
}
