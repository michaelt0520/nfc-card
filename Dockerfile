FROM golang:latest

LABEL maintainer="Michael Tran <mictran205@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./
COPY configs/.env_sample.sh configs/.env.production.sh

RUN go mod download

COPY . .

EXPOSE 80

CMD ["make", "start"]
