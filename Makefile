run:
	go run ./cmd/app
build:
	go build -o markbot ./cmd/app
dev:
	docker compose -f compose.dev.yaml up -d --build
dev.down:
	docker compose -f compose.dev.yaml down
prod:
	docker compose -f compose.prod.yaml up -d --build
prod.down:
	docker compose -f compose.prod.yaml down
