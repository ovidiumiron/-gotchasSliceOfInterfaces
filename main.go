package main

import (
	"fmt"

	"github.com/ovidiumiron/test_me/concrete"
	"github.com/ovidiumiron/test_me/inter"
)

func main() {
	cs := []concrete.Concrete{concrete.Concrete{}, concrete.Concrete{}}
	//adaptor
	x := make([]inter.Inter, len(cs))
	for idx, c := range cs {
		x[idx] = c
	}
	fmt.Printf("%d\n", inter.Func(x))
}
