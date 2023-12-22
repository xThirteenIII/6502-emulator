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
