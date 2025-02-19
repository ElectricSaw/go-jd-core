package model

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
	"hash/fnv"
)

/////////////////////////////////////////////////////////////////////////
//  New Functions
/////////////////////////////////////////////////////////////////////////

func NewTypeArguments() TypeArguments {
	return NewTypeArgumentsWithCapacity(0)
}

func NewTypeArgumentsWithCapacity(capacity int) TypeArguments {
	return TypeArguments{
		DefaultList: *util.NewDefaultListWithCapacity[ITypeArgument](capacity).(*util.DefaultList[ITypeArgument]),
	}
}

func NewDiamondTypeArgument() DiamondTypeArgument {
	return DiamondTypeArgument{}
}

func NewWildcardExtendsTypeArgument(typ IType) WildcardExtendsTypeArgument {
	return WildcardExtendsTypeArgument{
		Type: typ,
	}
}

func NewWildcardSuperTypeArgument(typ IType) WildcardSuperTypeArgument {
	return WildcardSuperTypeArgument{
		Type: typ,
	}
}

func NewWildcardTypeArgument() WildcardTypeArgument {
	return WildcardTypeArgument{}
}

/////////////////////////////////////////////////////////////////////////
//  Interfaces
/////////////////////////////////////////////////////////////////////////

type ITypeArgument interface {
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)

	TypeArgumentFirst() ITypeArgument            // ITypeArgument
	TypeArgumentList() util.IList[ITypeArgument] // ITypeArgument
	TypeArgumentSize() int
	IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool
	IsTypeArgumentList() bool
	IsGenericTypeArgument() bool
	IsInnerObjectTypeArgument() bool
	IsObjectTypeArgument() bool
	IsPrimitiveTypeArgument() bool
	IsWildcardExtendsTypeArgument() bool
	IsWildcardSuperTypeArgument() bool
	IsWildcardTypeArgument() bool

	GetType() IType

	HashCode() int
	Equals(o interface{}) bool
	String() string
}

type ITypeArgumentVisitable interface {
	AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor)
}

type ITypeArgumentVisitor interface {
	VisitTypeArguments(arguments *TypeArguments)
	VisitDiamondTypeArgument(argument *DiamondTypeArgument)
	VisitWildcardExtendsTypeArgument(argument *WildcardExtendsTypeArgument)
	VisitWildcardSuperTypeArgument(argument *WildcardSuperTypeArgument)
	VisitWildcardTypeArgument(argument *WildcardTypeArgument)
	VisitPrimitiveType(t *PrimitiveType)
	VisitObjectType(t IType)
	VisitInnerObjectType(t IType)
	VisitGenericType(t *GenericType)
}

/////////////////////////////////////////////////////////////////////////
//  Structures
/////////////////////////////////////////////////////////////////////////

type DiamondTypeArgument struct {
}

func (a *DiamondTypeArgument) TypeArgumentFirst() ITypeArgument {
	return a
}

func (a *DiamondTypeArgument) TypeArgumentList() util.IList[ITypeArgument] {
	return util.NewDefaultListWithElements[ITypeArgument](a)
}

func (a *DiamondTypeArgument) TypeArgumentSize() int {
	return 1
}

func (a *DiamondTypeArgument) GetType() IType {
	return &OtTypeUndefinedObject
}

func (a *DiamondTypeArgument) IsTypeArgumentAssignableFrom(_ map[string]IType, _ ITypeArgument) bool {
	return true
}

func (a *DiamondTypeArgument) IsTypeArgumentList() bool {
	return false
}

func (a *DiamondTypeArgument) IsGenericTypeArgument() bool {
	return false
}

func (a *DiamondTypeArgument) IsInnerObjectTypeArgument() bool {
	return false
}

func (a *DiamondTypeArgument) IsObjectTypeArgument() bool {
	return false
}

func (a *DiamondTypeArgument) IsPrimitiveTypeArgument() bool {
	return false
}

func (a *DiamondTypeArgument) IsWildcardExtendsTypeArgument() bool {
	return false
}

