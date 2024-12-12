package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/classfile/attribute"
	"github.com/ElectricSaw/go-jd-core/class/model/classfile/constant"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	"strings"
)

func NewUpdateBridgeMethodTypeVisitor(typeMaker intsrv.ITypeMaker) intsrv.IUpdateBridgeMethodTypeVisitor {
	return &UpdateBridgeMethodTypeVisitor{
		typeMaker: typeMaker,
	}
}

type UpdateBridgeMethodTypeVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	typeMaker intsrv.ITypeMaker
}

func (v *UpdateBridgeMethodTypeVisitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	bodyDeclaration := declaration.(intsrv.IClassFileBodyDeclaration)

	v.SafeAcceptListDeclaration(ConvertMethodDeclarations(bodyDeclaration.MethodDeclarations()))
	v.SafeAcceptListDeclaration(ConvertTypeDeclarations(bodyDeclaration.InnerTypeDeclarations()))
}

func (v *UpdateBridgeMethodTypeVisitor) VisitMethodDeclaration(declaration intmod.IMethodDeclaration) {
	if declaration.IsStatic() && declaration.ReturnedType().IsObjectType() && strings.HasPrefix(declaration.Name(), "access$") {
		typeTypes := v.typeMaker.MakeTypeTypes(declaration.ReturnedType().InternalName())

		if (typeTypes != nil) && (typeTypes.TypeParameters != nil) {
			cfmd := declaration.(intsrv.IClassFileMethodDeclaration)
			method := cfmd.Method()
			code := method.Attributes()["Code"].(attribute.AttributeCode).Code()
			offset := 0
			opcode := code[offset] & 255

			for (21 <= opcode && opcode <= 45) || // ILOAD, LLOAD, FLOAD, DLOAD, ..., ILOAD_0 ... ILOAD_3, ..., ALOAD_1, ..., ALOAD_3
				(89 <= opcode && opcode <= 95) { // DUP, ..., DUP2_X2, SWAP
				offset++
				opcode = code[offset] & 255
			}

			switch opcode {
			case 178, 179, 180, 181: // GETSTATIC, PUTSTATIC, GETFIELD, PUTFIELD
				offset++
				r1 := int(code[offset]&255) << 8
				offset++
				r2 := int(code[offset]&255) << 8
				index := r1 | r2

				constants := method.Constants()
				constantMemberRef := constants.Constant(index).(constant.ConstantMemberRef)
				typeName, _ := constants.ConstantTypeName(constantMemberRef.ClassIndex())
				constantNameAndType := constants.Constant(constantMemberRef.NameAndTypeIndex()).(constant.ConstantNameAndType)
				name, _ := constants.ConstantUtf8(constantNameAndType.NameIndex())
				descriptor, _ := constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
				typ := v.typeMaker.MakeFieldType(typeName, name, descriptor)
				// Update returned generic type of bridge method
				v.typeMaker.SetMethodReturnedType(typeName, cfmd.Name(), cfmd.Descriptor(), typ)
			case 182, 183, 184, 185: // INVOKEVIRTUAL, INVOKESPECIAL, INVOKESTATIC, INVOKEINTERFACE
				offset++
				r1 := int(code[offset]&255) << 8
				offset++
				r2 := int(code[offset]&255) << 8
				index := r1 | r2

				constants := method.Constants()
				constantMemberRef := constants.Constant(index).(constant.ConstantMemberRef)
				typeName, _ := constants.ConstantTypeName(constantMemberRef.ClassIndex())
				constantNameAndType := constants.Constant(constantMemberRef.NameAndTypeIndex()).(constant.ConstantNameAndType)
				name, _ := constants.ConstantUtf8(constantNameAndType.NameIndex())
				descriptor, _ := constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
				methodTypes := v.typeMaker.MakeMethodTypes2(typeName, name, descriptor)

				// Update returned generic type of bridge method
				v.typeMaker.SetMethodReturnedType(typeName, cfmd.Name(), cfmd.Descriptor(), methodTypes.ReturnedType())
			}
		}
	}
}

func (v *UpdateBridgeMethodTypeVisitor) VisitConstructorDeclaration(_ intmod.IConstructorDeclaration) {
}
func (v *UpdateBridgeMethodTypeVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
}

func (v *UpdateBridgeMethodTypeVisitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *UpdateBridgeMethodTypeVisitor) VisitAnnotationDeclaration(_ intmod.IAnnotationDeclaration) {
}
func (v *UpdateBridgeMethodTypeVisitor) VisitEnumDeclaration(_ intmod.IEnumDeclaration) {}

func ConvertMethodDeclarations(list []intsrv.IClassFileConstructorOrMethodDeclaration) []intmod.IDeclaration {
	ret := make([]intmod.IDeclaration, 0, len(list))
	for _, item := range list {
		ret = append(ret, item)
	}
	return ret
}

func ConvertFieldDeclarations(list []intsrv.IClassFileFieldDeclaration) []intmod.IDeclaration {
	ret := make([]intmod.IDeclaration, 0, len(list))
	for _, item := range list {
		ret = append(ret, item)
	}
	return ret
}

func ConvertTypeDeclarations(list []intsrv.IClassFileTypeDeclaration) []intmod.IDeclaration {
	ret := make([]intmod.IDeclaration, 0, len(list))
	for _, item := range list {
		ret = append(ret, item)
	}
	return ret
}

func ConvertMemberDeclaration(list []intsrv.IClassFileMemberDeclaration) []intmod.IDeclaration {
	ret := make([]intmod.IDeclaration, 0, len(list))
	for _, item := range list {
		ret = append(ret, item)
	}
	return ret
}
