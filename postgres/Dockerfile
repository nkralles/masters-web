FROM timescale/timescaledb:latest-pg14

LABEL maintainer="Nicholas Kralles <njkral@n-ask.com>"

ENV POSTGRES_INITDB_ARGS -E UTF8
ENV POSTGRES_HOST=0.0.0.0
ENV POSTGRES_PORT=5432
ENV POSTGRES_USER=master
ENV POSTGRES_PASSWORD=nick-masters
ENV POSTGRES_DBNAME=masters
ENV PGDATA=/var/lib/postgresql/data/pgdata

ADD initdb/* /docker-entrypoint-initdb.d/
EXPOSE ${POSTGRES_PORT}
