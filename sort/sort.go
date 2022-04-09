package main

import "fmt"

func main() {

	fmt.Printf("%v\n", quickSort([]int{5, 1, 3, 6, 7}))
	fmt.Printf("%v\n", quickSort([]int{5, 1, 3, 3, 6, 2, 7}))
}

func quickSort(sl []int) []int {
	if len(sl) < 2 {
		return sl
	}

	idx := len(sl) / 2
	pivot := sl[idx]

	var l, h []int

	for i, v := range sl {
		if i == idx {
			continue
		}

		if v < pivot {
			l = append(l, v)
		}

		if v >= pivot {
			h = append(h, v)
		}
	}
	x := []int{}
	x = append(x, quickSort(l)...)
	x = append(x, pivot)
	x = append(x, quickSort(h)...)
	return x
}
