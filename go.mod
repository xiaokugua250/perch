module perch

go 1.14

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/antchfx/xmlquery v1.3.2 // indirect
	github.com/coreos/bbolt v1.3.5 // indirect
	github.com/coreos/etcd v3.3.10+incompatible
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/corpix/uarand v0.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gdamore/tcell v1.3.0
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-redis/redis/v8 v8.4.0
	github.com/gocolly/colly v1.2.0
	github.com/golang/protobuf v1.4.3
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/googleapis/gnostic v0.5.1 // indirect
	github.com/gorilla/handlers v1.5.0
	github.com/gorilla/mux v1.7.4
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/icrowley/fake v0.0.0-20180203215853-4178557ae428
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/ipfs/go-cid v0.0.6
	github.com/ipfs/go-datastore v0.4.4
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/kr/pty v1.1.3
	github.com/libp2p/go-libp2p v0.9.6
	github.com/libp2p/go-libp2p-core v0.6.0
	github.com/libp2p/go-libp2p-discovery v0.4.0
	github.com/libp2p/go-libp2p-kad-dht v0.8.2
	github.com/libp2p/go-libp2p-mplex v0.2.3
	github.com/libp2p/go-libp2p-nat v0.0.6
	github.com/libp2p/go-libp2p-pubsub v0.3.2
	github.com/libp2p/go-libp2p-secio v0.2.2
	github.com/libp2p/go-libp2p-swarm v0.2.7
	github.com/libp2p/go-libp2p-yamux v0.2.8
	github.com/libp2p/go-tcp-transport v0.2.0
	github.com/libp2p/go-ws-transport v0.3.1
	github.com/libp2p/go-yamux v1.3.7
	github.com/mitchellh/mapstructure v1.1.2
	github.com/mottet-dev/medium-go-colly-basics v0.0.0-20190610193548-ce60257193c9
	github.com/multiformats/go-multiaddr v0.2.2
	github.com/multiformats/go-multiaddr-net v0.1.5
	github.com/multiformats/go-multibase v0.0.3
	github.com/multiformats/go-multihash v0.0.13
	github.com/pkg/errors v0.9.1
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/rivo/tview v0.0.0-20200528200248-fe953220389f
	github.com/shirou/gopsutil v2.20.8+incompatible
	github.com/sirupsen/logrus v1.6.0
	github.com/soheilhy/cmux v0.1.4 // indirect
	github.com/spf13/viper v1.3.2
	github.com/temoto/robotstxt v1.1.1 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200427203606-3cfed13b9966 // indirect
	github.com/urfave/cli/v2 v2.2.0
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	github.com/zellyn/kooky v0.0.0-20201108220156-bec09c12c339
	go.etcd.io/bbolt v1.3.5 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/tools v0.0.0-20200708183856-df98bc6d456c // indirect
	google.golang.org/grpc v1.29.1
	google.golang.org/protobuf v1.25.0
	gopkg.in/yaml.v2 v2.3.0
	gorm.io/driver/mysql v1.0.1
	gorm.io/gorm v1.20.1
	k8s.io/api v0.19.0
	k8s.io/apimachinery v0.19.0
	k8s.io/client-go v0.19.0
)

replace github.com/coreos/bbolt v1.3.5 => github.com/coreos/bbolt v1.3.3
