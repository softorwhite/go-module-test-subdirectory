package printer

import "fmt"

type Printer struct{}

func (*Printer) Print(str string) {
	fmt.Println(str)
}
