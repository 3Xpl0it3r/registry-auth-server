package server

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"l0calh0st.cn/registry-auth-server/configs"
	server2 "l0calh0st.cn/registry-auth-server/server"
)

var configPath string

func NewTokenServerCommand()*cobra.Command{
	cmd := &cobra.Command{
		Use: "server",
		Short: "run registry token auth server",
		PreRun: func(cmd *cobra.Command, args []string) {
			var err error
			configPath,err = cmd.Flags().GetString("config")
			if err!= nil || configPath == ""{
				logrus.WithField("Stage", "Load Server Config").Infoln("Config is not specialfied, will use default config file")
				return
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			cfg := configs.NewConfigs(configPath)
			fmt.Println(cfg)
			server := server2.NewRegistryAuthServer(cfg)
			if server == nil{
				return
			}
			ctx,cancel := context.WithCancel(context.Background())
			defer cancel()
			if err := server.Run(ctx);err != nil{
				logrus.Panicf("RunServerFailed: %s\n",err.Error())
			}
			<-ctx.Done()
		},
	}

	initFlags(cmd)
	return cmd
}


func initFlags(cmd *cobra.Command){
	cmd.Flags().StringVarP(&configPath, "config", "c", "", "--config/-c")
}


