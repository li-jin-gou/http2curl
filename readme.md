# http2curl

convert Request of [fasthttp](https://github.com/valyala/fasthttp), [hertz](https://github.com/cloudwego/hertz) and net/http to CURL command line and fork from [moul/http2curl](https://github.com/moul/http2curl)


# Install

```shell
go get github.com/li-jin-gou/http2curl
```

## Usage


### FastHttp

```go
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

```

### Hertz

```go
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
```

### net/http

```go
func NetHttpDemo() {
	req, _ := http.NewRequest(http.MethodPost, "https://example.com/index", bytes.NewBufferString(`{"a":"b"}`))
	req.Header.Set("Content-Type", "application/json")
	c, _ := http2curl.GetCurlCommand(req)
	fmt.Println(c)
	// Output: curl -k -X 'POST' -d '{"a":"b"}' -H 'Content-Type: application/json' 'https://example.com/index' --compressed
}
```

### Resty

```go
func RestyDemo() {
	var req *resty.Request
	client := resty.New()
	resp, _ := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody([]byte(`{"a":"b"}`)).
		Post("https://example.com/index")
	req = resp.Request
	c, _ := http2curl.GetCurlCommandResty(req)
	fmt.Println(c)
	// Output: curl -k -X 'POST' -d '{"a":"b"}' -H 'Accept: application/json' -H 'Content-Type: application/json' -H 'User-Agent: go-resty/2.12.0 (https://github.com/go-resty/resty)' 'https://example.com/index' --compressed
}
```