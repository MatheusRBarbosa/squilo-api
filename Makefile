deps-up:
	docker compose -p gofin up -d

deps-down:
	docker compose down

install:
	go mod vendor && go mod tidy