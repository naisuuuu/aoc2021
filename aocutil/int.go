package aocutil

func Max(ints ...int) int {
	m := ints[0]
	for _, i := range ints[1:] {
		if i > m {
			m = i
		}
	}
	return m
}

func Min(ints ...int) int {
	m := ints[0]
	for _, i := range ints[1:] {
		if i < m {
			m = i
		}
	}
	return m
}
