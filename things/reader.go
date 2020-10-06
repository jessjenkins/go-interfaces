package things

import "fmt"

type Reader struct {}

func (r Reader) Init() {
	fmt.Println("New Reader")
}

func (r Reader) Read() string {
	return "read value"
}