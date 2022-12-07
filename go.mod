module perch

go 1.14

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/antchfx/xmlquery v1.3.2 // indirect
	github.com/casbin/casbin/v2 v2.25.5
	github.com/casbin/gorm-adapter/v3 v3.2.5
	github.com/coreos/bbolt v1.3.5 // indirect
	github.com/coreos/etcd v3.3.13+incompatible
	github.com/corpix/uarand v0.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/dimiro1/banner v1.1.0
	github.com/gdamore/tcell v1.3.0
	github.com/go-ldap/ldap/v3 v3.2.4
	github.com/go-logr/logr v1.2.3
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/go-redis/redis/v8 v8.4.0
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gocolly/colly v1.2.0
	github.com/golang/protobuf v1.5.2
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/icrowley/fake v0.0.0-20180203215853-4178557ae428
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-datastore v0.5.0
	github.com/koding/websocketproxy v0.0.0-20181220232114-7ed82d81a28c
	github.com/kr/pty v1.1.8
	github.com/libp2p/go-libp2p v0.18.1
	github.com/libp2p/go-libp2p-core v0.14.0
	github.com/libp2p/go-libp2p-discovery v0.4.0
	github.com/libp2p/go-libp2p-kad-dht v0.8.2
	github.com/libp2p/go-libp2p-mplex v0.6.0
	github.com/libp2p/go-libp2p-nat v0.1.0
	github.com/libp2p/go-libp2p-pubsub v0.3.2
	github.com/libp2p/go-libp2p-secio v0.2.2
	github.com/libp2p/go-libp2p-swarm v0.10.2
	github.com/libp2p/go-libp2p-yamux v0.9.0
	github.com/libp2p/go-tcp-transport v0.5.1
	github.com/libp2p/go-ws-transport v0.6.0
	github.com/libp2p/go-yamux v1.4.1
	github.com/mattn/go-colorable v0.1.8
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/mitchellh/mapstructure v1.4.1
	github.com/mottet-dev/medium-go-colly-basics v0.0.0-20190610193548-ce60257193c9
	github.com/multiformats/go-multiaddr v0.5.0
	github.com/multiformats/go-multiaddr-net v0.2.0
	github.com/multiformats/go-multibase v0.0.3
	github.com/multiformats/go-multihash v0.0.15
	github.com/nsqio/go-nsq v1.0.8
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.19.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.12.2
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/rivo/tview v0.0.0-20200528200248-fe953220389f
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/shaj13/go-guardian/v2 v2.11.3
	github.com/shaj13/libcache v1.0.0
	github.com/shirou/gopsutil v2.20.8+incompatible
	github.com/sirupsen/logrus v1.8.1
	github.com/sony/sonyflake v1.0.0
	github.com/spf13/viper v1.7.0
	github.com/temoto/robotstxt v1.1.1 // indirect
	github.com/urfave/cli/v2 v2.2.0
	github.com/zellyn/kooky v0.0.0-20201108220156-bec09c12c339
	golang.org/x/crypto v0.0.0-20220315160706-3147a52a75dd
	golang.org/x/net v0.4.0 // indirect
	golang.org/x/time v0.0.0-20220609170525-579cf78fd858
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
	google.golang.org/grpc v1.50.1
	google.golang.org/protobuf v1.28.1
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.0.1
	gorm.io/gorm v1.20.7
	k8s.io/api v0.25.0
	k8s.io/apimachinery v0.25.0
	k8s.io/client-go v0.25.0
	sigs.k8s.io/controller-runtime v0.13.1
)

replace github.com/coreos/bbolt v1.3.5 => github.com/coreos/bbolt v1.3.3
