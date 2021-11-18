FROM golang:latest AS build

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GO111MODULE=on

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o out/ ./cmd/scraper

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/out/scraper .

CMD ["/app/scraper"]





