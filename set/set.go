package set

type Set[T comparable] interface {
	Add(T) Set[T]
	Has(T) bool
}

func New[T comparable](values ...T) Set[T] {
	s := set[T]{
		data: make(map[T]any, len(values)),
	}
	for _, v := range values {
		s.Add(v)
	}
	return &s
}

type set[T comparable] struct{
	data map[T]any
}

func (s *set[T]) Add(v T) Set[T] {
	s.data[v] = struct{}{}
	return s
}

func (s *set[T]) Has(v T) bool {
	_, ok := s.data[v]
	return ok
}
