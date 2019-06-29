package main

import (
	"fmt"

	"github.com/google/gopacket/pcap"
)

func main() {
	version := pcap.Version()
	fmt.Println(version)
	book := say.New("Node")
	fmt.Println("Say", book)
}

// http://www.somkiat.cc/hello-module-with-go/
// การจัดการ go module
// http://www.somkiat.cc/go-module-without-vcs/ Go module :: มาลองสร้าง module ใช้เองแบบ Local

// go mod init goelster
// go mod tidy
// go mod vendor
// https://medium.com/acoshift/go-mod-summary-fd8a41ef58a8
// https:// https://medium.com/@kaweel/เมื่อฉันได้รู้จัก-go-module-e7f0aed2d528
// https://prongbang.github.io/golang/2018/10/20/goodbye-gopath.html
// https://medium.com/odds-team/go-mod-init-3347ffc980a4  
// การใช้ vender
// สอน go lang
// https://www.youtube.com/watch?v=fj7xFW0rSHc


//https://github.com/avelino/awesome-go