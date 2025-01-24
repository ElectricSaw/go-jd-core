package model

import (
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
	"reflect"
)

const (
	EvUnknown = iota
	EvPrimitiveType
	EvEnumConstValue
	EvClassInfo
	EvAnnotationValue
	EvArrayValue
)

func NewAnnotationElementValue(reference *AnnotationReference) AnnotationElementValue {
	v := AnnotationElementValue{
		DefaultBase:       *util.NewDefaultBase[*AnnotationElementValue]().(*util.DefaultBase[*AnnotationElementValue]),
		Type:              reference.Type,
		ElementValue:      reference.ElementValue,
		ElementValuePairs: reference.ElementValuePairs,
	}
	v.SetValue(&v)
	return v
}

func NewAnnotationReference(type_ *ObjectType) AnnotationReference {
	return NewAnnotationReferenceWithEv(type_, nil)
}

func NewAnnotationReferenceWithEv(type_ *ObjectType, elementValue IElementValue) AnnotationReference {
	return AnnotationReference{
		Type:         type_,
		ElementValue: elementValue,
	}
}

func NewAnnotationReferenceWithPair(type_ *ObjectType, pairs *ElementValuePair) AnnotationReference {
	return NewAnnotationReferenceWithAll(type_, nil, pairs)
}

func NewAnnotationReferenceWithAll(type_ *ObjectType, elementValue IElementValue,
	elementValuePairs *ElementValuePair) AnnotationReference {
	return AnnotationReference{
		Type:              type_,
		ElementValue:      elementValue,
		ElementValuePairs: elementValuePairs,
	}
}

func NewAnnotationReferences() AnnotationReferences {
	return AnnotationReferences{
		DefaultList: *util.NewDefaultListWithCapacity[IAnnotationReference](0).(*util.DefaultList[IAnnotationReference]),
	}
}

func NewElementValuePair(name string, elementValue IElementValue) ElementValuePair {
	p := ElementValuePair{
		DefaultBase:  *util.NewDefaultBase[IElementValuePair]().(*util.DefaultBase[IElementValuePair]),
		name:         name,
		elementValue: elementValue,
	}
	p.SetValue(&p)
	return p
}

func NewElementValuePairs() ElementValuePairs {
	return ElementValuePairs{
		DefaultList: *util.NewDefaultListWithCapacity[*ElementValuePair](0).(*util.DefaultList[*ElementValuePair]),
	}
}

func NewElementValues() ElementValues {
	return ElementValues{
		DefaultList: *util.NewDefaultListWithCapacity[IElementValue](0).(*util.DefaultList[IElementValue]),
	}
}

func NewElementValuePairsWithCapacity(capacity int) ElementValuePairs {
	return ElementValuePairs{
		DefaultList: *util.NewDefaultListWithCapacity[*ElementValuePair](capacity).(*util.DefaultList[*ElementValuePair]),
	}
}

func NewElementValuePairsWithElements(pairs ...*ElementValuePair) ElementValuePairs {
	return ElementValuePairs{
		DefaultList: *util.NewDefaultListWithElements[*ElementValuePair](pairs...).(*util.DefaultList[*ElementValuePair]),
	}
}

func NewElementValuesWithCapacity(capacity int) ElementValues {
	return ElementValues{
		DefaultList: *util.NewDefaultListWithCapacity[IElementValue](0).(*util.DefaultList[IElementValue]),
	}
}

func NewElementValuesWithElements(pairs ...IElementValue) ElementValues {
	return ElementValues{
		DefaultList: *util.NewDefaultListWithCapacity[IElementValue](0).(*util.DefaultList[IElementValue]),
	}
}

func NewElementValueArrayInitializerElementValue(elementValueArrayInitializer IElementValue) ElementValueArrayInitializerElementValue {
	v := ElementValueArrayInitializerElementValue{
		DefaultBase:                  *util.NewDefaultBase[*ElementValueArrayInitializerElementValue]().(*util.DefaultBase[*ElementValueArrayInitializerElementValue]),
		ElementValueArrayInitializer: elementValueArrayInitializer,
	}
	v.SetValue(&v)
	return v
}

func NewElementValueArrayInitializerElementValueEmpty() ElementValueArrayInitializerElementValue {
	return NewElementValueArrayInitializerElementValue(nil)
}

