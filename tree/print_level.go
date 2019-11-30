package main

import (
	"fmt"
)
type Node
func printLevel(n *Node){
}
func main() {
	slice:=[]int{3,1,2,4,8}
	tree := &Node{1,
		&Node{2,
			&Node{4, nil, nil}, &Node{5, nil, nil},
		},
		&Node{3,
			&Node{6, nil, nil}, &Node{7, nil, nil},
		},
	}
	fmt.Printf(" %#+v \n", tree)
}
struct {
	V int
	L *Node
	R *Node
}