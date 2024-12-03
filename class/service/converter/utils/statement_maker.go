package utils

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/visitor"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

type IStatementMaker interface {
}

var GlobalFinallyExceptionExpression = expression.NewNullExpression(
	_type.NewObjectType("java/lang/Exception",
		"java.lang.Exception", "Exception"))
var GlobalMergeTryWithResourcesStatementVisitor = visitor.NewMergeTryWithResourcesStatementVisitor()

func NewStatementMaker(typeMaker intsrv.ITypeMaker, localVariableMaker intsrv.ILocalVariableMaker,
	comd intsrv.IClassFileConstructorOrMethodDeclaration) IStatementMaker {
	classFile := comd.ClassFile()

	m := &StatementMaker{
		typeMaker:                             typeMaker,
		typeBounds:                            comd.TypeBounds(),
		localVariableMaker:                    localVariableMaker,
		majorVersion:                          classFile.MajorVersion(),
		internalTypeName:                      classFile.InternalTypeName(),
		bodyDeclaration:                       comd.BodyDeclaration(),
		stack:                                 util.NewDefaultStack[intmod.IExpression](),
		byteCodeParser:                        NewByteCodeParser(typeMaker, localVariableMaker, classFile, comd.BodyDeclaration(), comd),
		removeFinallyStatementsVisitor:        visitor.NewRemoveFinallyStatementsVisitor(localVariableMaker),
		removeBinaryOpReturnStatementsVisitor: visitor.NewRemoveBinaryOpReturnStatementsVisitor(localVariableMaker),
		updateIntegerConstantTypeVisitor:      visitor.NewUpdateIntegerConstantTypeVisitor(comd.ReturnedType()),
		searchFirstLineNumberVisitor:          visitor.NewSearchFirstLineNumberVisitor(),
		memberVisitor:                         NewByteCodeParserMemberVisitor(),
		removeFinallyStatementsFlag:           false,
		mergeTryWithResourcesStatementFlag:    false,
	}

	return m
}

type StatementMaker struct {
	typeMaker                             intsrv.ITypeMaker
	typeBounds                            map[string]intmod.IType
	localVariableMaker                    intsrv.ILocalVariableMaker
	byteCodeParser                        intsrv.IByteCodeParser
	majorVersion                          int
	internalTypeName                      string
	bodyDeclaration                       intsrv.IClassFileBodyDeclaration
	stack                                 util.IStack[intmod.IExpression]
	removeFinallyStatementsVisitor        *visitor.RemoveFinallyStatementsVisitor
	removeBinaryOpReturnStatementsVisitor *visitor.RemoveBinaryOpReturnStatementsVisitor
	updateIntegerConstantTypeVisitor      *visitor.UpdateIntegerConstantTypeVisitor
	searchFirstLineNumberVisitor          *visitor.SearchFirstLineNumberVisitor
	memberVisitor                         *StatementMakerMemberVisitor
	removeFinallyStatementsFlag           bool
	mergeTryWithResourcesStatementFlag    bool
}

type SwitchCaseComparator struct {
}

func (c *SwitchCaseComparator) Compare(sc1, sc2 intsrv.ISwitchCase) int {
	diff := sc1.Offset() - sc2.Offset()

	if diff != 0 {
		return diff
	}

	return sc1.Value() - sc2.Value()
}

type StatementMakerMemberVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	name  string
	found bool
}

func (v *StatementMakerMemberVisitor) Init(name string) {
	v.name = name
	v.found = false
}

func (v *StatementMakerMemberVisitor) Found() bool {
	return v.found
}

func (v *StatementMakerMemberVisitor) VisitFieldDeclarator(declaration intmod.IFieldDeclarator) {
	v.found = v.found || declaration.Name() == v.name
}

func (v *StatementMakerMemberVisitor) VisitMethodDeclaration(declaration intmod.IMethodDeclaration) {
	v.found = v.found || declaration.Name() == v.name
}
