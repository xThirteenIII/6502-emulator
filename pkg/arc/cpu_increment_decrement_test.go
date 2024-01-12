package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

func TestINCZeroPageIncrementsTargetValueCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.PS.Z = 1
    cpu.PS.N = 1

    want := byte(0x35)

    cpu.Memory.Data[0xFFFC] = instructions.INS_INC_ZP
    cpu.Memory.Data[0xFFFD] = 0xD3
    cpu.Memory.Data[0x00D3] = 0x34

    cpuCopy := *cpu

    expectedCycles := 5
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.Memory.Data[0x00D3] != want {
        t.Error("Value at 0x00D3 should be ", want, "but got: ", cpu.Memory.Data[0x00D3])
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 0 but instead got 1")
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be 0 but instead got 1")
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestINCZeroPageXIncrementsTargetValueCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.PS.Z = 1
    cpu.PS.N = 1
    cpu.X = 0xC2

    want := byte(0x35)

    cpu.Memory.Data[0xFFFC] = instructions.INS_INC_ZPX
    cpu.Memory.Data[0xFFFD] = 0xD3
    cpu.Memory.Data[0x0095] = 0x34

    cpuCopy := *cpu

    expectedCycles := 6
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.Memory.Data[0x0095] != want {
        t.Error("Value at 0x0095 should be ", want, "but got: ", cpu.Memory.Data[0x0095])
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 0 but instead got 1")
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be 0 but instead got 1")
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestINCZeroPageAbsoluteIncrementsTargetValueCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.PS.Z = 1
    cpu.PS.N = 1

    want := byte(0x35)

    cpu.Memory.Data[0xFFFC] = instructions.INS_INC_ABS
    cpu.Memory.Data[0xFFFD] = 0xD3
    cpu.Memory.Data[0xFFFE] = 0x10
    cpu.Memory.Data[0x10D3] = 0x34

    cpuCopy := *cpu

    expectedCycles := 6
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.Memory.Data[0x10D3] != want {
        t.Error("Value at 0x10D3 should be ", want, "but got: ", cpu.Memory.Data[0x10D3])
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 0 but instead got 1")
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be 0 but instead got 1")
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}
