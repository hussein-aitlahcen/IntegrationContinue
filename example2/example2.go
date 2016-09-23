package example2

import "fmt"

type Tuple struct {
	v1 int
	v2 int
}

func Sample(myArray []Tuple) int {
	output := 0
	for _, tuple := range myArray {
		output += tuple.v1 + tuple.v2
	}
	return output
}

func main() {
	fmt.Printf("Yo la famille\n")
}
