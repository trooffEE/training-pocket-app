#!/usr/bin/env sh

TIMESTAMP=$(date +%Y-%m-%d_%H-%M-%S)
docker exec -it training-db-container sh -c "pg_dump -U ${DB_USER} ${DB_NAME} > ./tmp/dump_${TIMESTAMP}.sql"
docker cp training-db-container:/tmp/dump_${TIMESTAMP}.sql ./tmp/dump_${TIMESTAMP}.sql
docker exec -it training-db-container sh -c "rm ./tmp/dump_${TIMESTAMP}.sql"
