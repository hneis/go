// Package signature provides ...
package signature

import (
	"io/ioutil"
	"os"
	"strings"
)

type Encoder struct {
	hashSign   string
	fileSource string
	signature  Signature
}

func NewEncoder(fileSource string, fileHashSign string) (enc *Encoder, err error) {
	hashString, err := ioutil.ReadFile(fileHashSign)
	if err != nil {
		return nil, err
	}
	enc = &Encoder{
		hashSign:   strings.TrimSuffix(string(hashString), "\n"),
		fileSource: fileSource,
		signature:  nil,
	}

	return
}

func (enc *Encoder) EncryptSHA256() (err error) {
	file, err := os.Open(enc.fileSource)
	if err != nil {
		return
	}
	defer file.Close()

	enc.signature = NewFromFileSource(file, enc.hashSign)

	return
}

func (enc *Encoder) SaveToFile(path string) error {
	// fmt.Println(enc.signature.SignatureByte())
	return ioutil.WriteFile(path, enc.signature.SignatureByte(), 0644)
}
