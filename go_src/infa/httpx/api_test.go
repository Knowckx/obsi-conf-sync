package httpx

import (
	"go_trade/src/infa/ops"
	"go_trade/src/inner/log2/mylog"
	"net/url"
	"testing"
)

func Test_Get(t *testing.T) {
	cli, err := NewClient()
	ops.MustNoErrInTest(t, err)

	rawUrl := "https://postman-echo.com/get"
	params := url.Values{
		"wd": {"test"},
	}
	resp, err := cli.Get(rawUrl, params, nil)
	ops.MustNoErrInTest(t, err)
	mylog.Printf("%+v", resp)

	cookResp, err := cli.GetCookies(rawUrl)
	ops.MustNoErrInTest(t, err)
	mylog.Printf("resp Cookies: %+v", cookResp)
}

func Test_PostForm(t *testing.T) {
	cli, err := NewClient()
	ops.MustNoErrInTest(t, err)
	
	rawUrl := "https://postman-echo.com/post"
	params := url.Values{
		"wd": []string{"test"},
	}
	form := url.Values{
		"foo1": {"bar1"},
		"foo2": {"bar2"},
	}
	resp, err := cli.PostForm(rawUrl, params, form, nil)
	ops.MustNoErrInTest(t, err)
	mylog.Printf("%+v", resp)

	cookResp, err := cli.GetCookies(rawUrl)
	ops.MustNoErrInTest(t, err)
	mylog.Printf("resp Cookies: %+v", cookResp)
}

