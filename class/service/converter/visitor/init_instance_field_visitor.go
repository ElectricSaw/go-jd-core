package visitor

import (
	intmod "github.com/ElectricSaw/go-jd-core/class/interfaces/model"
	intsrv "github.com/ElectricSaw/go-jd-core/class/interfaces/service"
	"github.com/ElectricSaw/go-jd-core/class/model/javasyntax"
	moddecl "github.com/ElectricSaw/go-jd-core/class/model/javasyntax/declaration"
	"github.com/ElectricSaw/go-jd-core/class/util"
)

func NewInitInstanceFieldVisitor() *InitInstanceFieldVisitor {
	return &InitInstanceFieldVisitor{
		searchFirstLineNumberVisitor:   NewSearchFirstLineNumberVisitor(),
		fieldDeclarators:               make(map[string]intmod.IFieldDeclarator),
		datas:                          util.NewDefaultList[IData](),
		putFields:                      util.NewDefaultList[intmod.IExpression](),
		lineNumber:                     intmod.UnknownLineNumber,
		containsLocalVariableReference: false,
	}
}

type InitInstanceFieldVisitor struct {
	javasyntax.AbstractJavaSyntaxVisitor

	searchFirstLineNumberVisitor   *SearchFirstLineNumberVisitor
	fieldDeclarators               map[string]intmod.IFieldDeclarator
	datas                          util.IList[IData]
	putFields                      util.IList[intmod.IExpression]
	lineNumber                     int
	containsLocalVariableReference bool
}

func (v *InitInstanceFieldVisitor) VisitAnnotationDeclaration(declaration intmod.IAnnotationDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *InitInstanceFieldVisitor) VisitClassDeclaration(declaration intmod.IClassDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *InitInstanceFieldVisitor) VisitEnumDeclaration(declaration intmod.IEnumDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *InitInstanceFieldVisitor) VisitInterfaceDeclaration(declaration intmod.IInterfaceDeclaration) {
	v.SafeAcceptDeclaration(declaration.BodyDeclaration())
}

func (v *InitInstanceFieldVisitor) VisitBodyDeclaration(declaration intmod.IBodyDeclaration) {
	bodyDeclaration := declaration.(intsrv.IClassFileBodyDeclaration)

	// Init attributes
	v.fieldDeclarators = make(map[string]intmod.IFieldDeclarator)
	v.datas.Clear()
	v.putFields.Clear()
	// Visit fields
	tmp1 := make([]intmod.IDeclaration, 0)
	for _, item := range bodyDeclaration.FieldDeclarations() {
		tmp1 = append(tmp1, item)
	}
	v.SafeAcceptListDeclaration(tmp1)
	// Visit methods
	tmp2 := make([]intmod.IDeclaration, 0)
	for _, item := range bodyDeclaration.MethodDeclarations() {
		tmp2 = append(tmp2, item)
	}
	v.SafeAcceptListDeclaration(tmp2)
	// Init values
	v.updateFieldsAndConstructors()
}

func (v *InitInstanceFieldVisitor) VisitFieldDeclaration(declaration intmod.IFieldDeclaration) {
	if (declaration.Flags() & intmod.FlagStatic) == 0 {
		declaration.FieldDeclarators().AcceptDeclaration(v)
	}
}

