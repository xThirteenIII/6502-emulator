package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

func TestCLCClearsCarryFlagCorrectly(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.C = set

    cpu.Memory.Data[0xFFFC] = instructions.INS_CLC_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PS.C != cleared {
        t.Error("Carry flag should be clear instead is set")
    }
}

func TestCLDClearsDecimalFlagCorrectly(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.D = set

    cpu.Memory.Data[0xFFFC] = instructions.INS_CLD_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PS.D != cleared {
        t.Error("Decimal flag should be clear instead is set")
    }
}

// Clears the interrupt disable flag allowing normal interrupt requests to be serviced.
func TestCLIClearsInterruptDisableFlagCorrectly(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.I = set

    cpu.Memory.Data[0xFFFC] = instructions.INS_CLI_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PS.I != cleared {
        t.Error("Interrupt disable flag should be clear instead is set")
    }
}

func TestCLVClearsOverflowFlagCorrectly(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.V = set

    cpu.Memory.Data[0xFFFC] = instructions.INS_CLV_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PS.V != cleared {
        t.Error("Overflow flag should be clear instead is set")
    }
}

func TestSECSetCarryFlagCorrectly(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.C = cleared

    cpu.Memory.Data[0xFFFC] = instructions.INS_SEC_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PS.C != set {
        t.Error("Carry flag should be set instead is clear")
    }
}

func TestSEDSetDecimalFlagCorrectly(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.D = cleared

    cpu.Memory.Data[0xFFFC] = instructions.INS_SED_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PS.D != set {
        t.Error("Decimal flag should be set instead is clear")
    }
}

func TestSEISetDecimalFlagCorrectly(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.I = cleared

    cpu.Memory.Data[0xFFFC] = instructions.INS_SEI_IMP

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PS.I != set {
        t.Error("Interrupt disable flag should be set instead is clear")
    }
}
