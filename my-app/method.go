package main

type Value int

func (v Value) Add(n Value) Value {
	return v + n
}

func method() {
	v := Value(1)
	v = v.Add(2)
}
