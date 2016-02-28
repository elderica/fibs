# fibs

A golang port of https://rubygems.org/gems/fibs

As a command

```bash
$ go get github.com/sbwhitecap/fibs/cmd/fibs
$ fibs 10
1
1
2
3
5
8
13
21
34
55
```

As a library

```golang
package main
import (
	"github.com/sbwhitecap/fibs
	"fmt"
)

func main() {
	fmt.Println(fibs.New(1, 1).Take(10))
	// Output:
	// [1 1 2 3 6 8 13 21 34 55]

	fibs.New(1, 1).Each(func(x int) {
		// ...
	}
}
```
