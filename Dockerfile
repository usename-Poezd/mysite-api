FROM golang:1.16.7-alpine3.14 AS build
WORKDIR /src
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app ./cmd/app

######## Start a new stage from scratch #######
FROM alpine
WORKDIR /src

COPY --from=build /app  .
EXPOSE 8080

CMD ["./app"]