start:
	bash ./scripts/start.sh

test:
	go test -timeout 1000s -cover -a -v ./...
