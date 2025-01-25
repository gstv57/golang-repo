package main

func main() {
	println(sumAll(1, 2, 3, 4, 5))
}

// := is used to declare and assign values to variables
// =  is used to assign values to variables or update values
func split(a, b int) (res int, reminder int) {
	res = a / b
	reminder = a % b
	return res, reminder // could be return using naked return
}

// Higher order function
func sum(a int) func(int) int {
	// Closure or anonymous function
	return func(b int) int {
		return a + b
	}
}

// Variadic function (...) called ellipsis
func sumAll(nums ...int) int {
	var out int
	for _, num := range nums {
		out += num
	}
	return out
}
