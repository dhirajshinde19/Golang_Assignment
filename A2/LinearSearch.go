package main

import "fmt"

/*
understanding: given an integer array and target we have to iteratively find if the target exists into that array and if yes then
what's the index of the target
algo : using a for loop, compare each element of the array with the target
time complexity : O(n)
test cases : [5,2,7,1,9,3,6] target = 9 o/p= 4
*/

func main() {
	// Example array
	array := []int{5, 2, 7, 1, 9, 3, 6}

	target := 9

	index := linearSearch(array, target)

	if index != -1 {
		fmt.Printf("Target %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Target %d not found in the array\n", target)
	}
}

func linearSearch(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}
