package main

import (
	"fmt"
	TDACola "grafo/cola"
	TDAHash "grafo/hash"
	TDAHeap "grafo/heap"
	TDAUnionFind "grafo/unionfind"
	"sort"
)

var vertices = []string{"A", "B", "C", "D", "E", "F", "G", "H"}
var plan = []string{"Fisica I", "Fisica II", "Analisis Matematico II", "Algebra II", "Algoritmos y Programacion I", "Algoritmos y Programacion II", "Algoritmos y Programacion III", "Probabilidad y Estadistica", "Estructuras del Computador", "Analisis Numerico", "Taller", "Organizacion de Datos"}

func BFS[T comparable](grafo Grafo[T]) {
	visitados := TDAHash.CrearHash[T, *T]()
	padre := TDAHash.CrearHash[T, *T]()
	for _, vertice := range grafo.ObtenerVertices() {
		if !visitados.Pertenece(vertice) {
			visitados.Guardar(vertice, nil)
			bfs(grafo, vertice, padre, visitados)
		}
	}
}

func bfs[T comparable](grafo Grafo[T], vertice_inicial T, padre TDAHash.Diccionario[T, *T], visitados TDAHash.Diccionario[T, *T]) {
	cola := TDACola.CrearColaEnlazada[T]()
	cola.Encolar(vertice_inicial)
	for !cola.EstaVacia() {
		vertice := cola.Desencolar()
		for _, adyacente := range grafo.Adyacentes(vertice) {
			if !visitados.Pertenece(adyacente) {
				visitados.Guardar(adyacente, nil)
				padre.Guardar(adyacente, &vertice)
				cola.Encolar(adyacente)
			}
		}
	}
}

func DFS[T comparable](grafo Grafo[T]) {
	visitados := TDAHash.CrearHash[T, *T]()
	padre := TDAHash.CrearHash[T, *T]()
	for _, vertice := range grafo.ObtenerVertices() {
		if !visitados.Pertenece(vertice) {
			visitados.Guardar(vertice, nil)
			dfs(grafo, vertice, padre, visitados)
		}
	}
}

func dfs[T comparable](grafo Grafo[T], vertice_inicial T, padre TDAHash.Diccionario[T, *T], visitados TDAHash.Diccionario[T, *T]) {
	padre.Guardar(vertice_inicial, nil)
	for _, adyacente := range grafo.Adyacentes(vertice_inicial) {
		if !visitados.Pertenece(adyacente) {
			visitados.Guardar(adyacente, nil)
			padre.Guardar(adyacente, &vertice_inicial)
			dfs(grafo, adyacente, padre, visitados)
		}
	}
}

func EsBipartito[T comparable](grafo Grafo[T]) bool {
	colores := TDAHash.CrearHash[T, int]()
	for _, vertice := range grafo.ObtenerVertices() {
		if !colores.Pertenece(vertice) {
			if !esBipartito(grafo, vertice, colores) {
				return false
			}
		}
	}
	return true
}

func esBipartito[T comparable](grafo Grafo[T], vertice T, colores TDAHash.Diccionario[T, int]) bool {
	cola := TDACola.CrearColaEnlazada[T]()
	cola.Encolar(vertice)
	colores.Guardar(vertice, 0)
	for !cola.EstaVacia() {
		v := cola.Desencolar()
		for _, adyacente := range grafo.Adyacentes(v) {
			if colores.Pertenece(adyacente) {
				if colores.Obtener(adyacente) == colores.Obtener(v) {
					return false
				}
			} else {
				colores.Guardar(adyacente, colores.Obtener(v)-1)
				cola.Encolar(adyacente)
			}
		}
	}
	return true
}