func (v *InitInstanceFieldVisitor) VisitConstructorDeclaration(declaration intmod.IConstructorDeclaration) {
	cfcd := declaration.(intsrv.IClassFileConstructorDeclaration)

	if (cfcd.Statements() != nil) && cfcd.Statements().IsStatements() {
		statements := cfcd.Statements().(intmod.IStatements)
		iterator := statements.ListIterator()
		superConstructorCall := v.searchSuperConstructorCall(iterator)

		if superConstructorCall != nil {
			internalTypeName := cfcd.ClassFile().InternalTypeName()

			v.datas.Add(NewData(cfcd, statements, iterator.NextIndex()))

			if v.datas.Size() == 1 {
				var firstLineNumber int

				if (cfcd.Flags() & intmod.FlagSynthetic) != 0 {
					firstLineNumber = intmod.UnknownLineNumber
				} else {
					firstLineNumber = superConstructorCall.LineNumber()

					if (superConstructorCall.Descriptor() == "()V") &&
						(firstLineNumber != intmod.UnknownLineNumber) && iterator.HasNext() {
						if (v.lineNumber == intmod.UnknownLineNumber) || (v.lineNumber >= firstLineNumber) {
							v.searchFirstLineNumberVisitor.Init()
							iterator.Next().AcceptStatement(v.searchFirstLineNumberVisitor)
							iterator.Previous()

							ln := v.searchFirstLineNumberVisitor.LineNumber()
							if (ln != intmod.UnknownLineNumber) && (ln >= firstLineNumber) {
								firstLineNumber = intmod.UnknownLineNumber
							}
						}
					}
				}

				v.initPutFields(internalTypeName, firstLineNumber, iterator)
			} else {
				v.filterPutFields(internalTypeName, iterator)
			}
		}
	}
}

func (v *InitInstanceFieldVisitor) VisitMethodDeclaration(declaration intmod.IMethodDeclaration) {
	v.lineNumber = declaration.(intsrv.IClassFileMethodDeclaration).FirstLineNumber()
}

func (v *InitInstanceFieldVisitor) VisitNewExpression(expression intmod.INewExpression) {
	v.SafeAcceptExpression(expression.Parameters())
}

func (v *InitInstanceFieldVisitor) VisitStaticInitializerDeclaration(_ intmod.IStaticInitializerDeclaration) {
}

func (v *InitInstanceFieldVisitor) VisitFieldDeclarator(declaration intmod.IFieldDeclarator) {
	v.fieldDeclarators[declaration.Name()] = declaration
}

func (v *InitInstanceFieldVisitor) VisitLocalVariableReferenceExpression(_ intmod.ILocalVariableReferenceExpression) {
	v.containsLocalVariableReference = true
}

func (v *InitInstanceFieldVisitor) searchSuperConstructorCall(iterator util.IListIterator[intmod.IStatement]) intmod.ISuperConstructorInvocationExpression {
	for iterator.HasNext() {
		expression := iterator.Next().Expression()

		if expression.IsSuperConstructorInvocationExpression() {
			return expression.(intmod.ISuperConstructorInvocationExpression)
		}

		if expression.IsConstructorInvocationExpression() {
			break
		}
	}

	return nil
}

func (v *InitInstanceFieldVisitor) initPutFields(internalTypeName string, firstLineNumber int, iterator util.IListIterator[intmod.IStatement]) {
	fieldNames := util.NewSet[string]()
	var expression intmod.IExpression

	for iterator.HasNext() {
		statement := iterator.Next()

		if !statement.IsExpressionStatement() {
			break
		}

		expression = statement.Expression()

		if !expression.IsBinaryOperatorExpression() {
			break
		}

		if (expression.Operator() != "=") || !expression.LeftExpression().IsFieldReferenceExpression() {
			break
		}

		fre := expression.LeftExpression().(intmod.IFieldReferenceExpression)

		if fre.InternalTypeName() != internalTypeName || !fre.Expression().IsThisExpression() {
			break
		}

		fieldName := fre.Name()

		if fieldNames.Contains(fieldName) {
			break
		}

		v.containsLocalVariableReference = false
		expression.RightExpression().Accept(v)

		if v.containsLocalVariableReference {
			break
		}

		v.putFields.Add(expression)
		fieldNames.Add(fieldName)
		expression = nil
	}

	var lastLineNumber int

	if expression == nil {
		if firstLineNumber == intmod.UnknownLineNumber {
			lastLineNumber = intmod.UnknownLineNumber
		} else {
			lastLineNumber = firstLineNumber + 1
		}
	} else {
		lastLineNumber = expression.LineNumber()
	}

	if firstLineNumber < lastLineNumber {
		ite := v.putFields.Iterator()

		for ite.HasNext() {
			lineNumber := ite.Next().LineNumber()

			if (firstLineNumber <= lineNumber) && (lastLineNumber <= lastLineNumber) {
				lastLineNumber++
				_ = ite.Remove()
			}
		}
	}
}

