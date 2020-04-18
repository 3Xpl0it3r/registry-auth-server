package server

import (
	"l0calh0st.cn/registry-auth-server/api"
	"l0calh0st.cn/registry-auth-server/configs"
)

func tokenConfigHandler(config *configs.Configs) *api.TokenConfig {
	return &api.TokenConfig{
		Issuer:     config.Token.Issuer,
		CertFile:   config.Tls.Cert,
		KeyFile:    config.Tls.Key,
		Expiration: config.Token.Expiration,
		Claim:      api.TokenClaim{},
	}
}
