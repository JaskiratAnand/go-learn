package main

import "fmt"

func vals() (int, int) {
	return 3, 7
}

func sum(a, b, x, y int) (int, int) {
	return a + b, x + y
}

func main() {

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	d, e := sum(1, 2, 3, 4)
	fmt.Println(d, e)
}
