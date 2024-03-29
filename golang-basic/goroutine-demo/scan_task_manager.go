package goroutinedemo

import "fmt"

type scanTaskManager struct {
	tasks map[string]chan int
}

func NewScanTask() *scanTaskManager {
	return &scanTaskManager{tasks: make(map[string]chan int, 0)}
}

func (s *scanTaskManager) Start(taskName string) {
	if _, ok := s.tasks[taskName]; ok {
		return
	}

	ch := make(chan int)

	s.tasks[taskName] = ch

	go func(ch chan int, taskName string) {
		defer fmt.Printf("task[%v] done\n", taskName)

		for {
			select {
			case <-ch:
				return
			default:
				fmt.Printf("task[%v] running\n", taskName)
			}
		}

	}(ch, taskName)
}

func (s *scanTaskManager) Stop(taskName string) {
	if ch := s.tasks[taskName]; ch != nil {
		ch <- 1
	}
}
