package httpx

import (
	"context"
	"io"
	nhttp "net/http"
	"net/url"
	"strings"
	"time"

	fhttp "github.com/bogdanfinn/fhttp"
	"github.com/cockroachdb/errors"
)

// Get 发起 GET 请求。
func (t *Client) get(
	rawURL string, params url.Values, headers nhttp.Header,
) (*Response, error) {

	uURL, err := buildRequestURL(rawURL, params)
	if err != nil {
		return nil, errors.Wrap(err, "build request url")
	}
	rawUrl := uURL.String()

	ctx := context.Background()
	req, err := fhttp.NewRequestWithContext(ctx, nhttp.MethodGet, rawUrl, nil)
	if err != nil {
		return nil, errors.Wrap(err, "build request")
	}
	applyHeaders(req, headers)
	now := time.Now()
	resp, err := t.cli.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	cost := time.Since(now)

	return t.readResponse(resp, cost)
}

// buildRequestURL 在原始 URL 上追加 query 参数。
func buildRequestURL(rawURL string, params url.Values) (*url.URL, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	query := parsed.Query()
	for key, values := range params {
		query.Del(key)
		for _, value := range values {
			query.Add(key, value)
		}
	}
	parsed.RawQuery = query.Encode()
	return parsed, nil
}


// postForm 发起表单 POST 请求。
func (t *Client) postForm(
	rawURL string, params url.Values, form url.Values, headers nhttp.Header, 
) (*Response, error) {
	uURL, err := buildRequestURL(rawURL, params)
	if err != nil {
		return nil, errors.Wrap(err, "build request url")
	}
	rawUrl := uURL.String()

	body := buildFormBody(form)

	ctx := context.Background()
	req, err := fhttp.NewRequestWithContext(ctx, nhttp.MethodPost, rawUrl, body)
	if err != nil {
		return nil, errors.Wrap(err, "build request")
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	applyHeaders(req, headers)

	now := time.Now()
	resp, err := t.cli.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	cost := time.Since(now)
	return t.readResponse(resp, cost)
}

// buildFormBody 根据表单数据构造请求体。
func buildFormBody(form url.Values) io.Reader {
	if len(form) == 0 {
		return nil
	}
	return strings.NewReader(form.Encode())
}