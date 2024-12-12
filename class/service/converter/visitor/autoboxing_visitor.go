package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"strings"
)

func NewAutoboxingVisitor() intmod.IJavaSyntaxVisitor {
	return &AutoboxingVisitor{}
}

var ValueOfDescriptorMap = map[string]string{
	"java/lang/Byte":      "(B)Ljava/lang/Byte;",
	"java/lang/Character": "(C)Ljava/lang/Character;",
	"java/lang/Float":     "(F)Ljava/lang/Float;",
	"java/lang/Integer":   "(I)Ljava/lang/Integer;",
	"java/lang/Long":      "(J)Ljava/lang/Long;",
	"java/lang/Short":     "(S)Ljava/lang/Short;",
	"java/lang/Double":    "(D)Ljava/lang/Double;",
	"java/lang/Boolean":   "(Z)Ljava/lang/Boolean;",
}

var ValueDescriptorMap = map[string]string{
	"java/lang/Byte":      "()B",
	"java/lang/Character": "()C",
	"java/lang/Float":     "()F",
	"java/lang/Integer":   "()I",
	"java/lang/Long":      "()J",
	"java/lang/Short":     "()S",
	"java/lang/Double":    "()D",
	"java/lang/Boolean":   "()Z",
}

var ValueMethodNameMap = map[string]string{
	"java/lang/Byte":      "byteValue",
	"java/lang/Character": "charValue",
	"java/lang/Float":     "floatValue",
	"java/lang/Integer":   "intValue",
	"java/lang/Long":      "longValue",
	"java/lang/Short":     "shortValue",
	"java/lang/Double":    "doubleValue",
	"java/lang/Boolean":   "booleanValue",
}

type AutoboxingVisitor struct {
	AbstractUpdateExpressionVisitor
}

func (v *AutoboxingVisitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	cfbd := declaration.(intsrv.IClassFileBodyDeclaration)
	autoBoxingSupported := cfbd.ClassFile().MajorVersion() >= 49 // (majorVersion >= Java 5)

	if autoBoxingSupported {
		v.SafeAcceptDeclaration(declaration.MemberDeclarations())
	}
}

func (v *AutoboxingVisitor) updateExpression(expression intmod.IExpression) intmod.IExpression {
	if expression.IsMethodInvocationExpression() && strings.HasPrefix(expression.InternalTypeName(), "java/lang/") {
		var parameterSize int

		if expression.Parameters() != nil {
			parameterSize = expression.Parameters().Size()
		}

		if expression.Expression().IsObjectTypeReferenceExpression() {
			// static method invocation
			if parameterSize == 1 &&
				expression.Name() == "valueOf" &&
				expression.Descriptor() == ValueOfDescriptorMap[expression.InternalTypeName()] {
				return expression.Parameters().First()
			}
		} else {
			// non-static method invocation
			if (parameterSize == 0) &&
				expression.Name() == ValueMethodNameMap[expression.InternalTypeName()] &&
				expression.Descriptor() == ValueDescriptorMap[expression.InternalTypeName()] {
				return expression.Expression()
			}
		}
	}

	return expression
}
