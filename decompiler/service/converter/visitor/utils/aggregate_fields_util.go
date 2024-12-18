package utils

import (
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/declaration"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
)

func Aggregate(fields util.IList[intsrv.IClassFileFieldDeclaration]) {
	if fields != nil {
		size := fields.Size()

		if size > 1 {
			firstIndex := 0
			lastIndex := 0
			firstField := fields.Get(0)

			for index := 1; index < size; index++ {
				field := fields.Get(index)

				if firstField.FirstLineNumber() == 0 || firstField.Flags() != field.Flags() || firstField.Type() != field.Type() {
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

func aggregate(fields util.IList[intsrv.IClassFileFieldDeclaration], firstField intsrv.IClassFileFieldDeclaration, firstIndex, lastIndex int) {
	if firstIndex < lastIndex {
		sublist := fields.SubList(firstIndex+1, lastIndex+1)

		length := lastIndex - firstIndex
		declarators := declaration.NewFieldDeclaratorsWithCapacity(length)
		bfd := firstField.FieldDeclarators()

		if bfd.IsList() {
			declarators.AddAll(bfd.ToSlice())
		} else {
			declarators.Add(bfd.First())
		}

		for _, f := range sublist.ToSlice() {
			bfd = f.FieldDeclarators()

			if bfd.IsList() {
				declarators.AddAll(bfd.ToSlice())
			} else {
				declarators.Add(bfd.First())
			}
		}

		firstField.SetFieldDeclarators(declarators)
		sublist.Clear()
	}
}
