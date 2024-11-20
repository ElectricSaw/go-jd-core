package processor

import (
	"bitbucket.org/coontec/go-jd-core/class/model/message"
)

type UpdateJavaSyntaxTreeProcessor struct {
}

func (p *UpdateJavaSyntaxTreeProcessor) Process(message *message.Message) error {
	//typeMaker := message.Headers["typeMaker"]
	//compilationUnit := message.Body.(*model.CompilationUnit)

	//visitor.NewUpdateJavaSyntaxTreeStep0Visitor(typeMaker).Visit(compilationUnit)
	//visitor.NewUpdateJavaSyntaxTreeStep1Visitor(typeMaker).Visit(compilationUnit)
	//visitor.NewUpdateJavaSyntaxTreeStep2Visitor(typeMaker).Visit(compilationUnit)

	return nil
}
