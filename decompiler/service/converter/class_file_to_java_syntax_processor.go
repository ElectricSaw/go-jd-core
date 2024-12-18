package converter

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/api"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/message"
	"github.com/ElectricSaw/go-jd-core/decompiler/service/converter/processor"
	"github.com/ElectricSaw/go-jd-core/decompiler/service/converter/visitor"
)

var ConvertClassFileProcessor = processor.NewConvertClassFileProcessor()
var UpdateJavaSyntaxTreeProcessor = processor.NewUpdateJavaSyntaxTreeProcessor()

func NewClassFileToJavaSyntaxProcessor() *ClassFileToJavaSyntaxProcessor {
	return &ClassFileToJavaSyntaxProcessor{}
}

type ClassFileToJavaSyntaxProcessor struct {
}

func (p *ClassFileToJavaSyntaxProcessor) Process(message *message.Message) error {
	loader := message.Headers["loader"].(api.Loader)
	configuration := message.Headers["configuration"].(map[string]interface{})

	if configuration == nil {
		message.Headers["typeMaker"] = visitor.NewTypeMaker(loader)
	} else {
		typeMaker, ok := configuration["typeMaker"]

		if !ok {
			// Store the heavy weight object 'typeMaker' in 'configuration' to reuse it
			typeMaker = visitor.NewTypeMaker(loader)
			configuration["typeMaker"] = typeMaker
		}

		message.Headers["typeMaker"] = typeMaker
	}

	if err := ConvertClassFileProcessor.Process(message); err != nil {
		return err
	}
	if err := UpdateJavaSyntaxTreeProcessor.Process(message); err != nil {
		return err
	}

	return nil
}
