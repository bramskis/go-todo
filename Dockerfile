FROM golang:1-.20-alpine3.18 as build

WORKDIR /go/src/app

ENV CGO_ENABLED=0
COPY go.* ./
RUN go mod download

COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN --mount=type=cache,target=/root/.cache/go-build GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o app

FROM alpine:3.18
COPY --from=build /go/src/app/app /usr/local/bin/app

EXPOSE 3000
ENTRYPOINT ["/usr/local/bin/app"]