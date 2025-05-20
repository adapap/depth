package slicehelpers

// Any returns true if any element in the slice satisfies the predicate function.
func Any[T any](slice []T, predicate func(T) bool) bool {
	for _, v := range slice {
		if predicate(v) {
			return true
		}
	}
	return false
}
