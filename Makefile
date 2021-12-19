docker_build:
	docker build -t golang-rabbit:latest .

test:
	go test ./...

coverage:
	go test ./...