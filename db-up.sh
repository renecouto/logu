#!/bin/bash

docker-compose up -d

docker-compose exec -T db psql postgres postgres -f - < psql/schemas.sql