.PHONY: all up migrate

all: up migrate

up:
	docker-compose up -d

migrate:
	docker-compose exec db sh -c 'psql -U company < /db/migrations/company.sql'

