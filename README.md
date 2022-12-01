
migrate -path ./storage/migrations -database 'postgres://admin:qwerty123@localhost:5432/article_db?sslmode=disable' up

migrate -path ./storage/migrations -database 'postgres://admin:qwerty123@localhost:5432/article_db?sslmode=disable' down