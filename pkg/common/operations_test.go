package common

import "testing"

func TestInt8ToByteWorksCorrectlyWithNegativeNumber(t *testing.T){

    intToConvert := int8(-122)

    convertedByte := Int8ToByte(intToConvert)

    // -122 is: 10000110 as int8

    // 134 is: 10000110 as byte
    if convertedByte != 134{
        t.Error("Converted byte from int8: ", intToConvert, " should be 134 instead got: ", convertedByte)
    }

}

func TestInt8ToByteWorksCorrectlyWithPositiveNumber(t *testing.T){

    intToConvert := int8(122)

    convertedByte := Int8ToByte(intToConvert)

    if convertedByte != 122{
        t.Error("Converted byte from int8: ", intToConvert, " should be 122 instead got: ", convertedByte)
    }

}
func TestInt8ToByteWorksCorrectlyWithZero(t *testing.T){

    intToConvert := int8(0)

    convertedByte := Int8ToByte(intToConvert)

    if convertedByte != 0{
        t.Error("Converted byte from int8: ", intToConvert, " should be 0 instead got: ", convertedByte)
    }

}

func TestInt8AdditiveInverseWorksCorrectlyWithZero(t *testing.T){

    intToConvert := int8(0)

    convertedByte := Int8AdditiveInverse(intToConvert)

    if convertedByte != 0{
        t.Error("Inverse byte of: ", intToConvert, " should be 0 instead got: ", convertedByte)
    }
}

func TestInt8AdditiveInverseWorksCorrectlyWithPositiveNumber(t *testing.T){

    intToConvert := int8(127)

    convertedByte := Int8AdditiveInverse(intToConvert)

    if convertedByte != 0x81{
        t.Error("Inverse byte of: ", intToConvert, " should be 0x81(-127) instead got: ", convertedByte)
    }
}

func TestInt8AdditiveInverseWorksCorrectlyWithNegativeNumber(t *testing.T){

    intToConvert := int8(-127)

    convertedByte := Int8AdditiveInverse(intToConvert)

    if convertedByte != 0x7F{
        t.Error("Inverse byte of: ", intToConvert, " should be 0x7F(127) instead got: ", convertedByte)
    }
}
