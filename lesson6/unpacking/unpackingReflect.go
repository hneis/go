// Package unpacking provides ...
package unpacking

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"reflect"
)

func Pack(sig interface{}) (buf *bytes.Buffer, err error) {
	orderByte := binary.BigEndian
	v := reflect.ValueOf(sig).Elem()
	if v.Kind() != reflect.Struct {
		err = errors.New("Is not struct type")
		return
	}

	buffer := &bytes.Buffer{}
	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Type().Kind() {
		case reflect.Uint32, reflect.Uint64:
			err1 := binary.Write(buffer, orderByte, v.Field(i).Interface())
			if err != nil {
				err = err1
				return
			}
		case reflect.String:
			s := v.Field(i).String()
			b := []byte(s)
			err1 := binary.Write(buffer, orderByte, uint16(len(b)))

			if err1 != nil {
				err = err1
				return
			}
			_, err1 = buffer.Write(b)
			if err1 != nil {
				err = err1
				return
			}

		case reflect.Slice:
			buffer.Write(v.Field(i).Bytes())
		default:
			fmt.Printf("Bad format field %v\n", v.Field(i).Type().Kind().String())

		}
	}
	return nil
}

func Unpack(u interface{}, data[]byte) (err error) {
	orderByte := binary.BigEndian
	r := bytes.NewReader(data)
	val := reflect.ValueOf(u).Elem()
	for i:= 0; i< val.NumField(); i++ {
		switch val.Field(i).Type().Kind() {
		case reflect.Uint32, reflect.Uint64:
			var value uint64
			binary.Read(r, orderByte, &value)
			val.Field(i).Set(refreflect.ValueOf(uint64(value)))

		case reflect.String:
			var lenRaw uint16
			binary.Read(r,orderByte, &lenRaw)
			dataRaw:=make([]byte, lenRaw)
			binary.Read(r, orderByte, &dataRaw)
			val.Field(i).SetString(string(dataRaw))
		case reflect.Slice:
			var data []int
			binary.Read(r, orderByte, &data)
		}
		strings.Join()
}
