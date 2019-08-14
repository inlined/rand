package rand

import "math/rand"

// Rand is an injectable interface that matches the struct of math/rand.Rand
type Rand interface {
	ExpFloat64() float64
	Float32() float32
	Float64() float64
	Int() int
	Int31() int32
	Int31n(n int32) int32
	Int63() int64
	Int63n(n int64) int64
	Intn(n int) int
	NormFloat64() float64
	Perm(n int) []int
	Read(p []byte) (n int, err error)
	Seed(seed int64)
	Shuffle(n int, swap func(i, j int))
	Uint32() uint32
	Uint64() uint64
}

// New creates a new rand interface implementation based backed by the rand.Rand struct.
// The returned value will not be goroutine safe. To get a goroutine safe implementation,
// use NewLocked().
func New() Rand {
	// Note: All built-in Rand methods start with a seed of 1.
	// Production uses should call Seed with a realistic value (e.g. time.Now().Unix())
	return rand.New(rand.NewSource(1))
}

// NewLocked returns a goroutine safe implementation of the Rand interface.
func NewLocked() Rand {
	return rand.New(LockSource(rand.NewSource(1)))

}
