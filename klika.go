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
type top struct {
	vrx  int
	cost int
}

type state struct {
	cost      int
	heapIndex int
}

func dijkstra(airports *graph, start int) (costs []int) {
	// pole výsledků
	costs = make([]int, len(airports.flights)+airports.K+1)
	costs[start] = 0
	// pole vrcholů, ukládají se zde dosud spočítané nejkratší cesty
	var vertex []state = make([]state, len(airports.flights)+airports.K+1)
	for i := range vertex {
		vertex[i].cost = -1
		vertex[i].heapIndex = -1
	}
	// nastavení startu, indexuje se od 1
	vertex[start].cost = 0
	vertex[start].heapIndex = 1
	// halda na aktivní vrcholy
	Overtex := make([]top, 1)
	heapAdd(&Overtex, &vertex, start)
	for len(Overtex) > 1 {
		v := extractMin(&Overtex, &vertex)
		if v <= airports.K { // letiště v paktu
		}
		for i, w := range airports.flights[v] { // ostatní letiště
		}
	}
	return
}

func bubbleUp(Overtex *[]top, vertex *[]state, n int) {
	for n > 1 {
		parent := n / 2
		if (*Overtex)[parent].cost < (*Overtex)[n].cost {
			break
		} else {
			swp := (*Overtex)[parent]
			(*Overtex)[parent] = (*Overtex)[n]
			(*Overtex)[n] = swp
			(*vertex)[n].heapIndex = parent
			(*vertex)[parent].heapIndex = n
		}
	}
}

func bubbleDown(Overtex *[]top, vertex *[]state, n int) {
	for (2 * n) < len(*Overtex) {
		son := 2 * n
		if (son+1) < len((*Overtex)) && (*Overtex)[son+1].cost < (*Overtex)[son].cost {
			son++
		} else if (*Overtex)[n].cost < (*Overtex)[son].cost {
			break
		} else {
			swp := (*Overtex)[son]
			(*Overtex)[son] = (*Overtex)[n]
			(*Overtex)[n] = swp
			(*vertex)[n].heapIndex = son
			(*vertex)[son].heapIndex = n
		}
	}
}

func heapAdd(Overtex *[]top, vertex *[]state, vrx int) {
	*Overtex = append(*Overtex, top{0, 0})
	n := len(*Overtex) - 1
	(*Overtex)[n].vrx = vrx
	(*Overtex)[n].cost = (*vertex)[vrx].cost
	bubbleUp(Overtex, vertex, n)
}

func extractMin(Overtex *[]top, vertex *[]state) (vrx int) {
	swp := (*Overtex)[0]
	vrx = swp.vrx
	(*Overtex)[0] = (*Overtex)[len(*Overtex)-1]
	bubbleDown(Overtex, vertex, 1)
	*Overtex = (*Overtex)[:len(*Overtex)-1]
	return
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
