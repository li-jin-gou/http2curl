package main

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/li-jin-gou/http2curl"
	"github.com/valyala/fasthttp"
)

func main() {
	// hertz
	req := protocol.NewRequest(consts.MethodGet, "https://example.com/index", nil)
	req.URI().QueryArgs().Add("a", "1")
	req.URI().QueryArgs().Add("b", "2")
	req.Header.Set("a", "2")
	c, _ := http2curl.GetCurlCommandHertz(req)
	fmt.Println(c)
	// Output: curl -k -X 'GET' -H 'A: 2' -H 'Host: example.com' 'https://example.com/index?a=1&b=2' --compressed

	// fasthttp
	var req1 fasthttp.Request
	req1.SetRequestURI("https://example.com/index")
	req1.Header.SetMethod(fasthttp.MethodPost)
	req1.SetBody([]byte(`{"a":"b"}`))
	req1.Header.Set("Content-Type", "application/json")

	c, _ = http2curl.GetCurlCommandFastHttp(&req1)
	fmt.Println(c)
	// Output: curl -k -X 'POST' -d '{"a":"b"}' -H 'Content-Type: application/json' 'https://example.com/index' --compressed
}
