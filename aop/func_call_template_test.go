package aop

import (
	"encoding/json"
	"fmt"
	"testing"
)

var (
	template = &FuncCallTemplate{}
)

func add(a, b int) int {
	return a + b
}

type TestBean struct {
}

func (t *TestBean) mul(a, b, c int) int {
	return a * b * c
}

func myAround(originFunc *OriginalFuncMetaInfo) {
	// before
	fmt.Println(originFunc.GetFuncName())
	b, _ := json.Marshal(originFunc.GetParams())
	r, _ := json.Marshal(originFunc.GetResult())
	fmt.Println(string(b))
	fmt.Println(string(r))
	originFunc.Run()
	b, _ = json.Marshal(originFunc.GetParams())
	r, _ = json.Marshal(originFunc.GetResult())
	fmt.Println(string(b))
	fmt.Println(string(r))
}

func myAround2(originFunc *OriginalFuncMetaInfo) {
	// before
	fmt.Println(originFunc.GetFuncName())
	originFunc.Run()
}

func TestFuncCallMisc(t *testing.T) {
	template.AddAround(myAround, 3).AddAround(myAround2, 3)
	tb := &TestBean{}
	fmt.Println(template.Call(add).(func(int, int) int)(10, 20))
	fmt.Println(template.Call(tb.mul).(func(int, int, int) int)(2, 3, 4))
}
