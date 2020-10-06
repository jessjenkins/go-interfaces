package things

import "fmt"

type Writer struct {}

func (w Writer) Init() {
	fmt.Println("New Writer")
}

func (w Writer) Write() {
	fmt.Println("write value")
}