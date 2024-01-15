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
    cpu.Reset(0xFEFC)
    cpu.PS.Z = 1


    cpu.Memory.Data[0xFEFC] = instructions.INS_BEQ_REL
    cpu.Memory.Data[0xFEFD] = 0x33

    // Page crossing happens so it takes 4 cycles
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // Since at this point PC is 0xFFFE, 0xFFFE+0x33 = 0x31 with wrap around
    if cpu.PC != 0xFF31 {
        t.Error("PC should be 0xFF31 instead got: ", cpu.PC)
    }
}

func TestBEQSubtractsCorrectlyToProgramCounterWithoutPageCrossing(t *testing.T){
    
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

func TestBEQSubtractsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFF0C)
    cpu.PS.Z = 1

    cpu.Memory.Data[0xFF0C] = instructions.INS_BEQ_REL
    cpu.Memory.Data[0xFF0D] = common.Int8ToByte(-122)

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0xFE94 {
        t.Error("PC should be 0xFE94 instead got: ", cpu.PC)
    }
}

func TestBEQWorksWithAssembleProgram(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFF0C)
    cpu.PS.Z = 1

    /*
    loop
    lda #0
    beq loop
    */

    cpu.Memory.Data[0xFF0C] = instructions.INS_LDA_IM
    cpu.Memory.Data[0xFF0D] = 0x00
    cpu.Memory.Data[0xFF0E] = 0xF0
    cpu.Memory.Data[0xFF0F] = 0xFC // this goes backwards 4 bytes

    expectedCycles := 3+2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0xFF0C {
        t.Error("PC should be 0xFF0C instead got: ", cpu.PC)
    }
}

func TestBNESumsCorrectlyToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1000)
    cpu.PS.Z = 0

    cpu.Memory.Data[0x1000] = instructions.INS_BNE_REL
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

func TestBNEDoesNotModifyPCIfZeroFlagIsSet(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1000)
    cpu.PS.Z = 1

    cpu.Memory.Data[0x1000] = instructions.INS_BNE_REL
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

func TestBNESumsCorrectlyZeroToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1000)
    cpu.PS.Z = 0

    cpu.Memory.Data[0x1000] = instructions.INS_BNE_REL
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

func TestBNESumsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFEFC)
    cpu.PS.Z = 0


    cpu.Memory.Data[0xFEFC] = instructions.INS_BNE_REL
    cpu.Memory.Data[0xFEFD] = 0x33

    // Page crossing happens so it takes 4 cycles
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // Since at this point PC is 0xFFFE, 0xFFFE+0x33 = 0x31 with wrap around
    if cpu.PC != 0xFF31 {
        t.Error("PC should be 0xFF31 instead got: ", cpu.PC)
    }
}

func TestBNESubtractsCorrectlyToProgramCounterWithoutPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_BNE_REL
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

func TestBNESubtractsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFF0C)
    cpu.PS.Z = 0

    cpu.Memory.Data[0xFF0C] = instructions.INS_BNE_REL
    cpu.Memory.Data[0xFF0D] = common.Int8ToByte(-122)

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0xFE94 {
        t.Error("PC should be 0xFE94 instead got: ", cpu.PC)
    }
}

func TestBNEWorksWithAssembleProgram(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFF0C)
    cpu.PS.Z = 0

    /*
    loop
    lda #2
    bne loop
    */

    cpu.Memory.Data[0xFF0C] = instructions.INS_LDA_IM
    cpu.Memory.Data[0xFF0D] = 0x02
    cpu.Memory.Data[0xFF0E] = 0xD0
    cpu.Memory.Data[0xFF0F] = 0xFC // this goes backwards 4 bytes

    expectedCycles := 3+2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0xFF0C {
        t.Error("PC should be 0xFF0C instead got: ", cpu.PC)
    }
}
