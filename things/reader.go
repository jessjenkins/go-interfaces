package things

import (
	"fmt"
	"github.com/jessjenkins/go-interfaces/doer"
)

type Reader struct {}

func (r Reader) Init() {
	fmt.Println("New Reader")
}

func (r Reader) Read() string {
	return "read value"
}

func (r Reader) Another() doer.Reader {
	return Reader{}
}