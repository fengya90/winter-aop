//go:build !winter_aop
// +build !winter_aop

package dep

func add(a, b int) int {
	return a + b
}

//@WinterAOP(template=callTemplate)
func AddThreeNumber(a, b, c int) int {
	return add(a, add(b, c))
}
