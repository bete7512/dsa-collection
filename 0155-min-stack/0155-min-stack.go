type MinStack struct {
    values []int
    mins  []int
}


func Constructor() MinStack {
    return MinStack{}
}


func (this *MinStack) Push(value int)  {
    if len(this.mins) == 0{
        this.mins = append(this.mins,value)
    } else if (this.mins[len(this.mins)-1] > value) {
        this.mins = append(this.mins,value)
    } else {
                this.mins = append(this.mins,this.mins[len(this.mins)-1])
    }
    this.values = append(this.values, value)
}


func (this *MinStack) Pop()  {
    this.values =     this.values[:len(this.values)-1]
    this.mins =     this.mins[:len(this.mins)-1]
}


func (this *MinStack) Top() int {
    if len(this.values) <= 0{
        return 0
    }
    return this.values[len(this.values)-1]
}


func (this *MinStack) GetMin() int {
       if len(this.mins) <= 0{
        return 0
    }
    return this.mins[len(this.mins)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */