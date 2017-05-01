/*

Problem: given a binary search tree with nodes having right and left pointers,
write a function `setNext` that populates a node's next pointer such that
sibling nodes are ordered from left to right. For example:

input:
       A
      / \
     /   \
    B     C
     \    /\
     D   E  F
    /        \
   G          H

after processing:
       A
      / \
     /   \
    B --->C
     \    /\
     D-->E-->F
    /        \
   G--------->H
*/

package main

import "fmt"

type Node struct {
	left  Node
	right Node
	next  Node
}

func (n *Node) setNext(root Node) {
	// if left != null, set next to right

	// h=0
	// A.left = B
	// A.right = C
	// ---
	// B.next = C

	// h=1
	// B.right = D
	// C.left = E
	// C.right = F
	// ---
	// D.next = E
	// E.next = F

	// h=2
	// D.left = G
	// F.right = H

	// h=3
	// D.left = G
	// F.right = H

	/*
	  children = []
	  if root.left
	    children.append(root.left)
	    return setNext(root.left)
	  if root.right
	    children.append(root.right)
	    return setNext(root.right)

	  // children : contains all nodes at the current level
	  next_children = []
	  for c in children {
	    // accumulate the next level into next_children
	    // set .next pointers for nodes at this level
	    next_children.append(setNext(root.left))
	    // index check
	    c.next = children[i+1]
	  }
	  return next_children
	*/
}

func main() {
	cases = []struct {
		input  Node
		output Node
	}{}
	fmt.Println("Foo")
}
