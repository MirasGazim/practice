package factFib

func Factorial(num int)int{
	if num > 1 {
		return num * Factorial(num - 1)
	} else {
		return 1
	}

}