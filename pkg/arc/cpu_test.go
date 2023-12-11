package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

func TestExecute(t *testing.T){

    cpu := CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_JSR
    cpu.Memory.Data[0xFFFD] = 0x42
    cpu.Memory.Data[0xFFFE] = 0x42
    cpu.Memory.Data[0x4242] = instructions.INS_LDA_IM
    cpu.Memory.Data[0x4243] = 0x82
    // end - inline program

    cycles := 8
    cpu.Execute(&cycles)

    want := byte(0x82)

    got := cpu.A

    if cpu.A != want {
        t.Error("Want: ", want, " instead got: ", got)
    }

    

}
