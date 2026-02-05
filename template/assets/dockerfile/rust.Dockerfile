FROM rust:1.93-trixie AS build
WORKDIR /build

COPY Cargo.lock Cargo.toml ./
RUN mkdir src && echo "fn main() {}" > src/main.rs
RUN cargo fetch

COPY src src
RUN cargo build --release && \
    strip target/release/app && \
    cp ./target/release/app /app

# hadolint ignore=DL3007
FROM gcr.io/distroless/cc:nonroot AS deploy

COPY --from=build /app /

EXPOSE 3000
ENTRYPOINT ["/app"]
