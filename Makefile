.PHONY: migration-create migrate-up migrate-down migrate-force prepare create-docs init

PWD = $(shell pwd)
PORT = 5432

# Default number of migrations to execute up or down
N = 1
migration-create:
	@echo "---Creating migration files---"
	migrate create -ext sql -dir $(PWD)/migrations -seq -digits 5 $(NAME);

migrate-up:
	migrate -source file://$(PWD)/migrations -database postgres://postgres:123456@localhost:$(PORT)/users-api?sslmode=disable up $(N);

migrate-down:
	migrate -source file://$(PWD)/migrations -database postgres://postgres:123456@localhost:$(PORT)/users-api?sslmode=disable down $(N);

migrate-force:
	migrate -source file://$(PWD)/migrations -database postgres://postgres:123456@localhost:$(PORT)/users-api?sslmode=disable force $(VERSION);

prepare:
	go mod download && \
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
	docker-compose up -d postgres && \
	./docker/entrypoint.sh 127.0.0.1:5432 && \
	$(MAKE) migrate-up N= && \
	sudo chown -R $(shell echo ${USER}) ./.dbdata && \
	docker-compose down

create-docs:
	swag init main.go;

init:
	docker-compose up
