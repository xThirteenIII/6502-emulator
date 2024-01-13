package arc

import (
	"emulator/pkg/common"
	"emulator/pkg/instructions"
	"testing"
)

func TestBEQSumsCorrectlyToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1000)
    cpu.PS.Z = 1

    cpu.Memory.Data[0x1000] = instructions.INS_BEQ_REL
    cpu.Memory.Data[0x1001] = 0x33

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1035 {
        t.Error("PC should be 0x1035 instead got: ", cpu.PC)
    }
}

func TestBEQDoesNotModifyPCIfZeroFlagIsNotSet(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1000)
    cpu.PS.Z = 0

    cpu.Memory.Data[0x1000] = instructions.INS_BEQ_REL
    cpu.Memory.Data[0x1001] = 0x33

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1002 {
        t.Error("PC should be 0x1002 instead got: ", cpu.PC)
    }
}

func TestBEQSumsCorrectlyZeroToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1000)
    cpu.PS.Z = 1

    cpu.Memory.Data[0x1000] = instructions.INS_BEQ_REL
    cpu.Memory.Data[0x1001] = 0x00

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1002 {
        t.Error("PC should be 0x1002 instead got: ", cpu.PC)
    }
}

func TestBEQSumsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_BEQ_REL
    cpu.Memory.Data[0xFFFD] = 0x33

    // Page crossing happens so it takes 4 cycles
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // Since at this point PC is 0xFFFE, 0xFFFE+0x33 = 0x31 with wrap around
    if cpu.PC != 0x31 {
        t.Error("PC should be 0x31 instead got: ", cpu.PC)
    }
}

func TestBEQSubtractsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_BEQ_REL
    cpu.Memory.Data[0xFFFD] = common.Int8ToByte(-122)

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0xFF84 {
        t.Error("PC should be 0xFF84 instead got: ", cpu.PC)
    }
}
