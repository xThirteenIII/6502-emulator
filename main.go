package main

import (
	"emulator/pkg/arc"
	"emulator/pkg/instructions"
	"fmt"
)

// References

// http://www.6502.org/users/obelisk/
// https://sta.c64.org/cbm64mem.html
// https://www.c64-wiki.com/wiki/Reset_(Process)




func main() {



    cpu := arc.CPU{}
    cpu.Memory = arc.Memory{}

    cpu.Reset()

    fmt.Println("initial cpu values:")
    cpu.PrintValues()
    // start - inline program
    cpu.Memory.Data[0xFFFC] = instructions.INS_LDA_ZP
    cpu.Memory.Data[0xFFFD] = 0x42
    cpu.Memory.Data[0x0042] = 0x84
    // end - inline program

    // try to execute that instruction
    // fmt.Println(mem.Data[0xFFFD])
    cycles := 3
    cpu.Execute(&cycles)

    fmt.Println("A:",cpu.A)

    fmt.Println("final cpu values:")
    cpu.PrintValues()


}
