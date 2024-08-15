package gocharge

// List is a good way to use like csharp list
type List[T any] struct {
	items []T
}

// create a new list struct
func NewList[T any](data ...[]T) *List[T] {
	var gclist = List[T]{}
	if len(data) == 1 {
		gclist.items = data[0]
		return &gclist
	}

	gclist.items = make([]T, 0)
	return &gclist
}

// add new items
func (lst *List[T]) AddItems(data []T) {
	lst.items = append(lst.items, data...)
}

// add new item
func (lst *List[T]) AddItem(data T) {
	lst.items = append(lst.items, data)
}

// returns bool if the given element contains the given data
func (lst *List[T]) Contains(data T) (bool, error) {
	//todo
	return true, nil
}

// sum all the data if it is comparable
func (lst *List[T]) Sum() any {
	//todo
	return nil
}

// a linq function which returns a predicated list
func (lst *List[T]) Where(predicate func(T) bool) List[T] {

	var result = *NewList[T]()

	for _, data := range lst.items {
		if predicate(data) {
			result.AddItem(data)
		}
	}
	return result
}

// get the array list
func (lst *List[T]) ToArray() []T {
	return lst.items
}

// a linq function which returns a predicated list
func WhereT[T any](lst *[]T, predicate func(T) bool) []T {

	var result = make([]T, 0)

	for _, data := range *lst {
		if predicate(data) {
			result = append(result, data)
		}
	}
	return result
}
