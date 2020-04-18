# registry-auth-server

### 获取
```bash
$git clone https://github.com/3Xpl0it3r/registry-auth-server.git
```

### QuickStart
```bash
$cd registry-auth-server
$cd go build .
$ ./server
```


### 辅助功能
>> --ca 参数表明创建一非CA证书， 不指定则创建ca证书
>> --domain 给证书绑定域名
>> --ip ip1 ip2 ip3  给证书绑定IP地址

```bash
// 创建自签证书
registry  openssl generate -o <证书存放路径>
// 创建EndOfUser证书
registry openssl generate -o <证书存放路径> --ca <ca_path> [--ip 192.138.1.1 192.158.1.2] [--domain app1.example.com]
```