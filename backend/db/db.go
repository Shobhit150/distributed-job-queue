package db

import "time"

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

