package list

type List[T any] struct {
	slice []T
}

func NewList[T comparable](lst []T) *List[T] {
	return &List[T]{lst}
}

func (lst *List[T]) Contains(data T) (bool, error) {
	//todo
	return true, nil
}

func (lst *List[T]) Sum() any {
	//todo
	return nil
}
