package arc

import (
	"emulator/pkg/common"
	"emulator/pkg/instructions"
	"testing"
)

func TestSBCIMSubtractsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()
    cpu.PS.C = 1

    CheckSBCIMExecute(cpu, 0x00, 0x00, 0x00, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
}

func TestSBCIMSubtractsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()
    cpu.PS.C = 1

    CheckSBCIMExecute(cpu, 0x02, common.Int8ToByte(114), common.Int8ToByte(-112), t)

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

func TestSBCIMSubtractsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    cpu := Init6502()
    cpu.PS.C = 1

    CheckSBCIMExecute(cpu, 0x02, 0x01, 0x01, t)

    if cpu.PS.C != 1 {
        t.Error("Carry bit should be 1 but got ", cpu.PS.C)
    }

    if cpu.PS.V != 0 {
        t.Error("Overflow bit should be 0 but got ", cpu.PS.V)
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 1 but got ", cpu.PS.Z)
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be 0 but got ", cpu.PS.N)
    }
}

func TestSBCIMSubtractsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    // 127+1 = 128
    cpu := Init6502()
    cpu.PS.C = 1

    // 0x7F - (-1) - (1 - 1) = 0x80
    // 0x7F - (-1) - (1 - 0) = 0x7F
    CheckSBCIMExecute(cpu, 0x7F, common.Int8ToByte(-1), 0x80, t)

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

func TestSBCIMSubtractsCorrectlyWithPreviousCarryFlagClear(t *testing.T){

    // Given
    // -128-1-1 = +126
    cpu := Init6502()
    cpu.PS.C = 0
    CheckSBCIMExecute(cpu, common.Int8ToByte(-128), 0x01, 0x7E, t)

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
        t.Error("Negative flag should be 0 but got ", cpu.PS.N)
    }
}


func TestSBCZPSubtractsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()
    cpu.PS.C = 1

    CheckSBCZPExecute(cpu, 0x00, 0x00, 0x00, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
}

func TestSBCZPSubtractsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()
    cpu.PS.C = 1

    CheckSBCZPExecute(cpu, 0x02, common.Int8ToByte(114), common.Int8ToByte(-112), t)

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

func TestSBCZPSubtractsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    cpu := Init6502()
    cpu.PS.C = 1

    CheckSBCZPExecute(cpu, 0x02, 0x01, 0x01, t)

    if cpu.PS.C != 1 {
        t.Error("Carry bit should be 1 but got ", cpu.PS.C)
    }

    if cpu.PS.V != 0 {
        t.Error("Overflow bit should be 0 but got ", cpu.PS.V)
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 1 but got ", cpu.PS.Z)
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be 0 but got ", cpu.PS.N)
    }
}

func TestSBCZPSubtractsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    // 127+1 = 128
    cpu := Init6502()
    cpu.PS.C = 1

    // 0x7F - (-1) - (1 - 1) = 0x80
    // 0x7F - (-1) - (1 - 0) = 0x7F
    CheckSBCZPExecute(cpu, 0x7F, common.Int8ToByte(-1), 0x80, t)

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

func TestSBCZPSubtractsCorrectlyWithPreviousCarryFlagClear(t *testing.T){

    // Given
    // -128-1-1 = +126
    cpu := Init6502()
    cpu.PS.C = 0
    CheckSBCZPExecute(cpu, common.Int8ToByte(-128), 0x01, 0x7E, t)

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
        t.Error("Negative flag should be 0 but got ", cpu.PS.N)
    }
}

func TestSBCZPXSubtractsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()
    cpu.PS.C = 1

    CheckSBCZPXExecute(cpu, 0x00, 0x00, 0x00, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
}

func TestSBCZPXSubtractsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()
    cpu.PS.C = 1

    CheckSBCZPXExecute(cpu, 0x02, common.Int8ToByte(114), common.Int8ToByte(-112), t)

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

func TestSBCZPXSubtractsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    cpu := Init6502()
    cpu.PS.C = 1

    CheckSBCZPXExecute(cpu, 0x02, 0x01, 0x01, t)

    if cpu.PS.C != 1 {
        t.Error("Carry bit should be 1 but got ", cpu.PS.C)
    }

    if cpu.PS.V != 0 {
        t.Error("Overflow bit should be 0 but got ", cpu.PS.V)
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 1 but got ", cpu.PS.Z)
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be 0 but got ", cpu.PS.N)
    }
}

func TestSBCZPXSubtractsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    // 127+1 = 128
    cpu := Init6502()
    cpu.PS.C = 1

    // 0x7F - (-1) - (1 - 1) = 0x80
    // 0x7F - (-1) - (1 - 0) = 0x7F
    CheckSBCZPXExecute(cpu, 0x7F, common.Int8ToByte(-1), 0x80, t)

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

func TestSBCZPXSubtractsCorrectlyWithPreviousCarryFlagClear(t *testing.T){

    // Given
    // -128-1-1 = +126
    cpu := Init6502()
    cpu.PS.C = 0
    CheckSBCZPXExecute(cpu, common.Int8ToByte(-128), 0x01, 0x7E, t)

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
        t.Error("Negative flag should be 0 but got ", cpu.PS.N)
    }
}

func TestSBCABSSubtractsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()
    cpu.PS.C = 1

    CheckSBCABSExecute(cpu, 0x00, 0x00, 0x00, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
}

func TestSBCABSSubtractsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()
    cpu.PS.C = 1

    CheckSBCABSExecute(cpu, 0x02, common.Int8ToByte(114), common.Int8ToByte(-112), t)

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

func TestSBCABSSubtractsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    cpu := Init6502()
    cpu.PS.C = 1

    CheckSBCABSExecute(cpu, 0x02, 0x01, 0x01, t)

    if cpu.PS.C != 1 {
        t.Error("Carry bit should be 1 but got ", cpu.PS.C)
    }

    if cpu.PS.V != 0 {
        t.Error("Overflow bit should be 0 but got ", cpu.PS.V)
    }

    if cpu.PS.Z != 0 {
        t.Error("Zero flag should be 1 but got ", cpu.PS.Z)
    }

    if cpu.PS.N != 0 {
        t.Error("Negative flag should be 0 but got ", cpu.PS.N)
    }
}

func TestSBCABSSubtractsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    // 127+1 = 128
    cpu := Init6502()
    cpu.PS.C = 1

    // 0x7F - (-1) - (1 - 1) = 0x80
    // 0x7F - (-1) - (1 - 0) = 0x7F
    CheckSBCABSExecute(cpu, 0x7F, common.Int8ToByte(-1), 0x80, t)

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

func TestSBCABSSubtractsCorrectlyWithPreviousCarryFlagClear(t *testing.T){

    // Given
    // -128-1-1 = +126
    cpu := Init6502()
    cpu.PS.C = 0
    CheckSBCABSExecute(cpu, common.Int8ToByte(-128), 0x01, 0x7E, t)

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
        t.Error("Negative flag should be 0 but got ", cpu.PS.N)
    }
}
/*

func TestSBCABSXSubtractsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()

    CheckSBCABSXExecute(cpu, 0x00, 0x00, 0x00, 4, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
}

func TestSBCABSXSubtractsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()

    CheckSBCABSXExecute(cpu, 0x05, 0xF0, 0xF5, 4, t)

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

func TestSBCABSXSubtractsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    cpu := Init6502()
    CheckSBCABSXExecute(cpu, 0x05, 0xFB, 0x00, 4, t)

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

func TestSBCABSXSubtractsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    cpu := Init6502()
    CheckSBCABSXExecute(cpu, 0x7F, 0x01, 0x80, 4, t)

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
func TestSBCABSYSubtractsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()

    CheckSBCABSYExecute(cpu, 0x00, 0x00, 0x00, 4, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
}

func TestSBCABSYSubtractsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()

    CheckSBCABSYExecute(cpu, 0x05, 0xF0, 0xF5, 4, t)

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

func TestSBCABSYSubtractsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    cpu := Init6502()
    CheckSBCABSYExecute(cpu, 0x05, 0xFB, 0x00, 4, t)

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

func TestSBCABSYSubtractsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    cpu := Init6502()
    CheckSBCABSYExecute(cpu, 0x7F, 0x01, 0x80, 4, t)

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

func TestSBCINDXSubtractsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()

    CheckSBCINDXExecute(cpu, 0x00, 0x00, 0x00, 6, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
}

func TestSBCINDXSubtractsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()

    CheckSBCINDXExecute(cpu, 0x05, 0xF0, 0xF5, 6, t)

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

func TestSBCINDXSubtractsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    cpu := Init6502()
    CheckSBCINDXExecute(cpu, 0x05, 0xFB, 0x00, 6, t)

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

func TestSBCINDXSubtractsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    cpu := Init6502()
    CheckSBCINDXExecute(cpu, 0x7F, 0x01, 0x80, 6, t)

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
func TestSBCINDYSubtractsCorrectlyZeroToZero(t *testing.T){

    cpu := Init6502()

    CheckSBCINDYExecute(cpu, 0x00, 0x00, 0x00, 5, t)

    CheckIfFollowingFlagsAreSet(t, &cpu.PS.Z)
    CheckIfFollowingFlagsAreCleared(t, &cpu.PS.V, &cpu.PS.C, &cpu.PS.N)
}

func TestSBCINDYSubtractsCorrectlyWithNoCarryNorOverflow(t *testing.T){

    cpu := Init6502()

    CheckSBCINDYExecute(cpu, 0x05, 0xF0, 0xF5, 5, t)

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

func TestSBCINDYSubtractsCorrectlyWithCarryAndNoOverflow(t *testing.T){

    cpu := Init6502()
    CheckSBCINDYExecute(cpu, 0x05, 0xFB, 0x00, 5, t)

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

func TestSBCINDYSubtractsCorrectlyWithNoCarryAndOverflow(t *testing.T){

    // Given
    cpu := Init6502()
    CheckSBCINDYExecute(cpu, 0x7F, 0x01, 0x80, 5, t)

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

*/

