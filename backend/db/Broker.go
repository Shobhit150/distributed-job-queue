package db

import "context"

type Broker interface {
	Enqueue(ctx context.Context, j *Job) error
	Dequeue(ctx context.Context, consumer string, max int) error
	Ack(ctx context.Context, jobID string) error
	Nack(ctx context.Context, jobID string, requeue bool) error
}