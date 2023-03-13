package main

import "fmt"

func main() {
	a := []int{4, 2, 1, 3, 5}
	quickSort(a)
	fmt.Println(a)
}

func quickSort(arr []int) {
	quickSortProc(arr, 0, len(arr)-1)
}

func quickSortProc(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSortProc(arr, low, pi-1)
		quickSortProc(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}
