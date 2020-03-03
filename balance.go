package blockonomics

import (
	"net/http"
	"strings"
)

type Balance struct {
	Addr        string `json:"addr"`
	Confirmed   int64  `json:"confirmed"`
	Unconfirmed int64  `json:"unconfirmed"`
}

func (c *APIClient) Balance(addrs ...string) (balanceList []Balance, err error) {

	req, err := c.newRequest(http.MethodPost, "/api/balance",
		&struct {
			Addr string `json:"addr"`
		}{strings.Join(addrs, " ")})
	if err != nil {
		return nil, err
	}

	var data struct {
		Response []Balance `json:"response"`
	}
	err = c.send(req, &data)
	if err != nil {
		return nil, err
	}

	return data.Response, nil
}
