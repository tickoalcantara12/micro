module github.com/tickoalcantara12/micro/profile/platform

go 1.18

require (
	github.com/go-redis/redis/v8 v8.11.3
	github.com/tickoalcantara12/micro/plugin/etcd/v3 v3.0.0-20220816083658-7f875afc6a39
	github.com/tickoalcantara12/micro/plugin/postgres/v3 v3.0.0-20210825142032-d27318700a59
	github.com/tickoalcantara12/micro/plugin/prometheus/v3 v3.0.0-20201217215412-2f7ad18595ff
	github.com/tickoalcantara12/micro/plugin/redis/blocklist/v3 v3.0.0-20220908113659-758df06ca38b
	github.com/tickoalcantara12/micro/plugin/redis/broker/v3 v3.0.0-20210622092058-b12502169757
	github.com/tickoalcantara12/micro/plugin/redis/stream/v3 v3.0.0-20210622092058-b12502169757
	github.com/tickoalcantara12/micro/plugin/s3/v3 v3.0.0-20220203170347-2c86e5fb6d13
	github.com/tickoalcantara12/micro/v3 v3.13.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/urfave/cli/v2 v2.3.0
)

require (
	github.com/aws/aws-sdk-go v1.23.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/coreos/go-semver v0.2.0 // indirect
	github.com/coreos/go-systemd/v22 v22.0.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/evanphx/json-patch/v5 v5.0.0 // indirect
	github.com/go-acme/lego/v3 v3.4.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/improbable-eng/grpc-web v0.13.0 // indirect
	github.com/jmespath/go-jmespath v0.0.0-20180206201540-c2b33e8439af // indirect
	github.com/lib/pq v1.8.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/hashstructure v1.0.0 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.7.1 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.10.0 // indirect
	github.com/prometheus/procfs v0.2.0 // indirect
	github.com/rs/cors v1.7.0 // indirect
	github.com/russross/blackfriday/v2 v2.0.1 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/stoewer/go-strcase v1.2.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/uber/jaeger-client-go v2.29.1+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.etcd.io/bbolt v1.3.5 // indirect
	go.etcd.io/etcd v0.5.0-alpha.5.0.20200425165423-262c93980547 // indirect
	go.uber.org/atomic v1.6.0 // indirect
	go.uber.org/multierr v1.5.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/net v0.0.0-20210428140749-89ef3d95e781 // indirect
	golang.org/x/sys v0.0.0-20210423082822-04245dca01da // indirect
	golang.org/x/text v0.3.6 // indirect
	google.golang.org/genproto v0.0.0-20200806141610-86f49bd18e98 // indirect
	google.golang.org/grpc v1.40.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

replace github.com/soheilhy/cmux => github.com/soheilhy/cmux v0.1.5

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
