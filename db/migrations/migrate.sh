#!/bin/bash

set -e
until pg_isready -h localhost -p 5432; do
    echo "Waiting for PostgreSQL to start..."
    sleep 2
done

run_migration() {
    local file=$1
    echo "Running $file..."
    PGPASSWORD=$POSTGRES_PASSWORD psql -h localhost -U $POSTGRES_USER -d $POSTGRES_DB -f $file
    echo "$file executed."
}

files=$(ls /migrations/*.sql | sort -t'-' -k1,1n -k2,2n -k3,3n)

# Przejdź przez posortowane pliki SQL i wykonaj je
for file in $files; do
    run_migration $file
done