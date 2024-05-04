package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Graph struct {
	V       int
	myGraph map[int][]int
}

func (G *Graph) Init(V int) {
	G.V = V
}

func (G *Graph) Adj(v int) []int {
	return G.myGraph[v]
}

func (G *Graph) NodeCount(v int) int {
	return len(G.myGraph[v])
}

func (G *Graph) GraphBuild() {
	G.myGraph = map[int][]int{}
	file, err := os.Open("DataFile.txt")
	if err != nil {
		fmt.Printf("Error")
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		str := scanner.Text()
		if count == 0 {
			i, err := strconv.Atoi(str)
			if err != nil {
				log.Fatal(err)
			}
			G.Init(i)
		}
		if count > 1 {
			x := strings.Split(str, "-")
			v, errv := strconv.Atoi(x[0])
			w, errw := strconv.Atoi(x[1])
			if errw != nil {
				log.Fatal(errw)
			}
			if errv != nil {
				log.Fatal(errv)
			}
			if _, ok := G.myGraph[v]; !ok {
				G.myGraph[v] = []int{}
			}
			if _, ok := G.myGraph[w]; !ok {
				G.myGraph[w] = []int{}
			}
			G.myGraph[v] = append(G.myGraph[v], w)
			G.myGraph[w] = append(G.myGraph[w], v)
		}
		count++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func main() {

	G := Graph{}
	G.Init(9)
	G.GraphBuild()
	// fmt.Println(G.myGraph)
	// Bfstravel(G, 1)

	fmt.Println(DistanceTo(G, 1, 13))

}
