include .env

create_network:
	docker network create net_trinity

docker_build:
	docker-compose -f docker-compose.yml up --build -d

migrate_up:
	docker-compose -f docker-compose.yml exec be-trinity migrate -path migrations \
		-database "$(POSTGRES_DB_URL)" -verbose up 1
