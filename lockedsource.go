package rand

import (
	"math/rand"
	"sync"
)

// lockedSource implements rand.Source64
type lockedSource struct {
	rand.Source
	m sync.Mutex
}

func (l *lockedSource) Int63() int64 {
	l.m.Lock()
	n := l.Source.Int63()
	l.m.Unlock()
	return n
}

func (l *lockedSource) Seed(seed int64) {
	l.m.Lock()
	l.Source.Seed(seed)
	l.m.Unlock()
}

func (l *lockedSource) Uint64() uint64 {
	l.m.Lock()
	n := l.Source.(rand.Source64).Uint64()
	l.m.Unlock()
	return n
}

// LockSource wrapps an existing rand.Source to make it goroutine safe.
func LockSource(s rand.Source) rand.Source {
	return &lockedSource{Source: s}
}
