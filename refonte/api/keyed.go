package api

import "iter"

// Keyed is the interface that all keyed records must implement.
type Keyed interface {
	// ValueForKey returns the value for the given key.
	ValueForKey(key string) (value any, has bool)

	// SetValueForKey sets the value for the given key.
	SetValueForKey(key string, value any)

	// RemoveValueForKey removes the value for the given key.
	RemoveValueForKey(key string)

	// HasKey checks if the node has the given key.
	HasKey(key string) bool

	// Keys returns all keys of the node.
	Keys() []string

	// KeyValuePairs returns an iterator over the key-value pairs in the container.
	KeyValuePairs() iter.Seq2[string, any]
}
