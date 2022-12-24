package mapreduce

import (
	"sync"
)

func Map[inT, outT any](mapFunc func(inT) (outT, error), inCh <-chan inT) (<-chan outT, <-chan error) {

	outCh := make(chan outT)
	errCh := make(chan error, 100)

	go func() {
		defer func() {
			close(outCh)
			close(errCh)
		}()

		var wg sync.WaitGroup
		for in := range inCh {
			wg.Add(1)
			go func(in inT) {
				out, err := mapFunc(in)
				if err != nil {
					errCh <- err
					wg.Done()
					return
				}
				outCh <- out
				wg.Done()
			}(in)
		}
		wg.Wait()
		errCh <- nil
	}()

	return outCh, errCh
}
