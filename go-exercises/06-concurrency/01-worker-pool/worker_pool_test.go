package main

import (
    "errors"
    "sync"
    "testing"
    "time"
)

// TestTask is a simple task implementation for testing
type TestTask struct {
    ID       int
    Duration time.Duration
    Result   string
    Error    error
}

func (t TestTask) Execute() Result {
    if t.Duration > 0 {
        time.Sleep(t.Duration)
    }
    return Result{
        Value: t.Result,
        Error: t.Error,
    }
}

func TestNewWorkerPool(t *testing.T) {
    pool := NewWorkerPool(3)
    if pool == nil {
        t.Error("NewWorkerPool returned nil")
    }
    if pool.numWorkers != 3 {
        t.Errorf("Expected 3 workers, got %d", pool.numWorkers)
    }
}

func TestWorkerPoolBasic(t *testing.T) {
    pool := NewWorkerPool(2)
    pool.Start()
    defer pool.Stop()

    // Submit tasks
    tasks := []TestTask{
        {ID: 1, Result: "Task 1", Duration: 100 * time.Millisecond},
        {ID: 2, Result: "Task 2", Duration: 100 * time.Millisecond},
        {ID: 3, Result: "Task 3", Duration: 100 * time.Millisecond},
    }

    for _, task := range tasks {
        err := pool.Submit(task)
        if err != nil {
            t.Errorf("Failed to submit task %d: %v", task.ID, err)
        }
    }

    // Get results
    for range tasks {
        result, err := pool.GetResult()
        if err != nil {
            t.Errorf("Failed to get result: %v", err)
            continue
        }
        if result.Error != nil {
            t.Errorf("Task returned error: %v", result.Error)
        }
    }
}

func TestWorkerPoolConcurrent(t *testing.T) {
    pool := NewWorkerPool(5)
    pool.Start()
    defer pool.Stop()

    var wg sync.WaitGroup
    numTasks := 10
    results := make(chan string, numTasks)

    // Submit tasks concurrently
    for i := 0; i < numTasks; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            task := TestTask{
                ID:       id,
                Result:   "Task " + string(id),
                Duration: 50 * time.Millisecond,
            }
            err := pool.Submit(task)
            if err != nil {
                t.Errorf("Failed to submit task %d: %v", id, err)
                return
            }
            result, err := pool.GetResult()
            if err != nil {
                t.Errorf("Failed to get result for task %d: %v", id, err)
                return
            }
            if result.Error != nil {
                t.Errorf("Task %d returned error: %v", id, result.Error)
                return
            }
            results <- result.Value.(string)
        }(i)
    }

    wg.Wait()
    close(results)

    // Collect results
    var count int
    for range results {
        count++
    }
    if count != numTasks {
        t.Errorf("Expected %d results, got %d", numTasks, count)
    }
}

func TestWorkerPoolErrorHandling(t *testing.T) {
    pool := NewWorkerPool(2)
    pool.Start()
    defer pool.Stop()

    // Submit task with error
    task := TestTask{
        ID:     1,
        Result: "Error Task",
        Error:  errors.New("task error"),
    }
    err := pool.Submit(task)
    if err != nil {
        t.Errorf("Failed to submit task: %v", err)
    }

    // Get result and check error
    result, err := pool.GetResult()
    if err != nil {
        t.Errorf("Failed to get result: %v", err)
    }
    if result.Error == nil {
        t.Error("Expected error in result, got nil")
    }
    if result.Error.Error() != "task error" {
        t.Errorf("Expected error 'task error', got '%v'", result.Error)
    }
}

func TestWorkerPoolGracefulShutdown(t *testing.T) {
    pool := NewWorkerPool(3)
    pool.Start()

    // Submit some tasks
    for i := 0; i < 5; i++ {
        task := TestTask{
            ID:       i,
            Result:   "Task " + string(i),
            Duration: 100 * time.Millisecond,
        }
        err := pool.Submit(task)
        if err != nil {
            t.Errorf("Failed to submit task %d: %v", i, err)
        }
    }

    // Stop the pool
    pool.Stop()

    // Try to submit a task after stopping
    err := pool.Submit(TestTask{ID: 6, Result: "Task 6"})
    if err == nil {
        t.Error("Expected error when submitting task after stopping, got nil")
    }

    // Try to get results after stopping
    result, err := pool.GetResult()
    if err == nil {
        t.Error("Expected error when getting result after stopping, got nil")
    }
} 