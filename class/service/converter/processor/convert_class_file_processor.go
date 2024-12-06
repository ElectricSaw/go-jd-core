package processor

import (
	intcls "bitbucket.org/coontec/go-jd-core/class/interfaces/classpath"
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/classfile/attribute"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/expression"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
	"bitbucket.org/coontec/go-jd-core/class/model/message"
	srvdecl "bitbucket.org/coontec/go-jd-core/class/service/converter/model/javasyntax/declaration"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/utils"
	"bitbucket.org/coontec/go-jd-core/class/service/converter/visitor"
	"bitbucket.org/coontec/go-jd-core/class/util"
)

var populateBindingsWithTypeParameterVisitor = &CustomPopulateBindingsWithTypeParameterVisitor{}

func NewConvertClassFileProcessor() *ConvertClassFileProcessor {
	return &ConvertClassFileProcessor{
		populateBindingsWithTypeParameterVisitor: visitor.NewPopulateBindingsWithTypeParameterVisitor(),
	}
}

type ConvertClassFileProcessor struct {
	populateBindingsWithTypeParameterVisitor *visitor.PopulateBindingsWithTypeParameterVisitor
}

func (p *ConvertClassFileProcessor) Process(message *message.Message) error {
	typeMaker := message.Headers["typeMaker"].(intsrv.ITypeMaker)
	classFile := message.Body.(intcls.IClassFile)

	annotationConverter := utils.NewAnnotationConverter(typeMaker)

	var typeDeclaration intmod.ITypeDeclaration

	if classFile.IsEnum() {
		typeDeclaration = p.convertEnumDeclaration(typeMaker, annotationConverter, classFile, nil)
	}

	message.Headers["majorVersion"] = classFile.MajorVersion()
	message.Headers["minorVersion"] = classFile.MinorVersion()
	message.Body = javasyntax.NewCompilationUnit(typeDeclaration)

	return nil
}

func (p *ConvertClassFileProcessor) convertInterfaceDeclaration(
	parser intsrv.ITypeMaker,
	converter *utils.AnnotationConverter,
	classFile intcls.IClassFile,
	outerClassFileBodyDeclaration intsrv.IClassFileBodyDeclaration) intsrv.IClassFileInterfaceDeclaration {
	annotationReferences := p.convertAnnotationReferencesWithClass(converter, classFile)
	typeTypes := parser.ParseClassFileSignature(classFile)
	bodyDeclaration := p.convertBodyDeclaration(parser, converter, classFile, typeTypes.TypeParameters(), outerClassFileBodyDeclaration)

	return srvdecl.NewClassFileInterfaceDeclaration(annotationReferences, classFile.AccessFlags(),
		typeTypes.ThisType().InternalName(), typeTypes.ThisType().Name(),
		typeTypes.TypeParameters(), typeTypes.Interfaces(), bodyDeclaration)
}

func (p *ConvertClassFileProcessor) convertEnumDeclaration(
	parser intsrv.ITypeMaker,
	converter *utils.AnnotationConverter,
	classFile intcls.IClassFile,
	outerClassFileBodyDeclaration intsrv.IClassFileBodyDeclaration) intsrv.IClassFileEnumDeclaration {
	annotationReferences := p.convertAnnotationReferencesWithClass(converter, classFile)
	typeTypes := parser.ParseClassFileSignature(classFile)
	bodyDeclaration := p.convertBodyDeclaration(parser, converter, classFile, typeTypes.TypeParameters(), outerClassFileBodyDeclaration)

	return srvdecl.NewClassFileEnumDeclaration(annotationReferences, classFile.AccessFlags(),
		typeTypes.ThisType().InternalName(), typeTypes.ThisType().Name(), typeTypes.Interfaces(), bodyDeclaration)
}

