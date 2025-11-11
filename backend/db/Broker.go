package db

import (
	"context"
	"time"
)

type Broker interface {
	Enqueue(ctx context.Context, j *Job) error
	Dequeue(ctx context.Context, consumer string, max int) error
	Ack(ctx context.Context, jobID string) error
	Nack(ctx context.Context, jobID string, requeue bool) error
	Inspect(ctx context.Context, jobID string) (*Job, error)
	Cancel(ctx context.Context, jobID string) error
	Requeue(ctx context.Context, jobID string, runAfter time.Time) error
	Close() error  
}