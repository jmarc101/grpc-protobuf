
# Change this to your own Go module
GO_MODULE := github.com/jmarc101/grpc-protobuf


.PHONY: tidy
tidy:
	go mod tidy

.PHONY: clean
clean:
	rm -fR ./protogen


.PHONY: protoc
# (un)comment or add folder that contains proto as required
protoc:
	protoc \
    --go_opt=module=${GO_MODULE} \
    --go_out=. \
	--go-grpc_out=. \
	--go-grpc_opt=module=${GO_MODULE} \
    ./proto/greeting/*.proto \

.PHONY: build
build: clean protoc tidy


.PHONY: run
run:
	go run main.go


.PHONY: execute
execute: build run


.PHONY: protoc-validate
protoc-validate:
	protoc --validate_out="lang=go:./generated" --go_opt=module=${GO_MODULE} --go_out=. ./proto/car/*.proto
