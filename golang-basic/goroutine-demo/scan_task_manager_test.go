package goroutinedemo

import (
	"sync"
	"testing"
	"time"
)

func TestScanTaskManager(t *testing.T) {
	manager := NewScanTask()

	var wg sync.WaitGroup

	wg.Add(3)
	manager.Start("task1")
	manager.Start("task2")
	manager.Start("task3")

	time.Sleep(2 * time.Second)

	manager.Stop("task1")
	manager.Stop("task2")
	manager.Stop("task3")

	wg.Done()
	wg.Done()
	wg.Done()

	wg.Wait()
}
