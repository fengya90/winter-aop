package main

import (
	"fmt"

	"github.com/fengya90/winter-aop/example/tutorial/dep"
)

func main() {
	fmt.Println(dep.AddThreeNumber(2, 3, 4))
	ts := &dep.TestStruct{}
	ts.SayAny("hehe")
	ts.SayHello()
}
