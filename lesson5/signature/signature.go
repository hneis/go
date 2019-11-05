// Package signature provides ...
package signature

import (
	"bytes"
	"crypto/sha256"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

const DATE_FORMAT = "2006-01-02 15-04-05"

type Signature interface {
	Date() time.Time
	Size() string
	// source file
	Name() string
	SignatureByte() []byte
	Equals(s Signature) bool
}

type SignatureSHA256 struct {
	date      time.Time
	size      string
	name      string
	signature []byte
}

func NewFromFileSource(file *os.File, hashString string) Signature {
	sig := SignatureSHA256{}

	stat, _ := file.Stat()
	sig.size = strconv.FormatInt(stat.Size(), 10)
	sig.name = path.Base(file.Name())
	sig.date = stat.ModTime()

	var fileData = make([]byte, stat.Size())
	_, err := file.Read(fileData)
	if err != nil {
		panic(err)
	}

	data := strings.TrimSuffix(string(fileData), "\n") + hashString
	// fmt.Println("dataForSignString ", data)
	sig.signature = sig.encrypt(data)
	// fmt.Printf("%x \n", sig.signature)

	return &sig
}

func NewFromSignedFile(file *os.File) Signature {
	sig := SignatureSHA256{}
	fileData, err := ioutil.ReadFile(file.Name())
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fileData), separator)
	if len(lines) != 2 {
		panic("Sign file wrong coun lines != 2")
	}

	head := lines[0]
	sign := fileData[len(head)+len([]byte(separator)):]
	sig.signature = sign

	params := strings.Split(head, ":")
	sig.date, err = time.Parse(DATE_FORMAT, params[0])
	sig.size = params[1]
	sig.name = params[2]

	return &sig
}

func (sig *SignatureSHA256) encrypt(data string) []byte {
	s := sha256.New()
	s.Write([]byte(data))

	return s.Sum(nil)

}

func (sig *SignatureSHA256) Date() time.Time {
	return sig.date
}

func (sig *SignatureSHA256) Size() string {
	return sig.size
}

// source file
func (sig *SignatureSHA256) Name() string {
	return sig.name
}

const separator = "====sign===="

//yyyy-mm-dd hh-mm-ss
func (sig *SignatureSHA256) headString() string {
	return strings.Join(
		[]string{
			sig.Date().Format(DATE_FORMAT),
			sig.Size(),
			sig.Name(),
		}, ":")
}

func (sig *SignatureSHA256) SignatureByte() []byte {
	result := bytes.NewBufferString(sig.headString())
	result.WriteString(separator)
	result.Write(sig.signature)

	return result.Bytes()
}

func (sig *SignatureSHA256) Equals(s Signature) bool {
	if sig.name != s.Name() {
		return false
	}

	if sig.size != s.Size() {
		return false
	}

	if sig.date.Format(DATE_FORMAT) != s.Date().Format(DATE_FORMAT) {
		return false
	}

	if !bytes.Equal(sig.SignatureByte(), sig.SignatureByte()) {
		return false
	}

	return true
}
