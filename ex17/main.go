package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	index, status := binarySearch(2, arr)
	fmt.Println(index, status)
	index, status = binarySearch(10, arr)
	fmt.Println(index, status)
	index, status = binarySearch(-3, arr)
	fmt.Println(index, status)
}

func binarySearch(target int, arr []int) (int, bool) {
	start := 0
	end := len(arr)
	mid := (start + end) / 2
	for start != end {
		if arr[mid] < target {
			start = mid + 1
		} else if arr[mid] > target {
			end = mid
		} else {
			return mid, true
		}
		mid = (start + end) / 2
	}

	return mid, false
}