func (p *ConvertClassFileProcessor) convertAnnotationDeclaration(
	parser intsrv.ITypeMaker,
	converter *utils.AnnotationConverter,
	classFile intcls.IClassFile,
	outerClassFileBodyDeclaration intsrv.IClassFileBodyDeclaration) intsrv.IClassFileAnnotationDeclaration {
	annotationReferences := p.convertAnnotationReferencesWithClass(converter, classFile)
	typeTypes := parser.ParseClassFileSignature(classFile)
	bodyDeclaration := p.convertBodyDeclaration(parser, converter, classFile, typeTypes.TypeParameters(), outerClassFileBodyDeclaration)

	return srvdecl.NewClassFileAnnotationDeclaration(annotationReferences, classFile.AccessFlags(),
		typeTypes.ThisType().InternalName(), typeTypes.ThisType().Name(), bodyDeclaration)
}

func (p *ConvertClassFileProcessor) convertClassDeclaration(
	parser intsrv.ITypeMaker,
	converter *utils.AnnotationConverter,
	classFile intcls.IClassFile,
	outerClassFileBodyDeclaration intsrv.IClassFileBodyDeclaration) intsrv.IClassFileClassDeclaration {
	annotationReferences := p.convertAnnotationReferencesWithClass(converter, classFile)
	typeTypes := parser.ParseClassFileSignature(classFile)
	bodyDeclaration := p.convertBodyDeclaration(parser, converter, classFile, typeTypes.TypeParameters(), outerClassFileBodyDeclaration)

	return srvdecl.NewClassFileClassDeclaration(
		annotationReferences, classFile.AccessFlags(),
		typeTypes.ThisType().InternalName(), typeTypes.ThisType().Name(),
		typeTypes.TypeParameters(), typeTypes.SuperType(),
		typeTypes.Interfaces(), bodyDeclaration)
}

func (p *ConvertClassFileProcessor) convertBodyDeclaration(
	parser intsrv.ITypeMaker,
	converter *utils.AnnotationConverter,
	classFile intcls.IClassFile,
	typeParameters intmod.ITypeParameter,
	outerClassFileBodyDeclaration intsrv.IClassFileBodyDeclaration) intsrv.IClassFileBodyDeclaration {
	var bindings map[string]intmod.ITypeArgument
	var typeBounds map[string]intmod.IType

	if !classFile.IsStatic() && outerClassFileBodyDeclaration != nil {
		bindings = outerClassFileBodyDeclaration.Bindings()
		typeBounds = outerClassFileBodyDeclaration.TypeBounds()
	} else {
		bindings = make(map[string]intmod.ITypeArgument)
		typeBounds = make(map[string]intmod.IType)
	}

	if typeParameters != nil {
		p.populateBindingsWithTypeParameterVisitor.Init(make(map[string]intmod.ITypeArgument), make(map[string]intmod.IType))
		typeParameters.AcceptTypeParameterVisitor(p.populateBindingsWithTypeParameterVisitor)
	}

	bodyDeclaration := srvdecl.NewClassFileBodyDeclaration(classFile, bindings, typeBounds, outerClassFileBodyDeclaration)

	bodyDeclaration.SetFieldDeclarations(p.convertFields(parser, converter, classFile))
	bodyDeclaration.SetMethodDeclarations(p.convertMethods(parser, converter, bodyDeclaration, classFile))
	bodyDeclaration.SetInnerTypeDeclarations(p.convertInnerTypes(parser, converter, classFile, bodyDeclaration))

	return bodyDeclaration
}

func (p *ConvertClassFileProcessor) convertFields(
	parser intsrv.ITypeMaker,
	converter *utils.AnnotationConverter,
	classFile intcls.IClassFile) []intsrv.IClassFileFieldDeclaration {
	fields := classFile.Fields()

	if fields == nil {
		return nil
	} else {
		list := make([]intsrv.IClassFileFieldDeclaration, 0, len(fields))

		for _, field := range fields {
			annotationReference := p.convertAnnotationReferencesWithField(converter, field)
			typeField := parser.ParseFieldSignature(classFile, field)
			variableInitializer := p.convertFieldInitializer(field, typeField)
			fieldDeclarator := declaration.NewFieldDeclarator2(field.Name(), variableInitializer)
			list = append(list, srvdecl.NewClassFileFieldDeclaration3(annotationReference, field.AccessFlags(), typeField, fieldDeclarator))
		}

		return list
	}
}

