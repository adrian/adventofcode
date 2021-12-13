package utils

func Min(values []int) int {
	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}

func Max(values []int) int {
	max := values[0]
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
}

func SumNumbersUpTo(n int) int {
	return (n * (n - 1)) / 2
}
