migrate:
	migrate create -ext sql -dir db/migrations -seq $(name)

up:
	docker compose up

fmt:
	gofmt -s -l -w .
