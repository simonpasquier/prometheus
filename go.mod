module github.com/prometheus/prometheus

go 1.13

require (
	github.com/Azure/azure-sdk-for-go v45.1.0+incompatible
	github.com/Azure/go-autorest/autorest v0.11.4
	github.com/Azure/go-autorest/autorest/adal v0.9.2
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/alecthomas/units v0.0.0-20190717042225-c3de453c63f4
	github.com/aws/aws-sdk-go v1.25.48
	github.com/cespare/xxhash v1.1.0
	github.com/dgryski/go-sip13 v0.0.0-20190329191031-25c5027a8c7b
	github.com/edsrzf/mmap-go v1.0.0
	github.com/go-kit/kit v0.9.0
	github.com/go-logfmt/logfmt v0.4.0
	github.com/go-openapi/strfmt v0.19.2
	github.com/gogo/protobuf v1.3.1
	github.com/golang/snappy v0.0.1
	github.com/google/pprof v0.0.0-20190723021845-34ac40c74b70
	github.com/gophercloud/gophercloud v0.3.0
	github.com/grpc-ecosystem/grpc-gateway v1.9.5
	github.com/hashicorp/consul/api v1.1.0
	github.com/influxdata/influxdb v1.7.7
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/json-iterator/go v1.1.8
	github.com/miekg/dns v1.1.15
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f
	github.com/oklog/run v1.0.0
	github.com/oklog/ulid v1.3.1
	github.com/opentracing-contrib/go-stdlib v0.0.0-20190519235532-cf7a6c988dc9
	github.com/opentracing/opentracing-go v1.1.0
	github.com/pkg/errors v0.8.1
	github.com/prometheus/alertmanager v0.18.0
	github.com/prometheus/client_golang v1.2.0
	github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4
	github.com/prometheus/common v0.7.0
	github.com/samuel/go-zookeeper v0.0.0-20190810000440-0ceca61e4d75
	github.com/shurcooL/httpfs v0.0.0-20190707220628-8d4bc4ba7749
	github.com/shurcooL/vfsgen v0.0.0-20181202132449-6a9ea43bcacd
	github.com/soheilhy/cmux v0.1.4
	go.opencensus.io v0.22.0 // indirect
	golang.org/x/net v0.0.0-20191004110552-13f9640d40b9
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys v0.0.0-20191022100944-742c48ecaeb7
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools v0.0.0-20190918214516-5a1a30219888
	google.golang.org/api v0.8.0
	google.golang.org/genproto v0.0.0-20190801165951-fa694d86fc64
	google.golang.org/grpc v1.22.1
	gopkg.in/alecthomas/kingpin.v2 v2.2.6
	gopkg.in/fsnotify/fsnotify.v1 v1.4.7
	gopkg.in/yaml.v2 v2.2.8
	k8s.io/api v0.18.9
	k8s.io/apimachinery v0.18.9
	k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
	k8s.io/klog v1.0.0
)

replace (
	github.com/golang/lint => golang.org/x/lint v0.0.0-20190409202823-959b441ac422
	k8s.io/api => k8s.io/api v0.18.9
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.9
	k8s.io/client-go => k8s.io/client-go v0.18.9
	k8s.io/klog => github.com/simonpasquier/klog-gokit v0.1.0
)
