package http2curl

import (
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/valyala/fasthttp"
)

func TestGetCurlCommandFastHttp(t *testing.T) {
	t.Parallel()

	t.Run("test1", func(t *testing.T) {
		var req fasthttp.Request
		req.SetRequestURI("https://example.com/index")
		req.URI().QueryArgs().Add("a", "1")
		req.URI().QueryArgs().Add("b", "2")
		req.Header.Set("a", "2")
		c, err := GetCurlCommandFastHttp(&req)
		assert.DeepEqual(t, nil, err)
		assert.DeepEqual(t, "curl -k -X 'GET' -H 'A: 2' 'https://example.com/index?a=1&b=2' --compressed", c.String())
	})

	t.Run("test2", func(t *testing.T) {
		var req fasthttp.Request
		req.SetRequestURI("https://example.com/index")
		req.Header.SetMethod(fasthttp.MethodPost)
		req.SetBody([]byte(`{"a":"b"}`))
		req.Header.Set("Content-Type", "application/json")

		command, err := GetCurlCommandFastHttp(&req)
		assert.DeepEqual(t, nil, err)
		assert.DeepEqual(t, "curl -k -X 'POST' -d '{\"a\":\"b\"}' -H 'Content-Type: application/json' 'https://example.com/index' --compressed", command.String())
	})
}
