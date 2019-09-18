debug: build-cp
	./kci-server

build-cli:
	cd cmd/cli && go build -o ../../kci

build-cp:
	cd cmd/controlplane && go build -o ../../kci-server

generate-interface:
	protoc -I interface/ interface/main.proto --go_out=plugins=grpc:interface