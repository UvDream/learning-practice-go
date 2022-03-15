package main

import (
	"fmt"
	"os"
	"path/filepath"
)
func main() {
	a := false
	fmt.Printf("%T", a)
	filepath.Walk("/Users/wzj/Desktop/wzj/golang/learning-practice-go", getFile)
}

func getFile(path string, f os.FileInfo, err error) error {
	fmt.Println(path,f,err)
	return nil
}
