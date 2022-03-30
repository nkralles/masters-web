#!/bin/sh

set -e
docker stop masters-db || true
docker rm masters-db || true
docker run --restart=always -d \
  -e POSTGRES_USER=masters \
  -e POSTGRES_DB=masters \
  -e POSTGRES_PASSWORD=nick-masters \
  -p 55432:5432 \
  --name masters-db \
  -e PGDATA=/var/lib/postgresql/data/pgdata \
  -v ~/.mastersdb-data/:/var/lib/postgresql/data \
  masters-db
