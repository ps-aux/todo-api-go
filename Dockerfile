FROM alpine:3.11.2

COPY todo-api-go /app

ENTRYPOINT ["/app"]
