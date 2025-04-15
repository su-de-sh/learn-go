package main

import (
    "context"
    "sync"
)

// Task represents a unit of work that can be executed by a worker
type Task interface {
    Execute() Result
}

// Result represents the result of a task execution
type Result struct {
    Value interface{}
    Error error
}

// WorkerPool represents a pool of worker goroutines
type WorkerPool struct {
    numWorkers int
    tasks      chan Task
    results    chan Result
    done       chan struct{}
    wg         sync.WaitGroup
    ctx        context.Context
    cancel     context.CancelFunc
}

// NewWorkerPool creates a new worker pool with the specified number of workers
func NewWorkerPool(numWorkers int) *WorkerPool {
    ctx, cancel := context.WithCancel(context.Background())
    return &WorkerPool{
        numWorkers: numWorkers,
        tasks:      make(chan Task, numWorkers*2), // Buffered channel
        results:    make(chan Result, numWorkers*2), // Buffered channel
        done:       make(chan struct{}),
        ctx:        ctx,
        cancel:     cancel,
    }
}

// Start starts the worker pool
func (p *WorkerPool) Start() {
    for i := 0; i < p.numWorkers; i++ {
        p.wg.Add(1)
        go p.worker(i)
    }
}

// Stop stops the worker pool gracefully
func (p *WorkerPool) Stop() {
    p.cancel()
    close(p.done)
    p.wg.Wait()
    close(p.tasks)
    close(p.results)
}

// Submit submits a task to the worker pool
func (p *WorkerPool) Submit(task Task) error {
    select {
    case p.tasks <- task:
        return nil
    case <-p.ctx.Done():
        return p.ctx.Err()
    }
}

// GetResult retrieves a result from the worker pool
func (p *WorkerPool) GetResult() (Result, error) {
    select {
    case result := <-p.results:
        return result, nil
    case <-p.ctx.Done():
        return Result{}, p.ctx.Err()
    }
}

// worker processes tasks from the tasks channel
func (p *WorkerPool) worker(id int) {
    defer p.wg.Done()

    for {
        select {
        case task, ok := <-p.tasks:
            if !ok {
                return
            }
            result := task.Execute()
            select {
            case p.results <- result:
            case <-p.ctx.Done():
                return
            }
        case <-p.ctx.Done():
            return
        }
    }
}

// Example task implementation
type ExampleTask struct {
    Data string
}

func (t ExampleTask) Execute() Result {
    // Simulate some work
    return Result{
        Value: "Processed: " + t.Data,
        Error: nil,
    }
} 