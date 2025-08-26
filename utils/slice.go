package utils

func Find[T any](slice []T, predicate func(T) bool) (T, bool) {
	var found T

	for _, t := range slice {
		if predicate(t) {
			return t, true
		}
	}

	return found, false

}

func Map[T any, R any](slice []T, fn func(T) R) []R {
	newSlice := make([]R, len(slice))

	for i, t := range slice {
		newSlice[i] = fn(t)
	}

	return newSlice
}
