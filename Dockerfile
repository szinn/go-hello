# Build Project
FROM golang:1.22.4-alpine as build
WORKDIR /go/src/github.com/szinn/go-hello

ARG TARGETOS
ARG TARGETARCH
ARG TARGETVARIANT=""

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=${TARGETOS} \
    GOARCH=${TARGETARCH} \
    GOARM=${TARGETVARIANT}

COPY . ./
RUN go mod download
RUN go build -ldflags="-s -w" -o /go-hello

# Final Image
FROM gcr.io/distroless/static:nonroot
USER nonroot:nonroot
COPY --from=build --chown=nonroot:nonroot /go-hello /app/
EXPOSE 8080

CMD ["/app/go-hello"]
LABEL \
    org.opencontainers.image.title="go-hello" \
    org.opencontainers.image.source="https://github.com/szinn/go-hello"
