app:
	go run cmd/api/main.go

test:
	docker-compose up db -d
	go test github.com/grootkng/clean-arch-golang/tests