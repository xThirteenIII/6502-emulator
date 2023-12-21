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
        SP: 0xFD,
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
        
        t.Error("CPU not reset correctly")
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
    cpu.Reset(0xFFFC)

    return
}

func TestSPIsReturnedAs16BitAddressCorrectly(t *testing.T){
    cpu := Init6502()
    SP := cpu.SPTo16Address(cpu.SP)

    if SP != 0x01FD {
        t.Error("SP should be 0x01FD but got: ", SP)
    }
}

func TestPopByteFromStack(t *testing.T){

    cpu := Init6502()

    cpu.Memory.Data[0x01FD] = 0x32
    cpu.SP--

    expectedCycles := 1
    data := cpu.PopByteFromStack(&expectedCycles)

    if cpu.SP != 0xFD {
        t.Error("SP not incremented correctly, got: ", cpu.SP, "but want: 0xFD")
    }
    if data != 0x32 {
        t.Error ("Expected 0x32 but got: ", data)
    }


}

func TestPopWordFromStack(t *testing.T){

    cpu := Init6502()

    // MSB first since it's higher memory address
    cpu.Memory.Data[0x01FD] = 0x44
    cpu.Memory.Data[0x01FC] = 0x32
    cpu.SP -=2

    expectedCycles := 2
    data := cpu.PopWordFromStack(&expectedCycles)

    if cpu.SP != 0xFD {
        t.Error("SP not incremented correctly, got: ", cpu.SP, "but want: 0xFD")
    }

    if data != 0x4432 {
        t.Error ("Expected 0x4432 but got: ", data)
    }

}

func TestPushByteToStack(t *testing.T){

    cpu := Init6502()

    // MSB first since it's higher memory address
    expectedCycles := 1
    cpu.PushByteToStack(&expectedCycles, 0x3F)

    if cpu.SP != 0xFC {
        t.Error("SP not decremented correctly, got: ", cpu.SP, "but want: 0xFC")
    }

    if cpu.Memory.Data[0x01FD] != 0x3F{
        t.Error ("Expected 0x3F but got: ", cpu.Memory.Data[0x01FD])
    }
}

func TestPushWordToStack(t *testing.T){

    cpu := Init6502()

    // MSB first since it's higher memory address
    expectedCycles := 1
    cpu.PushWordToStack(&expectedCycles, 0x333F)

    if cpu.SP != 0xFB {
        t.Error("SP not decremented correctly, got: ", cpu.SP, "but want: 0xFB")
    }

    if cpu.Memory.Data[0x01FD] != 0x33 || cpu.Memory.Data[0x01FC] != 0x3F {
        t.Error("Pushed: ", uint16(cpu.Memory.Data[0x01FC]) | (uint16(cpu.Memory.Data[0x01FD]) << 8), "but want: 333F")
    }
}
