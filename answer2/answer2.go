// Answer no. 2
// go run AnswerSheet2.go if you want to test the code i made

package main

import (
	"fmt"
	"strconv"
)

// singleDigit function used to check the missing number for 1 digit input
func singleDigit(number string) int {
	var answer, i, j int
	rangeDigit := 1
	var a []string
	// convert given string into slife of string
	for i, j = 0, 1; j <= len(number); i, j = i+rangeDigit, j+rangeDigit {
		// check if there will be more than 1 digit,
		// if there's 9 in the string, next string will be count as 2 digit
		if number[i:j] == "9" {
			a = append(a, number[i:j])
			j++
			rangeDigit++
		} else {
			a = append(a, number[i:j])
		}

	}
	// Convert the array/slice of string into an integer using strconv package
	var b []int
	for _, k := range a {
		l, err := strconv.Atoi(k)
		if err != nil {
			panic(err)
		}
		b = append(b, l)
	}
	// check the missing number, if b[n] don't match with m, then b[n] is the missing number and loop will break
	if len(b) >= 1 {
		for m, n := b[0], 0; n < len(b); m++ {
			if b[n] != m {
				fmt.Println(m, " is the missing number")
				n--
				answer = m
				break
			}
			n++
		}
	}
	return answer
}

// twoDigit function used to check the missing number for 2 digit input
func twoDigit(number string) int {
	var answer, i, j int
	rangeDigit := 2
	var a []string
	// convert given string into slife of string
	for i, j = 0, 2; j <= len(number); i, j = i+rangeDigit, j+rangeDigit {
		// check if there will be more than 2 digit,
		// if there's 99 in the string, next string will be count as 3 digit
		if number[i:j] == "99" {
			a = append(a, number[i:j])
			j++
			rangeDigit++
		} else {
			a = append(a, number[i:j])
		}

	}
	// Convert the array/slice of string into an integer using strconv package
	var b []int
	for _, k := range a {
		l, err := strconv.Atoi(k)
		if err != nil {
			panic(err)
		}
		b = append(b, l)
	}
	// check the missing number, if b[n] don't match with m, then b[n] is the missing number and loop will break
	if len(b) >= 1 {
		for m, n := b[0], 0; n < len(b); m++ {
			if b[n] != m {
				fmt.Println(m, " is the missing number")
				n--
				answer = m
				break
			}
			n++
		}
	}
	return answer
}

// threeDigit function used to check the missing number for 3 digit input
func threeDigit(number string) int {
	var answer, i, j int
	rangeDigit := 3
	var a []string
	// convert given string into slife of string
	for i, j = 0, 3; j <= len(number); i, j = i+rangeDigit, j+rangeDigit {
		// check if there will be more than 3 digit,
		// if there's 999 in the string, next string will be count as 4 digit
		if number[i:j] == "999" {
			a = append(a, number[i:j])
			j++
			rangeDigit++
		} else {
			a = append(a, number[i:j])
		}

	}
	// Convert the array/slice of string into an integer using strconv package
	var b []int
	for _, k := range a {
		l, err := strconv.Atoi(k)
		if err != nil {
			panic(err)
		}
		b = append(b, l)
	}
	// check the missing number, if b[n] don't match with m, then b[n] is the missing number and loop will break
	if len(b) >= 1 {
		for m, n := b[0], 0; n < len(b); m++ {
			if b[n] != m {
				fmt.Println(m, " is the missing number")
				n--
				answer = m
				break
			}
			n++
		}
	}
	return answer
}

// fourDigit function used to check the missing number for 3 digit input
func fourDigit(number string) int {
	var answer, i, j int
	rangeDigit := 4
	var a []string
	// convert given string into slife of string
	for i, j = 0, 4; j <= len(number); i, j = i+rangeDigit, j+rangeDigit {
		// check if there will be more than 4 digit,
		// if there's 9999 in the string, next string will be count as 5 digit
		if number[i:j] == "9999" {
			a = append(a, number[i:j])
			j++
			rangeDigit++
		} else {
			a = append(a, number[i:j])
		}

	}
	// Convert the array/slice of string into an integer using strconv package
	var b []int
	for _, k := range a {
		l, err := strconv.Atoi(k)
		if err != nil {
			panic(err)
		}
		b = append(b, l)
	}
	// check the missing number, if b[n] don't match with m, then b[n] is the missing number and loop will break
	if len(b) >= 1 {
		for m, n := b[0], 0; n < len(b); m++ {
			if b[n] != m {
				fmt.Println(m, " is the missing number")
				n--
				answer = m
				break
			}
			n++
		}
	}
	return answer
}

