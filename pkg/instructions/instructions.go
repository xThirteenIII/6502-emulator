package instructions

// Opcodes
const (

    // Load/Store Operations
    INS_LDA_IM = 0xA9

    // An instruction using zero page addressing mode has only an 8 bit address operand.
    // This limits it to addressing only the first 256 bytes of memory (e.g. $0000 to $00FF) where the most significant byte of the address is always zero.
    // In zero page mode only the least significant byte of the address is held in the instruction making it shorter by one byte
    // (important for space saving) and one less memory fetch during execution (important for speed).
    INS_LDA_ZP = 0xA5
    INS_LDA_ZPX = 0xB5
    INS_LDA_ABS = 0xAD
    INS_LDA_ABSX = 0xBD
    INS_LDA_ABSY = 0xB9
    INS_LDA_INDX = 0xA1
    INS_LDA_INDY = 0xB1

    INS_LDX_IM = 0xA2
    INS_LDX_ZP = 0xA6
    INS_LDX_ZPY = 0xB6
    INS_LDX_ABS = 0xAE
    INS_LDX_ABSY = 0xBE

    INS_LDY_IM = 0xA0
    INS_LDY_ZP = 0xA4
    INS_LDY_ZPX = 0xB4
    INS_LDY_ABS = 0xAC
    INS_LDY_ABSX = 0xBC

    INS_STA_ZP = 0x85
    INS_STA_ZPX = 0x95
    INS_STA_ABS = 0x8D
    INS_STA_ABSX = 0x9D
    INS_STA_ABSY = 0x99
    INS_STA_INDX = 0x81
    INS_STA_INDY = 0x91

    // Register Transfers
    // Stack Operations
    // Logical
    // Arithmetic
    // Increments & Decrements
    // Shifts
    // Jump & Calls
    INS_JSR = 0x20

    // Branches
    // Status Flag Changes
    // System Functions

)
