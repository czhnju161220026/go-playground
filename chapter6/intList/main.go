package main

import "fmt"

type List struct {
	val  int
	next *List
}

func (this *List) Sum() int {
	if this == nil {
		return 0
	}
	return this.val + this.next.Sum()
}

func (this *List) Append(node *List) {
	this.next = node
}

func main() {
	head := List{1, nil}
	p := &head
	for i := 2; i < 10; i++ {
		node := &List{i, nil}
		p.Append(node)
		p = node
	}
	fmt.Println(head.Sum())
}