func NewExpressionElementValue(expression IExpression) ExpressionElementValue {
	v := ExpressionElementValue{
		DefaultBase: *util.NewDefaultBase[*ExpressionElementValue]().(*util.DefaultBase[*ExpressionElementValue]),
		Expression:  expression,
	}
	v.SetValue(&v)
	return v
}

func NewInnerObjectReference(internalName, qualifiedName, name string,
	outerType *ObjectType) InnerObjectReference {
	return NewInnerObjectReferenceWithAll(internalName, qualifiedName, name, nil, 0, outerType)
}

func NewInnerObjectReferenceWithDim(internalName, qualifiedName, name string,
	dimension int, outerType *ObjectType) InnerObjectReference {
	return NewInnerObjectReferenceWithAll(internalName, qualifiedName, name, nil, dimension, outerType)
}

func NewInnerObjectReferenceWithArgs(internalName, qualifiedName, name string,
	typeArguments ITypeArgument, outerType *ObjectType) InnerObjectReference {
	return NewInnerObjectReferenceWithAll(internalName, qualifiedName, name, typeArguments, 0, outerType)
}

func NewObjectReference(internalName, qualifiedName, name string) ObjectReference {
	return NewObjectReferenceWithAll(internalName, qualifiedName, name, nil, 0)
}

func NewObjectReferenceWithDim(internalName, qualifiedName, name string, dimension int) ObjectReference {
	return NewObjectReferenceWithAll(internalName, qualifiedName, name, nil, dimension)
}

func NewObjectReferenceWithArgs(internalName, qualifiedName, name string, typeArguments ITypeArgument) ObjectReference {
	return NewObjectReferenceWithAll(internalName, qualifiedName, name, typeArguments, 0)
}

func NewObjectReferenceWithAll(internalName, qualifiedName, name string, typeArguments ITypeArgument, dimension int) ObjectReference {
	r := ObjectReference{
		DefaultBase:   *util.NewDefaultBase[IType]().(*util.DefaultBase[IType]),
		internalName:  internalName,
		QualifiedName: qualifiedName,
		Name:          name,
		TypeArguments: typeArguments,
		Dimension:     dimension,
		Descriptor:    createDescriptor(fmt.Sprintf("L%s;", internalName), dimension),
	}
	r.SetValue(&r)
	return r
}

type IReference interface {
	Accept(visitor IReferenceVisitor)
	Equals(o interface{}) bool
	HashCode() int
	String() string
}

type IReferenceVisitor interface {
	VisitAnnotationElementValue(reference *AnnotationElementValue)
	VisitAnnotationReference(reference *AnnotationReference)
	VisitAnnotationReferences(references *AnnotationReferences)
	VisitElementValueArrayInitializerElementValue(reference *ElementValueArrayInitializerElementValue)
	VisitElementValues(references *ElementValues)
	VisitElementValuePair(reference *ElementValuePair)
	VisitElementValuePairs(references *ElementValuePairs)
	VisitExpressionElementValue(reference *ExpressionElementValue)
	VisitInnerObjectReference(reference *InnerObjectReference)
	VisitObjectReference(reference *ObjectReference)
}

type IAnnotationReference interface {
	Accept(visitor IReferenceVisitor)
	Equals(o interface{}) bool
	HashCode() int
	String() string

	IsAnnotationReference() bool
}

type IElementValue interface {
	Accept(visitor IReferenceVisitor)
	Equals(o interface{}) bool
	HashCode() int
	String() string

	IsElementValue() bool
}

type IElementValuePair interface {
	Accept(visitor IReferenceVisitor)
	Equals(o interface{}) bool
	HashCode() int
	String() string

	IsElementValuePair() bool
}

type AnnotationElementValue struct {
	util.DefaultBase[*AnnotationElementValue]
	Type              *ObjectType
	ElementValue      IElementValue
	ElementValuePairs *ElementValuePair
}

func (e *AnnotationElementValue) Accept(visitor IReferenceVisitor) {
	visitor.VisitAnnotationElementValue(e)
}

