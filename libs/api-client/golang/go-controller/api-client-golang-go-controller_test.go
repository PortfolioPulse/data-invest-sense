package api_client_golang_go_controller

import (
	"testing"
)

func TestGoController(t *testing.T) {
	result := GoController("works")
	if result != "GoController works" {
		t.Error("Expected GoController to append 'works'")
	}
}
