package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax"
	"bitbucket.org/coontec/go-jd-core/class/model/javasyntax/declaration"
	"math"
	"sort"
)

func NewSortMembersVisitor() *SortMembersVisitor {
	return &SortMembersVisitor{}
}

type SortMembersVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor
}

func (v *SortMembersVisitor) VisitAnnotationDeclaration(declaration intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *SortMembersVisitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	bodyDeclaration := declaration.(intsrv.IClassFileBodyDeclaration)
	innerTypes := bodyDeclaration.InnerTypeDeclarations()
	// Merge fields, getters & inner types
	members := Merge(ConvertField(bodyDeclaration.FieldDeclarations()),
		ConvertMethod(bodyDeclaration.MethodDeclarations()), ConvertTypes(innerTypes))
	bodyDeclaration.SetMemberDeclarations(members)
}

func (v *SortMembersVisitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *SortMembersVisitor) VisitEnumDeclaration(declaration intmod.IEnumDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *SortMembersVisitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func Merge(fields, methods, innerTypes []intsrv.IClassFileMemberDeclaration) intmod.IMemberDeclarations {
	var size int

	if fields != nil {
		size = len(fields)
	} else {
		size = 0
	}

	if methods != nil {
		size += len(methods)
	}

	if innerTypes != nil {
		size += len(innerTypes)
	}

	result := declaration.NewMemberDeclarationsWithSize(size)

	merge(result.Elements(), fields)
	merge(result.Elements(), methods)
	merge(result.Elements(), innerTypes)

	return result
}

func merge(result []intmod.IMemberDeclaration, members []intsrv.IClassFileMemberDeclaration) {
	if members != nil && len(members) != 0 {
		members = memberSort(members)

		if len(members) != 0 {
			for _, member := range members {
				result = append(result, member)
			}
		} else {
			resultIndex := 0
			resultLength := len(result)
			listStartIndex := 0
			listEndIndex := 0
			listLength := len(members)
			listLineNumber := 0

			for listEndIndex < listLength {

				for listEndIndex < listLength {
					listLineNumber = members[listEndIndex].FirstLineNumber()
					listEndIndex++
					if listLineNumber > 0 {
						break
					}
				}

				if listLineNumber == 0 {
					// Add end of list to result
					for _, member := range members[listStartIndex : listStartIndex+listEndIndex] {
						result = append(result, member)
					}
				} else {
					// Search insert index in result

					for resultIndex < listLength {
						member := result[resultIndex].(intsrv.IClassFileMemberDeclaration)
						resultLineNumber := member.FirstLineNumber()
						if resultLineNumber > listLineNumber {
							break
						}
						resultIndex++
					}

					// Add end of list to result
					tmp := make([]intmod.IMemberDeclaration, 0, len(result))
					tmp = append(tmp, result[:resultIndex]...)
					for i := listStartIndex; i < listEndIndex; i++ {
						result = append(result, members[i])
					}
					tmp = append(tmp, result[resultIndex:]...)

					subListLength := listEndIndex - listStartIndex
					resultIndex += subListLength
					resultLength += subListLength
					listStartIndex = listEndIndex
				}
			}
		}
	}
}

func memberSort(members []intsrv.IClassFileMemberDeclaration) []intsrv.IClassFileMemberDeclaration {
	order := 0
	lastLineNumber := 0

	// Detect order type
	for _, member := range members {
		lineNumber := member.FirstLineNumber()

		if (lineNumber > 0) && (lineNumber != lastLineNumber) {
			if lastLineNumber > 0 {
				if order == 0 { // Unknown order
					if lineNumber > lastLineNumber {
						order = 1
					} else {
						order = 2
					}
				} else if order == 1 { // Ascendant order
					if lineNumber < lastLineNumber {
						order = 3 // Random order
						break
					}
				} else if order == 2 { // Descendant order
					if lineNumber > lastLineNumber {
						order = 3 // Random order
						break
					}
				}
			}

			lastLineNumber = lineNumber
		}
	}

	// Sort
	switch order {
	case 2: // Descendant order
		reverse(members)
		break
	case 3: // Random order : ascendant memberSort and set unknown line number members at the end
		sort.Slice(members, func(i, j int) bool {
			lineNumber1 := members[i].FirstLineNumber()
			lineNumber2 := members[j].FirstLineNumber()

			if lineNumber1 == 0 {
				lineNumber1 = math.MaxInt
			}

			if lineNumber2 == 0 {
				lineNumber2 = math.MaxInt
			}

			return lineNumber1 < lineNumber2
		})
		break
	}

	return members
}

func reverse(slice []intsrv.IClassFileMemberDeclaration) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func ConvertField(list []intsrv.IClassFileFieldDeclaration) []intsrv.IClassFileMemberDeclaration {
	ret := make([]intsrv.IClassFileMemberDeclaration, 0, len(list))
	for _, item := range list {
		ret = append(ret, item)
	}
	return ret
}

func ConvertMethod(list []intsrv.IClassFileConstructorOrMethodDeclaration) []intsrv.IClassFileMemberDeclaration {
	ret := make([]intsrv.IClassFileMemberDeclaration, 0, len(list))
	for _, item := range list {
		ret = append(ret, item)
	}
	return ret
}

func ConvertTypes(list []intsrv.IClassFileTypeDeclaration) []intsrv.IClassFileMemberDeclaration {
	ret := make([]intsrv.IClassFileMemberDeclaration, 0, len(list))
	for _, item := range list {
		ret = append(ret, item)
	}
	return ret
}
