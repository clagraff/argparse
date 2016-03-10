// Package argparse provides functionallity to emulate Python's argparse for
// setting-up and parsing a program's options & arguments.
package argparse

import "fmt"

type Namespace struct {
	Mapping map[string]interface{}
	Keys    []string
}

func (n *Namespace) Bool(key string) (bool, error) {
	if n.KeyExists(key) != true {
		return false, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}
	if val, ok := n.Mapping[key].(bool); ok != true {
		return false, fmt.Errorf("Invalid type")
	} else {
		return val, nil
	}
}

func (n *Namespace) Float32(key string) (float32, error) {
	if n.KeyExists(key) != true {
		return 0.0, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}
	if val, ok := n.Mapping[key].(float32); ok != true {
		return 0.0, fmt.Errorf("Invalid type")
	} else {
		return val, nil
	}
}

func (n *Namespace) Float64(key string) (float64, error) {
	if n.KeyExists(key) != true {
		return 0.0, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}
	if val, ok := n.Mapping[key].(float64); ok != true {
		return 0.0, fmt.Errorf("Invalid type")
	} else {
		return val, nil
	}
}

func (n *Namespace) Get(key string) (interface{}, error) {
	if n.KeyExists(key) != true {
		return nil, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}

	return n.Mapping[key], nil
}

func (n *Namespace) Int(key string) (int, error) {
	if n.KeyExists(key) != true {
		return 0, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}
	if val, ok := n.Mapping[key].(int); ok != true {
		return 0, fmt.Errorf("Invalid type")
	} else {
		return val, nil
	}
}

func (n *Namespace) Int8(key string) (int8, error) {
	if n.KeyExists(key) != true {
		return 0, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}
	if val, ok := n.Mapping[key].(int8); ok != true {
		return 0, fmt.Errorf("Invalid type")
	} else {
		return val, nil
	}
}

func (n *Namespace) Int16(key string) (int16, error) {
	if n.KeyExists(key) != true {
		return 0, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}
	if val, ok := n.Mapping[key].(int16); ok != true {
		return 0, fmt.Errorf("Invalid type")
	} else {
		return val, nil
	}
}

func (n *Namespace) Int32(key string) (int32, error) {
	if n.KeyExists(key) != true {
		return 0, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}
	if val, ok := n.Mapping[key].(int32); ok != true {
		return 0, fmt.Errorf("Invalid type")
	} else {
		return val, nil
	}
}

func (n *Namespace) Int64(key string) (int64, error) {
	if n.KeyExists(key) != true {
		return 0, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}
	if val, ok := n.Mapping[key].(int64); ok != true {
		return 0, fmt.Errorf("Invalid type")
	} else {
		return val, nil
	}
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

func (n *Namespace) String(key string) (string, error) {
	if n.KeyExists(key) != true {
		return "", fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}
	if val, ok := n.Mapping[key].(string); ok != true {
		return "", fmt.Errorf("Invalid type")
	} else {
		return val, nil
	}
}

func (n *Namespace) Uint(key string) (uint, error) {
	if n.KeyExists(key) != true {
		return 0, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}
	if val, ok := n.Mapping[key].(uint); ok != true {
		return 0, fmt.Errorf("Invalid type")
	} else {
		return val, nil
	}
}

func (n *Namespace) Uint8(key string) (uint8, error) {
	if n.KeyExists(key) != true {
		return 0, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}
	if val, ok := n.Mapping[key].(uint8); ok != true {
		return 0, fmt.Errorf("Invalid type")
	} else {
		return val, nil
	}
}

func (n *Namespace) Uint16(key string) (uint16, error) {
	if n.KeyExists(key) != true {
		return 0, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}
	if val, ok := n.Mapping[key].(uint16); ok != true {
		return 0, fmt.Errorf("Invalid type")
	} else {
		return val, nil
	}
}

func (n *Namespace) Uint32(key string) (uint32, error) {
	if n.KeyExists(key) != true {
		return 0, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}
	if val, ok := n.Mapping[key].(uint32); ok != true {
		return 0, fmt.Errorf("Invalid type")
	} else {
		return val, nil
	}
}

func (n *Namespace) Uint64(key string) (uint64, error) {
	if n.KeyExists(key) != true {
		return 0, fmt.Errorf("Key \"%s\" does not exist in namespace", key)
	}
	if val, ok := n.Mapping[key].(uint64); ok != true {
		return 0, fmt.Errorf("Invalid type")
	} else {
		return val, nil
	}
}

func NewNamespace() *Namespace {
	n := new(Namespace)
	n.Mapping = make(map[string]interface{})
	return n
}
