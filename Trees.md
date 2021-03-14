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

	/* Keep popping the Nodes until queue is not empty i.e. all nodes are not traversed*/
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



