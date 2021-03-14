# Trees
## Binary Tree

<span>Unline stacks, arrays,queues, and linkedlists, Tree is a non-linear(hierarchical) data structure,
Something like filesystem, database these kinds of software systems are built using trees. </span>
</br>
We will discuss how to create a Binary Tree.
In a hierarchical data structure like Tree, we have concept of Nodes. where each Node has children , using which an hierarchy is built. 
An example is file system, like Folder contains sub-folders and files and so on. it can contain multiple folders inside it.

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
type Node struct{
    left *Node;
    right *Node;
    key int;
}


func(n* Node)Key(a int){
    n.key = a
}

func(n* Node)Left(leftNode *Node){
    n.left = leftNode
}

func(n* Node)Right(rightNode *Node){
    n.right = rightNode
}

```
