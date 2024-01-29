package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

func TestCMYIMSetsCarryFlagCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.Y = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMY_IM
    cpu.Memory.Data[0xFFFD] = 0x20

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Y != 0x30 {
        t.Error("Accumulator should be 30 but got", cpu.Y)
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

func TestCMYIMSetsCarryAndZeroFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.Y = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMY_IM
    cpu.Memory.Data[0xFFFD] = 0x30

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Y != 0x30 {
        t.Error("Accumulator should be 30 but got", cpu.Y)
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

func TestCMYIMSetsCarryAndNegativeFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.Y = 0xF0
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMY_IM
    cpu.Memory.Data[0xFFFD] = 0x01

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Y != 0xF0 {
        t.Error("Accumulator should be 0xF0 but got", cpu.Y)
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

func TestCMYZPSetsCarryFlagCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.Y = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMY_ZP
    cpu.Memory.Data[0xFFFD] = 0x20
    cpu.Memory.Data[0x0020] = 0x20

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Y != 0x30 {
        t.Error("Accumulator should be 30 but got", cpu.Y)
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

func TestCMYZPSetsCarryAndZeroFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.Y = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMY_ZP
    cpu.Memory.Data[0xFFFD] = 0x20
    cpu.Memory.Data[0x0020] = 0x30

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Y != 0x30 {
        t.Error("Accumulator should be 30 but got", cpu.Y)
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

func TestCMYZPSetsCarryAndNegativeFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.Y = 0xF0
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMY_ZP
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0x0050] = 0x01

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Y != 0xF0 {
        t.Error("Accumulator should be 0xF0 but got", cpu.Y)
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

func TestCMYABSSetsCarryFlagCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.Y = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMY_ABS
    cpu.Memory.Data[0xFFFD] = 0x80
    cpu.Memory.Data[0xFFFE] = 0x80
    cpu.Memory.Data[0x8080] = 0x20

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Y != 0x30 {
        t.Error("Accumulator should be 30 but got", cpu.Y)
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

func TestCMPYBSSetsCarryAndZeroFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.Y = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMY_ABS
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0xFFFE] = 0x50
    cpu.Memory.Data[0x5050] = 0x30

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Y != 0x30 {
        t.Error("Accumulator should be 30 but got", cpu.Y)
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

func TestCMPYBSSetsCarryAndNegativeFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.Y = 0xF0
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMY_ABS
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0xFFFE] = 0x50
    cpu.Memory.Data[0x5050] = 0x01

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.Y != 0xF0 {
        t.Error("Accumulator should be 0xF0 but got", cpu.Y)
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
