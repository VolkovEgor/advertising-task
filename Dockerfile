  
FROM golang:latest

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# go dependencies
RUN go mod download -x

# build go app
RUN go build -o app ./cmd/main.go

# swagger
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN make swag

CMD ./app docker_config


