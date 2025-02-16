package main

import "fmt"

func main() {
	var r1, s int
	fmt.Scanf("%d %d", &r1, &s)
	r2 := (2 * s) - r1
	fmt.Print(r2)
}