func (e *AnnotationElementValue) Equals(o interface{}) bool {
	if o == e {
		return true
	}

	if o == nil {
		return false
	}

	var other *AnnotationReference

	switch o := o.(type) {
	case AnnotationReference:
		other = &o
	case *AnnotationReference:
		other = o
	default:
		return false
	}

	if e.ElementValue != nil {
		if !e.ElementValue.Equals(other.ElementValue) {
			return false
		}
		if other.ElementValue != nil {
			return false
		}
	}

	if e.ElementValuePairs != nil {
		if !e.ElementValuePairs.Equals(other.ElementValuePairs) {
			return false
		}
		if other.ElementValuePairs != nil {
			return false
		}
	}

	if !e.Type.Equals(other.Type) {
		return false
	}

	return true
}

func (e *AnnotationElementValue) HashCode() int {
	result := 970748295 + e.Type.HashCode()
	result = 31 * result
	if e.ElementValue != nil {
		result += e.ElementValue.HashCode()
	}
	result = 31 * result
	if e.ElementValuePairs != nil {
		result += e.ElementValuePairs.HashCode()
	}
	return result
}

func (e *AnnotationElementValue) String() string {
	return fmt.Sprintf("AnnotationElementValue{type=%v, elementValue=%s, elementValuePairs=%v", e.Type, e.ElementValue, e.ElementValuePairs)
}

func (e *AnnotationElementValue) IsElementValue() bool {
	return true
}

func (e *AnnotationElementValue) IsAnnotationReference() bool {
	return true
}

type AnnotationReference struct {
	Type              *ObjectType
	ElementValue      IElementValue
	ElementValuePairs *ElementValuePair
}

func (r *AnnotationReference) Accept(visitor IReferenceVisitor) {
	visitor.VisitAnnotationReference(r)
}

func (r *AnnotationReference) Equals(o interface{}) bool {
	if o == r {
		return true
	}

	if o == nil {
		return false
	}

	var other *AnnotationReference

	switch o := o.(type) {
	case AnnotationReference:
		other = &o
	case *AnnotationReference:
		other = o
	default:
		return false
	}

	if r.ElementValue != nil {
		if !r.ElementValue.Equals(other.ElementValue) {
			return false
		}
		if other.ElementValue != nil {
			return false
		}
	}

	if r.ElementValuePairs != nil {
		if !r.ElementValuePairs.Equals(other.ElementValuePairs) {
			return false
		}
		if other.ElementValuePairs != nil {
			return false
		}
	}

	if !r.Type.Equals(other.Type) {
		return false
	}

	return true
}

func (r *AnnotationReference) HashCode() int {
	result := 970748295 + r.Type.HashCode()
	result = 31 * result
	if r.ElementValue != nil {
		result += r.ElementValue.HashCode()
	}
	result = 31 * result
	if r.ElementValuePairs != nil {
		result += r.ElementValuePairs.HashCode()
	}
	return result
}

func (r *AnnotationReference) String() string {
	return fmt.Sprintf("AnnotationReference{}")
}

func (r *AnnotationReference) IsAnnotationReference() bool {
	return true
}

type AnnotationReferences struct {
	util.DefaultList[IAnnotationReference]
}

func (r *AnnotationReferences) Accept(visitor IReferenceVisitor) {
	visitor.VisitAnnotationReferences(r)
}

func (r *AnnotationReferences) Equals(o interface{}) bool {
	//if o == e {
	//	return true
	//}
	//
	//if o == nil {
	//	return false
	//}
	//
	//var other *AnnotationReference
	//
	//switch o := o.(type) {
	//case AnnotationReference:
	//	other = &o
	//case *AnnotationReference:
	//	other = o
	//default:
	//	return false
	//}
	//
	//return false
	return r == o
}

func (r *AnnotationReferences) HashCode() int {
	return hashCodeWithStruct(r)
}

func (r *AnnotationReferences) String() string {
	return fmt.Sprintf("ElementValues{%v}", r.DefaultList)
}

func (r *AnnotationReferences) IsAnnotationReference() bool {
	return true
}

type ElementValuePair struct {
	util.DefaultBase[IElementValuePair]

	name         string
	elementValue IElementValue
}

func (e *ElementValuePair) Accept(visitor IReferenceVisitor) {
	visitor.VisitElementValuePair(e)
}

func (e *ElementValuePair) Equals(o interface{}) bool {
	//if o == e {
	//	return true
	//}
	//
	//if o == nil {
	//	return false
	//}
	//
	//var other *AnnotationReference
	//
	//switch o := o.(type) {
	//case AnnotationReference:
	//	other = &o
	//case *AnnotationReference:
	//	other = o
	//default:
	//	return false
	//}
	//
	//return false
	return e == o
}

