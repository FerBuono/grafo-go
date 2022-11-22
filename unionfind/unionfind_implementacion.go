package unionfind

import (
	TDAHash "grafo/hash"
)

type unionfind[T comparable] struct {
	grupos TDAHash.Diccionario[T, T]
}

func CrearUnionFind[T comparable](vertices []T) UnionFind[T] {
	u := new(unionfind[T])
	u.grupos = TDAHash.CrearHash[T, T]()
	for _, vertice := range vertices {
		u.grupos.Guardar(vertice, vertice)
	}
	return u
}

func (u *unionfind[T]) Find(vertice T) T {
	if u.grupos.Obtener(vertice) == vertice {
		return vertice
	}

	realGroup := u.Find(u.grupos.Obtener(vertice))
	u.grupos.Guardar(vertice, realGroup)
	return realGroup
}

func (u *unionfind[T]) Union(v1, v2 T) {
	newGroup := u.Find(v1)
	other := u.Find(v2)
	u.grupos.Guardar(other, newGroup)
}
