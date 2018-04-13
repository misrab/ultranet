package store

import (
	"testing"

	"io/ioutil"
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

func TestDatabasePut(t *testing.T) {
	db, _ := NewDatabase()

	// open test file
	file, err := ioutil.ReadFile(TEST_DATA_URL)
	if err != nil {
		t.Fatalf("Could not read file: %v\n", err)
	}

	db.Put(file)

}
