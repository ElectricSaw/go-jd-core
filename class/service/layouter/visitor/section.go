package visitor

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	"bitbucket.org/coontec/go-jd-core/class/util"
	"fmt"
	"math"
)

func NewSection(flexibleFragments util.IList[intmod.IFlexibleFragment],
	fixedFragment intmod.IFixedFragment, previousSection ISection) ISection {
	s := &Section{
		flexibleFragments: flexibleFragments,
		fixedFragment:     fixedFragment,
		previousSection:   previousSection,
		rate:              0,
		lastLineCount:     -1,
	}

	previousLineNumber := 0

	if previousSection == nil {
		previousLineNumber = 1
	} else {
		previousSection.SetNextSection(s)
		previousLineNumber = previousSection.FixedFragment().LastLineNumber()
	}

	if fixedFragment == nil {
		s.targetLineCount = 0
	} else {
		s.targetLineCount = fixedFragment.FirstLineNumber() - previousLineNumber
	}

	return s
}

type ISection interface {
	FlexibleFragments() util.IList[intmod.IFlexibleFragment]
	SetFlexibleFragments(flexibleFragments util.IList[intmod.IFlexibleFragment])
	FixedFragment() intmod.IFixedFragment
	SetFixedFragment(fixedFragment intmod.IFixedFragment)
	PreviousSection() ISection
	SetPreviousSection(previousSection ISection)
	NextSection() ISection
	SetNextSection(nextSection ISection)
	TargetLineCount() int
	SetTargetLineCount(targetLineCount int)
	Rate() int
	SetRate(rate int)
	LastLineCount() int
	SetLastLineCount(lastLineCount int)
	Delta() int
	SetDelta(delta int)

	UpdateRate()
	Layout(force bool) bool
	ReleaseConstraints(holder *VisitorsHolder) bool
}

type Section struct {
	flexibleFragments util.IList[intmod.IFlexibleFragment]
	fixedFragment     intmod.IFixedFragment
	previousSection   ISection
	nextSection       ISection
	targetLineCount   int
	rate              int
	lastLineCount     int
	delta             int
}

func (s *Section) FlexibleFragments() util.IList[intmod.IFlexibleFragment] {
	return s.flexibleFragments
}

func (s *Section) SetFlexibleFragments(flexibleFragments util.IList[intmod.IFlexibleFragment]) {
	s.flexibleFragments = flexibleFragments
}

func (s *Section) FixedFragment() intmod.IFixedFragment {
	return s.fixedFragment
}

func (s *Section) SetFixedFragment(fixedFragment intmod.IFixedFragment) {
	s.fixedFragment = fixedFragment
}

func (s *Section) PreviousSection() ISection {
	return s.previousSection
}

func (s *Section) SetPreviousSection(previousSection ISection) {
	s.previousSection = previousSection
}

func (s *Section) NextSection() ISection {
	return s.nextSection
}

func (s *Section) SetNextSection(nextSection ISection) {
	s.nextSection = nextSection
}

func (s *Section) TargetLineCount() int {
	return s.targetLineCount
}

func (s *Section) SetTargetLineCount(targetLineCount int) {
	s.targetLineCount = targetLineCount
}

func (s *Section) Rate() int {
	return s.rate
}

func (s *Section) SetRate(rate int) {
	s.rate = rate
}

func (s *Section) LastLineCount() int {
	return s.lastLineCount
}

func (s *Section) SetLastLineCount(lastLineCount int) {
	s.lastLineCount = lastLineCount
}

func (s *Section) Delta() int {
	return s.delta
}

func (s *Section) SetDelta(delta int) {
	s.delta = delta
}

func (s *Section) UpdateRate() {
	s.rate = 0

	for _, flexibleFragment := range s.flexibleFragments.ToSlice() {
		if flexibleFragment.InitialLineCount() > flexibleFragment.LineCount() {
			s.rate += flexibleFragment.InitialLineCount() - flexibleFragment.LineCount()
		}
	}
}

