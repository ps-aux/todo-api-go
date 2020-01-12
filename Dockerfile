#FROM alpine:3.11.2
 # TODO change to slimmer image which works
FROM ubuntu:18.04
#FROM golang:1.12.0-alpine3.9

COPY todo-api-go /app

ENTRYPOINT ["/app"]
