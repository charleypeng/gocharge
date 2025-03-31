package gocharge

import (
	"fmt"
	"sync"
)

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

// Count returns the number of elements in the list
func (lst *List[T]) Count() int {
	return len(lst.Items)
}

// Clear removes all elements from the list
func (lst *List[T]) Clear() {
	lst.Items = make([]T, 0)
}

// Contains checks if an element is in the list
func (lst *List[T]) Contains(data T) bool {
	for _, item := range lst.Items {
		if any(item) == any(data) {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the first occurrence of the specified element
func (lst *List[T]) IndexOf(data T) int {
	for i, item := range lst.Items {
		if any(item) == any(data) {
			return i
		}
	}
	return -1
}

// Insert adds an element at the specified index
func (lst *List[T]) Insert(index int, data T) error {
	if index < 0 || index > len(lst.Items) {
		return fmt.Errorf("index out of range")
	}

	lst.Items = append(lst.Items[:index], append([]T{data}, lst.Items[index:]...)...)
	return nil
}

// Remove removes the first occurrence of the specified element
func (lst *List[T]) Remove(data T) bool {
	index := lst.IndexOf(data)
	if index == -1 {
		return false
	}

	lst.Items = append(lst.Items[:index], lst.Items[index+1:]...)
	return true
}

// RemoveAt removes the element at the specified index
func (lst *List[T]) RemoveAt(index int) error {
	if index < 0 || index >= len(lst.Items) {
		return fmt.Errorf("index out of range")
	}

	lst.Items = append(lst.Items[:index], lst.Items[index+1:]...)
	return nil
}

// Find returns the first element that matches the predicate
func (lst *List[T]) Find(predicate func(T) bool) (T, bool) {
	for _, item := range lst.Items {
		if predicate(item) {
			return item, true
		}
	}
	var zero T
	return zero, false
}

// FindIndex returns the index of the first element that matches the predicate
func (lst *List[T]) FindIndex(predicate func(T) bool) int {
	for i, item := range lst.Items {
		if predicate(item) {
			return i
		}
	}
	return -1
}

// ForEach performs the specified action on each element
func (lst *List[T]) ForEach(action func(T)) {
	for _, item := range lst.Items {
		action(item)
	}
}

// a linq function which returns a predicated list
func (lst *List[T]) Where(predicate func(T) bool) *List[T] {
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
	return NewList(result)
}

// get the array list
func (lst *List[T]) ToArray() *[]T {
	return &lst.Items
}

// a linq function which returns a predicated list
func WhereT[T any](lst *[]T, predicate func(T) bool) *[]T {

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
	return &result
}
