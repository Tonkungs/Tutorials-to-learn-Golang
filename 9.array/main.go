package main

import "fmt"

// ตัวแปร
func main() {
	// อาเรย์ คือจำนวนกลุ่มของข้อมูลที่แน่นอน ต้องบอกขนาด เพราะจะต้องจองไว้ก่อน
	// var x [5]int
	// x[4] = 9
	x := [5]int{1, 2, 3, 4, 5}
	var total int

	for _, num := range x {
		total += num
	}
	fmt.Println(total / (int(len(x))))
}
