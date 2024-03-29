package arc

import (
	"emulator/pkg/common"
	"emulator/pkg/instructions"
	"testing"
)

func TestBEQSumsCorrectlyToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1000)
    cpu.X = 1

    cpu.Memory.Data[0x1000] = instructions.INS_DEX_IMP
    cpu.Memory.Data[0x1001] = instructions.INS_BEQ_REL
    cpu.Memory.Data[0x1002] = 0x33

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 2+3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1036 {
        t.Error("PC should be 0x1036 instead got: ", cpu.PC)
    }
}

func TestBEQDoesNotModifyPCIfZeroFlagIsClear(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1000)
    cpu.X = 0

    cpu.Memory.Data[0x1000] = instructions.INS_INX_IMP
    cpu.Memory.Data[0x1001] = instructions.INS_BEQ_REL
    cpu.Memory.Data[0x1002] = 0x33

    expectedCycles := 2+2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBEQSumsCorrectlyZeroToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1000)
    cpu.X = 1

    cpu.Memory.Data[0x1000] = instructions.INS_DEX_IMP
    cpu.Memory.Data[0x1001] = instructions.INS_BEQ_REL
    cpu.Memory.Data[0x1002] = 0x00

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 2+3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBEQSumsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFEFB)
    cpu.X = 1

    cpu.Memory.Data[0xFEFB] = instructions.INS_DEX_IMP
    cpu.Memory.Data[0xFEFC] = instructions.INS_BEQ_REL
    cpu.Memory.Data[0xFEFD] = 0x33

    // Page crossing happens so it takes 4 cycles
    expectedCycles := 2+4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // Since at this point PC is 0xFEFE, 0xFEFE+0x33 = 0xFF31 with wrap around
    if cpu.PC != 0xFF31 {
        t.Error("PC should be 0xFF31 instead got: ", cpu.PC)
    }
}

func TestBEQSubtractsCorrectlyToProgramCounterWithoutPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.X = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_DEX_IMP
    cpu.Memory.Data[0xFFFD] = instructions.INS_BEQ_REL
    cpu.Memory.Data[0xFFFE] = common.Int8ToByte(-122)


    expectedCycles := 2+3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0xFF85 {
        t.Error("PC should be 0xFF85 instead got: ", cpu.PC)
    }
}

func TestBEQSubtractsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFF0C)
    cpu.X = 1

    cpu.Memory.Data[0xFF0C] = instructions.INS_DEX_IMP
    cpu.Memory.Data[0xFF0D] = instructions.INS_BEQ_REL    
    cpu.Memory.Data[0xFF0E] = common.Int8ToByte(-122)

    expectedCycles := 2+4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0xFE95 {
        t.Error("PC should be 0xFE95 instead got: ", cpu.PC)
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
    cpu.X = 0

    cpu.Memory.Data[0x1000] = instructions.INS_INX_IMP
    cpu.Memory.Data[0x1001] = instructions.INS_BNE_REL
    cpu.Memory.Data[0x1002] = 0x33

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 2+3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1036 {
        t.Error("PC should be 0x1036 instead got: ", cpu.PC)
    }
}

func TestBNEDoesNotModifyPCIfZeroFlagIsSet(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1000)
    cpu.X = 1

    cpu.Memory.Data[0x1000] = instructions.INS_DEX_IMP
    cpu.Memory.Data[0x1001] = instructions.INS_BNE_REL
    cpu.Memory.Data[0x1002] = 0x33

    expectedCycles := 2+2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBNESumsCorrectlyZeroToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1000)
    cpu.X = 0

    cpu.Memory.Data[0x1000] = instructions.INS_INX_IMP
    cpu.Memory.Data[0x1001] = instructions.INS_BNE_REL
    cpu.Memory.Data[0x1002] = 0x00

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 2+3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBNESumsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFEFC)
    cpu.X = 0


    cpu.Memory.Data[0xFEFC] = instructions.INS_INX_IMP
    cpu.Memory.Data[0xFEFD] = instructions.INS_BNE_REL
    cpu.Memory.Data[0xFEFE] = 0x33

    // Page crossing happens so it takes 4 cycles
    expectedCycles := 2+4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // Since at this point PC is 0xFFFE, 0xFFFE+0x33 = 0x31 with wrap around
    if cpu.PC != 0xFF32 {
        t.Error("PC should be 0xFF32 instead got: ", cpu.PC)
    }
}

func TestBNESubtractsCorrectlyToProgramCounterWithoutPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.X = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_INX_IMP
    cpu.Memory.Data[0xFFFD] = instructions.INS_BNE_REL
    cpu.Memory.Data[0xFFFE] = common.Int8ToByte(-122)

    expectedCycles := 2+3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0xFF85 {
        t.Error("PC should be 0xFF85 instead got: ", cpu.PC)
    }
}

func TestBNESubtractsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFF0C)
    cpu.X = 0

    cpu.Memory.Data[0xFF0C] = instructions.INS_INX_IMP
    cpu.Memory.Data[0xFF0D] = instructions.INS_BNE_REL
    cpu.Memory.Data[0xFF0E] = common.Int8ToByte(-122)

    expectedCycles := 2+4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0xFE95 {
        t.Error("PC should be 0xFE95 instead got: ", cpu.PC)
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

func TestBCSSumsCorrectlyToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.C = 1

    cpu.Memory.Data[0x1001] = instructions.INS_BCS_REL
    cpu.Memory.Data[0x1002] = 0x33

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1036 {
        t.Error("PC should be 0x1036 instead got: ", cpu.PC)
    }
}

func TestBCSDoesNotModifyPCIfCarryFlagIsClear(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.C = 0

    cpu.Memory.Data[0x1001] = instructions.INS_BCS_REL
    cpu.Memory.Data[0x1002] = 0x33

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBCSSumsCorrectlyZeroToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.C = 1

    cpu.Memory.Data[0x1001] = instructions.INS_BCS_REL
    cpu.Memory.Data[0x1002] = 0x00

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBCSSumsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFEFD)
    cpu.PS.C = 1


    cpu.Memory.Data[0xFEFD] = instructions.INS_BCS_REL
    cpu.Memory.Data[0xFEFE] = 0x33

    // Page crossing happens so it takes 4 cycles
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // Since at this point PC is 0xFEFF, 0xFEFF+0x33 = 0xFF32 with wrap around
    if cpu.PC != 0xFF32 {
        t.Error("PC should be 0xFF32 instead got: ", cpu.PC)
    }
}

func TestBCSSubtractsCorrectlyToProgramCounterWithoutPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.C = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_BCS_REL
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

func TestBCSSubtractsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFF0C)
    cpu.PS.C = 1

    cpu.Memory.Data[0xFF0C] = instructions.INS_BCS_REL
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

func TestBCCSumsCorrectlyToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.C = 0

    cpu.Memory.Data[0x1001] = instructions.INS_BCC_REL
    cpu.Memory.Data[0x1002] = 0x33

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1036 {
        t.Error("PC should be 0x1036 instead got: ", cpu.PC)
    }
}

func TestBCCDoesNotModifyPCIfCarryFlagIsSet(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.C = 1

    cpu.Memory.Data[0x1001] = instructions.INS_BCC_REL
    cpu.Memory.Data[0x1002] = 0x33

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBCCSumsCorrectlyZeroToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.C = 0

    cpu.Memory.Data[0x1001] = instructions.INS_BCC_REL
    cpu.Memory.Data[0x1002] = 0x00

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBCCSumsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFEFD)
    cpu.PS.C = 0


    cpu.Memory.Data[0xFEFD] = instructions.INS_BCC_REL
    cpu.Memory.Data[0xFEFE] = 0x33

    // Page crossing happens so it takes 4 cycles
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // Since at this point PC is 0xFEFF, 0xFEFF+0x33 = 0xFF32 with wrap around
    if cpu.PC != 0xFF32 {
        t.Error("PC should be 0xFF32 instead got: ", cpu.PC)
    }
}

func TestBCCSubtractsCorrectlyToProgramCounterWithoutPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.C = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_BCC_REL
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

func TestBCCSubtractsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFF0C)
    cpu.PS.C = 0

    cpu.Memory.Data[0xFF0C] = instructions.INS_BCC_REL
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

func TestBMISumsCorrectlyToProgramCounter(t *testing.T){

    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.N = 1

    cpu.Memory.Data[0x1001] = instructions.INS_BMI_REL
    cpu.Memory.Data[0x1002] = 0x33

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1036 {
        t.Error("PC should be 0x1036 instead got: ", cpu.PC)
    }
}

func TestBMIDoesNotModifyPCIfNegativeFlagIsClear(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.N = 0

    cpu.Memory.Data[0x1001] = instructions.INS_BMI_REL
    cpu.Memory.Data[0x1002] = 0x33

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBMISumsCorrectlyZeroToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.N = 1

    cpu.Memory.Data[0x1001] = instructions.INS_BMI_REL
    cpu.Memory.Data[0x1002] = 0x00

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBMISumsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFEFD)
    cpu.PS.N = 1


    cpu.Memory.Data[0xFEFD] = instructions.INS_BMI_REL
    cpu.Memory.Data[0xFEFE] = 0x33

    // Page crossing happens so it takes 4 cycles
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // Since at this point PC is 0xFEFF, 0xFEFF+0x33 = 0xFF32 with wrap around
    if cpu.PC != 0xFF32 {
        t.Error("PC should be 0xFF32 instead got: ", cpu.PC)
    }
}

func TestBMISubtractsCorrectlyToProgramCounterWithoutPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.N = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_BMI_REL
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

func TestBMISubtractsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFF0C)
    cpu.PS.N = 1

    cpu.Memory.Data[0xFF0C] = instructions.INS_BMI_REL
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

