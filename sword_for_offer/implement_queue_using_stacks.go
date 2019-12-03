package main

type MyQueue struct {
	data_stack []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	queue := MyQueue {
		data_stack : make([]int, 0),
	}
	return queue
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.data_stack = append(this.data_stack, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.data_stack) == 0 {
		return 0
	}
	data := this.data_stack[0]
	this.data_stack = this.data_stack[1:]
	return data
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.data_stack) == 0 {
		return 0
	}
	data := this.data_stack[0]
	return data
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.data_stack) == 0 {
		return true
	}
	return false
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */