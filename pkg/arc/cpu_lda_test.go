package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

// Test if the LDA instruction loads a value succefully into the A register
func TestLDAImmCanLoadIntoARegister(t *testing.T){

    want := byte(0x82)

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

    got := cpu.A
    got1 := cpu.PS.Z
    got2 := cpu.PS.N

    // then
    if cpu.A != want {
        t.Error("A: Want ", want, " instead got ", got)
    }

    if cpu.PS.Z != 0 {
        t.Error("A: Want 0, instead got: ", got1)
    }

    if cpu.PS.N != 1 {
        t.Error("N: Want 1, instead got: ", got2)
    }

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

// Test if the LDA_ZP instruction loads a value succefully into the A register
func TestLDAZeroPageCanLoadIntoARegister(t *testing.T){

    want := byte(0x82)

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

    got := cpu.A

    if cpu.A != want {
        t.Error("Want: ", want, " instead got: ", got)
    }

    // TODO: find a way to handle duplicate code and error return
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

// Test if the LDA_ZPX instruction loads a value succefully into the A register
func TestLDAZeroXPageCanLoadIntoARegister(t *testing.T){

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

func CheckLDAUneffectedFlags(cpuCopy CPU, cpu *CPU) (err error){

    return
}
