package main

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
	i := 0
	for(l.Head != nil){
		l.Head = l.Head.Next
		i++;
	}
	return i
}

