package main

import "fmt"

type MyWriter interface {
	Write(string)
}
type MyReader interface {
	Read() string
}
type MyReadWriter interface {
	MyReader
	MyWriter
	ReadWriter()
}
type SreadWriter struct{}

func (s SreadWriter) Read() string {
	//TODO implement me
	fmt.Println("read")
	return ""
}

func (s SreadWriter) Write(s2 string) {
	//TODO implement me
	fmt.Println("write")
}

func (s SreadWriter) ReadWriter() {
	//TODO implement me
	fmt.Println("read writer")
}

func main() {
	var mrw MyReadWriter = &SreadWriter{}
	mrw.Read()
}
