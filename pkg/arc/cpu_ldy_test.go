package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

// Test if the LDY instruction loads a value succefully into the Y register
func TestLDYImmCanLoadIntoARegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    CheckLoadRegisterImmediate(cpu, instructions.INS_LDY_IM, &cpu.Y,  t)
}

// Test if the LDY instruction loads 0 succefully into the Y register
func TestLDYImmCanLoadZeroIntoYRegister(t *testing.T){

    want := byte(0x0)

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    cpu.PS.Z = 0
    cpu.PS.N = 1

    // Make a copy of the cpu to confront uneffected flags
    // after execution
    cpuCopy := *cpu

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDY_IM
    cpu.Memory.Data[0xFFFD] = want
    // end - inline program

    // when
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }

    gotY := cpu.Y
    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    // then
    if cpu.Y != want {
        t.Error("A: Want ", want, " instead got ", gotY)
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


// Test if the LDY_ZP instruction loads a value succefully into the Y register
func TestLDYZeroPageCanLoadIntoYRegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    CheckLoadRegisterZeroPage(cpu, instructions.INS_LDY_ZP, &cpu.Y, t)
}

// Test if the LDY_ZPX instruction loads a value succefully into the Y register
func TestLDYZeroPageXCanLoadIntoYRegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    CheckLoadRegisterZeroPageX(cpu, instructions.INS_LDY_ZPX, &cpu.Y, t)
}

// Test if the LDY_ABS instruction loads a value succefully into the Y register
func TestLDYAbsoluteCanLoadIntoYRegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    CheckLoadRegisterAbsolute(cpu, instructions.INS_LDY_ABS, &cpu.Y, t)
}

// Test if the LDY_ABSX instruction loads a value succefully into the Y register
func TestLDYAbsoluteXCanLoadIntoYRegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    CheckLoadRegisterAbsoluteX(cpu, instructions.INS_LDY_ABSX, &cpu.Y, t)
}

// Test if the LDY_ABSX instruction loads a value succefully into the Y register
func TestLDYAbsoluteXTakesAnExtraCycleWithPageCrossing(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    CheckLoadRegisterAbsoluteXWithPageCrossing(cpu, instructions.INS_LDY_ABSX, &cpu.Y, t)
}
