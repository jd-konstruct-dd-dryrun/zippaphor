# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY *.go ./
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} \
    go build -ldflags="-s -w" -o zippaphor .

# Runtime stage
FROM alpine:3.23
LABEL org.opencontainers.image.source=https://github.com/jd-konstruct-dd-dryrun/zippaphor
LABEL org.opencontainers.image.description="simple golang http server"
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/zippaphor /usr/bin/zippaphor
EXPOSE 8080
ENTRYPOINT ["/usr/bin/zippaphor"]
