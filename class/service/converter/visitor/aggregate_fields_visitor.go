package visitor

import (
	"bitbucket.org/coontec/javaClass/class/model/javasyntax"
	"bitbucket.org/coontec/javaClass/class/model/javasyntax/declaration"
	srvdecl "bitbucket.org/coontec/javaClass/class/service/converter/model/javasyntax/declaration"
)

type AggregateFieldsVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor
}

func (v *AggregateFieldsVisitor) VisitAnnotationDeclaration(declaration *declaration.AnnotationDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AggregateFieldsVisitor) VisitBodyDeclaration(declaration declaration.IBodyDeclaration) {
	bodyDeclaration := declaration.(*srvdecl.ClassFileBodyDeclaration)
	// Aggregate fields
	Aggregate(bodyDeclaration.FieldDeclarations())
}

func (v *AggregateFieldsVisitor) VisitClassDeclaration(declaration *declaration.ClassDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AggregateFieldsVisitor) VisitEnumDeclaration(declaration *declaration.EnumDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *AggregateFieldsVisitor) VisitInterfaceDeclaration(declaration *declaration.InterfaceDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func Aggregate(fields []srvdecl.ClassFileFieldDeclaration) {
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

func aggregate(fields []srvdecl.ClassFileFieldDeclaration, firstField srvdecl.ClassFileFieldDeclaration, firstIndex, lastIndex int) {
	if firstIndex < lastIndex {
		sublist := fields[firstIndex+1 : lastIndex+1]

		length := lastIndex - firstIndex
		declarators := declaration.NewFieldDeclarators(length)
		bfd := firstField.FieldDeclarators()

		if bfd.IsList() {
			l := bfd.List()
			declarators.AddAll(l)
		} else {
			declarators.Add(bfd.First())
		}

		for _, f := range sublist {
			bfd = f.FieldDeclarators()

			if bfd.IsList() {
				l := bfd.List()
				declarators.AddAll(l)
			} else {
				declarators.Add(bfd.First())
			}
		}

		firstField.SetFieldDeclarators(declarators)
	}
}
