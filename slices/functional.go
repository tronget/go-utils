package slices

func Map[T any, R any](sl []T, f func(T) R) []R {
	res := make([]R, len(sl))
	for i, el := range sl {
		res[i] = f(el)
	}
	return res
}

func Filter[T any](sl []T, f func(T) bool) []T {
	var res []T
	for _, el := range sl {
		if f(el) {
			res = append(res, el)
		}
	}
	return res
}

func Reduce[T any, R any](sl []T, f func(R, T) R, initial R) R {
	res := initial
	for _, el := range sl {
		res = f(res, el)
	}
	return res
}
