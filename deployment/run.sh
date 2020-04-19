#!/usr/bin/bash

# 安装目录
PREFIX=/opt/token-server
# 证书存放目录
CERT_PATH=/media/cert
# 证书绑定IP地址
IP=127.0.0.1
# 证书绑定域名
DOMAIN=example.com



ROOT_PATH=`pwd`

# 编译auth-tokenserver
cd $ROOT_PATH
export GOPROXY=https://goproxy.cn
go build l0calh0st.cn/registry-auth-server
mv registry-auth-server $PREFIX
mkdir $PREFIX/configs/
cp -rf configs/config.yaml $PREFIX/configs/
cd cli  && go build .
# 安装证书
mkdir $CERT_PATH -pv
./cli openssl generate -o $CERT_PATH
./cli openssl generate -o $CERT_PATH --ca $CERT_PATH --IP $IP --DOMAIN $DOMAIN

# 启动registry
docker run -d  --rm  -p 5000:5000 --name registry -v /media/cert:/cert \
    -e REGISTRY_AUTH=token  -e REGISTRY_AUTH_TOKEN_REALM=https://127.0.0.1:5050/auth  \
    -e REGISTRY_AUTH_TOKEN_SERVICE="registry.docker.io"  \
    -e REGISTRY_AUTH_TOKEN_ISSUER="distribution-token-server" \
    -e REGISTRY_STORAGE_DELETE_ENABLED=true \
    -e REGISTRY_AUTH_TOKEN_ROOTCERTBUNDLE=/cert/server.crt \
    -e REGISTRY_HTTP_TLS_CERTIFICATE=/cert/server.crt  \
    -e REGISTRY_HTTP_TLS_KEY=/cert/server.key registry

# 启动token server
cd $PREFIX
nohup ./registry-token-server &
