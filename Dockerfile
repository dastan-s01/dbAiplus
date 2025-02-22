FROM golang:1.23.3-alpine3.19 AS build-stage
WORKDIR /dbAiplus
COPY go.mod go.sum /
RUN go mod download
COPY . .
RUN go build -o main ./cmd/main.go

FROM alpine:3.19 AS runner
WORKDIR /app
COPY --from=build-stage /dbAiplus/db/migrations db/migrations
COPY --from=build-stage /dbAiplus/main .
EXPOSE 8080
CMD ["/app/main"]

