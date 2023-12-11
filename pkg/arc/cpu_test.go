package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

// Tests if the LDA instruction loads a value succeffully into the A register
func TestLDACanLoadIntoARegister(t *testing.T){

    want := byte(0x82)

    cpu := CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_IM
    cpu.Memory.Data[0xFFFD] = want
    // end - inline program

    cycles := 2
    cpu.Execute(&cycles)

    got := cpu.A

    if cpu.A != want {
        t.Error("Want: ", want, " instead got: ", got)
    }

    



}
