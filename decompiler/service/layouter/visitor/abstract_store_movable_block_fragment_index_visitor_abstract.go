package visitor

func NewAbstractStoreMovableBlockFragmentIndexVisitorAbstract() *AbstractStoreMovableBlockFragmentIndexVisitorAbstract {
	return &AbstractStoreMovableBlockFragmentIndexVisitorAbstract{
		AbstractSearchMovableBlockFragmentVisitor: *NewAbstractSearchMovableBlockFragmentVisitor(),
		indexes: make([]int, 10),
		size:    0,
		enabled: true,
	}
}

type IStoreMovableBlockFragmentIndexVisitorAbstract interface {
	ISearchMovableBlockFragmentVisitor

	IndexAt(index int) int
	Size() int
	IsEnabled() bool
	Reset()
}

type AbstractStoreMovableBlockFragmentIndexVisitorAbstract struct {
	AbstractSearchMovableBlockFragmentVisitor

	indexes []int
	size    int
	enabled bool
}

func (v *AbstractStoreMovableBlockFragmentIndexVisitorAbstract) IndexAt(index int) int {
	return v.indexes[index]
}

func (v *AbstractStoreMovableBlockFragmentIndexVisitorAbstract) Size() int {
	return v.size
}

func (v *AbstractStoreMovableBlockFragmentIndexVisitorAbstract) IsEnabled() bool {
	return v.enabled
}

func (v *AbstractStoreMovableBlockFragmentIndexVisitorAbstract) Reset() {
	v.size = 0
	v.depth = 1
	v.index = 0
	v.enabled = true
}

func (v *AbstractStoreMovableBlockFragmentIndexVisitorAbstract) storeIndex() {
	if v.size == len(v.indexes) {
		tmp := make([]int, len(v.indexes)*2)
		copy(tmp, v.indexes)
		v.indexes = tmp
	}

	v.indexes[v.size] = v.index
	v.size++
}
