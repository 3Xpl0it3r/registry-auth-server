package openssl

import "github.com/spf13/cobra"

func NewOpensslCommand()*cobra.Command{
	cmd := &cobra.Command{
		Use: "openssl",
		Short: "openssl create cert and key",
	}
	cmd.AddCommand(newOpensslGenerateCommand())
	cmd.SetHelpFunc(cmd.HelpFunc())
	return cmd
}
