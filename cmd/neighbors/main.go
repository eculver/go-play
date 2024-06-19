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

// TreeNode represents a node in the binary search tree.
type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
	Next  *TreeNode
}

// setNext populates each node's next pointer to point to its right sibling.
func setNext(root *TreeNode) {
	if root == nil {
		return
	}

	// Initialize a queue for level order traversal.
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// Get the number of nodes at the current level.
		levelSize := len(queue)

		// Iterate through nodes at the current level.
		for i := 0; i < levelSize; i++ {
			// Pop the front node from the queue.
			node := queue[0]
			queue = queue[1:]

			// Set the next pointer to the next node in the queue if it exists.
			if i < levelSize-1 {
				node.Next = queue[0]
			}

			// Enqueue the left and right children of the current node.
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
}

// Helper function to print the tree level by level with next pointers.
func printTreeWithNext(root *TreeNode) {
	if root == nil {
		return
	}

	// Initialize a queue for level order traversal.
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		// Get the number of nodes at the current level.
		levelSize := len(queue)

		// Iterate through nodes at the current level.
		for i := 0; i < levelSize; i++ {
			// Pop the front node from the queue.
			node := queue[0]
			queue = queue[1:]

			// Print the current node and its next pointer.
			if node.Next != nil {
				fmt.Printf("%s-->", node.Val)
			} else {
				fmt.Printf("%s\n", node.Val)
			}

			// Enqueue the left and right children of the current node.
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
}

func main() {
	// Construct the binary search tree.
	root := &TreeNode{Val: "A"}
	root.Left = &TreeNode{Val: "B"}
	root.Right = &TreeNode{Val: "C"}
	root.Left.Right = &TreeNode{Val: "D"}
	root.Right.Left = &TreeNode{Val: "E"}
	root.Right.Right = &TreeNode{Val: "F"}
	root.Left.Right.Left = &TreeNode{Val: "G"}
	root.Right.Right.Right = &TreeNode{Val: "H"}

	// Set the next pointers.
	setNext(root)

	// Print the tree with next pointers.
	printTreeWithNext(root)
}
