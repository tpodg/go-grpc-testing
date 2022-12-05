grpc-generate:
	protoc --proto_path=../proto --go_out=. --go-grpc_out=. ../proto/*.proto

docker-build:
	docker build -t grpc-client .

docker-build-arm:
	docker build -t grpc-client --build-arg TARGETARCH=arm64 .