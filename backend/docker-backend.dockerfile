FROM golang:1.21-alpine

RUN mkdir /app

COPY . /app

WORKDIR /app

CMD ["go", "run", "cmd/api/main.go"]
