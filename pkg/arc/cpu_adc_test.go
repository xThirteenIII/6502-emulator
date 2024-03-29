package arc

import (
	"emulator/pkg/common"
	"emulator/pkg/instructions"
	"testing"
)

func TestADCIMAddsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()

    CheckADCIMExecute(cpu, 0x00, 0x00, 0x00, 2, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
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
    // 127+1 = 128
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

func TestADCIMAddsCorrectlySignedWithCarryAndOverflow(t *testing.T){

    // Given
    // -128-1 = +127
    cpu := Init6502()
    CheckADCIMExecute(cpu, common.Int8ToByte(-128), common.Int8ToByte(-1), common.Int8ToByte(127), 2, t)

    if cpu.PS.C != 1 {
        t.Error("Carry bit should be 0 but got ", cpu.PS.C)
    }

    if cpu.PS.V != 1 {
        t.Error("Overflow bit should be 1 but got ", cpu.PS.V)
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 0 but got ", cpu.PS.Z)
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be 1 but got ", cpu.PS.N)
    }
}

func TestADCIMAddsCorrectlySignedWithCarryAndOverflowWithPreviousCarryFlagSet(t *testing.T){

    // Given
    // -128-1+1 = +128
    cpu := Init6502()
    cpu.PS.C = 1
    CheckADCIMExecute(cpu, common.Int8ToByte(-128), common.Int8ToByte(-1), 0x80, 2, t)

    if cpu.PS.C != 1 {
        t.Error("Carry bit should be 0 but got ", cpu.PS.C)
    }

    if cpu.PS.V != 0 {
        t.Error("Overflow bit should be 1 but got ", cpu.PS.V)
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 0 but got ", cpu.PS.Z)
    }

    if cpu.PS.N != 1 {
        t.Error("Negative flag should be 1 but got ", cpu.PS.N)
    }
}

func TestADCIMAddsCorrectlySignedWithCarryAndNoOverflow(t *testing.T){

    // Given
    // -127-1 = -128
    cpu := Init6502()
    CheckADCIMExecute(cpu, common.Int8ToByte(-127), common.Int8ToByte(-1), 0x80, 2, t)

    if cpu.PS.C != 1 {
        t.Error("Carry bit should be 0 but got ", cpu.PS.C)
    }

    if cpu.PS.V != 0 {
        t.Error("Overflow bit should be 1 but got ", cpu.PS.V)
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 0 but got ", cpu.PS.Z)
    }

    if cpu.PS.N != 1 {
        t.Error("Negative flag should be 1 but got ", cpu.PS.N)
    }
}

func TestADCZPAddsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()

    CheckADCZPExecute(cpu, 0x00, 0x00, 0x00, 3, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
}

func TestADCZPAddsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()

    CheckADCZPExecute(cpu, 0x05, 0xF0, 0xF5, 3, t)

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

func TestADCZPAddsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    cpu := Init6502()
    CheckADCZPExecute(cpu, 0x05, 0xFB, 0x00, 3, t)

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

func TestADCZPAddsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    cpu := Init6502()
    CheckADCZPExecute(cpu, 0x7F, 0x01, 0x80, 3, t)

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


func TestADCZPXAddsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()

    CheckADCZPXExecute(cpu, 0x00, 0x00, 0x00, 4, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
}

func TestADCZPXAddsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()

    CheckADCZPXExecute(cpu, 0x05, 0xF0, 0xF5, 4, t)

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

func TestADCZPXAddsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    cpu := Init6502()
    CheckADCZPXExecute(cpu, 0x05, 0xFB, 0x00, 4, t)

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

func TestADCZPXAddsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    cpu := Init6502()
    CheckADCZPXExecute(cpu, 0x7F, 0x01, 0x80, 4, t)

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

func TestADCABSAddsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()

    CheckADCABSExecute(cpu, 0x00, 0x00, 0x00, 4, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
}

func TestADCABSAddsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()

    CheckADCABSExecute(cpu, 0x05, 0xF0, 0xF5, 4, t)

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

func TestADCABSAddsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    cpu := Init6502()
    CheckADCABSExecute(cpu, 0x05, 0xFB, 0x00, 4, t)

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

func TestADCABSAddsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    cpu := Init6502()
    CheckADCABSExecute(cpu, 0x7F, 0x01, 0x80, 4, t)

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

func TestADCABSXAddsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()

    CheckADCABSXExecute(cpu, 0x00, 0x00, 0x00, 4, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
}

func TestADCABSXAddsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()

    CheckADCABSXExecute(cpu, 0x05, 0xF0, 0xF5, 4, t)

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

func TestADCABSXAddsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    cpu := Init6502()
    CheckADCABSXExecute(cpu, 0x05, 0xFB, 0x00, 4, t)

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

func TestADCABSXAddsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    cpu := Init6502()
    CheckADCABSXExecute(cpu, 0x7F, 0x01, 0x80, 4, t)

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
func TestADCABSYAddsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()

    CheckADCABSYExecute(cpu, 0x00, 0x00, 0x00, 4, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
}

func TestADCABSYAddsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()

    CheckADCABSYExecute(cpu, 0x05, 0xF0, 0xF5, 4, t)

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

func TestADCABSYAddsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    cpu := Init6502()
    CheckADCABSYExecute(cpu, 0x05, 0xFB, 0x00, 4, t)

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

func TestADCABSYAddsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    cpu := Init6502()
    CheckADCABSYExecute(cpu, 0x7F, 0x01, 0x80, 4, t)

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

func TestADCINDXAddsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()

    CheckADCINDXExecute(cpu, 0x00, 0x00, 0x00, 6, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
}

func TestADCINDXAddsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()

    CheckADCINDXExecute(cpu, 0x05, 0xF0, 0xF5, 6, t)

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

func TestADCINDXAddsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    cpu := Init6502()
    CheckADCINDXExecute(cpu, 0x05, 0xFB, 0x00, 6, t)

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

func TestADCINDXAddsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    cpu := Init6502()
    CheckADCINDXExecute(cpu, 0x7F, 0x01, 0x80, 6, t)

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
func TestADCINDYAddsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()

    CheckADCINDYExecute(cpu, 0x00, 0x00, 0x00, 5, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
}

func TestADCINDYAddsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()

    CheckADCINDYExecute(cpu, 0x05, 0xF0, 0xF5, 5, t)

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

func TestADCINDYAddsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    cpu := Init6502()
    CheckADCINDYExecute(cpu, 0x05, 0xFB, 0x00, 5, t)

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

func TestADCINDYAddsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    cpu := Init6502()
    CheckADCINDYExecute(cpu, 0x7F, 0x01, 0x80, 5, t)

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

// expectedResult: the value the accumulator should have after add operations,
// memValue: the value in the memory cell that is added to the accumulator,
// accumulator: initial value of the register,
// expectedCycles: the number of cycles expected from the instruction execution,
func CheckADCZPExecute(cpu *CPU, accumulator, memValue, expectedResult byte, expectedCycles int, t *testing.T){

    // Given
    cpu.A = accumulator

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ADC_ZP
    cpu.Memory.Data[0xFFFD] = 0x4F
    cpu.Memory.Data[0x004F] = memValue

    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    // Then
    if cpu.A != expectedResult {
        t.Error("Accumulator should be ", expectedResult, " but got: ", cpu.A)
    }
}

// expectedResult: the value the accumulator should have after add operations,
// memValue: the value in the memory cell that is added to the accumulator,
// accumulator: initial value of the register,
// expectedCycles: the number of cycles expected from the instruction execution,
func CheckADCZPXExecute(cpu *CPU, accumulator, memValue, expectedResult byte, expectedCycles int, t *testing.T){

    // Given
    cpu.A = accumulator
    cpu.X = 0x04

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ADC_ZPX
    cpu.Memory.Data[0xFFFD] = 0x4F
    cpu.Memory.Data[0x0053] = memValue

    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    // Then
    if cpu.A != expectedResult {
        t.Error("Accumulator should be ", expectedResult, " but got: ", cpu.A)
    }
}

// expectedResult: the value the accumulator should have after add operations,
// memValue: the value in the memory cell that is added to the accumulator,
// accumulator: initial value of the register,
// expectedCycles: the number of cycles expected from the instruction execution,
func CheckADCABSExecute(cpu *CPU, accumulator, memValue, expectedResult byte, expectedCycles int, t *testing.T){

    // Given
    cpu.A = accumulator

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ADC_ABS
    cpu.Memory.Data[0xFFFD] = 0x4F
    cpu.Memory.Data[0xFFFE] = 0x50
    cpu.Memory.Data[0x504F] = memValue

    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    // Then
    if cpu.A != expectedResult {
        t.Error("Accumulator should be ", expectedResult, " but got: ", cpu.A)
    }
}

// expectedResult: the value the accumulator should have after add operations,
// memValue: the value in the memory cell that is added to the accumulator,
// accumulator: initial value of the register,
// expectedCycles: the number of cycles expected from the instruction execution,
func CheckADCABSXExecute(cpu *CPU, accumulator, memValue, expectedResult byte, expectedCycles int, t *testing.T){

    // Given
    cpu.A = accumulator
    cpu.X = 0x04

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ADC_ABSX
    cpu.Memory.Data[0xFFFD] = 0x4F
    cpu.Memory.Data[0xFFFE] = 0x50
    cpu.Memory.Data[0x5053] = memValue

    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    // Then
    if cpu.A != expectedResult {
        t.Error("Accumulator should be ", expectedResult, " but got: ", cpu.A)
    }
}

// expectedResult: the value the accumulator should have after add operations,
// memValue: the value in the memory cell that is added to the accumulator,
// accumulator: initial value of the register,
// expectedCycles: the number of cycles expected from the instruction execution,
func CheckADCABSYExecute(cpu *CPU, accumulator, memValue, expectedResult byte, expectedCycles int, t *testing.T){

    // Given
    cpu.A = accumulator
    cpu.Y = 0x04

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ADC_ABSY
    cpu.Memory.Data[0xFFFD] = 0x4F
    cpu.Memory.Data[0xFFFE] = 0x50
    cpu.Memory.Data[0x5053] = memValue

    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    // Then
    if cpu.A != expectedResult {
        t.Error("Accumulator should be ", expectedResult, " but got: ", cpu.A)
    }
}

// expectedResult: the value the accumulator should have after add operations,
// memValue: the value in the memory cell that is added to the accumulator,
// accumulator: initial value of the register,
// expectedCycles: the number of cycles expected from the instruction execution,
func CheckADCINDXExecute(cpu *CPU, accumulator, memValue, expectedResult byte, expectedCycles int, t *testing.T){

    // Given
    cpu.A = accumulator
    cpu.X = 0x04

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ADC_INDX
    cpu.Memory.Data[0xFFFD] = 0x4F
    cpu.Memory.Data[0x0053] = 0x60
    cpu.Memory.Data[0x0054] = 0x60
    cpu.Memory.Data[0x6060] = memValue

    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", cyclesUsed, ", instead expected: ", expectedCycles)
    }

    // Then
    if cpu.A != expectedResult {
        t.Error("Accumulator should be ", expectedResult, " but got: ", cpu.A)
    }
}
// expectedResult: the value the accumulator should have after add operations,
// memValue: the value in the memory cell that is added to the accumulator,
// accumulator: initial value of the register,
// expectedCycles: the number of cycles expected from the instruction execution,
func CheckADCINDYExecute(cpu *CPU, accumulator, memValue, expectedResult byte, expectedCycles int, t *testing.T){

    // Given
    cpu.A = accumulator
    cpu.Y = 0x04

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ADC_INDY
    cpu.Memory.Data[0xFFFD] = 0x4F
    cpu.Memory.Data[0x004F] = 0x50
    cpu.Memory.Data[0x0050] = 0x50
    cpu.Memory.Data[0x5054] = memValue

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
