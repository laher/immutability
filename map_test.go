package immutability

import (
	"fmt"
	"testing"
)

func ExampleMap() {
	input := map[string]int{
		"A": 1,
		"B": 2,
	}
	m1 := NewMapOf(input)
	res := m1.Reduce(func(k string, v int, acc int) int {
		return v + acc
	}, 0)
	fmt.Printf("Sum: %d, ", res)

	input["C"] = 3
	res2 := m1.Reduce(func(k string, v int, acc int) int {
		return v + acc
	}, 0)
	fmt.Printf("Sum: %d, ", res2)
	// Output: Sum: 3, Sum: 3,
}

func TestMap(t *testing.T) {
	input := map[string]int{
		"A": 1,
		"B": 2,
	}
	m1 := NewMapOf(input)
	m2m := m1.Items()
	m2m["C"] = 3
	m2 := NewMapOf(m2m)
	if m1.Len() != 2 || m2.Len() != 3 {
		t.Errorf("lengths wrong - m1=%d, m2=%d", m1.Len(), m2.Len())
	}
}
