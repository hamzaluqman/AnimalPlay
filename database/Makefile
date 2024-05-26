include ../env/mysql.env

db_up:
	migrate -path=./migrations -database "mysql://${DB_USER}:${MYSQL_ROOT_PASSWORD}@(localhost:3306)/${MYSQL_DATABASE}" -verbose up

db_down:
	migrate -path=./migrations -database "mysql://${DB_USER}:${MYSQL_ROOT_PASSWORD}@(localhost:3306)/${MYSQL_DATABASE}" -verbose down

test_db_up:
	migrate -path=./migrations/test_db -database "mysql://${DB_USER}:${MYSQL_ROOT_PASSWORD}@(localhost:3306)/${MYSQL_TEST_DATABASE}" -verbose up

test_db_down:
	migrate -path=./migrations/test_db -database "mysql://${DB_USER}:${MYSQL_ROOT_PASSWORD}@(localhost:3306)/${MYSQL_TEST_DATABASE}" -verbose down

.PHONY: db_up db_down test_db_up test_db_down



