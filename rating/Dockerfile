# -------- Stage 1: Build --------
# FROM golang:1.24 AS builder

# RUN apt-get update && apt-get install -y librdkafka-dev

# WORKDIR /app

# COPY go.mod go.sum ./
# COPY . .
# RUN go mod download

# # Build from the correct sub-package
# RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /rating-app ./rating/cmd

# -------- Stage 2: Runtime --------
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y librdkafka1 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY main .

COPY configs/. .

# COPY --from=builder /rating-app .

EXPOSE 8082
CMD ["./main"]
