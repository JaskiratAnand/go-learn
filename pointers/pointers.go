package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func incr(x *int) {
	*x++
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

	var p *int
	k := 100
	p = &k
	fmt.Println("pointer:", *p)

	incr(&k)
	fmt.Println("incr:", k)
	incr(p)
	fmt.Println("incr:", k)

	var pointer *int

	if pointer == nil {
		fmt.Println("NullPointer is nil")
	} else {
		fmt.Println("NullPointer is not nil")
	}

}
