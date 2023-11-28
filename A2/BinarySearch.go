package main

import "fmt"

/*
understanding : Binary search is a divide-and-conquer algorithm used to efficiently find a target value in a sorted array.
It repeatedly divides the search space in half, comparing the target to the middle element and eliminating half of the remaining elements with each iteration. This process continues until the target is found
Algo : until the low index is less than or equal to high index:
Calculate mid index.
If the mid element is the target, return its index.
If the target is less than the mid element, update the high index to mid - 1.
If the target is greater than the mid element, update the low index to mid + 1.
If the target is not found, return -1.
time complexity : O(log n)
test case : [1,2,3,4,5,6,7,9] target = 5 o/p= 3
*/

func main() {
	
	array := []int{1, 2, 3, 5, 6, 7, 9}
	
	target := 5
	
	index := binarySearch(array, target)
	
	if index != -1 {
		fmt.Printf("Target %d found at index %d\n", target, index)
	} else {
		fmt.Printf("Target %d not found in the array\n", target)
	}
}

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == target {
			return mid 
		} else if arr[mid] < target {
			low = mid + 1 
		} else {
			high = mid - 1 
		}
	}

	return -1 
}
