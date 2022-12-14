# winter-aop

winter-aop is a Golang AOP library.

## Quick Start

1. Install command line tools

```shell
~ go install github.com/fengya90/winter-aop/wacli@latest
~ wacli
please run: wacli -h
```

2. source code

```go
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

```

3. compile and run

```
~ tree quickstart
quickstart
└── main.go
~ wacli gen -d quickstart
generate quickstart/main_winter_aop_gen.go
~ tree quickstart
quickstart
├── main.go
└── main_winter_aop_gen.go
~ cd quickstart
~  ./quickstart
start:say
hello
finish:say
```


## Document

see [documents](doc)

