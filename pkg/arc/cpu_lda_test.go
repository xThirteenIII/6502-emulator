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
    // Make sure executing the instruction changes the flags
    cpu.PS.Z = 1
    cpu.PS.N = 1

    CheckLoadRegisterImmediate(cpu, instructions.INS_LDA_IM, &cpu.A,  t)
}

// Test if the LDA instruction loads 0 succefully into the A register
func TestLDAImmCanLoadZeroIntoARegister(t *testing.T){

    want := byte(0x0)

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    // Make sure executing the instruction changes the flags
    cpu.PS.Z = 0
    cpu.PS.N = 1

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

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    // Make sure executing the instruction changes the flags
    cpu.PS.Z = 1
    cpu.PS.N = 1

    CheckLoadRegisterZeroPage(cpu, instructions.INS_LDA_ZP, &cpu.A, t)
}

// Test if the LDA_ZPX instruction loads a value succefully into the A register
func TestLDAZeroXPageCanLoadIntoARegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    // Make sure executing the instruction changes the flags
    cpu.PS.Z = 1
    cpu.PS.N = 1

    CheckLoadRegisterZeroPageX(cpu, instructions.INS_LDA_ZPX, &cpu.A, t)
}

// Test if the LDA_ABS instruction loads a value succefully into the A register
func TestLDAAbsoluteCanLoadIntoARegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    // Make sure executing the instruction changes the flags
    cpu.PS.Z = 1
    cpu.PS.N = 1

    CheckLoadRegisterAbsolute(cpu, instructions.INS_LDA_ABS, &cpu.A, t)
}

// Test if the LDA_ZPX instruction loads a value succefully into the A register
func TestLDAAbsoluteXCanLoadIntoARegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    // Make sure executing the instruction changes the flags
    cpu.PS.Z = 1
    cpu.PS.N = 1

    CheckLoadRegisterAbsoluteX(cpu, instructions.INS_LDA_ABSX, &cpu.A, t)
}

func TestLDAAsboluteXTakesAnExtraCycleWithPageCrossing(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    // Make sure executing the instruction changes the flags
    cpu.PS.Z = 1
    cpu.PS.N = 1

    CheckLoadRegisterAbsoluteXWithPageCrossing(cpu, instructions.INS_LDA_ABSX, &cpu.A, t)

}

// Test if the LDA_ABSY instruction loads a value succefully into the A register
func TestLDAAbsoluteYCanLoadIntoARegister(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    // Make sure executing the instruction changes the flags
    cpu.PS.Z = 1
    cpu.PS.N = 1

    CheckLoadRegisterAbsoluteY(cpu, instructions.INS_LDA_ABSY, &cpu.A, t)
}

func TestLDAAsboluteYTakesAnExtraCycleWithPageCrossing(t *testing.T){

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    // Make sure executing the instruction changes the flags
    cpu.PS.Z = 1
    cpu.PS.N = 1

    CheckLoadRegisterAbsoluteYWithPageCrossing(cpu, instructions.INS_LDA_ABSY, &cpu.A, t)

}

// Example:
// LDX #$04
// LDA ($02, X)

