FROM golang:1.22-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -ldflags "-w -s" -o /app

# hadolint ignore=DL3007
FROM alpine:latest AS build-release-stage

WORKDIR /opt/app

COPY --from=build-stage /app .

RUN chmod +x app

ENTRYPOINT ["./app"]
