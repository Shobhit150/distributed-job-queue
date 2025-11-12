package broker

import (
	"context"
	"time"
)

type Job struct {
	ID        string            `json:"id"`
	Type      string            `json:"type"`
	Payload   []byte            `json:"payload"`
	TenantID  string            `json:"tenant_id,omitempty"`
	Priority  int               `json:"priority,omitempty"`
	DedupKey  string            `json:"dedup_key,omitempty"`
	Meta      map[string]string `json:"meta,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
	RunAfter  time.Time         `json:"run_after,omitempty"`
	Attempts  int               `json:"attempts,omitempty"`
}


type Broker interface {
	Enqueue(ctx context.Context, j *Job) (string, error)
	Dequeue(ctx context.Context, consumer string) ([]*Job, error)
	Ack(ctx context.Context, jobID string) error
	Nack(ctx context.Context, jobID string, requeue bool) error
	// Inspect(ctx context.Context, jobID string) (*Job, error)
	// Cancel(ctx context.Context, jobID string) error
	// Requeue(ctx context.Context, jobID string, runAfter time.Time) error
	// Close() error  
}