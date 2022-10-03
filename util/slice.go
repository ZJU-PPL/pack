package util

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}


func Filter[T any](s []T, test func(T) bool) []T {
	var ret []T
	for _, x := range s {
		if test(x) {
			ret = append(ret, x)
		}
	}
	return ret
}