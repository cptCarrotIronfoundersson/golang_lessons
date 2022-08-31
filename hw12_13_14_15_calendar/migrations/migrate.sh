#!/usr/bin/env bash

export GOOSE_DRIVER=postgres
export GOOSE_DBSTRING="user=postgres dbname=otus sslmode=disable password=PASSWORD"

sleep 5
/usr/local/bin/goose -dir ./migrations  up
/usr/local/bin/goose -dir ./migrations  status
