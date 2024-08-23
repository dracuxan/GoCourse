package main

func factorial(n int64) int64 {
	if n <= 1 {
		return 1
	} else {
		var fact int64 = 1
		var i int64
		for i = 1; i <= n; i++ {
			fact = fact * i
		}
		return fact
	}
}
