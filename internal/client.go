package internal

import (
	"github.com/valyala/fasthttp"
)

type RemoteClient interface {
	Get(path string) ([]byte, error)
}

type httpClient struct {
	client *fasthttp.Client
}

func (h httpClient) Get(path string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.Header.SetMethod("GET")
	req.Header.SetContentType("application/json")
	req.SetRequestURI(path)

	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)

	if err := h.client.Do(req, res); err != nil {
		return nil, err
	}
	return res.Body(), nil
}

func newClient() RemoteClient {
	hClient := &fasthttp.Client{}
	return &httpClient{client: hClient}
}
