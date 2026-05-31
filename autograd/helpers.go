package autograd

type Value struct {
	Data float64
	Grad float64

	Prev     []*Value
	Backward func()
}
