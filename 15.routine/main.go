package main

import "fmt"

//  ทำพร้อมกับฟั่งชั่นอื่นได้
func main() {
	go f(0)
	var input string
	fmt.Scanln(&input)
}

func f(n int) {
	for index := 0; index < 10; index++ {
		fmt.Println(n, ":", index)
	}
}
