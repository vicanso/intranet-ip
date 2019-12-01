package main

import (
	"fmt"
	"net"

	intranetip "github.com/vicanso/intranet-ip"
)

func main() {
	fmt.Println(intranetip.Is(net.ParseIP("127.0.0.1")))
}
