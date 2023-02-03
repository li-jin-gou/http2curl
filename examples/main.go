package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/valyala/fasthttp"

	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/li-jin-gou/http2curl"
)

func main() {
	HertzDemo()
	FastHttpDemo()
	NetHttpDemo()
}

func HertzDemo() {
	// hertz
	req := protocol.NewRequest(consts.MethodGet, "https://example.com/index", nil)
	req.URI().QueryArgs().Add("a", "1")
	req.URI().QueryArgs().Add("b", "2")
	req.Header.Set("a", "2")
	c, _ := http2curl.GetCurlCommandHertz(req)
	fmt.Println(c)
	// Output: curl -k -X 'GET' -H 'A: 2' -H 'Host: example.com' 'https://example.com/index?a=1&b=2' --compressed
}

func FastHttpDemo() {
	// fasthttp
	var req fasthttp.Request
	req.SetRequestURI("https://example.com/index")
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetBody([]byte(`{"a":"b"}`))
	req.Header.Set("Content-Type", "application/json")

	c, _ := http2curl.GetCurlCommandFastHttp(&req)
	fmt.Println(c)
	// Output: curl -k -X 'POST' -d '{"a":"b"}' -H 'Content-Type: application/json' 'https://example.com/index' --compressed
}

func NetHttpDemo() {
	req, _ := http.NewRequest(http.MethodPost, "https://example.com/index", bytes.NewBufferString(`{"a":"b"}`))
	req.Header.Set("Content-Type", "application/json")
	c, _ := http2curl.GetCurlCommand(req)
	fmt.Println(c)
	// Output: curl -k -X 'POST' -d '{"a":"b"}' -H 'Content-Type: application/json' 'https://example.com/index' --compressed
}
