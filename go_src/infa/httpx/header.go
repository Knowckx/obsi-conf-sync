package httpx

import (
	nhttp "net/http"
	"sort"
	"strings"

	fhttp "github.com/bogdanfinn/fhttp"
)

var defaultHeads = map[string]string{
	"Accept":          "*/*",
	"Accept-Encoding": "gzip, deflate",
	"Accept-Language": "zh-CN,zh;q=0.9,en;q=0.8",
	"Connection":      "keep-alive",
	"User-Agent":      chromeUserAgent(defaultChromeVersion),
}

// chromeUserAgent 根据 Chrome 主版本生成 Windows Chrome User-Agent。
func chromeUserAgent(version string) string {
	ver := strings.TrimSpace(version)
	if !strings.Contains(ver, ".") {
		ver += ".0.0.0"
	}
	return "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/" + ver + " Safari/537.36"
}

// 伪装正常浏览器的head顺序
var browserHeaderOrder = []string{
	"Connection",
	"User-Agent",
	"Accept",
	"Accept-Encoding",
	"Accept-Language",
	"Content-Type",
	"Cookie",
}

// applyHeaders 把业务侧 header 覆盖到 tls-client 请求头。
func applyHeaders(req *fhttp.Request, headers nhttp.Header) {
	for key, value := range defaultHeads {
		req.Header.Set(key, value)
	}
	for key, values := range headers {
		req.Header[nhttp.CanonicalHeaderKey(key)] = append([]string(nil), values...)
	}
	// 调整顺序
	req.Header[fhttp.HeaderOrderKey] = buildHeaderOrderFromHeader(req.Header)
}

// buildHeaderOrderFromHeader 生成稳定的 header 顺序：常见浏览器 header 优先，其余按字典序追加。
func buildHeaderOrderFromHeader(headers fhttp.Header) []string {
	if len(headers) == 0 {
		return nil
	}

	order := make([]string, 0, len(headers))
	seen := make(map[string]struct{}, len(headers))

	for _, key := range browserHeaderOrder {
		canonicalKey := nhttp.CanonicalHeaderKey(key)
		if _, ok := headers[canonicalKey]; !ok {
			continue
		}

		order = append(order, canonicalKey)
		seen[canonicalKey] = struct{}{}
	}

	extras := make([]string, 0, len(headers))
	for key := range headers {
		if _, ok := seen[key]; ok {
			continue
		}
		extras = append(extras, key)
	}

	sort.Strings(extras)
	return append(order, extras...)
}
