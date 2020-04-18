package main

import (
	"github.com/spf13/cobra"
	"l0calh0st.cn/registry-auth-server/cli/openssl"
	"l0calh0st.cn/registry-auth-server/cli/server"
)

func main() {
	app := newRootCmd()
	app.Execute()
}


func newRootCmd()*cobra.Command{
	cmd := &cobra.Command{
		Use: "registry",
	}
	cmd.AddCommand(openssl.NewOpensslCommand())
	cmd.AddCommand(server.NewTokenServerCommand())
	return cmd
}