package store

import (
	"testing"

	"io/ioutil"

	"github.com/stretchr/testify/assert"
)

const (
	TEST_DATA_URL = "../test_data/test.txt"
)

func TestNewDatabase(t *testing.T) {
	_, err := NewDatabase()

	if err != nil {
		t.Fatalf("Could not open database: %v\n", err)
	}
}

/*
func TestDatabasePut(t *testing.T) {
	db, _ := NewDatabase()

	// open test file
	file, err := ioutil.ReadFile(TEST_DATA_URL)
	if err != nil {
		t.Fatalf("Could not read file: %v\n", err)
	}

	_, err = db.Put(file)
	if err != nil {
		t.Fatalf("Could not put: %v\n", err)
	}
}
*/
func TestDatabasePutGetRemove(t *testing.T) {
	db, _ := NewDatabase()

	// insert test file
	file, err := ioutil.ReadFile(TEST_DATA_URL)
	if err != nil {
		t.Fatalf("Could not read file: %v\n", err)
	}

	hash, err := db.Put(file)
	if err != nil {
		t.Fatalf("Could not put: %v\n", err)
	}

	// retrieve it
	blob, err := db.Get(hash)
	if err != nil {
		t.Fatalf("Could not get: %v\n", err)
	}
	assert.Equal(t, file, blob, "The file inserted did not match the file retrieved")

	// remove the file
	// db.Remove(hash)
}
