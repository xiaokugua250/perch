p2p 网络 模块

### 文档说明

|          |                   | 
|----------|:----------------: |
| 作者     |  liangdu          |  
| 联系邮箱 |liangdu@nscc-gz.cn或liangdu1992@gmail.com |
| 发布日期 |    2020-07-09     | 
| 版权所有 | nscc-gz@copyright | 
| 备注     | 无                |


## libp2p库说明与使用
### libp2p 简介
 libp2p 特性和适用场景: 
 -[x] Use Serval Transports 适用于多种协议,TCP/UDP/QUIC/WebRTC等 
 -[x] Native Roaming 自适应网络，网络发生变动时程序、服务不需要做额外配置
 -[x] Runtime Freedom 运行时无关，运行平台/软件不影响网络 
 -[x] Protocol Muxing 协议复用.网络连接复用  如: stream multiplexing
 -[x] Work Offline 可自行发现节点，不需要中心服务器或注册服务 如: mdns 节点发现
 -[x] Encrypted Connections 连接加密，通信链路加密和节点加密认证   如: peer identity
 -[x] Upgrade Without Compromises 无感升级
 -[x] Work In the brower 可浏览器中运行
 -[x] Good For High Latency Scenarios 可应用于高延迟场景
 
 libp2p 的实现版本
 -[x] go-libp2p 
 -[x] js-libp2p
 -[x] nodejs-libp2p
 -[x] rust-libp2p
### libp2p 来源
### libp2p 中核心数据结构
- [x] 分布式hash表(Distributed Hash Tables|dht)
- [x] Merkle DAGS
### libp2p核心概念
* 传输(Transport)
 参考:  
   - https://docs.libp2p.io/concepts/transport/
* NAT转换(NAT Traversal)
  参考:  
   - https://docs.libp2p.io/concepts/nat/
* 通信安全(Secure Communication)
  参考:  
   - https://docs.libp2p.io/concepts/nat/#automatic-router-configuration
* 中继传输(Circuit Relay)
  参考:  
   - https://docs.libp2p.io/concepts/circuit-relay/
* 传输协议(Protocols)  
  libp2p的核心协议有:
  1. ping 协议
  2. Identify 协议
  3. secio协议
  4. kad-dht 协议
  5. Circuit Relay 协议  
  参考:  
   - https://docs.libp2p.io/concepts/protocols/
* 节点标识(Peer Identify)
 参考:  
   - https://docs.libp2p.io/concepts/peer-id/
*  内容路由 (Content Routing)

* 节点路由 (Peer Routing)

* 地址标识(Addressing)
 参考:  
   - https://docs.libp2p.io/concepts/addressing/
* 安全性考虑(Security Considerations)
 参考:  
   - https://docs.libp2p.io/concepts/security-considerations/
* 发布订阅模式 (Publish/Subscribe)
 参考:  
   - https://docs.libp2p.io/concepts/stream-multiplexing/
* 流多路复用 (Stream Mutiplexing)
 参考:  
   - https://docs.libp2p.io/concepts/stream-multiplexing/

### libp2p 使用
- [x] libp2p 库使用通用流程

