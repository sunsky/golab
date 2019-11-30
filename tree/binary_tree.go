package main

import "fmt"

type Node struct {
	V int
	L *Node
	R *Node
}

//前序
func forwardLook(root *Node) {
	if root == nil {
		return
	}
	//输出行的位置在最前面
	fmt.Printf("node %v ", root.V)
	forwardLook(root.L)
	forwardLook(root.R)
}

//var i int
func forwardLoop(root *Node) {
	//需要一个堆保存走过的路径
	nodes := []*Node{}
	for len(nodes) != 0 || root != nil {
		//一直往左走
		for root != nil {
			nodes = append(nodes, root)
			fmt.Printf("node %v ", root.V)
			root = root.L
		}
		//说明左子结点为空，那么就看右结点
		if len(nodes) > 0 {
			root = nodes[len(nodes)-1]
			//用完最近一个结点后，删除它，删除后最后的结点一定是父结点
			nodes = nodes[:len(nodes)-1]
			//左子结点遍历完了，所以这里只看当看结点的右子结点
			root = root.R
		} else {
			root = nil
		}
	}
}

//中序
func middleLook(root *Node) {
	if root == nil {
		return
	}
	middleLook(root.L)
	//输出行的位置在中间
	fmt.Printf("node %v ", root.V)
	middleLook(root.R)
}

//后序
func backwardLook(root *Node) {
	if root == nil {
		return
	}
	//输出行的位置在
	backwardLook(root.L)
	backwardLook(root.R)
	fmt.Printf("node %v ", root.V)
}

func main() {
	tree := &Node{1,
		&Node{2,
			&Node{4, nil, nil}, &Node{5, nil, nil},
		},
		&Node{3,
			&Node{6, nil, nil}, &Node{7, nil, nil},
		},
	}
	fmt.Println("\nforwardLook ")
	forwardLook(tree)
	fmt.Println("\nforwardLoop ")
	forwardLoop(tree)
	fmt.Println("\nmiddleLook ")
	middleLook(tree)
	fmt.Println("\nbackwardLook ")
	backwardLook(tree)

	tree = &Node{1,
		&Node{2,
			nil,
			&Node{4,
				nil,
				&Node{6,
					&Node{7, nil, nil},
					&Node{8, nil, nil},
				},
			},
		},
		&Node{3,
			nil, &Node{5, nil, nil},
		},
	}
	fmt.Println("\nforwardLook ")
	forwardLook(tree)
	fmt.Println("\nforwardLoop ")
	forwardLoop(tree)
	fmt.Println("\nmiddleLook ")
	middleLook(tree)
	fmt.Println("\nbackwardLook ")
	backwardLook(tree)
}
