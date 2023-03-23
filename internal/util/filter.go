package util

func Filter[T any](ts []T, f func(t T) bool) []T {
	var rtn []T

	for _, t := range ts {
		if f(t) {
			rtn = append(rtn, t)
		}
	}

	return rtn
}
