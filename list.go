package gocharge

import "sync"

// List is a good way to use like csharp list
type List[T any] struct {
	//a slice of items
	Items []T
}

// create a new list struct
func NewList[T any](data ...[]T) *List[T] {
	var gclist = List[T]{}
	if len(data) == 1 {
		gclist.Items = data[0]
		return &gclist
	}

	gclist.Items = make([]T, 0)
	return &gclist
}

// add new items
func (lst *List[T]) AddRange(data []T) {
	lst.Items = append(lst.Items, data...)
}

// add new item
func (lst *List[T]) Add(data T) {
	lst.Items = append(lst.Items, data)
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
	// Estimate the capacity for the result slice
	estimatedCapacity := len(lst.Items) / 2
	result := make([]T, 0, estimatedCapacity)

	var mu sync.Mutex
	var wg sync.WaitGroup

	// Process the slice in parallel using goroutines
	for _, data := range lst.Items {
		wg.Add(1)
		go func(data T) {
			defer wg.Done()
			if predicate(data) {
				mu.Lock()
				result = append(result, data)
				mu.Unlock()
			}
		}(data)
	}

	wg.Wait()
	return *NewList(result)
}

// get the array list
func (lst *List[T]) ToArray() []T {
	return lst.Items
}

// a linq function which returns a predicated list
func WhereT[T any](lst *[]T, predicate func(T) bool) []T {

	// Estimate the capacity for the result slice
	estimatedCapacity := len(*lst) / 2
	result := make([]T, 0, estimatedCapacity)

	var mu sync.Mutex
	var wg sync.WaitGroup

	// Process the slice in parallel using goroutines
	for _, data := range *lst {
		wg.Add(1)
		go func(data T) {
			defer wg.Done()
			if predicate(data) {
				mu.Lock()
				result = append(result, data)
				mu.Unlock()
			}
		}(data)
	}

	wg.Wait()
	return result
}
