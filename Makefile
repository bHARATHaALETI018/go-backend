postgres:
	docker run --name opconnect-db -p 5432:5432 -e POSTGRES_PASSWORD=password -d  postgres:alpine

stopNdeletePostgresContainer:
	docker stop opconnect-db; docker rm opconnect-db

createDb:
	docker exec -it opconnect-db createdb -U postgres opconnect

useDatabase:
	docker exec -it opconnect-db psql -U postgres opconnect

dropDb: 
	docker exec -it opconnect-db dropdb -U postgres opconnect

migrateUp:
	migrate -path db/migration -database "postgres://postgres:password@localhost:5432/opconnect?sslmode=disable" -verbose up

migrateDown:
	migrate -path db/migration -database "postgres://postgres:password@localhost:5432/opconnect?sslmode=disable" -verbose down

downAll:
	make migrateDown; make dropDb; make stopNdeletePostgresContainer

.PHONY: postgres stopNdeletePostgresContainer createDb useDatabase dropDb migrateUp migrateDown