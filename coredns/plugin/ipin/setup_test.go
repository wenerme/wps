package ipin

import (
	"testing"

	"github.com/mholt/caddy"
)

func TestSetupWhoami(t *testing.T) {
	c := caddy.NewTestController("dns", `ipin`)
	if err := setup(c); err != nil {
		t.Fatalf("Expected no errors, but got: %v", err)
	}

	c = caddy.NewTestController("dns", `ipin example.org`)
	if err := setup(c); err == nil {
		t.Fatalf("Expected errors, but got: %v", err)
	}
}
