create-migrate:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrate-up:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/cassette?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/cassette?sslmode=disable" -verbose down
up:
	docker compose up

fmt:
	gofmt -s -l -w .

gen:
	./scripts/generate.sh
	
gensqlc-win:
	docker run --rm -v "%cd%:/src" -w /src sqlc/sqlc generate