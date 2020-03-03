package blockonomics

import (
	"net/http"
	"net/url"
	"strings"
)

func (c *APIClient) Price(currency string) (float64, error) {

	req, err := c.newRequest(http.MethodGet, "/api/price", nil)
	if err != nil {
		return 0, err
	}

	params := url.Values{}
	params.Add("currency", strings.ToUpper(currency))
	req.URL.RawQuery = params.Encode()

	var data struct {
		Price float64 `json:"price"`
	}
	err = c.send(req, &data)
	if err != nil {
		return 0, err
	}

	return data.Price, nil
}
