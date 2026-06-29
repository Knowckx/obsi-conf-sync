// package httpx 提供反爬虫的 HTTP 请求封装。
package httpx

import (
	"go_trade/src/infa/ops"

	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/bogdanfinn/tls-client/profiles"
	"github.com/cockroachdb/errors"
)

const (
	defaultHTTPTimeoutSec = 3
	defaultChromeVersion  = "133"
)

var (
	defaultClientProfile = profiles.Chrome_133 // 默认浏览器指纹
)

// newClient 创建带浏览器指纹的 tls-client 客户端。
func newClient(timeout int,) (*Client, error) {
	secs := ops.IfElse(timeout <= 0, defaultHTTPTimeoutSec, timeout)

	// tls_client 会自动管理 CookieJar 上下文
	jar := tls_client.NewCookieJar()

	opts := []tls_client.HttpClientOption{
		tls_client.WithTimeoutSeconds(secs),
		tls_client.WithCookieJar(jar),
		tls_client.WithClientProfile(defaultClientProfile),
	}

	cli, err := tls_client.NewHttpClient(
		tls_client.NewNoopLogger(), opts...,
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	out := &Client{}
	out.cli = cli
	return out, nil
}
