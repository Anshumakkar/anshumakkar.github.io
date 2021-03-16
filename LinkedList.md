<h1>  Linked List </h1>

Below is basic implementation of LinkedList - Singly Linked List which contains address to next Node in the list

```
package main

import "fmt"

type LLNode struct {
	data interface{}
	next *LLNode
}

func (node *LLNode) Print() {
	if node == nil {
		fmt.Printf("-- END --\n\n")
		return
	}

	fmt.Println(node)

	node.next.Print()
}

func New(data interface{}) *LLNode {
	return &LLNode{data: data, next: nil}
}

func Add(root *LLNode, data interface{}) *LLNode {
	return &LLNode{
		data: data,
		next: root,
	}
}

/* This is Recursive Implementation of Reversal
* of Linked List
*/
func Reverse(root *LLNode) *LLNode {
	if root == nil || root.next == nil {
		return root
	}
	rest := Reverse(root.next)
	root.next.next = root
	root.next = nil
	return rest
}

func main() {
	root := New(1)
	root = Add(root, 2)
	root = Add(root, 3)
	root = Add(root, 4)
	root = Add(root, 5)

	root.Print()
	reversed := Reverse(root)
	reversed.Print()
}

```
