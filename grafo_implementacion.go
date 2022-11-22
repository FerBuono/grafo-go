package main

import (
	TDAHash "grafo/hash"
	"math/rand"
)

type grafo[T comparable] struct {
	dicc     TDAHash.Diccionario[T, TDAHash.Diccionario[T, int]]
	dirigido bool
}

type iterGrafo[T comparable] struct {
	iterDicc TDAHash.IterDiccionario[T, TDAHash.Diccionario[T, int]]
}

func CrearGrafo[T comparable](es_dirigido bool, lista_vertices []T) Grafo[T] {
	grafo := new(grafo[T])
	grafo.dicc = TDAHash.CrearHash[T, TDAHash.Diccionario[T, int]]()
	if len(lista_vertices) > 0 {
		for _, vertice := range lista_vertices {
			grafo.dicc.Guardar(vertice, TDAHash.CrearHash[T, int]())
		}
	}
	grafo.dirigido = es_dirigido
	return grafo
}

func (g *grafo[T]) AgregarVertice(v T) {
	g.dicc.Guardar(v, TDAHash.CrearHash[T, int]())
}

func (g *grafo[T]) EliminarVertice(v T) {
	if !g.Pertenece(v) {
		panic("El vertice no pertenece al grafo")
	}
	g.dicc.Borrar(v)
	for iter := g.dicc.Iterador(); iter.HaySiguiente(); {
		_, diccAdyacentes := iter.VerActual()
		diccAdyacentes.Borrar(v)
		iter.Siguiente()
	}
}

func (g *grafo[T]) AgregarArista(v1 T, v2 T, peso int) {
	if !g.Pertenece(v1) || !g.Pertenece(v2) {
		panic("Un vertice no pertenece al grafo")
	}
	diccAdyacentes := g.dicc.Obtener(v1)
	diccAdyacentes.Guardar(v2, peso)
	if !g.dirigido {
		diccAdyacentes = g.dicc.Obtener(v2)
		diccAdyacentes.Guardar(v1, peso)
	}
}

func (g *grafo[T]) EliminarArista(v1 T, v2 T) {
	if !g.Pertenece(v1) || !g.Pertenece(v2) {
		panic("Un vertice no pertenece al grafo")
	}
	diccAdyacentes := g.dicc.Obtener(v1)
	diccAdyacentes.Borrar(v2)
	if !g.dirigido {
		diccAdyacentes = g.dicc.Obtener(v2)
		diccAdyacentes.Borrar(v1)
	}
}

func (g *grafo[T]) Peso(v1, v2 T) int {
	if !g.Pertenece(v1) || !g.Pertenece(v2) {
		panic("Un vertice no pertenece al grafo")
	}
	return g.dicc.Obtener(v1).Obtener(v2)
}

func (g *grafo[T]) Pertenece(v T) bool {
	return g.dicc.Pertenece(v)
}

func (g *grafo[T]) ObtenerVertices() []T {
	vertices := []T{}
	for iter := g.dicc.Iterador(); iter.HaySiguiente(); {
		vertices = append(vertices, iter.Siguiente())
	}
	return vertices
}

func (g *grafo[T]) Adyacentes(v T) []T {
	adyacentes := []T{}
	diccAdyacentes := g.dicc.Obtener(v)
	for iter := diccAdyacentes.Iterador(); iter.HaySiguiente(); {
		adyacentes = append(adyacentes, iter.Siguiente())
	}
	return adyacentes
}

func (g *grafo[T]) VerticeAleatorio() T {
	vertices := g.ObtenerVertices()
	return vertices[rand.Intn(len(vertices))]
}

func (g *grafo[T]) ContieneArista(v1, v2 T) bool {
	if !g.Pertenece(v1) || !g.Pertenece(v2) {
		panic("Un vertice no pertenece al grafo")
	}
	return g.dicc.Obtener(v1).Pertenece(v2)
}

func (g *grafo[T]) Iterador() IteradorGrafo[T] {
	iter := new(iterGrafo[T])
	iter.iterDicc = g.dicc.Iterador()
	return iter
}

func (i *iterGrafo[T]) VerActual() T {
	vertice, _ := i.iterDicc.VerActual()
	return vertice
}

func (i *iterGrafo[T]) HaySiguiente() bool {
	return i.iterDicc.HaySiguiente()
}

func (i *iterGrafo[T]) Siguiente() T {
	return i.iterDicc.Siguiente()
}
