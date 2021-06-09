FROM golang:latest AS build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on

WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
WORKDIR /app
RUN go build -o out/ ./cmd/scraper

FROM alpine:latest

RUN adduser --disabled-password --no-create-home --gecos '' user
RUN apk update \
    && apk --no-cache add ca-certificates

RUN mkdir /app && chown user:user /app
WORKDIR /app
USER user
COPY --from=build /app/out/scraper .

CMD ["/app/scraper"]





