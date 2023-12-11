# 6502-emulator
A cpu emulator for the 6502 architecture in Go

About 6502 cpu.

Launched in 1975, it's a simple 8-bit processor and despite being simple
it does have a lot of principles that operate the same as the modern processors.
By understanding how it works you understand how the modern processors work

It has few internal registers, only three. A register is a small area of memory that the computer can access.
It can address only 64Kb of memory so that means it's got a 16-bit address bus, so there's 16 pins coming out the processor
that can address the memory. The processor is little endian and expects addresses to be stored in memory least significant byte first.

This means if we have a 16-bit value '0xABCD' the 6502 stores it as 'CD AB'

Memory Address   | 0  | 1  | 2  | 3 <br />
-----------------+----+----+----+---- <br />
Little-endian    | CD | AB | -- | --  <br />
Big-endian       | AB | CD | -- | -- <br />

Back in the days where memory was expensive and computers didn't have much for that reason, the cpu has a bunch of different
addressing modes on most of the instructions, that allows to address stuff that is in the first 256 bytes (called Zero Page, which 
goes from $0000 to $00FF), one cycle quicker that you can do anywhere else.
(In computer architectures, addressing modes define how the CPU calculates the effective address of an operand.
Different addressing modes allow flexibility in specifying operands in instructions.)
For example, instead of using a 16-bit address for an operand, you can use an 8-bit address if the operand is in the zero page.
This means this 256 bytes behave like a bunch of registers, you've got lots of instructions that you can do there.
This was what made the 6502 fast, you could do a lot of things faster skipping one clock cycle, and when you had not so many 
cycles per tick you had to save as many as you could.

The second page of memory, the next 256 bytes are stack memory, which you can't relocate. (from $0100 to $01FF)
