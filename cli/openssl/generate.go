package openssl

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"l0calh0st.cn/registry-auth-server/pkg/openssl"
	"os"
	"path"
	"reflect"
)

var (
	OutPath    string
	CaPath     string
	DomainList []string
	IpSan      []string
)

func newOpensslGenerateCommand() *cobra.Command {
	var cert openssl.AbstractOpenssl

	var certName, keyName string
	cmd := &cobra.Command{
		Use: "generate",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			OutPath, err = cmd.Flags().GetString("out")
			if err != nil || OutPath == "" {
				logrus.Info("[!]Save Path is not specialfied, Default /tmp/")
				OutPath = "/tmp"
			}
			CaPath, err := cmd.Flags().GetString("ca")
			if CaPath == "" || err != nil {
				cert = openssl.NewSimpleRootCa(openssl.NewDefaultSimpleCertConfig())
				certName = "ca.pem"
				keyName = "ca.key"
			} else {
				certName = "server.pem"
				keyName = "server.key"
				simpleCertConfig := openssl.NewDefaultSimpleCertConfig()
				simpleCertConfig.DNSName = DomainList
				simpleCertConfig.IPAddress = IpSan
				cert = openssl.NewSimpleEndOfUserCert(simpleCertConfig, CaPath)
			}

			var isCheckedFalse bool
			certReflect := reflect.ValueOf(cert)
			if certReflect.Kind() == reflect.Ptr {
				isCheckedFalse = certReflect.IsNil()
			}
			if isCheckedFalse {
				logrus.Error("Build Concrete Cert Failed, Concrete CertBuilder is Nil\n")
				return errors.New("cert build is none")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			err := cert.Generate()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Generate  Cert failed: %s\n", err)
				return
			}
			err = cert.SaveToFile(path.Join(OutPath, certName), path.Join(OutPath, keyName))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Save Cert And Key to %s  Failed: %s\n", OutPath, err.Error())
				return
			}

		},
	}
	initCaFlags(cmd)
	return cmd
}

func initCaFlags(command *cobra.Command) {
	command.Flags().StringVarP(&OutPath, "out", "o", "", "the path to save  cert and  key")
	command.Flags().StringVarP(&CaPath, "ca", "c", "", "if special capath,create a endofuser cert \n"+
		"is not special capath, it will create a root ca ")
	command.Flags().StringArrayVar(&DomainList, "domain", []string{"example"}, "--domain")
	command.Flags().StringArrayVar(&IpSan, "ip", []string{"127.0.0.1"}, "--ip")
}
