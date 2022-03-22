package doc

import (
	"fmt"

	"github.com/alcor67/repo-Go-level-2-home-work/configuration"
)

func Example() {

	result, err := configuration.Load()
	fmt.Println(result.MyDubFile, err)
	/* Output:
	список дубликатов файлов:
	file1.txt
	file1.txt
	file1.txt <nil>
	*/
}
