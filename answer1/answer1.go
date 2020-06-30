// Answer no. 1
// go run AnswerSheet1.go if you want to test the code i made

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	//variable below used to store the  slice of book's list
	var science, arts, general, geohistory, language, literature, philosophy, religion, math, social, newField []string
	//while k used for loop variable
	var k, l, m, n int

	//write number of book's list that's going to be arranged.
	// bufio and stdin so it can scan the input even after the whitespace
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter list of books: ")
	text, _ := reader.ReadString('\n')

	// text variable being convert into slice by using this strings.field
	field := strings.Fields(text)

	// sort the height of the books first
	sort.Slice(field, func(i, j int) bool {
		return field[i][2:4] > field[j][2:4]
	})
	// sort the field if there are more than 3 books, the shortest book will not included.
	// Otherwise, put it  to newField slice
	for l = len(field) - 1; l > 0; l-- {
		n = 0
		for m = l; m > 0; m-- {
			if field[l][0:2] == field[m][0:2] {
				n++
			}
		}
		if n < 3 {
			newField = append(newField, field[l])
		}
	}
	// sort the books once more
	sort.Slice(newField, func(i, j int) bool {
		return newField[i][2:4] > newField[j][2:4]
	})

	//and put the books to their own category slice.
	// Using loop and if the number of first character  of the book match with the category
	//it will stored into its category
	for k = 0; k < len(newField); k++ {
		if newField[k][:1] == "6" {
			science = append(science, newField[k])
		}
		if newField[k][:1] == "7" {
			arts = append(arts, newField[k])
		}
		if newField[k][:1] == "0" {
			general = append(general, newField[k])
		}
		if newField[k][:1] == "9" {
			geohistory = append(geohistory, newField[k])
		}
		if newField[k][:1] == "4" {
			language = append(language, newField[k])
		}
		if newField[k][:1] == "8" {
			literature = append(literature, newField[k])
		}
		if newField[k][:1] == "1" {
			philosophy = append(philosophy, newField[k])
		}
		if newField[k][:1] == "2" {
			religion = append(religion, newField[k])
		}
		if newField[k][:1] == "5" {
			math = append(math, newField[k])
		}
		if newField[k][:1] == "3" {
			social = append(social, newField[k])
		}
	}

	//with package fmt, show the list of books by its category first.
	fmt.Print(strings.Join(science, "  "), " ", strings.Join(arts, "  "),
		" ", strings.Join(general, "  "), " ", strings.Join(geohistory, "  "),
		" ", strings.Join(language, "  "), " ", strings.Join(literature, "  "),
		" ", strings.Join(philosophy, "  "), " ", strings.Join(religion, "  "),
		" ", strings.Join(math, "  "), " ", strings.Join(social, "  "))

}
