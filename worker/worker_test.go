package worker_test

import (
	"go2/worker"
	"math/rand"
	"testing"
)

func TestWorkerPoolRandom(t *testing.T) {
	for i := 0; i < 100; i++ {
		want := rand.Intn(1000)
		have := worker.WorkerPool(want)
		if have != want {
			t.Fatalf("Expected: %d, recieved: %d", want, have)
		}
	}
}
