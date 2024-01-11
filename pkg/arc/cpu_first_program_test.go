package arc

import (
	"testing"
)

// Emulating program that we know that works:
// pc is @ $1000
// lda #$FF

// start
// sta $90
// sta $8000
// eor #$CC
// jmp start

// bytes: { 0x00, 0x10, 0xA9, 0xFF, 0x85, 0x90, 0x8D, 0x00, 0x80, 0x49, 0xCC, 0x4C, 0x02, 0x10}

var testProgram = []byte{0x00, 0x10, 0xA9, 0xFF, 0x85, 0x90, 0x8D, 0x00, 0x80, 0x49, 0xCC, 0x4C, 0x02, 0x10}


func TestProgramIsLoadedIntoMemory(t *testing.T){
    
    cpu := Init6502()

    // when
    cpu.LoadProgram(testProgram)

    // then
    if cpu.Memory.Data[0x1000] != 0xA9{
        t.Error("expected 169 at $1000 but instead got ", cpu.Memory.Data[0x1000])
    }
    if cpu.Memory.Data[0x1001] != 0xFF{
        t.Error("expected 255 at $1001 but instead got ", cpu.Memory.Data[0x1001])
    }
    if cpu.Memory.Data[0x1002] != 0x85{
        t.Error("expected 132 at $1002 but instead got ", cpu.Memory.Data[0x1002])
    }
    if cpu.Memory.Data[0x1003] != 0x90{
        t.Error("expected 144 at $1003 but instead got ", cpu.Memory.Data[0x1003])
    }
    // ...
    if cpu.Memory.Data[0x100A] != 0x02{
        t.Error("expected 2 at $100A but instead got ", cpu.Memory.Data[0x100A])
    }
    if cpu.Memory.Data[0x100B] != 0x10{
        t.Error("expected 16 at $100B but instead got ", cpu.Memory.Data[0x100B])
    }
}

func TestProgramIsLoadedIntoMemoryAndExecuted(t *testing.T){

    clock := 1000
    
    cpu := Init6502()

    // when
    cpu.LoadProgram(testProgram)

    for clock > 0 {

        clock -= cpu.Execute(1)
    }
}
