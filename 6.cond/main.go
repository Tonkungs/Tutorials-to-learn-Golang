package main

import "fmt"

// ตัวแปร
func main() {
	// ตัวเลข

	fmt.Println("Input Number :")
	var number int
	fmt.Scanf("%d", &number)

	// cond := number > 10
	// isMen := true

	// แบบเงือนไขเดียว
	// if cond {
	// 	fmt.Println(number * 5)
	// } else {
	// 	fmt.Println(number * 0)
	// }
	// if cond && isMen {
	// 	fmt.Println(number * 5)
	// } else {
	// 	fmt.Println(number * 0)
	// }

	// แบบ switch
	switch number {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("No")
	}
}
