package javafragment

import (
	intmod "github.com/ElectricSaw/go-jd-core/decompiler/interfaces/model"
	"github.com/ElectricSaw/go-jd-core/decompiler/model/fragment"
	"sort"
)

func NewImportsFragment(weight int) intmod.IImportsFragment {
	return &ImportsFragment{
		FlexibleFragment: *fragment.NewFlexibleFragment(0, -1,
			-1, weight, "Imports").(*fragment.FlexibleFragment),
		importMap: make(map[string]intmod.IImport),
	}
}

type ImportsFragment struct {
	fragment.FlexibleFragment

	importMap map[string]intmod.IImport
}

func (f *ImportsFragment) AddImport(internalName, qualifiedName string) {
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

func (f *ImportsFragment) IsEmpty() bool {
	return len(f.importMap) == 0
}

func (f *ImportsFragment) InitLineCounts() {
	f.SetMaximalLineCount(len(f.importMap))
	f.SetInitialLineCount(f.MaximalLineCount())
	f.SetLineCount(f.MaximalLineCount())
}

func (f *ImportsFragment) Contains(internalName string) bool {
	_, ok := f.importMap[internalName]
	return ok
}

func (f *ImportsFragment) Import(internalName string) intmod.IImport {
	if imp0rt, ok := f.importMap[internalName]; ok {
		return imp0rt
	}
	return nil
}

func (f *ImportsFragment) Imports() []intmod.IImport {
	lineCount := f.LineCount()
	size := len(f.importMap)

	imports := make([]intmod.IImport, 0, len(f.importMap))
	for _, v := range f.importMap {
		imports = append(imports, v)
	}

	if lineCount < size {
		sort.Slice(imports, func(i, j int) bool {
			return imports[i].Counter() > imports[j].Counter()
		})

		subList := imports[lineCount:size]

		for _, imp0rt := range subList {
			delete(f.importMap, imp0rt.InternalName())
		}
	}
	return imports
}

func (f *ImportsFragment) Accept(visitor intmod.IJavaFragmentVisitor) {
	visitor.VisitImportsFragment(f)
}

func NewImport(internalName string, qualifiedName string) intmod.IImport {
	return &Import{
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

func (i *Import) SetInternalName(internalName string) {
	i.internalName = internalName
}

func (i *Import) QualifiedName() string {
	return i.qualifiedName
}

func (i *Import) SetQualifiedName(qualifiedName string) {
	i.qualifiedName = qualifiedName
}

func (i *Import) Counter() int {
	return i.counter
}

func (i *Import) SetCounter(counter int) {
	i.counter = counter
}

func (i *Import) IncCounter() {
	i.counter++
}
