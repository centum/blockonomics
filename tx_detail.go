package blockonomics

import (
	"net/http"
	"net/url"
)

type AddressValue struct {
	Addr  string `json:"address"`
	Value int64  `json:"value"`
}

type TxDetail struct {
	VIN    []AddressValue `json:"vin"`
	VOUT   []AddressValue `json:"vout"`
	Status string         `json:"status"`
	Fee    int64          `json:"fee"`
	Time   int64          `json:"time"`
	Size   int64          `json:"size"`
}

func (c *APIClient) TxDetail(txid string) (data *TxDetail, err error) {

	req, err := c.newRequest(http.MethodGet, "/api/tx_detail", nil)
	if err != nil {
		return nil, err
	}

	params := url.Values{}
	params.Add("txid", txid)
	req.URL.RawQuery = params.Encode()

	err = c.send(req, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
