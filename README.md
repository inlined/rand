# Locked Source
Locked Source production-code companion to [xkcdrand](github.com/inlined/xkcdrand).

The `rand.Rand` type is a struct that is not inherently goroutine safe. The
top-level functions in `rand.Rand` are, thorough an unexported `rand.Source`
that locks on all operations.

This library allows you to improve testability with an injectable `rand.Rand`
and regain goroutine safety. When you don't need goroutine safety,
you can use `rand.NewSource`. This can speed up CPU-bound loops with random
number generation by removing locking.

## Usage

```golang
// Create a goroutine-safe rand.Rand with the default (fixed) seed.
// Seeds can be changed either with rand.Source.Seed or rand.Rand.Seed
rand.New(lockedsource.New())

// Wrap a goroutine-unsafe Source in a locking implementation:
rand.New(lockedsource.Wrap(mySource))
```

## Credit

This implementation, while trivial, is inspired directly by the lockedSource
type in [math/rand/rand.go](https://golang.org/src/math/rand/rand.go).