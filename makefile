run:
	go run .cmd/
proto:
	protoc internal/auth-svc/pb/*.proto --go_out=. --go-grpc_out=.