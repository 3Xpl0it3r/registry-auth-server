package api

type IToken interface {
	GenerateToken(claim *TokenClaim) (string, error)
}

type TokenConfig struct {
	Issuer     string
	CertFile   string
	KeyFile    string
	Expiration int64
	Claim      TokenClaim
}

type TokenClaim struct {
	Type    string // scope
	Account string
	Name    string //scope
	Service string //
	Actions []string
}
