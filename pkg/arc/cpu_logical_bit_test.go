package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

func TestBITZeroPageCorrectlySetsFlags(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.Z = 1
    cpu.PS.N = 1
    cpu.PS.V = 0
    cpu.A = 0x12

    cpu.Memory.Data[0xFFFC] = instructions.INS_BIT_ZP
    cpu.Memory.Data[0xFFFD] = 0x72
    cpu.Memory.Data[0x0072] = byte(114) // 01110010

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PS.Z != 0 {
        t.Error("Z: want 0,  but got: ", cpu.PS.Z)
    }

    if cpu.PS.N != 0 {
        t.Error("N: want 0,  but got: ", cpu.PS.N)
    }

    if cpu.PS.V != 1 {
        t.Error("V: want 0,  but got: ", cpu.PS.V)
    }
    

}

func TestBITAbsoluteCorrectlySetsFlags(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.Z = 1
    cpu.PS.N = 1
    cpu.PS.V = 0
    cpu.A = 0x12

    cpu.Memory.Data[0xFFFC] = instructions.INS_BIT_ABS
    cpu.Memory.Data[0xFFFD] = 0x72
    cpu.Memory.Data[0xFFFE] = 0x44
    cpu.Memory.Data[0x4472] = byte(114) // 01110010

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PS.Z != 0 {
        t.Error("Z: want 0,  but got: ", cpu.PS.Z)
    }

    if cpu.PS.N != 0 {
        t.Error("N: want 0,  but got: ", cpu.PS.N)
    }

    if cpu.PS.V != 1 {
        t.Error("V: want 0,  but got: ", cpu.PS.V)
    }
    

}

func TestBITZeroPageCorrectlySetsZeroFlag(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.Z = 0
    cpu.PS.N = 1
    cpu.PS.V = 0
    cpu.A = 0x00

    cpu.Memory.Data[0xFFFC] = instructions.INS_BIT_ZP
    cpu.Memory.Data[0xFFFD] = 0x72
    cpu.Memory.Data[0x0072] = byte(114) // 01110010

    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PS.Z != 1 {
        t.Error("Z: want 1,  but got: ", cpu.PS.Z)
    }

    if cpu.PS.N != 0 {
        t.Error("N: want 0,  but got: ", cpu.PS.N)
    }

    if cpu.PS.V != 1 {
        t.Error("V: want 0,  but got: ", cpu.PS.V)
    }
    

}

func TestBITAbsoluteCorrectlySetsZeroFlag(t *testing.T){
    
    cpu := Init6502()
    cpu.PS.Z = 0
    cpu.PS.N = 1
    cpu.PS.V = 0
    cpu.A = 0x00

    cpu.Memory.Data[0xFFFC] = instructions.INS_BIT_ABS
    cpu.Memory.Data[0xFFFD] = 0x72
    cpu.Memory.Data[0xFFFE] = 0x44
    cpu.Memory.Data[0x4472] = byte(114) // 01110010

    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if expectedCycles != cyclesUsed {
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.PS.Z != 1 {
        t.Error("Z: want 1,  but got: ", cpu.PS.Z)
    }

    if cpu.PS.N != 0 {
        t.Error("N: want 0,  but got: ", cpu.PS.N)
    }

    if cpu.PS.V != 1 {
        t.Error("V: want 0,  but got: ", cpu.PS.V)
    }
    

}

