migrate-up:
	@go run ./migrations/cmd/up/main.go
migrate-down:
	@go run ./migrations/cmd/down/main.go

# run main.go
run:
	@go run main.go