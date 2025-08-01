package pkg

import "encoding"

// An EncodeHasher is an object that can serialize and deserialize, and produce a unique hash
type EncodeHasher interface {
	Hash() string // unique
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

// A Store is a general purpose key-value store
type Store[K comparable, V EncodeHasher] interface {
	Get(K) (V, error)
	Insert(V) error
	Delete(K) error
	GetAll() map[K]V
}
