package main

import (
	"fmt"
	"math/cmplx"
)

var (
	Tobe bool = true
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
	 )

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, Tobe, Tobe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}