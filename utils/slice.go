package utils

func IntSliceAddition(sl []int) int {
	a := 0

	for _, i := range sl {
		a += i
	}

	return a
}
