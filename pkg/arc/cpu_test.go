package arc

import (
	"testing"
)

// Test if CPU register reset to default values correctly and memory initialises to 0.
func TestCPUResetsCorrectly(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    cpu.Memory.Initialise()

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

