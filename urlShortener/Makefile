postgresinit:
	docker run -d --name my-postgres -p 5438:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=pg_url postgres

postgres:
	docker exec -it my-postgres psql -U postgres -d pg_url

migrateup:
	 migrate -path ./migrations -database "postgres://postgres:postgres@localhost:5438/pg_url?sslmode=disable" up


.PHONY: postgresinit migrateup postgres