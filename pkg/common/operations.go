package common

func Int8ToByte(signed int8) (b byte){

    if signed < 0 {
        b = byte(signed & 0x7F) | 0x80
    }else{
        b = byte(signed)
    }
    
    return
}

func Int8AdditiveInverse(absoluteValue int8) byte{

    if absoluteValue != 0 {
        return (Int8ToByte(absoluteValue) ^ byte(0xFF)) + 1
    }

    return 0
}
    
