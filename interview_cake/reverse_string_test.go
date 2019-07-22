package main

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		in       []string
		expected []string
	}{
		{[]string{}, []string{}},
		{[]string{"a"}, []string{"a"}},
		{[]string{"a", "b"}, []string{"b", "a"}},
		{[]string{"a", "b", "c"}, []string{"c", "b", "a"}},
		{[]string{"a", "b", "c", "d"}, []string{"d", "c", "b", "a"}},
	}

	for _, tt := range tests {
		result := reverseString(tt.in)
		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("should be %v instead %v", tt.expected, result)
		}
	}
}

func reverseString(list []string) []string {
	if len(list) == 0 || len(list) == 1 {
		return list
	}

	start := 0
	end := len(list) - 1

	for start < end {
		tmp := list[start]
		list[start] = list[end]
		list[end] = tmp

		start += 1
		end -= 1
	}

	return list
}