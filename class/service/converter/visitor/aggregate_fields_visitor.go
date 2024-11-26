package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
)

func NewAggregateFieldsVisitor() *AggregateFieldsVisitor {
	return &AggregateFieldsVisitor{}
}

type AggregateFieldsVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor
}

func (v *AggregateFieldsVisitor) VisitAnnotationDeclaration(declaration intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AggregateFieldsVisitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	bodyDeclaration := declaration.(intsrv.IClassFileBodyDeclaration)
	// Aggregate fields
	Aggregate(bodyDeclaration.FieldDeclarations())
}

func (v *AggregateFieldsVisitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AggregateFieldsVisitor) VisitEnumDeclaration(declaration intmod.IEnumDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AggregateFieldsVisitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func Aggregate(fields []intsrv.IClassFileFieldDeclaration) {
	if fields != nil {
		size := len(fields)

		if size > 1 {
			var firstIndex, lastIndex int
			firstField := fields[0]

			for index := 1; index < size; index++ {
				field := fields[index]

				if firstField.FirstLineNumber() == 0 || firstField.Flags() != field.Flags() || !(firstField.Type() == field.Type()) {
					firstField = field
					firstIndex = index
					lastIndex = index
				} else {
					lineNumber := field.FirstLineNumber()

					if lineNumber > 0 {
						if lineNumber == firstField.FirstLineNumber() {
							// Compatible field -> Keep index
							lastIndex = index
						} else {
							// Aggregate declarators from 'firstIndex' to 'lastIndex'
							aggregate(fields, firstField, firstIndex, lastIndex)

							length := lastIndex - firstIndex
							index -= length
							size -= length

							firstField = field
							firstIndex = index
							lastIndex = index
						}
					}
				}
			}

			// Aggregate declarators from 'firstIndex' to 'lastIndex'
			aggregate(fields, firstField, firstIndex, lastIndex)
		}
	}
}

func aggregate(fields []intsrv.IClassFileFieldDeclaration, firstField intsrv.IClassFileFieldDeclaration, firstIndex, lastIndex int) {
	if firstIndex < lastIndex {
		sublist := fields[firstIndex+1 : lastIndex+1]

		length := lastIndex - firstIndex
		declarators := declaration.NewFieldDeclarators(length)
		bfd := firstField.FieldDeclarators()

		if bfd.IsList() {
			l := bfd.ToSlice()
			declarators.AddAll(l)
		} else {
			declarators.Add(bfd.First())
		}

		for _, f := range sublist {
			bfd = f.FieldDeclarators()

			if bfd.IsList() {
				l := bfd.ToSlice()
				declarators.AddAll(l)
			} else {
				declarators.Add(bfd.First())
			}
		}

		firstField.SetFieldDeclarators(declarators.(intmod.IFieldDeclarator))
	}
}