func (a *DiamondTypeArgument) IsWildcardSuperTypeArgument() bool {
	return false
}

func (a *DiamondTypeArgument) IsWildcardTypeArgument() bool {
	return false
}

func (a *DiamondTypeArgument) AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor) {
	visitor.VisitDiamondTypeArgument(a)
}

func (a *DiamondTypeArgument) Equals(o interface{}) bool {
	if o == nil {
		return false
	}

	var other *DiamondTypeArgument

	switch o := o.(type) {
	case DiamondTypeArgument:
		other = &o
	case *DiamondTypeArgument:
		other = o
	default:
		return false
	}

	if other == a {
		return true
	}

	return false
}

func (a *DiamondTypeArgument) HashCode() int {
	return hashCodeWithStruct(a)
}

func (a *DiamondTypeArgument) String() string {
	return "DiamondTypeArgument {}"
}

type TypeArguments struct {
	util.DefaultList[ITypeArgument]
}

func (t *TypeArguments) TypeArgumentFirst() ITypeArgument {
	return t.Get(0)
}

func (t *TypeArguments) TypeArgumentList() util.IList[ITypeArgument] {
	return &t.DefaultList
}

func (t *TypeArguments) TypeArgumentSize() int {
	return t.DefaultList.Size()
}

func (t *TypeArguments) GetType() IType {
	return &OtTypeUndefinedObject
}

func (t *TypeArguments) IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool {
	ata, ok := typeArgument.(*TypeArguments)
	if !ok {
		return false
	}

	if t.Size() != ata.Size() {
		return false
	}

	iterator1 := t.Iterator()
	iterator2 := ata.Iterator()

	for iterator1.HasNext() {
		if !iterator1.Next().IsTypeArgumentAssignableFrom(typeBounds, iterator2.Next()) {
			return false
		}
	}

	return true
}

func (t *TypeArguments) IsTypeArgumentList() bool {
	return true
}

func (t *TypeArguments) IsGenericTypeArgument() bool {
	return false
}

func (t *TypeArguments) IsInnerObjectTypeArgument() bool {
	return false
}

func (t *TypeArguments) IsObjectTypeArgument() bool {
	return false
}

func (t *TypeArguments) IsPrimitiveTypeArgument() bool {
	return false
}

func (t *TypeArguments) IsWildcardExtendsTypeArgument() bool {
	return false
}

func (t *TypeArguments) IsWildcardSuperTypeArgument() bool {
	return false
}

func (t *TypeArguments) IsWildcardTypeArgument() bool {
	return false
}

func (t *TypeArguments) AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor) {
	visitor.VisitTypeArguments(t)
}

func (t *TypeArguments) Equals(o interface{}) bool {
	if o == t {
		return true
	}

	if o == nil {
		return false
	}

	var other *TypeArguments

	switch o := o.(type) {
	case TypeArguments:
		other = &o
	case *TypeArguments:
		other = o
	default:
		return false
	}

	size := t.Size()
	equals := size == other.Size()

	if equals {
		es := t.ToSlice()
		otherEs := other.ToSlice()
		for i := 0; i < size; i++ {
			if es[i] != otherEs[i] {
				return false
			}
		}
	}

	return equals
}

func (t *TypeArguments) HashCode() int {
	result := 1
	for _, item := range t.ToSlice() {
		result = 31*result + item.HashCode()
	}
	return result
}

func (t *TypeArguments) String() string {
	return "TypeArguments {}"
}

type WildcardExtendsTypeArgument struct {
	Type IType
}

func (t *WildcardExtendsTypeArgument) TypeArgumentFirst() ITypeArgument {
	return t
}

func (t *WildcardExtendsTypeArgument) TypeArgumentList() util.IList[ITypeArgument] {
	return util.NewDefaultListWithElements[ITypeArgument](t)
}

func (t *WildcardExtendsTypeArgument) TypeArgumentSize() int {
	return 1
}

func (t *WildcardExtendsTypeArgument) GetType() IType {
	return t.Type
}

