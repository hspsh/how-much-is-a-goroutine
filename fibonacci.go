package how_much_is_a_goroutine

func GetFibonacciByIndex(index int) int64 {
	if index == 0 {
		return 0
	}
	if index <= 2 {
		return 1
	}
	var result int64
	var n1, n2 int64 = 1, 1
	for i := 3; i <= index; i++ {
		result = n1 + n2
		n2, n1 = n1, result
	}
	return result
}
