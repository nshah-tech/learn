package data

import "fmt"

var (
	Countries [10]string
	Slice     []int
	Codes     map[int]bool
)

func init() {
	Countries[0] = "USA"
	Countries[1] = "Canada"
	Countries[2] = "Mexico"
	Countries[3] = "Brazil"
	Countries[4] = "India"
	qty := len(Countries)
	fmt.Println("Countries Saved =", qty)
}
