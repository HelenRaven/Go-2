package mutex_test

import (
	mu "go2/mutex"
	"testing"
)

func Benchmark1090(b *testing.B) {
	var set = mu.NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Add(1, "Apple")
			}
		})
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 9; i++ {
					set.Has(i)
				}
			}
		})
	})
}

func BenchmarkRW1090(b *testing.B) {
	var setRw = mu.NewSetRW()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				setRw.Add(1, "Apple")
			}
		})
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 9; i++ {
					setRw.Has(i)
				}
			}
		})
	})
}

func Benchmark5050(b *testing.B) {
	var set = mu.NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 5; i++ {
					set.Add(i, "Apple")
				}
			}
		})
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 5; i++ {
					set.Has(i)
				}
			}
		})
	})
}

func BenchmarkRW5050(b *testing.B) {
	var setRW = mu.NewSetRW()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 5; i++ {
					setRW.Add(i, "Apple")
				}
			}
		})
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 5; i++ {
					setRW.Has(i)
				}
			}
		})
	})
}

func Benchmark9010(b *testing.B) {
	var set = mu.NewSet()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 9; i++ {
					set.Add(i, "Apple")
				}
			}
		})
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				set.Has(1)
			}
		})
	})
}

func BenchmarkRW9010(b *testing.B) {
	var setRW = mu.NewSetRW()

	b.Run("", func(b *testing.B) {
		b.SetParallelism(1000)
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				for i := 0; i < 9; i++ {
					setRW.Add(i, "Apple")
				}
			}
		})
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				setRW.Has(1)
			}
		})
	})
}
