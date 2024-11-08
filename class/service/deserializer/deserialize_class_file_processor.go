package deserializer

import (
	"bitbucket.org/coontec/javaClass/class/api"
	"bitbucket.org/coontec/javaClass/class/model/message"
)

func NewDeserializeClassFileProcessor() *DeserializeClassFileProcessor {
	return &DeserializeClassFileProcessor{}
}

type DeserializeClassFileProcessor struct {
	ClassFileDeserializer
}

func (p *DeserializeClassFileProcessor) Process(message *message.Message) error {
	loader := message.Headers["loader"].(api.Loader)
	internalTypeName := message.Headers["mainInternalTypeName"].(string)
	classFile, err := p.LoadClassFileWithRaw(loader, internalTypeName)
	if err != nil {
		return err
	}

	message.Body = classFile

	return nil
}