// In the above case X is loaded with four, so the vector is calculated with 
// $02 + $04 = $06 (resulting vector)
// If the zero page memory $06 contains: 00 80, then the effective address from the vector (06)
// would be $8000
// Test if the LDA_INDX instruction loads a value succefully into the A register
func TestLDAIndirectXCanLoadIntoARegister(t *testing.T){

    want := byte(0x72)

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()
    // Make sure executing the instruction changes the flags
    cpu.PS.Z = 1
    cpu.PS.N = 1

    cpuCopy := *cpu

    // given
    cpu.X = 0x04

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_INDX
    cpu.Memory.Data[0xFFFD] = 0x02
    cpu.Memory.Data[0x0006] = 0x00
    cpu.Memory.Data[0x0007] = 0x80
    cpu.Memory.Data[0x8000] = want
    // end - inline program

    expectedCycles := 6
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }


    got := cpu.A
    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    if cpu.A != want {
        t.Error("Want: ", want, " instead got: ", got)
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

// Test if the LDA_INDY instruction loads a value succefully into the A register
func TestLDAIndirectYCanLoadIntoARegister(t *testing.T){

    want := byte(0x72)

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    cpuCopy := *cpu

    // given
    cpu.Y = 0x04

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_INDY
    cpu.Memory.Data[0xFFFD] = 0x02
    cpu.Memory.Data[0x0002] = 0x00
    cpu.Memory.Data[0x0003] = 0x80
    cpu.Memory.Data[0x8004] = want
    // end - inline program

    expectedCycles := 5
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }


    got := cpu.A
    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    if cpu.A != want {
        t.Error("Want: ", want, " instead got: ", got)
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

// Test if the LDA_INDY instruction loads a value succefully into the A register
func TestLDAIndirectYCanLoadIntoARegisterWithPageCrossing(t *testing.T){

    want := byte(0x72)

    cpu := &CPU{}
    cpu.Memory = Memory{}
    cpu.Reset()

    cpuCopy := *cpu

    // given
    cpu.Y = 0x04

    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_INDY
    cpu.Memory.Data[0xFFFD] = 0x02
    cpu.Memory.Data[0x0002] = 0xFF
    cpu.Memory.Data[0x0003] = 0x80
    cpu.Memory.Data[0x8103] = want
    // end - inline program

    expectedCycles := 6
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }


    got := cpu.A
    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    if cpu.A != want {
        t.Error("Want: ", want, " instead got: ", got)
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

// This is used to avoid duplicate code
// TODO: PROBLEM: when using different instructions, cycles needed might change. 
// Need a way to handle that and operations
func CheckLoadRegisterZeroPage(cpu *CPU, opcode int, register *byte, t *testing.T){

    // given
    cpu.Memory.Data[0xFFFC] = byte(opcode)
    cpu.Memory.Data[0xFFFD] = 0x72
    cpu.Memory.Data[0x0072] = 0x44

    // when
    cpuCopy := *cpu
    expectedCycles := 3
    cyclesUsed := cpu.Execute(expectedCycles)


    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }

    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    // then
    if *register != 0x44 {
        t.Error("Want: 0x44, got: ", *register)
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

// This is used to avoid duplicate code
// TODO: PROBLEM: when using different instructions, cycles needed might change. 
// Need a way to handle that and operations
func CheckLoadRegisterZeroPageX(cpu *CPU, opcode int, register *byte, t *testing.T){

    // given
    cpu.X = 5
    cpu.Memory.Data[0xFFFC] = byte(opcode)
    cpu.Memory.Data[0xFFFD] = 0x72
    // Add 5 to 0x72
    cpu.Memory.Data[0x0077] = 0x44

    // when
    cpuCopy := *cpu
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)


    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }

    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    // then
    if *register != 0x44 {
        t.Error("Want: 0x44, got: ", *register)
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

func CheckLoadRegisterZeroPageY(cpu *CPU, opcode int, register *byte, t *testing.T){

    // given
    cpu.Y = 5
    cpu.Memory.Data[0xFFFC] = byte(opcode)
    cpu.Memory.Data[0xFFFD] = 0x72
    // Add 5 to 0x72
    cpu.Memory.Data[0x0077] = 0x44

    // when
    cpuCopy := *cpu
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)


    if expectedCycles != cyclesUsed{
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }

    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    // then
    if *register != 0x44 {
        t.Error("Want: 0x44, got: ", *register)
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

func CheckLoadRegisterAbsolute(cpu *CPU, opcode int, register *byte, t *testing.T){

    want := byte(0x32)

    // Given
    // start - inline program
    cpu.Memory.Data[0xFFFC] = byte(opcode)
    // LSB
    cpu.Memory.Data[0xFFFD] = 0x80
    // MSB
    cpu.Memory.Data[0xFFFE] = 0x44
    cpu.Memory.Data[0x4480] = want
    // end - inline program

    // When
    cpuCopy := *cpu
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }

    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    // then
    if *register != want {
        t.Error("A: Want ", want, " instead got ", *register)
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

func CheckLoadRegisterAbsoluteX(cpu *CPU, opcode int, register *byte, t *testing.T){

    want := byte(0x32)

    // Given
    // start - inline program
    cpu.X = 1
    cpu.Memory.Data[0xFFFC] = byte(opcode)
    // LSB
    cpu.Memory.Data[0xFFFD] = 0x80
    // MSB
    cpu.Memory.Data[0xFFFE] = 0x44
    cpu.Memory.Data[0x4481] = want
    // end - inline program

    // When
    cpuCopy := *cpu
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }

    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    // then
    if *register != want {
        t.Error("A: Want ", want, " instead got ", *register)
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
func CheckLoadRegisterAbsoluteY(cpu *CPU, opcode int, register *byte, t *testing.T){

    want := byte(0x32)

    // Given
    // start - inline program
    cpu.Y = 1
    cpu.Memory.Data[0xFFFC] = byte(opcode)
    // LSB
    cpu.Memory.Data[0xFFFD] = 0x80
    // MSB
    cpu.Memory.Data[0xFFFE] = 0x44
    cpu.Memory.Data[0x4481] = want
    // end - inline program

    // When
    cpuCopy := *cpu
    expectedCycles := 4
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }

    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    // then
    if *register != want {
        t.Error("A: Want ", want, " instead got ", *register)
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

func CheckLoadRegisterAbsoluteXWithPageCrossing(cpu *CPU, opcode int, register *byte, t *testing.T){

    want := byte(0x32)

    // Given
    // start - inline program
    cpu.X = 0xFF
    cpu.Memory.Data[0xFFFC] = byte(opcode)
    // LSB
    cpu.Memory.Data[0xFFFD] = 0x80
    // MSB
    cpu.Memory.Data[0xFFFE] = 0x44
    cpu.Memory.Data[0x457F] = want
    // end - inline program

    // When
    cpuCopy := *cpu
    // ExpectedCycles is 4+1 due to page crossing
    expectedCycles := 5
    cyclesUsed := cpu.Execute(expectedCycles)

    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }

    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    // then
    if *register != want {
        t.Error("A: Want ", want, " instead got ", *register)
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


func CheckLoadRegisterAbsoluteYWithPageCrossing(cpu *CPU, opcode int, register *byte, t *testing.T){

    want := byte(0x32)

    // Given
    // start - inline program
    cpu.Y = 0xFF
    cpu.Memory.Data[0xFFFC] = byte(opcode)
    // LSB
    cpu.Memory.Data[0xFFFD] = 0x80
    // MSB
    cpu.Memory.Data[0xFFFE] = 0x44
    cpu.Memory.Data[0x457F] = want
    // end - inline program

    // When
    cpuCopy := *cpu
    // ExpectedCycles is 4+1 due to page crossing
    expectedCycles := 5
    cyclesUsed := cpu.Execute(expectedCycles)
    if cyclesUsed != expectedCycles {
        t.Error("Cycles used: ", expectedCycles, ", instead got: ", cyclesUsed)
    }

    gotZ := cpu.PS.Z
    gotN := cpu.PS.N

    // then
    if *register != want {
        t.Error("A: Want ", want, " instead got ", *register)
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
