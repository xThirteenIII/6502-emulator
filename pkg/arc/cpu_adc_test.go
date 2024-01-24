package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

func TestADCIMAddsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0x05
    cpu.PS.Z = 1
    cpu.PS.C = 0
    cpu.PS.V = 1

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ADC_IM
    cpu.Memory.Data[0xFFFD] = 0xF0

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    // Then
    if cpu.A != 0xF5 {
        t.Error("Accumulator should be 0xF5 but got: ", cpu.A)
    }

    if cpu.PS.C != 0 {
        t.Error("Carry bit should be 0 but got ", cpu.PS.C)
    }

    if cpu.PS.V != 0 {
        t.Error("Overflow bit should be 0 but got ", cpu.PS.V)
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 0 but got ", cpu.PS.Z)
    }

    if cpu.PS.N != 1 {
        t.Error("Negative flag should be 1 but got ", cpu.PS.N)
    }
}

func TestADCIMAddsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0x05
    cpu.PS.Z = 1
    cpu.PS.C = 0
    cpu.PS.V = 1

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ADC_IM
    cpu.Memory.Data[0xFFFD] = 0xFB

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    // Then
    if cpu.A != 0x00 {
        t.Error("Accumulator should be 0x00 but got: ", cpu.A)
    }

    if cpu.PS.C != 1 {
        t.Error("Carry bit should be 1 but got ", cpu.PS.C)
    }

    if cpu.PS.V != 0 {
        t.Error("Overflow bit should be 0 but got ", cpu.PS.V)
    }

    if cpu.PS.Z != 1 {
        t.Error("Zero flag should be 1 but got ", cpu.PS.Z)
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be 0 but got ", cpu.PS.N)
    }
}

func TestADCIMAddsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0x7F
    cpu.PS.Z = 0
    cpu.PS.C = 0
    cpu.PS.V = 0

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ADC_IM
    cpu.Memory.Data[0xFFFD] = 0x01

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    // Then
    if cpu.A != 0x80 {
        t.Error("Accumulator should be 0xE0 but got: ", cpu.A)
    }

    if cpu.PS.C != 0 {
        t.Error("Carry bit should be 0 but got ", cpu.PS.C)
    }

    if cpu.PS.V != 1 {
        t.Error("Overflow bit should be 1 but got ", cpu.PS.V)
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 0 but got ", cpu.PS.Z)
    }

    if cpu.PS.N != 1 {
        t.Error("Negative flag should be 1 but got ", cpu.PS.N)
    }
}
