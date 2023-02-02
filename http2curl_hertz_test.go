package http2curl

import (
	"bytes"
	"testing"

	"github.com/cloudwego/hertz/pkg/common/test/assert"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func TestGetCurlCommandHertz(t *testing.T) {
	t.Parallel()

	t.Run("test1", func(t *testing.T) {
		req := protocol.NewRequest(consts.MethodGet, "https://example.com/index", nil)
		req.URI().QueryArgs().Add("a", "1")
		req.URI().QueryArgs().Add("b", "2")
		req.Header.Set("a", "2")
		c, err := GetCurlCommandHertz(req)
		assert.DeepEqual(t, nil, err)
		assert.DeepEqual(t, "curl -k -X 'GET' -H 'A: 2' -H 'Host: example.com' 'https://example.com/index?a=1&b=2' --compressed", c.String())
	})

	t.Run("test2", func(t *testing.T) {
		req := protocol.NewRequest(consts.MethodPost, "https://example.com/index", bytes.NewBufferString(`{"a":"b"}`))
		req.Header.Set("Content-Type", "application/json")

		command, err := GetCurlCommandHertz(req)
		assert.DeepEqual(t, nil, err)
		assert.DeepEqual(t, "curl -k -X 'POST' -d '{\"a\":\"b\"}' -H 'Content-Length: 9' -H 'Content-Type: application/json' -H 'Host: example.com' 'https://example.com/index' --compressed", command.String())
	})
}
