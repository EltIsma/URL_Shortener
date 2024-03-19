1- cd urlShortener
2.1 - Запуск(если используем локальное хранилище) - go run cmd/app/main.go
2.2 - Запуск(если используем postgres) - go run cmd/app/main.go -d  (До запуска нужно поднять бд( docker run -d --name my-postgres -p 5438:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=pg_url postgres))
