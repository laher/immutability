package immutability

type List[T any] struct {
	items []T
}

// NewList creates a List from the items received
// We don't want the caller to reach a pointer to the internal slice
func NewList[T any](ts ...T) *List[T] {
	return &List[T]{items: ts}
}

func (l *List[T]) Len() int {
	return len(l.items)
}

func (l *List[T]) Get(i int) T {
	return l.items[i]
}

// Items returns a copy of the items in the List.
// We don't want the caller to reach a pointer to the internal slice
func (l *List[T]) Items() []T {
	dst := make([]T, len(l.items))
	copy(dst, l.items)
	return dst
}

// Filter returns a new List containing only the items matched by f
func (l *List[T]) Filter(f func(T) bool) *List[T] {
	n := NewList[T]()
	for _, item := range l.items {
		if f(item) {
			n.items = append(n.items, item)
		}
	}
	return n
}

// Map returns a new List based on the result of the mapper for each element
func (l *List[T]) Map(mapper func(T, int) T) *List[T] {
	n := NewList[T]()
	for idx, item := range l.items {
		n.items = append(n.items, mapper(item, idx))
	}
	return n
}

// FlatMap returns a new List based on the result of the mapper for each element
func (l *List[T]) FlatMap(mapper func(T, int) []T) *List[T] {
	n := NewList[T]()
	for idx, item := range l.items {
		vs := mapper(item, idx)
		n.items = append(n.items, vs...)
	}
	return n
}

// Reduce returns a single item, accumulated by running mapper on each element
func (l *List[T]) Reduce(mapper func(item T, acc T, idx int) T, acc T) T {
	for idx, item := range l.items {
		acc = mapper(item, acc, idx)
	}
	return acc
}

// ForEach performs a func for each item in a list
func (l *List[T]) ForEach(t func(T, int)) {
	for idx, item := range l.items {
		t(item, idx)
	}
}

// TransformList transforms a given list to a list of another type
func TransformList[T any, U any](l *List[T], mapper func(T, int) U) *List[U] {
	n := &List[U]{items: make([]U, 0, len(l.items))} // prime list of appropriate size
	for idx, item := range l.items {
		n.items = append(n.items, mapper(item, idx))
	}
	return n
}
