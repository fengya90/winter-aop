//go:build !winter_aop
// +build !winter_aop

package main

import (
	"fmt"
	"github.com/fengya90/winter-aop/aop"
)

var (
	proxy = &aop.FuncCallTemplate{}
)

func myAroundFunc(originFunc *aop.OriginalFuncMetaInfo) {
	// do something before executing the function
	fmt.Printf("start:%v\n", originFunc.GetFuncName())
	// execute the function
	originFunc.Run()
	// do something after executing the function
	fmt.Printf("finish:%v\n", originFunc.GetFuncName())
}

func init() {
	proxy.AddAround(myAroundFunc, 1)
}

//@WinterAOP(template=proxy)
func Say(msg string) {
	fmt.Println(msg)
}

func main() {
	Say("hello")
}
