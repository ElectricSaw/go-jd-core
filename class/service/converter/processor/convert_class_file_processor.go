package processor

import (
	_type "bitbucket.org/coontec/javaClass/class/model/javasyntax/type"
	"bitbucket.org/coontec/javaClass/class/service/converter/visitor"
)

var populateBindingsWithTypeParameterVisitor = &CustomPopulateBindingsWithTypeParameterVisitor{}

type ConvertClassFileProcessor struct {
}

//func (p *ConvertClassFileProcessor) Process(message *message.Message) error {
//	typeMaker := message.Headers["typeMaker"].(*utils.TypeMaker)
//	classFile := message.Body.(*classfile.ClassFile)
//
//	annotationConverter := utils.NewAnnotationConverter(typeMaker)
//
//	var typeDeclaration declaration.ITypeDeclaration
//
//	if classFile.IsEnum() {
//		typeDeclaration = p.convertEnumDeclaration(typeMaker, annotationConverter, classFile, nil)
//	}
//
//	message.Headers["majorVersion"] = classFile.MajorVersion()
//	message.Headers["minorVersion"] = classFile.MinorVersion()
//	message.Body = javasyntax.NewCompilationUnit(typeDeclaration)
//
//	return nil
//}

//func (p *ConvertClassFileProcessor) convertInterfaceDeclaration(
//	parser utils.TypeMaker,
//	converter utils.AnnotationConverter,
//	classFile *classfile.ClassFile,
//	outerClassFileBodyDeclaration servdecl.ClassFileBodyDeclaration) ClassFileInterfaceDeclaration {
//BaseAnnotationReference annotationReferences = convertAnnotationReferences(converter, classFile);
//TypeMaker.TypeTypes typeTypes = parser.parseClassFileSignature(classFile);
//ClassFileBodyDeclaration bodyDeclaration = convertBodyDeclaration(parser, converter, classFile, typeTypes.typeParameters, outerClassFileBodyDeclaration);
//
//return new ClassFileInterfaceDeclaration(
//annotationReferences, classFile.getAccessFlags(),
//typeTypes.thisType.getInternalName(), typeTypes.thisType.getName(),
//typeTypes.typeParameters, typeTypes.interfaces, bodyDeclaration);
//}
//
//func (p *ConvertClassFileProcessor) convertEnumDeclaration(TypeMaker parser, AnnotationConverter converter, ClassFile classFile, ClassFileBodyDeclaration outerClassFileBodyDeclaration) ClassFileEnumDeclaration {
//BaseAnnotationReference annotationReferences = convertAnnotationReferences(converter, classFile);
//TypeMaker.TypeTypes typeTypes = parser.parseClassFileSignature(classFile);
//ClassFileBodyDeclaration bodyDeclaration = convertBodyDeclaration(parser, converter, classFile, typeTypes.typeParameters, outerClassFileBodyDeclaration);
//
//return new ClassFileEnumDeclaration(
//annotationReferences, classFile.getAccessFlags(),
//typeTypes.thisType.getInternalName(), typeTypes.thisType.getName(),
//typeTypes.interfaces, bodyDeclaration);
//}
//
//func (p *ConvertClassFileProcessor) convertAnnotationDeclaration(TypeMaker parser, AnnotationConverter converter, ClassFile classFile, ClassFileBodyDeclaration outerClassFileBodyDeclaration) ClassFileAnnotationDeclaration {
//BaseAnnotationReference annotationReferences = convertAnnotationReferences(converter, classFile);
//TypeMaker.TypeTypes typeTypes = parser.parseClassFileSignature(classFile);
//ClassFileBodyDeclaration bodyDeclaration = convertBodyDeclaration(parser, converter, classFile, typeTypes.typeParameters, outerClassFileBodyDeclaration);
//
//return new ClassFileAnnotationDeclaration(
//annotationReferences, classFile.getAccessFlags(),
//typeTypes.thisType.getInternalName(), typeTypes.thisType.getName(),
//bodyDeclaration);
//}
//
//func (p *ConvertClassFileProcessor) convertClassDeclaration(TypeMaker parser, AnnotationConverter converter, ClassFile classFile, ClassFileBodyDeclaration outerClassFileBodyDeclaration) ClassFileClassDeclaration {
//BaseAnnotationReference annotationReferences = convertAnnotationReferences(converter, classFile);
//TypeMaker.TypeTypes typeTypes = parser.parseClassFileSignature(classFile);
//ClassFileBodyDeclaration bodyDeclaration = convertBodyDeclaration(parser, converter, classFile, typeTypes.typeParameters, outerClassFileBodyDeclaration);
//
//return new ClassFileClassDeclaration(
//annotationReferences, classFile.getAccessFlags(),
//typeTypes.thisType.getInternalName(), typeTypes.thisType.getName(),
//typeTypes.typeParameters, typeTypes.superType,
//typeTypes.interfaces, bodyDeclaration);
//}
//
//func (p *ConvertClassFileProcessor) convertBodyDeclaration(TypeMaker parser, AnnotationConverter converter, ClassFile classFile, BaseTypeParameter typeParameters, ClassFileBodyDeclaration outerClassFileBodyDeclaration) ClassFileBodyDeclaration {
//Map<String, TypeArgument> bindings;
//Map<String, BaseType> typeBounds;
//
//if (!classFile.isStatic() && (outerClassFileBodyDeclaration != null)) {
//bindings = outerClassFileBodyDeclaration.getBindings();
//typeBounds = outerClassFileBodyDeclaration.getTypeBounds();
//} else {
//bindings = Collections.emptyMap();
//typeBounds = Collections.emptyMap();
//}
//
//if (typeParameters != null) {
//populateBindingsWithTypeParameterVisitor.init(bindings=new HashMap<>(bindings), typeBounds=new HashMap<>(typeBounds));
//typeParameters.accept(populateBindingsWithTypeParameterVisitor);
//}
//
//ClassFileBodyDeclaration bodyDeclaration = new ClassFileBodyDeclaration(classFile, bindings, typeBounds, outerClassFileBodyDeclaration);
//
//bodyDeclaration.setFieldDeclarations(convertFields(parser, converter, classFile));
//bodyDeclaration.setMethodDeclarations(convertMethods(parser, converter, bodyDeclaration, classFile));
//bodyDeclaration.setInnerTypeDeclarations(convertInnerTypes(parser, converter, classFile, bodyDeclaration));
//
//return bodyDeclaration;
//}
//
//func (p *ConvertClassFileProcessor) convertFields(TypeMaker parser, AnnotationConverter converter, ClassFile classFile) []ClassFileFieldDeclaration {
//Field[] fields = classFile.getFields();
//
//if (fields == null) {
//return null;
//} else {
//DefaultList<ClassFileFieldDeclaration> list = new DefaultList<>(fields.length);
//
//for (Field field : fields) {
//BaseAnnotationReference annotationReferences = convertAnnotationReferences(converter, field);
//Type typeField = parser.parseFieldSignature(classFile, field);
//ExpressionVariableInitializer variableInitializer = convertFieldInitializer(field, typeField);
//FieldDeclarator fieldDeclarator = new FieldDeclarator(field.getName(), variableInitializer);
//
//list.add(new ClassFileFieldDeclaration(annotationReferences, field.getAccessFlags(), typeField, fieldDeclarator));
//}
//
//return list;
//}
//}
//
//func (p *ConvertClassFileProcessor) convertMethods(TypeMaker parser, AnnotationConverter converter, ClassFileBodyDeclaration bodyDeclaration, ClassFile classFile) []ClassFileConstructorOrMethodDeclaration {
//Method[] methods = classFile.getMethods();
//
//if (methods == null) {
//return null;
//} else {
//DefaultList<ClassFileConstructorOrMethodDeclaration> list = new DefaultList<>(methods.length);
//
//for (Method method : methods) {
//String name = method.getName();
//BaseAnnotationReference annotationReferences = convertAnnotationReferences(converter, method);
//AttributeAnnotationDefault annotationDefault = method.getAttribute("AnnotationDefault");
//ElementValue defaultAnnotationValue = null;
//
//if (annotationDefault != null) {
//defaultAnnotationValue = converter.convert(annotationDefault.getDefaultValue());
//}
//
//TypeMaker.MethodTypes methodTypes = parser.parseMethodSignature(classFile, method);
//Map<String, TypeArgument> bindings;
//Map<String, BaseType> typeBounds;
//
//if ((method.getAccessFlags() & ACC_STATIC) == 0) {
//bindings = bodyDeclaration.getBindings();
//typeBounds = bodyDeclaration.getTypeBounds();
//} else {
//bindings = Collections.emptyMap();
//typeBounds = Collections.emptyMap();
//}
//
//if (methodTypes.typeParameters != null) {
//populateBindingsWithTypeParameterVisitor.init(bindings=new HashMap<>(bindings), typeBounds=new HashMap<>(typeBounds));
//methodTypes.typeParameters.accept(populateBindingsWithTypeParameterVisitor);
//}
//
//AttributeCode code = method.getAttribute("Code");
//int firstLineNumber = 0;
//
//if (code != null) {
//AttributeLineNumberTable lineNumberTable = code.getAttribute("LineNumberTable");
//if (lineNumberTable != null) {
//firstLineNumber = lineNumberTable.getLineNumberTable()[0].getLineNumber();
//}
//}
//
//if ("<init>".equals(name)) {
//list.add(new ClassFileConstructorDeclaration(
//bodyDeclaration, classFile, method, annotationReferences, methodTypes.typeParameters,
//methodTypes.parameterTypes, methodTypes.exceptionTypes, bindings, typeBounds, firstLineNumber));
//} else if ("<clinit>".equals(name)) {
//list.add(new ClassFileStaticInitializerDeclaration(bodyDeclaration, classFile, method, bindings, typeBounds, firstLineNumber));
//} else {
//ClassFileMethodDeclaration methodDeclaration = new ClassFileMethodDeclaration(
//bodyDeclaration, classFile, method, annotationReferences, name, methodTypes.typeParameters,
//methodTypes.returnedType, methodTypes.parameterTypes, methodTypes.exceptionTypes, defaultAnnotationValue,
//bindings, typeBounds, firstLineNumber);
//if (classFile.isInterface()) {
//if (methodDeclaration.getFlags() == Constants.ACC_PUBLIC) {
//// For interfaces, add 'default' access flag on public methods
//methodDeclaration.setFlags(Declaration.FLAG_PUBLIC|Declaration.FLAG_DEFAULT);
//}
//}
//list.add(methodDeclaration);
//}
//}
//
//return list;
//}
//}
//
//func (p *ConvertClassFileProcessor) convertInnerTypes(TypeMaker parser, AnnotationConverter converter, ClassFile classFile, ClassFileBodyDeclaration outerClassFileBodyDeclaration) []ClassFileTypeDeclaration {
//List<ClassFile> innerClassFiles = classFile.getInnerClassFiles();
//
//if (innerClassFiles == null) {
//return null;
//} else {
//DefaultList<ClassFileTypeDeclaration> list = new DefaultList<>(innerClassFiles.size());
//
//for (ClassFile innerClassFile : innerClassFiles) {
//ClassFileTypeDeclaration innerTypeDeclaration;
//
//if (innerClassFile.isEnum()) {
//innerTypeDeclaration = convertEnumDeclaration(parser, converter, innerClassFile, outerClassFileBodyDeclaration);
//} else if (innerClassFile.isAnnotation()) {
//innerTypeDeclaration = convertAnnotationDeclaration(parser, converter, innerClassFile, outerClassFileBodyDeclaration);
//} else if (innerClassFile.isInterface()) {
//innerTypeDeclaration = convertInterfaceDeclaration(parser, converter, innerClassFile, outerClassFileBodyDeclaration);
//} else {
//innerTypeDeclaration = convertClassDeclaration(parser, converter, innerClassFile, outerClassFileBodyDeclaration);
//}
//
//list.add(innerTypeDeclaration);
//}
//
//return list;
//}
//}
//
//func (p *ConvertClassFileProcessor) convertAnnotationReferences(AnnotationConverter converter, ClassFile classFile) BaseAnnotationReference {
//Annotations visibles = classFile.getAttribute("RuntimeVisibleAnnotations");
//Annotations invisibles = classFile.getAttribute("RuntimeInvisibleAnnotations");
//
//return converter.convert(visibles, invisibles);
//}
//
//func (p *ConvertClassFileProcessor) convertAnnotationReferences(AnnotationConverter converter, Field field) BaseAnnotationReference {
//Annotations visibles = field.getAttribute("RuntimeVisibleAnnotations");
//Annotations invisibles = field.getAttribute("RuntimeInvisibleAnnotations");
//
//return converter.convert(visibles, invisibles);
//}
//
//func (p *ConvertClassFileProcessor) convertAnnotationReferences(AnnotationConverter converter, Method method) BaseAnnotationReference {
//Annotations visibles = method.getAttribute("RuntimeVisibleAnnotations");
//Annotations invisibles = method.getAttribute("RuntimeInvisibleAnnotations");
//
//return converter.convert(visibles, invisibles);
//}
//
//func (p *ConvertClassFileProcessor) convertFieldInitializer(Field field, Type typeField) ExpressionVariableInitializer {
//AttributeConstantValue acv = field.getAttribute("ConstantValue");
//
//if (acv == null) {
//return null;
//} else {
//ConstantValue constantValue = acv.getConstantValue();
//Expression expression;
//
//switch (constantValue.getTag()) {
//case Constant.CONSTANT_Integer:
//expression = new IntegerConstantExpression(typeField, ((ConstantInteger)constantValue).getValue());
//break;
//case Constant.CONSTANT_Float:
//expression = new FloatConstantExpression(((ConstantFloat)constantValue).getValue());
//break;
//case Constant.CONSTANT_Long:
//expression = new LongConstantExpression(((ConstantLong)constantValue).getValue());
//break;
//case Constant.CONSTANT_Double:
//expression = new DoubleConstantExpression(((ConstantDouble)constantValue).getValue());
//break;
//case Constant.CONSTANT_Utf8:
//expression = new StringConstantExpression(((ConstantUtf8)constantValue).getValue());
//break;
//default:
//throw new ConvertClassFileException("Invalid attributes");
//}
//
//return new ExpressionVariableInitializer(expression);
//}
//}
//
//func (p *ConvertClassFileProcessor) convertModuleDeclaration(ClassFile classFile) ModuleDeclaration {
//AttributeModule attributeModule = classFile.getAttribute("Module");
//List<ModuleDeclaration.ModuleInfo> requires = convertModuleDeclarationModuleInfo(attributeModule.getRequires());
//List<ModuleDeclaration.PackageInfo> exports = convertModuleDeclarationPackageInfo(attributeModule.getExports());
//List<ModuleDeclaration.PackageInfo> opens = convertModuleDeclarationPackageInfo(attributeModule.getOpens());
//DefaultList<String> uses = new DefaultList<>(attributeModule.getUses());
//List<ModuleDeclaration.ServiceInfo> provides = convertModuleDeclarationServiceInfo(attributeModule.getProvides());
//
//return new ModuleDeclaration(
//attributeModule.getFlags(), classFile.getInternalTypeName(), attributeModule.getName(),
//attributeModule.getVersion(), requires, exports, opens, uses, provides);
//}
//
//func (p *ConvertClassFileProcessor) convertModuleDeclarationModuleInfo(ModuleInfo[] moduleInfos) []ModuleDeclaration.ModuleInfo {
//if ((moduleInfos == null) || (moduleInfos.length == 0)) {
//return null;
//} else {
//DefaultList<ModuleDeclaration.ModuleInfo> list = new DefaultList<>(moduleInfos.length);
//
//for (ModuleInfo moduleInfo : moduleInfos) {
//list.add(new ModuleDeclaration.ModuleInfo(moduleInfo.getName(), moduleInfo.getFlags(), moduleInfo.getVersion()));
//}
//
//return list;
//}
//}
//
//func (p *ConvertClassFileProcessor) convertModuleDeclarationPackageInfo(PackageInfo[] packageInfos) []ModuleDeclaration.PackageInfo {
//if ((packageInfos == null) || (packageInfos.length == 0)) {
//return null;
//} else {
//DefaultList<ModuleDeclaration.PackageInfo> list = new DefaultList<>(packageInfos.length);
//
//for (PackageInfo packageInfo : packageInfos) {
//DefaultList<String> moduleInfoNames = (packageInfo.getModuleInfoNames() == null) ?
//null : new DefaultList<String>(packageInfo.getModuleInfoNames());
//list.add(new ModuleDeclaration.PackageInfo(packageInfo.getInternalName(), packageInfo.getFlags(), moduleInfoNames));
//}
//
//return list;
//}
//}
//
//func (p *ConvertClassFileProcessor) convertModuleDeclarationServiceInfo(ServiceInfo[] serviceInfos) []ModuleDeclaration.ServiceInfo {
//if ((serviceInfos == null) || (serviceInfos.length == 0)) {
//return null;
//} else {
//DefaultList<ModuleDeclaration.ServiceInfo> list = new DefaultList<>(serviceInfos.length);
//
//for (ServiceInfo serviceInfo : serviceInfos) {
//DefaultList<String> implementationTypeNames = (serviceInfo.getImplementationTypeNames() == null) ?
//null : new DefaultList<String>(serviceInfo.getImplementationTypeNames());
//list.add(new ModuleDeclaration.ServiceInfo(serviceInfo.getInterfaceTypeName(), implementationTypeNames));
//}
//
//return list;
//}
//}

type CustomPopulateBindingsWithTypeParameterVisitor struct {
	visitor.PopulateBindingsWithTypeParameterVisitor
}

func (v *CustomPopulateBindingsWithTypeParameterVisitor) VisitTypeParameter(parameter *_type.TypeParameter) {
	v.Bindings[parameter.Identifier()] = _type.NewGenericType(parameter.Identifier(), 0)
}

func (v *CustomPopulateBindingsWithTypeParameterVisitor) VisitTypeParameterWithTypeBounds(parameter *_type.TypeParameterWithTypeBounds) {
	v.Bindings[parameter.Identifier()] = _type.NewGenericType(parameter.Identifier(), 0)
	v.TypeBounds[parameter.Identifier()] = parameter.TypeBounds()
}