func (e *ElementValuePair) HashCode() int {
	return hashCodeWithStruct(e)
}

func (e *ElementValuePair) String() string {
	return fmt.Sprintf("ElementValuePair{name=%s, elementValue=%s}", e.name, e.elementValue)
}

func (e *ElementValuePair) IsElementValuePair() bool {
	return true
}

type ElementValuePairs struct {
	util.DefaultList[*ElementValuePair]
}

func (e *ElementValuePairs) Accept(visitor IReferenceVisitor) {
	visitor.VisitElementValuePairs(e)
}

func (e *ElementValuePairs) Equals(o interface{}) bool {
	//if o == e {
	//	return true
	//}
	//
	//if o == nil {
	//	return false
	//}
	//
	//var other *AnnotationReference
	//
	//switch o := o.(type) {
	//case AnnotationReference:
	//	other = &o
	//case *AnnotationReference:
	//	other = o
	//default:
	//	return false
	//}
	//
	//return false
	return e == o
}

func (e *ElementValuePairs) HashCode() int {
	return hashCodeWithStruct(e)
}

func (e *ElementValuePairs) String() string {
	return fmt.Sprintf("ElementValuePairs{%v}", *e)
}

func (e *ElementValuePairs) IsElementValuePair() bool {
	return true
}

type ElementValues struct {
	util.DefaultList[IElementValue]
}

func (e *ElementValues) Accept(visitor IReferenceVisitor) {
	visitor.VisitElementValues(e)
}

func (e *ElementValues) Equals(o interface{}) bool {
	//if o == e {
	//	return true
	//}
	//
	//if o == nil {
	//	return false
	//}
	//
	//var other *AnnotationReference
	//
	//switch o := o.(type) {
	//case AnnotationReference:
	//	other = &o
	//case *AnnotationReference:
	//	other = o
	//default:
	//	return false
	//}
	//
	//return false
	return e == o
}

func (e *ElementValues) HashCode() int {
	return hashCodeWithStruct(e)
}

func (e *ElementValues) String() string {
	return fmt.Sprintf("ElementValues{%v}", e.DefaultList)
}

func (e *ElementValues) IsElementValue() bool {
	return true
}

type ElementValueArrayInitializerElementValue struct {
	util.DefaultBase[*ElementValueArrayInitializerElementValue]

	ElementValueArrayInitializer IElementValue
}

func (e *ElementValueArrayInitializerElementValue) Accept(visitor IReferenceVisitor) {
	visitor.VisitElementValueArrayInitializerElementValue(e)
}

func (e *ElementValueArrayInitializerElementValue) Equals(o interface{}) bool {
	//if o == e {
	//	return true
	//}
	//
	//if o == nil {
	//	return false
	//}
	//
	//var other *AnnotationReference
	//
	//switch o := o.(type) {
	//case AnnotationReference:
	//	other = &o
	//case *AnnotationReference:
	//	other = o
	//default:
	//	return false
	//}
	//
	//return false
	return e == o
}

func (e *ElementValueArrayInitializerElementValue) HashCode() int {
	return hashCodeWithStruct(e)
}

func (e *ElementValueArrayInitializerElementValue) String() string {
	return fmt.Sprintf("ElementValueArrayInitializerElementValue{%s}", e.ElementValueArrayInitializer)
}

func (e *ElementValueArrayInitializerElementValue) IsElementValue() bool {
	return true
}

type ExpressionElementValue struct {
	util.DefaultBase[*ExpressionElementValue]

	Expression IExpression
}

func (e *ExpressionElementValue) Accept(visitor IReferenceVisitor) {
	visitor.VisitExpressionElementValue(e)
}

func (e *ExpressionElementValue) Equals(o interface{}) bool {
	//if o == e {
	//	return true
	//}
	//
	//if o == nil {
	//	return false
	//}
	//
	//var other *AnnotationReference
	//
	//switch o := o.(type) {
	//case AnnotationReference:
	//	other = &o
	//case *AnnotationReference:
	//	other = o
	//default:
	//	return false
	//}
	//
	//return false
	return e == o
}

