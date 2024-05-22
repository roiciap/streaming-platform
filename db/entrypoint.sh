#!/bin/bash
set -e

# Uruchom domyślny entrypoint PostgreSQL
docker-entrypoint.sh postgres &

# Poczekaj na uruchomienie PostgreSQL
until pg_isready -h localhost -p 5432; do
    echo "Waiting for PostgreSQL to start..."
    sleep 2
done

# Uruchom migracje
/bin/bash /migrations/migrate.sh

# sprawdzić czy to potrzebne
wait $(pgrep -u postgres postgres)