func (p *ConvertClassFileProcessor) convertMethods(
	parser intsrv.ITypeMaker,
	converter *utils.AnnotationConverter,
	bodyDeclaration intsrv.IClassFileBodyDeclaration,
	classFile intcls.IClassFile) []intsrv.IClassFileConstructorOrMethodDeclaration {
	methods := classFile.Methods()

	if methods == nil {
		return nil
	} else {
		list := make([]intsrv.IClassFileConstructorOrMethodDeclaration, 0, len(methods))

		for _, method := range methods {
			name := method.Name()
			annotationReferences := p.convertAnnotationReferencesWithMethod(converter, method)
			annotationDefault := method.Attributes()["AnnotationDefault"].(intcls.IAttributeAnnotationDefault)
			var defaultAnnotationValue intmod.IElementValue

			if annotationDefault != nil {
				defaultAnnotationValue = converter.ConvertWithElementValue(annotationDefault.DefaultValue())
			}

			methodTypes := parser.ParseMethodSignature(classFile, method)
			var bindings map[string]intmod.ITypeArgument
			var typeBounds map[string]intmod.IType

			if (method.AccessFlags() & intcls.AccStatic) == 0 {
				bindings = bodyDeclaration.Bindings()
				typeBounds = bodyDeclaration.TypeBounds()
			} else {
				bindings = make(map[string]intmod.ITypeArgument)
				typeBounds = make(map[string]intmod.IType)
			}

			if methodTypes.TypeParameters() != nil {
				p.populateBindingsWithTypeParameterVisitor.Init(copyBindings(bindings), copyTypeBounds(typeBounds))
				methodTypes.TypeParameters().AcceptTypeParameterVisitor(p.populateBindingsWithTypeParameterVisitor)
			}

			code := method.Attributes()["Code"].(*attribute.AttributeCode)
			firstLineNumber := 0

			if code != nil {
				lineNumberTable := code.Attribute("LineNumberTable").(*attribute.AttributeLineNumberTable)
				if lineNumberTable != nil {
					firstLineNumber = lineNumberTable.LineNumberTable()[0].LineNumber()
				}
			}

			if name == "<init>" {
				list = append(list, srvdecl.NewClassFileConstructorDeclaration(
					bodyDeclaration, classFile, method, annotationReferences, methodTypes.TypeParameters(),
					methodTypes.ParameterTypes(), methodTypes.ExceptionTypes(), bindings, typeBounds, firstLineNumber))
			} else if name == "<clinit>" {
				list = append(list, srvdecl.NewClassFileStaticInitializerDeclaration(
					bodyDeclaration, classFile, method, bindings, typeBounds, firstLineNumber).(intsrv.IClassFileConstructorOrMethodDeclaration))
			} else {
				methodDeclaration := srvdecl.NewClassFileMethodDeclaration3(
					bodyDeclaration, classFile, method, annotationReferences, name, methodTypes.TypeParameters(),
					methodTypes.ReturnedType(), methodTypes.ParameterTypes(), methodTypes.ExceptionTypes(), defaultAnnotationValue,
					bindings, typeBounds, firstLineNumber)
				if classFile.IsInterface() {
					if methodDeclaration.Flags() == intcls.AccPublic {
						// For interfaces, add 'default' access flag on public methods
						methodDeclaration.SetFlags(intmod.FlagPublic | intmod.FlagDefault)
					}
				}

				list = append(list, methodDeclaration)
			}
		}

		return list
	}
}

