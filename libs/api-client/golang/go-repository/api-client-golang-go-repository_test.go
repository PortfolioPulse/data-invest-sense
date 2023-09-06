package api_client_golang_go_repository

import (
	"testing"
)

func TestGoRepository(t *testing.T) {
	result := GoRepository("works")
	if result != "GoRepository works" {
		t.Error("Expected GoRepository to append 'works'")
	}
}
