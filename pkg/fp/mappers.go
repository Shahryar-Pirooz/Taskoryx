package fp

func Map[T, U any](s []T, mapFunction func(T) *U) []U {
	result := make([]U, len(s))
	for _, ptr := range s {
		result = append(result, *mapFunction(ptr))
	}
	return result
}
