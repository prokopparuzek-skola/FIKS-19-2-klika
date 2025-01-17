// psáno v jazyce golang
package main

import "fmt"

type edge struct { // hrana grafu
	to   int
	cost int
}

type graph struct { // graf
	K       int
	X       int
	flights [][]edge
}
type top struct { // vrchol v haldě
	vrx  int
	cost int
}

type state struct { // vrchol
	cost      int
	heapIndex int
}

func dijkstra(airports *graph, start int) {
	// pole vrcholů, ukládají se zde dosud spočítané nejkratší cesty
	var vertex []state = make([]state, len(airports.flights))
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
			for i := 1; i <= airports.K; i++ {
				if v == i {
					continue
				} else if vertex[i].cost == -1 { // ještě jsem tu nebyl
					vertex[i].cost = vertex[v].cost + airports.X // cena
					vertex[i].heapIndex = len(Overtex)
					heapAdd(&Overtex, &vertex, i)
				} else if vertex[i].cost > (vertex[v].cost + airports.X) {
					vertex[i].cost = vertex[v].cost + airports.X // uprava ceny
					Overtex[vertex[i].heapIndex].cost = vertex[v].cost + airports.X
					bubbleUp(&Overtex, &vertex, vertex[i].heapIndex)
				}
			}
		}
		for _, w := range airports.flights[v] { // ostatní letiště
			if vertex[w.to].cost == -1 {
				vertex[w.to].cost = vertex[v].cost + w.cost // cena
				vertex[w.to].heapIndex = len(Overtex)
				heapAdd(&Overtex, &vertex, w.to)
			} else if vertex[w.to].cost > (vertex[v].cost + w.cost) {
				vertex[w.to].cost = vertex[v].cost + w.cost // uprava ceny
				Overtex[vertex[w.to].heapIndex].cost = vertex[v].cost + w.cost
				bubbleUp(&Overtex, &vertex, vertex[w.to].heapIndex)
			}
		}
	}
	for i := 1; i < len(vertex); i++ {
		if i < len(vertex)-1 {
			fmt.Printf("%d ", vertex[i].cost)
		} else {
			fmt.Printf("%d\n", vertex[i].cost)
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
			(*vertex)[(*Overtex)[n].vrx].heapIndex = parent
			(*vertex)[(*Overtex)[parent].vrx].heapIndex = n
			swp := (*Overtex)[parent]
			(*Overtex)[parent] = (*Overtex)[n]
			(*Overtex)[n] = swp
		}
		n /= 2
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
			(*vertex)[(*Overtex)[n].vrx].heapIndex = son
			(*vertex)[(*Overtex)[son].vrx].heapIndex = n
			swp := (*Overtex)[son]
			(*Overtex)[son] = (*Overtex)[n]
			(*Overtex)[n] = swp
		}
		n *= 2
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
	swp := (*Overtex)[1]
	vrx = swp.vrx
	(*Overtex)[1] = (*Overtex)[len(*Overtex)-1]
	*Overtex = (*Overtex)[:len(*Overtex)-1]
	(*vertex)[vrx].heapIndex = -1
	for i := 1; i < len(*vertex); i++ {
		if (*vertex)[i].heapIndex != -1 {
			(*vertex)[i].heapIndex--
		}
	}
	bubbleDown(Overtex, vertex, 1)
	return
}

func main() {
	var N, K, X, M, V int
	var airports graph

	fmt.Scanf("%d %d %d %d %d", &N, &K, &X, &M, &V)
	airports.K = K
	airports.X = X
	airports.flights = make([][]edge, N+1)
	for i := 0; i < N; i++ {
		airports.flights[i] = make([]edge, 0)
	}
	for i := 0; i < M; i++ { // načtení letů
		var from, to, cost int
		fmt.Scanf("%d %d %d", &from, &to, &cost)
		airports.flights[from] = append(airports.flights[from], edge{to, cost})
		airports.flights[to] = append(airports.flights[to], edge{from, cost})
	}
	dijkstra(&airports, V) // prohledání
}
