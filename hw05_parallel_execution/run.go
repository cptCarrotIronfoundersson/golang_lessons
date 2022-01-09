package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, workersCount, maxErrorsCount int) error {
	var totalErrorsCount int
	mu := sync.Mutex{}
	for i := 0; i < len(tasks); i += workersCount {
		wg := sync.WaitGroup{}
		var executeTasks []Task
		if i+workersCount <= len(tasks) {
			executeTasks = tasks[i : i+workersCount]
		} else {
			executeTasks = tasks[i:]
		}
		for _, function := range executeTasks {
			function := function
			wg.Add(1)
			go func() {
				defer mu.Unlock()
				defer wg.Done()

				err := function()
				mu.Lock()
				if err != nil {
					totalErrorsCount++
				}
			}()
		}
		wg.Wait()
		if totalErrorsCount >= maxErrorsCount {
			return ErrErrorsLimitExceeded
		}
	}

	return nil
}
