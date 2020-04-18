package api

type IToken interface {
	GenerateToken(config *ScopeRequest)(string,error)
}



type TokenConfig struct {
	Issuer string
	CertFile string
	KeyFile string
	Expiration int64
	Scope ScopeRequest
}


type ScopeRequest struct {
	Type string
	Name string
	Actions []string
}


