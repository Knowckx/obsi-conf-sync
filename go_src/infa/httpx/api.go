package httpx

import (
	nhttp "net/http"
	"net/url"

	fhttp "github.com/bogdanfinn/fhttp"
	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/cockroachdb/errors"
)

type Client struct {
	cli tls_client.HttpClient
}

// NewClient 创建可复用的 tls-client 客户端。
func NewClient() (*Client, error) {
	return newClient(defaultHTTPTimeoutSec)
}

// NewClientTimeout 创建可复用的 tls-client 客户端。
func NewClientTimeout(timeout int) (*Client, error) {
	return newClient(timeout)
}

// GetCookies 返回指定 URL 下的 Cookie。
func (c *Client) GetCookies(rawURL string) ([]*fhttp.Cookie, error) {
	uurl, err := url.Parse(rawURL)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return c.cli.GetCookies(uurl), nil
}

// Get 发起 GET 请求。
func (t *Client) Get(
	rawURL string, params url.Values, headers nhttp.Header, 
) (*Response, error) {
	return t.get(rawURL, params, headers,)
}


// PostForm 发起表单 POST 请求。
func (c *Client) PostForm(
	rawURL string, params url.Values, form url.Values, headers nhttp.Header, 
) (*Response, error) {
	return c.postForm(rawURL, params, form, headers)
}
