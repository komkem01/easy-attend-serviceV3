FROM docker.io/library/golang:1 AS builder
WORKDIR /app
COPY . /app
RUN mkdir -p /app/dist
RUN go work vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -modcacherw -o ./dist/ .;
FROM gcr.io/distroless/static AS serve
WORKDIR /app
COPY --from=builder /app/dist/ /app/
