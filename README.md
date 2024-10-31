## Setup
1. Install [golang](https://go.dev/)
2. Install [Buf](https://buf.build/docs/installation)
3. Install the generator cli 
4. Generate code from proto `buf generate`

### Installing generator cli
```bash
go install -mod=mod \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### Installing generator cli locally and generate
```bash
GOBIN=$PWD/.go/bin go install -mod=mod \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
PATH=$GOBIN:$PATH buf generate
```
