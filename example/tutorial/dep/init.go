package dep

import (
	"fmt"
	"time"

	"github.com/fengya90/winter-aop/aop"
)

var (
	callTemplate = &aop.FuncCallTemplate{}
)

func monitor(originFunc *aop.OriginalFuncMetaInfo) {
	// some code for monitoring
	defer func(begin time.Time) {
		// print the latency
		cost := time.Now().Sub(begin).Milliseconds()
		fmt.Printf("cost:%v\n", cost)

		// print the params
		fmt.Print("parameters: ")
		for _, param := range originFunc.GetParams() {
			fmt.Printf("%v\t", param)
		}
		fmt.Println()

		// print the result
		fmt.Print("result: ")
		for _, r := range originFunc.GetResult() {
			fmt.Printf("%v\t", r)
		}
		fmt.Println()
	}(time.Now())
	// do something before executing the function
	fmt.Printf("monitor start:%v\n", originFunc.GetFuncName())
	// execute the function
	originFunc.Run()
	// do something after executing the function
	fmt.Printf("monitor finish:%v\n", originFunc.GetFuncName())
}
func highPriorityAround(originFunc *aop.OriginalFuncMetaInfo) {
	// do something before executing the function
	fmt.Printf("highPriorityAround start:%v\n", originFunc.GetFuncName())
	// execute the function
	originFunc.Run()
	// do something after executing the function
	fmt.Printf("highPriorityAround finish:%v\n", originFunc.GetFuncName())
}

func init() {
	callTemplate.AddAround(monitor, 1)
	callTemplate.AddAround(highPriorityAround, 99)
}
