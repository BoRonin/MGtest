.PHONY: up down build

up:
	docker compose up -d

down:
	docker compose down

build:
	docker compose up --build

api:
	docker compose restart api-service