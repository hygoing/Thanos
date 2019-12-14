package main

type MinStack struct {
	nums []int
	temp []int
}


/** initialize your data structure here. */
func MinStackInit() MinStack {
	stack := MinStack {
		nums: make([]int, 0),
		temp: make([]int, 0),
	}
	return stack
}


func (this *MinStack) Push(x int)  {
	this.nums = append(this.nums, x)
	if len(this.temp) == 0 || this.temp[len(this.temp) - 1] >= x {
		this.temp = append(this.temp, x)
	}
}


func (this *MinStack) Pop()  {
	if len(this.nums)  == 0 {
		return
	}
	num := this.nums[len(this.nums) - 1]
	this.nums = this.nums[:len(this.nums) - 1]
	if num == this.temp[len(this.temp) - 1] {
		this.temp = this.temp[:len(this.temp) - 1]
	}
}


func (this *MinStack) Top() int {
	if len(this.nums) == 0 {
		return -1
	}
	return this.nums[len(this.nums) - 1]
}


func (this *MinStack) GetMin() int {
	return this.temp[len(this.temp) - 1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */