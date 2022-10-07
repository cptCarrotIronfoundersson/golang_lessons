FROM golang:1.16.5

WORKDIR /opt/src

ADD ../migrations/* .

RUN go get -u github.com/pressly/goose/cmd/goose

RUN chmod +x ./migrate.sh

CMD ./migrate.sh