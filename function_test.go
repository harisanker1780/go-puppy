package puppy

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	got := add(20, 10)
	want := 30
	if want != got {
		log.Fatalf("error - want %v and got %v", want, got)
	}
}
