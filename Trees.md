# Trees
## Binary Tree

<p>Unline stacks, arrays,queues, and linkedlists, Tree is a non-linear(hierarchical) data structure,
Something like filesystem, database these kinds of software systems are built using trees. </p>
</br>
We will discuss how to create a Binary Tree.
In a hierarchical data structure like Tree, we have concept of Nodes. where each Node has children , using which an hierarchy is built. 
An example is file system, like Folder contains sub-folders and files and so on. it can contain multiple folders inside it.

### Introduction
In a Binary Tree, Each Node has 2 children.<br>
a) Left Child Node <br>
b) Right Child Node <br>

<pre>

      tree
      ----
       j    &lt;-- root
     /   \
    f      k  
  /   \      \
 a     h      z    &lt;-- leaves
</pre>

 
```

/* Tree Node Contain Left and Right Child Nodes */
type Node struct{
    Left *Node;
    Right *Node;
    Key int;
}
```


## Traversal Binary Tree(Level-Order)

<p>We will BFS for Level Order Traversal, in which we will read elements at each level using queue.<br>
We will first implement Queue using slice and implement basic operations for Queue(push,pop, front) 
</p>

Code For Queue Implementation.
```
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

```
<i><b>Level Order Traversal Code</b></i>
     
```
func (n *Node) levelOrder() {
	if n == nil {
		return
	}
	fmt.Printf("\n")
	/* Create a queue and push the root Node to queue */
	q := &Queue{}
	q.push(n)

	/* Keep popping the Nodes until queue is not empty 
	 *  i.e. all nodes are not traversed
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
```

<i><b>Insert a Node in Tree </b></i>
<p> We will insert a node in Tree by taking use of Level Order Traversal, and while traversal where we find 1st child node which is empty(node for which left/right is nil) </p>

```
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
			/* Return from the point where we have inserted the node*/
		}

		if tempNode.Right != nil {
			q.push(tempNode.Right)
		} else {
			fmt.Println("Inserting in Right Side")
			tempNode.Right = &Node{Key: num, Left: nil, Right: nil}
			return
			/* Return from the point where we have inserted the node*/
		}

	}
}
```

<i><b>Delete a Node in Tree </b></i>
<ol>
	<li>Find the Node which is to be deleted using Level Order Traversal.</li>
	<li>Save the last Node as deepest Node in Level Order Traversal i.e. rightMost Node in Tree.</li>
	<li>Replace the Data of Deepest Node with KeyNode(To be deleted), and delete the Deepest Node.</li>
</ol>

```
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
```

<i><b>Level Order Traversal Using DFS </b></i>
<ol>
<li>Use Recursion to traverse to deeper nodes</li>
<li>This Method can be used to print nodes which are k distant from root using same methodology.
</ol>

```
/*Recurive Function*/
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

```

Call this Recurive Function for Each Level or obtain Height
of Tree and utilise that. 

```
	v := true
	level := 0
	for v {
		v = root.levelOrderDFS(level)
		level++
	}
```

```
//h = height of Tree
for level<h{
	 root.levelOrderDFS(level)
	level++
}
```


