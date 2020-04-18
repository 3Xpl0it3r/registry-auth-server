package main

import (
	"context"
	"l0calh0st.cn/registry-auth-server/configs"
	server2 "l0calh0st.cn/registry-auth-server/server"
	"log"
)

func main() {
	ctx,cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := configs.NewConfigs("config")
}
