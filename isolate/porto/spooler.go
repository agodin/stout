package porto

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
)

var (
	ErrCursorNoResult = errors.New("cursor: no more results")
)

// SpoolerConfig ...
type SpoolerConfig struct {
	Path string `json:"path"`
}

type Cursor interface {
	Name() string
	Next() (io.ReadCloser, error)
}

type cursor struct {
	ctx    context.Context
	name   string
	urls   []string
	client *http.Client
}

func newCursor(ctx context.Context, name string, urls []string, client *http.Client) *cursor {
	if client == nil {
		client = http.DefaultClient
	}
	cur := &cursor{
		ctx:    ctx,
		name:   name,
		urls:   urls,
		client: client,
	}

	return cur
}

func (c *cursor) Name() string {
	return c.name
}

func (c *cursor) Next() (io.ReadCloser, error) {
	if len(c.urls) == 0 {
		return nil, ErrCursorNoResult
	}

	var url string
	url, c.urls = c.urls[0], c.urls[1:]

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := ctxhttp.Do(c.ctx, c.client, req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("cursor: not OK %s", resp.Status)
	}

	return resp.Body, nil
}

// Spooler downloads images from Docker registry/distribution
type Spooler struct {
	config *SpoolerConfig
}

// NewSpooler creates new Spooler
func NewSpooler(ctx context.Context, config *SpoolerConfig) (*Spooler, error) {
	spooler := Spooler{
		config: config,
	}

	return &spooler, nil
}

func (s *Spooler) Cursor(name string) (*Cursor, error) {
	return nil, nil
}
