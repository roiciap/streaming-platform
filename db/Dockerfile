FROM postgres:16.3

COPY ./migrations/ migrations

COPY ./entrypoint.sh /docker-entrypoint-initdb.d/entrypoint.sh
RUN chmod +x /docker-entrypoint-initdb.d/entrypoint.sh
