package frappe

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Client struct
type Client struct {
	URL    string
	Key    string
	Secret string
}

// GetAll returns all documents
func (c *Client) GetAll(docType string, fields []string, filters []Filter, limit int, start int) ([]Document, error) {
	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}

	u.Path = "/api/resource/" + docType

	q := u.Query()

	if len(fields) > 0 {
		rawFields := "["
		for k, v := range fields {
			if k > 0 {
				rawFields += ","
			}
			rawFields += `"` + v + `"`
		}
		rawFields += "]"
		q.Set("fields", rawFields)
	}

	if len(filters) > 0 {
		rawFilters := "["
		for k, v := range filters {
			if k > 0 {
				rawFilters += ","
			}
			rawFilters += `["` + v.DocType + `","` + v.Field + `","` + v.Operator + `","` + v.Operand + `"]`

		}
		rawFilters += "]"
		q.Set("filters", rawFilters)
	}

	if limit > 0 {
		q.Set("limit_page_length", string(limit))
	}
	if start > 0 {
		q.Set("limit_start", string(start))
	}

	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+c.Key+":"+c.Secret)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	s := &struct {
		Data []Document
	}{}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(s); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("HTTP Status Code: %d (%s)", resp.StatusCode, resp.Status)
	}

	return s.Data, nil
}

// Get returns a document
func (c *Client) Get(docType string, name string, fields []string) (Document, error) {
	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}

	u.Path = "/api/resource/" + docType + "/" + name

	q := u.Query()

	if len(fields) > 0 {
		rawFields := "["
		for k, v := range fields {
			if k > 0 {
				rawFields += ","
			}
			rawFields += `"` + v + `"`
		}
		rawFields += "]"
		q.Set("fields", rawFields)
	}

	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "token "+c.Key+":"+c.Secret)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	respStruct := &struct {
		Data Document `json:"data"`
	}{}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(respStruct); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("HTTP Status Code: %d (%s)", resp.StatusCode, resp.Status)
	}

	return respStruct.Data, nil
}

// Post creates a document
func (c *Client) Post(docType string, doc Document) (Document, error) {
	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}

	u.Path = "/api/resource/" + docType

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(doc); err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+c.Key+":"+c.Secret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	respStruct := &struct {
		Data Document `json:"data"`
	}{}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(respStruct); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("HTTP Status Code: %d (%s)", resp.StatusCode, resp.Status)
	}

	return respStruct.Data, nil
}

// Put updates a document
func (c *Client) Put(docType string, name string, doc Document) (Document, error) {
	u, err := url.Parse(c.URL)
	if err != nil {
		return nil, err
	}

	u.Path = "/api/resource/" + docType + "/" + name

	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(doc); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+c.Key+":"+c.Secret)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	respStruct := &struct {
		Data Document `json:"data"`
	}{}

	if resp.StatusCode == http.StatusOK {
		if err := json.NewDecoder(resp.Body).Decode(respStruct); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("HTTP Status Code: %d (%s)", resp.StatusCode, resp.Status)
	}

	return respStruct.Data, nil
}
