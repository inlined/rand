// Package rand allows developers to inject rand.Rand interfaces rather the
// math/rand.Rand struct. This allows for easier testing with two benefits:
// 
// 1. Developers can inject either a locking or unlocking Rand implementation (the built-in
//    rand locks global methods and does not lock rand.Rand struct methods.
// 2. The rand.Rand interface allows more straightforward faking. The built-in rand.Rand
//    struct has differemt mechanisms for extracting various kinds of ints from bit fields,
//    which can make fakes unnecessarily hard to create.
//    For one such fake, see [xkcdrand](github.com/inlined/xkcdrand).
package rand