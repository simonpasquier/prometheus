module github.com/simonpasquier/prometheus/discovery

go 1.12

replace github.com/simonpasquier/prometheus/sdk => ../sdk

replace k8s.io/klog => github.com/simonpasquier/klog-gokit v0.1.0

require (
	contrib.go.opencensus.io/exporter/ocagent v0.5.1 // indirect
	github.com/Azure/azure-sdk-for-go v23.2.0+incompatible
	github.com/Azure/go-autorest v11.2.8+incompatible
	github.com/aws/aws-sdk-go v1.21.8
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-kit/kit v0.8.0
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/golang/groupcache v0.0.0-20180924190550-6f2cf27854a4 // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gophercloud/gophercloud v0.2.0
	github.com/hashicorp/consul/api v1.1.0
	github.com/hashicorp/go-msgpack v0.5.4 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/miekg/dns v1.1.15
	github.com/mwitkow/go-conntrack v0.0.0-20161129095857-cc309e4a2223
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.0.0
	github.com/prometheus/common v0.4.1
	github.com/samuel/go-zookeeper v0.0.0-20180130194729-c4fab1ac1bec
	github.com/simonpasquier/prometheus/sdk v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.3.0
	golang.org/x/crypto v0.0.0-20190325154230-a5d413f7728c // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/api v0.7.0
	gopkg.in/fsnotify/fsnotify.v1 v1.4.7
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/yaml.v2 v2.2.2
	k8s.io/api v0.0.0-20190718183219-b59d8169aab5
	k8s.io/apimachinery v0.0.0-20190612205821-1799e75a0719
	k8s.io/client-go v0.0.0-20190620085101-78d2af792bab
	k8s.io/utils v0.0.0-20190308190857-21c4ce38f2a7 // indirect
)
