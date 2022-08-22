package main

import (
	"fmt"
)

type Queue struct {
	data []interface{}
}

// add at the end of Queue
func (q *Queue) enqueue(element interface{}) {
	q.data = append(q.data, element)
}

// get Element at begining of Queue
func (q *Queue) dequeue() interface{} {
	length := len(q.data)
	if length > 0 {
		front := q.data[length-1]
		q.data = q.data[:length-1]
		return front
	}
	return nil
}

// Read the first Element in Queue
func (q *Queue) front() interface{} {
	length := len(q.data)
	if length > 0 {
		return q.data[length-1]
	}
	return nil
}

// get the last Element in Queue
func (q *Queue) back() interface{} {
	length := len(q.data)
	if length > 0 {
		return q.data[0]
	}
	return nil
}

// get length of Queue
func (q *Queue) length() int {
	return len(q.data)
}

// clear all Elements in Queue
func (q *Queue) clear() {
	q.data = []interface{}{}
}

// print all Elements in Queue
func (q *Queue) print() {
	fmt.Println("Queue: ", q.data)
}
func main() {
	var q Queue
	q.enqueue("Hello World")
	q.enqueue(1)
	q.enqueue(2)
	q.print()
	fmt.Println("Get the front Element: ", q.dequeue())
	q.print()
}
