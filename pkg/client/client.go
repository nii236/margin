package client

import (
	"net/http"
)

// ListURL is the endpoint for List
const ListURL = "/api/positions/list"

// OpenURL is the endpoint for Open
const OpenURL = "/api/positions/open"

// CloseURL is the endpoint for Close
const CloseURL = "/api/positions/close"

// Client is an instance of the server client
type Client struct {
	*http.Client
	Host    string
	Port    string
	Version string
	Scheme  string
}

// New returns a new client
func New(
	Scheme string,
	Host string,
	Port string,
	Version string,

) *Client {
	c := &http.Client{}
	return &Client{
		Client:  c,
		Scheme:  Scheme,
		Host:    Host,
		Port:    Port,
		Version: Version,
	}
}
