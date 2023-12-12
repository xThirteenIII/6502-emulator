package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

// Tests if the LDA instruction loads a value succeffully into the A register
func TestLDAImmCanLoadIntoARegister(t *testing.T){

    want := byte(0x82)

    cpu := CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_IM
    cpu.Memory.Data[0xFFFD] = want
    // end - inline program

    // when
    cycles := 2
    cpu.Execute(&cycles)

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
}

// Tests if the LDA_ZP instruction loads a value succeffully into the A register
func TestLDAZeroPageCanLoadIntoARegister(t *testing.T){

    want := byte(0x82)

    cpu := CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_ZP
    cpu.Memory.Data[0xFFFD] = 0x42
    cpu.Memory.Data[0x0042] = want
    // end - inline program

    cycles := 3
    cpu.Execute(&cycles)

    got := cpu.A

    if cpu.A != want {
        t.Error("Want: ", want, " instead got: ", got)
    }
}

// Tests if the LDA_ZPX instruction loads a value succeffully into the A register
func TestLDAZeroXPageCanLoadIntoARegister(t *testing.T){

    want := byte(0x82)

    cpu := CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    // given
    cpu.X = 0x0F

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_ZPX
    cpu.Memory.Data[0xFFFD] = 0x80
    cpu.Memory.Data[0x008F] = want
    // end - inline program

    cycles := 4
    cpu.Execute(&cycles)

    got := cpu.A

    if cpu.A != want {
        t.Error("Want: ", want, " instead got: ", got)
    }
}
