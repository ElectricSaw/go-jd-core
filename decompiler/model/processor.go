package model

type Processor interface {
	Process(message *Message) error
}
