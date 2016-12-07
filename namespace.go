package argparse

import "fmt"

// Namespace is a map of key-value pairs, used for storing pairings
// between options' destinations and their associated values. It will
// contain only `string` and `[]string` values.
type Namespace map[string]interface{}

// Get will retrieve either a string or a []string if the specified key
// exists in the mapping. Otherwise, an empty string is returned
func (n Namespace) Get(key string) interface{} {
	if n.KeyExists(key) != true {
		return nil
	}

	return n[key]
}

// KeyExists returns a bool indicating true if the key does exist in the mapping,
// or otherwise false.
func (n Namespace) KeyExists(key string) bool {
	if _, ok := n[key]; ok == true {
		return true
	}
	return false
}

// merge will take all the values from the other, provided Namespace and copy
// them to the current Namespace.
func (n *Namespace) merge(other *Namespace) {
	if other == nil {
		return
	}

	for key, value := range *other {
		(*n)[key] = value
	}
}

// Require will assert that all the specified keys exist in the namespace.
func (n Namespace) Require(keys ...string) error {
	for _, key := range keys {
		if n.KeyExists(key) != true {
			return fmt.Errorf("Missing option: %s", key)
		}
	}
	return nil
}

// Set will set the mapping's value at the desired key to the value provided.
func (n *Namespace) Set(key string, value interface{}) *Namespace {
	if n == nil {
		n = new(Namespace)
	}
	(*n)[key] = value

	return n
}

// Slice will retrieve either a string or a []string if the specified key
// exists in the mapping. Otherwise, an empty string is returned
func (n Namespace) Slice(key string) []string {
	if n.KeyExists(key) != true {
		return nil
	}

	slice, _ := n[key].([]string)
	return slice
}

// String will retrieve either a string or a []string if the specified key
// exists in the mapping. Otherwise, an empty string is returned
func (n Namespace) String(key string) string {
	if n.KeyExists(key) != true {
		return ""
	}

	return n[key].(string)
}

// Try will retrieve either a string or a []string if the specified key
// exists in the mapping. Otherwise, an error is returned.
func (n Namespace) Try(key string) (interface{}, error) {
	if n.KeyExists(key) != true {
		return nil, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}

	return n[key], nil
}

// Create a new pointer to an Namespace instance.
func NewNamespace() *Namespace {
	n := Namespace(make(map[string]interface{}))
	return &n
}