func (e *ExpressionElementValue) HashCode() int {
	return hashCodeWithStruct(e)
}

func (e *ExpressionElementValue) String() string {
	return fmt.Sprintf("ExpressionElementValue{%s}", *e)
}

func (e *ExpressionElementValue) IsElementValue() bool {
	return true
}

func NewInnerObjectReferenceWithAll(internalName, qualifiedName, name string,
	typeArguments ITypeArgument, dimension int, outerType *ObjectType) InnerObjectReference {
	t := InnerObjectReference{
		DefaultBase:   *util.NewDefaultBase[IType]().(*util.DefaultBase[IType]),
		internalName:  internalName,
		QualifiedName: qualifiedName,
		Name:          name,
		TypeArguments: typeArguments,
		Dimension:     dimension,
		Descriptor:    createDescriptor(fmt.Sprintf("L%s;", internalName), dimension),
		outerType:     outerType,
	}
	t.SetValue(&t)
	return t
}

type InnerObjectReference struct {
	util.DefaultBase[IType]

	internalName  string
	QualifiedName string
	Name          string
	TypeArguments ITypeArgument
	Dimension     int
	Descriptor    string
	outerType     *ObjectType
}

func (r *InnerObjectReference) CreateType(dimension int) IType {
	tmp := NewInnerObjectTypeWithAll(r.internalName, r.QualifiedName, r.Name, r.TypeArguments, dimension, r.OuterType())
	return &tmp
}

func (r *InnerObjectReference) IsGenericType() bool {
	return false
}

func (r *InnerObjectReference) IsInnerObjectType() bool {
	return true
}

func (r *InnerObjectReference) IsObjectType() bool {
	return true
}

func (r *InnerObjectReference) IsPrimitiveType() bool {
	return false
}

func (r *InnerObjectReference) IsTypes() bool {
	return false
}

func (r *InnerObjectReference) OuterType() *ObjectType {
	return r.outerType
}

func (r *InnerObjectReference) InternalName() string {
	return r.internalName
}

func (r *InnerObjectReference) HashCode() int {
	result := 735485092 + hashCodeWithString(r.internalName)
	result *= 31
	if r.TypeArguments != nil {
		result += r.TypeArguments.HashCode()
	}
	result = 31*result + r.Dimension
	result = 111476860 + result
	result = 31*result + r.OuterType().HashCode()
	return result
}

func (r *InnerObjectReference) AcceptTypeVisitor(visitor ITypeVisitor) {
	visitor.VisitInnerObjectType((*InnerObjectType)(r))
}

/////////////////////////////////////////////////////////////////////

func (r *InnerObjectReference) TypeArgumentFirst() ITypeArgument {
	return r
}

func (r *InnerObjectReference) TypeArgumentList() util.IList[ITypeArgument] {
	return util.NewDefaultListWithElements[ITypeArgument](r)
}

func (r *InnerObjectReference) TypeArgumentSize() int {
	return 1
}

func (r *InnerObjectReference) Type() IType {
	return &OtTypeUndefinedObject
}

func (r *InnerObjectReference) IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool {
	switch meta := typeArgument.(type) {
	case *ObjectType:
		if r.Dimension != meta.Dimension || r.internalName != meta.InternalName() {
			return false
		}

		if meta.TypeArguments == nil {
			return r.TypeArguments == nil
		} else if r.TypeArguments == nil {
			return false
		} else {
			return r.TypeArguments.IsTypeArgumentAssignableFrom(typeBounds, meta.TypeArguments)
		}
	case *InnerObjectType:
		if r.Dimension != meta.Dimension || r.internalName != meta.InternalName() {
			return false
		}

		if meta.TypeArguments == nil {
			return r.TypeArguments == nil
		} else if r.TypeArguments == nil {
			return false
		} else {
			return r.TypeArguments.IsTypeArgumentAssignableFrom(typeBounds, meta.TypeArguments)
		}
	case *GenericType:
		bt := typeBounds[meta.Name]
		ot, ok := bt.(*ObjectType)

		if ok {
			if r.internalName == ot.InternalName() {
				return true
			}
		}
	}

	return false
}

func (r *InnerObjectReference) IsTypeArgumentList() bool {
	return false
}

func (r *InnerObjectReference) IsGenericTypeArgument() bool {
	return false
}

