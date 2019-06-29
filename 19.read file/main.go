package main

import (
	"fmt"
	"io/ioutil"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main()  {
	// Read String
	txt, err := ioutil.ReadFile("./test-text.txt")
	check(err)
	// [84 69 83 84 32 82 69 65 68 32 84 88 84 32 70 73 76 69]
	fmt.Println("txt =",txt)
	// TEST READ TXT FILE
	fmt.Println("string(txt) =",string(txt))
	
}