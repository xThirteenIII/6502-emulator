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

func TestINCZeroPageAbsoluteXIncrementsTargetValueCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.PS.Z = 1
    cpu.PS.N = 1
    cpu.X = 0xC2

    want := byte(0x35)

    cpu.Memory.Data[0xFFFC] = instructions.INS_INC_ABSX
    cpu.Memory.Data[0xFFFD] = 0xD3
    cpu.Memory.Data[0xFFFE] = 0x10
    cpu.Memory.Data[0x1195] = 0x34

    cpuCopy := *cpu

    expectedCycles := 7
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.Memory.Data[0x1195] != want {
        t.Error("Value at 0x1195 should be ", want, "but got: ", cpu.Memory.Data[0x1195])
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 0 but instead got 1")
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be 0 but instead got 1")
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestINXIncrementsXRegisterCorrectly(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.Z = 1
    cpu.PS.N = 1
    cpu.X = 0x44
    cpuCopy := *cpu

    cpu.Memory.Data[0xFFFC] = instructions.INS_INX_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.X != 0x45 {
        t.Error("Expected X to be 0x45 instead got: ", cpu.X)
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 0 but instead got 1")
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be 0 but instead got 1")
    }
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestINYIncrementsXRegisterCorrectly(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.Z = 1
    cpu.PS.N = 1
    cpu.Y = 0x33
    cpuCopy := *cpu

    cpu.Memory.Data[0xFFFC] = instructions.INS_INY_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.Y != 0x34 {
        t.Error("Expected Y to be 0x34 instead got: ", cpu.Y)
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 0 but instead got 1")
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be 0 but instead got 1")
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestDECZeroPageDecrementsTargetValueCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.PS.Z = 1
    cpu.PS.N = 1

    want := byte(0x33)

    cpu.Memory.Data[0xFFFC] = instructions.INS_DEC_ZP
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

func TestDECZeroPageXDecrementsTargetValueCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.PS.Z = 1
    cpu.PS.N = 1
    cpu.X = 0xC2

    want := byte(0x33)

    cpu.Memory.Data[0xFFFC] = instructions.INS_DEC_ZPX
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

func TestDECZeroPageAbsoluteDecrementsTargetValueCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.PS.Z = 1
    cpu.PS.N = 1

    want := byte(0x33)

    cpu.Memory.Data[0xFFFC] = instructions.INS_DEC_ABS
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

func TestDECZeroPageAbsoluteXDecrementsTargetValueCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.PS.Z = 1
    cpu.PS.N = 1
    cpu.X = 0xC2

    want := byte(0x33)

    cpu.Memory.Data[0xFFFC] = instructions.INS_DEC_ABSX
    cpu.Memory.Data[0xFFFD] = 0xD3
    cpu.Memory.Data[0xFFFE] = 0x10
    cpu.Memory.Data[0x1195] = 0x34

    cpuCopy := *cpu

    expectedCycles := 7
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.Memory.Data[0x1195] != want {
        t.Error("Value at 0x1195 should be ", want, "but got: ", cpu.Memory.Data[0x1195])
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 0 but instead got 1")
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be 0 but instead got 1")
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestDEXDecrementsXRegisterCorrectly(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.Z = 1
    cpu.PS.N = 1
    cpu.X = 0x44
    cpuCopy := *cpu

    cpu.Memory.Data[0xFFFC] = instructions.INS_DEX_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.X != 0x43 {
        t.Error("Expected X to be 0x43 instead got: ", cpu.X)
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 0 but instead got 1")
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be 0 but instead got 1")
    }
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestDEYDecrementsXRegisterCorrectly(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.Z = 1
    cpu.PS.N = 1
    cpu.Y = 0x33
    cpuCopy := *cpu

    cpu.Memory.Data[0xFFFC] = instructions.INS_DEY_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.Y != 0x32 {
        t.Error("Expected Y to be 0x32 instead got: ", cpu.Y)
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 0 but instead got 1")
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be 0 but instead got 1")
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}