func (p *ConvertClassFileProcessor) convertInnerTypes(
	parser intsrv.ITypeMaker,
	converter *utils.AnnotationConverter,
	classFile intcls.IClassFile,
	outerClassFileBodyDeclaration intsrv.IClassFileBodyDeclaration) []intsrv.IClassFileTypeDeclaration {
	innerClassFiles := classFile.InnerClassFiles()

	if innerClassFiles == nil {
		return nil
	} else {
		list := make([]intsrv.IClassFileTypeDeclaration, 0, len(innerClassFiles))

		for _, innerClassFile := range innerClassFiles {
			var innerTypeDeclaration intsrv.IClassFileTypeDeclaration

			if innerClassFile.IsEnum() {
				innerTypeDeclaration = p.convertEnumDeclaration(parser, converter, innerClassFile, outerClassFileBodyDeclaration)
			} else if innerClassFile.IsAnnotation() {
				innerTypeDeclaration = p.convertAnnotationDeclaration(parser, converter, innerClassFile, outerClassFileBodyDeclaration)
			} else if innerClassFile.IsInterface() {
				innerTypeDeclaration = p.convertInterfaceDeclaration(parser, converter, innerClassFile, outerClassFileBodyDeclaration)
			} else {
				innerTypeDeclaration = p.convertClassDeclaration(parser, converter, innerClassFile, outerClassFileBodyDeclaration)
			}

			list = append(list, innerTypeDeclaration)
		}

		return list
	}
}

func (p *ConvertClassFileProcessor) convertAnnotationReferencesWithClass(
	converter *utils.AnnotationConverter, classFile intcls.IClassFile) intmod.IAnnotationReference {
	visibles := classFile.Attribute("RuntimeVisibleAnnotations").(*attribute.Annotations)
	invisibles := classFile.Attribute("RuntimeInvisibleAnnotations").(*attribute.Annotations)

	return converter.ConvertWithAnnotations2(visibles, invisibles)
}

func (p *ConvertClassFileProcessor) convertAnnotationReferencesWithField(
	converter *utils.AnnotationConverter, field intcls.IField) intmod.IAnnotationReference {
	visibles := field.Attribute("RuntimeVisibleAnnotations").(*attribute.Annotations)
	invisibles := field.Attribute("RuntimeInvisibleAnnotations").(*attribute.Annotations)

	return converter.ConvertWithAnnotations2(visibles, invisibles)
}

func (p *ConvertClassFileProcessor) convertAnnotationReferencesWithMethod(
	converter *utils.AnnotationConverter, method intcls.IMethod) intmod.IAnnotationReference {
	visibles := method.Attribute("RuntimeVisibleAnnotations").(*attribute.Annotations)
	invisibles := method.Attribute("RuntimeInvisibleAnnotations").(*attribute.Annotations)

	return converter.ConvertWithAnnotations2(visibles, invisibles)
}

func (p *ConvertClassFileProcessor) convertFieldInitializer(field intcls.IField, typeField intmod.IType) intmod.IExpressionVariableInitializer {
	acv := field.Attribute("ConstantValue").(*attribute.AttributeConstantValue)

	if acv == nil {
		return nil
	} else {
		constantValue := acv.ConstantValue()
		var expr intmod.IExpression

		switch constantValue.Tag() {
		case intcls.ConstTagInteger:
			expr = expression.NewIntegerConstantExpression(typeField, constantValue.(intcls.IConstantInteger).Value())
		case intcls.ConstTagFloat:
			expr = expression.NewFloatConstantExpression(constantValue.(intcls.IConstantFloat).Value())
		case intcls.ConstTagLong:
			expr = expression.NewLongConstantExpression(constantValue.(intcls.IConstantLong).Value())
		case intcls.ConstTagDouble:
			expr = expression.NewDoubleConstantExpression(constantValue.(intcls.IConstantDouble).Value())
		case intcls.ConstTagUtf8:
			expr = expression.NewStringConstantExpression(constantValue.(intcls.IConstantUtf8).Value())
		default:
			return nil
		}

		return declaration.NewExpressionVariableInitializer(expr)
	}
}

