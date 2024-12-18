package utils

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/javasyntax/declaration"
	"github.com/ElectricSaw/go-jd-core/decompiler/util"
	"math"
)

var GlobalMemberDeclarationComparator = NewMemberDeclarationComparator()

func Merge(fields util.IList[intsrv.IClassFileMemberDeclaration],
	methods util.IList[intsrv.IClassFileMemberDeclaration],
	innerTypes util.IList[intsrv.IClassFileMemberDeclaration]) intmod.IMemberDeclarations {
	size := 0

	if fields != nil {
		size = fields.Size()
	} else {
		size = 0
	}

	if methods != nil {
		size += methods.Size()
	}

	if innerTypes != nil {
		size += innerTypes.Size()
	}

	result := declaration.NewMemberDeclarationsWithCapacity(size)

	merge(result, fields)
	merge(result, methods)
	merge(result, innerTypes)

	return result
}

func merge(result util.IList[intmod.IMemberDeclaration], members util.IList[intsrv.IClassFileMemberDeclaration]) {
	if (members != nil) && !members.IsEmpty() {
		Sort(members)

		if result.IsEmpty() {
			tmp := make([]intmod.IMemberDeclaration, 0, members.Size())
			for _, item := range members.ToSlice() {
				tmp = append(tmp, item.(intmod.IMemberDeclaration))
			}
			result.AddAll(tmp)
		} else {
			resultIndex := 0
			resultLength := result.Size()
			listStartIndex := 0
			listEndIndex := 0
			listLength := members.Size()
			listLineNumber := 0

			for listEndIndex < listLength {
				// Search first line number > 0
				for listEndIndex < listLength {
					listLineNumber = members.Get(listEndIndex).FirstLineNumber()
					listEndIndex++
					if listLineNumber > 0 {
						break
					}
				}

				if listLineNumber == 0 {
					// Add end of list to result
					tmp := make([]intmod.IMemberDeclaration, 0, members.Size())
					for _, item := range members.SubList(listStartIndex, listEndIndex).ToSlice() {
						tmp = append(tmp, item.(intmod.IMemberDeclaration))
					}
					result.AddAll(tmp)
				} else {
					// Search insert index in result
					for resultIndex < resultLength {
						member := result.Get(resultIndex).(intsrv.IClassFileMemberDeclaration)
						resultLineNumber := member.FirstLineNumber()
						if resultLineNumber > listLineNumber {
							break
						}
						resultIndex++
					}

					// Add end of list to result
					tmp := make([]intmod.IMemberDeclaration, 0, members.Size())
					for _, item := range members.SubList(listStartIndex, listEndIndex).ToSlice() {
						tmp = append(tmp, item.(intmod.IMemberDeclaration))
					}
					result.AddAll(tmp)

					subListLength := listEndIndex - listStartIndex
					resultIndex += subListLength
					resultLength += subListLength
					listStartIndex = listEndIndex
				}
			}
		}
	}
}

func Sort(members util.IList[intsrv.IClassFileMemberDeclaration]) {
	order := 0
	lastLineNumber := 0

	// Detect order type
	for _, member := range members.ToSlice() {
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
		members.Reverse()
		break
	case 3: // Random order : ascendant sort and set unknown line number members at the end
		members.Sort(func(i, j int) bool {
			lineNumber1 := members.Get(i).FirstLineNumber()
			lineNumber2 := members.Get(j).FirstLineNumber()

			if lineNumber1 == 0 {
				lineNumber1 = math.MaxInt
			}

			if lineNumber2 == 0 {
				lineNumber2 = math.MaxInt
			}
			// FIXED: 체크 필요.
			return (lineNumber1 - lineNumber2) > 0
		})
		break
	}
}

func NewMemberDeclarationComparator() *MemberDeclarationComparator {
	return &MemberDeclarationComparator{}
}

type MemberDeclarationComparator struct {
}

func (c *MemberDeclarationComparator) compare(md1, md2 intsrv.IClassFileMemberDeclaration) int {
	lineNumber1 := md1.FirstLineNumber()
	lineNumber2 := md2.FirstLineNumber()

	if lineNumber1 == 0 {
		lineNumber1 = math.MaxInt
	}

	if lineNumber2 == 0 {
		lineNumber2 = math.MaxInt
	}

	return lineNumber1 - lineNumber2
}
