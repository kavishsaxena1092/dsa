# Trees
## Introduction
A tree data structure can be defined recursively (locally) as a **collection of nodes** (starting at a root node), where each **node is a data structure consisting of a value, together with a list of references to nodes (the "children")**, with the constraints that no reference is duplicated, and none points to the root.


A tree is a datastructure similiar to a linked list but instead of each node pinting to the next node in a linear fashion, each node points to a nmumber of nodes. A tree is an example of **non linear data structure**. A tree structure is an example of r**epresenting hiererichal nature of data in structure form**
![tree](/assets/tree.png)

## Glossary
- **root** - A node with no parents. Its the starting point of a tree. There can be atmost one root node in a tree.
- **edge** - An edge is a link from parent to child
- **siblings** - Children of same node
- **ancestor, descendent** - node p is called ancestor of node q if there exists a path from node q to root with node p in its path. in this case node q is called descendent of node p.  
- **level** - set of all nodes at a given depth is called a level
![Screen Shot 2018-06-02 at 12.35.26 AM](/assets/Screen%20Shot%202018-06-02%20at%2012.35.26%20AM.png)

- **depth** - depth of a node is length of the path from root node to that node. depth of a tree is the length of the path from the deepest node to the root node
- **height** - height of a node is the length of the path from the node to the deepest node in the tree. height of a tree is the length of the path from the root node to the deepest node in the tree.
    *height and depth of a tree are always same*
- **size** - size of a node is the number of descendents it has, including the node.
- **skew trees** - if in a tree every node has only one child, it's called a skew tree.
  - **left skew trees** - a skew tree in which every node has only left child
  - **right skew trees** - a skew tree in which every node has only right child
![Screen Shot 2018-06-02 at 12.44.31 AM](/assets/Screen%20Shot%202018-06-02%20at%2012.44.31%20AM.png)