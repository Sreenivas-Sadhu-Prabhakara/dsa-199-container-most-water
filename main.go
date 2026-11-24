package main

import "fmt"

// MaxArea returns the most water two lines can hold: width times the
// shorter line. Converging pointers always move the shorter side — the only
// move that can possibly improve the area.
func MaxArea(height []int) int {
	best := 0
	l, r := 0, len(height)-1
	for l < r {
		h := height[l]
		if height[r] < h {
			h = height[r]
		}
		if area := h * (r - l); area > best {
			best = area
		}
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return best
}

func main() {
	fmt.Println(MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
