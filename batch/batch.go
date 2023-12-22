package batch

// Func is a generic function type that processes batches of items.
// It accepts a slice of strings as input and returns a slice of type T.
// This allows for flexible processing functions that can return different types of slices.
type Func[T any] func(batch []string) []T

// Process is a generic function to process items in batches.
// It accepts a slice of strings, a batch size, and a processing function.
// The items slice is divided into batches of size batchSize and each batch is then processed
// using the provided processing function. The results are collected into a slice of type T and returned.
func Process[T any](items []string, batchSize int, process Func[T]) []T {
	var results []T
	for i := 0; i < len(items); i += batchSize {
		end := i + batchSize
		if end > len(items) {
			end = len(items)
		}
		batch := items[i:end]
		results = append(results, process(batch)...)
	}
	return results
}
