createdb:
	mysql -e 'CREATE DATABASE book_collection;' -u root
dropdb:
	mysql -e 'DROP DATABASE IF EXISTS book_collection' -u root
migrateup:
	migrate -database "mysql://root:@tcp(localhost:3306)/book_collection" -path db/migration -verbose up
migratedown:
	migrate -database "mysql://root:@tcp(localhost:3306)/book_collection" -path db/migration -verbose down

.PHONY: createdb dropdb migrateup migratedown