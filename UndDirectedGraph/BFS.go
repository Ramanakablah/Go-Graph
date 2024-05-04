package main

import (
	"fmt"
	queue "graph/Queue"
)

type BFS interface {
	Bfstravel(G *Graph, s int)
}

func DistanceTo(G Graph, s int, d int) int {
	Q := queue.Queue{}
	marked := map[int]bool{}
	Q.Init()
	Q.Push(s)
	count := 0
	marked[s] = true
	for Q.Size() != 0 {
		w := Q.Size()
		for i := 0; i < w; i++ {
			node := Q.First()
			for _, j := range G.myGraph[node] {
				_, ok := marked[j]
				if !ok {
					marked[j] = true
					if j == d {
						return count
					}
					Q.Push(j)
				}
			}
			Q.Pop()
		}
		count++
	}
	return -1

}

func Bfstravel(G Graph, s int) {
	Q := queue.Queue{}
	marked := map[int]bool{}
	Q.Init()
	Q.Push(s)
	marked[s] = true
	for Q.Size() != 0 {
		node := Q.First()
		fmt.Println(node)
		for _, j := range G.myGraph[node] {
			_, ok := marked[j]
			if !ok {
				marked[j] = true
				Q.Push(j)
			}
		}
		Q.Pop()
	}
}
