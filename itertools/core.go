package itertools

import (
	"sync"
)

func Map[T any, U any](function func(T) U, slice []T) *[]U {
	result := new([]U)
	var wg sync.WaitGroup

	outCh := make(chan U, len(slice))
	defer close(outCh)

	for _, sl := range slice {
		wg.Add(1)

		go func(i T) {
			defer wg.Done()
			outCh <- function(i)
		}(sl)
	}

	wg.Wait()

	wgReceiver := &sync.WaitGroup{}

	go func() {
		wgReceiver.Add(1)
		defer wgReceiver.Done()
		for out := range outCh {
			*result = append(*result, out)
		}
	}()

	wgReceiver.Wait()

	return result
}
