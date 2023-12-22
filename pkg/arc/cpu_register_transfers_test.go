package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

func TestTAXCopiesAccumulatorIntoXCorrectly(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0x04
    cpuCopy := *cpu

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_TAX_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but god: ", cyclesUsed)
    }

    if cpu.PC != 0xFFFD {
        t.Error("Program counter should be 0xFFFD but got: ", cpu.PC)
    }

    if cpu.X != cpu.A {
        t.Error("Should have X: ", cpu.A, "but got: ",cpu.X)
    }

    // Check every register apart from Z and N rests unmodified
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestTAYCopiesAccumulatorIntoYCorrectly(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0x04
    cpuCopy := *cpu

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_TAY_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but god: ", cyclesUsed)
    }

    if cpu.PC != 0xFFFD {
        t.Error("Program counter should be 0xFFFD but got: ", cpu.PC)
    }

    if cpu.Y != cpu.A {
        t.Error("Should have Y: ", cpu.A, "but got: ",cpu.Y)
    }

    // Check every register apart from Z and N rests unmodified
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestTXACopiesXRegisterIntoAccumulatorCorrectly(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.X = 0x04
    cpuCopy := *cpu

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_TXA_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but god: ", cyclesUsed)
    }

    if cpu.PC != 0xFFFD {
        t.Error("Program counter should be 0xFFFD but got: ", cpu.PC)
    }

    if cpu.A != cpu.X {
        t.Error("Should have A: ", cpu.X, "but got: ",cpu.A)
    }

    // Check every register apart from Z and N rests unmodified
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestTYACopiesYRegisterIntoAccumulatorCorrectly(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.Y = 0x04
    cpuCopy := *cpu

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_TYA_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but god: ", cyclesUsed)
    }

    if cpu.PC != 0xFFFD {
        t.Error("Program counter should be 0xFFFD but got: ", cpu.PC)
    }

    if cpu.A != cpu.Y {
        t.Error("Should have A: ", cpu.Y, "but got: ",cpu.A)
    }

    // Check every register apart from Z and N rests unmodified
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}
