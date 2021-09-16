.PHONY: migration-create migrate-up migrate-down migrate-force prepare init

PWD = $(shell pwd)
PORT = 5432

# Default number of migrations to execute up or down
N = 1
migration-create:
	@echo "---Creating migration files---"
	migrate create -ext sql -dir $(PWD)/migrations -seq -digits 5 $(NAME)

migrate-up:
	migrate -source file://$(PWD)/migrations -database postgres://postgres:123456@localhost:$(PORT)/users-api?sslmode=disable up $(N)

migrate-down:
	migrate -source file://$(PWD)/migrations -database postgres://postgres:123456@localhost:$(PORT)/users-api?sslmode=disable down $(N)

migrate-force:
	migrate -source file://$(PWD)/migrations -database postgres://postgres:123456@localhost:$(PORT)/users-api?sslmode=disable force $(VERSION)

prepare:
	cp .env.example .env && \
	go mod download && go mod verify && \
	docker-compose up -d postgres && \
	$(MAKE) migrate-up N= && \
	docker-compose down

init:
	docker-compose up