// fiveDigit function used to check the missing number for 3 digit input
func fiveDigit(number string) int {
	var answer, i, j int
	rangeDigit := 5
	var a []string
	// convert given string into slife of string
	for i, j = 0, 5; j <= len(number); i, j = i+rangeDigit, j+rangeDigit {
		// check if there will be more than 5 digit,
		// if there's 99999 in the string, next string will be count as 6 digit
		if number[i:j] == "99999" {
			a = append(a, number[i:j])
			j++
			rangeDigit++
		} else {
			a = append(a, number[i:j])
		}

	}
	// Convert the array/slice of string into an integer using strconv package
	var b []int
	for _, k := range a {
		l, err := strconv.Atoi(k)
		if err != nil {
			panic(err)
		}
		b = append(b, l)
	}
	// check the missing number, if b[n] don't match with m, then b[n] is the missing number and loop will break
	if len(b) >= 1 {
		for m, n := b[0], 0; n < len(b); m++ {
			if b[n] != m {
				fmt.Println(m, " is the missing number")
				n--
				answer = m
				break
			}
			n++
		}
	}
	return answer
}

// sixDigit function used to check the missing number for 3 digit input
func sixDigit(number string) int {
	var b []int
	var answer, i, j int
	//Split the input and put it into an array/slice of string
	rangeDigit := 6
	var a []string
	// convert given string into slife of string
	for i, j = 0, 6; j <= len(number); i, j = i+rangeDigit, j+rangeDigit {
		// check if there will be more than 6 digit,
		// if there's 999999 in the string, next string will be count as 7 digit
		if number[i:j] == "999999" {
			a = append(a, number[i:j])
			j++
			rangeDigit++
		} else {
			a = append(a, number[i:j])
		}

	}
	// Convert the array/slice of string into an integer using strconv package
	for _, k := range a {
		l, err := strconv.Atoi(k)
		if err != nil {
			panic(err)
		}
		b = append(b, l)
	}
	// check the missing number, if b[n] don't match with m, then b[n] is the missing number and loop will break
	if len(b) >= 1 {
		for m, n := b[0], 0; n < len(b); m++ {
			if b[n] != m {
				fmt.Println(m, " is the missing number")
				n--
				answer = m
				break
			}
			n++
		}
	}
	return answer
}

