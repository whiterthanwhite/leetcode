package minstack

type MinStack struct {
	first *Node
	min   *MinNode
}

type Node struct {
	item int
	next *Node
}

type MinNode struct {
	item *Node
	next *MinNode
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	oldFirst := this.first
	this.first = &Node{
		item: val,
		next: oldFirst,
	}

	if oldFirst == nil {
		this.min = &MinNode{
			item: this.first,
		}
	} else {
		if this.min.item.item > val {
			oldMin := this.min
			this.min = &MinNode{
				item: this.first,
				next: oldMin,
			}
		}
	}
}

func (this *MinStack) Pop() {
	if this.first == this.min.item {
		this.min = this.min.next
	}
	this.first = this.first.next
}

func (this *MinStack) Top() int {
	return this.first.item
}

func (this *MinStack) GetMin() int {
	return this.min.item.item
}
