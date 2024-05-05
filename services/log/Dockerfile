FROM golang:1.22.2-bookworm AS build

WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .

RUN go build -v -o log &&\
  chmod +x /app/log

FROM debian:bookworm-slim

COPY --from=build /app/log /app/log

ENTRYPOINT [ "/app/log" ]
