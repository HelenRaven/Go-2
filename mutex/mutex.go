package mutex

import "sync"

type Set struct {
	sync.Mutex
	data map[int]string
}

func NewSet() *Set {
	return &Set{
		data: map[int]string{},
	}
}

func (s *Set) Add(i int, str string) {
	s.Lock()
	s.data[i] = str
	s.Unlock()
}

func (s *Set) Has(i int) bool {
	s.Lock()
	defer s.Unlock()
	_, ok := s.data[i]
	return ok
}

type SetRW struct {
	sync.RWMutex
	data map[int]string
}

func NewSetRW() *SetRW {
	return &SetRW{
		data: map[int]string{},
	}
}

func (s *SetRW) Add(i int, str string) {
	s.Lock()
	s.data[i] = str
	s.Unlock()
}

func (s *SetRW) Has(i int) bool {
	s.RLock()
	defer s.RUnlock()
	_, ok := s.data[i]
	return ok
}
