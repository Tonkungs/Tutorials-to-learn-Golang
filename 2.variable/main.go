package main

import "fmt"

// ตัวแปร
func main() {
	// bool = true, false
	// string = “Hello World”, “Thailand”
	// int = 0, 1000, -3000
	// float64 = 11.23456, 0.0014, -45.5678
	// array = []
	// 	Short format
	// ประเภทตัวแปร
	// แบบเต็ม string
	var displayName string = "Tonkung"
	// แบบสั้น
	displayName2 := "Tonkung"

	// ตัวเลข
	var money int = 100
	// float32, float64
	var money2 float32 = 100.5
	// ฺBoolearn
	var isMen bool = true
	var isMin bool = 12 < 80
	age1, age2 := 25, 80

	fmt.Println(displayName, displayName2, money, money2, isMen)
	fmt.Println(age1, age2, isMin)
}
