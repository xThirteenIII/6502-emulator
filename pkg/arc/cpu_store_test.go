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

    cpuCopy := *cpu

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Memory.Data[0x0044] != *register {
        t.Error("Expected ", *register, "but got: ", cpu.Memory.Data[0x0044])
    }

    CheckUnmodifiedSTAFlags(cpuCopy, cpu, t)
}

func CheckStoreRegisterZeroPageX(cpu *CPU, opcode int, register *byte, t *testing.T){

    cpu.X = 0x02
    cpu.Memory.Data[0xFFFC] = byte(opcode)
    cpu.Memory.Data[0xFFFD] = 0x44

    cpuCopy := *cpu

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Memory.Data[0x0046] != *register {
        t.Error("Expected ", *register, "but got: ", cpu.Memory.Data[0x0044])
    }

    CheckUnmodifiedSTAFlags(cpuCopy, cpu, t)
}

func CheckStoreRegisterAbsolute(cpu *CPU, opcode int, register *byte, t *testing.T){

    cpu.Memory.Data[0xFFFC] = byte(opcode)
    cpu.Memory.Data[0xFFFD] = 0x44
    cpu.Memory.Data[0xFFFE] = 0x80

    cpuCopy := *cpu

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Memory.Data[0x8044] != *register {
        t.Error("Expected ", *register, "but got: ", cpu.Memory.Data[0x8044])
    }

    CheckUnmodifiedSTAFlags(cpuCopy, cpu, t)
    
}

func CheckStoreRegisterAbsoluteX(cpu *CPU, opcode int, register *byte, t *testing.T){

    cpu.X = 0x02
    cpu.Memory.Data[0xFFFC] = byte(opcode)
    cpu.Memory.Data[0xFFFD] = 0x44
    cpu.Memory.Data[0xFFFE] = 0x80

    cpuCopy := *cpu

    expectedCycles := 5
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Memory.Data[0x8046] != *register {
        t.Error("Expected ", *register, "but got: ", cpu.Memory.Data[0x8046])
    }
    
    CheckUnmodifiedSTAFlags(cpuCopy, cpu, t)
}

func CheckStoreRegisterAbsoluteY(cpu *CPU, opcode int, register *byte, t *testing.T){

    cpu.Y = 0x02
    cpu.Memory.Data[0xFFFC] = byte(opcode)
    cpu.Memory.Data[0xFFFD] = 0x44
    cpu.Memory.Data[0xFFFE] = 0x80
    cpuCopy := *cpu

    expectedCycles := 5
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Memory.Data[0x8046] != *register {
        t.Error("Expected ", *register, "but got: ", cpu.Memory.Data[0x8046])
    }
    
    CheckUnmodifiedSTAFlags(cpuCopy, cpu, t)
}

// Confront Initial PS Registers values with values after execution.
// These register shuould remain unmodified
func CheckUnmodifiedSTAFlags(cpuCopy CPU, cpu *CPU, t *testing.T){

    // Confront uneffected flags
    if cpu.PS.C != cpuCopy.PS.C {
        t.Error("PS.C: want: ", cpuCopy.PS.C, ", got: ", cpu.PS.C)
    }

    if cpu.PS.I != cpuCopy.PS.I {
        t.Error("PS.I: want: ", cpuCopy.PS.I, ", got: ", cpu.PS.I)
    }

    if cpu.PS.U != cpuCopy.PS.U {
        t.Error("PS.U: want: ", cpuCopy.PS.U, ", got: ", cpu.PS.U)
    }

    if cpu.PS.B != cpuCopy.PS.B {
        t.Error("PS.B: want: ", cpuCopy.PS.B, ", got: ", cpu.PS.B)
    }

    if cpu.PS.D != cpuCopy.PS.D {
        t.Error("PS.D: want: ", cpuCopy.PS.D, ", got: ", cpu.PS.D)
    }
    if cpu.PS.V != cpuCopy.PS.V {
        t.Error("PS.V: want: ", cpuCopy.PS.V, ", got: ", cpu.PS.V)
    }
    if cpu.PS.N != cpuCopy.PS.N {
        t.Error("PS.N: want: ", cpuCopy.PS.N, ", got: ", cpu.PS.N)
    }
    if cpu.PS.Z != cpuCopy.PS.Z {
        t.Error("PS.Z: want: ", cpuCopy.PS.Z, ", got: ", cpu.PS.Z)
    }
}
