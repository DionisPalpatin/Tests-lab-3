FROM golang:1.23 AS build

WORKDIR /go/src/app

COPY . .

RUN go build -o echo cmd/echo/main.go
RUN go build -o chi cmd/chi/main.go
RUN go build -o gin cmd/gin/main.go

# #=====#=====#=====#=====#=====#=====#=====#=====#=====#=====#=====#=====#=====

FROM debian:bookworm-slim

COPY --from=build /go/src/app/echo /app/echo
COPY --from=build /go/src/app/chi /app/chi
COPY --from=build /go/src/app/gin /app/gin

WORKDIR /app