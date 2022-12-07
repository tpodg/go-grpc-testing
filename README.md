## Server
* GRPC running on port 9090
* REST running on port 8081
  * ANY /rest

## Client
* REST running on port 8080
  * ANY /rest calls server over rest
  * ANY /grpc (with optional query parameter 'value') calls server over grpc

### Client Env
* GRPC_TARGET: Address of the grpc server (default: localhost:9090)
* GRPC_TLS: Use TLS (default: false)
* REST_TARGET: Address of the rest server (default: http://localhost:8081)