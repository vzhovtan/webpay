package webpay

import (
	"encoding/json"
	"io"
	"net/http"
)

type CheckoutResult struct {
	Status  int
	Balance int
	Err     string
}

type Cart struct {
	PaymentApiURL string
}

func (c *Cart) Checkout(id string) (*CheckoutResult, error) {
	url := c.PaymentApiURL + "?id=" + id
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &CheckoutResult{}

	err = json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
