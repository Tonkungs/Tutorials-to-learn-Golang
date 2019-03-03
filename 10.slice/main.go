package main

import "fmt"

// slice สามารถเพิ่มข้อมูลได้เรื่อยๆ แต่ก็เก็บเหมือยอารย์
func main() {
	// สร้างจากฟั่งชั้น make
	// x := make([]int, 5)
	// อีกแบบ
	x2 := []int{1, 2, 3}
	// ต่อข้อมูล
	x3 := append(x2, 1, 2, 3)
	fmt.Println(x3) //[1 2 3 1 2 3]

	// เปลี่ยน อาเรย์เป็น slice
	arr := [5]float64{1, 2, 3, 4, 5}
	x := arr[0:4]
	fmt.Println(x) // [1 2 3 4]

	// การ copy slice
	a1 := []int{1, 2, 3}
	a2 := make([]int, 2)
	copy(a2, a1)
	fmt.Println(a2) // [1 2]
}
