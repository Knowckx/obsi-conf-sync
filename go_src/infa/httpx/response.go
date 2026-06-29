package httpx

import (
	"bytes"
	"fmt"
	"go_trade/src/infa"
	"io"
	"log"
	"mime"
	"time"

	fhttp "github.com/bogdanfinn/fhttp"
	"github.com/cockroachdb/errors"
	"golang.org/x/text/encoding/htmlindex"
	"golang.org/x/text/transform"
)

// Response 只保留常被用到的 HTTP 响应字段
type Response struct {
	StatusCode int
	Body       []byte  // 原始字节 为图片、压缩、protobuf、文件下载预留
	Text       string   // 按 Content-Type charset 解码后的文本
	CostTime   time.Duration
}

func (t Response) String() string {
	out := fmt.Sprintf("http response: status_code=%d cost_time=%.3fs %s",
		t.StatusCode, t.CostTime.Seconds(), infa.PeekStr(t.Text))
	return out
}

// readResponse 读取底层响应体，并转换成上层使用的最小响应结构。
func (t *Client) readResponse(resp *fhttp.Response, cost time.Duration) (*Response, error) {
	if resp == nil {
		return nil, errors.New("nil response")
	}
	defer resp.Body.Close()
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		log.Printf("⚠️ http status %d", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}
	text, err := decodeResponseText(bodyBytes, resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	out := &Response{
		StatusCode: resp.StatusCode,
		Body:       bodyBytes,
		Text:       text,
		CostTime:   cost,
	}
	return out, nil
}

// 按response返回的编码解析数据
func decodeResponseText(body []byte, contentType string) (string, error) {
	if contentType == "" {
		return string(body), nil
	}
	_, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return "", errors.Wrap(err, "parse content type")
	}

	charsetName := params["charset"]
	if charsetName == "" {
		return string(body), nil
	}
	// log.Printf("http response 使用编码 %s", charsetName)
	enc, err := htmlindex.Get(charsetName)
	if err != nil {
		return "", errors.Wrapf(err, "get charset decoder: %s", charsetName)
	}
	textBytes, err := io.ReadAll(transform.NewReader(bytes.NewReader(body), enc.NewDecoder()))
	if err != nil {
		return "", errors.Wrapf(err, "decode response text: %s", charsetName)
	}
	return string(textBytes), nil
}
