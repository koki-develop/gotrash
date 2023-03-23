package util

func Contains[T comparable](ts []T, t T) bool {
	for _, r := range ts {
		if r == t {
			return true
		}
	}
	return false
}
