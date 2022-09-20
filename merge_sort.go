package main

import "fmt"

func mergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	midIdx := len(array) / 2
	leftSorted := mergeSort(array[:midIdx])
	rightSorted := mergeSort(array[midIdx:])
	return merge(leftSorted, rightSorted)
}

func merge(arrayOne, arrayTwo []int) []int {
	merged := []int{}
	for len(arrayOne) > 0 && len(arrayTwo) > 0 {
		if arrayOne[0] < arrayTwo[0] {
			merged = append(merged, arrayOne[0])
			arrayOne = arrayOne[1:]
		} else {
			merged = append(merged, arrayTwo[0])
			arrayTwo = arrayTwo[1:]
		}
	}

	merged = append(merged, arrayOne...)
	merged = append(merged, arrayTwo...)

	return merged
}

func main() {
	array := []int{9, 5, 2, 4, 11, 8, 22}
	fmt.Println(mergeSort(array))
}

// output: [2 4 5 8 9 11 22]
