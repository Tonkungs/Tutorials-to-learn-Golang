package main

import "fmt"

// ตัวแปรแบบเรียกใช้ได้ทุกคน
var gNum int = 300

// ตัวแปร
func main() {
	// ตัวเลข
	var lNumb int = 300

	fmt.Println(gNum)
	fmt.Println(lNumb)
	other()
}

func other() {
	fmt.Println(gNum)
	fmt.Println(lNumb)
}
