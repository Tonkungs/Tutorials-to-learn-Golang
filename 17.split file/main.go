package main

import (
	"fmt"
	"lib/book"
)

func main() {
	// ไฟล์อยู่ที่ C:\Users\Tonkung\go\src\lib\book
	book := book.New("Node")
	fmt.Println("book: ", book)
}

// https://www.youtube.com/watch?v=ePdZwye1JaQ&list=PL9GjjPCrE7xng2Llw6Gcxa2F5XeDLOKRL&index=4
// 04 Golang กับ Type Casting
