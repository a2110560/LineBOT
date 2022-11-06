.PHONY: init build dev
dev:
	go run . server
init:
	docker run -d -p 27017:27017 --name interview-mongo \
		-e MONGO_INITDB_ROOT_USERNAME=root \
		-e MONGO_INITDB_ROOT_PASSWORD=root \
		mongo:4.4
build:
	go build .
