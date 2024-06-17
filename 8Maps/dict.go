package main

import "errors"

// Made a new type Dictionary as map
type Dictionary map[string]string

// Initializing an error
var KeyAbsent = errors.New("given key is not present")

// func Search(dict Dictionary, key string) string {
// 	return dict[key]
// }

// Making Search a method to Dictionary
func (d Dictionary) Search(key string) (string, error) {
	return d[key], nil
}