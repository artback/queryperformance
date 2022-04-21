all : test  vet

test:
	go test ./...

test-integration:
	go test ./... -tags=integration -shuffle=on

vet :
	go vet ./...

