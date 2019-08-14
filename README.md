# Rand
The built-in `math/rand` library is not very condusive to injection or testability:

1. The way the rand.Rand library computes its values differs widely per type (e.g.
   an Int vs Int31 grabs different halves of the rand.Source's 64-bit field). This
   makes test fakes very impractical.
2. Use of the global rand methods prohibit test injection, but the struct interface
   isn't goroutine safe.

To fix these, this rand ipmlementation can fully supplant `math/rand`. The library
performs three primary functions:

1. It introduces an interface version of `rand.Rand` to allow better mocking (see
   [xkcdrand](github.com/inlined/xkcdrand) for one such fake)
2. It allows this interface to be created inline without importing both `inlined/rand`
   and `math/rand` in the same file.
3. It allows the interface to be created with a goroutine-safe source, preserving
   the benefits of the package global `math/rand` methods.

The `Rand` interface implementations in this library are trivial wrappers around
`math/rand` to avoid any loss of entropy. Since the library does not require a
`rand.Source` in its constructors, callers should remember to always call `Rand.Seed`
in production code. This is similar to the `math/rand` package methods, which also
default to a seed of 1.

## Usage

```golang
// Create an interface version of math/rand.Rand using the default (fixed) seed.
// Seeds can be changed with Rand.Seed.
rand.New()

// Create an interface version of math/rand.Rand that is goroutine safe, unlike
// math/rand.New()
rand.NewLocked()
```

## Credit

This implementation, while trivial, is inspired directly by the lockedSource
type in [math/rand/rand.go](https://golang.org/src/math/rand/rand.go).