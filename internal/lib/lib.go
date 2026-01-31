package lib

func Filter[T any](items []T, fn func(item T) bool) []T {
	filteredItems := []T{}
	for _, value := range items {
		if fn(value) {
			filteredItems = append(filteredItems, value)
		}
	}
	return filteredItems
}
