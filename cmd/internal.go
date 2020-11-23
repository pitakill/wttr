package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

// getWeather queries the URL
func getWeather(c *client) error {
	url := fmt.Sprintf("%s/%s", c.options.URL, c.options.City)

	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	request.URL.RawQuery = getURLQueryEncoded(request.URL, c.options)

	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("%d: %s returns status %s", response.StatusCode, c.options.URL, response.Status)
	}

	return format(response.Body)
}

func getURLQueryEncoded(url *url.URL, o *Options) string {
	q := url.Query()

	if o.Zero {
		q.Add("0", "")
	}
	if o.One {
		q.Add("1", "")
	}
	if o.Two {
		q.Add("2", "")
	}
	if o.A {
		q.Add("A", "")
	}
	if o.F {
		q.Add("F", "")
	}
	if o.N {
		q.Add("n", "")
	}
	if o.Q {
		q.Add("q", "")
	}
	if o.QQ {
		q.Add("Q", "")
	}
	if o.T {
		q.Add("T", "")
	}

	return q.Encode()
}

// format outputs to the terminal
func format(d io.ReadCloser) error {
	defer d.Close()

	body, err := ioutil.ReadAll(d)
	if err != nil {
		return err
	}

	fmt.Print(string(body))

	return nil
}
