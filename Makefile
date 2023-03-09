proto: 
	protoc pkg/**/pb/*.proto --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative 

postgres:
	docker run --name grpc-user -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine
	
server:
	go run cmd/main.go