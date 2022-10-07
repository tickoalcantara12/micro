module github.com/tickoalcantara12/micro/plugin/etcd/v3

go 1.15

require (
	github.com/tickoalcantara12/micro/v3 v3.0.4
	github.com/mitchellh/hashstructure v1.0.0
	go.etcd.io/etcd v0.5.0-alpha.5.0.20200425165423-262c93980547
	go.uber.org/zap v1.16.0
)

replace google.golang.org/grpc v1.30.0 => google.golang.org/grpc v1.26.0

replace google.golang.org/grpc v1.40.0 => google.golang.org/grpc v1.26.0

replace github.com/soheilhy/cmux => github.com/soheilhy/cmux v0.1.5

replace github.com/tickoalcantara12/micro/v3 => ../..

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
