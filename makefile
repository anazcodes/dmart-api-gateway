run:
	go run cmd/*.go
proto:
	protoc pkg/**/pb/*.proto --go_out=. --go-grpc_out=.