func TestBPLSumsCorrectlyToProgramCounter(t *testing.T){

    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.N = 0

    cpu.Memory.Data[0x1001] = instructions.INS_BPL_REL
    cpu.Memory.Data[0x1002] = 0x33

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1036 {
        t.Error("PC should be 0x1036 instead got: ", cpu.PC)
    }
}

func TestBPLDoesNotModifyPCIfNegativeFlagIsSet(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.N = 1

    cpu.Memory.Data[0x1001] = instructions.INS_BPL_REL
    cpu.Memory.Data[0x1002] = 0x33

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBPLSumsCorrectlyZeroToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.N = 0

    cpu.Memory.Data[0x1001] = instructions.INS_BPL_REL
    cpu.Memory.Data[0x1002] = 0x00

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBPLSumsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFEFD)
    cpu.PS.N = 0


    cpu.Memory.Data[0xFEFD] = instructions.INS_BPL_REL
    cpu.Memory.Data[0xFEFE] = 0x33

    // Page crossing happens so it takes 4 cycles
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // Since at this point PC is 0xFEFF, 0xFEFF+0x33 = 0xFF32 with wrap around
    if cpu.PC != 0xFF32 {
        t.Error("PC should be 0xFF32 instead got: ", cpu.PC)
    }
}

func TestBPLSubtractsCorrectlyToProgramCounterWithoutPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.N = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_BPL_REL
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

func TestBPLSubtractsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFF0C)
    cpu.PS.N = 0

    cpu.Memory.Data[0xFF0C] = instructions.INS_BPL_REL
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

func TestBVSSumsCorrectlyToProgramCounter(t *testing.T){

    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.V = 1

    cpu.Memory.Data[0x1001] = instructions.INS_BVS_REL
    cpu.Memory.Data[0x1002] = 0x33

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1036 {
        t.Error("PC should be 0x1036 instead got: ", cpu.PC)
    }
}

func TestBVSDoesNotModifyPCIfNegativeFlagIsClear(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.V = 0

    cpu.Memory.Data[0x1001] = instructions.INS_BVS_REL
    cpu.Memory.Data[0x1002] = 0x33

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBVSSumsCorrectlyZeroToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.V = 1

    cpu.Memory.Data[0x1001] = instructions.INS_BVS_REL
    cpu.Memory.Data[0x1002] = 0x00

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBVSSumsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFEFD)
    cpu.PS.V = 1


    cpu.Memory.Data[0xFEFD] = instructions.INS_BVS_REL
    cpu.Memory.Data[0xFEFE] = 0x33

    // Page crossing happens so it takes 4 cycles
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // Since at this point PC is 0xFEFF, 0xFEFF+0x33 = 0xFF32 with wrap around
    if cpu.PC != 0xFF32 {
        t.Error("PC should be 0xFF32 instead got: ", cpu.PC)
    }
}

func TestBVSSubtractsCorrectlyToProgramCounterWithoutPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.V = 1

    cpu.Memory.Data[0xFFFC] = instructions.INS_BVS_REL
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

func TestBVSSubtractsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFF0C)
    cpu.PS.V = 1

    cpu.Memory.Data[0xFF0C] = instructions.INS_BVS_REL
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

func TestBVCSumsCorrectlyToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.V = 0

    cpu.Memory.Data[0x1001] = instructions.INS_BVC_REL
    cpu.Memory.Data[0x1002] = 0x33

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1036 {
        t.Error("PC should be 0x1036 instead got: ", cpu.PC)
    }
}

func TestBVCDoesNotModifyPCIfCarryFlagIsSet(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.V = 1

    cpu.Memory.Data[0x1001] = instructions.INS_BVC_REL
    cpu.Memory.Data[0x1002] = 0x33

    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBVCSumsCorrectlyZeroToProgramCounter(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0x1001)
    cpu.PS.V = 0

    cpu.Memory.Data[0x1001] = instructions.INS_BVC_REL
    cpu.Memory.Data[0x1002] = 0x00

    // No Page crossing happens so it takes 3 cycles
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PC != 0x1003 {
        t.Error("PC should be 0x1003 instead got: ", cpu.PC)
    }
}

func TestBVCSumsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFEFD)
    cpu.PS.V = 0


    cpu.Memory.Data[0xFEFD] = instructions.INS_BVC_REL
    cpu.Memory.Data[0xFEFE] = 0x33

    // Page crossing happens so it takes 4 cycles
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // Since at this point PC is 0xFEFF, 0xFEFF+0x33 = 0xFF32 with wrap around
    if cpu.PC != 0xFF32 {
        t.Error("PC should be 0xFF32 instead got: ", cpu.PC)
    }
}

func TestBVCSubtractsCorrectlyToProgramCounterWithoutPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.V = 0

    cpu.Memory.Data[0xFFFC] = instructions.INS_BVC_REL
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

func TestBVCSubtractsCorrectlyToProgramCounterWithPageCrossing(t *testing.T){
    
    cpu := Init6502()
    cpu.Reset(0xFF0C)
    cpu.PS.V = 0

    cpu.Memory.Data[0xFF0C] = instructions.INS_BVC_REL
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
