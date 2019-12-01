# intranet-ip

[![Build Status](https://img.shields.io/travis/vicanso/intranet_ip.svg?label=linux+build)](https://travis-ci.org/vicanso/intranet_ip)

Determine if the IP address is an intranet IP

```go
package main

import (
	"fmt"
	"net"

	intranetip "github.com/vicanso/intranet-ip"
)

func main() {
	fmt.Println(intranetip.Is(net.ParseIP("127.0.0.1")))
}
```
