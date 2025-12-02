package broker

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
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

func (m *memBroker) Enqueue(ctx context.Context, j *Job) (string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if j.ID == "" {
		j.ID = uuid.NewString()
	}
	if j.CreatedAt.IsZero() {
		j.CreatedAt = time.Now()
	}
	m.queue = append(m.queue, j)
	return j.ID, nil
}

func (m *memBroker) Dequeue(ctx context.Context, consumer string) ([]*Job, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if len(m.queue) == 0 {
		return nil, nil
	}

	now := time.Now()
	var ready []*Job
	newQueue := make([]*Job, 0, len(m.queue))
	for _,j := range m.queue {
		if !j.RunAfter.IsZero() && j.RunAfter.After(now) {
			newQueue = append(newQueue, j)
			continue
		}
		m.inflight[j.ID] = j
		ready = append(ready, j)
	}
	m.queue = newQueue
	return ready, nil
}

func (m *memBroker) Ack( ctx context.Context, jobID string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
}

func (m *memBroker) Nack(ctx context.Context, jobID string, requeue bool) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	j,ok := m.inflight[jobID]

	if !ok {
		return ErrNotFound
	}
	delete(m.inflight, jobID)

	if requeue {
		j.Attempts++
		j.RunAfter = time.Now().Add(time.Second * time.Duration(1 << uint(j.Attempts)))
		m.queue = append(m.queue, j)
	} else {
		delete(m.store, jobID)
	}
	return nil
} 
