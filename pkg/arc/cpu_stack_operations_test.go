package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

func TestTSXCopiesStackRegisterIntoXCorrectly(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.SP = 0x04

    cpuCopy := *cpu

    // This should be set to 0 by execution
    cpu.PS.Z = 1
    cpu.PS.N = 1

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_TSX_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but god: ", cyclesUsed)
    }

    if cpu.PC != 0xFFFD {
        t.Error("Program counter should be 0xFFFD but got: ", cpu.PC)
    }

    if cpu.X != cpu.SP {
        t.Error("Should have X: ", cpu.SP, "but got: ",cpu.X)
    }

    // Check every register apart from Z and N rests unmodified
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestTXSCopiesXRegisterIntoStackPointerCorrectly(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.X = 0x04
    cpuCopy := *cpu

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_TXS_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0xFFFD {
        t.Error("Program counter should be 0xFFFD but got: ", cpu.PC)
    }

    if cpu.SP != cpu.X {
        t.Error("Should have SP: ", cpu.X, "but got: ",cpu.SP)
    }

    // Check every register rests unmodified
    CheckUnmodifiedlagsALL(cpuCopy, cpu, t)
}

func TestPHAPushesAccumulatorIntoStackCorrectly(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0x04
    cpuCopy := *cpu

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_PHA_IMP

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0xFFFD {
        t.Error("Program counter should be 0xFFFD but got: ", cpu.PC)
    }

    if cpu.Memory.Data[0x01FD] != cpu.A {
        t.Error("Stack pointer should contain: ", cpu.A, "but got: ", cpu.Memory.Data[0x01FD])
    }

    // Check every register rests unmodified
    CheckUnmodifiedlagsALL(cpuCopy, cpu, t)
}

func TestPHPPushesProcessorStatusIntoStackCorrectly(t *testing.T){

    // Given
    cpu := Init6502()
    // PS = 200
    cpu.PS = cpu.ByteToPS(200)
    cpuCopy := *cpu

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_PHP_IMP

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0xFFFD {
        t.Error("Program counter should be 0xFFFD but got: ", cpu.PC)
    }

    if cpu.Memory.Data[0x01FD] != cpu.PSToByte() {
        t.Error("Stack pointer should contain: ", cpu.PSToByte(), "but got: ", cpu.Memory.Data[0x01FD])
    }

    // Check every register rests unmodified
    CheckUnmodifiedlagsALL(cpuCopy, cpu, t)
}

func TestPLAPopsAccumulatorFromStackCorrectly(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.PS.Z = 1
    cpu.PS.N = 1
    cpu.A = 0x0F
    cpuCopy := *cpu

    // When
    // TODO: Add pha instruction to push into stack first
    cpu.Memory.Data[0xFFFC] = instructions.INS_PHA_IMP
    cpu.Memory.Data[0xFFFD] = instructions.INS_PLA_IMP

    expectedCycles := 3+4
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0xFFFE {
        t.Error("Program counter should be 0xFFFE but got: ", cpu.PC)
    }

    if cpu.Memory.Data[0x01FD] != cpu.A {
        t.Error("Accumulator should contain: ", cpu.Memory.Data[0x01FD], "but got: ", cpu.A)
    }

    // Check every register rests unmodified
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestPLPPopsAccumulatorFromStackCorrectly(t *testing.T){

    // Given
    cpu := Init6502()
    cpuCopy := *cpu
    cpu.PS.C = 1
    cpu.PS.Z = 1
    cpu.PS.B = 1

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_PHP_IMP
    cpu.Memory.Data[0xFFFD] = instructions.INS_PLP_IMP

    expectedCycles := 3+4
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0xFFFE {
        t.Error("Program counter should be 0xFFFE but got: ", cpu.PC)
    }

    if cpu.Memory.Data[0x01FD] != cpu.PSToByte() {
        t.Error("PS should be: ", cpu.Memory.Data[0x01FD], "but got: ", cpu.PSToByte())
    }

    if cpu.PS == cpuCopy.PS {
        t.Error("Processor stack read from stack pointer shoulnd't be the same as the original one: ", cpu.PS)
    }

}

