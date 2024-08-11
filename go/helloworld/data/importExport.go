package data

import (
	"fmt"
	"math"
)

// In Go, a name is exported if it begins with a capital letter.
// For example, Pizza is an exported name, as is Pi, which is exported from the math package.
func ExportThisFunction() {
	// fmt.Println(math.pi)  //pi is undefined
	fmt.Println(math.Pi)
}

func notExportThisFunction() {
	// fmt.Println(math.pi)  //pi is undefined
	fmt.Println(math.Pi)
}
