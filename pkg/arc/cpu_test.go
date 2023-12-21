package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

// Test if CPU register reset to default values correctly and memory initialises to 0.
func TestCPUResetsCorrectly(t *testing.T){

    cpu := Init6502()

    want := &CPU{
        PC: 0xFFFC,
        SP: 0x00,
        A: 0,
        X: 0,
        Y: 0,
        PS: ProcessorStatus{
            C: 0,
            Z: 0,
            I: 0,
            D: 0,
            B: 0,
            U: 0,
            V: 0,
            N: 0,
        },
    }
    for i := 0; i < MaxMem; i++{
        cpu.Memory.Data[i] = 0
    }

    if *cpu != *want {
        
        t.Error("Want: ", want, ", got: ", cpu)
    }
}

func TestCPUDoesNothingWhenWeExecuteZeroCycles(t *testing.T){

    // given
    const NUM_CYCLES = 0
    cpu := Init6502()

    // when 
    cyclesUsed := cpu.Execute(0)

    // then
    if cyclesUsed != 0 {
        t.Error("Executing with zero cycles should return 0")
    }
    

}

func TestCPUCanExecuteMoreCyclesThanRequestedIfRequiredByInstruction(t *testing.T){

    // given
    cpu := Init6502()

    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_IM
    cpu.Memory.Data[0xFFFD] = 0x84

    // when 
    // LDA_IM should execute 2 cycles anyways
    const NUM_CYCLES = 1
    cyclesUsed := cpu.Execute(NUM_CYCLES)

    // then
    // When executing Execute(1) returns 1 - (-1) = 2 and that should pass the test
    // Even the parameter passed is 1, it executes required cycles anyways
    if cyclesUsed != 2 {
        t.Error("Couldn't run at least 1 cycle")
    }
}

func Init6502() (cpu *CPU){
    cpu = &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    return
}
