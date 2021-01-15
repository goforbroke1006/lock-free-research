package tcp_counter_lock_free

import (
	"github.com/goforbroke1006/lock-free-research/internal"
	"sync/atomic"
)

func NewLockFreeNumberCounter() *lockFreeNumberCounter {
	return &lockFreeNumberCounter{}
}

var _ internal.NumbersCounter = &lockFreeNumberCounter{}

type lockFreeNumberCounter struct {
	storage int64
}

func (nc *lockFreeNumberCounter) Add(n int64) error {
	for {
		old := nc.storage
		fresh := old + n
		if atomic.CompareAndSwapInt64(&nc.storage, old, fresh) {
			break
		}

		// TODO: create live-lock fix
	}

	return nil
}

func (nc *lockFreeNumberCounter) Get() (int64, error) {
	return nc.storage, nil
}
