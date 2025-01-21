package fragmenter

import (
	"github.com/ElectricSaw/go-jd-core/decompiler/api"
	"github.com/ElectricSaw/go-jd-core/decompiler/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/service/fragmenter/visitor"
)

func NewJavaSyntaxToJavaFragmentProcessor() *JavaSyntaxToJavaFragmentProcessor {
	return &JavaSyntaxToJavaFragmentProcessor{}
}

type JavaSyntaxToJavaFragmentProcessor struct {
}

func (p *JavaSyntaxToJavaFragmentProcessor) Process(message *model.Message) error {
	loader := message.Headers["loader"].(api.Loader)
	mainInternalTypeName := message.Headers["mainInternalTypeName"].(string)
	majorVersion := message.Headers["majorVersion"].(int)
	compilationUnit := message.Body.(*model.CompilationUnit)

	importsVisitor := visitor.NewSearchImportsVisitor(loader, mainInternalTypeName)
	importsVisitor.VisitCompilationUnit(compilationUnit)
	importsFragment := importsVisitor.ImportsFragment()
	message.Headers["maxLineNumber"] = importsVisitor.MaxLineNumber()

	compilationUnitVisitor := visitor.NewCompilationUnitVisitor(loader, mainInternalTypeName, majorVersion, importsFragment)
	compilationUnitVisitor.VisitCompilationUnit(compilationUnit)
	message.Body = compilationUnitVisitor.Fragments()

	return nil
}
