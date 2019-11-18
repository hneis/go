package chat

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"io"
)

type MessageReader func(r io.Reader) Message

var (
	MessageSet = map[uint16]MessageReader{
		1: NewBroadcastMessageFromReader,
		2: NewTextMessageReader,
		3: NewChangeNameMessageFromReader,
		4: NewClientListMessageFromReader,
	}
)

func GetMessage(r io.Reader) (message Message, err error) {
	var mType uint16
	binary.Read(r, binary.BigEndian, &mType)
	messageReader, ok := MessageSet[mType]
	if !ok {
		err = fmt.Errorf("Message with type %d does not exists", mType)
		return
	}

	message = messageReader(r)

	return
}

type BroadcastMessage struct {
	User    string
	Message string
	mType   uint16
}

func (b BroadcastMessage) Type() uint16 {
	return b.mType
}

func NewBroadcastMessageFromReader(r io.Reader) Message {
	var size uint16
	binary.Read(r, binary.BigEndian, &size)
	dataRaw := make([]byte, size)
	binary.Read(r, binary.BigEndian, &dataRaw)

	binary.Read(r, binary.BigEndian, &size)
	user := make([]byte, size)
	binary.Read(r, binary.BigEndian, &user)
	return NewBroadcastMessage(string(dataRaw), string(user))
}

func NewBroadcastMessage(message string, user string) Message {
	return BroadcastMessage{
		mType:   1,
		Message: message,
		User:    user,
	}
}

func (b BroadcastMessage) Data() []byte {
	buffer := &bytes.Buffer{}
	bm := []byte(b.Message)
	binary.Write(buffer, binary.BigEndian, uint16(b.mType))
	binary.Write(buffer, binary.BigEndian, uint16(len(bm)))
	buffer.Write(bm)
	bm = []byte(b.User)
	binary.Write(buffer, binary.BigEndian, uint16(len(bm)))
	buffer.Write(bm)

	return buffer.Bytes()
}

type TextMessage struct {
	mType uint16
}

func NewTextMessageReader(r io.Reader) Message {
	return TextMessage{}
}

func (t TextMessage) Data() []byte {
	return []byte("data")
}

func (t TextMessage) Type() uint16 {
	return t.mType
}

type ChangeNameMessage struct {
	mType    uint16
	Username string
}

func NewChangeNameMessageFromReader(r io.Reader) Message {
	var size uint16
	binary.Read(r, binary.BigEndian, &size)
	dataRaw := make([]byte, size)
	binary.Read(r, binary.BigEndian, &dataRaw)

	return NewChangeNameMessage(string(dataRaw))
}

func NewChangeNameMessage(username string) Message {
	return ChangeNameMessage{
		mType:    3,
		Username: username,
	}
}

func (t ChangeNameMessage) Data() []byte {
	buffer := &bytes.Buffer{}
	bm := []byte(t.Username)
	binary.Write(buffer, binary.BigEndian, uint16(t.mType))
	binary.Write(buffer, binary.BigEndian, uint16(len(bm)))
	buffer.Write(bm)

	return buffer.Bytes()
}

func (t ChangeNameMessage) Type() uint16 {
	return t.mType
}

type ClientListMessage struct {
	mType   uint16
	Clients []string
}

func NewClientListMessageFromReader(r io.Reader) Message {
	tempClients := []string{}
	gob.NewDecoder(r).Decode(&tempClients)

	return NewClientListMessage(tempClients)
}

func NewClientListMessage(clients []string) Message {
	return ClientListMessage{
		mType:   4,
		Clients: clients,
	}
}

func (c ClientListMessage) Data() []byte {
	buf := &bytes.Buffer{}
	binary.Write(buf, binary.BigEndian, uint16(c.mType))
	gob.NewEncoder(buf).Encode(c.Clients)

	return buf.Bytes()
}

func (c ClientListMessage) Type() uint16 {
	return c.mType
}
