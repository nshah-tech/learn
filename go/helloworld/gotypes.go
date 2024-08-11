// bool
//
// string
//
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
//
// byte // alias for uint8
//
// rune // alias for int32
//      // represents a Unicode code point
//
// float32 float64
//
// complex64 complex128
//
//
// Collections
// Array
// [5]int
//
// Slices
// []int
//
//Maps: key/value pairs
//map[keyType]valueType
//
// Generics 1.18
//

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

var message1231 string

func mainGoTypes() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
