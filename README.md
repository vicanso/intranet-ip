# intranet-ip

[![Build Status](https://github.com/vicanso/intranet-ip/workflows/Test/badge.svg)](https://github.com/vicanso/intranet-ip/actions)

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
