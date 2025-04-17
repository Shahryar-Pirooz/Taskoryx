package fp

func Map[T, U any](s []T, mapFunction func(T) U) []U {
	result := make([]U, len(s))
	for index := range s {
		result[index] = mapFunction(s[index])
	}
	return result
}
