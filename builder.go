package httpx

type builder interface {
	Add(key, value string)
	Set(key, value string)
	Del(key string)
}

type Builder[T builder] struct {
	builder T
}

func NewBuilder[T builder](builder T) Builder[T] {
	return Builder[T]{builder}
}

func (h Builder[T]) Add(key, value string) Builder[T] {
	h.builder.Add(key, value)
	return h
}

func (h Builder[T]) Set(key, value string) Builder[T] {
	h.builder.Set(key, value)
	return h
}

func (h Builder[T]) Del(key string) Builder[T] {
	h.builder.Del(key)
	return h
}

func (h Builder[T]) Build() T {
	return h.builder
}
