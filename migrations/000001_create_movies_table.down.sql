-- migrate create -seq -ext=.sql -dir=./migrations create_movies_table
-- migrate -path="./migrations" -database "postgres://greenlight:0001@127.0.0.1:5432/greenlight?sslmode=disable" up

DROP TABLE IF EXISTS movies;
