package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/classfile/attribute"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
	"strings"
)

func NewCreateInstructionsVisitor(typeMaker intsrv.ITypeMaker) *CreateInstructionsVisitor {
	return &CreateInstructionsVisitor{
		typeMaker: typeMaker,
	}
}

type CreateInstructionsVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	typeMaker intsrv.ITypeMaker
}


func (v *CreateInstructionsVisitor) VisitAnnotationDeclaration( decl intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration());
}

func (v *CreateInstructionsVisitor) VisitBodyDeclaration( decl intmod.IBodyDeclaration) {
	bodyDeclaration := decl.(intsrv.IClassFileBodyDeclaration)

	// Parse byte code
	methods := bodyDeclaration.MethodDeclarations();

	if methods != nil {
		for _, method := range methods {
			if (method.Flags() & (intmod.FlagSynthetic|intmod.FlagBridge)) != 0 {
				method.Accept(v);
			} else if (method.Flags() & (intmod.FlagStatic|intmod.FlagBridge)) == intmod.FlagStatic {
				if strings.HasPrefix(method.Method().Name(), "access$") {
					// Accessor -> bridge method
					method.SetFlags(method.Flags() | intmod.FlagBridge);
					method.Accept(v);
				}
			} else if method.ParameterTypes() != nil {
				if method.ParameterTypes().IsList() {
					for _, item  := range method.ParameterTypes().ToSlice() {
						if item.IsObjectType() && (item.Name() == "") {
						// Synthetic type in parameters -> synthetic method
						method.SetFlags(method.Flags() | intmod.FlagSynthetic);
						method.Accept(v);
						break;
						}
					}
				} else {
					typ := method.ParameterTypes().First();
					if typ.IsObjectType() && (typ.Name() == "") {
					// Synthetic type in parameters -> synthetic method
					method.SetFlags(method.Flags() | intmod.FlagSynthetic);
					method.Accept(v);
					break;
					}
				}
			}
		}

		for _, method := range methods {
			if (method.Flags() & (intmod.FlagSynthetic|intmod.FlagBridge)) == 0 {
				method.Accept(v);
			}
		}
	}
}

func (v *CreateInstructionsVisitor) VisitFieldDeclaration( _ intmod.IFieldDeclaration) {}

func (v *CreateInstructionsVisitor) VisitConstructorDeclaration( decl intmod.IConstructorDeclaration) {
	v.createParametersVariablesAndStatements(decl.(intsrv.IClassFileConstructorOrMethodDeclaration), true);
}

func (v *CreateInstructionsVisitor) VisitMethodDeclaration( decl intmod.IMethodDeclaration) {
	v.createParametersVariablesAndStatements(decl.(intsrv.IClassFileConstructorOrMethodDeclaration), false);
}

func (v *CreateInstructionsVisitor) VisitStaticInitializerDeclaration( decl intmod.IStaticInitializerDeclaration) {
	v.createParametersVariablesAndStatements(decl.(intsrv.IClassFileConstructorOrMethodDeclaration), false);
}

func (v *CreateInstructionsVisitor) VisitClassDeclaration( decl intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration());
}

func (v *CreateInstructionsVisitor) VisitEnumDeclaration( decl intmod.IEnumDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration());
}

func (v *CreateInstructionsVisitor) VisitInterfaceDeclaration( decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration());
}

func (v *CreateInstructionsVisitor) createParametersVariablesAndStatements(comd intsrv.IClassFileConstructorOrMethodDeclaration, constructor bool) {
classFile := comd.ClassFile();
method := comd.Method();
attributeCode := method.Attribute("Code").(*attribute.AttributeCode)
localVariableMaker := utils.NewLocalVariableMaker(v.typeMaker, comd, constructor);

if attributeCode == nil {
localVariableMaker.Make(false, v.typeMaker);
} else {
statementMaker := utils.NewStatementMaker(v.typeMaker, localVariableMaker, comd);
containsLineNumber := attributeCode.Attribute("LineNumberTable") != nil;

ControlFlowGraph cfg = ControlFlowGraphMaker.make(method);

if (cfg != nil) {
ControlFlowGraphGotoReducer.reduce(cfg);
ControlFlowGraphLoopReducer.reduce(cfg);

if (ControlFlowGraphReducer.reduce(cfg)) {
comd.setStatements(statementMaker.make(cfg));
} else {
comd.setStatements(new ByteCodeStatement(ByteCodeWriter.write("// ", method)));
}
}
} catch (Exception e) {
assert ExceptionUtil.printStackTrace(e);
comd.setStatements(new ByteCodeStatement(ByteCodeWriter.write("// ", method)));
}

localVariableMaker.make(containsLineNumber, typeMaker);
}

comd.setFormalParameters(localVariableMaker.FormalParameters());

if (classFile.isInterface()) {
comd.setFlags(comd.Flags() & ~(FLAG_PUBLIC|FLAG_ABSTRACT));
}
}