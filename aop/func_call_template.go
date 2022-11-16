package aop

import (
	"context"
	"github.com/pkg/errors"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

type AroundFunc func(originFunc *OriginalFuncMetaInfo)

type aroundObj struct {
	priority uint64
	around   AroundFunc
}

type OriginalFuncMetaInfo struct {
	fnValue         reflect.Value
	inputs          []reflect.Value
	ret             []reflect.Value
	originalFnValue reflect.Value
	ext             map[string]interface{}
}

func (o *OriginalFuncMetaInfo) Run() {
	o.ret = o.fnValue.Call(o.inputs)
}

func (o *OriginalFuncMetaInfo) GetFuncName() string {
	name := runtime.FuncForPC(o.originalFnValue.Pointer()).Name()
	tmp := strings.Split(name, ".")
	name = tmp[len(tmp)-1]
	name = strings.TrimSuffix(name, "-fm")
	return name
}

func (o *OriginalFuncMetaInfo) GetExt() map[string]interface{} {
	return o.ext
}

func prettyInterface(i interface{}) interface{} {
	if bs, ok := i.([]byte); ok {
		l := strconv.FormatUint(uint64(len(bs)), 10)
		return "byte:len=" + l
	}
	return i
}

func (o *OriginalFuncMetaInfo) GetParams() []interface{} {
	params := []interface{}{}
	for _, i := range o.inputs {
		params = append(params, prettyInterface(i.Interface()))
	}
	return params
}

func (o *OriginalFuncMetaInfo) GetContext() (context.Context, error) {
	for _, i := range o.inputs {
		ctxI := i.Interface()
		if ctx, ok := ctxI.(context.Context); ok {
			return ctx, nil
		}
	}
	return nil, errors.New("internal error")
}

func (o *OriginalFuncMetaInfo) GetInput() []reflect.Value {
	return o.inputs
}

func (o *OriginalFuncMetaInfo) GetFnValue() reflect.Value {
	return o.fnValue
}

func (o *OriginalFuncMetaInfo) SetRet(ret []reflect.Value) {
	o.ret = ret
}

func (o *OriginalFuncMetaInfo) GetResult() []interface{} {
	result := []interface{}{}
	for _, i := range o.ret {
		result = append(result, prettyInterface(i.Interface()))
	}
	return result
}

func (o *OriginalFuncMetaInfo) HasError() bool {
	if o.ret != nil && len(o.ret) > 0 {
		lastV := o.ret[len(o.ret)-1]
		if lastV.Interface() != nil {
			if _, ok := lastV.Interface().(error); ok {
				return true
			}
		}
	}
	return false
}

type FuncCallTemplate struct {
	aroundObjs []*aroundObj
}

func (t *FuncCallTemplate) call(Around AroundFunc, fnValue reflect.Value, originFunc *OriginalFuncMetaInfo) (method interface{}) {
	realFn := func(inputs []reflect.Value) (ret []reflect.Value) {
		if Around == nil {
			return fnValue.Call(inputs)
		} else {
			originFunc.fnValue = fnValue
			originFunc.inputs = inputs
			Around(originFunc)
			return originFunc.ret
		}
	}
	return reflect.MakeFunc(fnValue.Type(), realFn).Interface()
}

func (t *FuncCallTemplate) AddAround(around AroundFunc, priority uint64) *FuncCallTemplate {
	if around != nil {
		obj := &aroundObj{
			around:   around,
			priority: priority,
		}
		t.aroundObjs = append(t.aroundObjs, obj)
		sort.SliceStable(t.aroundObjs, func(i, j int) bool {
			return t.aroundObjs[i].priority < t.aroundObjs[j].priority
		})
	}
	return t
}

func (t *FuncCallTemplate) Call(fn interface{}) (method interface{}) {
	method = fn
	originFunc := &OriginalFuncMetaInfo{
		originalFnValue: reflect.ValueOf(fn),
		ext:             nil,
	}
	for _, obj := range t.aroundObjs {
		fnValue := reflect.ValueOf(method)
		method = t.call(obj.around, fnValue, originFunc)
	}
	return
}

func (t *FuncCallTemplate) CallWithExt(fn interface{}, ext map[string]interface{}) (method interface{}) {
	method = fn
	originFunc := &OriginalFuncMetaInfo{
		originalFnValue: reflect.ValueOf(fn),
		ext:             ext,
	}
	for _, obj := range t.aroundObjs {
		fnValue := reflect.ValueOf(method)
		method = t.call(obj.around, fnValue, originFunc)
	}
	return
}