func (r *InnerObjectReference) IsInnerObjectTypeArgument() bool {
	return true
}

func (r *InnerObjectReference) IsObjectTypeArgument() bool {
	return true
}

func (r *InnerObjectReference) IsPrimitiveTypeArgument() bool {
	return false
}

func (r *InnerObjectReference) IsWildcardExtendsTypeArgument() bool {
	return false
}

func (r *InnerObjectReference) IsWildcardSuperTypeArgument() bool {
	return false
}

func (r *InnerObjectReference) IsWildcardTypeArgument() bool {
	return false
}

func (r *InnerObjectReference) AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor) {
	visitor.VisitInnerObjectType((*InnerObjectType)(r))
}

/////////////////////////////////////////////////////////////////////

func (r *InnerObjectReference) Equals(o interface{}) bool {
	if r == o {
		return true
	}

	if o == nil {
		return false
	}

	var other *InnerObjectReference

	switch o := o.(type) {
	case InnerObjectReference:
		other = &o
	case *InnerObjectReference:
		other = o
	default:
		return false
	}

	if !r.outerType.Equals(other.Dimension) {
		return false
	}

	return true
}

func (r *InnerObjectReference) String() string {
	if r.TypeArguments == nil {
		return fmt.Sprintf("InnerObjectType { %s.%s }", r.outerType, r.Descriptor)
	} else {
		return fmt.Sprintf("InnerObjectType { %s.%s<%s> }", r.outerType, r.Descriptor, r.TypeArguments)
	}
}

/////////////////////////////////////////////////////////////////////

func (r *InnerObjectReference) CreateTypeWithArg(typeArguments ITypeArgument) IType {
	tmp := NewInnerObjectTypeWithAll(r.internalName, r.QualifiedName, r.Name, typeArguments, r.Dimension, r.outerType)
	return &tmp
}

func (r *InnerObjectReference) Accept(visitor IReferenceVisitor) {
	visitor.VisitInnerObjectReference(r)
}

type ObjectReference struct {
	util.DefaultBase[IType]

	internalName  string
	QualifiedName string
	Name          string
	TypeArguments ITypeArgument
	Dimension     int
	Descriptor    string
}

func (r *ObjectReference) CreateType(dimension int) IType {
	if r.Dimension == dimension {
		return r
	} else if r.Descriptor[len(r.Descriptor)-1] != ';' {
		if dimension == 0 {
			tmp := GetPrimitiveType(int(r.Descriptor[r.Dimension]))
			return &tmp
		} else {
			tmp := NewObjectTypeWithDescAndDim(r.internalName, r.Dimension)
			return &tmp
		}
	} else {
		tmp := NewObjectTypeWithAll(r.internalName, r.QualifiedName, r.Name, r.TypeArguments, dimension)
		return &tmp
	}
}

func (r *ObjectReference) IsGenericType() bool {
	return false
}

func (r *ObjectReference) IsInnerObjectType() bool {
	return false
}

func (r *ObjectReference) IsObjectType() bool {
	return true
}

func (r *ObjectReference) IsPrimitiveType() bool {
	return false
}

func (r *ObjectReference) IsTypes() bool {
	return false
}

func (r *ObjectReference) OuterType() *ObjectType {
	return &OtTypeUndefinedObject
}

func (r *ObjectReference) InternalName() string {
	return r.internalName
}

func (r *ObjectReference) HashCode() int {
	result := 735485092 + hashCodeWithString(r.internalName)
	result *= 31
	if r.TypeArguments != nil {
		result += r.TypeArguments.HashCode()
	}
	result = 31*result + r.Dimension
	return result
}

func (r *ObjectReference) AcceptTypeVisitor(visitor ITypeVisitor) {
	visitor.VisitObjectType((*ObjectType)(r))
}

/////////////////////////////////////////////////////////////////////

func (r *ObjectReference) TypeArgumentFirst() ITypeArgument {
	return r
}

func (r *ObjectReference) TypeArgumentList() util.IList[ITypeArgument] {
	return util.NewDefaultListWithElements[ITypeArgument](r)
}

func (r *ObjectReference) TypeArgumentSize() int {
	return 1
}

func (r *ObjectReference) Type() IType {
	return &OtTypeUndefinedObject
}

