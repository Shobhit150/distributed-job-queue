package broker

import (
	"errors"
	"sync"
	
)

var ErrNotFound = errors.New("Job Not Found")

type memBroker struct {
	mu       sync.Mutex
	queue    []*Job
	inflight map[string]*Job
	store    map[string]*Job
}

func NewMemBroker() Broker {
	return &memBroker{
		queue:    make([]*Job, 0),
		inflight: make(map[string]*Job),
		store:    make(map[string]*Job),
	}
}