* 传输(Transport)
* NAT转换(NAT Traversal)
* 通信安全(Secure Communication)
- [x] 启用节点公私钥认证和传输流量加密
```$xslt
	security := libp2p.Security(secio.ID, secio.New) // 传输加密
	if options.NodeIdentifyPrivStr != "" { //节点身份ID设置
		priv, _, err := utils.GenSecurekeysByStr(options.NodeIdentifyPrivStr)
		p2pnetworkhostOptions = append(p2pnetworkhostOptions, libp2p.Identity(priv))
	} 
	//设置节点启动配置 basic node host
	p2pnetworkhostOptions = append(p2pnetworkhostOptions, security, muxers, transports, listenAddrOptions)
	networknode, err = p2pnetwork.CreatenetworkNodehost(ctx, dhtObj, p2pnetworkhostOptions)
}
```
* 中继传输(Circuit Relay)
 -[x] 使用中继节点方式进行节点联通与传输
 ```$xslt
	if option.PeerID != "" { //从命令行参数中获取中转节点PEER
		if option.RelayID != "" { //存在中继节点，则通过中继节点连接
			rawPeerIDWithIP := strings.Split(option.PeerID, "/")
			rawPeerID := rawPeerIDWithIP[len(rawPeerIDWithIP)-1]
			ID, err := peerstore.Decode(rawPeerID)
			logger.Printf("relayID is %s,and peerID is %s and targeID is %s\n", option.RelayID, option.PeerID, ID)
			relayAddr, err := ma.NewMultiaddr(option.RelayID + "/p2p-circuit" + option.PeerID)
			//根据地址信息拿到
			//peer, err := peerstore.AddrInfoFromP2pAddr(relayAddr)
			peer := peerstore.AddrInfo{
				ID:    ID,
				Addrs: []ma.Multiaddr{relayAddr},
			}
			if err = p2pNode.Connect(option.ctx, peer); err != nil {
				logger.Errorf("node %s connect to peer node %s with relayaddr  %s failed ,error is %s ...\n", p2pNode.ID(), option.PeerID, relayAddr, err.Error())
			}
		}
```
* 传输协议(Protocols)
 -[x] 传输协议标识方法
 ```$xslt
	transports := libp2p.ChainOptions(
		libp2p.Transport(tcp.NewTCPTransport),
		libp2p.Transport(ws.New),
	)
	listenAddr := libp2p.ListenAddrStrings(
		"/ip4/172.16.171.94/tcp/0", "/ip6/::/tcp/0/ws")
	p2pOptions = append(p2pOptions, muxers, security, listenAddr, transports)
	
	basicHost, err := libp2p.New(ctx, p2pOptions...)
		
```
* 节点标识(Peer Identify)
 -[x] 采用自定义公私钥方式进行节点标识
 ```$xslt
	// Generate a key pair for this host
	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		return nil, nil, err
	}

	ctx := context.Background()

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", listenPort)),
		libp2p.Identity(priv),
		libp2p.DefaultTransports,
		libp2p.DefaultMuxers,
		libp2p.DefaultSecurity,
		libp2p.NATPortMap(),
	}

	basicHost, err := libp2p.New(ctx, opts...)
	if err != nil {
		return nil, nil, err
	}
```
*  内容路由 (Content Routing)
 -[x] 基于DHT数据结构的内容发现
 ```$xslt
    dstore := dsync.MutexWrap(ds.NewMapDatastore())
	dhtObj, err := dht.New(ctx, basicHost, dht.Datastore(dstore), dht.Mode(dht.ModeServer))

	// step 3构建 需要存储在分布式dht 数据结果中的数据信息
	//dhtObj := dht.NewDHT(ctx, basicHost, dstore)
	data := []byte("this is some test content")
	hash, _ := mh.Sum(data, mh.SHA2_256, -1)
	contentId := cid.NewCidV1(cid.DagCBOR, hash)
	// step4 发布/共享该数据
	if err = dhtObj.Provide(ctx, contentId, false); err != nil {
		log.Error(err)
	}
	// Make the routed host
	//构造 routed host ，routedhost 为包含dht信息的basichost
	routedHost := rhost.Wrap(basicHost, dhtObj)
	// Bootstrap the host
    dhtObj.Bootstrap(ctx)
	// Build host multiaddress
	hostAddr, _ := ma.NewMultiaddr(fmt.Sprintf("/ipfs/%s", routedHost.ID().Pretty()))

	//dht_1 := dht.New(ctx, basicHost_1, dstore)
	basicHost_addressinfo := peer.AddrInfo{
		ID:    routedHost.ID(),
		Addrs: routedHost.Addrs(),
	}
	//works both 创建相同的dht 对象
	//dhtobj1, err := dht.New(ctx, basicHost_1, dht.Datastore(dstore_1), dht.Mode(dht.ModeServer), dht.BootstrapPeers(basicHost_addressinfo))
	dhtobj1, err := dht.New(ctx, basicHost_1, dht.Datastore(dstore_1), dht.Mode(dht.ModeServer))

	//dtbatch,err := dstore.Batch()
	//dhtobj1  = dht.NewDHTClient(ctx, basicHost_1, dtb)
	routedhost1 := rhost.Wrap(basicHost_1, dhtobj1) //关键点

	if err = dhtobj1.Bootstrap(ctx); err != nil { //关键点
		log.Error(err)
	}
	// 将两个节点进行连接，注意此处是目标节点信息 可通过节点发现、固定参数等形式获取
	if err = routedhost1.Connect(ctx, basicHost_addressinfo); err != nil { //关键点
		log.Error(err)
	}


	for {
		time.Sleep(5 * time.Second)
		//查询提供该内容的节点对象
		peers, err := dhtobj1.FindProviders(ctx, contentId) //关键点
		if err != nil {
			log.Error(err)
		}
		if len(peers) <= 0 {
			log.Println("found zero peers....", peers)
		}
		for _, peer := range peers {
			fmt.Printf("found peer %s provider contedt %s\n", peer, contentId.String())
		}

	}
```
* 节点路由 (Peer Routing)
 -[x] 基于mdns的节点发现与节点路由(原始)
 ```$xslt

type discoveryNotifee struct {
	PeerChan chan peer.AddrInfo
}

//interface to be called when new  peer is found
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	n.PeerChan <- pi
}

//Initialize the MDNS service
func initMDNS(ctx context.Context, peerhost host.Host, rendezvous string) chan peer.AddrInfo {
	// An hour might be a long long period in practical applications. But this is fine for us
	ser, err := discovery.NewMdnsService(ctx, peerhost, time.Hour, rendezvous)
	if err != nil {
		panic(err)
	}

	//register with service so that we get notified about peer discovery
	n := &discoveryNotifee{}
	n.PeerChan = make(chan peer.AddrInfo)

	ser.RegisterNotifee(n)
	return n.PeerChan
}

```
 -[x] 基于mdns的节点发现与节点路由(dht)
 ```$xslt
// Start a DHT, for use in peer discovery. We can't just make a new DHT
	// client because we want each peer to maintain its own local copy of the
	// DHT, so that the bootstrapping node of the DHT can go down without
	// inhibiting future peer discovery.
	//kademliaDHT, err := dht.New(ctx, host)
	kademliaDHT, err := dht.New(option.ctx, p2pNode, dht.Mode(dht.ModeServer))
	//kademliaDHT, err :=dht.New(ctx, host,dht.Mode(dht.ModeServer),dht.ProtocolPrefix(protocol.ID(config.ProtocolID)))


	// Bootstrap the DHT. In the default configuration, this spawns a Background
	// thread that will refresh the peer table every five minutes.
	logger.Debug("Bootstrapping the DHT")
	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		panic(err)
	}
	// We use a rendezvous point "meet me here" to announce our location.
	// This is like telling your friends to meet you at the Eiffel Tower.
	logger.Info("Announcing ourselves...")
	routingDiscovery := discovery.NewRoutingDiscovery(kademliaDHT)

	discovery.Advertise(ctx, routingDiscovery, config.RendezvousString)
	peerChan, err := routingDiscovery.FindPeers(ctx, config.RendezvousString)

```
* 地址标识(Addressing)
 -[x] 标识节点监听地址
 ```$xslt
	// Generate a key pair for this host
	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		return nil, nil, err
	}
	ctx := context.Background()
	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", listenPort)),
		libp2p.Identity(priv),
	}
	basicHost, err := libp2p.New(ctx, opts...)
	if err != nil {
		return nil, nil, err
	}

```

