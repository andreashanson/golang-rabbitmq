docker_build:
	docker build -t golang-rabbit:latest .

docker_run:
	docker container run -it --env-file .env golang-rabbit:latest

test:
	go test ./...

coverage:
	go test ./...