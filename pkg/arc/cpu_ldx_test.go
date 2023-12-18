package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

// Test if the LDX instruction loads a value succefully into the X register
func TestLDXImmCanLoadIntoARegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    CheckLoadRegisterImmediate(cpu, instructions.INS_LDX_IM, &cpu.X,  t)
}

// Test if the LDX instruction loads 0 succefully into the X register
func TestLDXImmCanLoadZeroIntoXRegister(t *testing.T){

    want := byte(0x0)

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    // Make sure flags registers are changed when executing to correct values
    cpu.PS.Z = 0
    cpu.PS.N = 1

    // Make a copy of the cpu to confront uneffected flags
    // after execution
    cpuCopy := *cpu

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDX_IM
    cpu.Memory.Data[0xFFFD] = want
    // end - inline program

    // when
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }

    gotX := cpu.X
    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    // then
    if cpu.X != want {
        t.Error("A: Want ", want, " instead got ", gotX)
    }

    if cpu.PS.Z != 1 {
        t.Error("Z: Want 1, instead got: ", gotZ)
    }

    if cpu.PS.N != 0 {
        t.Error("N: Want 0, instead got: ", gotN)
    }

    // Confront uneffected flags
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}


// Test if the LDX_ZP instruction loads a value succefully into the X register
func TestLDXZeroPageCanLoadIntoXRegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    CheckLoadRegisterZeroPage(cpu, instructions.INS_LDX_ZP, &cpu.X, t)
}

// Test if the LDX_ZPY instruction loads a value succefully into the X register
func TestLDXZeroPageYCanLoadIntoXRegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    CheckLoadRegisterZeroPageY(cpu, instructions.INS_LDX_ZPY, &cpu.X, t)
}

// Test if the LDX_ABS instruction loads a value succefully into the X register
func TestLDXAbsoluteCanLoadIntoXRegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    CheckLoadRegisterAbsolute(cpu, instructions.INS_LDX_ABS, &cpu.X, t)
}

// Test if the LDX_ABSY instruction loads a value succefully into the X register
func TestLDXAbsoluteYCanLoadIntoXRegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    CheckLoadRegisterAbsoluteY(cpu, instructions.INS_LDX_ABSY, &cpu.X, t)
}

// Test if the LDX_ABSY instruction loads a value succefully into the X register
func TestLDXAbsoluteYTakesAnExtraCycleWithPageCrossing(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    CheckLoadRegisterAbsoluteYWithPageCrossing(cpu, instructions.INS_LDX_ABSY, &cpu.X, t)
}
