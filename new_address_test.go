package blockonomics

import (
	"testing"
	"time"
)

// nolint:gocyclo
func TestNewAddress(t *testing.T) {
	b := NewTestBlockonomicsHandler()

	c := NewClient(b.token, WithTimeout(time.Duration(30)*time.Second))
	if c == nil {
		t.Fatal("client is nil")
	}
	c.APIBase = b.server.URL

	t.Run("new address", func(t *testing.T) {
		a, err := c.NewAddress("", false)
		if err != nil {
			t.Error("error is not nil")
		}
		if a == nil {
			t.Fatal("address is nil")
		}
		if a.Address == "" {
			t.Fatal("address is empty")
		}
	})

	t.Run("new address with reset", func(t *testing.T) {
		a, err := c.NewAddress("", true)
		if err != nil {
			t.Error("error is not nil")
		}
		if a == nil {
			t.Fatal("address is nil")
		}
		if a.Address == "" {
			t.Fatal("address is empty")
		}
		if a.Reset != 1 {
			t.Fatal("Reset is not equal 1")
		}
	})

	t.Run("new address with another account", func(t *testing.T) {
		a, err := c.NewAddress("xpub6D9qFCtaxyyP3aAMy6D9qFCtaxyyP3aAMy", false)
		if err != nil {
			t.Error("error is not nil")
		}
		if a == nil {
			t.Fatal("address is nil")
		}
		if a.Address == "" {
			t.Fatal("address is empty")
		}
		if a.Account == "" {
			t.Fatal("Account is empty")
		}
	})
}
