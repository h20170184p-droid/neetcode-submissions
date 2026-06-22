type MinStack struct {
	contents []int
	mins     []int
}

func Constructor() MinStack {
	return MinStack{contents: []int{}, mins: []int{}}
}

func (this *MinStack) Push(val int) {
	this.contents = append(this.contents, val)
	if len(this.mins) == 0 || val <= this.mins[len(this.mins)-1] {
		this.mins = append(this.mins, val)
	}
}

func (this *MinStack) Pop() {
	val := this.Top()
	this.contents = this.contents[: (len(this.contents) - 1)]
	if val == this.mins[len(this.mins)-1] {
		this.mins = this.mins[: (len(this.mins) - 1)]
	}
}

func (this *MinStack) Top() int {
	return this.contents[len(this.contents) - 1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
