package linker

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"time"

	"github.com/seibert-media/golibs/log"
	"go.uber.org/zap"

	"github.com/sethgrid/pester"

	"github.com/pkg/errors"

	"github.com/canonical-debate-lab/argument-analysis-research/pkg/document"
)

// Rater provides the interface for rating two segments against eachother
type Rater interface {
	Rate(ctx context.Context, a, b *document.Segment) (float32, error)
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
		MaxRetries: 5,

		client: client,
	}
}

// Rate by sending json to URL and retry on timeout
// TODO(kwiesmueller): add two-way rating
func (r *HTTPRater) Rate(ctx context.Context, a, b *document.Segment) (float32, error) {
	type Request struct {
		A string `json:"text1"`
		B string `json:"text2"`
	}

	type Response struct {
		Dist float32 `json:"value"`
	}

	req := Request{a.Text, b.Text}

	data, err := json.Marshal(req)
	if err != nil {
		return 0, errors.Wrap(err, "encoding request")
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

	var response *Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.From(ctx).Error("decoding response", zap.String("url", r.URL), zap.ByteString("body", body), zap.Error(err))
		return 0, errors.Wrap(err, "decoding response")
	}

	return response.Dist, nil
}
