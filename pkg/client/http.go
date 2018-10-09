package client

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/nii236/margin/pkg/positions"
)

func postHelper(u *url.URL, target interface{}) error {
	resp, err := http.Post(u.String(), "application/json", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("non 200 response: " + strconv.Itoa(resp.StatusCode))
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

// List returns a list of open positions
func (c *Client) List() (*positions.ListResponse, error) {
	u, err := url.Parse(fmt.Sprintf("%s://%s:%s/%s%s", c.Scheme, c.Host, c.Port, c.Version, ListURL))
	if err != nil {
		return nil, err
	}
	result := &positions.ListResponse{}
	err = postHelper(u, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Open creates a new position
func (c *Client) Open(position positions.Type, pair positions.Pair, leverage float64, stack int) (*positions.OpenResponse, error) {
	u, err := url.Parse(fmt.Sprintf("%s://%s:%s/%s", c.Scheme, c.Host, c.Port, OpenURL))
	if err != nil {
		return nil, err
	}
	result := &positions.OpenResponse{}
	err = postHelper(u, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Close creates a new position
func (c *Client) Close(position positions.Type, pair positions.Pair, leverage float64, stack int) (*positions.CloseResponse, error) {
	u, err := url.Parse(fmt.Sprintf("%s://%s:%s/%s", c.Scheme, c.Host, c.Port, OpenURL))
	if err != nil {
		return nil, err
	}
	result := &positions.CloseResponse{}
	err = postHelper(u, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
