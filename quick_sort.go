package main

import (
	"fmt"
)

func main() {
	arr := []int{80, 20, 60, 70, 90, 30, 40, 50, 100, 2, 50, 6, 77, 30}
	quickSort(0, len(arr)-1, arr...)
	fmt.Println(arr)

}
func quickSort(low int, high int, a ...int) {
	if low < high {
		pi := partition(low, high, a...)
		quickSort(low, pi-1, a...)
		quickSort(pi+1, high, a...)
	}
}
func partition(low int, high int, a ...int) int {
	pivot := a[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if a[j] < pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[i+1], a[high] = a[high], a[i+1]
	return i + 1
}