func main() {
	//number used for contain the input
	//a and b as a slice variable
	var number string
	var a []string
	var b []int
	fmt.Print("Enter the code between 1 - 1.000.000 : ")
	fmt.Scan(&number)

	//  check the number's digit
	if len(number) >= 18 {
		var a []string
		var b []int
		if number[0:6] == "999999" {
			a = append(a, number[0:6])
			a = append(a, number[6:13])

		} else {
			a = append(a, number[0:6])
			a = append(a, number[6:12])

		}

		for _, k := range a {
			l, err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}
			b = append(b, l)
		}
		if b[0]+1 == b[1] {
			sixDigit(number)
		} else if b[0]+2 == b[1] {
			// special occurence when the second number is the answer
			fmt.Println("missing number is = ", b[0]+1)
		} else {
			var a []string
			var b []int
			if number[0:5] == "99999" {
				a = append(a, number[0:5])
				a = append(a, number[5:11])
			} else {
				a = append(a, number[0:5])
				a = append(a, number[5:10])
			}
			for _, k := range a {
				l, err := strconv.Atoi(k)
				if err != nil {
					panic(err)
				}
				b = append(b, l)
			}
			if b[0]+1 == b[1] {
				fiveDigit(number)
			} else if b[0]+2 == b[1] {
				// special occurence when the second number is the answer
				fmt.Println("missing number is = ", b[0]+1)

			} else {
				var a []string
				var b []int
				if number[0:4] == "9999" {
					a = append(a, number[0:4])
					a = append(a, number[4:9])
				} else {
					a = append(a, number[0:4])
					a = append(a, number[4:8])
				}
				for _, k := range a {
					l, err := strconv.Atoi(k)
					if err != nil {
						panic(err)
					}
					b = append(b, l)
				}
				if b[0]+1 == b[1] {
					fourDigit(number)
				} else if b[0]+2 == b[1] {
					// special occurence when the second number is the answer
					fmt.Println("missing number is = ", b[0]+1)
				} else {
					var a []string
					var b []int
					if number[0:3] == "999" {
						a = append(a, number[0:3])
						a = append(a, number[3:7])

					} else {
						a = append(a, number[0:3])
						a = append(a, number[3:6])
					}
					for _, k := range a {
						l, err := strconv.Atoi(k)
						if err != nil {
							panic(err)
						}
						b = append(b, l)
					}
					if b[0]+1 == b[1] {
						threeDigit(number)
					} else if b[0]+2 == b[1] {
						// special occurence when the second number is the answer
						fmt.Println("missing number is = ", b[0]+1)
					} else {
						var a []string
						var b []int
						if number[0:2] == "99" {
							a = append(a, number[0:2])
							a = append(a, number[2:5])
						} else {
							a = append(a, number[0:2])
							a = append(a, number[2:4])
						}
						for _, k := range a {
							l, err := strconv.Atoi(k)
							if err != nil {
								panic(err)
							}
							b = append(b, l)
						}
						if b[0]+1 == b[1] {
							twoDigit(number)
						} else if b[0]+2 == b[1] {
							// special occurence when the second number is the answer
							fmt.Println("missing number is = ", b[0]+1)

						} else {
							var a []string
							var b []int
							if number[0:1] == "9" {
								a = append(a, number[0:1])
								a = append(a, number[1:3])
							} else {
								a = append(a, number[0:1])
								a = append(a, number[1:2])
							}
							for _, k := range a {
								l, err := strconv.Atoi(k)
								if err != nil {
									panic(err)
								}
								b = append(b, l)
							}
							if b[0]+1 == b[1] {
								singleDigit(number)
							} else if b[0]+2 == b[1] {
								// special occurence when the second number is the answer
								fmt.Println("missing number is = ", b[0]+1)

							} else {
								fmt.Println("Couldn't solve the problem")
							}
						}
					}
				}
			}
		}
	} else if len(number) >= 15 {
		//  check the number's digit
		a = append(a, number[0:5])
		a = append(a, number[5:10])
		for _, k := range a {
			l, err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}
			b = append(b, l)
		}
		if b[0]+1 == b[1] {
			fiveDigit(number)
		} else if b[0]+2 == b[1] {
			// special occurence when the second number is the answer
			fmt.Println("missing number is = ", b[0]+1)
		} else {
			var a []string
			var b []int
			if number[0:4] == "9999" {
				a = append(a, number[0:4])
				a = append(a, number[4:9])
			} else {
				a = append(a, number[0:4])
				a = append(a, number[4:8])
			}
			for _, k := range a {
				l, err := strconv.Atoi(k)
				if err != nil {
					panic(err)
				}
				b = append(b, l)
			}
			if b[0]+1 == b[1] {
				fourDigit(number)

			} else if b[0]+2 == b[1] {
				// special occurence when the second number is the answer
				fmt.Println("missing number is = ", b[0]+1)
			} else {
				var a []string
				var b []int
				if number[0:3] == "999" {
					a = append(a, number[0:3])
					a = append(a, number[3:7])
				} else {
					a = append(a, number[0:3])
					a = append(a, number[3:6])
				}
				for _, k := range a {
					l, err := strconv.Atoi(k)
					if err != nil {
						panic(err)
					}
					b = append(b, l)
				}
				if b[0]+1 == b[1] {
					threeDigit(number)

				} else if b[0]+2 == b[1] {
					// special occurence when the second number is the answer
					fmt.Println("missing number is = ", b[0]+1)
				} else {
					var a []string
					var b []int
					if number[0:2] == "99" {
						a = append(a, number[0:2])
						a = append(a, number[2:5])
					} else {
						a = append(a, number[0:2])
						a = append(a, number[2:4])
					}
					for _, k := range a {
						l, err := strconv.Atoi(k)
						if err != nil {
							panic(err)
						}
						b = append(b, l)
					}
					if b[0]+1 == b[1] {
						twoDigit(number)
					} else if b[0]+2 == b[1] {
						// special occurence when the second number is the answer
						fmt.Println("missing number is = ", b[0]+1)
					} else {
						var a []string
						var b []int
						if number[0:1] == "9" {
							a = append(a, number[0:1])
							a = append(a, number[1:3])
						} else {
							a = append(a, number[0:1])
							a = append(a, number[1:2])
						}
						for _, k := range a {
							l, err := strconv.Atoi(k)
							if err != nil {
								panic(err)
							}
							b = append(b, l)
						}
						if b[0]+1 == b[1] {
							singleDigit(number)
						} else if b[0]+2 == b[1] {
							// special occurence when the second number is the answer
							fmt.Println("missing number is = ", b[0]+1)
						} else {
							fmt.Println("couldn't solve the problem")
						}
					}
				}
			}
		}
	} else if len(number) >= 12 {
		//  check the number's digit
		a = append(a, number[0:4])
		a = append(a, number[4:8])
		for _, k := range a {
			l, err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}
			b = append(b, l)
		}
		if b[0]+1 == b[1] {
			fourDigit(number)

		} else if b[0]+2 == b[1] {
			// special occurence when the second number is the answer
			fmt.Println("missing number is = ", b[0]+1)
		} else {
			var a []string
			var b []int
			if number[0:3] == "999" {
				a = append(a, number[0:3])
				a = append(a, number[3:7])
			} else {
				a = append(a, number[0:3])
				a = append(a, number[3:6])
			}
			for _, k := range a {
				l, err := strconv.Atoi(k)
				if err != nil {
					panic(err)
				}
				b = append(b, l)
			}
			if b[0]+1 == b[1] {
				threeDigit(number)
			} else if b[0]+2 == b[1] {
				// special occurence when the second number is the answer
				fmt.Println("missing number is = ", b[0]+1)
			} else {
				var a []string
				var b []int
				if number[0:2] == "99" {
					a = append(a, number[0:2])
					a = append(a, number[2:5])
				} else {
					a = append(a, number[0:2])
					a = append(a, number[2:4])
				}
				for _, k := range a {
					l, err := strconv.Atoi(k)
					if err != nil {
						panic(err)
					}
					b = append(b, l)
				}
				if b[0]+1 == b[1] {
					twoDigit(number)
				} else if b[0]+2 == b[1] {
					// special occurence when the second number is the answer
					fmt.Println("missing number is = ", b[0]+1)
				} else {
					var a []string
					var b []int
					if number[0:1] == "9" {
						a = append(a, number[0:1])
						a = append(a, number[1:3])

					} else {
						a = append(a, number[0:1])
						a = append(a, number[1:2])

					}

					for _, k := range a {
						l, err := strconv.Atoi(k)
						if err != nil {
							panic(err)
						}
						b = append(b, l)
					}
					if b[0]+1 == b[1] {
						singleDigit(number)

					} else if b[0]+2 == b[1] {
						// special occurence when the second number is the answer
						fmt.Println("missing number is = ", b[0]+1)

					} else {
						fmt.Println("couldn't solve the problem")
					}
				}
			}

		}
	} else if len(number) >= 9 {
		//  check the number's digit
		a = append(a, number[0:3])
		a = append(a, number[3:6])

		for _, k := range a {
			l, err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}
			b = append(b, l)
		}

		if b[0]+1 == b[1] {
			threeDigit(number)

		} else if b[0]+2 == b[1] {
			// special occurence when the second number is the answer
			fmt.Println("missing number is = ", b[0]+1)

		} else {
			var a []string
			var b []int
			if number[0:2] == "99" {
				a = append(a, number[0:2])
				a = append(a, number[2:5])

			} else {
				a = append(a, number[0:2])
				a = append(a, number[2:4])

			}

			for _, k := range a {
				l, err := strconv.Atoi(k)
				if err != nil {
					panic(err)
				}
				b = append(b, l)
			}
			if b[0]+1 == b[1] {
				twoDigit(number)

			} else if b[0]+2 == b[1] {
				// special occurence when the second number is the answer
				fmt.Println("missing number is = ", b[0]+1)

			} else {
				var a []string
				var b []int
				if number[0:1] == "9" {
					a = append(a, number[0:1])
					a = append(a, number[1:3])

				} else {
					a = append(a, number[0:1])
					a = append(a, number[1:2])

				}

				for _, k := range a {
					l, err := strconv.Atoi(k)
					if err != nil {
						panic(err)
					}
					b = append(b, l)
				}
				if b[0]+1 == b[1] {
					singleDigit(number)

				} else if b[0]+2 == b[1] {
					// special occurence when the second number is the answer
					fmt.Println("missing number is = ", b[0]+1)

				} else {
					fmt.Println("couldn't solve the problem")
				}
			}
		}
	} else if len(number) >= 6 {
		//  check the number's digit
		a = append(a, number[0:2])
		a = append(a, number[2:4])

		for _, k := range a {
			l, err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}
			b = append(b, l)
		}

		if b[0]+1 == b[1] {
			twoDigit(number)

		} else if b[0]+2 == b[1] {
			// special occurence when the second number is the answer
			fmt.Println("missing number is = ", b[0]+1)

		} else {
			var a []string
			var b []int
			if number[0:1] == "9" {
				a = append(a, number[0:1])
				a = append(a, number[1:3])

			} else {
				a = append(a, number[0:1])
				a = append(a, number[1:2])

			}

			for _, k := range a {
				l, err := strconv.Atoi(k)
				if err != nil {
					panic(err)
				}
				b = append(b, l)
			}
			if b[0]+1 == b[1] {
				singleDigit(number)

			} else if b[0]+2 == b[1] {
				// special occurence when the second number is the answer
				fmt.Println("missing number is = ", b[0]+1)

			} else {
				fmt.Println("couldn't solve the problem")
			}
		}
	} else if len(number) >= 3 {
		//  check the number's digit
		a = append(a, number[0:1])
		a = append(a, number[1:2])

		for _, k := range a {
			l, err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}
			b = append(b, l)
		}

		if b[0]+1 == b[1] {
			singleDigit(number)
		} else if b[0]+2 == b[1] {
			// special occurence when the second number is the answer
			fmt.Println("missing number is = ", b[0]+1)

		} else {
			fmt.Println("couldn't solve the problem")
		}
	}
}
