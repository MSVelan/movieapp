FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y librdkafka1 && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY main .

COPY configs/. .

# COPY --from=builder /rating-app .

EXPOSE 8082
