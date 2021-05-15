package ipin

import (
	"testing"

	"github.com/coredns/caddy"
)

func TestSetupWhoami(t *testing.T) {
	c := caddy.NewTestController("dns", `ipin`)
	if err := setup(c); err != nil {
		t.Fatalf("Expected no errors, but got: %v", err)
	}

	c = caddy.NewTestController("dns", `ipin {
fallback
}`)
	if err := setup(c); err != nil {
		t.Fatalf("Expected no errors, but got: %v", err)
	}

	c = caddy.NewTestController("dns", `ipin example.org`)
	if err := setup(c); err == nil {
		t.Fatalf("Expected errors, but got: %v", err)
	}

	c = caddy.NewTestController("dns", `ipin {
fallback invalid
}`)
	if err := setup(c); err == nil {
		t.Fatalf("Expected errors, but got: %v", err)
	}
}
