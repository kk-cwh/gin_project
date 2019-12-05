package util

import (
	"encoding/base64"
    "golang.org/x/crypto/scrypt"
	"log"
	"reflect"
)

//字符串加密 这里用于password加密
func ScryptStr(str string) (string, error){
	salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}
	dk, err := scrypt.Key([]byte(str), salt, 32768, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(dk),err
}

//两个相似结构体的复制 src =》dst
func CopyStruct(src, dst interface{}) {
	sVal := reflect.ValueOf(src).Elem()
	dVal := reflect.ValueOf(dst).Elem()

	for i := 0; i < sVal.NumField(); i++ {
		value := sVal.Field(i)
		name := sVal.Type().Field(i).Name

		dvalue := dVal.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		dvalue.Set(value) //这里默认共同成员的类型一样，否则这个地方可能导致 panic，需要简单修改一下。
	}
}