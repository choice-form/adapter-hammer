package utils

func Filter[T []E, E any](t T, cb func(E) bool) T {
	list := make([]E, 0)
	for _, v := range t {
		if cb(v) {
			list = append(list, v)
		}
	}
	return list
}
