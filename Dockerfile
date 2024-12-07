FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=1 go build -o main -a -ldflags '-linkmode external -extldflags "-static"' .


FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
COPY static/ static/
COPY templates/ templates/
EXPOSE 1323
CMD ["./main"]