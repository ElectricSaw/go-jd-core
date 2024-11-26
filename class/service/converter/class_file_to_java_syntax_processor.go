package converter

import (
	"bitbucket.org/coontec/go-jd-core/class/model/message"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/processor"
)

var ConvertClassFileProcessor = processor.NewConvertClassFileProcessor()

func NewClassFileToJavaSyntaxProcessor() *ClassFileToJavaSyntaxProcessor {
	return &ClassFileToJavaSyntaxProcessor{}
}

type ClassFileToJavaSyntaxProcessor struct {
}

func (p *ClassFileToJavaSyntaxProcessor) Process(message *message.Message) error {
	return nil
}
