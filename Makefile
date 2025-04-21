deps:
	@echo "Installing dependencies..."
	go mod tidy
	go mod download
.PHONY: deps

DSN="postgres://pcappuser:PagasaWarrior23@localhost:5432/pagasa-centre-db?sslmode=disable"

migrate-up:
	@goose -dir ./migrations postgres ${DSN} up
.PHONY: migrate-up

migrate-down:
	@goose -dir ./migrations postgres ${DSN} down
.PHONY: migrate-down

migrate-create:
	@cd ./migrations && goose create add_ministryid_column_to_approvals sql
.PHONY: migrate-create


mock:
	go generate ./...
.PHONY: mock

entity:
	@sqlboiler psql -c ./sqlboiler.toml
.PHONY: entity

down:
	docker-compose down -v
.PHONY: down

start:
	docker-compose up --build
.PHONY: start

lint:
	go mod tidy
	go vet ./...
	gci write -s standard -s default -s "prefix(github.com/Pagasa-Centre/Pagasa-Centre-Mobile-App-API)" .
	gofumpt -l -w .
	wsl -fix ./... 2> /dev/null || true
	golangci-lint run $(p)
	go fmt ./...
.PHONY: lint
