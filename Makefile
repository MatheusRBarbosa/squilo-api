deps-up:
	docker compose -p squilo up -d

deps-down:
	docker compose down

install:
	go mod vendor && go mod tidy