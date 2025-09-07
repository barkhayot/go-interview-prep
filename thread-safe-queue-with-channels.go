package main

import (
	"errors"
)

var ErrQueueFull = errors.New("queue is full")

type Queue struct{
	channel chan int
}

func NewQueue(size int) *Queue {
  return &Queue{
		channel : make(chan int, size),
	}
}

func (q *Queue) Push(val int) error {
  select {
	case q.channel <- val:
		return nil
	default:
		return ErrQueueFull
	}
}

func (q *Queue) Pop() int {
	select {
	case val := <- q.channel:
		return val
	default:
		return -1
	}
}
