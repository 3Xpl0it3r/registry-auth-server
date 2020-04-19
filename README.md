# registry-auth-server




### 下载安装
> 1. 下载源码包
```bash
[root@k3s ~]# git clone https://github.com/3Xpl0it3r/registry-auth-server.git
Cloning into 'registry-auth-server'...
remote: Enumerating objects: 85, done.
remote: Counting objects: 100% (85/85), done.
remote: Compressing objects: 100% (64/64), done.
remote: Total 85 (delta 19), reused 78 (delta 16), pack-reused 0
Unpacking objects: 100% (85/85), done.
[root@k3s ~]# 
```

### 快速安装
```bash
./sdeployment/run.sh
```
> 2. 配置go proxy
```bash
[root@k3s ~]# export GOPROXY=https://goproxy.cn
[root@k3s ~]# 
```
> 3. 编译token server 和 cli命令行(非必须)
```bash
[root@k3s ~]# cd registry-auth-server/
[root@k3s registry-auth-server]# go build l0calh0st.cn/registry-auth-server
go: downloading github.com/sirupsen/logrus v1.5.0
go: downloading github.com/spf13/viper v1.6.3
.......
[root@k3s registry-auth-server]# ls
api  cli  configs  deployment  go.mod  go.sum  LICENSE  main.go  pkg  README.md  registry-auth-server  server
[root@k3s registry-auth-server]# cd cli/
[root@k3s cli]# go build .
[root@k3s cli]# ls
cli  openssl  registry.go  server
[root@k3s cli]# 
```
>> `cli命令行工具提供serve命令，用于运行token auth server  和一个openssl 命令，openssl命令用于创建一个简易证书`
>
>4. 设置环境变量
```bash
[root@k3s cli]# cd ..
[root@k3s registry-auth-server]# export PATH=$PATH:`pwd`:`pwd`/cli
[root@k3s registry-auth-server]# cli
Usage:
  registry [command]

Available Commands:
  help        Help about any command
  openssl     openssl create cert and key
  server      run registry token auth server

Flags:
  -h, --help   help for registry

Use "registry [command] --help" for more information about a command.
```

### 修改配置文件

### QuickStart(非证书通信)
*docker daemon默认以https方式和registry server通信，所以以非证书模式需要修改docker daemon配置文件，以http方式发送请求*
> 1. 运行 token-auth 服务
```bash

```
> 2. 启动registry 容器
> 3. 修改docker 配置文件
> 4. 重启docker

### QuickStart(默认证书通信)
> 1. 创建证书
> 2. 修改token server配置文件


### QuickStart (自定义证书通信)



### 配置文件


### 证书创建工具
`证书创建工具提供一种快速创建证书的途径`
>> --ca 参数表明创建一非CA证书， 不指定则创建ca证书
>> --domain 给证书绑定域名
>> --ip ip1 ip2 ip3  给证书绑定IP地址

```bash
// 创建自签证书
registry  openssl generate -o <证书存放路径>
// 创建EndOfUser证书
registry openssl generate -o <证书存放路径> --ca <ca_path> [--ip 192.138.1.1 192.158.1.2] [--domain app1.example.com]
```