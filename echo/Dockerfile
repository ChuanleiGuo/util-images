FROM golang:1.14-alpine AS Builder
COPY . /src
WORKDIR /src
RUN go build -o echo .

FROM alpine:latest
COPY --from=Builder /src/echo /app/
WORKDIR /app
ENTRYPOINT ["/app/echo"]