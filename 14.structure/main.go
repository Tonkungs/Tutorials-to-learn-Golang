package main

import "fmt"

type Books struct {
	title  string
	author string
	lang   string
	price  float64
}

//structure
func main() {
	// แบบทั่วไป
	// var GoLang Books
	// GoLang.title = "GO Lang"
	// GoLang.author = "Tonkung"
	// GoLang.lang = "Th"
	// GoLang.price = 1000

	GoLang := Books{
		title:  "GO Lang",
		author: "Tonkung",
		lang:   "Th",
		price:  1000,
	}

	fmt.Println(GoLang)
}
