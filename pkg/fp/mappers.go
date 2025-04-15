package fp

func Map[T, U any](s []T, mapFunction func(T) U) []U {
	result := make([]U, len(s))
	for _, item := range s {
		result = append(result, mapFunction(item))
	}
	return result
}
