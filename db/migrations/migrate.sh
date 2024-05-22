#!/bin/bash

set -e

run_migration() {
    local file=$1
    echo "Running $file..."
    PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -U $DB_USER -d $DB_NAME -f $file
    echo "$file executed."
}

