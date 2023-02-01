package immutability

import (
	"fmt"
	"testing"
)

func ExampleSet() {
	input := []int{1, 2, 3}
	s1 := NewSet(input...)
	res := s1.Reduce(func(k int, acc int) int {
		return acc + k
	}, 0)
	input[0] = 100 // underlying slice should not affect Set
	res2 := s1.Reduce(func(k int, acc int) int {
		return acc + k
	}, 0)
	fmt.Print(res, ";", res2)
	// Output: 6;6
}

func TestSet(t *testing.T) {
	input := []string{"A", "B", "C"}
	s1 := NewSet(input...)
	s1i := s1.Items()
	s1i[0] = "Z"
	s2 := NewSet(s1i...)
	if s1.Len() != 3 && s2.Len() != 4 {
		t.Errorf("lengths wrong - s1=%d, s2=%d", s1.Len(), s2.Len())
	}
}
