type MinStack struct {
	val []int
	min []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	lm := len(this.min) - 1
	if len(this.min) == 0 || this.min[lm] >= val {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, this.min[lm])
	}
	this.val = append(this.val, val)
}

func (this *MinStack) Pop() {
	l := len(this.val) - 1
	lm := len(this.min) - 1
	this.val = this.val[:l]
	this.min = this.min[:lm]
}

func (this *MinStack) Top() int {
	return this.val[len(this.val)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}