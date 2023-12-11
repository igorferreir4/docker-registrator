FROM --platform=$BUILDPLATFORM golang:1.21.5-alpine3.19 AS builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM

LABEL org.opencontainers.image.version="v8.1"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.title="Consul Registrator"
LABEL org.opencontainers.image.url="https://github.com/igorferreir4/docker/tree/main/imagens/registrator"
LABEL org.regitrator.maintainer="Igor Ferreira"

WORKDIR /go/src/github.com/mario-ezquerro/registrator/
COPY . .
RUN \
	apk add --no-cache curl git \
	&& curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
	&& dep ensure -vendor-only \
	&& CGO_ENABLED=0 GOOS=linux go build \
		-a -installsuffix cgo \
		-mod=mod \
		-ldflags "-X main.Version=$(cat VERSION)" \
		-o bin/registrator \
		.

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
COPY --from=builder /go/src/github.com/mario-ezquerro/registrator/bin/registrator /bin/registrator

ENTRYPOINT ["/bin/registrator"]