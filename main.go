package main

import (
	"fmt"
	"github.com/jessjenkins/go-interfaces/doer"
	"github.com/jessjenkins/go-interfaces/things"
)

type Thing interface {
	Init()
}

func main()  {
	fmt.Println("Starting…")

	reader := getReader()
	writer:= getWriter()

	reader.Init()
	writer.Init()

	var w doer.Reader = reader
	
	doer := doer.NewDoer(w, writer)  // I know this fails, I just want to think about it
	doer.Do()
}

func getReader() Thing {
	return things.Reader{}
}

func getWriter() Thing {
	return things.Writer{}
}