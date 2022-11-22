package unionfind

type UnionFind[T comparable] interface {
	Find(T) T

	Union(T, T)
}
