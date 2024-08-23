# VimeWorldGo

**VimeWorldGo** is a package for [Go](https://go.dev/) that provides a client to interact with the [VimeWorld API](https://api.vimeworld.com/).

## Installation

```sh
go get github.com/iwajezhgf/vimeworldgo
```

## Examples

```go
package main

import (
	"fmt"
	"log"

	"github.com/iwajezhgf/vimeworldgo"
)

func main() {
	client := vimeworldgo.NewClient(vimeworldgo.Options{})

	res, err := client.GetLeaderboardList()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res[0])
}
```

- [VimeWorldGo Examples](https://github.com/iwajezhgf/vimeworldgo/tree/master/examples) - a list of examples
