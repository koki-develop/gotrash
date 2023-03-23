package util

func Some[T any](ts []T, f func(t T) bool) bool {
	for _, t := range ts {
		if f(t) {
			return true
		}
	}
	return false
}
