# distributed-job-queue

### ⚡️ High-Performance Distributed Job Queue System (Go + Kafka + PostgreSQL)

> **A production-grade, scalable job queue platform with real-time processing, reliable delivery, and modern cloud-native architecture.**

---

## 🚀 Overview

This project is a distributed job queue system designed to handle high-throughput, reliable background processing at scale—similar to infrastructure powering large tech companies. It leverages **Kafka** for message brokering, **Go** for concurrency, and **PostgreSQL** for persistence and analytics.

---

## 🏗️ Architecture

![Discributed_img](Distributive_system.png)

## Folder structure

<pre>
distributed-job-queue
  backend
    cmd
        main.go
    worker
    db
      db.go       //Structure of job here
      Broker.go   //Defined Interface
    kafka
      comsumer/
      producer/
    proto/
  docker-compose.yml
</pre>

## Job structure

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