func OrdenTopologico[T comparable](grafo Grafo[T]) []T {
	gradosEntrada := TDAHash.CrearHash[T, int]()
	for _, vertice := range grafo.ObtenerVertices() {
		gradosEntrada.Guardar(vertice, 0)
		for _, adyacente := range grafo.Adyacentes(vertice) {
			if !gradosEntrada.Pertenece(adyacente) {
				gradosEntrada.Guardar(adyacente, 1)
			} else {
				gradosEntrada.Guardar(adyacente, gradosEntrada.Obtener(adyacente)+1)
			}
		}
	}

	cola := TDACola.CrearColaEnlazada[T]()
	for _, vertice := range grafo.ObtenerVertices() {
		if gradosEntrada.Obtener(vertice) == 0 {
			cola.Encolar(vertice)
		}
	}

	salida := []T{}

	for !cola.EstaVacia() {
		v := cola.Desencolar()
		salida = append(salida, v)
		for _, adyacente := range grafo.Adyacentes(v) {
			gradosEntrada.Guardar(adyacente, gradosEntrada.Obtener(adyacente)-1)
			if gradosEntrada.Obtener(adyacente) == 0 {
				cola.Encolar(adyacente)
			}
		}
	}
	return salida
}

func CaminoMinimo[T comparable](origen T, grafo Grafo[T]) (TDAHash.Diccionario[T, T], TDAHash.Diccionario[T, int]) {
	var NADIE T
	distancia := TDAHash.CrearHash[T, int]()
	padre := TDAHash.CrearHash[T, T]()
	visitados := TDAHash.CrearHash[T, bool]()

	for _, vertice := range grafo.ObtenerVertices() {
		distancia.Guardar(vertice, int(^uint(0)>>1))
	}

	distancia.Guardar(origen, 0)
	padre.Guardar(origen, NADIE)
	visitados.Guardar(origen, true)

	cola := TDACola.CrearColaEnlazada[T]()
	cola.Encolar(origen)

	for !cola.EstaVacia() {
		v := cola.Desencolar()
		for _, adyacente := range grafo.Adyacentes(v) {
			if !visitados.Pertenece(adyacente) {
				distancia.Guardar(adyacente, distancia.Obtener(v)+1)
				padre.Guardar(adyacente, v)
				visitados.Guardar(adyacente, true)
				cola.Encolar(adyacente)
			}
		}
	}
	return padre, distancia
}

type dist[T comparable] struct {
	vertice T
	peso    int
}

func CaminoMinimoDijkstra[T comparable](origen T, grafo Grafo[T]) (TDAHash.Diccionario[T, T], TDAHash.Diccionario[T, int]) {
	var NADIE T
	distancia := TDAHash.CrearHash[T, int]()
	padre := TDAHash.CrearHash[T, T]()

	for _, vertice := range grafo.ObtenerVertices() {
		distancia.Guardar(vertice, int(^uint(0)>>1))
	}

	distancia.Guardar(origen, 0)
	padre.Guardar(origen, NADIE)

	heap := TDAHeap.CrearHeap(func(a, b dist[T]) int { return b.peso - a.peso })
	heap.Encolar(dist[T]{origen, 0})

	for !heap.EstaVacia() {
		v := heap.Desencolar().vertice
		for _, adyacente := range grafo.Adyacentes(v) {
			if distancia.Obtener(v)+grafo.Peso(v, adyacente) < distancia.Obtener(adyacente) {
				distancia.Guardar(adyacente, distancia.Obtener(v)+grafo.Peso(v, adyacente))
				padre.Guardar(adyacente, v)
				heap.Encolar(dist[T]{adyacente, distancia.Obtener(adyacente)})
			}
		}
	}

	resultado := []dist[T]{}
	for iter := distancia.Iterador(); iter.HaySiguiente(); {
		v, p := iter.VerActual()
		resultado = append(resultado, dist[T]{v, p})
		iter.Siguiente()
	}
	return padre, distancia
}

