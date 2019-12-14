package main


func validateStackSequences(pushed []int, popped []int) bool {
	push_ptr := 0
	pop_ptr := 0
	tempStack := make([]int, 0)

	for push_ptr < len(pushed) && pop_ptr < len(popped) {
		if pushed[push_ptr] != popped[pop_ptr] {
			if len(tempStack) > 0 && tempStack[len(tempStack) - 1] == popped[pop_ptr] {
				tempStack = tempStack[:len(tempStack) - 1]
				pop_ptr++
			}else {
				tempStack = append(tempStack, pushed[push_ptr])
				push_ptr++
			}
		} else {
			push_ptr++
			pop_ptr++
		}
	}
	for i := len(tempStack) - 1; i >= 0;i-- {
		if tempStack[i] != popped[pop_ptr] {
			return false
		}
		pop_ptr++
	}
	return true
}