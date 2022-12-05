FROM golang:1.19 as build

ARG TARGETOS=linux
ARG TARGETARCH=amd64
ENV CGO_ENABLED=0 \
    GRPC_HEALTH_PROBE_VERSION=v0.4.14

RUN wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-$TARGETOS-$TARGETARCH && \
    chmod +x /bin/grpc_health_probe

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN go build -o client *.go


FROM alpine:3.17
COPY --from=build /bin/grpc_health_probe /bin/grpc_health_probe
WORKDIR /app
COPY --from=build /app/client .

ENTRYPOINT [ "./client" ]