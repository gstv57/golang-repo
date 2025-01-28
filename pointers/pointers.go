package pointers

import "fmt"

func main() {
	var x *int

	fmt.Println(x)
}

func take(x *int) {
	*x = 100
}
