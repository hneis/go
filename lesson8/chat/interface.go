package chat

type Message interface {
	Data() []byte
	Type() uint16
}
