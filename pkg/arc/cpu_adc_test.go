package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

func TestADCIMAddsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()

    CheckADCIMExecute(cpu, 0x00, 0x00, 0x00, 2, t)

    if cpu.PS.C != 0 {
        t.Error("Carry bit should be 0 but got ", cpu.PS.C)
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

func TestADCIMAddsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()

    CheckADCIMExecute(cpu, 0x05, 0xF0, 0xF5, 2, t)

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

    cpu := Init6502()
    CheckADCIMExecute(cpu, 0x05, 0xFB, 0x00, 2, t)

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
    CheckADCIMExecute(cpu, 0x7F, 0x01, 0x80, 2, t)

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

// expectedResult: the value the accumulator should have after add operations,
// memValue: the value in the memory cell that is added to the accumulator,
// accumulator: initial value of the register,
// expectedCycles: the number of cycles expected from the instruction execution,
func CheckADCIMExecute(cpu *CPU, accumulator, memValue, expectedResult byte, expectedCycles int, t *testing.T){

    // Given
    cpu.A = accumulator

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ADC_IM
    cpu.Memory.Data[0xFFFD] = memValue

    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    // Then
    if cpu.A != expectedResult {
        t.Error("Accumulator should be ", expectedResult, " but got: ", cpu.A)
    }
}

func CheckIfCarryFlagIs(expectedStatus uint, cpu CPU, t *testing.T){

    if cpu.PS.C != expectedStatus {
        t.Error("Carry bit should be ", expectedStatus, " but got ", cpu.PS.C)
    }
}

func CheckIfOverflowFlagIs(expectedStatus uint, cpu CPU, t *testing.T){
    if cpu.PS.V != expectedStatus {
        t.Error("Overflow bit should be ", expectedStatus, " but got ", cpu.PS.V)
    }
}

func CheckIfZeroFlagIs(expectedStatus uint, cpu CPU, t *testing.T){
    if cpu.PS.Z != expectedStatus {
        t.Error("Zero bit should be ", expectedStatus, " but got ", cpu.PS.Z)
    }
}

func CheckIfNegativeFlagIs(expectedStatus uint, cpu CPU, t *testing.T){
    if cpu.PS.N != expectedStatus {
        t.Error("Negative bit should be ", expectedStatus, " but got ", cpu.PS.N)
    }
}
