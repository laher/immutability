package immutability

import "testing"

func TestMap(t *testing.T) {
	m1 := NewMapOf(map[string]int{
		"A": 1,
		"B": 2,
	})

	m2m := m1.Items()
	m2m["C"] = 3
	m2 := NewMapOf(m2m)
	if m1.Len() != 2 || m2.Len() != 3 {
		t.Errorf("lengths wrong - m1=%d, m2=%d", m1.Len(), m2.Len())
	}

}
