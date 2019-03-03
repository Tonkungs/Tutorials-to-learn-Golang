package main

import "fmt"
import s "strings"

var p = fmt.Println

// ตัวแปร
func main() {
	// ตัวเลข
	numb1 := 300
	numb2 := 200
	fmt.Println("+ = ", numb1+numb2)
	fmt.Println("- = ", numb1-numb2)
	fmt.Println("* = ", numb1*numb2)
	fmt.Println("/ = ", numb1/numb2)
	fmt.Println("% = ", numb1%numb2)
	fmt.Println("----------------")
	// ตัวอักษร

	text1 := "Hello"
	text2 := "world"

	// ต่อตัวอักษร
	fmt.Println(text1 + text2)
	// เข้าถึงทีละตัว เริ่มตัวที่ 1 ถึงก่อนตัวที่ 3
	fmt.Println(text1[1:3]) // el
	fmt.Println(text1[1:])  // ello

	p("Contains: ", s.Contains("test", "es"))       //true
	p("Count: ", s.Count("test", "t"))              // 2
	p("HasPrefix: ", s.HasPrefix("test", "te"))     // true
	p("HasSuffix: ", s.HasSuffix("test", "st"))     // true
	p("Index: ", s.Index("test", "st"))             // 2
	p("Join: ", s.Join([]string{"a", "b"}, "test")) //atestb
	p("Repeat: ", s.Repeat("test", 5))              //testtesttesttesttest
	p("Replace: ", s.Replace("foo", "o", "0", -1))  //f00
	p("Replace: ", s.Replace("fooo", "o", "0", 2))  // f00o  ถ้าพบ 2 ตัวก็แทน 2 ตัวเลย
	p("Split: ", s.Split("a-b-c", "-"))             // [a b c]
	p("ToLower: ", s.ToLower("test"))               // test
	p("ToUpper: ", s.ToUpper("test"))               // TEST
	p("len: ", len("test"))                         // 4
}
