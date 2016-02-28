package main

import (
	"flag"
	"fmt"
	"github.com/sbwhitecap/fibs"
	"strconv"
)

func main() {
	flag.Parse()

	s := flag.Arg(0)
	if i, err := strconv.Atoi(s); err == nil {
		for _, e := range fibs.New(1, 1).Take(uint(i)) {
			fmt.Println(e)
		}
	} else {
		fmt.Println("invalid arg")
	}
}
