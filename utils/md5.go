package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(unEncry string) (encryption string){
	byte16 := []byte(unEncry)
	encryption = fmt.Sprintf( "%x",md5.Sum(byte16) )
	return
}
