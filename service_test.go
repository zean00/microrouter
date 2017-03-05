package common

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {
	svc := NewService("test", "0.1")
	srv := &BaseService{service: svc}
	svc.Get("/hello", func(req *Request, resp *Response, p Params) {
		resp.Code = 200
		resp.Payload = []byte("OK")
	})

	rsp := &Response{}
	srv.Get(context.TODO(), &Request{Path: "/hello"}, rsp)
	assert.Equal(t, int32(200), rsp.Code)
	assert.Equal(t, []byte("OK"), rsp.Payload)
}

func TestPath(t *testing.T) {
	svc := NewService("test", "0.1")
	srv := &BaseService{service: svc}
	rsp := &Response{}
	srv.Get(context.TODO(), &Request{Path: "/hello"}, rsp)
	assert.Equal(t, int32(404), rsp.Code)
}

func TestMissingPath(t *testing.T) {
	svc := NewService("test", "0.1")
	svc.Get("/hello", func(req *Request, resp *Response, p Params) {
		resp.Code = 200
		resp.Payload = []byte("OK")
	})
	srv := &BaseService{service: svc}
	rsp := &Response{}
	srv.Get(context.TODO(), &Request{Path: "/test"}, rsp)
	assert.Equal(t, int32(404), rsp.Code)
}

func TestParams(t *testing.T) {
	svc := NewService("test", "0.1")
	svc.Get("/hello/:id", func(req *Request, resp *Response, p Params) {
		resp.Code = 200
		if len(p) > 0 {
			resp.Payload = []byte(p[0].Key + ":" + p[0].Value)
		}
	})
	srv := &BaseService{service: svc}
	rsp := &Response{}
	srv.Get(context.TODO(), &Request{Path: "/hello/zean"}, rsp)
	assert.Equal(t, int32(200), rsp.Code)
	assert.Equal(t, []byte("id:zean"), rsp.Payload)
}
