FROM golang:1.23 as build

WORKDIR /app

RUN apt-get update -y && \
    apt-get install -y ffmpeg

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]