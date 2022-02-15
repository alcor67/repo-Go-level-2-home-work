package doc

import (
	"fmt"

	"github.com/alcor67/repo-Go-level-2-home-work/projectcalculator/calculator"
)

func Example() {
	result, err := calculator.Calc("+", 300, 500)
	fmt.Println(result, err)
	// Output: 800 <nil>
}
