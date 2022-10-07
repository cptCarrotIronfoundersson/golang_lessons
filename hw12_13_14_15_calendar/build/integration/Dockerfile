# Собираем в гошке
FROM golang:1.16.2 as build
WORKDIR tests
COPY . .
ENTRYPOINT ["go", "test",  "-count=1" ,"-race", "./tests/integration/..."]