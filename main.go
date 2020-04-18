package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"l0calh0st.cn/registry-auth-server/configs"
	server2 "l0calh0st.cn/registry-auth-server/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := server2.NewRegistryAuthServer(configs.NewConfigs("config"))
	if err := server.Run(ctx); err != nil {
		logrus.Error(err.Error())
	}
}
