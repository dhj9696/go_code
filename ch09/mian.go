package main

import "fmt"

type MyWriter interface {
	Write(string) error
}
type MyCloser interface {
	Close() error
}
type writerCloser struct {
	MyWriter
}
type fileWriter struct {
	filePath string
}
type databaseWriter struct {
	host string
	port int
	db   string
}

func (fw *fileWriter) Write(string) error {
	fmt.Println("write string to file")
	return nil
}
func (dw *databaseWriter) Write(string) error {
	fmt.Println("write string to database")
	return nil
}

//	func (wc *writerCloser) Write(string) error {
//		fmt.Println("write string")
//		return nil
//	}
func (wc *writerCloser) Close() error {
	fmt.Println("close")
	return nil
}

func main() {
	var mw MyWriter = &writerCloser{
		&fileWriter{},
	}
	mw.Write("hello")
}