* 安全性考虑(Security Considerations)
* 发布订阅模式 (Publish/Subscribe)
 -[x]发布订阅模式使用实例
 ```$xslt
// 发布订阅模式 ，该处可选
	//	pubs ,err := pubsub.NewGossipSub(ctx,p2pnetwork.NetworkBasicHost)
	//pubs, err := pubsub.NewGossipSub(ctx, p2pnetwork.NetworkBasicHost)
	pubs, err := PubsubgossipGen(ctx, networknode.BasicNodeHost)
	if err != nil {
		log.Error(err)
	}
	sub, topsub, err := PubsubtopicsJoin(pubs, Pubsub_Default_Topic)
	if err != nil {
		log.Error(err)
	}

	// 设置mdns发现处理方法
	err = MdnsDiscoverySetup(ctx, networknode.BasicNodeHost, DiscoveryInterval, DiscoveryServiceTag)
	if err != nil {
		log.Error(err)
	}

	go func() {
		for {
			msg := new(PubsubMessage)
			msg.SenderPeer = networknode.BasicNodeHost.ID().Pretty()
			msg.PMessageStr = "hello world"
			msg.SenderFrom = "from localhost"
			err = PubsubTopicPubish(ctx, *msg, topsub, nil)
			if err != nil {
				log.Error(err)
			}
			time.Sleep(3 * time.Second)
		}

	}()
	msgChan := make(chan interface{})

	go PubsubMsgHandler(sub, ctx, networknode.BasicNodeHost, msgChan)

	signalChan := make(chan os.Signal, 1)
	errChan := make(chan error, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case msg := <-msgChan:
			fmt.Println("msg from msg chan is:", msg)
			for _, peersFind := range PubsubPeersList(pubs, Pubsub_Default_Topic) {
				peerinfo, err := dhtObj.FindPeer(ctx, peersFind)
				if err != nil {
					continue
				}
				if err = networknode.BasicNodeHost.Connect(ctx, peerinfo); err != nil {
					fmt.Print("connect b")
				}
				fmt.Printf("peer found by dht is %s\n", peerinfo.String())
			}
		}
	}
}

```
* 流多路复用 (Stream Mutiplexing)

核心示例代码如下:
`

`




## 参考资料
1. https://libp2p.io/  [官方网站]
2. https://github.com/libp2p [官方github]
3. https://github.com/libp2p/go-libp2p [libp2p golang 语言实现]
4. https://github.com/libp2p/go-libp2p-examples [go-libp2p 官方example 库]
5. https://docs.ipfs.io/concepts/what-is-ipfs/ [ipfs 官方网站]


