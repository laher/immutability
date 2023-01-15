package immutability

import (
	"fmt"
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	input := []string{"1", "2", "a", "b", "c"}
	l1 := NewList(input...)
	l2 := l1.Map(func(t string, i int) string {
		return fmt.Sprintf("n%d: %s", i, t)
	})
	expectedLen := 5
	if l1.Len() != expectedLen || l1.Len() != l2.Len() {
		t.Errorf("lengths wrong - expected %d. l1=%d, l2=%d", expectedLen, l1.Len(), l2.Len())
	}
	if !reflect.DeepEqual(l1.Items(), input) {
		t.Errorf("input [%v] should be same as returned items", input)
	}
	l2Expected := []string{"n0: 1", "n1: 2", "n2: a", "n3: b", "n4: c"}
	if !reflect.DeepEqual(l2.Items(), l2Expected) {
		t.Errorf("[%v] should be same as mapped items [%v]", l2Expected, l2.Items())
	}
	l3 := TransformList(l1, func(t string, i int) int {
		return i
	})
	l3Expected := []int{0, 1, 2, 3, 4}
	if !reflect.DeepEqual(l3.Items(), l3Expected) {
		t.Errorf("[%v] should be same as mapped items [%v]", l3Expected, l3.Items())
	}
}
