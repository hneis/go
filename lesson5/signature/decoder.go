// Package decoder provides ...
package signature

import (
	"errors"
	"io/ioutil"
	"os"
)

type EncryptCheck struct {
	fileSource string
	hashString string
	fileSigned string
}

func NewEncryptCheck(fileSource, fileHash, fileSigned string) (dec *EncryptCheck, err error) {
	hashString, err := ioutil.ReadFile(fileSigned)
	if err != nil {
		return
	}

	dec = &EncryptCheck{
		hashString: string(hashString),
		fileSource: fileSource,
		fileSigned: fileSigned,
	}

	return
}

func (dec *EncryptCheck) Check() (err error) {
	file, err := os.Open(dec.fileSource)
	if err != nil {
		return
	}

	defer file.Close()
	signatureSource := NewFromFileSource(file, dec.hashString)
	_ = signatureSource

	fileSign, err := os.Open(dec.fileSigned)
	if err != nil {
		return
	}
	defer fileSign.Close()

	signatureSignedFile := NewFromSignedFile(fileSign)

	if !signatureSource.Equals(signatureSignedFile) {
		err = errors.New("Signature wrong")
	}

	return
}
