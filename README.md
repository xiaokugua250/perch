# perch

<div align="center"><img src ="./asserts/logo/perch_logo.jpg"/></div>

`Golang Code And Others`,`k8s`,`colly`,`grpc`,`mysql`,`micro-service`,`docker`,`docker-compose`
## 项目内容说明
 &emsp;&emsp; 本项目为个人项目，项目以Golang为主，包括go-web、go-spider、grpc等基础语言编码内容以及kubernetes、etcd等高阶内容。主要涉及到的内容有：
 1. 基于`Golang`的微服务开发
 2. 基于`docker`,`kubernetes`的容器管理平台开发
 3. 基于`colly`的go语言爬虫开发
 4. 基于`grpc`的分布式服务调用和任务分发
 
&emsp;&emsp;项目主要目的是对自己技能的总结和部分想法的实现。  
&emsp;&emsp;目前项目部署实例为 [https://z-gour.com](https://z-gour.com)    
&emsp;&emsp;部署方式为[minkube](https://minikube.sigs.k8s.io/)集群中以kubernete容器方式进行部署。采用到的kubernetes资源有
- pod  运行底层前端服务和后端服务
- service 集群内网络访问
- ingress  允许外部网络访问
- [cert-manager](https://cert-manager.io/) 证书签发和续期

在[minikube](https://minikube.sigs.k8s.io/) 或[kubernete](https://kubernetes.io/zh/)集群中部署可以参考deploy目录中的k8s_deploy目录。

下面对项目中各目录对应信息进行说明:
* api:  项目API
* cmd: 部分项目入口程序和可执行程序
* configs: 项目配置
* database: 数据库相关代码
* deploy: 项目部署文件，包括dockerfile、docker-compose.yaml等
* docs: 项目文档
* internal: 内部使用库
* pkg: 项目内部库
* scripts:  项目脚本
* third_party: 第三方工具，如`ETCD`等
* tools: 项目通用工具包
* web: 项目内嵌微服务集合
* website: 项目前端文件
* .gitignore: ignore文件
* Makefile: Makefile文件
* go.mod: go mod文件
* version.sh: 生成go_version.go脚本 
  
### 后端
本项目的后端采用Golang语言开发，涉及到的工具包有
* [gorilla/mux](github.com/gorilla/mux)
* [jwt-go](https://github.com/dgrijalva/jwt-go)
* [spf13/viper](github.com/spf13/viper)
* [logrus](github.com/sirupsen/logrus)
* [gorm ](https://github.com/go-gorm/gorm)
* [client-go](https://github.com/kubernetes/client-go)
* [grpc/grpc-go](https://github.com/grpc/grpc-go)
* [gocolly/colly](https://github.com/gocolly/colly)
* [gopsutils](https://github.com/shirou/gopsutil)
*  等等...
### 前端
&emsp;&emsp;本项目前端采用vue.js,具体框架采用[vue-element-admin](https://github.com/PanJiaChen/vue-element-admin),对原作者[PanJiaChen](https://github.com/PanJiaChen)表示感谢.  

&emsp;&emsp;前端计划采用[buefy](https://github.com/buefy/buefy)或原生[bulma](https://github.com/jgthms/bulma)进行调整优化


## 已完成

## 待完善
* [x]  minikube集群部署(cert-manager 证书管理)
* [ ]  利用[ttyd](https://github.com/tsl0922/ttyd)提供容器内终端访问
* [x] 基于jwt-go的jwt 认证
* [ ] 基于RBAC的用户管理
    * [ ] 采用[casbin](https://casbin.org/) 进行混合rbac和restful权限管理
* [x] 基于Makefile和Dockerfile的服务构建
    * [x] 编译golang项目时打包版本信息[采用build.sh 脚本生成api版本信息]
* [ ] web 框架调整和优化
    * [x] 基于`mux`的微服务框架
    * [x] 日志功能
    * [ ] Error错误处理
    * [x] 采用中间件模式添加服务[prometheus](https://prometheus.io/) 指标
    * [x] web配置文件解析(`yaml`配置文件的解析)
## TODO
* [x] golang 配置文件解析
* [x] Linux 监控服务API开发
* [ ] 前端开发
* [ ] 用户管理与RBAC
* [ ] kubernetes集群管理
* [ ] kubernetes YAML应用部署和访问
* [ ] 熔断器和限流器   
    *[x] [限流器](https://pkg.go.dev/golang.org/x/time/rate) 已完成
* [ ] golang 代理服务
* [ ] ssh /tcp 代理，多层代理
* [ ] 基于NSQ的消息订阅发布模式
* [ ] 采用[go-guard](https://github.com/shaj13/go-guardian) 统一认证
 

项目布局参考:  
https://github.com/golang-standards/project-layout

* [x] 去除src目录
* [ ]  使用expvar包进行服务监控


## 参考资料
[1].http://www.hatlonely.com/2018/06/21/%E5%BE%AE%E6%9C%8D%E5%8A%A1%E7%BB%84%E4%BB%B6%E4%B9%8B%E9%99%90%E6%B5%81%E5%99%A8%E4%B8%8E%E7%86%94%E6%96%AD%E5%99%A8/
