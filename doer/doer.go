package doer

import "fmt"

type Reader interface {
	Read() string
}

type Writer interface {
	Write() string
}

type Doer struct {
	Writer Writer
	Reader Reader
}

func NewDoer(r Reader, w Writer) Doer {
	return Doer{
		Writer: w,
		Reader: r,
	}
}

func (d Doer) Do() {
	value := d.Reader.Read()
	fmt.Println(value)
	d.Writer.Write()
}