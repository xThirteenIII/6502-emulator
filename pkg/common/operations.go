package common

func Int8ToByte(signed int8) (b byte){

    if signed < 0 {
        b = byte(signed & 0x7F) | 0x80
    }else{
        b = byte(signed)
    }
    
    return
}
