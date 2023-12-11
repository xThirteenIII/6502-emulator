package instructions

// Opcodes
const (
    INS_LDA_IM = 0xA9

    // An instruction using zero page addressing mode has only an 8 bit address operand.
    // This limits it to addressing only the first 256 bytes of memory (e.g. $0000 to $00FF) where the most significant byte of the address is always zero.
    // In zero page mode only the least significant byte of the address is held in the instruction making it shorter by one byte
    // (important for space saving) and one less memory fetch during execution (important for speed).
    INS_LDA_ZP = 0xA5
    INS_LDA_ZPX = 0xB5
    INS_JSR = 0x20
)
