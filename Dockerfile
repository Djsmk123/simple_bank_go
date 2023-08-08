# Build Stage
FROM golang:1.20 AS Builder
WORKDIR /app

COPY . .
RUN CGO_ENABLED=0  go build --o main main.go

# Run Stage
FROM alpine:3.18.3
WORKDIR /app
COPY --from=Builder /app/main .
COPY app.env .

EXPOSE 8080
CMD [ "/app/main" ]