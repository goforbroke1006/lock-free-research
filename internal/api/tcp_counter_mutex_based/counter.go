package tcp_counter_mutex_based

import (
	"github.com/goforbroke1006/lock-free-research/internal"
	"sync"
)

func NewMutexBasedNumberCounter() *mutexBasedNumberCounter {
	return &mutexBasedNumberCounter{}
}

var _ internal.NumbersCounter = &mutexBasedNumberCounter{}

type mutexBasedNumberCounter struct {
	storage   int64
	storageMx sync.RWMutex
}

func (nc *mutexBasedNumberCounter) Add(n int64) error {
	nc.storageMx.Lock()
	defer nc.storageMx.Unlock()

	nc.storage += n

	return nil
}

func (nc *mutexBasedNumberCounter) Get() (int64, error) {
	nc.storageMx.RLock()
	defer nc.storageMx.RUnlock()

	return nc.storage, nil
}
