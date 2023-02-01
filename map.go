package immutability

// Map stores a map of keys and values
type Map[K comparable, V any] struct {
	items map[K]V
}

// NewMapOf initialises a Map
func NewMapOf[K comparable, V any](in map[K]V) *Map[K, V] {
	if in == nil { // note that nil is used to indicate 'empty'. It's not valid to set a nil map
		return &Map[K, V]{items: map[K]V{}}
	}
	m := &Map[K, V]{items: make(map[K]V, len(in))}
	for k, v := range in {
		m.items[k] = v
	}
	return m
}

func NewMapEmpty[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{items: map[K]V{}}
}

func (m *Map[K, V]) Len() int {
	return len(m.items)
}

func (m *Map[K, V]) Get(k K) V {
	return m.items[k]
}

func (m *Map[K, V]) Items() map[K]V {
	out := make(map[K]V, len(m.items))
	// copy map
	for k, v := range m.items {
		out[k] = v
	}
	return out
}

func (m *Map[K, V]) Keys() *Set[K] {
	n := &Map[K, struct{}]{items: make(map[K]struct{}, m.Len())}
	for k := range m.items {
		n.items[k] = struct{}{}
	}
	return &Set[K]{items: n}
}

func (m *Map[K, V]) Values() []V {
	n := make([]V, 0, m.Len())
	for _, v := range m.items {
		n = append(n, v)
	}
	return n
}

// Filter returns a new Map containing only the items matched by f
func (m *Map[K, V]) Filter(f func(K, V) bool) *Map[K, V] {
	n := NewMapEmpty[K, V]()
	for k, v := range m.items {
		if f(k, v) {
			n.items[k] = v
		}
	}
	return n
}

// Map returns a new Map based on the result of the mapper for each key/value pair
func (m *Map[K, V]) Map(mapper func(K, V) (K, V)) *Map[K, V] {
	n := &Map[K, V]{items: make(map[K]V, len(m.items))}
	for k, v := range m.items {
		k, v := mapper(k, v)
		n.items[k] = v
	}
	return n
}

// Reduce returns a single value, accumulated by running mapper on each key/value pair
func (m *Map[K, V]) Reduce(mapper func(K, V, V) V, acc V) V {
	for k, v := range m.items {
		acc = mapper(k, v, acc)
	}
	return acc
}

// Reduce returns a single key, accumulated by running mapper on each key/value pair
func (m *Map[K, V]) ReduceKey(mapper func(K, V, K) K, acc K) K {
	for k, v := range m.items {
		acc = mapper(k, v, acc)
	}
	return acc
}

// ForEach performs a func for each key/value pair
func (m *Map[K, V]) ForEach(t func(K, V)) {
	for k, v := range m.items {
		t(k, v)
	}
}

// TransformMapValues transforms a given map to a map with values of another type
func TransformMapValues[K comparable, V any, V2 any](m *Map[K, V], mapper func(V, K) V2) *Map[K, V2] {
	n := &Map[K, V2]{items: make(map[K]V2, len(m.items))}
	for k, v := range m.items {
		v2 := mapper(v, k)
		n.items[k] = v2
	}
	return n
}

// TransformMapKeys transforms a given map to a map with keys of another type
func TransformMapKeys[K comparable, V any, K2 comparable](m *Map[K, V], mapper func(K, V) K2) *Map[K2, V] {
	n := &Map[K2, V]{items: make(map[K2]V, len(m.items))}
	for k, v := range m.items {
		k2 := mapper(k, v)
		n.items[k2] = v
	}
	return n
}
