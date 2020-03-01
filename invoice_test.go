package blockonomics

import (
	"testing"
	"time"
)

func TestInvoice(t *testing.T) {
	b := NewTestBlockonomicsHandler()

	c := NewClient(b.token, WithTimeout(time.Duration(30)*time.Second))
	if c == nil {
		t.Fatal("client is nil")
	}
	c.APIBase = b.server.URL

	t.Run("get invoice", func(t *testing.T) {
		addr := string(genPass(16))
		amount := 10.50
		currency := "USD"
		description := "description transaction"
		expiresAt := time.Now().Add(30 * time.Minute)
		urlCheckout, err := c.Invoice(addr, amount, currency, description, expiresAt)
		if err != nil {
			t.Error("error is not nil")
		}
		if urlCheckout == "" {
			t.Fatal("urlCheckout is empty")
		}
	})

}
