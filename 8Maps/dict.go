package main

import "errors"

// Made a new type Dictionary as map
type Dictionary map[string]string

// Initializing an error
var ErrKeyAbsent = errors.New("given key is not present")

// func Search(dict Dictionary, key string) string {
// 	return dict[key]
// }

// Making Search a method to Dictionary
func (d Dictionary) Search(key string) (string, error) {

	// This approach is used to check whether a specific key is present in map or not
	// here ok variable stores a bool val representing if the key is present in the map or not
	value, ok := d[key]

	if ok {
		return value, nil
	}

	return "", ErrKeyAbsent
}

func (d Dictionary) Add(key, value string)  {
	// Adding to a map is similar to adding to an array
	d[key] = value
}