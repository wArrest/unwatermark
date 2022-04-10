package unwatermark

import (
	"bytes"
	"mvdan.cc/xurls/v2"
	"net/http"
)

func NewReq(url string, body []byte) *http.Request {
	var req *http.Request
	if body == nil {
		req, _ = http.NewRequest("GET", url, nil)
	} else {
		req, _ = http.NewRequest("POST", url, bytes.NewReader(body))
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Add("Connection", "keep-alive")
	return req
}

//提取有用的链接部分
func SimpleTxt(txt string) string {
	//提取url
	xurlsStrict := xurls.Strict()
	output := xurlsStrict.FindAllString(txt, -1)
	if len(output) != 1 {
		return ""
	}
	return output[0]
}
