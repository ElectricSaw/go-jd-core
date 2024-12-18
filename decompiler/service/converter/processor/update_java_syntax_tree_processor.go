package processor

import (
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/message"
	"github.com/ElectricSaw/go-jd-core/decompiler/service/converter/visitor"
)

func NewUpdateJavaSyntaxTreeProcessor() *UpdateJavaSyntaxTreeProcessor {
	return &UpdateJavaSyntaxTreeProcessor{}
}

type UpdateJavaSyntaxTreeProcessor struct {
}

func (p *UpdateJavaSyntaxTreeProcessor) Process(message *message.Message) error {
	typeMaker := message.Headers["typeMaker"].(intsrv.ITypeMaker)
	compilationUnit := message.Body.(*javasyntax.CompilationUnit)

	visitor.NewUpdateJavaSyntaxTreeStep0Visitor(typeMaker).VisitCompilationUnit(compilationUnit)
	visitor.NewUpdateJavaSyntaxTreeStep1Visitor(typeMaker).VisitCompilationUnit(compilationUnit)
	visitor.NewUpdateJavaSyntaxTreeStep2Visitor(typeMaker).VisitCompilationUnit(compilationUnit)

	return nil
}
