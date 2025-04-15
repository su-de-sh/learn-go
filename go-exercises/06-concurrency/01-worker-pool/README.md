# Worker Pool

Implement a worker pool pattern that processes tasks concurrently using goroutines and channels.

## Requirements

1. Create a `WorkerPool` struct with the following fields:

   - `numWorkers`: number of worker goroutines
   - `tasks`: channel for incoming tasks
   - `results`: channel for task results
   - `done`: channel for signaling completion

2. Implement the following methods:

   - `NewWorkerPool(numWorkers int) *WorkerPool`: Creates a new worker pool
   - `Start()`: Starts the worker pool
   - `Stop()`: Stops the worker pool gracefully
   - `Submit(task Task)`: Submits a task to the pool
   - `GetResult() Result`: Gets a result from the pool

3. Create a `Task` interface:

   ```go
   type Task interface {
       Execute() Result
   }
   ```

4. Create a `Result` struct to hold task results:
   ```go
   type Result struct {
       Value interface{}
       Error error
   }
   ```

## Example Usage

```go
pool := NewWorkerPool(3)
pool.Start()
defer pool.Stop()

// Submit tasks
pool.Submit(MyTask{data: "task1"})
pool.Submit(MyTask{data: "task2"})

// Get results
result := pool.GetResult()
if result.Error != nil {
    log.Fatal(result.Error)
}
fmt.Println(result.Value)
```

## Requirements

1. The worker pool should:

   - Process tasks concurrently
   - Handle task submission and result retrieval
   - Gracefully shut down when stopped
   - Handle errors appropriately

2. Tasks should:

   - Implement the Task interface
   - Return results through the Result struct
   - Handle errors gracefully

3. The implementation should:
   - Use buffered channels appropriately
   - Prevent goroutine leaks
   - Handle concurrent access safely
   - Provide proper error handling

## Tips

- Use `select` statements for channel operations
- Consider using `context.Context` for cancellation
- Use `sync.WaitGroup` for coordinating goroutines
- Remember to close channels when done
- Handle panics in goroutines
