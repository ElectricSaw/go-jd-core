package localvariable

import (
	intmod "bitbucket.org/coontec/go-jd-core/class/interfaces/model"
	intsrv "bitbucket.org/coontec/go-jd-core/class/interfaces/service"
	_type "bitbucket.org/coontec/go-jd-core/class/model/javasyntax/type"
)

func NewLocalVariableSet() intsrv.ILocalVariableSet {
	return &LocalVariableSet{
		array: make([]intsrv.ILocalVariable, 0, 10),
		size:  0,
	}
}

type LocalVariableSet struct {
	array []intsrv.ILocalVariable
	size  int
}

func (s *LocalVariableSet) Add(index int, newLV intsrv.ILocalVariable) {
	if index >= len(s.array) {
		// Increases array
		tmp := make([]intsrv.ILocalVariable, index*2)
		copy(tmp, s.array)
		s.array = tmp
		// Store
		s.array[index] = newLV
	} else {
		lv := s.array[index]

		if lv == nil {
			s.array[index] = newLV
		} else if lv.FromOffset() < newLV.FromOffset() {
			if newLV != lv {
				return
			}
			newLV.SetNext(lv)
			s.array[index] = newLV
		} else {
			previous := lv

			lv = lv.Next()

			for lv != nil && (lv.FromOffset() > newLV.FromOffset()) {
				previous = lv
				lv = lv.Next()
			}

			previous.SetNext(newLV)
			newLV.SetNext(lv)
		}
	}

	s.size++
}

func (s *LocalVariableSet) Root(index int) intsrv.ILocalVariable {
	if index < len(s.array) {
		lv := s.array[index]

		if lv != nil {
			for lv.Next() != nil {
				lv = lv.Next()
			}
			return lv
		}
	}

	return nil
}

func (s *LocalVariableSet) Remove(index, offset int) intsrv.ILocalVariable {
	if index < len(s.array) {
		var previous intsrv.ILocalVariable
		lv := s.array[index]

		for lv != nil {
			if lv.FromOffset() <= offset {
				if previous == nil {
					s.array[index] = lv.Next()
				} else {
					previous.SetNext(lv.Next())
				}

				s.size--
				lv.SetNext(nil)

				return lv
			}
			previous = lv
			lv = lv.Next()
		}
	}
	return nil
}

func (s *LocalVariableSet) Get(index, offset int) intsrv.ILocalVariable {
	if index < len(s.array) {
		lv := s.array[index]

		for lv != nil {
			if lv.FromOffset() <= offset {
				return lv
			}
			lv = lv.Next()
		}
	}

	return nil
}

func (s *LocalVariableSet) IsEmpty() bool {
	return s.size == 0
}

func (s *LocalVariableSet) Update(index, offset int, typ intmod.IObjectType) {
	if index < len(s.array) {
		lv := s.array[index]

		for lv != nil {
			if lv.FromOffset() == offset {
				olv := lv.(*ObjectLocalVariable)
				olv.typ = typ.(intmod.IType)
				break
			}

			lv = lv.Next()
		}
	}
}

func (s *LocalVariableSet) update(index, offset int, typ *_type.GenericType) {
	if index < len(s.array) {
		var previous intsrv.ILocalVariable
		lv := s.array[index]

		for lv != nil {
			if lv.FromOffset() == offset {
				glv := NewGenericLocalVariableWithAll(index, lv.FromOffset(), typ, lv.Name())
				glv.SetNext(lv.Next())

				if previous == nil {
					s.array[index] = glv
				} else {
					previous.SetNext(glv)
				}
				break
			}
		}

		previous = lv
		if lv == nil {
			return
		}
		lv = lv.Next()
	}
}

func (s *LocalVariableSet) initialize(rootFrame intsrv.IFrame) []intsrv.ILocalVariable {
	cache := make([]intsrv.ILocalVariable, 0, s.size)

	for index := len(s.array) - 1; index >= 0; index-- {
		lv := s.array[index]

		if lv != nil {
			var previous intsrv.ILocalVariable

			for lv.Next() != nil {
				previous = lv
				lv = lv.Next()
			}

			if lv.FromOffset() == 0 {
				if previous == nil {
					s.array[index] = lv.Next()
				} else {
					previous.SetNext(lv.Next())
				}

				s.size--
				lv.SetNext(nil)
				rootFrame.AddLocalVariable(lv)
				cache = append(cache, lv)
			}
		}
	}

	return cache
}
