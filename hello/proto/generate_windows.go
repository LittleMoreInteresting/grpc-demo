//go:build windows

// 生成 proto
//go:generate protoc -I=. -I=../../third_party --go_out=paths=source_relative:./ ./*.proto
//go:generate protoc -I=. -I=../../third_party --go_out=. --go_opt=paths=source_relative ./*.proto
package proto
