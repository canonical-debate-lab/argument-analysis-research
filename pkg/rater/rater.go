package rater

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/pkg/errors"
	"github.com/seibert-media/golibs/log"
	"github.com/sethgrid/pester"
	"go.uber.org/zap"
)

// Rater provides the interface for rating two segments against eachother
type Rater interface {
	Rate(ctx context.Context, a, b string) (float32, error)
}

// HTTPRater implements Rater and executes the rating through a http request which it retries on timeout
type HTTPRater struct {
	URL        string
	MaxRetries int
	Timeout    time.Duration

	client *pester.Client
}

// NewHTTPRater for URL
func NewHTTPRater(url string) *HTTPRater {
	client := pester.New()
	client.MaxRetries = 5
	client.Backoff = pester.ExponentialBackoff
	client.Timeout = 60 * time.Second

	return &HTTPRater{
		URL:        url,
		MaxRetries: 10,

		client: client,
	}
}

// Rate by sending json to URL and retry on timeout
// TODO(kwiesmueller): add two-way rating
func (r *HTTPRater) Rate(ctx context.Context, a, b string) (float32, error) {
	type Request struct {
		A string `json:"text1"`
		B string `json:"text2"`
	}

	type Response struct {
		Dist float32 `json:"value"`
	}

	req := Request{a, b}

	data, err := json.Marshal(req)
	if err != nil {
		return 0, errors.Wrap(err, "encoding request")
	}

	if len(req.A) == 0 || len(req.B) == 0 {
		return 0, errors.New("invalid request")
	}

	resp, err := r.client.Post(r.URL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.From(ctx).Error("sending request", zap.String("url", r.URL), zap.Error(err))
		return 0, errors.Wrap(err, "sending request")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.From(ctx).Error("reading response", zap.String("url", r.URL), zap.Error(err))
		return 0, errors.Wrap(err, "reading response")
	}

	if resp.StatusCode != 200 {
		log.From(ctx).Error("sending request", zap.Int("status", resp.StatusCode), zap.String("url", r.URL), zap.Error(err))
		return 0, fmt.Errorf("request failed: status %d: %s", resp.StatusCode, body)
	}

	var response *Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.From(ctx).Error("decoding response", zap.String("url", r.URL), zap.ByteString("body", body), zap.Error(err))
		return 0, errors.Wrap(err, "decoding response")
	}

	return response.Dist, nil
}
