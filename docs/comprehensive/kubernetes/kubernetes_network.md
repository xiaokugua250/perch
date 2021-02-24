# k8s 网络

## K8S CNI 
### CNI简介
Kubernetes中的网络插件有几种类型：
 - CNI 插件：遵守容器网络接口（Container Network Interface，CNI） 规范，其设计上偏重互操作性。
Kubernetes 遵从 CNI 规范的 v0.4.0 版本。
- Kubenet 插件：使用 bridge 和 host-local CNI 插件实现了基本的 cbr0。
  
CNI 插件必须实现一个可执行文件，这个文件可以被容器管理系统（例如 rkt 或 Kubernetes）调用。
CNI 插件负责将网络接口插入容器网络命名空间（例如，veth 对的一端），并在主机上进行任何必要的改变（例如将 veth 的另一端连接到网桥）。然后将 IP 分配给接口，并通过调用适当的 IPAM 插件来设置与 “IP 地址管理” 部分一致的路由。
### CNI要点
- 接口定义
>>> CNI 的接口中包括以下几个方法：
type CNI interface {AddNetworkList (net *NetworkConfigList, rt *RuntimeConf) (types.Result, error)
    DelNetworkList (net *NetworkConfigList, rt *RuntimeConf) error
    AddNetwork (net *NetworkConfig, rt *RuntimeConf) (types.Result, error)
    DelNetwork (net *NetworkConfig, rt *RuntimeConf) error
}
该接口只有四个方法，添加网络、删除网络、添加网络列表、删除网络列表。
- 设计考虑
>>>CNI 设计的时候考虑了以下问题：
1.容器运行时必须在调用任何插件之前为容器创建一个新的网络命名空间。
2.然后，运行时必须确定这个容器应属于哪个网络，并为每个网络确定哪些插件必须被执行。
3.网络配置采用 JSON 格式，可以很容易地存储在文件中。网络配置包括必填字段，如 name 和 type 以及插件（类型）。网络配置允许字段在调用之间改变值。为此，有一个可选的字段 args，必须包含不同的信息。
4.容器运行时必须按顺序为每个网络执行相应的插件，将容器添加到每个网络中。
5.在完成容器生命周期后，运行时必须以相反的顺序执行插件（相对于执行添加容器的顺序）以将容器与网络断开连接。
6.容器运行时不能为同一容器调用并行操作，但可以为不同的容器调用并行操作。
7.容器运行时必须为容器订阅 ADD 和 DEL 操作，这样 ADD 后面总是跟着相应的 DEL。 DEL 可能跟着额外的 DEL，但是，插件应该允许处理多个 DEL（即插件 DEL 应该是幂等的）。
8.容器必须由 ContainerID 唯一标识。存储状态的插件应该使用（网络名称，容器 ID）的主键来完成。
运行时不能调用同一个网络名称或容器 ID 执行两次 ADD（没有相应的 DEL）。换句话说，给定的容器 ID 必须只能添加到特定的网络一次。