// expectedResult: the value the accumulator should have after add operations,
// memValue: the value in the memory cell that is added to the accumulator,
// accumulator: initial value of the register,
// expectedCycles: the number of cycles expected from the instruction execution,
func CheckSBCIMExecute(cpu *CPU, accumulator, memValue, expectedResult byte, t *testing.T){

    // Given
    cpu.A = accumulator

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_SBC_IM
    cpu.Memory.Data[0xFFFD] = memValue

    expectedCycles := 2
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
func CheckSBCZPExecute(cpu *CPU, accumulator, memValue, expectedResult byte, t *testing.T){

    // Given
    cpu.A = accumulator

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_SBC_ZP
    cpu.Memory.Data[0xFFFD] = 0x4F
    cpu.Memory.Data[0x004F] = memValue

    expectedCycles := 3
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
func CheckSBCZPXExecute(cpu *CPU, accumulator, memValue, expectedResult byte, t *testing.T){

    // Given
    cpu.A = accumulator
    cpu.X = 0x04

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_SBC_ZPX
    cpu.Memory.Data[0xFFFD] = 0x4F
    cpu.Memory.Data[0x0053] = memValue

    expectedCycles := 4
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
func CheckSBCABSExecute(cpu *CPU, accumulator, memValue, expectedResult byte, t *testing.T){

    // Given
    cpu.A = accumulator

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_SBC_ABS
    cpu.Memory.Data[0xFFFD] = 0x4F
    cpu.Memory.Data[0xFFFE] = 0x50
    cpu.Memory.Data[0x504F] = memValue

    expectedCycles := 4
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
func CheckSBCABSXExecute(cpu *CPU, accumulator, memValue, expectedResult byte, expectedCycles int, t *testing.T){

    // Given
    cpu.A = accumulator
    cpu.X = 0x04

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_SBC_ABSX
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
func CheckSBCABSYExecute(cpu *CPU, accumulator, memValue, expectedResult byte, expectedCycles int, t *testing.T){

    // Given
    cpu.A = accumulator
    cpu.Y = 0x04

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_SBC_ABSY
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
func CheckSBCINDXExecute(cpu *CPU, accumulator, memValue, expectedResult byte, expectedCycles int, t *testing.T){

    // Given
    cpu.A = accumulator
    cpu.X = 0x04

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_SBC_INDX
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
func CheckSBCINDYExecute(cpu *CPU, accumulator, memValue, expectedResult byte, expectedCycles int, t *testing.T){

    // Given
    cpu.A = accumulator
    cpu.Y = 0x04

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_SBC_INDY
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
