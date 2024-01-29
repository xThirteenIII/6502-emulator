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
        t.Error("Accumulator should be 10 but got", cpu.A)
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
        t.Error("Accumulator should be 0 but got", cpu.A)
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
        t.Error("Accumulator should be 0xEF but got", cpu.A)
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
        t.Error("Accumulator should be 10 but got", cpu.A)
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
        t.Error("Accumulator should be 0 but got", cpu.A)
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
        t.Error("Accumulator should be 0xEF but got", cpu.A)
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
        t.Error("Accumulator should be 10 but got", cpu.A)
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
        t.Error("Accumulator should be 0 but got", cpu.A)
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
        t.Error("Accumulator should be 0xEF but got", cpu.A)
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

func TestCMPABSSetsCarryFlagCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_ABS
    cpu.Memory.Data[0xFFFD] = 0x80
    cpu.Memory.Data[0xFFFE] = 0x80
    cpu.Memory.Data[0x8080] = 0x20

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0x10 {
        t.Error("Accumulator should be 10 but got", cpu.A)
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

func TestCMPABSSetsCarryAndZeroFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0x30
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_ABS
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0xFFFE] = 0x50
    cpu.Memory.Data[0x5050] = 0x30

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0x00 {
        t.Error("Accumulator should be 0 but got", cpu.A)
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

func TestCMPABSSetsCarryAndNegativeFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0xF0
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_ABS
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0xFFFE] = 0x50
    cpu.Memory.Data[0x5050] = 0x01

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0xEF {
        t.Error("Accumulator should be 0xEF but got", cpu.A)
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

func TestCMPABSXSetsCarryFlagCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0x30
    cpu.X = 0x10
    cpu.PS.C = 0
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_ABSX
    cpu.Memory.Data[0xFFFD] = 0xFF
    cpu.Memory.Data[0xFFFE] = 0x20
    cpu.Memory.Data[0x210F] = 0x20

    expectedCycles := 5
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0x10 {
        t.Error("Accumulator should be 0x10 but got", cpu.A)
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

func TestCMPABSXSetsCarryAndZeroFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0x30
    cpu.X = 0x05
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_ABSX
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0xFFFE] = 0x50
    cpu.Memory.Data[0x5055] = 0x30

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0x00 {
        t.Error("Accumulator should be 0 but got", cpu.A)
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

func TestCMPABSXSetsCarryAndNegativeFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0xF0
    cpu.X = 0x05
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_ABSX
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0xFFFE] = 0x50
    cpu.Memory.Data[0x5055] = 0x01

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0xEF {
        t.Error("Accumulator should be 0xEF but got", cpu.A)
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

func TestCMPABSYSetsCarryFlagCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0x30
    cpu.Y = 0x10
    cpu.PS.C = 0
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_ABSY
    cpu.Memory.Data[0xFFFD] = 0xFF
    cpu.Memory.Data[0xFFFE] = 0x20
    cpu.Memory.Data[0x210F] = 0x20

    expectedCycles := 5
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0x10 {
        t.Error("Accumulator should be 0x10 but got", cpu.A)
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

func TestCMPABSYSetsCarryAndZeroFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0x30
    cpu.Y = 0x05
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_ABSY
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0xFFFE] = 0x50
    cpu.Memory.Data[0x5055] = 0x30

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0x00 {
        t.Error("Accumulator should be 0 but got", cpu.A)
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

func TestCMPABSYSetsCarryAndNegativeFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0xF0
    cpu.Y = 0x05
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_ABSY
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0xFFFE] = 0x50
    cpu.Memory.Data[0x5055] = 0x01

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0xEF {
        t.Error("Accumulator should be 0xEF but got", cpu.A)
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

func TestCMPINDXSetsCarryFlagCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0x30
    cpu.X = 0x04
    cpu.PS.C = 0
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_INDX
    cpu.Memory.Data[0xFFFD] = 0x4F
    cpu.Memory.Data[0x0053] = 0x60
    cpu.Memory.Data[0x0054] = 0x60
    cpu.Memory.Data[0x6060] = 0x20

    expectedCycles := 6
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0x10 {
        t.Error("Accumulator should be 0x10 but got", cpu.A)
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

func TestCMPINDXSetsCarryAndZeroFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0x30
    cpu.X = 0x05
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_INDX
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0x0055] = 0x50
    cpu.Memory.Data[0x0056] = 0x30
    cpu.Memory.Data[0x3050] = 0x30

    expectedCycles := 6
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0x00 {
        t.Error("Accumulator should be 0 but got", cpu.A)
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

func TestCMPINDXSetsCarryAndNegativeFlagsCorrectly(t *testing.T){

    cpu := Init6502()
    cpu.A = 0xF0
    cpu.X = 0x05
    cpu.PS.C = 0
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_CMP_INDX
    cpu.Memory.Data[0xFFFD] = 0x50
    cpu.Memory.Data[0x0055] = 0x50
    cpu.Memory.Data[0x0056] = 0x61
    cpu.Memory.Data[0x6150] = 0x01

    expectedCycles := 6
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    if cpu.A != 0xEF {
        t.Error("Accumulator should be 0xEF but got", cpu.A)
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
