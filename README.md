# microrouter
httprouter for go micro service


Based on httprouter https://github.com/julienschmidt/httprouter

and go-micro framework https://github.com/micro/go-micro/

Support for nats transport


###Usage

```
//Create Service 
svc := NewService("test.service", "0.1")

//Use nats transport
//svc.UseNats("127.0.0.1:4222")

svc.Get("/hello/:id", func(req *Request, resp *Response, p Params) {
		resp.Code = 200
        msg := string(req.Payload)
		if len(p) > 0 {
			resp.Payload = []byte(msg + " " +  p[0].Value)
		}
	})

//Add init function
svc.AddInit(func(){
    //Do something
})

//Run server
svc.Run()

//Set client
client := NewClient("test.service")
//client := NewNatsClient("test.service", "127.0.0.1:4222")
resp, err := client.Get("/hello/zean", &Request{Payload:[]byte("Hello"),Headers:map[string]string{"key":"value"}})
fmt.Println(resp.Code)
fmt.Println(resp.Headers)
fmt.Println(resp.Payload)
fmt.Println(resp.Error)

```
