FROM rust:1.81-alpine3.20 AS build

WORKDIR /app

COPY Cargo.lock Cargo.toml ./
RUN mkdir src && echo "fn main() {}" > src/main.rs
RUN cargo fetch

RUN apk update && apk add musl-dev libressl-dev --no-cache

COPY src src
RUN cargo build --release && \
    strip target/release/qa-api-rs && \
    cp ./target/release/qa-api-rs /qa-api-rs

# hadolint ignore=DL3007
FROM alpine:latest AS deploy
COPY --from=build /qa-api-rs /

CMD ["/qa-api-rs"]
