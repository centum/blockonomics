package blockonomics

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Invoice struct {
	Address     string `json:"addr"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Description string `json:"desc"`
	ExpiresAt   int64  `json:"expiry"`
	Timestamp   int64  `json:"timestamp"`
}

func (c *APIClient) Invoice(
	address string,
	amount float64,
	currency,
	description string,
	expiresAt time.Time) (urlCheckout string, err error) {

	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("%w: %s", ErrInternal, e)
		}
	}()

	passphrase := genPass(8)

	invoice := Invoice{
		Address:     address,
		Amount:      fmt.Sprintf("%.2f", amount),
		Currency:    currency,
		Description: description,
		ExpiresAt:   expiresAt.Unix(),
		Timestamp:   time.Now().Unix(),
	}

	data, err := json.Marshal(&invoice)
	if err != nil {
		return "", err
	}
	encryptedInvoice := Encrypt(data, passphrase)

	req, err := c.newRequest(
		http.MethodPost,
		"/api/invoice",
		&struct {
			Content string `json:"content"`
		}{string(encryptedInvoice)})
	if err != nil {
		return "", err
	}

	var resp struct {
		Number int64 `json:"number"`
	}
	err = c.send(req, &resp)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/invoice/%d/#/?key=%s", c.APIBase, resp.Number, string(passphrase)), nil
}
