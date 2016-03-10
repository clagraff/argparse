// Package argparse provides functionallity to emulate Python's argparse for
// setting-up and parsing a program's options & arguments.
package argparse

import "fmt"

type Namespace struct {
	Mapping map[string]interface{}
	Keys    []string
}

func (n *Namespace) Get(key string) (interface{}, error) {
	if n.KeyExists(key) != true {
		return nil, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}

	return n.Mapping[key], nil
}
func (n *Namespace) KeyExists(key string) bool {
	for _, s := range n.Keys {
		if s == key {
			return true
		}
	}

	return false
}

func (n *Namespace) Set(key string, value interface{}) *Namespace {
	if _, ok := n.Mapping[key]; !ok {
		n.Keys = append(n.Keys, key)
	}
	n.Mapping[key] = value

	return n
}

func NewNamespace() *Namespace {
	n := new(Namespace)
	n.Mapping = make(map[string]interface{})
	return n
}
