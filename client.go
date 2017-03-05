package common

import (
	"context"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-plugins/transport/nats"
)

//ServiceClient client for generic service
type ServiceClient struct {
	client GenericServiceClient
}

//NewClient http client
func NewClient(service string) *ServiceClient {
	return &ServiceClient{client: NewGenericServiceClient(service, client.DefaultClient)}
}

//NewNatsClient new nats client
func NewNatsClient(service, address string) *ServiceClient {
	trans := nats.NewTransport()
	if address == "" {
		address = "127.0.0.1:4222"
	}
	trans.Listen(address)
	cl := client.NewClient(client.Transport(trans))
	return &ServiceClient{client: NewGenericServiceClient(service, cl)}
}

//Get request get
func (c *ServiceClient) Get(path string, req *Request) (*Response, error) {
	req.Path = path
	req.Method = "GET"
	return c.client.Get(context.TODO(), req)
}

//Post request get
func (c *ServiceClient) Post(path string, req *Request) (*Response, error) {
	req.Path = path
	req.Method = "POST"
	return c.client.Post(context.TODO(), req)
}

//Put request get
func (c *ServiceClient) Put(path string, req *Request) (*Response, error) {
	req.Path = path
	req.Method = "PUT"
	return c.client.Put(context.TODO(), req)
}

//Delete request get
func (c *ServiceClient) Delete(path string, req *Request) (*Response, error) {
	req.Path = path
	req.Method = "DELETE"
	return c.client.Delete(context.TODO(), req)
}
