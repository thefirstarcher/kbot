FROM quay.io/projectquay/golang:1.27 AS builder

WORKDIR /go/src/app
COPY . .

ARG VERSION
ARG TARGETARCH
RUN make build TARGETARCH=$TARGETARCH VERSION=$VERSION

FROM scratch
COPY --from=builder /go/src/app/kbot .

COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

ENTRYPOINT ["./kbot", "start"]

