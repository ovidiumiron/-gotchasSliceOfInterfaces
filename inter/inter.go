package inter

type Inter interface {
	Foo() int
}

func Func(bar []Inter) int {
	return 100 + bar[0].Foo()
}