func (t *WildcardExtendsTypeArgument) IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool {
	if typeArgument.IsWildcardExtendsTypeArgument() {
		return t.Type.(ITypeArgument).IsTypeArgumentAssignableFrom(typeBounds, typeArgument.GetType().(ITypeArgument))
	} else if _, ok := typeArgument.(ITypeArgument); ok {
		return t.Type.(ITypeArgument).IsTypeArgumentAssignableFrom(typeBounds, typeArgument)
	}
	return false
}

func (t *WildcardExtendsTypeArgument) IsTypeArgumentList() bool {
	return false
}

func (t *WildcardExtendsTypeArgument) IsGenericTypeArgument() bool {
	return false
}

func (t *WildcardExtendsTypeArgument) IsInnerObjectTypeArgument() bool {
	return false
}

func (t *WildcardExtendsTypeArgument) IsObjectTypeArgument() bool {
	return false
}

func (t *WildcardExtendsTypeArgument) IsPrimitiveTypeArgument() bool {
	return false
}

func (t *WildcardExtendsTypeArgument) IsWildcardExtendsTypeArgument() bool {
	return true
}

func (t *WildcardExtendsTypeArgument) IsWildcardSuperTypeArgument() bool {
	return false
}

func (t *WildcardExtendsTypeArgument) IsWildcardTypeArgument() bool {
	return false
}

func (t *WildcardExtendsTypeArgument) AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor) {
	visitor.VisitWildcardExtendsTypeArgument(t)
}

func (t *WildcardExtendsTypeArgument) Equals(o interface{}) bool {
	if o == t {
		return true
	}

	if o == nil {
		return false
	}

	var other *WildcardExtendsTypeArgument

	switch o := o.(type) {
	case WildcardExtendsTypeArgument:
		other = &o
	case *WildcardExtendsTypeArgument:
		other = o
	default:
		return false
	}

	if t.Type != nil {
		return t.Type == other.Type
	}

	return other.Type == nil
}

func (t *WildcardExtendsTypeArgument) HashCode() int {
	if t.Type == nil {
		return 957014778
	}

	return 957014778 + t.Type.HashCode()
}

func (t *WildcardExtendsTypeArgument) String() string {
	return fmt.Sprintf("WildcardExtendsTypeArgument{? extends %s }", t.Type)
}

type WildcardSuperTypeArgument struct {
	Type IType
}

func (t *WildcardSuperTypeArgument) TypeArgumentFirst() ITypeArgument {
	return t
}

func (t *WildcardSuperTypeArgument) TypeArgumentList() util.IList[ITypeArgument] {
	return util.NewDefaultListWithElements[ITypeArgument](t)
}

func (t *WildcardSuperTypeArgument) TypeArgumentSize() int {
	return 1
}

func (t *WildcardSuperTypeArgument) GetType() IType {
	return t.Type
}

func (t *WildcardSuperTypeArgument) IsTypeArgumentAssignableFrom(typeBounds map[string]IType, typeArgument ITypeArgument) bool {
	if typeArgument.IsWildcardSuperTypeArgument() {
		return t.Type.(ITypeArgument).IsTypeArgumentAssignableFrom(typeBounds, typeArgument.GetType().(ITypeArgument))
	} else if _, ok := typeArgument.(ITypeArgument); ok {
		return t.Type.(ITypeArgument).IsTypeArgumentAssignableFrom(typeBounds, typeArgument)
	}
	return false
}

func (t *WildcardSuperTypeArgument) IsTypeArgumentList() bool {
	return false
}

func (t *WildcardSuperTypeArgument) IsGenericTypeArgument() bool {
	return false
}

func (t *WildcardSuperTypeArgument) IsInnerObjectTypeArgument() bool {
	return false
}

func (t *WildcardSuperTypeArgument) IsObjectTypeArgument() bool {
	return false
}

func (t *WildcardSuperTypeArgument) IsPrimitiveTypeArgument() bool {
	return false
}

func (t *WildcardSuperTypeArgument) IsWildcardExtendsTypeArgument() bool {
	return false
}

func (t *WildcardSuperTypeArgument) IsWildcardSuperTypeArgument() bool {
	return true
}

