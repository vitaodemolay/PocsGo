package main

import "fmt"

func main() {
	array := []int{60, 50, 95, 80, 70}
	fmt.Println("Unsorted array:", array)
	fmt.Println("Sorting array using Bubble Sort...")
	bubbleSort(&array)
	fmt.Println("Sorted array:", array)
}

func bubbleSort(array *[]int) {
	for i := 0; i < len(*array)-1; i++ {
		hasSwapped := false
		for j := 0; j < len(*array)-1-i; j++ {
			hasSwapped = swap(&(*array)[j], &(*array)[j+1])
		}
		if !hasSwapped {
			fmt.Println("Algorithm finished early. Idx=", i)
			break
		}
	}
}

func swap(i, j *int) bool {
	if *i > *j {
		*i, *j = *j, *i
		return true
	}
	return false
}
