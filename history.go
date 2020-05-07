package blockonomics

import (
	"net/http"
	"strings"
)

type Transaction struct {
	TxID  string   `json:"txid"`
	Addr  []string `json:"addr"`
	Value int64    `json:"value"`
	Time  int64    `json:"time"`
}

type PendingTransaction struct {
	Transaction
	Status int `json:"status"`
}

type History struct {
	Pending []PendingTransaction `json:"pending"`
	History []Transaction        `json:"history"`
}

func (c *APIClient) SearchHistory(addrs ...string) (data *History, err error) {

	req, err := c.newRequest(http.MethodPost, "/api/searchhistory",
		&struct {
			Addr string `json:"addr"`
		}{strings.Join(addrs, " ")})
	if err != nil {
		return nil, err
	}

	c.auth(req)

	err = c.send(req, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
