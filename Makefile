
run:
	go run cmd/main.go

swag-gen:
	swag init -g api/api.go -o api/docs

