package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"
)

//go:generate mockery --all

type (
	Requester interface {
		Post(url string, contentType string, body io.Reader) (*http.Response, error)
		Get(url string) (resp *http.Response, err error)
	}
	Client struct {
		http Requester
	}
	Response struct {
		StatusCode int
		Result     any
	}
)

func NewRequester() Requester {
	return &http.Client{
		Timeout: 1 * time.Second,
	}
}

func NewClient(http Requester) Client {
	return Client{
		http: http,
	}
}

func (c Client) Get(ctx context.Context, url string) (Response, error) {
	duration := time.Now()
	var (
		resp *http.Response
		err  error
	)
	defer func() {
		c.logRequest(ctx, url, duration, err, resp)
		_ = resp.Body.Close()
	}()
	resp, err = c.http.Get(url)
	if err != nil {
		return Response{}, err
	}
	return c.decodeResponse(resp)
}

func (c Client) Post(ctx context.Context, url string, body any) (Response, error) {
	duration := time.Now()
	var (
		resp *http.Response
		err  error
	)
	defer func() {
		c.logRequest(ctx, url, duration, err, resp)
		_ = resp.Body.Close()
	}()
	marshal, err := json.Marshal(body)
	if err != nil {
		return Response{}, err
	}
	resp, err = c.http.Post(url, "application/json", bytes.NewBuffer(marshal))
	if err != nil {
		return Response{}, err
	}
	return c.decodeResponse(resp)
}

func (c Client) decodeResponse(resp *http.Response) (Response, error) {
	result := Response{
		StatusCode: resp.StatusCode,
	}
	if err := json.NewDecoder(resp.Body).Decode(&result.Result); err != nil {
		return Response{}, err
	}
	return result, nil
}

func (c Client) logRequest(ctx context.Context, url string, duration time.Time, err error, resp *http.Response) {
	var errorMsg string
	if err != nil {
		errorMsg = err.Error()
	}
	var status string
	if resp != nil {
		status = resp.Status
	}
	slog.InfoContext(ctx,
		"request info",
		slog.String("url", url),
		slog.Duration("duration", time.Since(duration)),
		slog.String("error_msg", errorMsg),
		slog.String("status", status),
	)
}