func (s *Section) Layout(force bool) bool {
	// Skip layout of last section
	if s.fixedFragment != nil {
		// Compute line count
		currentLineCount := 0

		for _, flexibleFragment := range s.flexibleFragments.ToSlice() {
			currentLineCount += flexibleFragment.LineCount()
		}

		// Do not re-layout if nothing has changed
		if force || (s.lastLineCount != currentLineCount) {
			s.lastLineCount = currentLineCount

			if s.targetLineCount != currentLineCount {
				filteredFlexibleFragments := NewAutoGrowthList(s)
				constrainedFlexibleFragments := util.NewDefaultListWithCapacity[intmod.IFlexibleFragment](s.flexibleFragments.Size())

				if s.targetLineCount > currentLineCount {
					// Expands fragments
					oldDelta := s.targetLineCount - currentLineCount
					delta := oldDelta

					for _, flexibleFragment := range s.flexibleFragments.ToSlice() {
						if flexibleFragment.LineCount() < flexibleFragment.MaximalLineCount() {
							// Keep only expandable fragments
							filteredFlexibleFragments.Get(flexibleFragment.Weight()).Add(flexibleFragment)
						}
					}

					// First, expand compacted fragments
					for _, flexibleFragments := range filteredFlexibleFragments.ToSlice() {
						constrainedFlexibleFragments.Clear()

						for _, flexibleFragment := range flexibleFragments.ToSlice() {
							if flexibleFragment.LineCount() < flexibleFragment.InitialLineCount() {
								// Store compacted flexibleFragments
								constrainedFlexibleFragments.Add(flexibleFragment)
							}
						}

						s.expand(constrainedFlexibleFragments, force)
						if delta == 0 {
							break
						}
					}

					// Next, expand all
					if delta > 0 {
						for _, flexibleFragments := range filteredFlexibleFragments.ToSlice() {
							s.expand(flexibleFragments, force)
							if delta == 0 {
								break
							}
						}
					}

					// Something changed ?
					return oldDelta != delta
				} else {
					// Compacts fragments
					s.delta = currentLineCount - s.targetLineCount
					oldDelta := s.delta

					for _, flexibleFragment := range s.flexibleFragments.ToSlice() {
						if flexibleFragment.MinimalLineCount() < flexibleFragment.LineCount() {
							// Keep only compactable fragments
							filteredFlexibleFragments.Get(flexibleFragment.Weight()).Add(flexibleFragment)
						}
					}

					// First, compact expanded fragments
					for _, flexibleFragments := range filteredFlexibleFragments.ToSlice() {
						constrainedFlexibleFragments.Clear()

						for _, flexibleFragment := range flexibleFragments.ToSlice() {
							if flexibleFragment.LineCount() > flexibleFragment.InitialLineCount() {
								// Store expanded flexibleFragments
								constrainedFlexibleFragments.Add(flexibleFragment)
							}
						}

						s.compact(constrainedFlexibleFragments, force)
						if s.delta == 0 {
							break
						}
					}

					// Next, compact all
					if s.delta > 0 {
						for _, flexibleFragments := range filteredFlexibleFragments.ToSlice() {
							s.compact(flexibleFragments, force)
							if s.delta == 0 {
								break
							}
						}
					}

					// Something changed ?
					return oldDelta != s.delta
				}
			}
		}
	}

	return false
}

func (s *Section) expand(flexibleFragments util.IList[intmod.IFlexibleFragment], force bool) {
	oldDelta := math.MaxInt

	for (s.delta > 0) && (s.delta < oldDelta) {
		oldDelta = s.delta

		for _, flexibleFragment := range flexibleFragments.ToSlice() {
			if flexibleFragment.IncLineCount(force) {
				s.delta--
				if s.delta == 0 {
					break
				}
			}
		}
	}
}

func (s *Section) compact(flexibleFragments util.IList[intmod.IFlexibleFragment], force bool) {
	oldDelta := math.MaxInt

	for (s.delta > 0) && (s.delta < oldDelta) {
		oldDelta = s.delta

		for _, flexibleFragment := range flexibleFragments.ToSlice() {
			if flexibleFragment.DecLineCount(force) {
				s.delta--
				if s.delta == 0 {
					break
				}
			}
		}
	}
}