func Centralidad[T comparable](grafo Grafo[T]) []dist[T] {
	cent := TDAHash.CrearHash[T, int]()
	for _, vertice := range grafo.ObtenerVertices() {
		cent.Guardar(vertice, 0)
	}
	for _, v := range grafo.ObtenerVertices() {
		padre, _ := CaminoMinimo(v, grafo)
		for _, w := range grafo.ObtenerVertices() {
			if v == w {
				continue
			}
			if !padre.Pertenece(w) {
				continue
			}
			actual := padre.Obtener(w)
			for actual != v {
				cent.Guardar(actual, cent.Obtener(actual)+1)
				actual = padre.Obtener(actual)
			}
		}
	}
	resultado := []dist[T]{}
	for iter := cent.Iterador(); iter.HaySiguiente(); {
		v, p := iter.VerActual()
		resultado = append(resultado, dist[T]{v, p / 2})
		iter.Siguiente()
	}
	return resultado
}

func CantidadMinimaInversiones[T comparable](grafo Grafo[T], s, t T) int {
	grafoPesado := CrearGrafo(true, []T{})
	for _, vertice := range grafo.ObtenerVertices() {
		if !grafoPesado.Pertenece(vertice) {
			grafoPesado.AgregarVertice(vertice)
		}
		for _, adyacente := range grafo.Adyacentes(vertice) {
			if !grafoPesado.Pertenece(adyacente) {
				grafoPesado.AgregarVertice(adyacente)
			}
			grafoPesado.AgregarArista(vertice, adyacente, 0)
			if !grafo.ContieneArista(adyacente, vertice) {
				grafoPesado.AgregarArista(adyacente, vertice, 1)
			}
		}
	}
	_, camino := CaminoMinimoDijkstra(s, grafoPesado)
	return camino.Obtener(t)
}

type arista[T comparable] struct {
	origen  T
	destino T
	peso    int
}

func MST_Prim[T comparable](grafo Grafo[T]) Grafo[T] {
	origen := grafo.VerticeAleatorio()
	visitados := TDAHash.CrearHash[T, bool]()
	visitados.Guardar(origen, true)

	heap := TDAHeap.CrearHeap(func(a, b arista[T]) int { return b.peso - a.peso })
	for _, adyacente := range grafo.Adyacentes(origen) {
		heap.Encolar(arista[T]{origen, adyacente, grafo.Peso(origen, adyacente)})
	}

	arbol := CrearGrafo[T](false, grafo.ObtenerVertices())
	for _, vertice := range grafo.ObtenerVertices() {
		arbol.AgregarVertice(vertice)
	}

	for !heap.EstaVacia() {
		fmt.Println(heap)
		a := heap.Desencolar()
		if visitados.Pertenece(a.destino) {
			continue
		}
		fmt.Println(a)
		arbol.AgregarArista(a.origen, a.destino, a.peso)
		visitados.Guardar(a.destino, true)
		for _, adyacente := range grafo.Adyacentes(a.destino) {
			if !visitados.Pertenece(adyacente) {
				heap.Encolar(arista[T]{a.destino, adyacente, grafo.Peso(a.destino, adyacente)})
			}
		}
	}
	return arbol
}

func ObtenerAristas[T comparable](grafo Grafo[T]) []arista[T] {
	a := []arista[T]{}
	visitados := TDAHash.CrearHash[T, bool]()
	for _, vertice := range grafo.ObtenerVertices() {
		for _, adyacente := range grafo.Adyacentes(vertice) {
			if !visitados.Pertenece(adyacente) {
				a = append(a, arista[T]{vertice, adyacente, grafo.Peso(vertice, adyacente)})
			}
		}
		visitados.Guardar(vertice, true)
	}
	return a
}

