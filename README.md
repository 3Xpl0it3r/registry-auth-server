# registry-auth-server

[toc]


### 0x00 快速安装
1. **下载源码包**
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
2. **快速部署**
```bash
[root@k3s registry-auth-server]# ./deployment/run.sh 192.168.1.23
INFO[0000] Generate Ca Cert and Key Successfully        
INFO[0000] Save Ca cert:/media/cert/ca.crt	Key:/media/cert/ca.key	Successfully 
INFO[0000] Save Ca cert:/media/cert/server.crt	Key:/media/cert/server.key	Successfully 
dcc2e0a47c80c81cd441282b1fbf4e1fbb74a202c4db7fc6c11c3ac33870a7db
[root@k3s registry-auth-server]# ./registry-auth-server 
INFO[0000] Load Authenticator Controller Successfully   
INFO[0000] Load Authorization Controller Successfully   
INFO[0000] Load token Controller Successfully           
INFO[0000] Docker registry token server begin running..... 
INFO[0000] Docker Registry Auth server Run as TLS Module 
```

### 0x01 docker 证书问题
> ·默认情况下docker是通过https的方式来登陆验证的，因此我们需要做docker 做一些配置， 通过又两种方式(任选其一)·
> 1. 忽略证书
```bash
# 编辑docker daemon.json配置文件
[root@bogon ~]# vim /etc/docker/daemon.json
[root@bogon ~]# cat /etc/docker/daemon.json 
{
"insecure-registries":["example.com:5000"]
}
# 由于证书设置了域名，因此需要把域名解析填写到hosts文件里面
[root@bogon ~]# cat /etc/hosts
127.0.0.1   localhost localhost.localdomain localhost4 localhost4.localdomain4
::1         localhost localhost.localdomain localhost6 localhost6.localdomain6
192.168.1.23 reg.example.com example.com
[root@bogon ~]# 
[root@bogon ~]# docker login example.com:5000
Username: root
Password: 
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded

```
> 2. 信任证书
```bash
#
[root@bogon ~]# cat /etc/docker/daemon.json 
{
}
[root@bogon ~]# docker logout example.com:5000
Not logged in to example.com:5000
[root@bogon ~]# docker login example.com:5000
Username: root
Password: 
Error response from daemon: Get https://example.com:5000/v2/: x509: certificate signed by unknown authority
[root@bogon ~]# 

# 将ca证书拷贝到/etc/docker/certs.d/<域名/Ip地址>/
[root@bogon ~]# scp root@192.168.1.23:/media/cert/ca.crt /etc/docker/certs.d/example.com\:5000/
The authenticity of host '192.168.1.23 (192.168.1.23)' can't be established.
ECDSA key fingerprint is SHA256:64fbUNTDF+Yp4gjAD2Zst19CULJoAri2RLVpvQCh4w4.
ECDSA key fingerprint is MD5:8d:c6:d3:ff:fb:3f:34:c4:5f:3d:61:9d:1f:f0:75:37.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added '192.168.1.23' (ECDSA) to the list of known hosts.
root@192.168.1.23's password: 
ca.crt                                                  
[root@bogon ~]# docker login example.com:5000
Username: root
Password: 
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
[root@bogon ~]# 

```