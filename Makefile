.PHONY: proto
proto:
	protoc -I. \
        -IC:/Users/dimek/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-IC:/Users/dimek/go/src/github.com/grpc-ecosystem/grpc-gateway \
        --openapiv2_out=disable_default_errors=true,allow_merge=true:. --go_out=:. --micro_out=components="micro|http|grpc|gorilla":. proto/*.proto