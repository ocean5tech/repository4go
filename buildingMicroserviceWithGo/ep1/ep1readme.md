go mod init github.com/ocean5tech/repository4go/buildingMicroserviceWithGo/ep1

用来测试get
curl -v localhost:9090

用来测试post
curl -d 'wy'  localhost:9090  
给出详细信息 -v
curl -v -d 'wy'  localhost:9090  

https://go.dev/src/net/http/server.go?s=61509%3A61556#L2378


https://pkg.go.dev/net/http#HandleFunc
https://pkg.go.dev/net/http#ServeMux