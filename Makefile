.POHNY: build
build:
	go build -o calendar-backend cmd/backend/main.go 

.POHNY: test
test:
	go test ./...