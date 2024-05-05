FROM golang:1.22.2-bookworm AS build

WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .

RUN go build -v -o mastodon &&\
  chmod +x /app/mastodon

FROM debian:bookworm-slim

COPY --from=build /app/mastodon /app/mastodon

ENTRYPOINT [ "/app/mastodon" ]

