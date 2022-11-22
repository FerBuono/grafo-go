package main

type Grafo[T comparable] interface {
	AgregarVertice(T)

	EliminarVertice(T)

	AgregarArista(T, T, int)

	EliminarArista(T, T)

	Peso(T, T) int

	Pertenece(T) bool

	ObtenerVertices() []T

	Adyacentes(T) []T

	ContieneArista(T, T) bool

	VerticeAleatorio() T

	Iterador() IteradorGrafo[T]
}

type IteradorGrafo[T comparable] interface {
	VerActual() T

	HaySiguiente() bool

	Siguiente() T
}
