# perch
Golang Code And Others
## 项目内容说明

### 后端
本项目的后端采用Golang语言开发，涉及到的工具包有
* [jwt-go](https://github.com/dgrijalva/jwt-go)
*  [gopsutils](https://github.com/shirou/gopsutil)
* [gorm ](https://github.com/go-gorm/gorm)
* [client-go](https://github.com/kubernetes/client-go)
### 前端
本项目前端采用vue.js,具体框架采用[vue-element-admin](https://github.com/PanJiaChen/vue-element-admin),对原作者[PanJiaChen](https://github.com/PanJiaChen)表示感谢.  
前端涉及到的数据包有  
* [datav](http://datav.jiaminghi.com/guide/) 进行数据可视化

## 已完成

## 待完善
* [ ] 基于jwt-go的jwt 认证
* [ ] 基于RBAC的用户管理
* [ ] 基于Makefile 和docker-compose的环境部署
* [ ] web 框架调整和优化
    * [ ] 日志功能
    * [ ] Error错误处理
    * [x] web配置文件解析
      * go-yaml进行配置文件的解析
## TODO
* [ ] golang 配置文件解析
* [ ] Linux 监控服务API开发
* [ ] 前端开发
* [ ] 用户管理与RBAC
* [ ] kubernetes集群管理
* [ ] kubernetes YAML应用部署和访问
* [ ] golang 代理服务
 *[ ] ssh /tcp 代理，多层代理

项目布局参考:  
https://github.com/golang-standards/project-layout

* [ ] 去除src目录
* [ ]  使用expvar包进行服务监控
