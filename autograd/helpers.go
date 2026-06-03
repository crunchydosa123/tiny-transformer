package autograd

import "math"

type Value struct {
	Data float64
	Grad float64

	Prev     []*Value
	Backward func()
}

func Add(a, b *Value) *Value {
	out := &Value{
		Data: a.Data + b.Data,
		Prev: []*Value{a, b},
	}

	out.Backward = func() {
		a.Grad += out.Grad
		b.Grad += out.Grad
	}

	return out
}

func Mul(a, b *Value) *Value {
	out := &Value{
		Data: a.Data * b.Data,
		Prev: []*Value{a, b},
	}

	out.Backward = func() {
		a.Grad += b.Data * out.Grad
		b.Grad += a.Data * out.Grad
	}

	return out
}

func topo(
	v *Value,
	visited map[*Value]bool,
	order *[]*Value,
) {
	if visited[v] {
		return
	}

	visited[v] = true

	for _, child := range v.Prev {
		topo(child, visited, order)
	}

	*order = append(*order, v)
}

func (v *Value) BackwardAll() {
	var order []*Value

	topo(v, map[*Value]bool{}, &order)

	v.Grad = 1.0

	for i := len(order) - 1; i >= 0; i-- {
		if order[i].Backward != nil {
			order[i].Backward()
		}
	}
}

func ReLU(v *Value) *Value {
	val := math.Max(0, v.Data)

	out := &Value{
		Data: val,
		Prev: []*Value{v},
	}

	out.Backward = func() {
		if out.Data > 0 {
			v.Grad += out.Grad
		}
	}

	return out
}

func Pow(v *Value, p float64) *Value {
	out := &Value{
		Data: math.Pow(v.Data, p),
		Prev: []*Value{v},
	}

	out.Backward = func() {
		v.Grad += p *
			math.Pow(v.Data, p-1) *
			out.Grad
	}

	return out
}
