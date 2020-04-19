#!/usr/bin/bash

if [ $# != 1 ];
then
    echo "Usage: run.sh  <ipaddress>"
    echo "IP是registry 连接token server的地址"
    exit
fi

# 证书存放目录
CERT_PATH=/media/cert
# 证书绑定IP地址
IP=127.0.0.1
# 证书绑定域名
DOMAIN=example.com

ROOT_PATH=`pwd`

IP=$1

# 编译auth-tokenserver
cd $ROOT_PATH
export GOPROXY=https://goproxy.cn
go build l0calh0st.cn/registry-auth-server
cd cli  && go build .
# 安装证书
mkdir $CERT_PATH -pv
./cli openssl generate -o $CERT_PATH
./cli openssl generate -o $CERT_PATH --ca $CERT_PATH --ip $IP --domain $DOMAIN

cd ..
# 启动registry
docker run -d    -p 5000:5000 --name registry -v /media/cert:/cert \
    -e REGISTRY_AUTH=token  -e REGISTRY_AUTH_TOKEN_REALM=https://$IP:5050/auth  \
    -e REGISTRY_AUTH_TOKEN_SERVICE="registry.docker.io"  \
    -e REGISTRY_AUTH_TOKEN_ISSUER="distribution-token-server" \
    -e REGISTRY_STORAGE_DELETE_ENABLED=true \
    -e REGISTRY_AUTH_TOKEN_ROOTCERTBUNDLE=/cert/server.crt \
    -e REGISTRY_HTTP_TLS_CERTIFICATE=/cert/server.crt  \
    -e REGISTRY_HTTP_TLS_KEY=/cert/server.key registry

# 启动token server
#cd $PREFIX
#nohup ./registry-token-server &
