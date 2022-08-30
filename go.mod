module github.com/grpc-demo

go 1.18

require (
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	go.etcd.io/etcd/api/v3 v3.5.4
	go.etcd.io/etcd/client/v3 v3.5.4
	golang.org/x/net v0.0.0-20220822230855-b0a4917ee28c
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.4 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/sys v0.0.0-20220818161305-2296e01440c6 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc // indirect
)

//replace github.com/coreos/bbolt v1.3.4 => go.etcd.io/bbolt v1.3.4
//replace go.etcd.io/bbolt v1.3.4 => github.com/coreos/bbolt v1.3.4
//replace google.golang.org/grpc => google.golang.org/grpc v1.40.0
