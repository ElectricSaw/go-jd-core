package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/classfile/attribute"
	"bitbucket.org/coontec/go-jd-core/class/model/classfile/constant"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
)

func NewUpdateOuterFieldTypeVisitor(typeMaker *utils.TypeMaker) *UpdateOuterFieldTypeVisitor {
	return &UpdateOuterFieldTypeVisitor{
		typeMaker: typeMaker,
	}
}

type UpdateOuterFieldTypeVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	typeMaker          *utils.TypeMaker
	searchFieldVisitor SearchFieldVisitor
}

func (v *UpdateOuterFieldTypeVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	bodyDeclaration := decl.(intsrv.IClassFileBodyDeclaration)
	if !bodyDeclaration.ClassFile().IsStatic() {
		v.SafeAcceptListDeclaration(ConvertMethodDeclarations(bodyDeclaration.MethodDeclarations()))
	}
	v.SafeAcceptListDeclaration(ConvertInnerTypeDeclarations(bodyDeclaration.InnerTypeDeclarations()))
}

func (v *UpdateOuterFieldTypeVisitor) VisitConstructorDeclaration(decl intmod.IConstructorDeclaration) {
	if !decl.IsStatic() {
		cfcd := decl.(intsrv.IClassFileConstructorDeclaration)

		if cfcd.ClassFile().OuterClassFile() != nil && !decl.IsStatic() {
			method := cfcd.Method()
			code := method.Attribute("Code").(attribute.AttributeCode).Code()
			offset := 0
			opcode := code[offset] & 255

			if opcode != 42 { // ALOAD_0
				return
			}

			offset++
			opcode = code[offset] & 255

			if opcode != 43 { // ALOAD_1
				return
			}
			offset++
			opcode = code[offset] & 255

			if opcode != 181 { // PUTFIELD
				return
			}

			index := 0
			offset++
			index = int(code[offset]&255) << 8
			offset++
			index = index | int(code[offset]&255)

			constants := method.Constants()
			constantMemberRef := constants.Constant(index).(constant.ConstantMemberRef)
			typeName, _ := constants.ConstantTypeName(constantMemberRef.ClassIndex())
			constantNameAndType := constants.Constant(constantMemberRef.NameAndTypeIndex()).(constant.ConstantNameAndType)
			descriptor, _ := constants.ConstantUtf8(constantNameAndType.DescriptorIndex())
			typeTypes := v.typeMaker.MakeTypeTypes(descriptor[1:])

			if (typeTypes != nil) && (typeTypes.TypeParameters != nil) {
				name, _ := constants.ConstantUtf8(constantNameAndType.NameIndex())
				v.searchFieldVisitor.Init(name)

				for _, field := range cfcd.BodyDeclaration().FieldDeclarations() {
					field.FieldDeclarators().Accept(&v.searchFieldVisitor)
					if v.searchFieldVisitor.Found() {
						var typeArguments intmod.ITypeArgument

						if typeTypes.TypeParameters.IsList() {
							tas := _type.NewTypeArgumentsWithSize(typeTypes.TypeParameters.Size())
							for _, typeParameter := range typeTypes.TypeParameters.List() {
								tas.Add(_type.NewGenericType(typeParameter.Identifier()))
							}
							typeArguments = tas
						} else {
							typeArguments = _type.NewGenericType(typeTypes.TypeParameters.First().Identifier())
						}

						// Update generic type of outer field reference
						v.typeMaker.SetFieldType(typeName, name, typeTypes.ThisType.CreateTypeWithArgs(typeArguments).(intmod.IType))
						break
					}
				}
			}
		}
	}
}

func (v *UpdateOuterFieldTypeVisitor) VisitMethodDeclaration(_ intmod.IMethodDeclaration) {
}

func (v *UpdateOuterFieldTypeVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
}

func (v *UpdateOuterFieldTypeVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateOuterFieldTypeVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *UpdateOuterFieldTypeVisitor) VisitAnnotationDeclaration(_ intmod.IAnnotationDeclaration) {
}

func (v *UpdateOuterFieldTypeVisitor) VisitEnumDeclaration(_ intmod.IEnumDeclaration) {
}

type SearchFieldVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	name  string
	found bool
}

func (v *SearchFieldVisitor) Init(name string) {
	v.name = name
	v.found = false
}

func (v *SearchFieldVisitor) Found() bool {
	return v.found
}

func (v *SearchFieldVisitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.found = v.found || decl.Name() == v.name
}
