FROM golang:1.13 AS build
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go test ./internal/... -cover fmt \
    && CGO_ENABLE=0 go build -tags netgo -ldflags="-d -s -w" -o main ./cmd/iris-backend.go

FROM alpine:latest AS runtime
RUN apk add --update --no-cache tzdata \
    && cp -rf /usr/share/zoneinfo/Europe/Paris /etc/localtime \
    && echo "Europe/Paris" >/etc/timezone
COPY --from=build /app/main /app/main
CMD ["/app/main"]
