package main

import (
	"bytes"
	"testing"
)

func TestStore(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}
	s := NewStore(opts)
	buf := bytes.NewReader([]byte("fdsfds"))

	
}
