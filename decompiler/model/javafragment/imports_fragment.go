package javafragment

import (
	"sort"
)

func NewImportsFragment(weight int) ImportsFragment {
	return ImportsFragment{
		FlexibleFragment: NewFlexibleFragment(0, -1,
			-1, weight, "Imports"),
		ImportMap: make(map[string]Import),
	}
}

type ImportsFragment struct {
	FlexibleFragment

	ImportMap map[string]Import
}

func (f *ImportsFragment) AddImport(internalName, qualifiedName string) {
	imp, ok := f.ImportMap[internalName]
	if ok {
		imp.IncCounter()
	} else {
		f.ImportMap[internalName] = NewImport(internalName, qualifiedName)
	}
}

func (f *ImportsFragment) IncCounter(internalName string) bool {
	imp, ok := f.ImportMap[internalName]
	if ok {
		imp.IncCounter()
		return true
	} else {
		return false
	}
}

func (f *ImportsFragment) IsEmpty() bool {
	return len(f.ImportMap) == 0
}

func (f *ImportsFragment) InitLineCounts() {
	f.MaximalLineCount = len(f.ImportMap)
	f.InitialLineCount = f.MaximalLineCount
	f.LineCount = f.MaximalLineCount
}

func (f *ImportsFragment) Contains(internalName string) bool {
	_, ok := f.ImportMap[internalName]
	return ok
}

func (f *ImportsFragment) Import(internalName string) (Import, bool) {
	if imp0rt, ok := f.ImportMap[internalName]; ok {
		return imp0rt, true
	}
	var zero Import
	return zero, false
}

func (f *ImportsFragment) Imports() []Import {
	lineCount := f.LineCount
	size := len(f.ImportMap)

	imports := make([]Import, 0, len(f.ImportMap))
	for _, v := range f.ImportMap {
		imports = append(imports, v)
	}

	if lineCount < size {
		sort.Slice(imports, func(i, j int) bool {
			return imports[i].Counter > imports[j].Counter
		})

		subList := imports[lineCount:size]

		for _, imp0rt := range subList {
			delete(f.ImportMap, imp0rt.InternalName)
		}
	}
	return imports
}

func (f *ImportsFragment) Accept(visitor IJavaFragmentVisitor) {
	visitor.VisitImportsFragment(f)
}

func NewImport(internalName string, qualifiedName string) Import {
	return Import{
		InternalName:  internalName,
		QualifiedName: qualifiedName,
		Counter:       1,
	}
}

type Import struct {
	InternalName  string
	QualifiedName string
	Counter       int
}

func (i *Import) IncCounter() {
	i.Counter++
}
