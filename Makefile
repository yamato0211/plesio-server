
.PHONY: lint lint-fix run logs

lint:
	golangci-lint run ./...

lint-fix:
	golangci-lint run ./...  --fix

run:
	docker compose up --build -d

logs:
	docker compose logs -f
