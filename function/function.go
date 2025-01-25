package function

// Split := is used to declare and assign values to variables
// =  is used to assign values to variables or update values
func Split(a, b int) (res int, reminder int) {
	res = a / b
	reminder = a % b
	return res, reminder // could be return using naked return
}

// Sum Higher order function
func Sum(a int) func(int) int {
	// Closure or anonymous function
	return func(b int) int {
		return a + b
	}
}

// SumAll Variadic function (...) called ellipsis
func SumAll(nums ...int) int {
	var out int
	for _, num := range nums {
		out += num
	}
	return out
}
