package arc

import (
	"emulator/pkg/instructions"
    "testing"
)

func TestNOPDoesNotAffectProcessorStatus(t *testing.T){

    cpu := Init6502()
    cpu.PS.C = set
    cpu.PS.N = set
    cpu.PS.V = set

    cpu.Memory.Data[0xFFFC] = instructions.INS_CLC_IMP
    cpu.Memory.Data[0xFFFD] = instructions.INS_CLV_IMP
    cpu.Memory.Data[0xFFFE] = instructions.INS_NOP_IMP

    expectedCycles := 2+2+2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PS.C != cleared {
        t.Error("Carry flag should be clear instead is still set")
    }

    if cpu.PS.V != cleared {
        t.Error("Overflow flag should be clear instead is still set")
    }

    if cpu.PS.N != set {
        t.Error("Negative flag should be set instead is clear")
    }

    if cpu.PS.I != cleared {
        t.Error("Interrupt disable flag should be clear instead is set")
    }

    if cpu.PS.D != cleared {
        t.Error("Decimal disable flag should be clear instead is set")
    }

    if cpu.PS.Z != cleared {
        t.Error("Zero flag should be clear instead is set")
    }

    if cpu.PS.B != cleared {
        t.Error("Break command flag should be clear instead is set")
    }
}

