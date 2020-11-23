package doer

import "fmt"

type Reader interface {
	Read() string
	Another() Reader
}

type Writer interface {
	Write()
}

type Doer struct {
	Writer Writer
	Reader Reader
}

func NewDoer(r,w interface{}) Doer {
    doer := Doer{}
    if v, ok := r.(Reader) ; ok {
    	doer.Reader = v
    	fmt.Println("Is a reader")
	}
	if v, ok := w.(Writer) ; ok {
		doer.Writer = v
		fmt.Println("Is a writer")
	}
	fmt.Printf("Doer: %v\n",doer)
    return doer
}

func (d Doer) Do() {
	value := d.Reader.Read()
	fmt.Println(value)
	d.Writer.Write()
}