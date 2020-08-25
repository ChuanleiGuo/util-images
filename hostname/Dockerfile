FROM golang:1.14-alpine AS Builder
COPY . /src
WORKDIR /src
RUN go build -o hostname .

FROM alpine:latest
COPY --from=Builder /src/hostname /app/
WORKDIR /app
ENTRYPOINT ["/app/hostname"]