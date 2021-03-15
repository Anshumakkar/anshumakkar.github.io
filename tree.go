package main

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Key   int
}

type Queue struct {
	a []*Node
}

func (q *Queue) push(a *Node) {
	q.a = append(q.a, a)
}
func (q Queue) front() *Node {
	return q.a[0]
}

func (q Queue) isEmpty() bool {
	if len(q.a) == 0 {
		return true
	}
	return false
}

func (q *Queue) pop() {
	q.a = q.a[1:]
}
func (n *Node) inorder() {
	if n == nil {
		return
	}
	if n.Left != nil {
		n.Left.inorder()
	}
	fmt.Printf("%d  ", n.Key)
	if n.Right != nil {
		n.Right.inorder()
	}
}
func (n *Node) levelOrder() {
	if n == nil {
		return
	}
	fmt.Printf("\n")
	/* Create a queue and push the root Node to queue */
	q := &Queue{}
	q.push(n)

	/* Keep popping the Nodes until queue is not empty i.e. all nodes are not
	   traversed
	*/
	for !q.isEmpty() {
		temp := q.front()
		q.pop()

		fmt.Printf("%d ", temp.Key)
		//Push the Left Node if not nil
		if temp.Left != nil {
			q.push(temp.Left)
		}
		//Push the Right Node if not nil
		if temp.Right != nil {
			q.push(temp.Right)
		}
	}
	fmt.Printf("\n")
}

func (n *Node) deleteDeepest(deepNode *Node) {
	if n == nil {
		return
	}

	q := &Queue{}
	q.push(n)

	for !q.isEmpty() {
		temp := q.front()
		q.pop()

		if temp == deepNode {
			temp = nil
		}

		if temp.Left != nil {
			if temp.Left == deepNode {
				temp.Left = nil
				return
			} else {
				q.push(temp.Left)
			}
		}

		if temp.Right != nil {
			if temp.Right == deepNode {
				temp.Right = nil
				return
			} else {
				q.push(temp.Right)
			}
		}

	}
}

func (n *Node) Delete(num int) {
	//Find the RightMost Node and Replace the content of Node to be Found
	if n == nil {
		return
	}

	q := &Queue{}
	q.push(n)
	var keyNode *Node
	var temp *Node
	for !q.isEmpty() {
		temp = q.front()
		q.pop()

		if temp.Key == num {
			keyNode = temp
		}

		if temp.Left != nil {
			q.push(temp.Left)
		}

		if temp.Right != nil {
			q.push(temp.Right)
		}

	}

	if keyNode != nil {
		deepestNode := temp
		deepestNodeKey := temp.Key
		n.deleteDeepest(deepestNode)
		keyNode.Key = deepestNodeKey
	}
}

func (n *Node) Insert(num int) {

	if n == nil {
		return
	}

	q := &Queue{}

	q.push(n)

	for !q.isEmpty() {
		tempNode := q.front()
		q.pop()

		if tempNode.Left != nil {
			q.push(tempNode.Left)
		} else {
			fmt.Println("Inserting in Left Side")
			tempNode.Left = &Node{Key: num, Left: nil, Right: nil}
			return
		}

		if tempNode.Right != nil {
			q.push(tempNode.Right)
		} else {
			fmt.Println("Inserting in Right Side")
			tempNode.Right = &Node{Key: num, Left: nil, Right: nil}
			return
		}

	}
}

func (n *Node) levelOrderDFS(level int) bool {
	if n == nil {
		return false
	}

	if level == 0 {
		fmt.Println(n.Key)
		return true
	}
	left := false
	right := false
	/*Use Recursion to go to deeper nodes*/
	if n.Left != nil {
		left = n.Left.levelOrderDFS(level - 1)
	}
	if n.Right != nil {
		right = n.Right.levelOrderDFS(level - 1)
	}
	return (left || right)
}

func main() {
	root := &Node{Key: 10}

	root.Left = &Node{Key: 11}
	root.Right = &Node{Key: 12}

	root.Left.Left = &Node{Key: 13}
	// root.Left.Right = &Node{Key: 14}

	root.Right.Left = &Node{Key: 15}
	root.Right.Right = &Node{Key: 16}

	root.levelOrder()

	fmt.Println("DFS level order")
	v := true
	level := 0
	for v {
		v = root.levelOrderDFS(level)
		level++
	}
	fmt.Println(" ")

	root.inorder()

	root.Insert(14)

	root.inorder()

	root.Delete(10)
	fmt.Println()
	root.inorder()

}
