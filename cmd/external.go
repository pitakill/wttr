package cmd

import (
	"net/http"
	"time"
)

var (
	defaultCity   = "La+Tortuga+Queretaro"
	defaultUrl    = "https://wttr.in"
	defaultClient = &http.Client{
		Timeout: time.Second * 5,
	}
)

// client object to configure the http client and the options for the request
type client struct {
	httpClient *http.Client
	options    *Options
}

// Options object to specify the URL, City, etc.
type Options struct {
	URL  string
	City string
	Zero bool
	One  bool
	Two  bool
	A    bool
	F    bool
	N    bool
	Q    bool
	QQ   bool
	T    bool
}

// NewClient
func NewClient(o *Options) *client {
	if o == nil {
		o = &Options{}
	}

	if len(o.City) == 0 {
		o.City = defaultCity
	}

	if len(o.URL) == 0 {
		o.URL = defaultUrl
	}

	return &client{
		httpClient: defaultClient,
		options:    o,
	}
}

// GetWeather
func (c *client) GetWeather() error {
	return getWeather(c)
}
