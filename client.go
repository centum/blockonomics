package blockonomics

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
)

const APIBase = "https://www.blockonomics.co"

type APIClient struct {
	APIBase string
	Client  *http.Client
	token   string
	timeout time.Duration
}

func NewClient(token string, opts ...Option) *APIClient {
	c := &APIClient{
		APIBase: APIBase,
		token:   token,
	}
	for _, o := range opts {
		o(c)
	}
	c.Client = &http.Client{
		Timeout: c.timeout,
	}
	return c
}

type Option func(s *APIClient)

func WithTimeout(timeout time.Duration) Option {
	return func(s *APIClient) {
		s.timeout = timeout
	}
}

func (c *APIClient) newRequest(method, urlEndpoint string, payload interface{}) (*http.Request, error) {
	u, err := url.Parse(c.APIBase)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, urlEndpoint)

	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}
		buf = bytes.NewBuffer(b)
	}
	return http.NewRequest(method, u.String(), buf)
}

func (c *APIClient) auth(req *http.Request) {
	req.Header.Set("Authorization", "Bearer "+c.token)
}

type Error struct {
	err error
}

func (e *Error) Error() string {
	return "APIClient: " + e.err.Error()
}

func (e *Error) Unwrap() error {
	return e.err
}

var (
	ErrUnauthorised = errors.New("unauthorised")
	ErrBadRequest   = errors.New("bad request")
	ErrServer       = errors.New("server error")
	ErrInternal     = errors.New("internal error")
)

func (c *APIClient) send(req *http.Request, v interface{}) error {
	var (
		err  error
		resp *http.Response
	)

	// Set default headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en_US")

	// Default values for headers
	if req.Header.Get("Content-type") == "" {
		req.Header.Set("Content-type", "application/json")
	}

	resp, err = c.Client.Do(req)

	if err != nil {
		return err
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	switch resp.StatusCode {
	case http.StatusOK:
		if v == nil {
			return nil
		}
		if w, ok := v.(io.Writer); ok {
			_, err = io.Copy(w, resp.Body)
			return err
		}
		return json.NewDecoder(resp.Body).Decode(v)

	case http.StatusUnauthorized:
		return &Error{
			err: ErrUnauthorised,
		}
	case http.StatusBadRequest:
		return &Error{
			err: ErrBadRequest,
		}
	default:
		return &Error{
			err: ErrServer,
		}
	}

}
