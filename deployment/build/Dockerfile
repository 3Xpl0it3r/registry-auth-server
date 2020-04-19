FROM alpine

ARG BUILD_PATH=/tmp/build
ARG INSTALL_PATH=/opt
ARG PROJECT_NAME=registry-auth-server
ARG CERT_PATH=/media/cert/
RUN apk update && \
    apk add git go && \
    mkdir  $BUILD_PATH -pv && cd $BUILD_PATH && \
    git clone https://github.com/3Xpl0it3r/registry-auth-server.git && \
    cd $PROJECT_NAME && \
    go build . && chmod +x $PROJECT_NAME && \
    mkdir $INSTALL_PATH/$PROJECT_NAME/configs -pv && mv $PROJECT_NAME $INSTALL_PATH/$PROJECT_NAME/  \
    cp configs/config.yaml $INSTALL_PATH/$PROJECT_NAME/configs/  \
    # 准备创建证书
    cd cli && go build . && chmod +x cli && \
    ./cli openssl generate -o $CERT_PATH && ./cli openssl generate -o $CERT_PATH --ca $CERT_PATH && \
    rm -rf $BUILD_PATH &&  \
    apk remove git go && apk clean && rm -rf /var/lib/apk/*

WORKDIR $INSTALL_PATH:$PROJECT_NAME

# ENTRYPOINT ["./$PROJECT_NAME"]
ENTRYPOINT ["/sbin/init"]