func (s *Section) ReleaseConstraints(holder *VisitorsHolder) bool {
	flexibleCount := s.flexibleFragments.Size()
	backwardSearchStartIndexesVisitor := holder.BackwardSearchStartIndexesVisitor()
	forwardSearchEndIndexesVisitor := holder.ForwardSearchEndIndexesVisitor()
	forwardSearchVisitor := holder.ForwardSearchVisitor()
	backwardSearchVisitor := holder.BackwardSearchVisitor()
	iterator := s.flexibleFragments.ListIterator()
	iterator.SetCursor(flexibleCount)

	backwardSearchStartIndexesVisitor.Reset()
	forwardSearchEndIndexesVisitor.Reset()

	for iterator.HasPrevious() && backwardSearchStartIndexesVisitor.IsEnabled() {
		iterator.Previous().AcceptFragmentVisitor(backwardSearchStartIndexesVisitor)
	}

	for _, flexibleFragment := range s.flexibleFragments.ToSlice() {
		flexibleFragment.AcceptFragmentVisitor(forwardSearchEndIndexesVisitor)
		if !forwardSearchEndIndexesVisitor.IsEnabled() {
			break
		}
	}

	size := backwardSearchStartIndexesVisitor.Size()
	nextSection := s.searchNextSection(forwardSearchVisitor)

	if (size > 1) && (nextSection != nil) {
		index1 := flexibleCount - 1 - backwardSearchStartIndexesVisitor.IndexAt(size/2)
		index2 := flexibleCount - 1 - backwardSearchStartIndexesVisitor.IndexAt(0)
		nextIndex := forwardSearchVisitor.Index()

		size = forwardSearchEndIndexesVisitor.Size()

		if size > 1 {
			index3 := forwardSearchEndIndexesVisitor.IndexAt(0) + 1
			index4 := forwardSearchEndIndexesVisitor.IndexAt(size/2) + 1
			previousSection := s.searchPreviousSection(backwardSearchVisitor)

			if nextSection.Rate() > previousSection.Rate() {
				index := previousSection.FlexibleFragments().Size() - backwardSearchVisitor.Index()
				previousSection.addFragmentsAtEnd(holder, index, s.extract(index3, index4))
			} else {
				nextSection.addFragmentsAtBeginning(holder, nextIndex, s.extract(index1, index2))
			}
		} else {
			nextSection.addFragmentsAtBeginning(holder, nextIndex, s.extract(index1, index2))
		}

		return true
	} else {
		size = forwardSearchEndIndexesVisitor.Size()

		if size > 1 {
			index3 := forwardSearchEndIndexesVisitor.IndexAt(0) + 1
			index4 := forwardSearchEndIndexesVisitor.IndexAt(size/2) + 1
			previousSection := s.searchPreviousSection(backwardSearchVisitor)

			if (size > 1) && (previousSection != nil) {
				index := previousSection.FlexibleFragments().Size() - backwardSearchVisitor.Index()
				previousSection.addFragmentsAtEnd(holder, index, s.extract(index3, index4))
				return true
			}
		}
	}

	return false
}

func (s *Section) searchNextSection(visitor ISearchMovableBlockFragmentVisitor) *Section {
	section := s.NextSection()

	visitor.Reset()

	for section != nil {
		visitor.ResetIndex()

		for _, flexibleFragment := range section.FlexibleFragments().ToSlice() {
			flexibleFragment.AcceptFragmentVisitor(visitor)
			if visitor.Depth() == 0 {
				return section.(*Section)
			}
		}

		section = section.NextSection()
	}

	return nil
}

func (s *Section) searchPreviousSection(visitor ISearchMovableBlockFragmentVisitor) *Section {
	section := s.PreviousSection()

	visitor.Reset()

	for section != nil {
		flexibleFragments := section.FlexibleFragments()
		iterator := flexibleFragments.ListIterator()
		iterator.SetCursor(flexibleFragments.Size())

		visitor.ResetIndex()

		for iterator.HasPrevious() {
			iterator.Previous().AcceptFragmentVisitor(visitor)
			if visitor.Depth() == 0 {
				return section.(*Section)
			}
		}

		section = section.PreviousSection()
	}

	return nil
}

func (s *Section) addFragmentsAtBeginning(holder *VisitorsHolder, index int, flexibleFragments util.IList[intmod.IFlexibleFragment]) {
	visitor := holder.ForwardSearchVisitor()
	iterator := flexibleFragments.ListIterator()
	iterator.SetCursor(flexibleFragments.Size())

	// Extract separators
	visitor.Reset()

	for iterator.HasPrevious() {
		iterator.Previous().AcceptFragmentVisitor(visitor)
		if visitor.Depth() == 0 {
			break
		}
	}

	// assert (visitor.Index() < flexibleFragments.Size()) && (visitor.Index() > 1);

	index1 := flexibleFragments.Size() + 1 - visitor.Index()

	// Insert other fragments
	_ = s.flexibleFragments.AddAllAt(index, flexibleFragments.SubList(0, index1).ToSlice())
	// Insert separator at beginning

	_ = s.flexibleFragments.AddAllAt(index, flexibleFragments.SubList(index1, flexibleFragments.Size()).ToSlice())

	s.resetLineCount()
}

