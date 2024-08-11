package main

import "fmt"

func loops() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println("i = ", i)
	// A classic initial/condition/after for loop.
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	// j is undefined here
	// fmt.Println("j = ", j)

	// Another way of accomplishing the basic “do this N times” iteration is range over an integer.
	for i := range 3 {
		fmt.Println("range", i)
	}
	// for without a condition will loop repeatedly until you break out of the loop or return from the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// You can also continue to the next iteration of the loop.
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
