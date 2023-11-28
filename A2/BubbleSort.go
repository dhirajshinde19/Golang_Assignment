package main

import "fmt"

/*
understanding:Bubble Sort repeatedly compares and swaps adjacent elements, moving the largest elements towards the end, until the entire list is sorted.
algo : using two for loops, Outer loop controls the number of passes through the array. Inner loop performs comparisons and swaps.
time complexity : O(n^2)
test case : [64,25,12,22,11] sorted to [11,12,22,25,64]
*/

func main() {

	array := []int{64, 25, 12, 22, 11}

	bubbleSort(array)

	fmt.Println("Sorted array:", array)
}

func bubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {

				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}

		}
	}
}