func (p *ConvertClassFileProcessor) convertModuleDeclaration(classFile intcls.IClassFile) intmod.IModuleDeclaration {
	attributeModule := classFile.Attribute("Module").(*attribute.AttributeModule)
	requires := p.convertModuleDeclarationModuleInfo(attributeModule.Requires())
	exports := p.convertModuleDeclarationPackageInfo(attributeModule.Exports())
	opens := p.convertModuleDeclarationPackageInfo(attributeModule.Opens())
	uses := util.NewDefaultListWithCapacity[string](len(attributeModule.Uses()))
	for _, use := range attributeModule.Uses() {
		uses.Add(use)
	}
	provides := p.convertModuleDeclarationServiceInfo(attributeModule.Provides())

	return declaration.NewModuleDeclaration(
		attributeModule.Flags(), classFile.InternalTypeName(), attributeModule.Name(),
		attributeModule.Version(), requires, exports, opens, uses, provides)
}

func (p *ConvertClassFileProcessor) convertModuleDeclarationModuleInfo(moduleInfos []intcls.IModuleInfo) util.IList[intmod.IModuleInfo] {
	if (moduleInfos == nil) || (len(moduleInfos) == 0) {
		return nil
	} else {
		list := util.NewDefaultListWithCapacity[intmod.IModuleInfo](len(moduleInfos))
		for _, moduleInfo := range moduleInfos {
			list.Add(declaration.NewModuleInfo(moduleInfo.Name(), moduleInfo.Flags(), moduleInfo.Version()))
		}
		return list
	}
}

func (p *ConvertClassFileProcessor) convertModuleDeclarationPackageInfo(packageInfos []intcls.IPackageInfo) util.IList[intmod.IPackageInfo] {
	if (packageInfos == nil) || (len(packageInfos) == 0) {
		return nil
	} else {
		list := util.NewDefaultListWithCapacity[intmod.IPackageInfo](len(packageInfos))
		for _, packageInfo := range packageInfos {
			var moduleInfoNames []string
			if packageInfo.ModuleInfoNames() != nil {
				moduleInfoNames = copyStrings(packageInfo.ModuleInfoNames())
			}
			list.Add(declaration.NewPackageInfo(packageInfo.InternalName(), packageInfo.Flags(), moduleInfoNames))
		}
		return list
	}
}

func (p *ConvertClassFileProcessor) convertModuleDeclarationServiceInfo(serviceInfos []intcls.IServiceInfo) util.IList[intmod.IServiceInfo] {
	if (serviceInfos == nil) || (len(serviceInfos) == 0) {
		return nil
	} else {
		list := util.NewDefaultListWithCapacity[intmod.IServiceInfo](len(serviceInfos))
		for _, serviceInfo := range serviceInfos {
			var implementationTypeNames []string
			if serviceInfo.ImplementationTypeNames() != nil {
				implementationTypeNames = copyStrings(serviceInfo.ImplementationTypeNames())
			}
			list.Add(declaration.NewServiceInfo(serviceInfo.InterfaceTypeName(), implementationTypeNames))
		}
		return list
	}
}

type CustomPopulateBindingsWithTypeParameterVisitor struct {
	visitor.PopulateBindingsWithTypeParameterVisitor
}

func (v *CustomPopulateBindingsWithTypeParameterVisitor) VisitTypeParameter(parameter intmod.ITypeParameter) {
	v.Bindings[parameter.Identifier()] = _type.NewGenericTypeWithAll(parameter.Identifier(), 0)
}

func (v *CustomPopulateBindingsWithTypeParameterVisitor) VisitTypeParameterWithTypeBounds(parameter intmod.ITypeParameterWithTypeBounds) {
	v.Bindings[parameter.Identifier()] = _type.NewGenericTypeWithAll(parameter.Identifier(), 0)
	v.TypeBounds[parameter.Identifier()] = parameter.TypeBounds()
}

func copyStrings(names []string) []string {
	ret := make([]string, 0, len(names))
	for _, value := range names {
		ret = append(ret, value)
	}
	return ret
}

func copyBindings(bindings map[string]intmod.ITypeArgument) map[string]intmod.ITypeArgument {
	ret := make(map[string]intmod.ITypeArgument)
	for key, value := range bindings {
		ret[key] = value
	}
	return ret
}

func copyTypeBounds(typeBounds map[string]intmod.IType) map[string]intmod.IType {
	ret := make(map[string]intmod.IType)
	for key, value := range typeBounds {
		ret[key] = value
	}
	return ret
}
