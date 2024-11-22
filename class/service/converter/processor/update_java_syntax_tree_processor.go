package processor

import (
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/model/message"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/visitor"
)

type UpdateJavaSyntaxTreeProcessor struct {
}

func (p *UpdateJavaSyntaxTreeProcessor) Process(message *message.Message) error {
	typeMaker := message.Headers["typeMaker"]
	compilationUnit := message.Body.(*javasyntax.CompilationUnit)

	visitor.NewUpdateJavaSyntaxTreeStep0Visitor(typeMaker).Visit(compilationUnit)
	visitor.NewUpdateJavaSyntaxTreeStep1Visitor(typeMaker).Visit(compilationUnit)
	visitor.NewUpdateJavaSyntaxTreeStep2Visitor(typeMaker).Visit(compilationUnit)

	return nil
}
