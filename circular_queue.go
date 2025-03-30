package main

type MyCircularQueue struct {
	front, size, capacity int
	buffer                []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{capacity: k, buffer: make([]int, k)}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.size == this.capacity {
		return false
	}

	this.buffer[(this.front+this.size)%this.capacity] = value
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.size == 0 {
		return false
	}

	this.front = (this.front + 1) % this.capacity
	this.size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.size == 0 {
		return -1
	}
	return this.buffer[this.front]
}

func (this *MyCircularQueue) Rear() int {
	if this.size == 0 {
		return -1
	}
	return this.buffer[(this.front+this.size-1)%this.capacity]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == this.capacity
}
