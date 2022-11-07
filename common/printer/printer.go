package printer

import (
	"fmt"

	"github.com/softorwhite/go-module-test-subdirectory/common/printer/utility"
)

type Printer struct{}

func (*Printer) Print(str string) {
	fmt.Println(utility.StringGenerator())
}
