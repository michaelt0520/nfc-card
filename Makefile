start:
	bash ./scripts/start.sh port=$(port) env=$(env) host=$(host)

test:
	go test -timeout 1000s -cover -a -v ./...
