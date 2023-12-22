package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

func TestJumpToSubroutineAndComeBack(t *testing.T){

    want := byte(0x33)
    // Given
    cpu := Init6502()

    // Need a lower resetVector address to run instructions
    cpu.Reset(0xFF00)
    cpuCopy := *cpu

    // When
    cpu.Memory.Data[0xFF00] = instructions.INS_JSR_ABS
    cpu.Memory.Data[0xFF01] = 0x00
    cpu.Memory.Data[0xFF02] = 0x80
    cpu.Memory.Data[0x8000] = instructions.INS_RTS_IMP
    cpu.Memory.Data[0xFF03] = instructions.INS_LDA_IM
    cpu.Memory.Data[0xFF04] = want

    expectedCycles := 6+6+2
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != want{
        t.Error("Want A: ", want, " but got: ", cpu.A)
    }

    // Since we pushed and popped same amount of bytes from the stack
    // the stack pointer should not change
    if cpu.SP != cpuCopy.SP {
        t.Error("The stack pointer SP shouldn't have changed. Want ", cpuCopy.SP, " but got: ", cpu.SP)
    }
}

func TestJSRDoesNotAffectProcessorStatus(t *testing.T){

    // Given
    cpu := Init6502()

    // Need a lower resetVector address to run instructions
    cpu.Reset(0xFF00)
    cpuCopy := *cpu

    // When
    cpu.Memory.Data[0xFF00] = instructions.INS_JSR_ABS
    cpu.Memory.Data[0xFF01] = 0x00
    cpu.Memory.Data[0xFF02] = 0x80
    cpu.Memory.Data[0x8000] = instructions.INS_RTS_IMP

    expectedCycles := 6+6
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    CheckUnmodifiedlagsALL(cpuCopy, cpu, t)

}

func TestRTSDoesNotAffectProcessorStatus(t *testing.T){

    // Given
    cpu := Init6502()

    // Need a lower resetVector address to run instructions
    cpu.Reset(0xFF00)
    cpuCopy := *cpu

    // When
    cpu.Memory.Data[0xFF00] = instructions.INS_JSR_ABS
    cpu.Memory.Data[0xFF01] = 0x00
    cpu.Memory.Data[0xFF02] = 0x80

    expectedCycles := 6
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    CheckUnmodifiedlagsALL(cpuCopy, cpu, t)

}

func TestJMPAbsoluteJumpsCorrectlyToNewAddress(t *testing.T){

    // Given
    cpu := Init6502()
    cpuCopy := *cpu

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_JMP_ABS
    cpu.Memory.Data[0xFFFD] = 0x00
    cpu.Memory.Data[0xFFFE] = 0x80

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.PC != 0x8000 {
        t.Error("PC should be 0x8000, but got: ", cpu.PC)
    }

    CheckUnmodifiedlagsALL(cpuCopy, cpu, t)
}

func TestJMPIndirectJumpsCorrectlyToNewAddress(t *testing.T){

    // Given
    cpu := Init6502()
    cpuCopy := *cpu

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_JMP_IND
    cpu.Memory.Data[0xFFFD] = 0x00
    cpu.Memory.Data[0xFFFE] = 0x80
    cpu.Memory.Data[0x8000] = 0x33
    cpu.Memory.Data[0x8001] = 0x44

    expectedCycles := 5
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.PC != 0x4433 {
        t.Error("PC should be 0x4433, but got: ", cpu.PC)
    }

    CheckUnmodifiedlagsALL(cpuCopy, cpu, t)
}

func TestJSRAbsolute(t *testing.T){

    cpu := Init6502()

    cpu.Memory.Data[0xFFFC] = instructions.INS_JSR_ABS
    cpu.Memory.Data[0xFFFD] = 0x00
    cpu.Memory.Data[0xFFFE] = 0x80
    // After reset SP = 0xFD
    // cpu.Memory.Data[0x01FD] = 0xFC
    // cpu.Memory.Data[0x01FC] = 0xFF

    expectedCycles := 6
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x8000 {
        t.Error("PC should be 8000 but got", cpu.PC)
    }
}