func MST_Kruskal[T comparable](grafo Grafo[T]) Grafo[T] {
	conjuntos := TDAUnionFind.CrearUnionFind(grafo.ObtenerVertices())
	aristas := ObtenerAristas(grafo)
	sort.Slice(aristas, func(i, j int) bool { return aristas[i].peso < aristas[j].peso })
	arbol := CrearGrafo[T](false, grafo.ObtenerVertices())
	for _, arista := range aristas {

		if conjuntos.Find(arista.origen) == conjuntos.Find(arista.destino) {
			continue
		}
		arbol.AgregarArista(arista.origen, arista.destino, arista.peso)
		conjuntos.Union(arista.origen, arista.destino)
	}
	return arbol
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfsPuntosArticulacion[T comparable](grafo Grafo[T], v T, visitados TDAHash.Diccionario[T, bool], padre TDAHash.Diccionario[T, T], orden TDAHash.Diccionario[T, int], masBajo TDAHash.Diccionario[T, int], puntos TDAHash.Diccionario[T, T], esRaiz bool) {
	hijos := 0
	masBajo.Guardar(v, orden.Obtener(v))
	for _, w := range grafo.Adyacentes(v) {
		if !visitados.Pertenece(w) {
			hijos++
			orden.Guardar(w, orden.Obtener(v)+1)
			padre.Guardar(w, v)
			visitados.Guardar(w, true)
			dfsPuntosArticulacion(grafo, w, visitados, padre, orden, masBajo, puntos, false)

			if masBajo.Obtener(w) >= orden.Obtener(v) && !esRaiz {
				puntos.Guardar(v, v)
			}

			masBajo.Guardar(v, min(masBajo.Obtener(v), masBajo.Obtener(w)))
		} else if padre.Obtener(v) != w {
			masBajo.Guardar(v, min(masBajo.Obtener(v), orden.Obtener(w)))
		}
	}

	if esRaiz && hijos > 1 {
		puntos.Guardar(v, v)
	}
}

func PuntosArticulacion[T comparable](grafo Grafo[T]) TDAHash.Diccionario[T, T] {
	var NADIE T
	origen := grafo.VerticeAleatorio()
	fmt.Println(grafo.ObtenerVertices())
	visitados := TDAHash.CrearHash[T, bool]()
	padre := TDAHash.CrearHash[T, T]()
	orden := TDAHash.CrearHash[T, int]()
	masBajo := TDAHash.CrearHash[T, int]()
	puntosArticulacion := TDAHash.CrearHash[T, T]()

	visitados.Guardar(origen, true)
	padre.Guardar(origen, NADIE)
	orden.Guardar(origen, 0)

	dfsPuntosArticulacion(grafo, origen, visitados, padre, orden, masBajo, puntosArticulacion, true)

	return puntosArticulacion
}

func main() {
	/**
		grafo := CrearGrafo(false, vertices)
		grafo.AgregarArista("A", "B", 1)
		grafo.AgregarArista("A", "J", 1)
		grafo.AgregarArista("A", "H", 1)
		grafo.AgregarArista("A", "G", 1)
		grafo.AgregarArista("H", "I", 1)
		grafo.AgregarArista("I", "B", 1)
		grafo.AgregarArista("B", "F", 1)
		grafo.AgregarArista("B", "C", 1)
		grafo.AgregarArista("F", "E", 1)
		grafo.AgregarArista("C", "E", 1)
		grafo.AgregarArista("C", "D", 1)
		grafo.AgregarArista("J", "K", 1)
		grafo.AgregarArista("J", "Ñ", 1)
		grafo.AgregarArista("K", "L", 1)
		grafo.AgregarArista("L", "M", 1)
		grafo.AgregarArista("Ñ", "O", 1)

		materias := CrearGrafo(true, plan)
		materias.AgregarArista("Fisica I", "Fisica II", 1)
		materias.AgregarArista("Fisica II", "Estructuras del Computador", 1)
		materias.AgregarArista("Estructuras del Computador", "Organizacion de Datos", 1)
		materias.AgregarArista("Estructuras del Computador", "Taller", 1)
		materias.AgregarArista("Analisis Matematico II", "Fisica II", 1)
		materias.AgregarArista("Analisis Matematico II", "Probabilidad y Estadistica", 1)
		materias.AgregarArista("Analisis Matematico II", "Analisis Numerico", 1)
		materias.AgregarArista("Algebra II", "Probabilidad y Estadistica", 1)
		materias.AgregarArista("Algebra II", "Analisis Numerico", 1)
		materias.AgregarArista("Analisis Numerico", "Organizacion de Datos", 1)
		materias.AgregarArista("Algoritmos y Programacion I", "Algoritmos y Programacion II", 1)
		materias.AgregarArista("Algoritmos y Programacion II", "Algoritmos y Programacion III", 1)
		materias.AgregarArista("Algoritmos y Programacion II", "Analisis Numerico", 1)

		numeros := CrearGrafo(false, []int{})
		for i := 0; i <= 7; i++ {
			numeros.AgregarVertice(i)
		}
		numeros.AgregarArista(0, 1, 1)
		numeros.AgregarArista(1, 2, 1)
		numeros.AgregarArista(1, 5, 1)
		numeros.AgregarArista(4, 3, 1)
		numeros.AgregarArista(4, 7, 1)
		numeros.AgregarArista(5, 3, 1)
		numeros.AgregarArista(5, 4, 1)
		numeros.AgregarArista(6, 5, 1)
		numeros.AgregarArista(6, 7, 1)

		dijkstra := CrearGrafo(false, vertices)
		dijkstra.AgregarArista("A", "B", 6)
		dijkstra.AgregarArista("A", "C", 1)
		dijkstra.AgregarArista("B", "C", 2)
		dijkstra.AgregarArista("B", "D", 2)
		dijkstra.AgregarArista("B", "E", 5)
		dijkstra.AgregarArista("C", "D", 1)
		dijkstra.AgregarArista("D", "E", 5)

		cent := CrearGrafo(false, []string{})
		cent.AgregarVertice("Misiones")
		cent.AgregarVertice("Corrientes")
		cent.AgregarVertice("Cordoba")
		cent.AgregarVertice("CABA")
		cent.AgregarVertice("La Plata")
		cent.AgregarVertice("Mar del Plata")
		cent.AgregarVertice("Las Toninas")
		cent.AgregarArista("Misiones", "Corrientes", 1)
		cent.AgregarArista("Cordoba", "Corrientes", 1)
		cent.AgregarArista("Misiones", "Cordoba", 1)
		cent.AgregarArista("CABA", "Cordoba", 1)
		cent.AgregarArista("CABA", "Mar del Plata", 1)
		cent.AgregarArista("La Plata", "Mar del Plata", 1)

		invMinima := CrearGrafo(true, vertices)
		invMinima.AgregarArista("B", "A", 0)
		invMinima.AgregarArista("B", "C", 0)
		invMinima.AgregarArista("C", "D", 0)
		invMinima.AgregarArista("E", "C", 0)
		invMinima.AgregarArista("E", "F", 0)
		invMinima.AgregarArista("G", "F", 0)
		invMinima.AgregarArista("G", "H", 0)
		invMinima.AgregarArista("H", "I", 0)
		invMinima.AgregarArista("I", "J", 0)
		invMinima.AgregarArista("J", "A", 0)

		mst := CrearGrafo(false, vertices)
		mst.AgregarArista("A", "B", 3)
		mst.AgregarArista("A", "D", 1)
		mst.AgregarArista("C", "D", 2)
		mst.AgregarArista("C", "B", 1)
		mst.AgregarArista("D", "E", 1)
		mst.AgregarArista("D", "G", 2)
		mst.AgregarArista("I", "E", 1)
		mst.AgregarArista("I", "F", 5)
		mst.AgregarArista("G", "F", 3)
		mst.AgregarArista("H", "F", 1)
	**/
	articulacion := CrearGrafo(false, vertices)

	articulacion.AgregarArista("G", "A", 0)
	articulacion.AgregarArista("E", "A", 0)
	articulacion.AgregarArista("B", "A", 0)
	articulacion.AgregarArista("B", "C", 0)
	articulacion.AgregarArista("B", "D", 0)
	articulacion.AgregarArista("H", "D", 0)
	articulacion.AgregarArista("F", "A", 0)

	fmt.Println(PuntosArticulacion(articulacion))
}
