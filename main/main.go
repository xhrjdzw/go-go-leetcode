package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func ping(pings chan<- string, msg string) {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func loop() {
	for i := 0; i < 10; i++ {
		fmt.Print("%d ", i)
	}
}

func Filters(s []int, fn func(int) bool) []int {
	var p []int
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
			//	fmt.Println(v)
		}
		//	fmt.Println(p)
	}
	fmt.Println(p)
	return p
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	test := ListNode{}
	deleteNode(&test)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
*/
func maxDepth(root *TreeNode) int {

}
