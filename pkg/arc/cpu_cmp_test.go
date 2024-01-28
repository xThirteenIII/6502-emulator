package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

func TestCMPIMSetsCarryFlagCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_IM
    cpu.Memory.Data[0xFFFD] = 0x20

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0x10 {
        t.Error("Accumulator should be", cpu.A - 0x20, "but got", cpu.A)
    }

    if cpu.PS.C != 1 {
        t.Error("Carry flag should be set")
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be clear")
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be clear")
    }
}

func TestCMPIMSetsCarryAndZeroFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_IM
    cpu.Memory.Data[0xFFFD] = 0x30

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0x00 {
        t.Error("Accumulator should be", cpu.A - 0x30, "but got", cpu.A)
    }

    if cpu.PS.C != 1 {
        t.Error("Carry flag should be set")
    }

    if cpu.PS.Z != 1 {
        t.Error("Zero flag should be set")
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be clear")
    }
}

func TestCMPIMSetsCarryAndNegativeFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0xF0
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_IM
    cpu.Memory.Data[0xFFFD] = 0x01

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0xEF {
        t.Error("Accumulator should be", cpu.A - 0x01, "but got", cpu.A)
    }

    if cpu.PS.C != 1 {
        t.Error("Carry flag should be set")
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be clear")
    }

    if cpu.PS.N != 1 {
        t.Error("Negative flag should be set")
    }
}

func TestCMPZPSetsCarryFlagCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_ZP
    cpu.Memory.Data[0xFFFD] = 0x20
    cpu.Memory.Data[0x0020] = 0x20

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0x10 {
        t.Error("Accumulator should be", cpu.A - 0x20, "but got", cpu.A)
    }

    if cpu.PS.C != 1 {
        t.Error("Carry flag should be set")
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be clear")
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be clear")
    }
}

func TestCMPZPSetsCarryAndZeroFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_ZP
    cpu.Memory.Data[0xFFFD] = 0x20
    cpu.Memory.Data[0x0020] = 0x30

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0x00 {
        t.Error("Accumulator should be", cpu.A - 0x30, "but got", cpu.A)
    }

    if cpu.PS.C != 1 {
        t.Error("Carry flag should be set")
    }

    if cpu.PS.Z != 1 {
        t.Error("Zero flag should be set")
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be clear")
    }
}

func TestCMPZPSetsCarryAndNegativeFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0xF0
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_ZP
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0x0050] = 0x01

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0xEF {
        t.Error("Accumulator should be", cpu.A - 0x01, "but got", cpu.A)
    }

    if cpu.PS.C != 1 {
        t.Error("Carry flag should be set")
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be clear")
    }

    if cpu.PS.N != 1 {
        t.Error("Negative flag should be set")
    }
}

func TestCMPZPXSetsCarryFlagCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0x30
    cpu.X = 0x10
    cpu.PS.C = 0
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_ZPX
    cpu.Memory.Data[0xFFFD] = 0x20
    cpu.Memory.Data[0x0030] = 0x20

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0x10 {
        t.Error("Accumulator should be", cpu.A - 0x20, "but got", cpu.A)
    }

    if cpu.PS.C != 1 {
        t.Error("Carry flag should be set")
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be clear")
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be clear")
    }
}

func TestCMPZPXSetsCarryAndZeroFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0x30
    cpu.X = 0x05
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_ZPX
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0x0055] = 0x30

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0x00 {
        t.Error("Accumulator should be", cpu.A - 0x30, "but got", cpu.A)
    }

    if cpu.PS.C != 1 {
        t.Error("Carry flag should be set")
    }

    if cpu.PS.Z != 1 {
        t.Error("Zero flag should be set")
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be clear")
    }
}

func TestCMPZPXSetsCarryAndNegativeFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0xF0
    cpu.X = 0x05
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_ZPX
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0x0055] = 0x01

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0xEF {
        t.Error("Accumulator should be", cpu.A - 0x01, "but got", cpu.A)
    }

    if cpu.PS.C != 1 {
        t.Error("Carry flag should be set")
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be clear")
    }

    if cpu.PS.N != 1 {
        t.Error("Negative flag should be set")
    }
}
