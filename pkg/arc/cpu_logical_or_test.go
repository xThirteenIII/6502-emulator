package arc

import (
	"emulator/pkg/instructions"
	"testing"
)

// TODO: add zero and negative values checks
func TestORAImmediateCanPerformLogicalAnd(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0xFF
    cpuCopy := *cpu

    // These should be modified to 0 by execution
    cpu.PS.Z = 1
    cpu.PS.N = 1

    value := byte(0x04)

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ORA_IM
    cpu.Memory.Data[0xFFFD] = value
    
    expectedCycles := 2
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles!=cyclesUsed{
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // A should be FF ^ 04 = 11111111 | 00000100 = 11111111 = 255
    if cpu.A != (cpuCopy.A | value){
        t.Error(cpuCopy.A, " OR ", value, " should result", cpuCopy.A | value, ", instead got: ", cpu.A)
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestORAZeroPageCanPerformLogicalAnd(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0xFF
    cpuCopy := *cpu

    // These should be modified to 0 by execution
    cpu.PS.Z = 1
    cpu.PS.N = 1

    value := byte(0x04)

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ORA_ZP
    cpu.Memory.Data[0xFFFD] = 0xF3
    cpu.Memory.Data[0x00F3] = value
    
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles!=cyclesUsed{
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.A != (cpuCopy.A | value){
        t.Error(cpuCopy.A, " OR ", value, " should result", cpuCopy.A | value, ", instead got: ", cpu.A)
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestORAZeroPageXCanPerformLogicalAnd(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0xFF
    cpu.X = 0x02
    cpuCopy := *cpu

    // These should be modified to 0 by execution
    cpu.PS.Z = 1
    cpu.PS.N = 1

    value := byte(0x04)

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ORA_ZPX
    cpu.Memory.Data[0xFFFD] = 0xF3
    cpu.Memory.Data[0x00F5] = value
    
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles!=cyclesUsed{
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.A != (cpuCopy.A | value){
        t.Error(cpuCopy.A, " OR ", value, " should result", cpuCopy.A | value, ", instead got: ", cpu.A)
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestORAAbsoluteCanPerformLogicalAnd(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0xFF
    cpuCopy := *cpu

    // These should be modified to 0 by execution
    cpu.PS.Z = 1
    cpu.PS.N = 1

    value := byte(0x04)

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ORA_ABS
    cpu.Memory.Data[0xFFFD] = 0x00
    cpu.Memory.Data[0xFFFE] = 0x80
    cpu.Memory.Data[0x8000] = value
    
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles!=cyclesUsed{
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.A != (cpuCopy.A | value){
        t.Error(cpuCopy.A, " OR ", value, " should result", cpuCopy.A ^ value, ", instead got: ", cpu.A)
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestORAAbsoluteXCanPerformLogicalAnd(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0xFF
    cpu.X = 0x02
    cpuCopy := *cpu

    // These should be modified to 0 by execution
    cpu.PS.Z = 1
    cpu.PS.N = 1

    value := byte(0x04)

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ORA_ABSX
    cpu.Memory.Data[0xFFFD] = 0x00
    cpu.Memory.Data[0xFFFE] = 0x80
    cpu.Memory.Data[0x8002] = value
    
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles!=cyclesUsed{
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // A should be 04 & 02 = 00
    if cpu.A != (cpuCopy.A | value){
        t.Error(cpuCopy.A, " OR ", value, " should result", cpuCopy.A | value, ", instead got: ", cpu.A)
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestORAAbsoluteXCanPerformLogicalAndWithPageCrossing(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0xFF
    cpu.X = 0x02
    cpuCopy := *cpu

    // These should be modified to 0 by execution
    cpu.PS.Z = 1
    cpu.PS.N = 1

    value := byte(0x04)

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ORA_ABSX
    cpu.Memory.Data[0xFFFD] = 0xFF
    cpu.Memory.Data[0xFFFE] = 0x80
    cpu.Memory.Data[0x8101] = value
    
    expectedCycles := 5
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles!=cyclesUsed{
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // A should be 04 & 02 = 00
    if cpu.A != (cpuCopy.A | value){
        t.Error(cpuCopy.A, " XOR ", value, " should result", cpuCopy.A | value, ", instead got: ", cpu.A)
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestORAAbsoluteYCanPerformLogicalAnd(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0xFF
    cpu.Y = 0x02
    cpuCopy := *cpu

    // These should be modified to 0 by execution
    cpu.PS.Z = 1
    cpu.PS.N = 1

    value := byte(0x04)

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ORA_ABSY
    cpu.Memory.Data[0xFFFD] = 0x00
    cpu.Memory.Data[0xFFFE] = 0x80
    cpu.Memory.Data[0x8002] = value
    
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles!=cyclesUsed{
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.A != (cpuCopy.A | value){
        t.Error(cpuCopy.A, " OR ", value, " should result", cpuCopy.A | value, ", instead got: ", cpu.A)
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestORAAbsoluteYCanPerformLogicalAndWithPageCrossing(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0xFF
    cpu.Y = 0x02
    cpuCopy := *cpu

    // These should be modified to 0 by execution
    cpu.PS.Z = 1
    cpu.PS.N = 1

    value := byte(0x04)

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ORA_ABSY
    cpu.Memory.Data[0xFFFD] = 0xFF
    cpu.Memory.Data[0xFFFE] = 0x80
    cpu.Memory.Data[0x8101] = value
    
    expectedCycles := 5
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles!=cyclesUsed{
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // A should be 04 & 02 = 00
    if cpu.A != (cpuCopy.A | value){
        t.Error(cpuCopy.A, " OR ", value, " should result", cpuCopy.A | value, ", instead got: ", cpu.A)
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestORAIndirectXCanPerformLogicalAnd(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0xFF
    cpu.X = 0x02
    cpuCopy := *cpu

    // These should be modified to 0 by execution
    cpu.PS.Z = 1
    cpu.PS.N = 1

    value := byte(0x04)

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ORA_INDX
    cpu.Memory.Data[0xFFFD] = 0x32
    cpu.Memory.Data[0x0034] = 0x00
    cpu.Memory.Data[0x0035] = 0x80
    cpu.Memory.Data[0x8000] = value
    
    expectedCycles := 6
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles!=cyclesUsed{
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.A != (cpuCopy.A | value){
        t.Error(cpuCopy.A, " OR ", value, " should result", cpuCopy.A | value, ", instead got: ", cpu.A)
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestORAIndirectYCanPerformLogicalAnd(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0xFF
    cpu.Y = 0x02
    cpuCopy := *cpu

    // These should be modified to 0 by execution
    cpu.PS.Z = 1
    cpu.PS.N = 1

    value := byte(0x04)

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ORA_INDY
    cpu.Memory.Data[0xFFFD] = 0x32
    cpu.Memory.Data[0x0032] = 0x00
    cpu.Memory.Data[0x0033] = 0x80
    cpu.Memory.Data[0x8002] = value
    
    expectedCycles := 5
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles!=cyclesUsed{
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    if cpu.A != (cpuCopy.A | value){
        t.Error(cpuCopy.A, " OR ", value, " should result", cpuCopy.A | value, ", instead got: ", cpu.A)
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}

func TestORAIndirectYCanPerformLogicalAndWithPageCrossing(t *testing.T){

    // Given
    cpu := Init6502()
    cpu.A = 0xFF
    cpu.Y = 0x02
    cpuCopy := *cpu

    // These should be modified to 0 by execution
    cpu.PS.Z = 1
    cpu.PS.N = 1

    value := byte(0x04)

    // When
    cpu.Memory.Data[0xFFFC] = instructions.INS_ORA_INDY
    cpu.Memory.Data[0xFFFD] = 0x32
    cpu.Memory.Data[0x0032] = 0xFF
    cpu.Memory.Data[0x0033] = 0x80
    cpu.Memory.Data[0x8101] = value
    
    expectedCycles := 6
    cyclesUsed := cpu.Execute(expectedCycles)

    // Then
    if expectedCycles!=cyclesUsed{
        t.Error("Expected cycles: ", expectedCycles, "but got: ", cyclesUsed)
    }

    // A should be 04 & 02 = 00
    if cpu.A != (cpuCopy.A | value){
        t.Error(cpuCopy.A, " OR ", value, " should result", cpuCopy.A | value, ", instead got: ", cpu.A)
    }

    CheckUnmodifiedLDAFlags(cpuCopy, cpu, t)
}
