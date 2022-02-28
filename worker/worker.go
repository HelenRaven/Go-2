package worker

func WorkerPool(i int) int {
	pool := make(chan struct{}, 1)
	res := make(chan int, 1)
	res <- 0

	for j := 0; j <= i; j++ {
		pool <- struct{}{}

		go func() {
			defer func() {
				<-pool
			}()
			r := <-res
			r++
			res <- r
		}()
	}
	r := <-res
	return r
}
