package javafragment

import (
	"bitbucket.org/coontec/javaClass/java/model/fragment"
	"sort"
)

var CountComparator = ImportsCountComparator{}

func NewImportsFragment(weight int) ImportsFragment {
	return ImportsFragment{
		FlexibleFragment: fragment.NewFlexibleFragment(0, -1, -1, weight, "Imports"),
	}
}

type ImportsFragment struct {
	fragment.FlexibleFragment

	importMap map[string]Import
}

func (f *ImportsFragment) addImport(internalName string, qualifiedName string) {
	imp, ok := f.importMap[internalName]
	if ok {
		imp.IncCounter()
	} else {
		f.importMap[internalName] = NewImport(internalName, qualifiedName)
	}
}

func (f *ImportsFragment) IncCounter(internalName string) bool {
	imp, ok := f.importMap[internalName]
	if ok {
		imp.IncCounter()
		return true
	} else {
		return false
	}
}

func (f *ImportsFragment) isEmpty() bool {
	return len(f.importMap) == 0
}

func (f *ImportsFragment) InitLineCounts() {
	f.MaximalLineCount = len(f.importMap)
	f.InitialLineCount = f.MaximalLineCount
	f.SetLineCount(f.MaximalLineCount)
}

func (f *ImportsFragment) Contains(internalName string) bool {
	_, ok := f.importMap[internalName]
	return ok
}

func (f *ImportsFragment) Imports() []Import {
	lineCount := f.LineCount()
	size := len(f.importMap)

	imports := make([]Import, 0, len(f.importMap))
	for _, v := range f.importMap {
		imports = append(imports, v)
	}

	if lineCount < size {
		sort.Slice(imports, func(i, j int) bool {
			return imports[i].counter > imports[j].counter
		})

		subList := imports[lineCount:size]

		for _, imp0rt := range subList {
			delete(f.importMap, imp0rt.internalName)
		}
	}
	return imports
}

func (f *ImportsFragment) Accept(visitor JavaFragmentVisitor) {
	visitor.VisitImportsFragment(f)
}

func NewImport(internalName string, qualifiedName string) Import {
	return Import{
		internalName:  internalName,
		qualifiedName: qualifiedName,
		counter:       1,
	}
}

type Import struct {
	internalName  string
	qualifiedName string
	counter       int
}

func (i *Import) InternalName() string {
	return i.internalName
}

func (i *Import) QualifiedName() string {
	return i.qualifiedName
}

func (i *Import) Counter() int {
	return i.counter
}

func (i *Import) IncCounter() {
	i.counter++
}

type Comparator interface {
	Compare(a, b Import) int
}

type ImportsCountComparator struct {
}

func (c ImportsCountComparator) Compare(tr1 Import, tr2 Import) int {
	return tr2.counter - tr1.counter
}
