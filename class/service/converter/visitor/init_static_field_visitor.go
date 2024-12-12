package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	moddec "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/declaration"
	modsts "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/statement"
	_type "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/type"
	srvdecl "github.com/ElectricSaw/go-jd-core/class/service/converter/model/javasyntax/declaration"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewInitStaticFieldVisitor() intsrv.IInitStaticFieldVisitor {
	return &InitStaticFieldVisitor{
		searchFirstLineNumberVisitor:        NewSearchFirstLineNumberVisitor(),
		searchLocalVariableReferenceVisitor: NewSearchLocalVariableReferenceVisitor(),
		fields:                              make(map[string]intmod.IFieldDeclarator),
	}
}

type InitStaticFieldVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	searchFirstLineNumberVisitor        intsrv.ISearchFirstLineNumberVisitor
	searchLocalVariableReferenceVisitor intsrv.ISearchLocalVariableReferenceVisitor
	internalTypeName                    string
	fields                              map[string]intmod.IFieldDeclarator
	methods                             util.IList[intsrv.IClassFileConstructorOrMethodDeclaration]
	deleteStaticDeclaration             bool
}

func (v *InitStaticFieldVisitor) SetInternalTypeName(internalTypeName string) {
	v.internalTypeName = internalTypeName
}

func (v *InitStaticFieldVisitor) VisitAnnotationDeclaration(decl intmod.IAnnotationDeclaration) {
	v.internalTypeName = decl.InternalTypeName()
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *InitStaticFieldVisitor) VisitClassDeclaration(decl intmod.IClassDeclaration) {
	v.internalTypeName = decl.InternalTypeName()
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *InitStaticFieldVisitor) VisitEnumDeclaration(decl intmod.IEnumDeclaration) {
	v.internalTypeName = decl.InternalTypeName()
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *InitStaticFieldVisitor) VisitInterfaceDeclaration(decl intmod.IInterfaceDeclaration) {
	v.internalTypeName = decl.InternalTypeName()
	v.SafeAcceptDeclaration(decl.BodyDeclaration())
}

func (v *InitStaticFieldVisitor) VisitBodyDeclaration(decl intmod.IBodyDeclaration) {
	bodyDeclaration := decl.(intsrv.IClassFileBodyDeclaration)

	// Store field declarations
	v.fields = make(map[string]intmod.IFieldDeclarator)
	tmp := make([]intmod.IDeclaration, 0)
	for _, item := range bodyDeclaration.FieldDeclarations() {
		tmp = append(tmp, item)
	}
	v.SafeAcceptListDeclaration(tmp)

	if !(len(v.fields) == 0) {
		methods := util.NewDefaultListWithSlice(bodyDeclaration.MethodDeclarations())

		if methods != nil {
			v.deleteStaticDeclaration = false

			length := methods.Size()
			for i := 0; i < length; i++ {
				methods.Get(i).AcceptDeclaration(v)

				if !v.deleteStaticDeclaration {
					if v.deleteStaticDeclaration {
						methods.RemoveAt(i)
					}
					break
				}
			}
		}
	}

	tmp = make([]intmod.IDeclaration, 0)
	for _, item := range bodyDeclaration.InnerTypeDeclarations() {
		tmp = append(tmp, item)
	}
	v.SafeAcceptListDeclaration(tmp)
}

func (v *InitStaticFieldVisitor) VisitFieldDeclarator(decl intmod.IFieldDeclarator) {
	v.fields[decl.Name()] = decl
}

func (v *InitStaticFieldVisitor) VisitConstructorDeclaration(_ intmod.IConstructorDeclaration) {}

func (v *InitStaticFieldVisitor) VisitMethodDeclaration(_ intmod.IMethodDeclaration) {}

func (v *InitStaticFieldVisitor) VisitInstanceInitializerDeclaration(_ intmod.IInstanceInitializerDeclaration) {
}

