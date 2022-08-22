
# 安装 grpcurl
# go get github.com/fullstorydev/grpcurl
# go install github.com/fullstorydev/grpcurl/cmd/grpcurl

grpcurl -plaintext localhost:8000 list
grpcurl -plaintext localhost:8000 describe helloworld.Greeter
grpcurl -plaintext -d '{"name":"gofer"}'  localhost:8000 helloworld.Greeter/SayHello


## 拦截器测试
grpcurl -plaintext -d '{"name":"grpc","content":"grpc"}'  localhost:9090 interceptor.Speaker/Speak
# {
#   "message": "grpc : grpc"
# }
# 2022/08/22 16:37:16 Hello
# 2022/08/22 16:37:16 Speak>grpc : grpc
# 2022/08/22 16:37:16 Bye bye
