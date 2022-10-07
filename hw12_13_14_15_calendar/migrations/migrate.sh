#!/bin/bash

DBSTRING="postgresql://postgres:PASSWORD@db/otus?sslmode=disable"

goose postgres "$DBSTRING" up
goose postgres "$DBSTRING" status
#ping db