func (v *InitStaticFieldVisitor) VisitStaticInitializerDeclaration(decl intmod.IStaticInitializerDeclaration) {
	sid := decl.(intsrv.IClassFileStaticInitializerDeclaration)
	statements := sid.Statements()

	if statements != nil {
		if statements.IsList() {
			list := statements.ToList()

			// Multiple statements
			if (list.Size() > 0) && v.isAssertionsDisabledStatement(list.First()) {
				// Remove assert initialization statement
				list.RemoveFirst()
			}

			length := list.Size()
			for i := 0; i < length; i++ {
				if v.setStaticFieldInitializer(list.Get(i)) {
					if i > 0 {
						// Split 'static' block
						var newStatements intmod.IStatement

						if i == 1 {
							newStatements = list.RemoveFirst()
						} else {
							subList := list.SubList(0, i)
							newStatements = modsts.NewStatementsWithList(subList)
							subList.Clear()
						}

						// Removes statements from original list
						length -= newStatements.Size()
						i = 0

						v.addStaticInitializerDeclaration(sid, v.getFirstLineNumber(newStatements), newStatements)
					}
					// Remove field initialization statement
					list.RemoveAt(i)
					i--
					length--
				}
			}
		} else {
			// Single statement
			if v.isAssertionsDisabledStatement(statements.First()) {
				// Remove assert initialization statement
				statements = nil
			}
			if (statements != nil) && v.setStaticFieldInitializer(statements.First()) {
				// Remove field initialization statement
				statements = nil
			}
		}

		if (statements == nil) || (statements.Size() == 0) {
			v.deleteStaticDeclaration = true
		} else {
			firstLineNumber := v.getFirstLineNumber(statements)
			if firstLineNumber == -1 {
				sid.SetFirstLineNumber(0)
			} else {
				sid.SetFirstLineNumber(firstLineNumber)
			}
			v.deleteStaticDeclaration = false
		}
	}
}

func (v *InitStaticFieldVisitor) isAssertionsDisabledStatement(state intmod.IStatement) bool {
	expr := state.Expression()

	if expr.LeftExpression().IsFieldReferenceExpression() {
		fre := expr.LeftExpression().(intmod.IFieldReferenceExpression)

		if fre.Type() == _type.PtTypeBoolean &&
			fre.InternalTypeName() == v.internalTypeName &&
			fre.Name() == "$assertionsDisabled" {
			return true
		}
	}

	return false
}

func (v *InitStaticFieldVisitor) setStaticFieldInitializer(state intmod.IStatement) bool {
	expr := state.Expression()

	if expr.LeftExpression().IsFieldReferenceExpression() {
		fre := expr.LeftExpression().(intmod.IFieldReferenceExpression)

		if fre.InternalTypeName() == v.internalTypeName {
			fdr := v.fields[fre.Name()].(intmod.IFieldDeclarator)

			if fdr != nil && fdr.VariableInitializer() == nil {
				fdn := fdr.FieldDeclaration()

				if fdn.Flags()&intmod.FlagStatic != 0 && fdn.Type().Descriptor() == fre.Descriptor() {
					expr = expr.RightExpression()

					v.searchLocalVariableReferenceVisitor.Init(-1)
					expr.Accept(v.searchLocalVariableReferenceVisitor)

					if !v.searchLocalVariableReferenceVisitor.ContainsReference() {
						fdr.SetVariableInitializer(moddec.NewExpressionVariableInitializer(expr))
						fdr.FieldDeclaration().(intsrv.IClassFileFieldDeclaration).SetFirstLineNumber(expr.LineNumber())
						return true
					}
				}
			}
		}
	}

	return false
}

func (v *InitStaticFieldVisitor) getFirstLineNumber(state intmod.IStatement) int {
	v.searchFirstLineNumberVisitor.Init()
	state.AcceptStatement(v.searchFirstLineNumberVisitor)
	return v.searchFirstLineNumberVisitor.LineNumber()
}

func (v *InitStaticFieldVisitor) addStaticInitializerDeclaration(
	sid intsrv.IClassFileStaticInitializerDeclaration,
	lineNumber int, statements intmod.IStatement) {
	v.methods.Add(srvdecl.NewClassFileStaticInitializerDeclaration2(
		sid.BodyDeclaration(), sid.ClassFile(), sid.Method(), sid.Bindings(),
		sid.TypeBounds(), lineNumber, statements))
}
