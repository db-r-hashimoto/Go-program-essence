package test1

import "fmt"

func main() {
	var n int
	n = 1
	fmt.Println(n)

	var a = int
	x := 1
	y := 1.2

	// aはintなのでエラーになる
	// go build -buildvcs=false
	// # my-app
	// ./test1.go:10:10: int (type) is not an expression
	a = 1 + ((float64(x) + 2) * float64(y))
	fmt.Println(a)
}