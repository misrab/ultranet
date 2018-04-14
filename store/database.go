package store

import (
	"github.com/multiformats/go-multihash"
)

type Blob []byte

type Database interface {
	Get(multihash.Multihash) (Blob, error)
	Put(Blob) (multihash.Multihash, error)

	Remove(multihash.Multihash) error
}

func NewDatabase() (Database, error) {
	db := new(database)

	return db, nil
}

type database struct {
}

func (db *database) Get(hash multihash.Multihash) (Blob, error) {
	return nil, nil
}

func (db *database) Put(Blob) (multihash.Multihash, error) {
	// hash the file

	return nil, nil
}

func (db *database) Remove(multihash.Multihash) error {
	// TODO
	return nil
}
