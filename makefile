proto:
		protoc pkg/user/pb/*.proto --go-grpc_out=. --go_out=. 
		protoc pkg/service2/pb/*.proto --go-grpc_out=. --go_out=. 
server:
		go run cmd/main.go	

swag: 
	swag init -g pkg/user/routes.go -o ./cmd/api/docs