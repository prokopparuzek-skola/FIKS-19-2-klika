package main

import "fmt"

type edge struct {
	to   int
	cost int
}

type graph struct {
	K       int
	X       int
	flights [][]edge
}
type vertex struct {
	vrx  int
	cost int
}

func djakstra(airports *graph, start int) (costs []int) {
	costs = make([]int, len(airports.flights)+airports.K+1)
	costs[start] = 0
	var vertexes []vertex = make([]vertex, len(airports.flights)+airports.K+1)
	for i := range vertexes {
		vertexes[i].vrx = i
		vertexes[i].cost = -1
	}
	vertexes[start].cost = 0
	return
}

func bubbleUp(vertexes *[]vertex, n int) {
	for n > 1 {
		parent := n / 2
		if (*vertexes)[parent].cost < (*vertexes)[n].cost {
			break
		} else {
			swp := (*vertexes)[parent]
			(*vertexes)[parent] = (*vertexes)[n]
			(*vertexes)[n] = swp
		}
	}
}

func bubbleDown(vertexes *[]vertex, n int) {
	for (2 * n) < len(*vertexes) {
		son := 2 * n
		if (son+1) < len((*vertexes)) && (*vertexes)[son+1].cost < (*vertexes)[son].cost {
			son++
		} else if (*vertexes)[n].cost < (*vertexes)[son].cost {
			break
		} else {
			swp := (*vertexes)[son]
			(*vertexes)[son] = (*vertexes)[n]
			(*vertexes)[n] = swp
		}
	}
}

func main() {
	var N, K, X, M, V int
	var airports graph

	fmt.Scanf("%d %d %d %d %d", &N, &K, &X, &M, &V)
	airports.K = K
	airports.X = X
	airports.flights = make([][]edge, N)
	for i := 0; i < N; i++ {
		airports.flights[i] = make([]edge, 0)
	}
	for i := 0; i < M; i++ {
		var from, to, cost int
		fmt.Scanf("%d %d %d", &from, &to, &cost)
		airports.flights[from] = append(airports.flights[from], edge{to, cost})
		airports.flights[to] = append(airports.flights[to], edge{from, cost})
	}
}
