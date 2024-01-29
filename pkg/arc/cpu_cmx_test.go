package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

func TestCMXIMSetsCarryFlagCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.X = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMX_IM
    cpu.Memory.Data[0xFFFD] = 0x20

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.X != 0x30 {
        t.Error("Accumulator should be 30 but got", cpu.X)
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

func TestCMXIMSetsCarryAndZeroFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.X = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMX_IM
    cpu.Memory.Data[0xFFFD] = 0x30

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.X != 0x30 {
        t.Error("Accumulator should be 30 but got", cpu.X)
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

func TestCMXIMSetsCarryAndNegativeFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.X = 0xF0
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMX_IM
    cpu.Memory.Data[0xFFFD] = 0x01

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.X != 0xF0 {
        t.Error("Accumulator should be 0xF0 but got", cpu.X)
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

func TestCMXZPSetsCarryFlagCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.X = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMX_ZP
    cpu.Memory.Data[0xFFFD] = 0x20
    cpu.Memory.Data[0x0020] = 0x20

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.X != 0x30 {
        t.Error("Accumulator should be 30 but got", cpu.X)
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

func TestCMXZPSetsCarryAndZeroFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.X = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMX_ZP
    cpu.Memory.Data[0xFFFD] = 0x20
    cpu.Memory.Data[0x0020] = 0x30

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.X != 0x30 {
        t.Error("Accumulator should be 30 but got", cpu.X)
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

func TestCMXZPSetsCarryAndNegativeFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.X = 0xF0
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMX_ZP
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0x0050] = 0x01

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.X != 0xF0 {
        t.Error("Accumulator should be 0xF0 but got", cpu.X)
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

func TestCMXABSSetsCarryFlagCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.X = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMX_ABS
    cpu.Memory.Data[0xFFFD] = 0x80
    cpu.Memory.Data[0xFFFE] = 0x80
    cpu.Memory.Data[0x8080] = 0x20

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.X != 0x30 {
        t.Error("Accumulator should be 30 but got", cpu.X)
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

func TestCMPXBSSetsCarryAndZeroFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.X = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMX_ABS
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0xFFFE] = 0x50
    cpu.Memory.Data[0x5050] = 0x30

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.X != 0x30 {
        t.Error("Accumulator should be 30 but got", cpu.X)
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

func TestCMPXBSSetsCarryAndNegativeFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.X = 0xF0
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMX_ABS
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0xFFFE] = 0x50
    cpu.Memory.Data[0x5050] = 0x01

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.X != 0xF0 {
        t.Error("Accumulator should be 0xF0 but got", cpu.X)
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
