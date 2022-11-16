# Annotation and special tags

## Annotation

Currently, we only have one annotation

```go
//@WinterAOP(template=<your proxy template>)
```

For the template, please refer to the example in this library.

## Tags for conditional compilation

For every source file including the annotation, you need to add the special comment at the head of the file:

```go
//go:build !winter_aop
// +build !winter_aop
```

the `wacli gen`	will create the code files like `xxx_winter_aop_gen.go which include `

```go
//go:build winter_aop
// +build winter_aop
```

If you use `go build`, the original source file will be compiled and only the business logic will be executed. 

If you use `go build -tags=winter_aop`	, the `xxx_winter_aop_gen.go`	will be compiled, both your business logic and `around function` will be executed.
