package main

import (
  "errors"
  "sync"
)

var ErrQueueFull = errors.New("queue is full")

type Queue struct{
  size int
  vals []int
  mu sync.Mutex
}

func NewQueue(size int) *Queue {
  q := Queue{
    size : size,
    vals: []int{},
  }
  return &q
}

func (q *Queue) Push(val int) error {
  q.mu.Lock()
  defer q.mu.Unlock()

  if len(q.vals) == q.size{
    return ErrQueueFull
  }

  q.vals = append(q.vals, val)
  return nil
}

func (q *Queue) Pop() int {
  q.mu.Lock()
  defer q.mu.Unlock()

  if len(q.vals) > 0{
    val := q.vals[0]
    q.vals = q.vals[1:]

    return val
  }
  return 0
}