func (t *WildcardSuperTypeArgument) IsWildcardTypeArgument() bool {
	return false
}

func (t *WildcardSuperTypeArgument) AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor) {
	visitor.VisitWildcardSuperTypeArgument(t)
}

func (t *WildcardSuperTypeArgument) Equals(o interface{}) bool {
	if o == t {
		return true
	}

	if o == nil {
		return false
	}

	var other *WildcardSuperTypeArgument

	switch o := o.(type) {
	case WildcardSuperTypeArgument:
		other = &o
	case *WildcardSuperTypeArgument:
		other = o
	default:
		return false
	}

	if t.Type != nil {
		return t.Type == other.Type
	}

	return other.Type == nil
}

func (t *WildcardSuperTypeArgument) HashCode() int {
	if t.Type == nil {
		return 979510081
	}

	return 979510081 + t.Type.HashCode()
}

func (t *WildcardSuperTypeArgument) String() string {
	return fmt.Sprintf("WildcardSuperTypeArgument{? super %s }", t.Type)
}

/////////////////////////////////////////////////////////////////////

type WildcardTypeArgument struct {
}

/////////////////////////////////////////////////////////////////////

func (t *WildcardTypeArgument) TypeArgumentFirst() ITypeArgument {
	return t
}

func (t *WildcardTypeArgument) TypeArgumentList() util.IList[ITypeArgument] {
	return util.NewDefaultListWithElements[ITypeArgument](t)
}

func (t *WildcardTypeArgument) TypeArgumentSize() int {
	return 1
}

func (t *WildcardTypeArgument) GetType() IType {
	return &OtTypeUndefinedObject
}

func (t *WildcardTypeArgument) IsTypeArgumentAssignableFrom(_ map[string]IType, _ ITypeArgument) bool {
	return true
}

func (t *WildcardTypeArgument) IsTypeArgumentList() bool {
	return false
}

func (t *WildcardTypeArgument) IsGenericTypeArgument() bool {
	return false
}

func (t *WildcardTypeArgument) IsInnerObjectTypeArgument() bool {
	return false
}

func (t *WildcardTypeArgument) IsObjectTypeArgument() bool {
	return false
}

func (t *WildcardTypeArgument) IsPrimitiveTypeArgument() bool {
	return false
}

func (t *WildcardTypeArgument) IsWildcardExtendsTypeArgument() bool {
	return false
}

func (t *WildcardTypeArgument) IsWildcardSuperTypeArgument() bool {
	return false
}

func (t *WildcardTypeArgument) IsWildcardTypeArgument() bool {
	return true
}

func (t *WildcardTypeArgument) AcceptTypeArgumentVisitor(visitor ITypeArgumentVisitor) {
	visitor.VisitWildcardTypeArgument(t)
}

func (t *WildcardTypeArgument) Equals(o interface{}) bool {
	if o == t {
		return true
	}

	if o == nil {
		return false
	}

	var other *WildcardTypeArgument

	switch o := o.(type) {
	case WildcardTypeArgument:
		other = &o
	case *WildcardTypeArgument:
		other = o
	default:
		return false
	}

	return t == other
}

func (t *WildcardTypeArgument) HashCode() int {
	return hashCodeWithStruct(t)
}

func (t *WildcardTypeArgument) String() string {
	return "Wildcard{?}"
}

/////////////////////////////////////////////////////////////////////////
//  Functions
/////////////////////////////////////////////////////////////////////////

func hashCodeWithString(str string) int {
	h := fnv.New32a()
	_, err := h.Write([]byte(str))
	if err != nil {
		return -1
	}
	return int(h.Sum32())
}

func hashCodeWithStruct(data any) int {
	byteArray := toBytes(data)
	if byteArray == nil {
		return -1
	}

	h := fnv.New32a()
	_, err := h.Write(byteArray)
	if err != nil {
		return -1
	}

	return int(h.Sum32())
}

func toBytes(data any) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(data); err != nil {
		return nil
	}
	return buf.Bytes()
}
