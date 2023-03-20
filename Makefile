.PHONY: generate_pb

build_proto_builder:
	@docker build -f Dockerfile.protos --tag developer-protos .

build_go: 
	@docker run --rm -v "$(PWD)/developer-protos-src:/developer-protos" developer-protos

generate_pb:
	@protoc -I ./src/DeveloperApi --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import ./src/DeveloperApi/*.proto