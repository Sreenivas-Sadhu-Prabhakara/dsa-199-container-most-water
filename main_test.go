package main

import "testing"

func TestMaxArea(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
		{[]int{5}, 0}, // one line holds nothing
		{[]int{}, 0},
		{[]int{0, 0}, 0},
	}
	for _, c := range cases {
		if got := MaxArea(c.in); got != c.want {
			t.Errorf("MaxArea(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
