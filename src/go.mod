module github.com/goharbor/harbor/src

go 1.20

require (
	github.com/FZambia/sentinel v1.1.0
	github.com/cenkalti/backoff/v4 v4.1.2 // indirect
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/uuid v1.3.0
	github.com/jackc/pgconn v1.9.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/prometheus/client_golang v1.14.0 // indirect
	github.com/stretchr/testify v1.8.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.22.0
	go.opentelemetry.io/otel v1.8.0
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.3.0
	go.opentelemetry.io/otel/sdk v1.8.0
	go.opentelemetry.io/otel/trace v1.8.0
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/net v0.7.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

require github.com/beego/beego/v2 v2.0.6

require (
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/errors v0.20.3
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.1.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/mapstructure v1.4.3 // indirect
	github.com/prometheus/client_model v0.3.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/shiena/ansicolor v0.0.0-20200904210342-c7312218db18 // indirect
	go.opentelemetry.io/contrib v0.22.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.3.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.3.0 // indirect
	go.opentelemetry.io/otel/internal/metric v0.22.0 // indirect
	go.opentelemetry.io/otel/metric v0.22.0 // indirect
	go.opentelemetry.io/proto/otlp v0.11.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	google.golang.org/genproto v0.0.0-20220502173005-c8bf987b8c21 // indirect
	google.golang.org/grpc v1.49.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v14.2.0+incompatible
	github.com/docker/distribution => github.com/distribution/distribution v2.8.1+incompatible
	github.com/goharbor/harbor => ../
	github.com/gomodule/redigo => github.com/gomodule/redigo v1.8.8
	google.golang.org/api => google.golang.org/api v0.0.0-20160322025152-9bf6e6e569ff
)
