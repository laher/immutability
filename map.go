package immutability

type Map[K comparable, V any] struct {
	m map[K]V
}

func NewMapOf[K comparable, V any](in map[K]V) *Map[K, V] {
	if in == nil { // note that nil is used to indicate 'empty'. It's not valid to set a nil map
		return &Map[K, V]{m: map[K]V{}}
	}
	m := &Map[K, V]{m: make(map[K]V, len(in))}
	// copy map
	for k, v := range in {
		m.m[k] = v
	}
	return m
}

func NewMapEmpty[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{m: map[K]V{}}
}

func (m *Map[K, V]) Len() int {
	return len(m.m)
}

func (m *Map[K, V]) Get(k K) V {
	return m.m[k]
}

func (m *Map[K, V]) Items() map[K]V {
	out := make(map[K]V, len(m.m))
	// copy map
	for k, v := range m.m {
		out[k] = v
	}
	return out
}

// Filter returns a new Map containing only the items matched by f
func (m *Map[K, V]) Filter(f func(K, V) bool) *Map[K, V] {
	n := NewMapEmpty[K, V]()
	for k, v := range m.m {
		if f(k, v) {
			n.m[k] = v
		}
	}
	return n
}

// Map returns a new Map based on the result of the mapper for each key/value pair
func (m *Map[K, V]) Map(mapper func(K, V) (K, V)) *Map[K, V] {
	n := &Map[K, V]{m: make(map[K]V, len(m.m))}
	for k, v := range m.m {
		k, v := mapper(k, v)
		n.m[k] = v
	}
	return n
}

// Reduce returns a single value, accumulated by running mapper on each key/value pair
func (m *Map[K, V]) Reduce(mapper func(K, V, V) V, acc V) V {
	for k, v := range m.m {
		acc = mapper(k, v, acc)
	}
	return acc
}

// ForEach performs a func for each key/value pair
func (m *Map[K, V]) ForEach(t func(K, V)) {
	for k, v := range m.m {
		t(k, v)
	}
}

// TransformMapValues transforms a given map to a map with values of another type
func TransformMapValues[K comparable, V any, V2 any](m *Map[K, V], mapper func(V, K) V2) *Map[K, V2] {
	n := &Map[K, V2]{m: make(map[K]V2, len(m.m))}
	for k, v := range m.m {
		v2 := mapper(v, k)
		n.m[k] = v2
	}
	return n
}

// TransformMapKeys transforms a given map to a map with keys of another type
func TransformMapKeys[K comparable, V any, K2 comparable](m *Map[K, V], mapper func(K, V) K2) *Map[K2, V] {
	n := &Map[K2, V]{m: make(map[K2]V, len(m.m))}
	for k, v := range m.m {
		k2 := mapper(k, v)
		n.m[k2] = v
	}
	return n
}
