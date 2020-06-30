// Answer no. 3
// Used palindrome function as an output.
// However, i put static input in minimal & maximal variable.
// the server will only show the output  of the given minimal & maximal

package webtest

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/test", answer)
	http.ListenAndServe(":8080", nil)
}
func answer(w http.ResponseWriter, r *http.Request) {
	var remainder, i, minimal, maximal int
	sum := 0
	// You can change minimal & maximal manually if you want to check the palindrome function.
	minimal = 1
	maximal = 50
	for i = minimal; i <= maximal; i++ {
		var reverse int = 0
		temp := i
		for {
			remainder = temp % 10
			reverse = reverse*10 + remainder
			temp = temp / 10
			if temp == 0 {
				break
			}
		}
		if reverse == i {
			sum++
		}
	}
	fmt.Fprintf(w, "there are %v palindrome between  %v and %v", sum, minimal, maximal)
}
