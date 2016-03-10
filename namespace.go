// Package argparse provides functionallity to emulate Python's argparse for
// setting-up and parsing a program's options & arguments.
package argparse

import "fmt"

// Namespace is a struct for storing the key-value pairings between
// options' destinations and their associated values.
type Namespace struct {
	Mapping map[string]interface{}
}

// Get will retrieve either a string or a []string if the specified key
// exists in the mapping. Otherwise, an error is returned.
func (n *Namespace) Get(key string) (interface{}, error) {
	if n.KeyExists(key) != true {
		return nil, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}

	return n.Mapping[key], nil
}

// KeyExists returns a bool indicating true if the key does exist in the mapping,
// or otherwise false.
func (n *Namespace) KeyExists(key string) bool {
	if _, ok := n.Mapping[key]; ok == true {
		return true
	}
	return false
}

// Set will set the mapping's value at the desired key to the value provided.
func (n *Namespace) Set(key string, value interface{}) *Namespace {
	n.Mapping[key] = value

	return n
}

// Create a new pointer to an Namespace instance.
func NewNamespace() *Namespace {
	n := new(Namespace)
	n.Mapping = make(map[string]interface{})
	return n
}
