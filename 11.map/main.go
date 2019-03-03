package main

import "fmt"

func main() {
	// x := make(map[string]string)
	// x["TH"] = "Thai"
	// x["EN"] = "English"

	// สร้างแบบสั้น
	x := map[string]string{
		"TH": "Thai",
		"EN": "English",
	}

	fmt.Println(x)       // map[EN:English TH:Thai]
	fmt.Println(x["EN"]) // English
}
