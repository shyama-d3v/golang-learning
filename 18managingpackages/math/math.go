package math

// PI (exported)
const PI = 3.14159

// Add (exported)
func Add(x, y int) int {
	return x + y
}

// substract (unexported)

func substract(a, b int) int {
	return a - b
}
