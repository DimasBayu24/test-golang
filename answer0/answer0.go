// Answer no. 0
// go run AnswerSheet0.go if you want to test the code i made

package main

import "fmt"

func main() {

	//remainder is used tp store the leftover of the arithmetic result
	//i for loop
	//minimal and maximal to store the requirement input
	var remainder, i, minimal, maximal int
	sum := 0

	// 2 input for minimum range of number and its maximum
	// using fmt package command 'scan'
	fmt.Print("Enter any positive integer for minimal range : ")
	fmt.Scan(&minimal)
	fmt.Print("Enter any positive integer for maximal range : ")
	fmt.Scan(&maximal)

	//Then use loop to get every number in the given range include
	//to check they're palindrome or not
	for i = minimal; i <= maximal; i++ {
		var reverse int = 0
		temp := i

		//this  loop for  check the palindrome by using a simple arithmetic
		for {
			remainder = temp % 10
			reverse = reverse*10 + remainder
			temp = temp / 10

			//if temp variable equal 0, the loop stop
			if temp == 0 {
				break
			}
		}

		// if the reverse equal with the current looping number, that it is palindrome
		// and it will accumulated as +1 to sum variable
		if reverse == i {
			sum++
		}
	}

	//fmt.Print so it could show the number of how much palindrome in the given range of number
	fmt.Printf("there are %d palindrome\n", sum)
}