func (s *Section) addFragmentsAtEnd(holder *VisitorsHolder, index int, flexibleFragments util.IList[intmod.IFlexibleFragment]) {
	visitor := holder.ForwardSearchVisitor()

	// Extract separators
	visitor.Reset()

	for _, flexibleFragment := range flexibleFragments.ToSlice() {
		flexibleFragment.AcceptFragmentVisitor(visitor)
		if visitor.Depth() == 2 {
			break
		}
	}

	// assert (visitor.Index() < flexibleFragments.Size()) && (visitor.Index() > 1);

	index1 := visitor.Index() - 1

	// Insert other fragments
	_ = s.flexibleFragments.AddAllAt(index, flexibleFragments.SubList(0, index1).ToSlice())
	// Insert separator at end
	_ = s.flexibleFragments.AddAllAt(index, flexibleFragments.SubList(index1, flexibleFragments.Size()).ToSlice())

	s.resetLineCount()
}

func (s *Section) extract(index1, index2 int) util.IList[intmod.IFlexibleFragment] {
	s.resetLineCount()

	subList := s.flexibleFragments.SubList(index1, index2)
	fragmentsToMove := util.NewDefaultListWithSlice(subList.ToSlice())

	subList.Clear()

	return fragmentsToMove
}

func (s *Section) resetLineCount() {
	for _, flexibleFragment := range s.flexibleFragments.ToSlice() {
		flexibleFragment.ResetLineCount()
	}
}

func (s *Section) String() string {
	sb := fmt.Sprintf("Section{flexibleFragments.size=%d, ", s.flexibleFragments.Size())
	if s.fixedFragment == nil {
		sb += "fixedFragment.firstLineNumber=undefined"
	} else {
		sb += fmt.Sprintf("fixedFragment.firstLineNumber=%d, ", s.fixedFragment.FirstLineNumber())
	}
	sb += fmt.Sprintf("rate=%d}", s.rate)
	return sb
}

func NewAutoGrowthList(parent ISection) *AutoGrowthList {
	return &AutoGrowthList{
		parent:        parent,
		elements:      make([]util.IList[intmod.IFlexibleFragment], 21),
		iteratorIndex: 0,
	}
}

type AutoGrowthList struct {
	parent        ISection
	elements      []util.IList[intmod.IFlexibleFragment]
	iteratorIndex int
}

func (l *AutoGrowthList) Set(index int, element util.IList[intmod.IFlexibleFragment]) {
	l.ensureCapacity(index)
	l.elements[index] = element
}

func (l *AutoGrowthList) Get(index int) util.IList[intmod.IFlexibleFragment] {
	l.ensureCapacity(index)

	element := l.elements[index]
	if element == nil {
		element = util.NewDefaultListWithCapacity[intmod.IFlexibleFragment](l.parent.FlexibleFragments().Size())
		l.elements[index] = element
	}

	return element
}

func (l *AutoGrowthList) Reverse() {
	for i, j := 0, len(l.elements)-1; i < j; i, j = i+1, j-1 {
		l.elements[i], l.elements[j] = l.elements[j], l.elements[i]
	}
}

func (l *AutoGrowthList) ensureCapacity(minCapacity int) {
	if len(l.elements) < minCapacity {
		tmp := make([]util.IList[intmod.IFlexibleFragment], minCapacity+10)
		copy(tmp, l.elements)
		l.elements = tmp
	}
}

func (l *AutoGrowthList) Iterator() util.IIterator[util.IList[intmod.IFlexibleFragment]] {
	length := len(l.elements)
	l.iteratorIndex = 0

	for l.iteratorIndex < length && l.elements[l.iteratorIndex] == nil {
		l.iteratorIndex++
	}

	return util.NewIteratorWithSlice(l.elements)
}

func (l *AutoGrowthList) HasNext() bool {
	return l.iteratorIndex < len(l.elements)
}

func (l *AutoGrowthList) Next() util.IList[intmod.IFlexibleFragment] {
	element := l.elements[l.iteratorIndex]
	l.iteratorIndex++
	length := len(l.elements)

	for l.iteratorIndex < length && l.elements[l.iteratorIndex] == nil {
		l.iteratorIndex++
	}

	return element
}

func (l *AutoGrowthList) Remove() {
}

func (l *AutoGrowthList) ToSlice() []util.IList[intmod.IFlexibleFragment] {
	return l.elements
}
