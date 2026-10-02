FROM golang:1.27-alpine AS builder

WORKDIR /piss
#COPY go.mod go.sum ./
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /piss
COPY --from=builder /piss/main ./main
RUN chmod +x main

CMD ["./main"]
