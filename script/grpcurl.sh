
# 安装 grpcurl
# go get github.com/fullstorydev/grpcurl
# go install github.com/fullstorydev/grpcurl/cmd/grpcurl

grpcurl -plaintext localhost:8000 list
grpcurl -plaintext localhost:8000 describe helloworld.Greeter
grpcurl -plaintext -d '{"name":"gofer"}'  localhost:8000 helloworld.Greeter/SayHello