run:
	go run cmd/*.go
proto:
	protoc pkg/**/pb/*.proto --go_out=. --go-grpc_out=.
swag: ## Generate swagger docs
	swag init -g cmd/main.go -o ./api/docs