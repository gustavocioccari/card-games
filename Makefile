migrate:
	go run ./cmd/migration/main.go

seed:
	@ go run ./cmd/seed/main.go

test:
	@ go test ./... -v -count=1

deploy: 
	go run ./cmd/api/main.go
