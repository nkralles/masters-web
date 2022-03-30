#!/usr/bin/env bash

mkdir -p "backups"
##mp-pg is the name of the container we are using
docker exec masters-db pg_dump -U mp -F t mp >./backups/masters-db-$(date +%Y-%m-%d).tar