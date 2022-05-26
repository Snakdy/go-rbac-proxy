.PHONY: install
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

.PHONY: generate
generate:
	protoc \
    		-I. \
    		-I${GOPATH}/pkg/mod \
    		--go_out=. --go_opt=paths=source_relative \
            --go-grpc_out=. --go-grpc_opt=paths=source_relative \
            pkg/rbac/authority.proto
	# https://stackoverflow.com/a/37335452
	# removes the 'omitempty' json tag
	ls pkg/rbac/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
