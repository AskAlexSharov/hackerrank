package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf *bytes.Buffer

	fmt.Println(buf)
	fmt.Printf("%t\n", buf)
	if buf != nil {
		fmt.Println("NOT NIL!!!!")
	}
}
