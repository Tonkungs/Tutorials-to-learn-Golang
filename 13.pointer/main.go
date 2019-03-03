package main

import "fmt"

// เอาไว้อ้างอิงตำแหน่งของ address ใน memory
func main() {
	a := 10
	fmt.Printf("value a %d\n", a)

	fmt.Printf("Address a %x\n ", &a)

	var t *int
	t = &a // ชี้ไปยัง address ที่ตัวแปร x เก็บอยู่

	fmt.Printf("Poiter a %x\n ", &t)
	fmt.Println(t)
	*t = 50
	fmt.Printf("Poiter a %d\n ", a)

}
