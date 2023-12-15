package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

// Test if the LDA instruction loads a value succefully into the A register
func TestLDAImmCanLoadIntoARegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    CheckLoadRegisterImmediate(cpu, instructions.INS_LDA_IM, &cpu.A,  t)
}

// Test if the LDA instruction loads 0 succefully into the A register
func TestLDAImmCanLoadZeroIntoARegister(t *testing.T){

    want := byte(0x0)

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    // Make a copy of the cpu to confront uneffected flags
    // after execution
    cpuCopy := *cpu

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_IM
    cpu.Memory.Data[0xFFFD] = want
    // end - inline program

    // when
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }

    gotA := cpu.A
    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    // then
    if cpu.A != want {
        t.Error("A: Want ", want, " instead got ", gotA)
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


// Test if the LDA_ZP instruction loads a value succefully into the A register
func TestLDAZeroPageCanLoadIntoARegister(t *testing.T){

    want := byte(0x72)

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    cpuCopy := *cpu

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_ZP
    cpu.Memory.Data[0xFFFD] = 0x42
    cpu.Memory.Data[0x0042] = want
    // end - inline program

    // when
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }

    gotA := cpu.A
    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    // then
    if cpu.A != want {
        t.Error("A: Want ", want, " instead got ", gotA)
    }

    if cpu.PS.Z != 0 {
        t.Error("Z: Want 0, instead got: ", gotZ)
    }

    if cpu.PS.N != 0 {
        t.Error("N: Want 0, instead got: ", gotN)
    }

    // Confront uneffected flags
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

// Test if the LDA_ZPX instruction loads a value succefully into the A register
func TestLDAZeroXPageCanLoadIntoARegister(t *testing.T){

    want := byte(0x72)

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    cpuCopy := *cpu

    // given
    cpu.X = 0x0F

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_ZPX
    cpu.Memory.Data[0xFFFD] = 0x80
    cpu.Memory.Data[0x008F] = want
    // end - inline program

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }


    gotA := cpu.A
    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    // then
    if cpu.A != want {
        t.Error("A: Want ", want, " instead got ", gotA)
    }

    if cpu.PS.Z != 0 {
        t.Error("Z: Want 0, instead got: ", gotZ)
    }

    if cpu.PS.N != 0 {
        t.Error("N: Want 0, instead got: ", gotN)
    }

    // Confront uneffected flags
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

// Test if the LDA_ABS instruction loads a value succefully into the A register
func TestLDAAbsoluteCanLoadIntoARegister(t *testing.T){

    want := byte(0x72)

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    cpuCopy := *cpu

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_ABS
    // LSB
    cpu.Memory.Data[0xFFFD] = 0x80
    // MSB
    cpu.Memory.Data[0xFFFE] = 0x44
    cpu.Memory.Data[0x4480] = want
    // end - inline program

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }


    gotA := cpu.A
    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    // then
    if cpu.A != want {
        t.Error("A: Want ", want, " instead got ", gotA)
    }

    if cpu.PS.Z != 0 {
        t.Error("Z: Want 0, instead got: ", gotZ)
    }

    if cpu.PS.N != 0 {
        t.Error("N: Want 0, instead got: ", gotN)
    }

    // Confront uneffected flags
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}
// Test if the LDA_ZPX instruction loads a value succefully into the A register
func TestLDAAbsoluteXCanLoadIntoARegister(t *testing.T){

    want := byte(0x82)

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    cpuCopy := *cpu

    // given
    cpu.X = 0x0F

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_ZPX
    cpu.Memory.Data[0xFFFD] = 0x80
    cpu.Memory.Data[0x008F] = want
    // end - inline program

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }


    got := cpu.A

    if cpu.A != want {
        t.Error("Want: ", want, " instead got: ", got)
    }

    // Confront uneffected flags
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}
// Test if the LDA_ABSY instruction loads a value succefully into the A register
func TestLDAAbsoluteYCanLoadIntoARegister(t *testing.T){

    want := byte(0x82)

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    cpuCopy := *cpu

    // given
    cpu.X = 0x0F

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_ZPX
    cpu.Memory.Data[0xFFFD] = 0x80
    cpu.Memory.Data[0x008F] = want
    // end - inline program

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }


    got := cpu.A

    if cpu.A != want {
        t.Error("Want: ", want, " instead got: ", got)
    }

    // Confront uneffected flags
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}
// Test if the LDA_ZPX instruction loads a value succefully into the A register
func TestLDAIndirectXCanLoadIntoARegister(t *testing.T){

    want := byte(0x82)

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    cpuCopy := *cpu

    // given
    cpu.X = 0x0F

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_ZPX
    cpu.Memory.Data[0xFFFD] = 0x80
    cpu.Memory.Data[0x008F] = want
    // end - inline program

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }


    got := cpu.A

    if cpu.A != want {
        t.Error("Want: ", want, " instead got: ", got)
    }

    // Confront uneffected flags
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

// Test if the LDA_ZPX instruction loads a value succefully into the A register
func TestLDAIndirectYCanLoadIntoARegister(t *testing.T){

    want := byte(0x82)

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    cpuCopy := *cpu

    // given
    cpu.X = 0x0F

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_ZPX
    cpu.Memory.Data[0xFFFD] = 0x80
    cpu.Memory.Data[0x008F] = want
    // end - inline program

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }


    got := cpu.A

    if cpu.A != want {
        t.Error("Want: ", want, " instead got: ", got)
    }

    // Confront uneffected flags
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)

}

// This is used to avoid duplicate code
// TODO: PROBLEM: when using different instructions, cycles needed might change. 
// Need a way to handle that and operations
func CheckLoadRegisterImmediate(cpu *CPU, opcode int, register *byte, t *testing.T){

    // given
    cpu.Memory.Data[0xFFFC] = byte(opcode)
    cpu.Memory.Data[0xFFFD] = 0x72

    // when
    cpuCopy := *cpu
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)


    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }

    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    // then
    if *register != 0x72 {
        t.Error("Want: 0x72, got: ", *register)
    }

    if cpu.PS.Z != 0 {
        t.Error("Z: Want 0, instead got: ", gotZ)
    }

    if cpu.PS.N != 0 {
        t.Error("N: Want 0, instead got: ", gotN)
    }

    // Confront uneffected flags
    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)

}

// Confront Initial PS Registers values with values after execution.
// These register shuould remain unmodified
func CheckUnmodifiedLDAFlags(cpuCopy CPU, cpu *CPU, t *testing.T){

    // Confront uneffected flags
    if cpu.PS.C != cpuCopy.PS.C {
        t.Error("PS.C: want: ", cpuCopy.PS.C, ", got: ", cpu.PS.C)
    }

    if cpu.PS.I != cpuCopy.PS.I {
        t.Error("PS.C: want: ", cpuCopy.PS.C, ", got: ", cpu.PS.C)
    }

    if cpu.PS.U != cpuCopy.PS.U {
        t.Error("PS.C: want: ", cpuCopy.PS.C, ", got: ", cpu.PS.C)
    }

    if cpu.PS.B != cpuCopy.PS.B {
        t.Error("PS.C: want: ", cpuCopy.PS.C, ", got: ", cpu.PS.C)
    }

    if cpu.PS.D != cpuCopy.PS.D {
        t.Error("PS.C: want: ", cpuCopy.PS.C, ", got: ", cpu.PS.C)
    }
    if cpu.PS.V != cpuCopy.PS.V {
        t.Error("PS.C: want: ", cpuCopy.PS.C, ", got: ", cpu.PS.C)
    }
}