func (v *InitInstanceFieldVisitor) filterPutFields(internalTypeName string, iterator util.IListIterator[intmod.IStatement]) {
	putFieldIterator := v.putFields.Iterator()
	index := 0

	for iterator.HasNext() && putFieldIterator.HasNext() {
		expression := iterator.Next().Expression()

		if !expression.IsBinaryOperatorExpression() {
			break
		}

		if expression.Operator() != "=" || !expression.LeftExpression().IsFieldReferenceExpression() {
			break
		}

		fre := expression.LeftExpression().(intmod.IFieldReferenceExpression)

		if fre.InternalTypeName() != internalTypeName {
			break
		}

		putField := putFieldIterator.Next()

		if expression.LineNumber() != putField.LineNumber() {
			break
		}

		if fre.Name() != putField.LeftExpression().Name() {
			break
		}

		index++
	}

	if index < v.putFields.Size() {
		// Cut extra putFields
		v.putFields.SubList(index, v.putFields.Size()).Clear()
	}
}

func (v *InitInstanceFieldVisitor) updateFieldsAndConstructors() {
	count := v.putFields.Size()

	if count > 0 {
		// Init values
		for _, putField := range v.putFields.ToSlice() {
			decl := v.fieldDeclarators[putField.LeftExpression().Name()]

			if decl != nil {
				expression := putField.RightExpression()
				decl.SetVariableInitializer(moddecl.NewExpressionVariableInitializer(expression))
				decl.FieldDeclaration().(intsrv.IClassFileFieldDeclaration).SetFirstLineNumber(expression.LineNumber())
			}
		}

		// Update data : remove init field statements
		for _, data := range v.datas.ToSlice() {
			data.Statements().SubList(data.Index(), data.Index()+count).Clear()

			if data.Statements().IsEmpty() {
				data.Declaration().SetStatements(nil)
				data.Declaration().SetFirstLineNumber(0)
			} else {
				v.searchFirstLineNumberVisitor.Init()
				v.searchFirstLineNumberVisitor.VisitStatements(data.Statements())

				firstLineNumber := v.searchFirstLineNumberVisitor.LineNumber()

				if firstLineNumber == -1 {
					data.Declaration().SetFirstLineNumber(0)
				} else {
					data.Declaration().SetFirstLineNumber(firstLineNumber)
				}
			}
		}
	}
}

func NewData(declaration intsrv.IClassFileConstructorDeclaration, statements intmod.IStatements, index int) IData {
	return &Data{
		declaration: declaration,
		statements:  statements,
		index:       index,
	}
}

type IData interface {
	Declaration() intsrv.IClassFileConstructorDeclaration
	SetDeclaration(declaration intsrv.IClassFileConstructorDeclaration)
	Statements() intmod.IStatements
	SetStatements(statements intmod.IStatements)
	Index() int
	SetIndex(index int)
}

type Data struct {
	declaration intsrv.IClassFileConstructorDeclaration
	statements  intmod.IStatements
	index       int
}

func (d *Data) Declaration() intsrv.IClassFileConstructorDeclaration {
	return d.declaration
}

func (d *Data) SetDeclaration(declaration intsrv.IClassFileConstructorDeclaration) {
	d.declaration = declaration
}

func (d *Data) Statements() intmod.IStatements {
	return d.statements
}

func (d *Data) SetStatements(statements intmod.IStatements) {
	d.statements = statements
}

func (d *Data) Index() int {
	return d.index
}

func (d *Data) SetIndex(index int) {
	d.index = index
}
