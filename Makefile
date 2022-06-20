.PHONY: run
run:
	go run cmd/main.go

.PHONY: docker-up
docker-up:
	docker-compose up --build --remove-orphans

.PHONY: proto-gen
proto-gen:
	protoc --proto_path=proto proto/*.proto --go_out=proto/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=proto/

.PHONY: proto-clean
proto-clean:
	rm proto/*.go

.PHONY: fill-db
fill-db:
	docker exec -i mysql mysql -uroot -p"mysql" -D book < internal/repository/dump.sql

.PHONY: test
test:
	go test -v ./...

.DEFAULT_GOAL = run

