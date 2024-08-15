package gocharge

// List is a good way to use like csharp list
type List[T any] struct {
	items []T
}

func NewList[T any](data ...[]T) *List[T] {
	var gclist = List[T]{}
	if len(data) == 1 {
		gclist.items = data[0]
	}
	gclist.items = make([]T, 0)
	return &gclist
}

func (lst *List[T]) AddItems(data []T) {
	lst.items = append(lst.items, data...)
}

func (lst *List[T]) AddItem(data T) {
	lst.items = append(lst.items, data)
}

func (lst *List[T]) Contains(data T) (bool, error) {
	//todo
	return true, nil
}

func (lst *List[T]) Sum() any {
	//todo
	return nil
}

func (lst *List[T]) Where(predicate func(T) bool) List[T] {

	var result = *NewList[T]()

	for _, data := range lst.items {
		if predicate(data) {
			result.AddItem(data)
		}
	}
	return result
}

func WhereT[T any](lst *[]T, predicate func(T) bool) []T {

	var result = make([]T, 0)

	for _, data := range *lst {
		if predicate(data) {
			result = append(result, data)
		}
	}
	return result
}
