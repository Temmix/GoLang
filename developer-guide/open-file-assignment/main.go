package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("The error is: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
