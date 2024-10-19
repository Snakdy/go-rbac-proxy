# Go RBAC Proxy

The `go-rbac-proxy` is a sidecar that offloads authorization decisions away from your application.

## Getting started

The `go-rbac-proxy` exposes a gRPC interface that you can use to create and query roles.

### Configuration

```yaml
globals:
  SUPER:
    - SUDO
adapter:
  mode: redis
  redis:
    addrs:
      - localhost:6379
  postgres:
    dsn: "host=localhost port=5432 user=postgres password=hunter2 sslmode=disable"
```