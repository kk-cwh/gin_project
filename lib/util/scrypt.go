package util

import (
	"encoding/base64"
    "golang.org/x/crypto/scrypt"
	"log"
)

func ScryptStr(str string) (string, error){
	salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}
	dk, err := scrypt.Key([]byte(str), salt, 32768, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(dk),err
}
