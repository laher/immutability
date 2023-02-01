package immutability

type Set[T comparable] struct {
	items *Map[T, struct{}]
}

func NewSet[T comparable](ts ...T) *Set[T] {
	m := &Map[T, struct{}]{items: make(map[T]struct{}, len(ts))}
	for _, t := range ts {
		m.items[t] = struct{}{}
	}
	return &Set[T]{items: m}
}

func (s *Set[T]) Len() int {
	return len(s.items.items)
}

func (s *Set[T]) Has(i T) bool {
	_, ok := s.items.items[i]
	return ok
}

func (s *Set[T]) Items() []T {
	n := make([]T, 0, s.Len())
	for k := range s.items.items {
		n = append(n, k)
	}
	return n
}

func (s *Set[T]) Filter(f func(T) bool) *Set[T] {
	n := NewSet[T]()
	for i := range s.items.items {
		if f(i) {
			n.items.items[i] = struct{}{}
		}
	}
	return n
}

func (s *Set[T]) Map(mapper func(T) T) *Set[T] {
	n := &Set[T]{items: &Map[T, struct{}]{items: make(map[T]struct{}, s.Len())}}
	for i := range s.items.items {
		n.items.items[mapper(i)] = struct{}{}
	}
	return n
}

func (s *Set[T]) Reduce(mapper func(T, T) T, acc T) T {
	for i := range s.items.items {
		acc = mapper(i, acc)
	}
	return acc
}

func (s *Set[T]) ForEach(t func(T)) {
	for i := range s.items.items {
		t(i)
	}
}

func TransformSet[T comparable, T2 comparable](s *Set[T], mapper func(T) T2) *Set[T2] {
	n := &Set[T2]{items: &Map[T2, struct{}]{items: make(map[T2]struct{}, s.Len())}}
	for i := range s.items.items {
		n.items.items[mapper(i)] = struct{}{}
	}
	return n
}