func (r *ObjectReference) IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool {
	switch meta := typeArgument.(type) {
	case *ObjectType:
		if r.Dimension != meta.Dimension || r.internalName != meta.InternalName() {
			return false
		}

		if meta.TypeArguments == nil {
			return r.TypeArguments == nil
		} else if r.TypeArguments == nil {
			return false
		} else {
			return r.TypeArguments.IsTypeArgumentAssignableFrom(typeBounds, meta.TypeArguments)
		}
	case *InnerObjectType:
		if r.Dimension != meta.Dimension || r.internalName != meta.InternalName() {
			return false
		}

		if meta.TypeArguments == nil {
			return r.TypeArguments == nil
		} else if r.TypeArguments == nil {
			return false
		} else {
			return r.TypeArguments.IsTypeArgumentAssignableFrom(typeBounds, meta.TypeArguments)
		}
	case *GenericType:
		bt := typeBounds[meta.Name]
		ot, ok := bt.(*ObjectType)

		if ok {
			if r.internalName == ot.InternalName() {
				return true
			}
		}
	}

	return false
}

func (r *ObjectReference) IsTypeArgumentList() bool {
	return false
}

func (r *ObjectReference) IsGenericTypeArgument() bool {
	return false
}

func (r *ObjectReference) IsInnerObjectTypeArgument() bool {
	return false
}

func (r *ObjectReference) IsObjectTypeArgument() bool {
	return true
}

func (r *ObjectReference) IsPrimitiveTypeArgument() bool {
	return false
}

func (r *ObjectReference) IsWildcardExtendsTypeArgument() bool {
	return false
}

func (r *ObjectReference) IsWildcardSuperTypeArgument() bool {
	return false
}

func (r *ObjectReference) IsWildcardTypeArgument() bool {
	return false
}

func (r *ObjectReference) AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor) {
	visitor.VisitObjectType((*ObjectType)(r))
}

/////////////////////////////////////////////////////////////////////

func (r *ObjectReference) Equals(o interface{}) bool {
	if r == o {
		return true
	}

	if o == nil {
		return false
	}

	var other *ObjectReference

	switch o := o.(type) {
	case ObjectReference:
		other = &o
	case *ObjectReference:
		other = o
	default:
		return false
	}

	if r.Dimension != other.Dimension {
		return false
	}

	if r.internalName != other.internalName {
		return false
	}

	if r.internalName == "jara/lang/Class" {
		wildcard1 := (r.TypeArguments == nil) || (reflect.TypeOf(r.TypeArguments) == reflect.TypeOf(WildcardTypeArgument{}))
		wildcard2 := (other.TypeArguments == nil) || (reflect.TypeOf(other.TypeArguments) == reflect.TypeOf(WildcardTypeArgument{}))

		if wildcard1 || wildcard2 {
			return true
		}
	}

	if r.TypeArguments != nil {
		return r.TypeArguments.Equals(other.TypeArguments)
	}

	return r.TypeArguments == nil
}

func (r *ObjectReference) String() string {
	msg := fmt.Sprintf("ObjectType{ %s", r.internalName)
	if r.TypeArguments != nil {
		msg += fmt.Sprintf("<%s>", r.TypeArguments)
	}
	if r.Dimension > 0 {
		msg += fmt.Sprintf(", %d", r.Dimension)
	}
	msg += " }"
	return msg
}

/////////////////////////////////////////////////////////////////////

func (r *ObjectReference) IsTypeArgumentAssignableFromWithObj(typeBounds map[string]IType, objectType ObjectType) bool {
	if r.Dimension != objectType.Dimension || r.internalName != objectType.InternalName() {
		return false
	}

	if objectType.TypeArguments == nil {
		return r.TypeArguments == nil
	} else if r.TypeArguments == nil {
		return false
	} else {
		return r.TypeArguments.IsTypeArgumentAssignableFrom(typeBounds, objectType.TypeArguments)
	}
}

func (r *ObjectReference) CreateTypeWithArgs(typeArguments ITypeArgument) IType {
	if r.TypeArguments == typeArguments {
		return r
	} else {
		tmp := NewObjectTypeWithAll(r.internalName, r.QualifiedName, r.Name, typeArguments, r.Dimension)
		return &tmp
	}
}

func (r *ObjectReference) Accept(visitor IReferenceVisitor) {
	visitor.VisitObjectReference(r)
}
