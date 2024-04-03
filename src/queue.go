package src

import "sync"

type Queue struct {
	list  []string
	front int
	last  int
	sync.RWMutex
}

func NewQueue() *Queue {
	return &Queue{
		list:  make([]string, 0, 10000),
		front: -1,
		last:  -1,
	}
}
