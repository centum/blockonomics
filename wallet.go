package blockonomics

import (
	"net/http"
)

type Address struct {
	Address   string  `json:"address"`
	Balance   int64   `json:"balance"`
	Tag       string  `json:"tag"`
	CreatedOn float64 `json:"createdon"`
}

func (c *APIClient) AddrMonList() ([]Address, error) {

	req, err := c.newRequest(http.MethodGet, "/api/address", nil)
	if err != nil {
		return nil, err
	}
	c.auth(req)

	var data []Address
	err = c.send(req, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *APIClient) AddrMonitor(addr string, tag string) error {

	req, err := c.newRequest(http.MethodPost, "/api/address", &struct {
		Addr string `json:"addr"`
		Tag  string `json:"tag"`
	}{
		addr,
		tag,
	})
	if err != nil {
		return err
	}
	c.auth(req)

	err = c.send(req, nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *APIClient) AddrMonDelete(addr string) error {

	req, err := c.newRequest(http.MethodPost, "/api/delete_address", &struct {
		Addr string `json:"addr"`
	}{
		addr,
	})
	if err != nil {
		return err
	}
	c.auth(req)

	err = c.send(req, nil)
	if err != nil {
		return err
	}

	return nil
}
