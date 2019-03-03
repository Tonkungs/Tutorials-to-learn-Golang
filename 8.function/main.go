package main

import "fmt"

// ตัวแปร
func main() {
	fmt.Println(1)
	fmt.Print("getFn1() =")
	getFn1()

	fmt.Print("getFn2() =")
	getFn2("Koko")

	fmt.Print("getFn3() =")
	fmt.Println(getFn3())

	fmt.Print("getFn4() =")
	name := getFn4("Tonkung")
	fmt.Println(getFn4(name))

	fmt.Println("var", sumData(1, 5, 9, 5))

	fmt.Println("factorial", factorial(5))
}

// ไม่รับค่า ไม่คืนค่า
func getFn1() {
	fmt.Println(1)
}

// รับค่า ไม่คืนค่า
func getFn2(name string) {
	fmt.Println(name)
}

// ไม่รับค่า คืนค่า  ต้องบอกว่าเวลาส่งค่ากลับ ส่งกลับแบบไหน
func getFn3() string {
	return "tonkung"
}

// รับค่า คืนค่า
func getFn4(name string) string {
	return name
}

// Variadic Function
// รับค่าไม่จำกัดจำนวน
func sumData(nums ...int) int {
	var count int
	// ไม่ระบค่าเริ่มต้น
	for _, num := range nums {
		count += num
	}
	return count
}

// ฟั่งชั้นที่เรียกใช้งานตัวเอง
func factorial(num int) int {
	if num == 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}
}
