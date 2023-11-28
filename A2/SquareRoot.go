package main

import (
	"fmt"
)

/*
Understanding: given a number write a program to find its square root.  we Cannot calculate square root of negative number
algo : we use newton raphson method
time complexity : O(n)
test cases: 25 = 5, 9=3, 1784 = 42
*/

func main() {
	var num int
	fmt.Println("Enter the number to find the square root: ")
	fmt.Scanln(&num)

	if num < 0 {
		fmt.Println("Cannot calculate square root of negative number")
	} else {

		sqrt := num / 2
		temp := 0

		for temp != sqrt {
			//by newton raphson method
			temp = sqrt
			sqrt = (temp + (num / temp)) / 2
		}

		fmt.Println("square root of ", num, " is: ", sqrt)
	}
	// fmt.Printf("Square root of %d is %d: ",num,sqrt)
}
