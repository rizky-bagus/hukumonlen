test:
	go test -v -race ./...

dep:
	env GO111MODULE=on go mod download

tidy:
	env GO111MODULE=on go mod tidy

vendor:
	env GO111MODULE=on go mod vendor

cover:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out

format:
	scripts/format.sh

check.import:
	scripts/check-import.sh

cleanlintcache:
	golangci-lint cache clean

lint: cleanlintcache
	golangci-lint run ./...

pretty: tidy format lint

.PHONY: migration
migration:
	migrate create -ext sql -dir db/migrations $(name)

.PHONY: migrate
migrate:
	migrate -path db/migrations -database 'mysql://root:Secret1234!@localhost:3306/hukumonlen' -verbose up

.PHONY: rollback
rollback:
	migrate -path db/migrations -database "$(url)" -verbose down 1

.PHONY: rollback-all
rollback-all:
	migrate -path db/migrations -database "$(url)" -verbose down -all

.PHONY: force-migrate
force-migrate:
	migrate -path db/migrations -database "$(url)" -verbose force $(version)

mockgen:
	scripts/generate-mock.sh