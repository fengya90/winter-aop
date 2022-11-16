//go:build !winter_aop
// +build !winter_aop

package dep

import "fmt"

type TestStruct struct {
}

func (t *TestStruct) SayHello() {
	fmt.Println("hello")
}

//@WinterAOP(template=callTemplate)
func (t *TestStruct) SayAny(s string) {
	fmt.Println(s)
}
