module github.com/simonpasquier/prometheus

replace github.com/simonpasquier/prometheus/discovery => ./discovery

replace github.com/simonpasquier/prometheus/sdk => ./sdk

require (
	github.com/OneOfOne/xxhash v1.2.5 // indirect
	github.com/alecthomas/units v0.0.0-20151022065526-2efee857e7cf
	github.com/cespare/xxhash v1.1.0
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/go-kit/kit v0.8.0
	github.com/go-logfmt/logfmt v0.4.0
	github.com/go-openapi/strfmt v0.19.0
	github.com/gogo/protobuf v1.2.1
	github.com/golang/snappy v0.0.1
	github.com/google/pprof v0.0.0-20181206194817-3ea8567a2e57
	github.com/grpc-ecosystem/grpc-gateway v1.8.5
	github.com/ianlancetaylor/demangle v0.0.0-20181102032728-5e5cf60278f6 // indirect
	github.com/influxdata/influxdb v1.7.7
	github.com/json-iterator/go v1.1.6
	github.com/mwitkow/go-conntrack v0.0.0-20161129095857-cc309e4a2223
	github.com/oklog/run v1.0.0
	github.com/opentracing-contrib/go-stdlib v0.0.0-20170113013457-1de4cc2120e7
	github.com/opentracing/opentracing-go v1.0.2
	github.com/pkg/errors v0.8.1
	github.com/prometheus/alertmanager v0.17.0
	github.com/prometheus/client_golang v1.0.0
	github.com/prometheus/client_model v0.0.0-20190129233127-fd36f4220a90
	github.com/prometheus/common v0.4.1
	github.com/prometheus/tsdb v0.10.0
	github.com/shurcooL/httpfs v0.0.0-20171119174359-809beceb2371
	github.com/shurcooL/vfsgen v0.0.0-20180825020608-02ddb050ef6b
	github.com/simonpasquier/prometheus/discovery v0.0.0-00010101000000-000000000000
	github.com/soheilhy/cmux v0.1.4
	github.com/stretchr/testify v1.3.0
	golang.org/x/net v0.0.0-20190613194153-d28f0bde5980
	golang.org/x/time v0.0.0-20181108054448-85acf8d2951c
	golang.org/x/tools v0.0.0-20190506145303-2d16b83fe98c
	google.golang.org/genproto v0.0.0-20190502173448-54afdca5d873
	google.golang.org/grpc v1.21.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/klog v0.3.1
)

replace k8s.io/klog => github.com/simonpasquier/klog-gokit v0.1.0
