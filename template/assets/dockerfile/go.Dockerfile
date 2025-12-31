FROM golang:1.25-alpine AS build
## for CGO
#FROM golang:1.25-trixie AS build

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -ldflags "-w -s" -o app
## for CGO
#RUN CGO_ENABLED=1 go build -ldflags "-w -s" -o /app

FROM gcr.io/distroless/static-debian13:nonroot AS deploy
## for CGO
#FROM gcr.io/distroless/base-debian13:nonroot AS deploy

COPY --from=build /build/app /

EXPOSE 3000
ENTRYPOINT ["/app"]
