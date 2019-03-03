package main

import "fmt"

// การสร้างฟั่งชั้นแบบไม่ต้องมีชื่อ
func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(11, 50))
}
