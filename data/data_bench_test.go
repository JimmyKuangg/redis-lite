package data

import (
	"strconv"
	"testing"
)

func BenchmarkSet(b *testing.B) {
	db := NewDatabase()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		key := "key" + strconv.Itoa(i)
		db.Set(key, "value")
	}
}

func BenchmarkGet(b *testing.B) {
	db := NewDatabase()
	db.Set("key", "value")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = db.Get("key")
	}
}

// Benchmark concurrent reads/writes
func BenchmarkConcurrentSetGet(b *testing.B) {
	db := NewDatabase()

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := "key" + strconv.Itoa(i)
			if i%2 == 0 {
				db.Set(key, "val")
			} else {
				_, _ = db.Get(key)
			}
			i++
		}
	})
}
