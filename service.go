package common

import (
	"context"
	"errors"
	"log"

	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-micro/transport"
	"github.com/micro/go-plugins/transport/nats"
)

//Service service object
type Service struct {
	Name        string
	Version     string
	ID          string
	nats        bool
	natsAddress string
	seed        []byte
	inits       []InitFunc
	trees       map[string]*node
}

//InitFunc init function
type InitFunc func()

//HandlerFunc request handler function
type HandlerFunc func(*Request, *Response, Params)

//BaseService base exposed servce API
type BaseService struct {
	service *Service
}

// Param is a single URL parameter, consisting of a key and a value.
type Param struct {
	Key   string
	Value string
}

//Params parameter
type Params []Param

//Get get resource
func (s *BaseService) Get(ctx context.Context, req *Request, resp *Response) error {
	return s.serve("GET", ctx, req, resp)
}

//Post post resource
func (s *BaseService) Post(ctx context.Context, req *Request, resp *Response) error {
	return s.serve("POST", ctx, req, resp)
}

//Put put resource
func (s *BaseService) Put(ctx context.Context, req *Request, resp *Response) error {
	return s.serve("PUT", ctx, req, resp)
}

//Delete delete resource
func (s *BaseService) Delete(ctx context.Context, req *Request, resp *Response) error {
	return s.serve("DELETE", ctx, req, resp)
}

//Head check resource
func (s *BaseService) Head(ctx context.Context, req *Request, resp *Response) error {
	return s.serve("HEAD", ctx, req, resp)
}

//Options resource option
func (s *BaseService) Options(ctx context.Context, req *Request, resp *Response) error {
	return s.serve("OPTIONS", ctx, req, resp)
}

func (s *BaseService) serve(method string, ctx context.Context, req *Request, resp *Response) error {
	path := req.Path
	if root := s.service.trees[method]; root != nil {
		if handle, ps, _ := root.getValue(path); handle != nil {
			handle(req, resp, ps)
			return nil
		}
	}
	resp.Code = 404
	resp.Error = "Invalid path"
	return errors.New(resp.Error)
}

//AddInit add init function
func (s *Service) AddInit(fn InitFunc) {
	s.inits = append(s.inits, fn)
}

//Get add get handler
func (s *Service) Get(path string, fn HandlerFunc) {
	s.Handle("GET", path, fn)
}

//Post add post handler
func (s *Service) Post(path string, fn HandlerFunc) {
	s.Handle("POST", path, fn)
}

//Put add put handler
func (s *Service) Put(path string, fn HandlerFunc) {
	s.Handle("PUT", path, fn)
}

//Delete add delete handler
func (s *Service) Delete(path string, fn HandlerFunc) {
	s.Handle("DELETE", path, fn)
}

//Handle add handler
func (s *Service) Handle(method, path string, handle HandlerFunc) {
	if path[0] != '/' {
		panic("path must begin with '/' in path '" + path + "'")
	}

	if s.trees == nil {
		s.trees = make(map[string]*node)
	}

	root := s.trees[method]
	if root == nil {
		root = new(node)
		s.trees[method] = root
	}

	root.addRoute(path, handle)
}

//UseNats enable nats transport
func (s *Service) UseNats(address string) {
	s.nats = true
	s.natsAddress = "127.0.0.1:4222"
	if address != "" {
		s.natsAddress = address
	}

}

//NewService instantiate service
func NewService(name, version string) *Service {
	svc := &Service{
		Name:    name,
		Version: version,
	}
	svc.inits = make([]InitFunc, 0)
	svc.trees = make(map[string]*node)
	return svc
}

//Run run service
func (s *Service) Run() {

	for _, init := range s.inits {
		init()
	}

	trans := transport.DefaultTransport
	if s.nats {
		trans = nats.NewTransport()
		trans.Listen(s.natsAddress)
	}

	cmd.Init()

	// Initialise Server
	server.Init(
		server.Name(s.Name),
		server.Id(s.ID),
		server.Version(s.Version),
		server.Metadata(map[string]string{
			"node": s.ID,
		}),
		server.Transport(trans),
	)

	// Register Handlers
	server.Handle(server.NewHandler(&BaseService{service: s}))

	// Run server
	if err := server.Run(); err != nil {
		log.Panic(err)
	}
}

//GetValue get parameter value
func (p Params) GetValue(key string) string {
	for _, param := range p {
		if param.Key == key {
			return param.Value
		}
	}
	return ""
}

//SetError set error message and code
func (r *Response) SetError(code int32, message interface{}) {
	r.Code = code
	r.Error = ToString(message)
}
