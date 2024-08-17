FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -ldflags "-w -s" -o /app

# hadolint ignore=DL3007
FROM gcr.io/distroless/static-debian11:latest AS deploy
COPY --from=build /app /

CMD ["/app"]

## use this if your app is CGO_ENABLED=1
#FROM alpine:latest AS deploy
#
#WORKDIR /opt/app
#COPY --from=build /app .
#
#RUN chmod +x app
#ENTRYPOINT ["/app"]
