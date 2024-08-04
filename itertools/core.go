package itertools

import (
	"context"
	"sync"
)

func Map[T any, U any](function func(T) U, slice []T) []U {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	result := make([]U, 0, len(slice))
	var wg sync.WaitGroup
	wg.Add(len(slice))

	outCh := make(chan U)
	errorCh := make(chan error)

	for _, sl := range slice {
		go func(i T) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				defer func() {
					if r := recover(); r != nil {
						errorCh <- r.(error)
					}
				}()
				out := function(i)
				outCh <- out
			}
		}(sl)
	}

	go func() {
		wg.Wait()
		close(outCh)
	}()

	go func() {
		for out := range outCh {
			result = append(result, out)
		}
	}()

	wg.Wait()

	// Check for errors
	if len(errorCh) > 0 {
		panic(<-errorCh)
	}

	return result
}
